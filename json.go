package togglr

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
)

var jsonData map[string]interface{}

func Init(path string) error {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}
	jsonData = make(map[string]interface{})
	json.Unmarshal(file, &jsonData)
	return nil
}

func createFeatureFromJson(field reflect.StructField) (Feature, bool) {

	key := field.Tag.Get("json")
	if key == "" {
		key = field.Name
	}
	//todo detect the type
	if val, ok := jsonData[key]; ok {
		if val == true {
			return staticValueFeature{true}, true
		} else if val == false {
			return staticValueFeature{false}, true
		}
		prop := val.(map[string]interface{})
		if feature, ok := createEnabledOnOrAfter(prop); ok {
			return feature, ok
		}
		return staticValueFeature{val == true}, true
	}
	return nil, false
}
