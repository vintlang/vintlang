package module

import (
	"os"
	"strings"

	"github.com/ekilie/vint-lang/object"
	"github.com/ekilie/vint-lang/toolkit"
)

var CliFunctions = map[string]object.ModuleFunction{}

func init() {
	CliFunctions["parseArgs"] = parseArgs
	CliFunctions["getArgValue"] = getArgValue
	CliFunctions["hasArg"] = hasArg
	CliFunctions["getArgs"] = getArgs
}
//Returns an array/list of args as strings []strings
func getArgs(args []object.Object, defs map[string]object.Object)object.Object{
	if len(args) > 0{
		return &object.Error{Message:"getArgs does not accept any arguments"}
	}

	cliArgs := &object.Array{}

	for _, arg := range toolkit.GetCliArgs() {
		cliArgs.Elements = append(cliArgs.Elements,&object.String{Value:arg})
	}

	return cliArgs
}

func parseArgs(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || len(defs) != 0 {
		return &object.Error{Message: "We need exactly one argument: a slice of strings representing the arguments"}
	}

	// Get the command-line arguments as a string
	argStr := args[0].Inspect()
	argList := strings.Split(argStr, " ")

	// Convert the list of arguments into a VintLang List object
	parsedArgs := &object.Array{}
	for _, arg := range argList {
		parsedArgs.Elements = append(parsedArgs.Elements, &object.String{Value: arg})
	}

	return parsedArgs
}

func getArgValue(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || len(defs) != 0 {
		return &object.Error{Message: "We need exactly one argument: the argument name"}
	}

	argName := args[0].Inspect()

	// Loop through the command-line arguments (os.Args[1:] because os.Args[0] is the program name)
	for _, arg := range os.Args[1:] {
		// If the argument matches the provided argument name, return it
		if strings.HasPrefix(arg, argName+"=") {
			return &object.String{Value: strings.Split(arg, "=")[1]}
		}
	}

	// If the argument is not found, return null
	return &object.Null{}
}

func hasArg(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || len(defs) != 0 {
		return &object.Error{Message: "We need exactly one argument: the argument name"}
	}

	argName := args[0].Inspect()

	// Check if the argument is present in os.Args
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, argName+"=") {
			return &object.Boolean{Value: true}
		}
	}

	return &object.Boolean{Value: false}
}
