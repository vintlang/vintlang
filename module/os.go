package module

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

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

	// File permissions and ownership
	OsFunctions["chmod"] = chmod
	OsFunctions["chown"] = chown
	OsFunctions["lchown"] = lchown
	OsFunctions["chtimes"] = chtimes

	// Process information
	OsFunctions["getpid"] = getpid
	OsFunctions["getppid"] = getppid
	OsFunctions["getuid"] = getuid
	OsFunctions["getgid"] = getgid
	OsFunctions["geteuid"] = geteuid
	OsFunctions["getegid"] = getegid
	OsFunctions["getgroups"] = getgroups
	OsFunctions["getpagesize"] = getpagesize

	// Environment functions
	OsFunctions["environ"] = environ
	OsFunctions["clearenv"] = clearenv
	OsFunctions["lookupEnv"] = lookupEnv
	OsFunctions["expand"] = expand
	OsFunctions["expandEnv"] = expandEnv

	// File/directory info and operations
	OsFunctions["stat"] = stat
	OsFunctions["lstat"] = lstat
	OsFunctions["readDir"] = readDir
	OsFunctions["isExist"] = isExist
	OsFunctions["isNotExist"] = isNotExist
	OsFunctions["isPermission"] = isPermission
	OsFunctions["isTimeout"] = isTimeout
	OsFunctions["sameFile"] = sameFile
	OsFunctions["isPathSeparator"] = isPathSeparator

	// Links
	OsFunctions["link"] = link
	OsFunctions["symlink"] = symlink
	OsFunctions["readlink"] = readlink

	OsFunctions["truncate"] = truncate
	OsFunctions["mkdirAll"] = mkdirAll
	OsFunctions["mkdirTemp"] = mkdirTemp
	OsFunctions["createTemp"] = createTemp
	OsFunctions["executable"] = executable
	OsFunctions["rename"] = rename
	OsFunctions["remove"] = remove
	OsFunctions["removeAll"] = removeAll

	// User directories
	OsFunctions["userCacheDir"] = userCacheDir
	OsFunctions["userConfigDir"] = userConfigDir
	OsFunctions["userHomeDir"] = userHomeDir

	OsFunctions["tempDir"] = tempDir
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

	entries, err := ioutil.ReadDir(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to list directory: " + err.Error()}
	}

	var fileObjects []object.VintObject
	for _, entry := range entries {
		// We Skip directories, and only include files
		if !entry.IsDir() {
			fileObjects = append(fileObjects, &object.String{Value: entry.Name()})
		}
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

// File permissions and ownership functions

// chmod changes file mode/permissions
func chmod(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "chmod",
			"2 arguments (file path, mode)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.chmod("file.txt", 0o644)`,
		)
	}

	path, ok1 := args[0].(*object.String)
	mode, ok2 := args[1].(*object.Integer)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "chmod",
			"string and integer arguments for path and mode",
			fmt.Sprintf("path: %s, mode: %s", args[0].Type(), args[1].Type()),
			`os.chmod("file.txt", 0o644)`,
		)
	}

	err := os.Chmod(path.Value, os.FileMode(mode.Value))
	if err != nil {
		return &object.Error{Message: "Failed to change file mode: " + err.Error()}
	}

	return &object.String{Value: "File mode changed successfully"}
}

// chown changes file owner and group
func chown(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 {
		return ErrorMessage(
			"os", "chown",
			"3 arguments (file path, uid, gid)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.chown("file.txt", 1000, 1000)`,
		)
	}

	path, ok1 := args[0].(*object.String)
	uid, ok2 := args[1].(*object.Integer)
	gid, ok3 := args[2].(*object.Integer)
	if !ok1 || !ok2 || !ok3 {
		return ErrorMessage(
			"os", "chown",
			"string and integer arguments for path, uid, and gid",
			fmt.Sprintf("path: %s, uid: %s, gid: %s", args[0].Type(), args[1].Type(), args[2].Type()),
			`os.chown("file.txt", 1000, 1000)`,
		)
	}

	err := os.Chown(path.Value, int(uid.Value), int(gid.Value))
	if err != nil {
		return &object.Error{Message: "Failed to change file owner: " + err.Error()}
	}

	return &object.String{Value: "File owner changed successfully"}
}

