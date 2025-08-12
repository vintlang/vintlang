# Command Line Argument Parsing in Vint

Vint now provides built-in support for parsing command line arguments, making it easy to build CLI applications.

## Basic Usage

### Get All Arguments

```js
// Get all command line arguments as an array
let allArgs = args()
print(allArgs)  // ["input.txt", "--verbose", "--output", "result.txt"]
```

### Using the CLI Module

```js
import cli

// Get all arguments (same as args())
let cliArgs = cli.args()

// Get only positional arguments (non-flags)
let positional = cli.getPositional()
print(positional)  // ["input.txt", "extra_file.txt"]

// Get flags as a dictionary
let flags = cli.getFlags()
print(flags)  // {verbose: true, output: "result.txt"}

// Check if a specific flag is present
if (cli.hasArg("--verbose")) {
    print("Verbose mode enabled")
}

// Get the value of a specific flag
let outputFile = cli.getArgValue("--output")
if (outputFile) {
    print("Output file:", outputFile)
}
```

## Supported Flag Formats

- `--flag` - Boolean flags (sets to true)
- `--flag value` - Flag with value (space-separated)
- `--flag=value` - Flag with value (equals-separated)
- `-v` - Short flags (boolean)

## Complete Example

```js
// example: vint myapp.vint input.txt --output result.txt --verbose --format json

import cli

// Check for help
if (cli.hasArg("--help")) {
    print("Usage: myapp <input> [options]")
    exit(0)
}

// Get input file from positional arguments
let positional = cli.getPositional()
let inputFile = positional[0] || "stdin"

// Get options
let outputFile = cli.getArgValue("--output") || "stdout"
let format = cli.getArgValue("--format") || "txt"
let verbose = cli.hasArg("--verbose")

if (verbose) {
    print("Processing", inputFile, "->", outputFile, "in", format, "format")
}
```

## Available Functions

### Built-in Functions
- `args()` - Returns array of all command line arguments

### CLI Module Functions
- `cli.args()` - Same as `args()`
- `cli.getPositional()` - Returns array of positional arguments
- `cli.getFlags()` - Returns dictionary of flags
- `cli.hasArg(flag)` - Checks if flag is present
- `cli.getArgValue(flag)` - Gets value of flag
- `cli.prompt(message)` - Prompts user for input
- `cli.confirm(message)` - Prompts for yes/no confirmation
- `cli.execCommand(cmd)` - Executes shell command
- `cli.cliExit(code)` - Exits with status code

## Notes

- Positional arguments are those that don't start with `-` or `--`
- Flag values are automatically detected (space or equals separated)
- Boolean flags are set to `true` when present
- The `args()` builtin and `cli.args()` return the same result
- All functions handle quotes in arguments properly
