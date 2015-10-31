package togglr

import (
	"os"
	"reflect"
)

type envSource struct {
}

func (source envSource) GetConfig(name string, tag reflect.StructTag) (interface{}, bool) {
	var varName = tag.Get("env")
	if varName == "" {
		return nil, false
	}
	return os.Getenv(varName) == "1", true
}

var singltonEnvSource = envSource{}
