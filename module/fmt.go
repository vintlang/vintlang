package module

import (
	"fmt"

	"github.com/vintlang/vintlang/object"
)

var FmtFunctions = map[string]object.ModuleFunction{}

func init() {
	FmtFunctions["sprintf"] = sprintf
}

// sprintf formats a string using the provided format and arguments
func sprintf(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"fmt", "sprintf",
			"at least 1 string argument (format string)",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.sprintf("Hello, %s!", "World") -> returns "Hello, World!"`,
		)
	}

	format := args[0].(*object.String).Value
	var formatArgs []interface{}
	for _, arg := range args[1:] {
		formatArgs = append(formatArgs, VintObjectToInterface(arg))
	}

	result := fmt.Sprintf(format, formatArgs...)
	return &object.String{Value: result}
}
