package module

import "github.com/vintlang/vintlang/object"

var TypesFunctions = map[string]object.ModuleFunction{}

func init() {
	TypesFunctions["isInt"] = isInt
	TypesFunctions["isFloat"] = isFloat
	TypesFunctions["isString"] = isString
	TypesFunctions["isBool"] = isBool
	TypesFunctions["isArray"] = isArray
	TypesFunctions["isDict"] = isDict
	TypesFunctions["isNull"] = isNull
}

func isInt(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 1 {
		return ErrorMessage(
			"types", "isInt",
			"any object",
			formatArgs(args),
			`types.isInt(42) -> true`,
		)
	}

	return &object.Boolean{Value: args[0].Type() == object.INTEGER_OBJ}
}

func isFloat(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 1 {
		return ErrorMessage(
			"types", "isFloat",
			"any object",
			formatArgs(args),
			`types.isFloat(3.14) -> true`,
		)
	}

	return &object.Boolean{Value: args[0].Type() == object.FLOAT_OBJ}
}

func isString(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 1 {
		return ErrorMessage(
			"types", "isString",
			"any object",
			formatArgs(args),
			`types.isString("hello") -> true`,
		)
	}

	return &object.Boolean{Value: args[0].Type() == object.STRING_OBJ}
}

func isBool(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 1 {
		return ErrorMessage(
			"types", "isBool",
			"any object",
			formatArgs(args),
			`types.isBool(true) -> true`,
		)
	}

	return &object.Boolean{Value: args[0].Type() == object.BOOLEAN_OBJ}
}

func isArray(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 1 {
		return ErrorMessage(
			"types", "isArray",
			"any object",
			formatArgs(args),
			`types.isArray([1, 2, 3]) -> true`,
		)
	}

	return &object.Boolean{Value: args[0].Type() == object.ARRAY_OBJ}
}

func isDict(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 1 {
		return ErrorMessage(
			"types", "isDict",
			"any object",
			formatArgs(args),
			`types.isDict({"key": "value"}) -> true`,
		)
	}

	return &object.Boolean{Value: args[0].Type() == object.DICT_OBJ}
}

func isNull(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 1 {
		return ErrorMessage(
			"types", "isNull",
			"any object",
			formatArgs(args),
			`types.isNull(null) -> true`,
		)
	}

	return &object.Boolean{Value: args[0].Type() == object.NULL_OBJ}
}