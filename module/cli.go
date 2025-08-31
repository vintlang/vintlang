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
	CliFunctions["help"] = cliHelp
	CliFunctions["version"] = cliVersion
}

// getArgs returns an array of command line arguments
func getArgs(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 0 {
		return ErrorMessage(
			"cli", "getArgs",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"cli.getArgs()",
		)
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
		return ErrorMessage(
			"cli", "getFlags",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"cli.getFlags()",
		)
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
		return ErrorMessage(
			"cli", "getPositional",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"cli.getPositional() -> returns array of positional arguments",
		)
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
		return ErrorMessage(
			"cli", "prompt",
			"1 string argument (prompt message)",
			fmt.Sprintf("%d arguments", len(args)),
			`cli.prompt("Enter your name: ") -> returns user input`,
		)
	}

	message, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"cli", "prompt",
			"string argument for prompt message",
			string(args[0].Type()),
			`cli.prompt("Enter your name: ") -> returns user input`,
		)
	}

	fmt.Print(message.Value)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return &object.Error{
			Message: "\033[1;31m -> cli.prompt()\033[0m:\n" +
				"  Failed to read user input from terminal.\n" +
				"  This may indicate terminal settings or input stream issues.\n" +
				"  Usage: cli.prompt(\"Enter your name: \") -> returns user input\n",
		}
	}

	return &object.String{Value: strings.TrimSpace(input)}
}

// confirm prompts the user for a yes/no response
func confirm(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"cli", "confirm",
			"1 string argument (confirmation message)",
			fmt.Sprintf("%d arguments", len(args)),
			`cli.confirm("Continue with operation?") -> returns true/false`,
		)
	}

	message, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"cli", "confirm",
			"string argument for confirmation message",
			string(args[0].Type()),
			`cli.confirm("Continue with operation?") -> returns true/false`,
		)
	}

	fmt.Printf("%s (y/n): ", message.Value)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("cli.confirm() failed to read user input: %v. Please check your terminal settings and try again.", err)}
	}

	input = strings.ToLower(strings.TrimSpace(input))
	return &object.Boolean{Value: input == "y" || input == "yes"}
}

// execCommand executes a shell command and returns its output
func execCommand(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"cli", "execCommand",
			"1 string argument (command string)",
			fmt.Sprintf("%d arguments", len(args)),
			`cli.execCommand("ls -la")`,
		)
	}

	command, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"cli", "execCommand",
			"string argument for shell command",
			string(args[0].Type()),
			`cli.execCommand("ls -la")`,
		)
	}

	if strings.TrimSpace(command.Value) == "" {
		return ErrorMessage(
			"cli", "execCommand",
			"non-empty command string",
			"empty command",
			`cli.execCommand("ls -la") - provide a valid shell command`,
		)
	}

	cmd := exec.Command("sh", "-c", command.Value)
	output, err := cmd.CombinedOutput()
	if err != nil {
		exitErr, isExitError := err.(*exec.ExitError)
		if isExitError {
			return &object.Error{Message: fmt.Sprintf("cli.execCommand() failed to execute '%s': command exited with status %d. Output: %s", command.Value, exitErr.ExitCode(), string(output))}
		}
		return &object.Error{Message: fmt.Sprintf("cli.execCommand() failed to execute '%s': %v", command.Value, err)}
	}

	return &object.String{Value: string(output)}
}

// cliExit terminates the program with the given status code
func cliExit(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"cli", "exit",
			"1 integer argument (status code)",
			fmt.Sprintf("%d arguments", len(args)),
			"cli.exit(0) for success or cli.exit(1) for error",
		)
	}

	code, ok := args[0].(*object.Integer)
	if !ok {
		return ErrorMessage(
			"cli", "exit",
			"integer status code",
			string(args[0].Type()),
			"cli.exit(0) for success or cli.exit(1) for error",
		)
	}

	if code.Value < 0 || code.Value > 255 {
		return ErrorMessage(
			"cli", "exit",
			"status code between 0 and 255",
			fmt.Sprintf("status code %d", code.Value),
			"cli.exit(0) for success, 1-255 for various error conditions",
		)
	}

	os.Exit(int(code.Value))
	return &object.Null{}
}

