package module

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vintlang/vintlang/object"
)

var UuidFunctions = map[string]object.ModuleFunction{}

func init() {
	UuidFunctions["generate"] = generateUUID
}

func generateUUID(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"uuid", "generate",
			"No arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`uuid.generate() -> "550e8400-e29b-41d4-a716-446655440000"`,
		)
	}
	return &object.String{Value: uuid.New().String()}
}
