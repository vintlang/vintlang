package module

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/vintlang/vintlang/object"
)

var OsFunctions = map[string]object.ModuleFunction{}

func init() {
	OsFunctions["exit"] = exit
	OsFunctions["run"] = run
	OsFunctions["getEnv"] = getEnv
	OsFunctions["setEnv"] = setEnv
	OsFunctions["unsetEnv"] = unsetEnv
	OsFunctions["readFile"] = readFile
	OsFunctions["writeFile"] = writeFile
	OsFunctions["listDir"] = listDir
	OsFunctions["listFiles"] = listFiles
	OsFunctions["deleteFile"] = deleteFile
	OsFunctions["makeDir"] = makeDir
	OsFunctions["removeDir"] = removeDir
	OsFunctions["currentDir"] = currentDir
	OsFunctions["changeDir"] = changeDir
	OsFunctions["fileExists"] = fileExists
	OsFunctions["readLines"] = readLines
	OsFunctions["getwd"] = getwd
	OsFunctions["homedir"] = homedir
	OsFunctions["tmpdir"] = tmpdir
	OsFunctions["cpuCount"] = cpuCount
	OsFunctions["hostname"] = hostname
	OsFunctions["copy"] = copyFile
	OsFunctions["move"] = moveFile
}

func exit(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) > 1 {
		return ErrorMessage(
			"os", "exit",
			"0 or 1 argument (optional status code)",
			fmt.Sprintf("%d arguments", len(args)),
			"os.exit() or os.exit(0)",
		)
	}

	if len(args) == 1 {
		status, ok := args[0].(*object.Integer)
		if !ok {
			return ErrorMessage(
				"os", "exit",
				"integer argument for status code",
				string(args[0].Type()),
				"os.exit(0) or os.exit(1)",
			)
		}
		os.Exit(int(status.Value))
		return nil
	}

	os.Exit(0)

	return nil
}

func run(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
			Message: "\033[1;31m -> os.run()\033[0m:\n" +
				"  Cannot execute an empty command.\n" +
				"  Please provide a valid shell command.\n" +
				"  Usage: os.run(\"ls -la\") -> returns command output\n",
		}
	}

	cmdParts := strings.Split(cmd.Value, " ")
	command := cmdParts[0]
	cmdArgs := cmdParts[1:]

	out, err := exec.Command(command, cmdArgs...).Output()
	if err != nil {
		exitErr, isExitError := err.(*exec.ExitError)
		if isExitError {
			return &object.Error{
				Message: fmt.Sprintf("\033[1;31m -> os.run()\033[0m:\n"+
					"  Command '%s' exited with status %d.\n"+
					"  This usually indicates the command encountered an error.\n"+
					"  Usage: os.run(\"ls -la\") -> returns command output\n",
					cmd.Value, exitErr.ExitCode()),
			}
		}
		return &object.Error{
			Message: fmt.Sprintf("\033[1;31m -> os.run()\033[0m:\n"+
				"  Failed to execute '%s': %v\n"+
				"  Please check if the command exists and is executable.\n"+
				"  Usage: os.run(\"ls -la\") -> returns command output\n",
				cmd.Value, err),
		}
	}

	return &object.String{Value: string(out)}
}

func getEnv(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "getEnv",
			"1 string argument (environment variable name)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.getEnv("PATH") -> returns environment variable value`,
		)
	}

	key, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "getEnv",
			"string argument for environment variable name",
			string(args[0].Type()),
			`os.getEnv("PATH") -> returns environment variable value`,
		)
	}

	if strings.TrimSpace(key.Value) == "" {
		return &object.Error{
			Message: "\033[1;31m -> os.getEnv()\033[0m:\n" +
				"  Cannot retrieve an environment variable with an empty name.\n" +
				"  Please provide a valid variable name.\n" +
				"  Usage: os.getEnv(\"PATH\") -> returns environment variable value\n",
		}
	}

	value := os.Getenv(key.Value)
	return &object.String{Value: value}
}

