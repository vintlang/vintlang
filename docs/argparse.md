# VintLang Argparse Module

The `argparse` module provides a powerful command-line argument parsing system for VintLang applications. It allows you to define command-line arguments and flags, generate help text, and parse command-line input.

## Features

- Define positional arguments with validation
- Define optional flags with short and long names
- Type checking and validation
- Automatic help text generation
- Version information support
- Default values for arguments and flags
- Required arguments and flags

## Usage

### Basic Example

```js
import argparse

// Create a new parser
let parser = argparse.newParser("myapp", "My awesome application")

// Add arguments
argparse.addArgument(parser, "input", {
    description: "Input file",
    required: true
})

// Add flags
argparse.addFlag(parser, "verbose", {
    short: "v",
    description: "Enable verbose output"
})

argparse.addFlag(parser, "output", {
    short: "o",
    description: "Output file",
    type: "string"
})

// Set version information
argparse.version(parser, "1.0.0")

// Parse arguments
let args = argparse.parse(parser)

// Access parsed arguments
let inputFile = args["input"]
let verbose = args["verbose"]
let outputFile = args["output"]

if (verbose) {
    print("Verbose mode enabled")
}

print("Input file:", inputFile)
if (outputFile) {
    print("Output file:", outputFile)
}
```

## API Reference

### newParser(name, description)

Creates a new argument parser.

**Parameters:**
- `name` (string): The name of the parser
- `description` (string, optional): A description of the application

**Returns:**
- A parser ID string that can be used with other functions

### addArgument(parser, name, options)

Adds a positional argument to the parser.

**Parameters:**
- `parser` (string): The parser ID
- `name` (string): The name of the argument
- `options` (dict, optional): Options for the argument
  - `description` (string): Description of the argument
  - `required` (boolean): Whether the argument is required (default: false)
  - `default` (any): Default value if the argument is not provided
  - `type` (string): Type of the argument ("string", "integer", "float", "boolean")
  - `choices` (array): List of valid values for the argument

**Returns:**
- `true` if the argument was added successfully

### addFlag(parser, name, options)

Adds a flag (optional named argument) to the parser.

**Parameters:**
- `parser` (string): The parser ID
- `name` (string): The name of the flag
- `options` (dict, optional): Options for the flag
  - `short` (string): Short name for the flag (single character)
  - `description` (string): Description of the flag
  - `required` (boolean): Whether the flag is required (default: false)
  - `default` (any): Default value if the flag is not provided
  - `type` (string): Type of the flag ("string", "integer", "float", "boolean")

**Returns:**
- `true` if the flag was added successfully

### parse(parser, args)

Parses command line arguments according to the parser definition.

**Parameters:**
- `parser` (string): The parser ID
- `args` (array, optional): Array of strings representing the arguments to parse. If not provided, the system arguments will be used.

**Returns:**
- A dictionary containing the parsed arguments and flags

### help(parser)

Generates help text for the parser.

**Parameters:**
- `parser` (string): The parser ID

**Returns:**
- A string containing the help text

### version(parser, versionString)

Sets the version information for the parser.

**Parameters:**
- `parser` (string): The parser ID
- `versionString` (string): The version string

**Returns:**
- `true` if the version was set successfully

## Examples

### Command-line Calculator

```js
import argparse

// Create a new parser
let parser = argparse.newParser("calc", "A simple command-line calculator")

// Add arguments
argparse.addArgument(parser, "operation", {
    description: "Operation to perform (add, subtract, multiply, divide)",
    choices: ["add", "subtract", "multiply", "divide"]
})

argparse.addArgument(parser, "a", {
    description: "First number",
    type: "float"
})

argparse.addArgument(parser, "b", {
    description: "Second number",
    type: "float"
})

// Add flags
argparse.addFlag(parser, "precision", {
    short: "p",
    description: "Number of decimal places",
    type: "integer",
    default: 2
})

// Set version information
argparse.version(parser, "1.0.0")

// Parse arguments
let args = argparse.parse(parser)

// Get values
let operation = args["operation"]
let a = args["a"]
let b = args["b"]
let precision = args["precision"]

// Perform calculation
let result = 0
if (operation == "add") {
    result = a + b
} else if (operation == "subtract") {
    result = a - b
} else if (operation == "multiply") {
    result = a * b
} else if (operation == "divide") {
    if (b == 0) {
        print("Error: Division by zero")
        exit(1)
    }
    result = a / b
}

// Format result with specified precision
print(result.toFixed(precision))
```

### File Processor

```js
import argparse
import os

// Create a new parser
let parser = argparse.newParser("fileproc", "A file processing utility")

// Add arguments
argparse.addArgument(parser, "input", {
    description: "Input file",
    required: true
})

// Add flags
argparse.addFlag(parser, "output", {
    short: "o",
    description: "Output file",
    type: "string"
})

argparse.addFlag(parser, "uppercase", {
    short: "u",
    description: "Convert to uppercase"
})

argparse.addFlag(parser, "lowercase", {
    short: "l",
    description: "Convert to lowercase"
})

argparse.addFlag(parser, "count", {
    short: "c",
    description: "Count lines, words, and characters"
})

// Parse arguments
let args = argparse.parse(parser)

// Get values
let inputFile = args["input"]
let outputFile = args["output"]
let uppercase = args["uppercase"]
let lowercase = args["lowercase"]
let count = args["count"]

// Read input file
let content = os.readFile(inputFile)

// Process content
if (uppercase) {
    content = content.toUpperCase()
}
if (lowercase) {
    content = content.toLowerCase()
}

// Write output
if (outputFile) {
    os.writeFile(outputFile, content)
    print("Processed content written to", outputFile)
}

// Count statistics
if (count) {
    let lines = content.split("\n").length
    let words = content.split(/\s+/).length
    let chars = content.length
    
    print("Lines:", lines)
    print("Words:", words)
    print("Characters:", chars)
}
```