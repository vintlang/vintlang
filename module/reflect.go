package module

import (
	"github.com/vintlang/vintlang/object"
)

var ReflectFunctions = map[string]object.ModuleFunction{}

func init() {
	ReflectFunctions["typeOf"] = typeOf
	ReflectFunctions["valueOf"] = valueOf
	ReflectFunctions["isNil"] = isNil
	ReflectFunctions["isArray"] = isArray
	ReflectFunctions["isObject"] = isObject
	ReflectFunctions["isFunction"] = isFunction
}

// typeOf returns the type name of the given value
func typeOf(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"reflect",
			"typeOf",
			"1 argument",
			formatArgs(args),
			`reflect.typeOf("Hello") -> "STRING"`,
		)
	}
	return &object.String{Value: string(args[0].Type())}
}

// valueOf returns the raw value (already implemented)
func valueOf(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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

// isNil checks if the value is null
func isNil(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"reflect",
			"isNil",
			"1 argument",
			formatArgs(args),
			`reflect.isNil(null) -> true`,
		)
	}
	_, ok := args[0].(*object.Null)
	return &object.Boolean{Value: ok}
}

// isArray checks if the value is an array
func isArray(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"reflect",
			"isArray",
			"1 argument",
			formatArgs(args),
			`reflect.isArray([1,2,3]) -> true`,
		)
	}
	_, ok := args[0].(*object.Array)
	return &object.Boolean{Value: ok}
}

// isObject checks if the value is a dictionary/object
func isObject(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"reflect",
			"isObject",
			"1 argument",
			formatArgs(args),
			`reflect.isObject({"a":1}) -> true`,
		)
	}
	_, ok := args[0].(*object.Dict)
	return &object.Boolean{Value: ok}
}

// isFunction checks if the value is a function
func isFunction(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"reflect",
			"isFunction",
			"1 argument",
			formatArgs(args),
			`reflect.isFunction(func() {}) -> true`,
		)
	}
	_, ok := args[0].(*object.Function)
	return &object.Boolean{Value: ok}
}
