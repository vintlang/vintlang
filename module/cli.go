package module

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/toolkit"
)

var CliFunctions = map[string]object.ModuleFunction{}

func init() {
	CliFunctions["parseArgs"] = parseArgs
	CliFunctions["getArgValue"] = getArgValue
	CliFunctions["hasArg"] = hasArg
	CliFunctions["getArgs"] = getArgs
	CliFunctions["getFlags"] = getFlags
	CliFunctions["getPositional"] = getPositional
	CliFunctions["args"] = args
	CliFunctions["parse"] = argsParse
	CliFunctions["prompt"] = prompt
	CliFunctions["confirm"] = confirm
	CliFunctions["execCommand"] = execCommand
	CliFunctions["cliExit"] = cliExit
}

// getArgs returns an array of command line arguments
func getArgs(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 0 {
		return &object.Error{Message: "getArgs does not accept any arguments"}
	}

	cliArgs := &object.Array{}
	for _, arg := range toolkit.GetCliArgs() {
		cliArgs.Elements = append(cliArgs.Elements, &object.String{Value: arg})
	}
	return cliArgs
}

// getFlags parses command line flags into a dictionary
func getFlags(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 0 {
		return &object.Error{Message: "getFlags does not accept any arguments"}
	}

	flags := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	cliArgs := toolkit.GetCliArgs()

	for i := 0; i < len(cliArgs); i++ {
		arg := cliArgs[i]
		if strings.HasPrefix(arg, "--") {
			key := strings.TrimPrefix(arg, "--")
			var value object.Object = &object.Boolean{Value: true}

			// Check if next arg is a value (not a flag)
			if i+1 < len(cliArgs) && !strings.HasPrefix(cliArgs[i+1], "--") {
				value = &object.String{Value: cliArgs[i+1]}
				i++ // Skip the value in next iteration
			}

			hashKey := object.HashKey{Type: object.STRING_OBJ, Value: uint64(len(flags.Pairs))}
			flags.Pairs[hashKey] = object.DictPair{
				Key:   &object.String{Value: key},
				Value: value,
			}
		}
	}

	return flags
}

// getPositional returns an array of positional (non-flag) arguments
func getPositional(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 0 {
		return &object.Error{Message: "getPositional does not accept any arguments"}
	}

	positional := &object.Array{}
	cliArgs := toolkit.GetCliArgs()

	for i := 0; i < len(cliArgs); i++ {
		arg := cliArgs[i]
		if !strings.HasPrefix(arg, "-") {
			// This is a positional argument
			positional.Elements = append(positional.Elements, &object.String{Value: arg})
		} else if strings.HasPrefix(arg, "--") {
			// Skip the next argument if it's a value for this flag
			if i+1 < len(cliArgs) && !strings.HasPrefix(cliArgs[i+1], "-") {
				i++ // Skip the value
			}
		}
	}

	return positional
}

// prompt displays a message and reads user input
func prompt(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "prompt requires exactly one argument: the prompt message"}
	}

	message, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "prompt message must be a string"}
	}

	fmt.Print(message.Value)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Error reading input: %v", err)}
	}

	return &object.String{Value: strings.TrimSpace(input)}
}

// confirm prompts the user for a yes/no response
func confirm(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "confirm requires exactly one argument: the prompt message"}
	}

	message, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "confirm message must be a string"}
	}

	fmt.Printf("%s (y/n): ", message.Value)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Error reading input: %v", err)}
	}

	input = strings.ToLower(strings.TrimSpace(input))
	return &object.Boolean{Value: input == "y" || input == "yes"}
}

// execCommand executes a shell command and returns its output
func execCommand(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "execCommand requires exactly one argument: the command to execute"}
	}

	command, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "command must be a string"}
	}

	cmd := exec.Command("sh", "-c", command.Value)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Error executing command: %v", err)}
	}

	return &object.String{Value: string(output)}
}

// cliExit terminates the program with the given status code
func cliExit(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "exit requires exactly one argument: the status code"}
	}

	code, ok := args[0].(*object.Integer)
	if !ok {
		return &object.Error{Message: "status code must be an integer"}
	}

	os.Exit(int(code.Value))
	return &object.Null{}
}

// parseArgs parses command line arguments into a structured format
func parseArgs(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || len(defs) != 0 {
		return &object.Error{Message: "We need exactly one argument: a slice of strings representing the arguments"}
	}

	argStr := args[0].Inspect()
	argList := strings.Split(argStr, " ")

	parsedArgs := &object.Array{}
	for _, arg := range argList {
		parsedArgs.Elements = append(parsedArgs.Elements, &object.String{Value: arg})
	}

	return parsedArgs
}

// getArgValue gets the value of a named argument
func getArgValue(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || len(defs) != 0 {
		return &object.Error{Message: "We need exactly one argument: the argument name"}
	}

	argName := args[0].Inspect()
	// Remove quotes if present
	argName = strings.Trim(argName, `"'`)

	cliArgs := toolkit.GetCliArgs()

	// Check for --flag=value format first
	for _, arg := range cliArgs {
		if strings.HasPrefix(arg, argName+"=") {
			return &object.String{Value: strings.Split(arg, "=")[1]}
		}
	}

	// Check for --flag value format
	for i, arg := range cliArgs {
		if arg == argName && i+1 < len(cliArgs) {
			// Make sure next arg is not another flag
			if !strings.HasPrefix(cliArgs[i+1], "-") {
				return &object.String{Value: cliArgs[i+1]}
			}
		}
	}

	return &object.Null{}
}

