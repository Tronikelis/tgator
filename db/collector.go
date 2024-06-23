package db

import (
	"fmt"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type StructFieldMap map[string]reflect.Value

func initNilPointer(val reflect.Value) {
	if val.Kind() != reflect.Pointer || !val.IsNil() {
		return
	}

	val.Set(reflect.New(val.Type().Elem()))
}

func structToMap(v any) (StructFieldMap, error) {
	mp := StructFieldMap{}

	val := reflect.ValueOf(v).Elem()
	typ := reflect.TypeOf(v).Elem()

	if val.Kind() == reflect.Pointer && val.IsNil() {
		initNilPointer(val)

		val = val.Elem()
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return StructFieldMap{}, fmt.Errorf("wanted struct, got %v", typ.Kind().String())
	}

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldTyp := typ.Field(i)

		name := fieldTyp.Tag.Get("db")
		if name == "" {
			name = fieldTyp.Name
		} else if name == "-" {
			continue
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
			// the struct does not have a corresponding field
			var trash any
			pointers = append(pointers, trash)
			continue
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
			currentId = tableId

			split = append(split, description)
			description = []pgconn.FieldDescription{}
		}

		description = append(description, d)
	}

	split = append(split, description)

	return split
}

func nestedStructsPtrs(strct any) []any {
	val := reflect.ValueOf(strct).Elem()
	typ := reflect.TypeOf(strct).Elem()

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
		typ = typ.Elem()
	}

	nested := []any{}

	for i := 0; i < val.NumField(); i++ {
		fieldTyp := typ.Field(i)
		fieldVal := val.Field(i)

		if _, exists := fieldTyp.Tag.Lookup("embedded"); !exists {
			continue
		}

		nested = append(nested, fieldVal.Addr().Interface())
	}

	return nested
}

func RowToStruct[T any](row pgx.CollectableRow) (T, error) {
	descriptions := splitDescriptionsByTable(row.FieldDescriptions())

	var t T
	toBeScanned := []any{}

	if _, err := traverseDescriptions(&toBeScanned, descriptions, 0, &t); err != nil {
		return t, err
	}

	if err := row.Scan(toBeScanned...); err != nil {
		return t, err
	}

	return t, nil
}

func traverseDescriptions(
	toBeScanned *[]any,
	descriptions [][]pgconn.FieldDescription,
	index int,
	strct any,
) (int, error) {
	if index >= len(descriptions) {
		return 0, nil
	}

	pointers, err := descriptionsToPointers(descriptions[index], strct)
	if err != nil {
		return 0, err
	}

	*toBeScanned = append(*toBeScanned, pointers...)

	for _, nested := range nestedStructsPtrs(strct) {
		index, err = traverseDescriptions(toBeScanned, descriptions, index+1, nested)
		if err != nil {
			return 0, err
		}
	}

	return index, nil
}
