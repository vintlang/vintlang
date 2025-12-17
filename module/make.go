package module

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/vintlang/vintlang/object"
)

var MakeFunctions = map[string]object.ModuleFunction{}

func init() {
	MakeFunctions["env"] = makeEnv
	MakeFunctions["exec"] = makeExec
	MakeFunctions["check"] = makeCheck
	MakeFunctions["echo"] = makeEcho
}

// getShell returns the appropriate shell for the current OS
func getShell() (string, string) {
	if runtime.GOOS == "windows" {
		return "cmd", "/C"
	}
	return "sh", "-c"
}

// env sets an environment variable for the current process
// Usage: make.env("GOOS", "linux")
func makeEnv(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"make",
			"env",
			"2 string arguments (key, value)",
			formatArgs(args),
			`make.env("GOOS", "linux")`,
		)
	}

	key := args[0].(*object.String).Value
	value := args[1].(*object.String).Value

	err := os.Setenv(key, value)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to set environment variable: %s", err.Error())}
	}

	return &object.Boolean{Value: true}
}

// exec executes a shell command and returns the output
// Usage: make.exec("ls -la")
func makeExec(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"make",
			"exec",
			"1 string argument (command)",
			formatArgs(args),
			`make.exec("ls -la")`,
		)
	}

	cmd := args[0].(*object.String).Value
	shell, flag := getShell()
	output, err := exec.Command(shell, flag, cmd).CombinedOutput()
	
	if err != nil {
		// Return both output and error information
		errorMsg := fmt.Sprintf("Command failed: %s\nOutput: %s", err.Error(), string(output))
		return &object.Error{Message: errorMsg}
	}

	return &object.String{Value: string(output)}
}

// check verifies if a command exists in PATH
// Usage: make.check("upx")
func makeCheck(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"make",
			"check",
			"1 string argument (command name)",
			formatArgs(args),
			`make.check("upx")`,
		)
	}

	cmd := args[0].(*object.String).Value
	_, err := exec.LookPath(cmd)

	return &object.Boolean{Value: err == nil}
}

// echo prints a message (useful for build scripts)
// Usage: make.echo("Building...")
func makeEcho(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"make",
			"echo",
			"1 string argument (message)",
			formatArgs(args),
			`make.echo("Building linux binary...")`,
		)
	}

	message := args[0].(*object.String).Value
	fmt.Println(message)

	return &object.Null{}
}


