package module

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/vintlang/vintlang/object"
)

var ReflectFunctions = map[string]object.ModuleFunction{}

func init() {
	ReflectFunctions["typeof"] = reflectTypeof
	ReflectFunctions["inspect"] = reflectInspect
	ReflectFunctions["isType"] = reflectIsType
	ReflectFunctions["size"] = reflectSize
	ReflectFunctions["keys"] = reflectKeys
	ReflectFunctions["hasMethod"] = reflectHasMethod
}

// reflectTypeof returns the type of an object as a string
func reflectTypeof(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return ErrorMessage(
			"reflect", "typeof",
			"1 argument",
			"keyword arguments provided",
			`reflect.typeof(42) -> "INTEGER"`,
		)
	}
	if len(args) != 1 {
		return ErrorMessage(
			"reflect", "typeof",
			"1 argument",
			FormatArgs(args),
			`reflect.typeof(42) -> "INTEGER"`,
		)
	}
	
	return &object.String{Value: string(args[0].Type())}
}

// reflectInspect returns the inspection string of an object
func reflectInspect(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return ErrorMessage(
			"reflect", "inspect",
			"1 argument",
			"keyword arguments provided",
			`reflect.inspect([1, 2, 3]) -> "[1, 2, 3]"`,
		)
	}
	if len(args) != 1 {
		return ErrorMessage(
			"reflect", "inspect",
			"1 argument",
			FormatArgs(args),
			`reflect.inspect([1, 2, 3]) -> "[1, 2, 3]"`,
		)
	}
	
	return &object.String{Value: args[0].Inspect()}
}

// reflectIsType checks if an object is of a specific type
func reflectIsType(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return ErrorMessage(
			"reflect", "isType",
			"2 arguments (object, type_string)",
			"keyword arguments provided",
			`reflect.isType(42, "INTEGER") -> true`,
		)
	}
	if len(args) != 2 {
		return ErrorMessage(
			"reflect", "isType",
			"2 arguments (object, type_string)",
			FormatArgs(args),
			`reflect.isType(42, "INTEGER") -> true`,
		)
	}
	if args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"reflect", "isType",
			"2 arguments (object, type_string)",
			fmt.Sprintf("object, %s", args[1].Type()),
			`reflect.isType(42, "INTEGER") -> true`,
		)
	}
	
	expectedType := args[1].(*object.String).Value
	actualType := string(args[0].Type())
	
	return &object.Boolean{Value: actualType == expectedType}
}

// reflectSize returns the size/length of collections
func reflectSize(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return ErrorMessage(
			"reflect", "size",
			"1 argument",
			"keyword arguments provided",
			`reflect.size([1, 2, 3]) -> 3`,
		)
	}
	if len(args) != 1 {
		return ErrorMessage(
			"reflect", "size",
			"1 argument",
			FormatArgs(args),
			`reflect.size([1, 2, 3]) -> 3`,
		)
	}
	
	obj := args[0]
	switch o := obj.(type) {
	case *object.Array:
		return &object.Integer{Value: int64(len(o.Elements))}
	case *object.Dict:
		return &object.Integer{Value: int64(len(o.Pairs))}
	case *object.String:
		return &object.Integer{Value: int64(len(o.Value))}
	default:
		return ErrorMessage(
			"reflect", "size",
			"sizeable object (array, dict, string)",
			string(obj.Type()),
			`reflect.size([1, 2, 3]) -> 3`,
		)
	}
}

// reflectKeys returns keys of dict-like objects
func reflectKeys(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return ErrorMessage(
			"reflect", "keys",
			"1 argument",
			"keyword arguments provided",
			`reflect.keys({"a": 1, "b": 2}) -> ["a", "b"]`,
		)
	}
	if len(args) != 1 {
		return ErrorMessage(
			"reflect", "keys",
			"1 argument",
			FormatArgs(args),
			`reflect.keys({"a": 1, "b": 2}) -> ["a", "b"]`,
		)
	}
	
	obj := args[0]
	switch o := obj.(type) {
	case *object.Dict:
		keys := make([]object.Object, 0, len(o.Pairs))
		for _, pair := range o.Pairs {
			keys = append(keys, pair.Key)
		}
		return &object.Array{Elements: keys}
	default:
		return ErrorMessage(
			"reflect", "keys",
			"dict object",
			string(obj.Type()),
			`reflect.keys({"a": 1, "b": 2}) -> ["a", "b"]`,
		)
	}
}

// reflectHasMethod checks if an object has a specific method
func reflectHasMethod(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return ErrorMessage(
			"reflect", "hasMethod",
			"2 arguments (object, method_name)",
			"keyword arguments provided",
			`reflect.hasMethod([1, 2, 3], "push") -> true`,
		)
	}
	if len(args) != 2 {
		return ErrorMessage(
			"reflect", "hasMethod",
			"2 arguments (object, method_name)",
			FormatArgs(args),
			`reflect.hasMethod([1, 2, 3], "push") -> true`,
		)
	}
	if args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"reflect", "hasMethod",
			"2 arguments (object, method_name)",
			fmt.Sprintf("object, %s", args[1].Type()),
			`reflect.hasMethod([1, 2, 3], "push") -> true`,
		)
	}
	
	obj := args[0]
	methodName := args[1].(*object.String).Value
	
	// Use Go reflection to check if the object has a Method function
	objValue := reflect.ValueOf(obj)
	methodFunc := objValue.MethodByName("Method")
	
	if !methodFunc.IsValid() {
		// Object doesn't have any methods
		return &object.Boolean{Value: false}
	}
	
	// Try to call the method and see if it returns an error indicating the method doesn't exist
	// We'll call with empty args to test if the method exists
	methodArgs := []reflect.Value{
		reflect.ValueOf(methodName),
		reflect.ValueOf([]object.Object{}),
	}
	
	result := methodFunc.Call(methodArgs)
	if len(result) > 0 {
		if errorObj, ok := result[0].Interface().(*object.Error); ok {
			// Check if the error message indicates method not found
			errorMsg := strings.ToLower(errorObj.Message)
			if strings.Contains(errorMsg, "not supported") || 
			   strings.Contains(errorMsg, "unknown method") ||
			   strings.Contains(errorMsg, "method") && strings.Contains(errorMsg, methodName) {
				return &object.Boolean{Value: false}
			}
		}
	}
	
	return &object.Boolean{Value: true}
}