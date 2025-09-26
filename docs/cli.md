
# CLI Module

The `cli` module provides a comprehensive set of tools for building command-line applications in VintLang. It allows you to parse arguments, handle flags, prompt for user input, execute external commands, and more.

## Functions

### `getArgs()`

Returns an array of all command-line arguments passed to the script.

**Usage:**

```vint
import cli

let allArgs = cli.getArgs()
println("All arguments:", allArgs)
```

### `getFlags()`

Parses command-line arguments and returns a dictionary of flags (arguments starting with `--`). If a flag is followed by a value that doesn't start with `-`, it's treated as the flag's value. Otherwise, the flag's value is `true`.

**Usage:**

```vint
import cli

// Command: vint my_script.vint --verbose --output "file.txt"
let flags = cli.getFlags()
println("Flags:", flags)
// Output: Flags: {"verbose": true, "output": "file.txt"}
```

### `getPositional()`

Returns an array of positional arguments (arguments that are not flags or their values).

**Usage:**

```vint
import cli

// Command: vint my_script.vint my_file.txt --verbose
let positional = cli.getPositional()
println("Positional arguments:", positional)
// Output: Positional arguments: ["my_file.txt"]
```

### `getArgValue(flagName)`

Gets the value of a named argument (flag). It supports both `--flag=value` and `--flag value` formats.

- `flagName` (string): The name of the flag to get the value for (e.g., `"--output"`).

**Usage:**

```vint
import cli

// Command: vint my_script.vint --output="report.txt"
let outputFile = cli.getArgValue("--output")
println("Output file:", outputFile) // "report.txt"
```

### `hasArg(flagName)`

Checks if a named argument (flag) is present in the command-line arguments.

- `flagName` (string): The name of the flag to check for (e.g., `"--verbose"`).

**Usage:**

```vint
import cli

// Command: vint my_script.vint --verbose
if (cli.hasArg("--verbose")) {
    println("Verbose mode enabled.")
}
```

### `parse()`

A more advanced argument parser that returns a dictionary containing parsed flags, positional arguments, and helper methods (`has`, `get`, `positional`).

**Usage:**

```vint
import cli

// Command: vint my_script.vint --input="data.csv" process
let args = cli.parse()

println("Flags:", args.flags)
println("Positional:", args.positional())

if (args.has("--input")) {
    println("Input file:", args.get("--input"))
}
```

### `prompt(message)`

Displays a message to the user and waits for them to enter a line of text.

- `message` (string): The prompt message to display.

**Usage:**

```vint
import cli

let name = cli.prompt("Enter your name: ")
println("Hello, " + name)
```

### `confirm(message)`

Asks the user a yes/no question and returns `true` for "yes" and `false` for "no".

- `message` (string): The confirmation message to display.

**Usage:**

```vint
import cli

if (cli.confirm("Are you sure you want to continue?")) {
    println("Proceeding...")
} else {
    println("Operation cancelled.")
}
```

### `execCommand(command)`

Executes a shell command and returns its combined standard output and standard error.

- `command` (string): The command to execute.

**Usage:**t

```vint
import cli

let files = cli.execCommand("ls -l")
println(files)
```

### `exit(statusCode)`

Terminates the script with a given status code.

- `statusCode` (integer): The exit status code (0 for success, non-zero for error).

**Usage:**

```vint
import cli

if (error) {
    cli.exit(1)
}
```

### `help(appName, description)`

Generates and prints a standard help message for a CLI application.

- `appName` (string, optional): The name of the application.
- `description` (string, optional): A brief description of the application.

**Usage:**

```vint
import cli

cli.help("My Awesome App", "This app does awesome things.")
```

### `version(appName, version)`

Prints version information for the CLI application.

- `appName` (string, optional): The name of the application.
- `version` (string, optional): The version number.

**Usage:**

```vint
import cli

cli.version("My Awesome App", "1.0.0")
```
