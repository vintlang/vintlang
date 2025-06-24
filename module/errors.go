package module

import (
	"fmt"

	"github.com/vintlang/vintlang/object"
)

var ErrorFunctions = map[string]object.ModuleFunction{}

func init() {
	ErrorFunctions["new"] = errorsNew
}

func errorsNew(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("errors.new() expects 1 argument, got %d", len(args))}
	}
	if len(defs) != 0 {
		return &object.Error{Message: "errors.new() does not support keyword arguments"}
	}
	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("errors.new() expects a string argument, got %s", args[0].Type())}
	}
	return &object.Error{Message: msg.Value}
}
