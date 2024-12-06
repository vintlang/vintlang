package module

import (
	"encoding/base64"

	"github.com/ekilie/vint-lang/object"
)

var EncodingFunctions = map[string]object.ModuleFunction{}

func init() {
	EncodingFunctions["base64Encode"] = base64Encode
	EncodingFunctions["base64Decode"] = base64Decode
}

func base64Encode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "base64Encode requires one argument: the string to encode"}
	}

	input := args[0].Inspect()
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	return &object.String{Value: encoded}
}

func base64Decode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "base64Decode requires one argument: the string to decode"}
	}

	input := args[0].Inspect()
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return &object.Error{Message: "Invalid base64 string"}
	}

	return &object.String{Value: string(decoded)}
}
