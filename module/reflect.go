package module

import "github.com/vintlang/vintlang/object"

var ReflectFunctions = map[string]object.ModuleFunction{}

func init(){
	ReflectFunctions["typeOf"] = typeOf
	ReflectFunctions["valueOf"] = valueOf
	ReflectFunctions["isNil"] = isNil
	ReflectFunctions["isArray"] = isArray
	ReflectFunctions["isObject"] = isObject
	ReflectFunctions["isFunction"] = isFunction
}

func 