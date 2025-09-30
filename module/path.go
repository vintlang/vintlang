package module

import (
	"path/filepath"

	"github.com/vintlang/vintlang/object"
)

var PathFunctions = map[string]object.ModuleFunction{}

func init() {
	PathFunctions["join"] = joinPaths
	PathFunctions["basename"] = basePath
	PathFunctions["dirname"] = dirName
	PathFunctions["ext"] = extName
	PathFunctions["isAbs"] = isAbs
}

func joinPaths(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) == 0 {
		return &object.Error{Message: "join() expects at least one argument"}
	}
	paths := make([]string, len(args))
	for i, arg := range args {
		if arg.Type() != object.STRING_OBJ {
			return &object.Error{Message: "All arguments to join() must be strings"}
		}
		paths[i] = arg.(*object.String).Value
	}
	return &object.String{Value: filepath.Join(paths...)}
}

func basePath(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "basename() expects a single string argument"}
	}
	path := args[0].(*object.String).Value
	return &object.String{Value: filepath.Base(path)}
}

func dirName(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "dirname() expects a single string argument"}
	}
	path := args[0].(*object.String).Value
	return &object.String{Value: filepath.Dir(path)}
}

func extName(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "ext() expects a single string argument"}
	}
	path := args[0].(*object.String).Value
	return &object.String{Value: filepath.Ext(path)}
}

func isAbs(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "isAbs() expects a single string argument"}
	}
	path := args[0].(*object.String).Value
	return &object.Boolean{Value: filepath.IsAbs(path)}
}
