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
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, argName+"=") {
			return &object.String{Value: strings.Split(arg, "=")[1]}
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
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, argName+"=") {
			return &object.Boolean{Value: true}
		}
	}

	return &object.Boolean{Value: false}
}