// lchown changes symlink owner and group (doesn't follow the link)
func lchown(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 {
		return ErrorMessage(
			"os", "lchown",
			"3 arguments (file path, uid, gid)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.lchown("symlink", 1000, 1000)`,
		)
	}

	path, ok1 := args[0].(*object.String)
	uid, ok2 := args[1].(*object.Integer)
	gid, ok3 := args[2].(*object.Integer)
	if !ok1 || !ok2 || !ok3 {
		return ErrorMessage(
			"os", "lchown",
			"string and integer arguments for path, uid, and gid",
			fmt.Sprintf("path: %s, uid: %s, gid: %s", args[0].Type(), args[1].Type(), args[2].Type()),
			`os.lchown("symlink", 1000, 1000)`,
		)
	}

	err := os.Lchown(path.Value, int(uid.Value), int(gid.Value))
	if err != nil {
		return &object.Error{Message: "Failed to change symlink owner: " + err.Error()}
	}

	return &object.String{Value: "Symlink owner changed successfully"}
}

// chtimes changes file access and modification times
func chtimes(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 3 {
		return ErrorMessage(
			"os", "chtimes",
			"3 arguments (file path, access time, modification time)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.chtimes("file.txt", 1640995200, 1640995200)`,
		)
	}

	path, ok1 := args[0].(*object.String)
	atime, ok2 := args[1].(*object.Integer)
	mtime, ok3 := args[2].(*object.Integer)
	if !ok1 || !ok2 || !ok3 {
		return ErrorMessage(
			"os", "chtimes",
			"string and integer arguments for path, atime, and mtime",
			fmt.Sprintf("path: %s, atime: %s, mtime: %s", args[0].Type(), args[1].Type(), args[2].Type()),
			`os.chtimes("file.txt", 1640995200, 1640995200)`,
		)
	}

	atimeVal := time.Unix(atime.Value, 0)
	mtimeVal := time.Unix(mtime.Value, 0)

	err := os.Chtimes(path.Value, atimeVal, mtimeVal)
	if err != nil {
		return &object.Error{Message: "Failed to change file times: " + err.Error()}
	}

	return &object.String{Value: "File times changed successfully"}
}

// Process information functions

// getpid returns the process ID
func getpid(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "getpid",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.getpid()",
		)
	}

	return &object.Integer{Value: int64(os.Getpid())}
}

// getppid returns the parent process ID
func getppid(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "getppid",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.getppid()",
		)
	}

	return &object.Integer{Value: int64(os.Getppid())}
}

// getuid returns the user ID
func getuid(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "getuid",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.getuid()",
		)
	}

	return &object.Integer{Value: int64(os.Getuid())}
}

// getgid returns the group ID
func getgid(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "getgid",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.getgid()",
		)
	}

	return &object.Integer{Value: int64(os.Getgid())}
}

// geteuid returns the effective user ID
func geteuid(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "geteuid",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.geteuid()",
		)
	}

	return &object.Integer{Value: int64(os.Geteuid())}
}

// getegid returns the effective group ID
func getegid(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "getegid",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.getegid()",
		)
	}

	return &object.Integer{Value: int64(os.Getegid())}
}

// getgroups returns the list of supplemental group IDs
func getgroups(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "getgroups",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.getgroups()",
		)
	}

	groups, err := os.Getgroups()
	if err != nil {
		return &object.Error{Message: "Failed to get groups: " + err.Error()}
	}

	groupObjects := make([]object.VintObject, len(groups))
	for i, group := range groups {
		groupObjects[i] = &object.Integer{Value: int64(group)}
	}

	return &object.Array{Elements: groupObjects}
}

// getpagesize returns the system page size
func getpagesize(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "getpagesize",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.getpagesize()",
		)
	}

	return &object.Integer{Value: int64(os.Getpagesize())}
}

// Environment functions

// environ returns all environment variables
func environ(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "environ",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.environ()",
		)
	}

	envVars := os.Environ()
	envObjects := make([]object.VintObject, len(envVars))
	for i, env := range envVars {
		envObjects[i] = &object.String{Value: env}
	}

	return &object.Array{Elements: envObjects}
}

// clearenv clears all environment variables
func clearenv(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "clearenv",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.clearenv()",
		)
	}

	os.Clearenv()
	return &object.String{Value: "Environment cleared successfully"}
}

