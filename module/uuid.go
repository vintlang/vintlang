package module

import (
	"github.com/google/uuid"
	"github.com/ekilie/vint-lang/object"
)

var UuidFunctions = map[string]object.ModuleFunction{}

func init() {
	UuidFunctions["generate"] = generateUUID
}

func generateUUID(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 || len(defs) != 0 {
		return &object.Error{Message: "No arguments required for UUID generation"}
	}

	// Generate a new UUID using the Google UUID package
	id := uuid.New().String()

	return &object.String{Value: id}
}