// hasArg checks if a named argument is present
func hasArg(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || len(defs) != 0 {
		return &object.Error{Message: "We need exactly one argument: the argument name"}
	}

	argName := args[0].Inspect()
	// Remove quotes if present
	argName = strings.Trim(argName, `"'`)

	cliArgs := toolkit.GetCliArgs()

	for _, arg := range cliArgs {
		// Check for exact match or --flag=value format
		if arg == argName || strings.HasPrefix(arg, argName+"=") {
			return &object.Boolean{Value: true}
		}
	}

	return &object.Boolean{Value: false}
}

// args returns an array of all command line arguments (alias for getArgs)
func args(args []object.Object, defs map[string]object.Object) object.Object {
	return getArgs(args, defs)
}

// argsParse returns a parsed arguments object with helper methods
func argsParse(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 0 {
		return &object.Error{Message: "parse does not accept any arguments"}
	}

	cliArgs := toolkit.GetCliArgs()

	// Create a dictionary with parsed arguments and helper methods
	result := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Add flags to the result
	flags := make(map[string]object.Object)
	positionalArgs := []object.Object{}

	for i := 0; i < len(cliArgs); i++ {
		arg := cliArgs[i]
		if strings.HasPrefix(arg, "--") {
			key := strings.TrimPrefix(arg, "--")
			var value object.Object = &object.Boolean{Value: true}

			// Check if next arg is a value (not a flag)
			if i+1 < len(cliArgs) && !strings.HasPrefix(cliArgs[i+1], "--") && !strings.HasPrefix(cliArgs[i+1], "-") {
				value = &object.String{Value: cliArgs[i+1]}
				i++ // Skip the value in next iteration
			}
			flags[key] = value
		} else if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			// Handle short flags like -v
			key := strings.TrimPrefix(arg, "-")
			flags[key] = &object.Boolean{Value: true}
		} else {
			// Positional argument
			positionalArgs = append(positionalArgs, &object.String{Value: arg})
		}
	}

	// Add flags to result dict
	flagsHashKey := object.HashKey{Type: object.STRING_OBJ, Value: 0}
	result.Pairs[flagsHashKey] = object.DictPair{
		Key:   &object.String{Value: "flags"},
		Value: createDictFromMap(flags),
	}

	// Add positional arguments
	posArgsArray := &object.Array{Elements: positionalArgs}
	posHashKey := object.HashKey{Type: object.STRING_OBJ, Value: 1}
	result.Pairs[posHashKey] = object.DictPair{
		Key:   &object.String{Value: "positional"},
		Value: posArgsArray,
	}

	// Add has method
	hasHashKey := object.HashKey{Type: object.STRING_OBJ, Value: 2}
	result.Pairs[hasHashKey] = object.DictPair{
		Key:   &object.String{Value: "has"},
		Value: &object.Builtin{Fn: createHasFunction(flags)},
	}

	// Add get method
	getHashKey := object.HashKey{Type: object.STRING_OBJ, Value: 3}
	result.Pairs[getHashKey] = object.DictPair{
		Key:   &object.String{Value: "get"},
		Value: &object.Builtin{Fn: createGetFunction(flags)},
	}

	// Add positional method
	positionalHashKey := object.HashKey{Type: object.STRING_OBJ, Value: 4}
	result.Pairs[positionalHashKey] = object.DictPair{
		Key:   &object.String{Value: "positional"},
		Value: &object.Builtin{Fn: createPositionalFunction(positionalArgs)},
	}

	return result
}

// Helper function to create a dictionary from a map
func createDictFromMap(flags map[string]object.Object) *object.Dict {
	dict := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	i := uint64(0)
	for key, value := range flags {
		hashKey := object.HashKey{Type: object.STRING_OBJ, Value: i}
		dict.Pairs[hashKey] = object.DictPair{
			Key:   &object.String{Value: key},
			Value: value,
		}
		i++
	}
	return dict
}

// Helper function to create the has method
func createHasFunction(flags map[string]object.Object) func(...object.Object) object.Object {
	return func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return &object.Error{Message: "has requires exactly one argument: the flag name"}
		}

		flagName, ok := args[0].(*object.String)
		if !ok {
			return &object.Error{Message: "flag name must be a string"}
		}

		name := strings.TrimPrefix(strings.TrimPrefix(flagName.Value, "--"), "-")
		_, exists := flags[name]
		return &object.Boolean{Value: exists}
	}
}

// Helper function to create the get method
func createGetFunction(flags map[string]object.Object) func(...object.Object) object.Object {
	return func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return &object.Error{Message: "get requires exactly one argument: the flag name"}
		}

		flagName, ok := args[0].(*object.String)
		if !ok {
			return &object.Error{Message: "flag name must be a string"}
		}

		name := strings.TrimPrefix(strings.TrimPrefix(flagName.Value, "--"), "-")
		if value, exists := flags[name]; exists {
			return value
		}
		return &object.Null{}
	}
}

// Helper function to create the positional method
func createPositionalFunction(positionalArgs []object.Object) func(...object.Object) object.Object {
	return func(args ...object.Object) object.Object {
		if len(args) > 0 {
			return &object.Error{Message: "positional does not accept any arguments"}
		}

		return &object.Array{Elements: positionalArgs}
	}
}
