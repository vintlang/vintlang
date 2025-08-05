package module

import (
	"os/exec"
	"github.com/vintlang/vintlang/object"
)

var ShellFunctions = map[string]object.ModuleFunction{}

func init() {
	ShellFunctions["run"] = runCommand
	ShellFunctions["exists"] = commandExists
}

func runCommand(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"shell",
			"run",
			"1 string argument (command)",
			formatArgs(args),
			`shell.run("ls -la") -> string`,
		)
	}
	cmd := args[0].(*object.String).Value
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return &object.Error{Message: err.Error()}
	}
	return &object.String{Value: string(output)}
}

func commandExists(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"shell",
			"exists",
			"1 string argument (command name)",
			formatArgs(args),
			`shell.exists("git") -> true`,
		)
	}
	cmd := args[0].(*object.String).Value
	_, err := exec.LookPath(cmd)
	return &object.Boolean{Value: err == nil}
}
