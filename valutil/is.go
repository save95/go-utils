package valutil

import (
	"reflect"
)

func IsNil(v interface{}) bool {
	vi := reflect.ValueOf(v)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
