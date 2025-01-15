package module

import (
	"os/exec"
	// "strings"

	"github.com/vintlang/vintlang/object"
)

var ShellFunctions = map[string]object.ModuleFunction{}

func init() {
	ShellFunctions["run"] = runCommand
	ShellFunctions["exists"] = commandExists
}

func runCommand(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || len(defs) != 0 {
		return &object.Error{Message: "run requires exactly one argument: the command string"}
	}

	cmd := args[0].Inspect()
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return &object.Error{Message: err.Error()}
	}

	return &object.String{Value: string(output)}
}

func commandExists(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 || len(defs) != 0 {
		return &object.Error{Message: "exists requires exactly one argument: the command name"}
	}

	cmd := args[0].Inspect()
	_, err := exec.LookPath(cmd)
	return &object.Boolean{Value: err == nil}
}
