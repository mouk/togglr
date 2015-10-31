package togglr

import "reflect"

//This tinterface represents a pluggable configuration source.
//You can plug in your own sources or use one of the built-in.
type ConfigSource interface {
	GetConfig(name string, tag reflect.StructTag) (interface{}, bool)
}
