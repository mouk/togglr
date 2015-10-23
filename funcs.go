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
			fieldName := typ.Field(i)
			createField(fieldValue, fieldName)

		}
	}
}

func createField(v reflect.Value, field reflect.StructField) {
	var varName = field.Tag.Get("env")
	if varName == "" {
		varName = field.Name
	}
	envVariable := os.Getenv(varName) == "1"
	v.Set(reflect.ValueOf(staticValueFeatur{envVariable}))
}

type staticValueFeatur struct {
	Value bool
}

func (d staticValueFeatur) IsEnabled() bool {
	return d.Value
}