// lookupEnv looks up an environment variable and returns whether it exists
func lookupEnv(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "lookupEnv",
			"1 string argument (environment variable name)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.lookupEnv("PATH")`,
		)
	}

	key, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "lookupEnv",
			"string argument for environment variable name",
			string(args[0].Type()),
			`os.lookupEnv("PATH")`,
		)
	}

	value, exists := os.LookupEnv(key.Value)
	valuePair := object.DictPair{
		Key:   &object.String{Value: "value"},
		Value: &object.String{Value: value},
	}
	existsPair := object.DictPair{
		Key:   &object.String{Value: "exists"},
		Value: &object.Boolean{Value: exists},
	}

	pairs := map[object.HashKey]object.DictPair{
		(&object.String{Value: "value"}).HashKey():  valuePair,
		(&object.String{Value: "exists"}).HashKey(): existsPair,
	}

	return &object.Dict{Pairs: pairs}
}

// expand expands variables in string using provided mapping function
func expand(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "expand",
			"1 string argument (string with variables)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.expand("$HOME/file.txt")`,
		)
	}

	str, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "expand",
			"string argument",
			string(args[0].Type()),
			`os.expand("$HOME/file.txt")`,
		)
	}

	expanded := os.Expand(str.Value, os.Getenv)
	return &object.String{Value: expanded}
}

// expandEnv expands environment variables in string
func expandEnv(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "expandEnv",
			"1 string argument (string with environment variables)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.expandEnv("$HOME/file.txt")`,
		)
	}

	str, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "expandEnv",
			"string argument",
			string(args[0].Type()),
			`os.expandEnv("$HOME/file.txt")`,
		)
	}

	expanded := os.ExpandEnv(str.Value)
	return &object.String{Value: expanded}
}

// File/directory info and operations

// stat returns file info
func stat(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "stat",
			"1 string argument (file path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.stat("file.txt")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "stat",
			"string argument for file path",
			string(args[0].Type()),
			`os.stat("file.txt")`,
		)
	}

	info, err := os.Stat(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to stat file: " + err.Error()}
	}

	namePair := object.DictPair{Key: &object.String{Value: "name"}, Value: &object.String{Value: info.Name()}}
	sizePair := object.DictPair{Key: &object.String{Value: "size"}, Value: &object.Integer{Value: info.Size()}}
	modePair := object.DictPair{Key: &object.String{Value: "mode"}, Value: &object.Integer{Value: int64(info.Mode())}}
	modTimePair := object.DictPair{Key: &object.String{Value: "modTime"}, Value: &object.Integer{Value: info.ModTime().Unix()}}
	isDirPair := object.DictPair{Key: &object.String{Value: "isDir"}, Value: &object.Boolean{Value: info.IsDir()}}

	pairs := map[object.HashKey]object.DictPair{
		(&object.String{Value: "name"}).HashKey():    namePair,
		(&object.String{Value: "size"}).HashKey():    sizePair,
		(&object.String{Value: "mode"}).HashKey():    modePair,
		(&object.String{Value: "modTime"}).HashKey(): modTimePair,
		(&object.String{Value: "isDir"}).HashKey():   isDirPair,
	}

	return &object.Dict{Pairs: pairs}
}

// lstat returns file info (doesn't follow symlinks)
func lstat(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "lstat",
			"1 string argument (file path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.lstat("symlink")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "lstat",
			"string argument for file path",
			string(args[0].Type()),
			`os.lstat("symlink")`,
		)
	}

	info, err := os.Lstat(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to lstat file: " + err.Error()}
	}

	namePair := object.DictPair{Key: &object.String{Value: "name"}, Value: &object.String{Value: info.Name()}}
	sizePair := object.DictPair{Key: &object.String{Value: "size"}, Value: &object.Integer{Value: info.Size()}}
	modePair := object.DictPair{Key: &object.String{Value: "mode"}, Value: &object.Integer{Value: int64(info.Mode())}}
	modTimePair := object.DictPair{Key: &object.String{Value: "modTime"}, Value: &object.Integer{Value: info.ModTime().Unix()}}
	isDirPair := object.DictPair{Key: &object.String{Value: "isDir"}, Value: &object.Boolean{Value: info.IsDir()}}

	pairs := map[object.HashKey]object.DictPair{
		(&object.String{Value: "name"}).HashKey():    namePair,
		(&object.String{Value: "size"}).HashKey():    sizePair,
		(&object.String{Value: "mode"}).HashKey():    modePair,
		(&object.String{Value: "modTime"}).HashKey(): modTimePair,
		(&object.String{Value: "isDir"}).HashKey():   isDirPair,
	}

	return &object.Dict{Pairs: pairs}
}

