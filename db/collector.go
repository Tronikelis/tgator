package db

import (
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type StructFieldMap map[string]reflect.Value

func structToMap(v any) StructFieldMap {
	mp := StructFieldMap{}

	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Struct {
		panic("wanted struct")
	}

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldTyp := typ.Field(i)

		name := fieldTyp.Tag.Get("db")
		if name == "" {
			name = fieldTyp.Name
		}

		mp[name] = fieldVal
	}

	return mp
}

func descriptionsToPointers(ds []pgconn.FieldDescription, strct any) []any {
	structMap := structToMap(strct)

	pointers := []any{}

	for _, d := range ds {
		value, exists := structMap[d.Name]
		if !exists {
			panic("map does not have value")
		}

		pointers = append(pointers, value.Addr().Pointer())
	}

	return pointers
}

func splitDescriptionsByTable(ds []pgconn.FieldDescription) [][]pgconn.FieldDescription {
	descriptions := [][]pgconn.FieldDescription{}

	var currentId uint32
	description := []pgconn.FieldDescription{}

	for _, d := range ds {
		tableId := d.TableOID

		if currentId == 0 {
			currentId = tableId
		} else if currentId != tableId {
			descriptions = append(descriptions, description)
			description = []pgconn.FieldDescription{}
		}

		description = append(description, d)
	}

	return descriptions
}

func RowToStruct[T any](row pgx.CollectableRow) (T, error) {
	descriptions := splitDescriptionsByTable(row.FieldDescriptions())

	var t T

	first := descriptions[0]

	pointers := descriptionsToPointers(first, &t)
	if err := row.Scan(pointers...); err != nil {
		return t, err
	}

	return t, nil
}
