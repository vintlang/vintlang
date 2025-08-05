package module

import (
	"regexp"
	"github.com/vintlang/vintlang/object"
)

var RegexFunctions = map[string]object.ModuleFunction{
	"match":         match,
	"replaceString": replaceString,
	"splitString":   splitString,
}

func match(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"regex",
			"match",
			"2 string arguments (pattern, text)",
			formatArgs(args),
			`regex.match("\\d+", "abc123") -> true`,
		)
	}
	pattern := args[0].(*object.String).Value
	input := args[1].(*object.String).Value

	re, err := regexp.Compile(pattern)
	if err != nil {
		return &object.Error{Message: "Invalid regex pattern: " + err.Error()}
	}
	return &object.Boolean{Value: re.MatchString(input)}
}

func replaceString(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 3 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"regex",
			"replaceString",
			"3 string arguments (pattern, replacement, text)",
			formatArgs(args),
			`regex.replaceString("\\d+", "#", "abc123") -> "abc#"`,
		)
	}
	pattern := args[0].(*object.String).Value
	replacement := args[1].(*object.String).Value
	input := args[2].(*object.String).Value

	re, err := regexp.Compile(pattern)
	if err != nil {
		return &object.Error{Message: "Invalid regex pattern: " + err.Error()}
	}
	return &object.String{Value: re.ReplaceAllString(input, replacement)}
}

func splitString(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"regex",
			"splitString",
			"2 string arguments (pattern, text)",
			formatArgs(args),
			`regex.splitString("\\s+", "a b  c") -> ["a", "b", "c"]`,
		)
	}
	pattern := args[0].(*object.String).Value
	input := args[1].(*object.String).Value

	re, err := regexp.Compile(pattern)
	if err != nil {
		return &object.Error{Message: "Invalid regex pattern: " + err.Error()}
	}
	result := re.Split(input, -1)
	var objs []object.Object
	for _, s := range result {
		objs = append(objs, &object.String{Value: s})
	}
	return &object.Array{Elements: objs}
}
