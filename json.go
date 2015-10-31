package togglr

import (
	"reflect"
)

var source ConfigSource
var factories = make([]FeatureFactory, 0, 2)

func Init(path string) {
	source = NewJsonConfigSource(path)
}

func createFeature(field reflect.StructField) (Feature, bool) {
	name := field.Name
	tag := field.Tag
	f := append(factories, DefaulFactory)
	for _, factory := range f {
		if val, ok := getConfig(name, tag); ok {
			if feature, ok := factory(val); ok {
				return feature, ok
			}
		}
	}
	return nil, false
}

func getConfig(name string, tag reflect.StructTag) (interface{}, bool) {
	if val, ok := singltonEnvSource.GetConfig(name, tag); ok {
		return val, ok
	}
	if val, ok := source.GetConfig(name, tag); ok {
		return val, ok
	}
	return nil, false
}
