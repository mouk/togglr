package togglr

import (
	"os"
	"reflect"
)

func Read(obj interface{}) {
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}
	fType := reflect.TypeOf((*Feature)(nil)).Elem()

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Type.Implements(fType) {
			fieldValue := val.Field(i)
			fieldName := typ.Field(i).Name
			createField(fieldValue, fieldName)

		}
	}
}

func createField(v reflect.Value, name string) {
	envVariable := os.Getenv(name) == "1"
	v.Set(reflect.ValueOf(staticValueFeatur{envVariable}))
}

type staticValueFeatur struct {
	Value bool
}

func (d staticValueFeatur) IsEnabled() bool {
	return d.Value
}