func setEnv(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "setEnv",
			"2 string arguments (key, value)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.setEnv("PATH", "/usr/bin")`,
		)
	}

	key, ok1 := args[0].(*object.String)
	value, ok2 := args[1].(*object.String)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "setEnv",
			"2 string arguments for key and value",
			fmt.Sprintf("key: %s, value: %s", args[0].Type(), args[1].Type()),
			`os.setEnv("PATH", "/usr/bin")`,
		)
	}

	err := os.Setenv(key.Value, value.Value)
	if err != nil {
		return &object.Error{Message: "Failed to set environment variable: " + err.Error()}
	}

	return &object.String{Value: "Environment variable set successfully"}
}

func readFile(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "readFile",
			"1 string argument (file path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.readFile("file.txt")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "readFile",
			"string argument for file path",
			string(args[0].Type()),
			`os.readFile("file.txt")`,
		)
	}

	content, err := ioutil.ReadFile(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to read file: " + err.Error()}
	}

	return &object.String{Value: string(content)}
}

func getwd(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "getwd",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`os.getwd()`,
		)
	}

	dir, err := os.Getwd()
	if err != nil {
		return &object.Error{Message: "Failed to get current working directory: " + err.Error()}
	}
	return &object.String{Value: dir}
}

func writeFile(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "writeFile",
			"2 string arguments (file path, content)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.writeFile("file.txt", "content")`,
		)
	}

	path, ok1 := args[0].(*object.String)
	content, ok2 := args[1].(*object.String)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "writeFile",
			"2 string arguments for file path and content",
			fmt.Sprintf("path: %s, content: %s", args[0].Type(), args[1].Type()),
			`os.writeFile("file.txt", "content")`,
		)
	}

	err := ioutil.WriteFile(path.Value, []byte(content.Value), 0644)
	if err != nil {
		return &object.Error{Message: "Failed to write file: " + err.Error()}
	}

	return &object.String{Value: "File written successfully"}
}

func listDir(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "listDir",
			"1 string argument (directory path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.listDir("/path/to/directory")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "listDir",
			"string argument for directory path",
			string(args[0].Type()),
			`os.listDir("/path/to/directory")`,
		)
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

func listFiles(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "listFiles",
			"1 string argument (directory path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.listFiles("/path/to/directory")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "listFiles",
			"string argument for directory path",
			string(args[0].Type()),
			`os.listFiles("/path/to/directory")`,
		)
	}

	files, err := ioutil.ReadDir(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to list directory: " + err.Error()}
	}

	fileObjects := make([]object.VintObject, len(files))
	for i, file := range files {
		fileObjects[i] = &object.String{Value: file.Name()}
	}

	return &object.Array{Elements: fileObjects}
}

func deleteFile(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "deleteFile",
			"1 string argument (file path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.deleteFile("file.txt")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "deleteFile",
			"string argument for file path",
			string(args[0].Type()),
			`os.deleteFile("file.txt")`,
		)
	}

	err := os.Remove(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to delete file: " + err.Error()}
	}

	return &object.String{Value: "File deleted successfully"}
}

// This makeDir method Still has an issue with what
// path the new dir is saved will fix this
func makeDir(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "makeDir",
			"1 string argument (directory path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.makeDir("new_directory")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "makeDir",
			"string argument for directory path",
			string(args[0].Type()),
			`os.makeDir("new_directory")`,
		)
	}

	err := os.Mkdir(path.Value, 0755)
	if err != nil {
		return &object.Error{Message: "Failed to create directory: " + err.Error()}
	}

	return &object.String{Value: "Directory created successfully"}
}

func removeDir(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "removeDir",
			"1 string argument (directory path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.removeDir("directory")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "removeDir",
			"string argument for directory path",
			string(args[0].Type()),
			`os.removeDir("directory")`,
		)
	}

	err := os.RemoveAll(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to remove directory: " + err.Error()}
	}

	return &object.String{Value: "Directory removed successfully"}
}

func currentDir(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "currentDir",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.currentDir()",
		)
	}

	dir, err := os.Getwd()
	if err != nil {
		return &object.Error{Message: "Failed to get current directory: " + err.Error()}
	}

	return &object.String{Value: dir}
}

func changeDir(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "changeDir",
			"1 string argument (directory path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.changeDir("/path/to/directory")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "changeDir",
			"string argument for directory path",
			string(args[0].Type()),
			`os.changeDir("/path/to/directory")`,
		)
	}

	err := os.Chdir(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to change directory: " + err.Error()}
	}

	return &object.String{Value: "Directory changed successfully"}
}

