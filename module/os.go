package module

import (
	"os"
	"os/exec"
	"strings"

	"github.com/ekilie/vint-lang/object"
)

var OsFunctions = map[string]object.ModuleFunction{}

func init() {
	OsFunctions["exit"] = exit
	OsFunctions["run"] = run
}

func exit(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 1 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	if len(args) == 1 {
		status, ok := args[0].(*object.Integer)
		if !ok {
			return &object.Error{Message: "Argument must be a number"}
		}
		os.Exit(int(status.Value))
		return nil
	}

	os.Exit(0)

	return nil
}

func run(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	cmd, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Argument must be a string"}
	}
	cmdMain := cmd.Value
	cmdArgs := strings.Split(cmdMain, " ")
	cmdArgs = cmdArgs[1:]

	out, err := exec.Command(cmdMain, cmdArgs...).Output()
	if err != nil {
		return &object.Error{Message: "Failed to execute command"}
	}

	return &object.String{Value: string(out)}
}
