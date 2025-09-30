package module

import (
	"fmt"
	"strings"

	"github.com/vintlang/vintlang/object"
)

var ArgparseFunctions = map[string]object.ModuleFunction{}

func init() {
	ArgparseFunctions["newParser"] = newParser
	ArgparseFunctions["addArgument"] = addArgument
	ArgparseFunctions["addFlag"] = addFlag
	ArgparseFunctions["parse"] = parse
	ArgparseFunctions["help"] = generateHelp
	ArgparseFunctions["version"] = setVersion
}

// Argument definition structure
type argDef struct {
	name        string
	shortName   string
	description string
	required    bool
	defaultVal  object.VintObject
	valueType   string
	choices     []string
}

// Parser structure to store argument definitions
var parsers = make(map[string]*object.Dict)
var parserArgs = make(map[string][]argDef)
var parserFlags = make(map[string][]argDef)
var parserVersions = make(map[string]string)
var parserDescriptions = make(map[string]string)

// newParser creates a new argument parser
func newParser(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || len(args) > 2 {
		return ErrorMessage("argparse", "newParser", "1-2 arguments: name and optional description", "", "")
	}

	name, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage("argparse", "newParser", "parser name must be a string", "", "")
	}

	// Create a new parser
	parserID := name.Value
	parsers[parserID] = &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	parserArgs[parserID] = []argDef{}
	parserFlags[parserID] = []argDef{}

	// Set description if provided
	if len(args) == 2 {
		desc, ok := args[1].(*object.String)
		if !ok {
			return ErrorMessage("argparse", "newParser", "description must be a string", "", "")
		}
		parserDescriptions[parserID] = desc.Value
	}

	return &object.String{Value: parserID}
}

// addArgument adds a positional argument to the parser
func addArgument(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 {
		return ErrorMessage("argparse", "addArgument", "addArgument requires at least 2 arguments: parser and name", "", "")
	}

	parserID, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage("argparse", "addArgument", "parser must be a string", "", "")
	}

	name, ok := args[1].(*object.String)
	if !ok {
		return ErrorMessage("argparse", "addArgument", "argument name must be a string", "", "")
	}

	// Check if parser exists
	if _, exists := parsers[parserID.Value]; !exists {
		return ErrorMessage("argparse", "addArgument", fmt.Sprintf("parser '%s' not found", parserID.Value), "", "")
	}

	// Create argument definition
	arg := argDef{
		name:        name.Value,
		description: "",
		required:    false,
		defaultVal:  &object.Null{},
		valueType:   "string",
	}

	// Process optional parameters
	for k, v := range defs {
		switch k {
		case "description":
			if desc, ok := v.(*object.String); ok {
				arg.description = desc.Value
			}
		case "required":
			if req, ok := v.(*object.Boolean); ok {
				arg.required = req.Value
			}
		case "default":
			arg.defaultVal = v
		case "type":
			if t, ok := v.(*object.String); ok {
				arg.valueType = t.Value
			}
		case "choices":
			if choices, ok := v.(*object.Array); ok {
				for _, choice := range choices.Elements {
					if c, ok := choice.(*object.String); ok {
						arg.choices = append(arg.choices, c.Value)
					}
				}
			}
		}
	}

	// Add argument to parser
	parserArgs[parserID.Value] = append(parserArgs[parserID.Value], arg)
	return &object.Boolean{Value: true}
}