// readDir reads directory contents
func readDir(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "readDir",
			"1 string argument (directory path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.readDir("/path/to/directory")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "readDir",
			"string argument for directory path",
			string(args[0].Type()),
			`os.readDir("/path/to/directory")`,
		)
	}

	entries, err := os.ReadDir(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to read directory: " + err.Error()}
	}

	entryObjects := make([]object.VintObject, len(entries))
	for i, entry := range entries {
		info, _ := entry.Info()

		namePair := object.DictPair{Key: &object.String{Value: "name"}, Value: &object.String{Value: entry.Name()}}
		isDirPair := object.DictPair{Key: &object.String{Value: "isDir"}, Value: &object.Boolean{Value: entry.IsDir()}}

		pairs := map[object.HashKey]object.DictPair{
			(&object.String{Value: "name"}).HashKey():  namePair,
			(&object.String{Value: "isDir"}).HashKey(): isDirPair,
		}

		if info != nil {
			sizePair := object.DictPair{Key: &object.String{Value: "size"}, Value: &object.Integer{Value: info.Size()}}
			modePair := object.DictPair{Key: &object.String{Value: "mode"}, Value: &object.Integer{Value: int64(info.Mode())}}
			modTimePair := object.DictPair{Key: &object.String{Value: "modTime"}, Value: &object.Integer{Value: info.ModTime().Unix()}}

			pairs[(&object.String{Value: "size"}).HashKey()] = sizePair
			pairs[(&object.String{Value: "mode"}).HashKey()] = modePair
			pairs[(&object.String{Value: "modTime"}).HashKey()] = modTimePair
		}

		entryObjects[i] = &object.Dict{Pairs: pairs}
	}

	return &object.Array{Elements: entryObjects}
}

// Error checking functions
func isExist(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "isExist",
			"1 string argument (error message)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.isExist("file already exists")`,
		)
	}

	errMsg, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "isExist",
			"string argument for error message",
			string(args[0].Type()),
			`os.isExist("file already exists")`,
		)
	}

	// Simple check for common "exists" error messages
	exists := strings.Contains(strings.ToLower(errMsg.Value), "exist") &&
		!strings.Contains(strings.ToLower(errMsg.Value), "not exist")

	return &object.Boolean{Value: exists}
}

func isNotExist(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "isNotExist",
			"1 string argument (error message)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.isNotExist("no such file")`,
		)
	}

	errMsg, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "isNotExist",
			"string argument for error message",
			string(args[0].Type()),
			`os.isNotExist("no such file")`,
		)
	}

	// Simple check for common "not exists" error messages
	notExists := strings.Contains(strings.ToLower(errMsg.Value), "not exist") ||
		strings.Contains(strings.ToLower(errMsg.Value), "no such")

	return &object.Boolean{Value: notExists}
}

func isPermission(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "isPermission",
			"1 string argument (error message)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.isPermission("permission denied")`,
		)
	}

	errMsg, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "isPermission",
			"string argument for error message",
			string(args[0].Type()),
			`os.isPermission("permission denied")`,
		)
	}

	// Simple check for common permission error messages
	permissionDenied := strings.Contains(strings.ToLower(errMsg.Value), "permission") ||
		strings.Contains(strings.ToLower(errMsg.Value), "denied") ||
		strings.Contains(strings.ToLower(errMsg.Value), "access")

	return &object.Boolean{Value: permissionDenied}
}

func isTimeout(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "isTimeout",
			"1 string argument (error message)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.isTimeout("timeout occurred")`,
		)
	}

	errMsg, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "isTimeout",
			"string argument for error message",
			string(args[0].Type()),
			`os.isTimeout("timeout occurred")`,
		)
	}

	// Simple check for common timeout error messages
	timeout := strings.Contains(strings.ToLower(errMsg.Value), "timeout") ||
		strings.Contains(strings.ToLower(errMsg.Value), "deadline")

	return &object.Boolean{Value: timeout}
}

