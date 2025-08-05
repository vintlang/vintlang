package module

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/vintlang/vintlang/object"
)

var OsFunctions = map[string]object.ModuleFunction{}

func init() {
	OsFunctions["exit"] = exit
	OsFunctions["run"] = run
	OsFunctions["getEnv"] = getEnv
	OsFunctions["setEnv"] = setEnv
	OsFunctions["readFile"] = readFile
	OsFunctions["writeFile"] = writeFile
	OsFunctions["listDir"] = listDir
	OsFunctions["deleteFile"] = deleteFile
	OsFunctions["makeDir"] = makeDir
	OsFunctions["removeDir"] = removeDir
	OsFunctions["currentDir"] = currentDir
	OsFunctions["changeDir"] = changeDir
	OsFunctions["fileExists"] = fileExists
	OsFunctions["readLines"] = readLines
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
		return ErrorMessage(
			"os", "run",
			"1 string argument (command to execute)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.run("ls -la") -> returns command output`,
		)
	}

	cmd, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "run",
			"string argument for shell command",
			string(args[0].Type()),
			`os.run("ls -la") -> returns command output`,
		)
	}
	
	if strings.TrimSpace(cmd.Value) == "" {
		return &object.Error{
			Message: "\033[1;31mError in os.run()\033[0m:\n" +
				"  Cannot execute an empty command.\n" +
				"  Please provide a valid shell command.\n" +
				"  Usage: os.run(\"ls -la\") -> returns command output\n",
		}
	}	cmdParts := strings.Split(cmd.Value, " ")
	command := cmdParts[0]
	cmdArgs := cmdParts[1:]

	out, err := exec.Command(command, cmdArgs...).Output()
	if err != nil {
		exitErr, isExitError := err.(*exec.ExitError)
		if isExitError {
			return &object.Error{Message: fmt.Sprintf("os.run() failed to execute '%s': command exited with non-zero status %d. This usually indicates the command encountered an error.", cmd.Value, exitErr.ExitCode())}
		}
		return &object.Error{Message: fmt.Sprintf("os.run() failed to execute '%s': %v. Please check if the command exists and is executable.", cmd.Value, err)}
	}

	return &object.String{Value: string(out)}
}

func getEnv(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("os.getEnv() expects exactly 1 argument (environment variable name), but received %d. Usage: os.getEnv(\"PATH\")", len(args))}
	}

	key, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("os.getEnv() expects a string argument, but received %s. Usage: os.getEnv(\"PATH\")", args[0].Type())}
	}

	if strings.TrimSpace(key.Value) == "" {
		return &object.Error{Message: "os.getEnv() cannot retrieve an environment variable with an empty name. Please provide a valid variable name."}
	}

	value := os.Getenv(key.Value)
	return &object.String{Value: value}
}

func setEnv(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	key, ok1 := args[0].(*object.String)
	value, ok2 := args[1].(*object.String)
	if !ok1 || !ok2 {
		return &object.Error{Message: "Arguments must be strings"}
	}

	err := os.Setenv(key.Value, value.Value)
	if err != nil {
		return &object.Error{Message: "Failed to set environment variable: " + err.Error()}
	}

	return &object.String{Value: "Environment variable set successfully"}
}

func readFile(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Argument must be a string"}
	}

	content, err := ioutil.ReadFile(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to read file: " + err.Error()}
	}

	return &object.String{Value: string(content)}
}

func writeFile(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	path, ok1 := args[0].(*object.String)
	content, ok2 := args[1].(*object.String)
	if !ok1 || !ok2 {
		return &object.Error{Message: "Arguments must be strings"}
	}

	err := ioutil.WriteFile(path.Value, []byte(content.Value), 0644)
	if err != nil {
		return &object.Error{Message: "Failed to write file: " + err.Error()}
	}

	return &object.String{Value: "File written successfully"}
}

func listDir(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Argument must be a string"}
	}

	files, err := ioutil.ReadDir(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to list directory: " + err.Error()}
	}

	var fileList []string
	for _, file := range files {
		fileList = append(fileList, file.Name())
	}

	return &object.String{Value: strings.Join(fileList, ", ")}
}

func deleteFile(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Argument must be a string"}
	}

	err := os.Remove(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to delete file: " + err.Error()}
	}

	return &object.String{Value: "File deleted successfully"}
}

// This makeDir method Still has an issue with what
// path the new dir is saved will fix this
func makeDir(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Argument must be a string"}
	}

	err := os.Mkdir(path.Value, 0755)
	if err != nil {
		return &object.Error{Message: "Failed to create directory: " + err.Error()}
	}

	return &object.String{Value: "Directory created successfully"}
}

func removeDir(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Argument must be a string"}
	}

	err := os.RemoveAll(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to remove directory: " + err.Error()}
	}

	return &object.String{Value: "Directory removed successfully"}
}

func currentDir(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	dir, err := os.Getwd()
	if err != nil {
		return &object.Error{Message: "Failed to get current directory: " + err.Error()}
	}

	return &object.String{Value: dir}
}

func changeDir(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Argument must be a string"}
	}

	err := os.Chdir(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to change directory: " + err.Error()}
	}

	return &object.String{Value: "Directory changed successfully"}
}

func fileExists(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Argument must be a string"}
	}

	if _, err := os.Stat(path.Value); os.IsNotExist(err) {
		return &object.Boolean{Value: false}
	}

	return &object.Boolean{Value: true}
}

func readLines(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "Incorrect number of arguments"}
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Argument must be a string"}
	}

	content, err := ioutil.ReadFile(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to read file: " + err.Error()}
	}

	lines := strings.Split(string(content), "\n")
	lineObjects := make([]object.Object, len(lines))
	for i, line := range lines {
		lineObjects[i] = &object.String{Value: line}
	}

	return &object.Array{Elements: lineObjects}
}
