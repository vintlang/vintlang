package module

import (
	"github.com/vintlang/vintlang/config"
	"github.com/vintlang/vintlang/object"
)

var ErrorsModule = object.NewModule("errors", nil)

func init() {
	//Note: Sample usage of the new module definition
	ErrorsModule.RegisterFunction("new", errorsNew)
	// ErrorsModule.RegisterVariable("defaultCode", &object.Integer{Value: 1000})
	ErrorsModule.Doc = "Provides error creation and handling utilities."
	ErrorsModule.Version = config.VINT_VERSION
	ErrorsModule.Author = "VintLang Core"
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
