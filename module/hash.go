package module

import (
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/vintlang/vintlang/object"
)

var HashFunctions = map[string]object.ModuleFunction{}

func init() {
	HashFunctions["sha1"] = hashSHA1
	HashFunctions["sha512"] = hashSHA512
}

func hashSHA1(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"hash", "sha1",
			"1 argument: data (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`hash.sha1("hello") -> "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"`,
		)
	}

	data := args[0]
	if data.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"hash", "sha1",
			"string data",
			string(data.Type()),
			`hash.sha1("hello") -> "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"`,
		)
	}

	input := data.(*object.String).Value
	hasher := sha1.New()
	hasher.Write([]byte(input))
	hash := hex.EncodeToString(hasher.Sum(nil))
	
	return &object.String{Value: hash}
}

func hashSHA512(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"hash", "sha512",
			"1 argument: data (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`hash.sha512("hello") -> "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"`,
		)
	}

	data := args[0]
	if data.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"hash", "sha512",
			"string data",
			string(data.Type()),
			`hash.sha512("hello") -> "9b71d224bd62f3785d96d46ad3ea3d73..."`,
		)
	}

	input := data.(*object.String).Value
	hasher := sha512.New()
	hasher.Write([]byte(input))
	hash := hex.EncodeToString(hasher.Sum(nil))
	
	return &object.String{Value: hash}
}