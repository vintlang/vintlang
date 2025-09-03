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

func typeOf(args []object.Object, defs map[string]object.Object) object.Object {
	//
}

func valueOf(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"reflect",
			"valueOf",
			"1 argument",
			formatArgs(args),
			`reflect.valueOf("Hello") -> "Hello"`,
		)
	}
	return args[0]
}
