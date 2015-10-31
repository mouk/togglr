package togglr

import (
	"reflect"
)

func GetFeatureSpanshots(obj interface{}) map[string]bool {
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}
	fType := reflect.TypeOf((*Feature)(nil)).Elem()

	ret := make(map[string]bool)
	for i := 0; i < typ.NumField(); i++ {
		t := typ.Field(i).Type
		feature := val.Field(i).Interface()
		if t.Implements(fType) && feature != nil {
			name := typ.Field(i).Name
			featureValue := feature.(Feature).IsEnabled()
			ret[name] = featureValue
		}
	}
	return ret
}
