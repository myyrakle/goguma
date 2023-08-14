package cast

import (
	"reflect"
)

func ToInterfaceSlice[T any](array []T) []interface{} {
	v := reflect.ValueOf(array)
	if v.Kind() != reflect.Slice {
		panic("ToInterfaceSlice only accepts slices")
	}

	result := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = v.Index(i).Interface()
	}

	return result
}
