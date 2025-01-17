# Shell Module in Vint

The Shell module in Vint provides a way to interact with the system's shell environment, allowing you to execute commands and check the existence of commands or files.

---

## Importing the Shell Module

To use the Shell module, import it as follows:

```js
import shell
```

---

## Functions and Examples

### 1. Run a Shell Command (`run`)
The `run` function allows you to execute a shell command and capture its output.

**Syntax**:
```js
run(command)
```

- `command` (string): The shell command to execute (e.g., `echo Hello` or `ls`).

**Example**:
```js
import shell

output = shell.run("echo Hello, Shell!")
print(output)
// Output: "Hello, Shell!"
```

In the example, the `echo` command prints the string `Hello, Shell!` to the terminal, and the output is captured by `shell.run` and printed.

---

### 2. Check if a Command Exists (`exists`)
The `exists` function checks whether a given command is available on the system.

**Syntax**:
```js
exists(command)
```

- `command` (string): The name of the command to check (e.g., `ls`, `python`, `echo`).

**Example**:
```js
import shell

exists_ls = shell.exists("ls")
print(exists_ls)
// Output: true if the 'ls' command exists

exists_python = shell.exists("python")
print(exists_python)
// Output: true if the 'python' command exists

exists_nonexistent = shell.exists("nonexistent_command")
print(exists_nonexistent)
// Output: false if the 'nonexistent_command' does not exist
```

In the example, `exists("ls")` checks if the `ls` command is available, returning `true` if it exists, and `false` otherwise.

---

### 3. Running Commands with Parameters
You can also pass arguments to shell commands within the `run` function.

**Example**:
```js
import shell

output = shell.run("ls -l")
print(output)
// Output: The list of files and directories in the current directory with detailed info.
```

This runs the `ls` command with the `-l` flag to list files and directories with additional details (e.g., permissions, size, etc.).

---

### Summary of Functions

| Function          | Description                                    | Example Output                              |
|-------------------|------------------------------------------------|---------------------------------------------|
| `run`             | Executes a shell command and returns its output. | `"Hello, Shell!"`                           |
| `exists`          | Checks if a command is available in the shell.  | `true` or `false` depending on command existence |

---

The Shell module is a powerful tool for integrating system-level shell commands directly within your Vint programs. It is especially useful for automation tasks, system monitoring, or running external scripts.