// addFlag adds a flag (optional named argument) to the parser
func addFlag(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 {
		return ErrorMessage("argparse", "addFlag", "addFlag requires at least 2 arguments: parser and name", "", "")
	}

	parserID, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage("argparse", "addFlag", "parser must be a string", "", "")
	}

	name, ok := args[1].(*object.String)
	if !ok {
		return ErrorMessage("argparse", "addFlag", "flag name must be a string", "", "")
	}

	// Check if parser exists
	if _, exists := parsers[parserID.Value]; !exists {
		return ErrorMessage("argparse", "addFlag", fmt.Sprintf("parser '%s' not found", parserID.Value), "", "")
	}

	// Create flag definition
	flag := argDef{
		name:        name.Value,
		shortName:   "",
		description: "",
		required:    false,
		defaultVal:  &object.Boolean{Value: false},
		valueType:   "boolean",
	}

	// Process optional parameters
	for k, v := range defs {
		switch k {
		case "short":
			if short, ok := v.(*object.String); ok {
				flag.shortName = short.Value
			}
		case "description":
			if desc, ok := v.(*object.String); ok {
				flag.description = desc.Value
			}
		case "required":
			if req, ok := v.(*object.Boolean); ok {
				flag.required = req.Value
			}
		case "default":
			flag.defaultVal = v
		case "type":
			if t, ok := v.(*object.String); ok {
				flag.valueType = t.Value
			}
		}
	}

	// Add flag to parser
	parserFlags[parserID.Value] = append(parserFlags[parserID.Value], flag)
	return &object.Boolean{Value: true}
}

