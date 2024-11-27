package module

import (
	"strings"
	"github.com/ekilie/vint-lang/object"
)

var StringFunctions = map[string]object.ModuleFunction{}

func init() {
	StringFunctions["trim"] = trim
	StringFunctions["contains"] = contains
	StringFunctions["toUpper"] = toUpper
	StringFunctions["toLower"] = toLower
}

func trim(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "We need one argument: the string to trim"}
	}

	str := args[0].Inspect()
	return &object.String{Value: strings.TrimSpace(str)}
}

func contains(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "We need two arguments: the string and the substring"}
	}

	str := args[0].Inspect()
	substr := args[1].Inspect()
	return &object.Boolean{Value: strings.Contains(str, substr)}
}

func toUpper(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "We need one argument: the string to convert"}
	}

	str := args[0].Inspect()
	return &object.String{Value: strings.ToUpper(str)}
}

func toLower(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "We need one argument: the string to convert"}
	}

	str := args[0].Inspect()
	return &object.String{Value: strings.ToLower(str)}
}
