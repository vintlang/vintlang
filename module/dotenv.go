package module

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/vintlang/vintlang/object"
)

var DotenvFunctions = map[string]object.ModuleFunction{}

func init() {
	DotenvFunctions["load"] = loadEnv
	DotenvFunctions["get"] = getDotenvValue
}

func loadEnv(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) > 1 || len(defs) != 0 {
		return ErrorMessage(
			"dotenv", "load",
			"0 or 1 string argument (optional file path to .env)",
			fmt.Sprintf("%d arguments", len(args)),
			`dotenv.load() -> loads environment variables from .env file in current directory
dotenv.load("path/to/.env") -> loads from specified file`,
		)
	}

	filePath := ""
	if len(args) == 1 {
		filePath = args[0].Inspect()
	}

	err := godotenv.Load(filePath)
	if err != nil {
		return ErrorMessage(
			"dotenv", "load",
			"valid .env file path",
			fmt.Sprintf("error: %v", err),
			`dotenv.load() -> loads environment variables from .env file in current directory
dotenv.load("path/to/.env") -> loads from specified file`,
		)
	}

	return &object.Boolean{Value: true}
}

func getDotenvValue(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || len(defs) != 0 {
		return ErrorMessage(
			"dotenv", "get",
			"1 string argument (environment variable key)",
			fmt.Sprintf("%d arguments", len(args)),
			`dotenv.get("MY_VAR") -> returns the value of MY_VAR from the environment`,
		)
	}

	key := args[0].Inspect()
	value := os.Getenv(strings.Trim(key, `"'`))
	return &object.String{Value: value}
}