// parse parses command line arguments according to the parser definition
func parse(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 {
		return ErrorMessage("argparse", "parse", "parse requires at least 1 argument: parser", "", "")
	}

	parserID, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage("argparse", "parse", "parser must be a string", "", "")
	}

	// Check if parser exists
	_, exists := parsers[parserID.Value]
	if !exists {
		return ErrorMessage("argparse", "parse", fmt.Sprintf("parser '%s' not found", parserID.Value), "", "")
	}

	// Get command line arguments
	var cliArgs []string
	if len(args) > 1 {
		// Use provided arguments
		if arr, ok := args[1].(*object.Array); ok {
			for _, arg := range arr.Elements {
				if str, ok := arg.(*object.String); ok {
					cliArgs = append(cliArgs, str.Value)
				}
			}
		} else {
			return ErrorMessage("argparse", "parse", "arguments must be an array of strings", "", "")
		}
	} else {
		// Use system arguments
		cliArgs = []string{}
		for _, arg := range args {
			cliArgs = append(cliArgs, arg.Inspect())
		}
	}

	// Parse arguments
	result := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	// Set default values
	for _, arg := range parserArgs[parserID.Value] {
		hashKey := object.HashKey{Type: object.STRING_OBJ, Value: uint64(len(result.Pairs))}
		result.Pairs[hashKey] = object.DictPair{
			Key:   &object.String{Value: arg.name},
			Value: arg.defaultVal,
		}
	}

	for _, flag := range parserFlags[parserID.Value] {
		hashKey := object.HashKey{Type: object.STRING_OBJ, Value: uint64(len(result.Pairs))}
		result.Pairs[hashKey] = object.DictPair{
			Key:   &object.String{Value: flag.name},
			Value: flag.defaultVal,
		}
	}

	// Process arguments
	positionalIndex := 0
	for i := 0; i < len(cliArgs); i++ {
		arg := cliArgs[i]

		// Handle flags
		if strings.HasPrefix(arg, "--") {
			flagName := strings.TrimPrefix(arg, "--")

			// Check if it's a help flag
			if flagName == "help" {
				// Print help and exit
				fmt.Println(generateHelpText(parserID.Value))
				return result
			}

			// Check if it's a version flag
			if flagName == "version" && parserVersions[parserID.Value] != "" {
				fmt.Println(parserVersions[parserID.Value])
				return result
			}

			// Find the flag definition
			var flagDef *argDef
			for j := range parserFlags[parserID.Value] {
				if parserFlags[parserID.Value][j].name == flagName {
					flagDef = &parserFlags[parserID.Value][j]
					break
				}
			}

			if flagDef == nil {
				return ErrorMessage("argparse", "parse", fmt.Sprintf("unknown flag: %s", flagName), "", "")
			}

			// Handle flag value
			if flagDef.valueType == "boolean" {
				// Boolean flag
				hashKey := object.HashKey{Type: object.STRING_OBJ, Value: uint64(len(result.Pairs))}
				result.Pairs[hashKey] = object.DictPair{
					Key:   &object.String{Value: flagDef.name},
					Value: &object.Boolean{Value: true},
				}
			} else {
				// Flag with value
				if i+1 >= len(cliArgs) {
					return ErrorMessage("argparse", "parse", fmt.Sprintf("flag %s requires a value", flagName), "", "")
				}

				value := cliArgs[i+1]
				i++ // Skip the value in next iteration

				// Convert value based on type
				var objValue object.VintObject
				switch flagDef.valueType {
				case "string":
					objValue = &object.String{Value: value}
				case "integer":
					// TODO: Implement proper conversion
					objValue = &object.String{Value: value}
				case "float":
					// TODO: Implement proper conversion
					objValue = &object.String{Value: value}
				default:
					objValue = &object.String{Value: value}
				}

				hashKey := object.HashKey{Type: object.STRING_OBJ, Value: uint64(len(result.Pairs))}
				result.Pairs[hashKey] = object.DictPair{
					Key:   &object.String{Value: flagDef.name},
					Value: objValue,
				}
			}
		} else if strings.HasPrefix(arg, "-") {
			// Short flag
			shortName := strings.TrimPrefix(arg, "-")

			// Find the flag definition
			var flagDef *argDef
			for j := range parserFlags[parserID.Value] {
				if parserFlags[parserID.Value][j].shortName == shortName {
					flagDef = &parserFlags[parserID.Value][j]
					break
				}
			}

			if flagDef == nil {
				return ErrorMessage("argparse", "parse", fmt.Sprintf("unknown flag: -%s", shortName), "", "")
			}

			// Handle flag value (same logic as for long flags)
			if flagDef.valueType == "boolean" {
				hashKey := object.HashKey{Type: object.STRING_OBJ, Value: uint64(len(result.Pairs))}
				result.Pairs[hashKey] = object.DictPair{
					Key:   &object.String{Value: flagDef.name},
					Value: &object.Boolean{Value: true},
				}
			} else {
				if i+1 >= len(cliArgs) {
					return ErrorMessage("argparse", "parse", fmt.Sprintf("flag -%s requires a value", shortName), "", "")
				}

				value := cliArgs[i+1]
				i++ // Skip the value in next iteration

				var objValue object.VintObject
				switch flagDef.valueType {
				case "string":
					objValue = &object.String{Value: value}
				case "integer":
					// TODO: Implement proper conversion
					objValue = &object.String{Value: value}
				case "float":
					// TODO: Implement proper conversion
					objValue = &object.String{Value: value}
				default:
					objValue = &object.String{Value: value}
				}

				hashKey := object.HashKey{Type: object.STRING_OBJ, Value: uint64(len(result.Pairs))}
				result.Pairs[hashKey] = object.DictPair{
					Key:   &object.String{Value: flagDef.name},
					Value: objValue,
				}
			}
		} else {
			// Positional argument
			if positionalIndex >= len(parserArgs[parserID.Value]) {
				return ErrorMessage("argparse", "parse", "too many positional arguments", "", "")
			}

			argDef := parserArgs[parserID.Value][positionalIndex]

			// Convert value based on type
			var objValue object.VintObject
			switch argDef.valueType {
			case "string":
				objValue = &object.String{Value: arg}
			case "integer":
				// TODO: Implement proper conversion
				objValue = &object.String{Value: arg}
			case "float":
				// TODO: Implement proper conversion
				objValue = &object.String{Value: arg}
			default:
				objValue = &object.String{Value: arg}
			}

			hashKey := object.HashKey{Type: object.STRING_OBJ, Value: uint64(len(result.Pairs))}
			result.Pairs[hashKey] = object.DictPair{
				Key:   &object.String{Value: argDef.name},
				Value: objValue,
			}

			positionalIndex++
		}
	}

	// Check required arguments
	for _, arg := range parserArgs[parserID.Value] {
		if arg.required {
			found := false
			for _, pair := range result.Pairs {
				if pair.Key.(*object.String).Value == arg.name {
					if _, ok := pair.Value.(*object.Null); !ok {
						found = true
						break
					}
				}
			}

			if !found {
				return ErrorMessage("argparse", "parse", fmt.Sprintf("required argument '%s' not provided", arg.name), "", "")
			}
		}
	}

	// Check required flags
	for _, flag := range parserFlags[parserID.Value] {
		if flag.required {
			found := false
			for _, pair := range result.Pairs {
				if pair.Key.(*object.String).Value == flag.name {
					if _, ok := pair.Value.(*object.Null); !ok {
						found = true
						break
					}
				}
			}

			if !found {
				return ErrorMessage("argparse", "parse", fmt.Sprintf("required flag '--%s' not provided", flag.name), "", "")
			}
		}
	}

	return result
}

