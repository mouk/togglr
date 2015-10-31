package togglr

import (
	"reflect"
)

var source ConfigSource

func Init(path string) {
	source = NewJsonConfigSource(path)
}

func createFeatureFromJson(field reflect.StructField) (Feature, bool) {
	name := field.Name
	tag := field.Tag
	//todo detect the type
	if val, ok := source.GetConfig(name, tag); ok {
		if val == true {
			return staticValueFeature{true}, true
		} else if val == false {
			return staticValueFeature{false}, true
		}
		prop := val.(map[string]interface{})
		if feature, ok := createEnabledOnOrAfter(prop); ok {
			return feature, ok
		}
		if feature, ok := createEnabledOnOrBefore(prop); ok {
			return feature, ok
		}
		if feature, ok := createPercentFeature(prop); ok {
			return feature, ok
		}
		if feature, ok := createWeekdaysFeature(prop); ok {
			return feature, ok
		}
	}
	return nil, false
}
