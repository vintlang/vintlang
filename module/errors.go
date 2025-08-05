package module

import (
	"github.com/vintlang/vintlang/object"
)

var ErrorFunctions = map[string]object.ModuleFunction{}

func init() {
	ErrorFunctions["new"] = errorsNew
}

func errorsNew(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return ErrorMessage(
			"errors",
			"new",
			"1 string argument",
			"keyword arguments provided",
			`errors.new("Something went wrong") -> error`,
		)
	}
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"errors",
			"new",
			"1 string argument",
			formatArgs(args),
			`errors.new("Something went wrong") -> error`,
		)
	}
	msg := args[0].(*object.String).Value
	return &object.Error{Message: msg}
}
