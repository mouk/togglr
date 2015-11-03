package togglr

import (
	"reflect"
)

//Populate the passed object with feature using the already
//specified ConfigSources and FeatureFactories
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

func createField(v reflect.Value, field reflect.StructField) {
	if featur, ok := createFeature(field); ok {
		v.Set(reflect.ValueOf(featur))
		return
	}
}

type staticValueFeature struct {
	Value bool
}

//Create a new Feature with a static value that cannot be changed afterwards
func NewStaticFeature(val bool) Feature {
	return &staticValueFeature{val}
}

func (d *staticValueFeature) IsEnabled() bool {
	return d.Value
}
