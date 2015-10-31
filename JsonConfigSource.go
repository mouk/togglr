package togglr

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
)

type jsonConfigSource struct {
	jsonData map[string]interface{}
}

func NewJsonConfigSource(path string) ConfigSource {
	file, _ := ioutil.ReadFile(path)
	//todo handles errors
	jsonData := make(map[string]interface{})
	json.Unmarshal(file, &jsonData)
	return jsonConfigSource{jsonData}
}

func (source jsonConfigSource) GetConfig(name string, tag reflect.StructTag) (interface{}, bool) {
	key := tag.Get("json")
	if key == "" {
		key = name
	}
	val, ok := source.jsonData[key]

	return val, ok
}