func sameFile(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "sameFile",
			"2 string arguments (file paths)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.sameFile("file1.txt", "file2.txt")`,
		)
	}

	path1, ok1 := args[0].(*object.String)
	path2, ok2 := args[1].(*object.String)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "sameFile",
			"string arguments for file paths",
			fmt.Sprintf("path1: %s, path2: %s", args[0].Type(), args[1].Type()),
			`os.sameFile("file1.txt", "file2.txt")`,
		)
	}

	info1, err1 := os.Stat(path1.Value)
	info2, err2 := os.Stat(path2.Value)

	if err1 != nil || err2 != nil {
		return &object.Boolean{Value: false}
	}

	same := os.SameFile(info1, info2)
	return &object.Boolean{Value: same}
}

func isPathSeparator(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "isPathSeparator",
			"1 string argument (character)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.isPathSeparator("/")`,
		)
	}

	char, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "isPathSeparator",
			"string argument for character",
			string(args[0].Type()),
			`os.isPathSeparator("/")`,
		)
	}

	if len(char.Value) != 1 {
		return &object.Boolean{Value: false}
	}

	isSep := os.IsPathSeparator(char.Value[0])
	return &object.Boolean{Value: isSep}
}

// Link functions

func link(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "link",
			"2 string arguments (oldname, newname)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.link("file.txt", "link.txt")`,
		)
	}

	oldname, ok1 := args[0].(*object.String)
	newname, ok2 := args[1].(*object.String)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "link",
			"string arguments for oldname and newname",
			fmt.Sprintf("oldname: %s, newname: %s", args[0].Type(), args[1].Type()),
			`os.link("file.txt", "link.txt")`,
		)
	}

	err := os.Link(oldname.Value, newname.Value)
	if err != nil {
		return &object.Error{Message: "Failed to create hard link: " + err.Error()}
	}

	return &object.String{Value: "Hard link created successfully"}
}

func symlink(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "symlink",
			"2 string arguments (oldname, newname)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.symlink("file.txt", "symlink.txt")`,
		)
	}

	oldname, ok1 := args[0].(*object.String)
	newname, ok2 := args[1].(*object.String)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "symlink",
			"string arguments for oldname and newname",
			fmt.Sprintf("oldname: %s, newname: %s", args[0].Type(), args[1].Type()),
			`os.symlink("file.txt", "symlink.txt")`,
		)
	}

	err := os.Symlink(oldname.Value, newname.Value)
	if err != nil {
		return &object.Error{Message: "Failed to create symbolic link: " + err.Error()}
	}

	return &object.String{Value: "Symbolic link created successfully"}
}

func readlink(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "readlink",
			"1 string argument (symlink path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.readlink("symlink.txt")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "readlink",
			"string argument for symlink path",
			string(args[0].Type()),
			`os.readlink("symlink.txt")`,
		)
	}

	target, err := os.Readlink(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to read symbolic link: " + err.Error()}
	}

	return &object.String{Value: target}
}

// Advanced file operations

func truncate(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "truncate",
			"2 arguments (file path, size)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.truncate("file.txt", 100)`,
		)
	}

	path, ok1 := args[0].(*object.String)
	size, ok2 := args[1].(*object.Integer)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "truncate",
			"string and integer arguments for path and size",
			fmt.Sprintf("path: %s, size: %s", args[0].Type(), args[1].Type()),
			`os.truncate("file.txt", 100)`,
		)
	}

	err := os.Truncate(path.Value, size.Value)
	if err != nil {
		return &object.Error{Message: "Failed to truncate file: " + err.Error()}
	}

	return &object.String{Value: "File truncated successfully"}
}

func mkdirAll(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "mkdirAll",
			"2 arguments (directory path, mode)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.mkdirAll("/path/to/nested/dir", 0o755)`,
		)
	}

	path, ok1 := args[0].(*object.String)
	mode, ok2 := args[1].(*object.Integer)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "mkdirAll",
			"string and integer arguments for path and mode",
			fmt.Sprintf("path: %s, mode: %s", args[0].Type(), args[1].Type()),
			`os.mkdirAll("/path/to/nested/dir", 0o755)`,
		)
	}

	err := os.MkdirAll(path.Value, os.FileMode(mode.Value))
	if err != nil {
		return &object.Error{Message: "Failed to create directory tree: " + err.Error()}
	}

	return &object.String{Value: "Directory tree created successfully"}
}

