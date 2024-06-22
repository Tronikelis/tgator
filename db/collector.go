package db

import (
	"fmt"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type StructFieldMap map[string]reflect.Value

func structToMap(v any) (StructFieldMap, error) {
	mp := StructFieldMap{}

	val := reflect.ValueOf(v).Elem()
	typ := reflect.TypeOf(v).Elem()

	if typ.Kind() != reflect.Struct {
		return StructFieldMap{}, fmt.Errorf("wanted struct, got %v", typ.Kind().String())
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

	return mp, nil
}

func descriptionsToPointers(ds []pgconn.FieldDescription, strct any) ([]any, error) {
	structMap, err := structToMap(strct)
	if err != nil {
		return []any{}, err
	}

	pointers := []any{}

	for _, d := range ds {
		value, exists := structMap[d.Name]
		if !exists {
			return []any{}, fmt.Errorf("%v does not exist in struct", d.Name)
		}

		pointers = append(pointers, value.Addr().Interface())
	}

	return pointers, nil
}

func splitDescriptionsByTable(ds []pgconn.FieldDescription) [][]pgconn.FieldDescription {
	split := [][]pgconn.FieldDescription{}

	var currentId uint32
	description := []pgconn.FieldDescription{}

	for _, d := range ds {
		tableId := d.TableOID

		if currentId == 0 {
			currentId = tableId
		} else if currentId != tableId {
			split = append(split, description)
			description = []pgconn.FieldDescription{}
		}

		description = append(description, d)
	}

	split = append(split, description)

	return split
}

func RowToStruct[T any](row pgx.CollectableRow) (T, error) {
	descriptions := splitDescriptionsByTable(row.FieldDescriptions())

	var t T

	first := descriptions[0]

	pointers, err := descriptionsToPointers(first, &t)
	if err != nil {
		return t, err
	}

	if err := row.Scan(pointers...); err != nil {
		return t, err
	}

	return t, nil
}
