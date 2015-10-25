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
	json.Unmarshal(file, &jsonData)
	return nil
}

func createFeatureFromJson(field reflect.StructField) (Feature, bool) {

	key := field.Tag.Get("json")
	if key == "" {
		key = field.Name
	}

	if val, ok := jsonData[key]; ok {
		return staticValueFeature{val == true}, true
	}
	return AlwaysDisabledFeature, true
}