func fileExists(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "fileExists",
			"1 string argument (file path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.fileExists("file.txt")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "fileExists",
			"string argument for file path",
			string(args[0].Type()),
			`os.fileExists("file.txt")`,
		)
	}

	if _, err := os.Stat(path.Value); os.IsNotExist(err) {
		return &object.Boolean{Value: false}
	}

	return &object.Boolean{Value: true}
}

func readLines(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "readLines",
			"1 string argument (file path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.readLines("file.txt")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "readLines",
			"string argument for file path",
			string(args[0].Type()),
			`os.readLines("file.txt")`,
		)
	}

	content, err := ioutil.ReadFile(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to read file: " + err.Error()}
	}

	lines := strings.Split(string(content), "\n")
	lineObjects := make([]object.VintObject, len(lines))
	for i, line := range lines {
		lineObjects[i] = &object.String{Value: line}
	}

	return &object.Array{Elements: lineObjects}
}

// unsetEnv removes an environment variable
func unsetEnv(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "unsetEnv",
			"1 string argument (environment variable name)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.unsetEnv("TEMP_VAR")`,
		)
	}

	key, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "unsetEnv",
			"string argument for environment variable name",
			string(args[0].Type()),
			`os.unsetEnv("TEMP_VAR")`,
		)
	}

	if strings.TrimSpace(key.Value) == "" {
		return &object.Error{
			Message: "\033[1;31m -> os.unsetEnv()\033[0m:\n" +
				"  Cannot unset an environment variable with an empty name.\n" +
				"  Please provide a valid variable name.\n" +
				"  Usage: os.unsetEnv(\"TEMP_VAR\")\n",
		}
	}

	err := os.Unsetenv(key.Value)
	if err != nil {
		return &object.Error{Message: "Failed to unset environment variable: " + err.Error()}
	}

	return &object.String{Value: "Environment variable unset successfully"}
}

// homedir returns the user's home directory
func homedir(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "homedir",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.homedir()",
		)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return &object.Error{Message: "Failed to get home directory: " + err.Error()}
	}

	return &object.String{Value: home}
}

// tmpdir returns the system's temporary directory
func tmpdir(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "tmpdir",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.tmpdir()",
		)
	}

	tmpdir := os.TempDir()
	return &object.String{Value: tmpdir}
}

// cpuCount returns the number of logical CPUs
func cpuCount(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "cpuCount",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.cpuCount()",
		)
	}

	count := runtime.NumCPU()
	return &object.Integer{Value: int64(count)}
}

// hostname returns the system's hostname
func hostname(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "hostname",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.hostname()",
		)
	}

	hostname, err := os.Hostname()
	if err != nil {
		return &object.Error{Message: "Failed to get hostname: " + err.Error()}
	}

	return &object.String{Value: hostname}
}

// copyFile copies a file from source to destination
func copyFile(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "copy",
			"2 string arguments (source path, destination path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.copy("source.txt", "destination.txt")`,
		)
	}

	source, ok1 := args[0].(*object.String)
	destination, ok2 := args[1].(*object.String)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "copy",
			"2 string arguments for source and destination paths",
			fmt.Sprintf("source: %s, destination: %s", args[0].Type(), args[1].Type()),
			`os.copy("source.txt", "destination.txt")`,
		)
	}

	sourceFile, err := os.Open(source.Value)
	if err != nil {
		return &object.Error{Message: "Failed to open source file: " + err.Error()}
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destination.Value)
	if err != nil {
		return &object.Error{Message: "Failed to create destination file: " + err.Error()}
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return &object.Error{Message: "Failed to copy file: " + err.Error()}
	}

	return &object.String{Value: "File copied successfully"}
}

// moveFile moves or renames a file from source to destination
func moveFile(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "move",
			"2 string arguments (source path, destination path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.move("old_name.txt", "new_name.txt")`,
		)
	}

	source, ok1 := args[0].(*object.String)
	destination, ok2 := args[1].(*object.String)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "move",
			"2 string arguments for source and destination paths",
			fmt.Sprintf("source: %s, destination: %s", args[0].Type(), args[1].Type()),
			`os.move("old_name.txt", "new_name.txt")`,
		)
	}

	err := os.Rename(source.Value, destination.Value)
	if err != nil {
		return &object.Error{Message: "Failed to move file: " + err.Error()}
	}

	return &object.String{Value: "File moved successfully"}
}