// generateHelp generates help text for the parser
func generateHelp(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage("argparse", "help", "help requires exactly 1 argument: parser", "", "")
	}

	parserID, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage("argparse", "help", "parser must be a string", "", "")
	}

	// Check if parser exists
	if _, exists := parsers[parserID.Value]; !exists {
		return ErrorMessage("argparse", "help", fmt.Sprintf("parser '%s' not found", parserID.Value), "", "")
	}

	helpText := generateHelpText(parserID.Value)
	return &object.String{Value: helpText}
}

// generateHelpText generates the help text for a parser
func generateHelpText(parserID string) string {
	var sb strings.Builder

	// Add description if available
	if desc, ok := parserDescriptions[parserID]; ok && desc != "" {
		sb.WriteString(desc)
		sb.WriteString("\n\n")
	}

	// Add usage
	sb.WriteString("Usage:\n")
	sb.WriteString("  program")

	// Add positional arguments to usage
	for _, arg := range parserArgs[parserID] {
		if arg.required {
			sb.WriteString(fmt.Sprintf(" <%s>", arg.name))
		} else {
			sb.WriteString(fmt.Sprintf(" [%s]", arg.name))
		}
	}

	// Add flags to usage
	if len(parserFlags[parserID]) > 0 {
		sb.WriteString(" [options]")
	}

	sb.WriteString("\n\n")

	// Add positional arguments section
	if len(parserArgs[parserID]) > 0 {
		sb.WriteString("Arguments:\n")
		for _, arg := range parserArgs[parserID] {
			sb.WriteString(fmt.Sprintf("  %-20s %s", arg.name, arg.description))
			if !arg.required && arg.defaultVal.Type() != object.NULL_OBJ {
				sb.WriteString(fmt.Sprintf(" (default: %s)", arg.defaultVal.Inspect()))
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	// Add flags section
	if len(parserFlags[parserID]) > 0 {
		sb.WriteString("Options:\n")
		for _, flag := range parserFlags[parserID] {
			if flag.shortName != "" {
				sb.WriteString(fmt.Sprintf("  -%-1s, --%-16s %s", flag.shortName, flag.name, flag.description))
			} else {
				sb.WriteString(fmt.Sprintf("      --%-16s %s", flag.name, flag.description))
			}

			if !flag.required && flag.defaultVal.Type() != object.NULL_OBJ {
				sb.WriteString(fmt.Sprintf(" (default: %s)", flag.defaultVal.Inspect()))
			}
			sb.WriteString("\n")
		}

		// Add built-in help and version flags
		sb.WriteString("      --help              Show this help message and exit\n")
		if parserVersions[parserID] != "" {
			sb.WriteString("      --version           Show version information and exit\n")
		}
	}

	return sb.String()
}

// setVersion sets the version information for the parser
func setVersion(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage("argparse", "version", "version requires exactly 2 arguments: parser and version string", "", "")
	}

	parserID, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage("argparse", "version", "parser must be a string", "", "")
	}

	version, ok := args[1].(*object.String)
	if !ok {
		return ErrorMessage("argparse", "version", "version must be a string", "", "")
	}

	// Check if parser exists
	if _, exists := parsers[parserID.Value]; !exists {
		return ErrorMessage("argparse", "version", fmt.Sprintf("parser '%s' not found", parserID.Value), "", "")
	}

	parserVersions[parserID.Value] = version.Value
	return &object.Boolean{Value: true}
}
