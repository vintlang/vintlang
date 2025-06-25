package module

import "github.com/vintlang/vintlang/object"

var MySQLFunctions = map[string]object.ModuleFunction{}

func init() {
	MySQLFunctions["open"] = openConnection
	MySQLFunctions["close"] = closeConnection
}

func openConnection(args []object.Object, def map[string]object.Object) object.Object {}

func closeConnection(args []object.Object, def map[string]object.Object) object.Object {}