// parseArgs parses command line arguments into a structured format
func parseArgs(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || len(defs) != 0 {
		return ErrorMessage(
			"cli", "parseArgs",
			"1 array argument (slice of strings representing arguments)",
			fmt.Sprintf("%d arguments", len(args)),
			"cli.parseArgs([\"--flag\", \"value\", \"positional\"])",
		)
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
		return ErrorMessage(
			"cli", "getArgValue",
			"1 string argument (flag name)",
			fmt.Sprintf("%d arguments", len(args)),
			`cli.getArgValue("--output") -> returns flag value or null`,
		)
	}

	argNameObj, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"cli", "getArgValue",
			"string argument for flag name",
			string(args[0].Type()),
			`cli.getArgValue("--output") -> returns flag value or null`,
		)
	}

	argName := argNameObj.Value
	// Remove quotes if present
	argName = strings.Trim(argName, `"'`)

	if strings.TrimSpace(argName) == "" {
		return &object.Error{
			Message: "\033[1;31m -> cli.getArgValue()\033[0m:\n" +
				"  Cannot search for an empty flag name.\n" +
				"  Please provide a valid flag like \"--output\" or \"-o\".\n" +
				"  Usage: cli.getArgValue(\"--output\") -> returns flag value or null\n",
		}
	}

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
		return ErrorMessage(
			"cli", "hasArg",
			"1 string argument (flag name)",
			fmt.Sprintf("%d arguments", len(args)),
			`cli.hasArg("--verbose") -> returns true/false`,
		)
	}

	argNameObj, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"cli", "hasArg",
			"string argument for flag name",
			string(args[0].Type()),
			`cli.hasArg("--verbose") -> returns true/false`,
		)
	}

	argName := argNameObj.Value
	// Remove quotes if present
	argName = strings.Trim(argName, `"'`)

	if strings.TrimSpace(argName) == "" {
		return &object.Error{
			Message: "\033[1;31m -> cli.hasArg()\033[0m:\n" +
				"  Cannot search for an empty flag name.\n" +
				"  Please provide a valid flag like \"--verbose\" or \"-v\".\n" +
				"  Usage: cli.hasArg(\"--verbose\") -> returns true/false\n",
		}
	}

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
		return ErrorMessage(
			"cli", "parse",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"cli.parse()",
		)
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
			return ErrorMessage(
				"cli", "has",
				"1 string argument (flag name)",
				fmt.Sprintf("%d arguments", len(args)),
				"parsed.has(\"flagname\")",
			)
		}

		flagName, ok := args[0].(*object.String)
		if !ok {
			return ErrorMessage(
				"cli", "has",
				"string argument for flag name",
				string(args[0].Type()),
				"parsed.has(\"flagname\")",
			)
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
			return ErrorMessage(
				"cli", "get",
				"1 string argument (flag name)",
				fmt.Sprintf("%d arguments", len(args)),
				"parsed.get(\"flagname\")",
			)
		}

		flagName, ok := args[0].(*object.String)
		if !ok {
			return ErrorMessage(
				"cli", "get",
				"string argument for flag name",
				string(args[0].Type()),
				"parsed.get(\"flagname\")",
			)
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
			return ErrorMessage(
				"cli", "positional",
				"no arguments",
				fmt.Sprintf("%d arguments", len(args)),
				"parsed.positional()",
			)
		}

		return &object.Array{Elements: positionalArgs}
	}
}

// cliHelp generates and displays help text for CLI applications
func cliHelp(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 2 {
		return ErrorMessage(
			"cli", "help",
			"0-2 arguments: optional app name and description",
			fmt.Sprintf("%d arguments", len(args)),
			"cli.help() or cli.help(\"myapp\", \"My application description\")",
		)
	}

	appName := "app"
	description := "A VintLang command-line application"

	if len(args) >= 1 {
		if name, ok := args[0].(*object.String); ok {
			appName = name.Value
		}
	}

	if len(args) >= 2 {
		if desc, ok := args[1].(*object.String); ok {
			description = desc.Value
		}
	}

	helpText := fmt.Sprintf(`%s

%s

Usage:
  %s [options] [arguments]

Options:
  --help, -h     Show this help message
  --version, -v  Show version information
  --verbose      Enable verbose output
  --output FILE  Specify output file
  --input FILE   Specify input file

Examples:
  %s --help
  %s --verbose input.txt
  %s --output result.txt --input data.txt

For more information, visit: https://github.com/vintlang/vintlang
`, appName, description, appName, appName, appName, appName)

	fmt.Print(helpText)
	return &object.Null{}
}

// cliVersion displays version information  
func cliVersion(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 2 {
		return ErrorMessage(
			"cli", "version",
			"0-2 arguments: optional app name and version",
			fmt.Sprintf("%d arguments", len(args)),
			"cli.version() or cli.version(\"myapp\", \"1.0.0\")",
		)
	}

	appName := "VintLang CLI Application"
	version := "1.0.0"

	if len(args) >= 1 {
		if name, ok := args[0].(*object.String); ok {
			appName = name.Value
		}
	}

	if len(args) >= 2 {
		if ver, ok := args[1].(*object.String); ok {
			version = ver.Value
		}
	}

	versionText := fmt.Sprintf("%s v%s\n", appName, version)
	fmt.Print(versionText)
	return &object.Null{}
}
