package module

import (
	"encoding/base64"

	"github.com/vintlang/vintlang/object"
)

var EncodingFunctions = map[string]object.ModuleFunction{}

func init() {
	EncodingFunctions["base64Encode"] = base64Encode
	EncodingFunctions["base64Decode"] = base64Decode
}

func base64Encode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"encoding",
			"base64Encode",
			"1 string argument (text to encode)",
			formatArgs(args),
			`encoding.base64Encode("Hello") -> "SGVsbG8="`,
		)
	}

	input := args[0].(*object.String).Value
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	return &object.String{Value: encoded}
}

func base64Decode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"encoding",
			"base64Decode",
			"1 string argument (base64 encoded text)",
			formatArgs(args),
			`encoding.base64Decode("SGVsbG8=") -> "Hello"`,
		)
	}

	input := args[0].(*object.String).Value
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return &object.Error{
			Message: "\033[1;31mError in encoding.base64Decode()\033[0m:\n" +
				"  The provided string is not valid Base64.\n" +
				"  Usage: encoding.base64Decode(\"SGVsbG8=\") -> \"Hello\"\n",
		}
	}

	return &object.String{Value: string(decoded)}
}

// formatArgs converts the provided arguments into a string representation
// so we can display them in error messages.
func formatArgs(args []object.Object) string {
	if len(args) == 0 {
		return "no arguments"
	}
	result := ""
	for i, arg := range args {
		if i > 0 {
			result += ", "
		}
		result += string(arg.Type())
	}
	return result
}
