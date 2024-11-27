package module

import (
	"regexp"
	"github.com/ekilie/vint-lang/object"
)

// RegexFunctions is a map that holds the available functions in the Regex module.
var RegexFunctions = map[string]object.ModuleFunction{
	"match":        match,
	"replaceString": replaceString,
	"splitString":   splitString,
}


// match checks if the given pattern matches the input string.
// It returns true if there is a match, otherwise returns false.
func match(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "We need two arguments: the pattern and the input string"}
	}

	// Get the regex pattern and input string values
	pattern := args[0].Inspect()
	input := args[1].Inspect()

	// Compile the regex pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		return &object.Error{Message: err.Error()}
	}

	// Check if the pattern matches the input string
	return &object.Boolean{Value: re.MatchString(input)}
}

// replaceString replaces the first occurrence of the regex pattern in the input string with the replacement string.
func replaceString(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "We need three arguments: the pattern, the replacement, and the input string"}
	}

	// Get the regex pattern, replacement string, and input string
	pattern := args[0].Inspect()
	replacement := args[1].Inspect()
	input := args[2].Inspect()

	// Compile the regex pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		return &object.Error{Message: err.Error()}
	}

	// Replace the first occurrence of the pattern with the replacement
	return &object.String{Value: re.ReplaceAllString(input, replacement)}
}

// splitString splits the input string by the regex pattern and returns an array of substrings.
func splitString(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "We need two arguments: the pattern and the input string"}
	}

	// Get the regex pattern and input string
	pattern := args[0].Inspect()
	input := args[1].Inspect()

	// Compile the regex pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		return &object.Error{Message: err.Error()}
	}

	// Split the input string by the pattern
	result := re.Split(input, -1)

	// Convert the result to an array of strings
	var resultObjects []object.Object
	for _, item := range result {
		resultObjects = append(resultObjects, &object.String{Value: item})
	}

	// Return the array of substrings
	return &object.Array{Elements: resultObjects}
}