func mkdirTemp(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "mkdirTemp",
			"2 string arguments (dir, pattern)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.mkdirTemp("", "temp-*")`,
		)
	}

	dir, ok1 := args[0].(*object.String)
	pattern, ok2 := args[1].(*object.String)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "mkdirTemp",
			"string arguments for dir and pattern",
			fmt.Sprintf("dir: %s, pattern: %s", args[0].Type(), args[1].Type()),
			`os.mkdirTemp("", "temp-*")`,
		)
	}

	tempDir, err := os.MkdirTemp(dir.Value, pattern.Value)
	if err != nil {
		return &object.Error{Message: "Failed to create temporary directory: " + err.Error()}
	}

	return &object.String{Value: tempDir}
}

func createTemp(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "createTemp",
			"2 string arguments (dir, pattern)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.createTemp("", "temp-*.txt")`,
		)
	}

	dir, ok1 := args[0].(*object.String)
	pattern, ok2 := args[1].(*object.String)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "createTemp",
			"string arguments for dir and pattern",
			fmt.Sprintf("dir: %s, pattern: %s", args[0].Type(), args[1].Type()),
			`os.createTemp("", "temp-*.txt")`,
		)
	}

	file, err := os.CreateTemp(dir.Value, pattern.Value)
	if err != nil {
		return &object.Error{Message: "Failed to create temporary file: " + err.Error()}
	}

	// Close the file immediately and return just the name
	file.Close()

	return &object.String{Value: file.Name()}
}

func executable(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "executable",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.executable()",
		)
	}

	exec, err := os.Executable()
	if err != nil {
		return &object.Error{Message: "Failed to get executable path: " + err.Error()}
	}

	return &object.String{Value: exec}
}

func rename(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return ErrorMessage(
			"os", "rename",
			"2 string arguments (oldpath, newpath)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.rename("old.txt", "new.txt")`,
		)
	}

	oldpath, ok1 := args[0].(*object.String)
	newpath, ok2 := args[1].(*object.String)
	if !ok1 || !ok2 {
		return ErrorMessage(
			"os", "rename",
			"string arguments for oldpath and newpath",
			fmt.Sprintf("oldpath: %s, newpath: %s", args[0].Type(), args[1].Type()),
			`os.rename("old.txt", "new.txt")`,
		)
	}

	err := os.Rename(oldpath.Value, newpath.Value)
	if err != nil {
		return &object.Error{Message: "Failed to rename: " + err.Error()}
	}

	return &object.String{Value: "File renamed successfully"}
}

func remove(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "remove",
			"1 string argument (path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.remove("file.txt")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "remove",
			"string argument for path",
			string(args[0].Type()),
			`os.remove("file.txt")`,
		)
	}

	err := os.Remove(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to remove: " + err.Error()}
	}

	return &object.String{Value: "File removed successfully"}
}

func removeAll(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return ErrorMessage(
			"os", "removeAll",
			"1 string argument (path)",
			fmt.Sprintf("%d arguments", len(args)),
			`os.removeAll("directory")`,
		)
	}

	path, ok := args[0].(*object.String)
	if !ok {
		return ErrorMessage(
			"os", "removeAll",
			"string argument for path",
			string(args[0].Type()),
			`os.removeAll("directory")`,
		)
	}

	err := os.RemoveAll(path.Value)
	if err != nil {
		return &object.Error{Message: "Failed to remove all: " + err.Error()}
	}

	return &object.String{Value: "Path removed successfully"}
}

// User directory functions

func userCacheDir(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "userCacheDir",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.userCacheDir()",
		)
	}

	dir, err := os.UserCacheDir()
	if err != nil {
		return &object.Error{Message: "Failed to get user cache directory: " + err.Error()}
	}

	return &object.String{Value: dir}
}

func userConfigDir(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "userConfigDir",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.userConfigDir()",
		)
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return &object.Error{Message: "Failed to get user config directory: " + err.Error()}
	}

	return &object.String{Value: dir}
}

func userHomeDir(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "userHomeDir",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.userHomeDir()",
		)
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		return &object.Error{Message: "Failed to get user home directory: " + err.Error()}
	}

	return &object.String{Value: dir}
}

// Path utility functions

func tempDir(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 0 {
		return ErrorMessage(
			"os", "tempDir",
			"no arguments",
			fmt.Sprintf("%d arguments", len(args)),
			"os.tempDir()",
		)
	}

	return &object.String{Value: os.TempDir()}
}
