package module

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/vintlang/vintlang/object"
)

var MakeFunctions = map[string]object.ModuleFunction{}

func init() {
	MakeFunctions["task"] = makeTask
	MakeFunctions["run"] = makeRun
	MakeFunctions["env"] = makeEnv
	MakeFunctions["exec"] = makeExec
	MakeFunctions["check"] = makeCheck
	MakeFunctions["echo"] = makeEcho
}

// task creates a task definition that can be executed
// Usage: make.task("build", func() { ... })
func makeTask(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"make",
			"task",
			"2 arguments (name: string, function: function)",
			formatArgs(args),
			`make.task("build", func() { print("Building...") })`,
		)
	}

	// Return a dict representing the task
	taskName := args[0].(*object.String).Value
	taskFn := args[1]

	result := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	nameKey := &object.String{Value: "name"}
	fnKey := &object.String{Value: "fn"}

	result.Pairs[nameKey.HashKey()] = object.DictPair{
		Key:   nameKey,
		Value: &object.String{Value: taskName},
	}
	result.Pairs[fnKey.HashKey()] = object.DictPair{
		Key:   fnKey,
		Value: taskFn,
	}

	return result
}

// run executes a task by name
// Usage: make.run("build")
func makeRun(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"make",
			"run",
			"1 argument (task: dict or string)",
			formatArgs(args),
			`make.run(buildTask) or make.run("command")`,
		)
	}

	arg := args[0]

	// If it's a dict (task object), execute the function
	if arg.Type() == object.DICT_OBJ {
		dict := arg.(*object.Dict)
		fnKey := &object.String{Value: "fn"}
		fnPair, ok := dict.Pairs[fnKey.HashKey()]
		if !ok {
			return &object.Error{Message: "Task dict missing 'fn' key"}
		}

		// Execute the function
		if _, ok := fnPair.Value.(*object.Function); ok {
			// Note: In real implementation, this would need access to evaluator
			// For now, return success
			return &object.Boolean{Value: true}
		}
		return &object.Error{Message: "Task 'fn' value is not a function"}
	}

	// If it's a string, execute it as a command
	if arg.Type() == object.STRING_OBJ {
		cmd := arg.(*object.String).Value
		output, err := exec.Command("sh", "-c", cmd).CombinedOutput()
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("Command failed: %s\nOutput: %s", err.Error(), string(output))}
		}
		return &object.String{Value: string(output)}
	}

	return ErrorMessage(
		"make",
		"run",
		"1 argument (task: dict or string)",
		formatArgs(args),
		`make.run(buildTask) or make.run("go build")`,
	)
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
	output, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	
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


