package module

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/ekilie/vint-lang/object"
)

var DotenvFunctions = map[string]object.ModuleFunction{}

func init() {
	DotenvFunctions["load"] = loadEnv
	DotenvFunctions["get"] = getDotenvValue
}

func loadEnv(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 1 || len(defs) != 0 {
		return &object.Error{Message: "load requires zero or one argument: the .env file path"}
	}

	filePath := ""
	if len(args) == 1 {
		filePath = args[0].Inspect()
	}

	err := godotenv.Load(filePath)
	if err != nil {
		return &object.Error{Message: "Failed to load .env file"}
	}

	return &object.Boolean{Value: true}
}

func getDotenvValue(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || len(defs) != 0 {
		return &object.Error{Message: "get requires exactly one argument: the key name"}
	}

	key := args[0].Inspect()
	value := os.Getenv(strings.Trim(key, `"'`))
	return &object.String{Value: value}
}
