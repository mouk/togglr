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
		t := typ.Field(i).Type
		if t.Kind() == reflect.Interface && t.Implements(fType) {
			fieldValue := val.Field(i)
			fieldName := typ.Field(i)
			createField(fieldValue, fieldName)

		}
	}
}

func createFeatureFromEnv(field reflect.StructField) (Feature, bool) {
	var varName = field.Tag.Get("env")
	if varName == "" {
		return nil, false
	}
	envVariable := os.Getenv(varName) == "1"
	return staticValueFeature{envVariable}, true
}
func createField(v reflect.Value, field reflect.StructField) {
	if featur, ok := createFeatureFromEnv(field); ok {
		v.Set(reflect.ValueOf(featur))
		return
	}

	if featur, ok := createFeatureFromJson(field); ok {
		v.Set(reflect.ValueOf(featur))
		return
	}
	v.Set(reflect.ValueOf(AlwaysDisabledFeature))
}

type staticValueFeature struct {
	Value bool
}

func (d staticValueFeature) IsEnabled() bool {
	return d.Value
}

func CreateSimpleValueToggle(enabled bool) Feature {
	if enabled {
		return AlwaysEnabledFeature{}
	}
	return AlwaysDisabledFeature
}

var AlwaysDisabledFeature = staticValueFeature{false}
