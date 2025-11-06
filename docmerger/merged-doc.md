# Merged Documentation

## RELEASE_PROCESS.md

```markdown
# VintLang Release Process

This document explains how to create and publish releases for VintLang using GoReleaser and GitHub Actions.

## Overview

VintLang uses [GoReleaser](https://goreleaser.com) to automate the build and release process. When a new version tag is pushed to the repository, GitHub Actions automatically:

1. Builds binaries for multiple platforms (Linux, macOS, Windows)
2. Creates archives (tar.gz for Linux/macOS, zip for Windows)
3. Generates Linux packages (deb, rpm, apk)
4. Calculates checksums
5. Creates a GitHub release with all artifacts
6. Generates release notes from the changelog

## Prerequisites

- Maintainer access to the repository
- Git configured with proper credentials
- Go 1.21 or later installed (for local testing)

## Release Process

### 1. Prepare for Release

Before creating a release, ensure:

- All changes are merged to the `main` branch
- All tests are passing
- Documentation is up to date
- CHANGELOG.md is updated with the new version

### 2. Create a Release (Using the Release Script)

The easiest way to create a release is using the provided release script:

```bash
./scripts/release.sh v0.3.0
```

This script will:

- Validate the version format
- Check you're on the main branch
- Check for uncommitted changes
- Create and push the version tag
- Trigger the GitHub Actions workflow

#### Version Format

Versions must follow the format: `vX.Y.Z` (e.g., `v0.3.0`, `v1.0.0`)

Optional suffixes are supported: `v0.3.0-beta.1`, `v0.3.0-rc.1`

### 3. Manual Release Process

If you prefer to create a release manually:

```bash
# Ensure you're on main branch
git checkout main

# Pull latest changes
git pull origin main

# Create and push the tag
git tag -a v0.3.0 -m "Release v0.3.0"
git push origin main
git push origin v0.3.0
```

### 4. Monitor the Release

After pushing the tag:

1. Go to the [Actions tab](https://github.com/vintlang/vintlang/actions)
2. Find the "Release" workflow for your tag
3. Monitor the build progress
4. Once complete, verify the [release page](https://github.com/vintlang/vintlang/releases)

## Testing Releases Locally

Before creating an official release, you can test the build process locally using GoReleaser:

### Using the Test Script

```bash
./scripts/test-goreleaser.sh
```

This script will:

- Install GoReleaser if not already installed
- Build binaries for all platforms in snapshot mode
- Create archives and packages
- Generate checksums

All artifacts will be in the `dist/` directory.

### Manual Testing

```bash
# Install goreleaser (if not installed)
go install github.com/goreleaser/goreleaser@latest

# Ensure it's in your PATH
export PATH="$HOME/go/bin:$PATH"

# Test the configuration
goreleaser check

# Build for a single platform (fast)
goreleaser build --snapshot --clean --single-target

# Build for all platforms (slower)
goreleaser release --snapshot --clean --skip=publish
```

## Configuration

### GoReleaser Configuration

The GoReleaser configuration is in `.goreleaser.yml`. Key settings:

- **Builds**: Configured for Linux, macOS, and Windows (amd64, arm64, 386)
- **Archives**: tar.gz for Unix-like systems, zip for Windows
- **Packages**: Generates deb, rpm, and apk packages for Linux
- **Checksums**: SHA256 checksums for all artifacts

### GitHub Actions Workflow

The release workflow is in `.github/workflows/build.yml`. It:

- Triggers on tags matching `v*`
- Runs tests before building
- Uses GoReleaser to build and publish
- Requires `GITHUB_TOKEN` (automatically provided)

## Troubleshooting

### Build Fails

If the build fails:

1. Check the Actions logs for detailed error messages
2. Test locally with `./scripts/test-goreleaser.sh`
3. Ensure `.goreleaser.yml` is valid with `goreleaser check`

### Missing Artifacts

If some artifacts are missing from the release:

1. Check the GoReleaser configuration for the platform
2. Verify the build succeeded for that platform in the Actions logs
3. Test locally with `goreleaser release --snapshot --clean --skip=publish`

### Wrong Version Number

If the version number is incorrect:

1. Check that the tag name is correct (should start with `v`)
2. Verify the ldflags in `.goreleaser.yml` are set correctly
3. The version is injected at build time from the Git tag

## Platform Support

VintLang is built for the following platforms:

| OS      | Architectures        | Formats          |
|---------|---------------------|------------------|
| Linux   | amd64, arm64, 386   | tar.gz, deb, rpm, apk |
| macOS   | amd64, arm64        | tar.gz           |
| Windows | amd64, 386          | zip              |

Note: Windows arm64 and macOS 386 are excluded due to lack of support.

## Release Artifacts

Each release includes:

1. **Binaries**: Pre-built executables for each platform
2. **Archives**: Compressed archives containing the binary and documentation
3. **Linux Packages**: Native packages for Debian/Ubuntu (deb), Red Hat/Fedora (rpm), and Alpine (apk)
4. **Checksums**: SHA256 checksums for verifying downloads
5. **Release Notes**: Auto-generated from commit messages

## Best Practices

1. **Always test locally first**: Run `./scripts/test-goreleaser.sh` before pushing a tag
2. **Follow semantic versioning**: Use major.minor.patch (e.g., v1.2.3)
3. **Update documentation**: Ensure README.md and CHANGELOG.md are current
4. **Test the binaries**: Download and test artifacts from the release page
5. **Announce the release**: Update the community about new releases

## Additional Resources

- [GoReleaser Documentation](https://goreleaser.com)
- [Semantic Versioning](https://semver.org)
- [VintLang Releases](https://github.com/vintlang/vintlang/releases)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)

```

## SHOWCASE_README.md

```markdown
# VintLang Showcase Applications

This directory contains comprehensive showcase applications demonstrating VintLang's real-world capabilities.

## 🎯 Comprehensive Feature Showcase

**File:** `comprehensive_showcase.vint`

A complete demonstration of VintLang's core features including:

- **Data Structures**: Variables, arrays, dictionaries
- **String Processing**: Manipulation, splitting, transformation
- **Functions**: Definition, calls, parameters
- **JSON Operations**: Encoding, decoding, data processing
- **File I/O**: Reading, writing, directory management
- **Time Operations**: Formatting, timestamps
- **UUID Generation**: Unique identifiers
- **Data Analysis**: Statistics, processing
- **Report Generation**: Multiple output formats

### Usage

```bash
vint comprehensive_showcase.vint
```

### Output

- Creates employee database with JSON processing
- Generates comprehensive reports
- Demonstrates data analysis capabilities
- Shows file management operations

## 🗂️ Personal Information Manager

**File:** `vintlang_showcase.vint`

A practical application showcasing:

- Contact management system
- JSON data persistence
- CSV export functionality
- Report generation
- Statistical analysis
- File operations

### Usage

```bash
vint vintlang_showcase.vint
```

### Features Demonstrated

- UUID generation for unique records
- Time formatting and timestamps
- JSON encoding/decoding
- String manipulation (splitting, processing)
- Data structures (arrays, dictionaries)
- File I/O operations
- Statistical calculations

## 🔧 Feature Testing Suite

**File:** `feature_test.vint`

Basic feature validation covering:

- Time functions
- UUID generation
- JSON operations
- File operations
- Arrays and loops
- String operations

### Usage

```bash
vint feature_test.vint
```

## 🌐 Web Data Fetcher

**File:** `web_fetcher.vint`

Network capabilities demonstration:

- HTTP GET requests
- JSON response processing
- Error handling
- Performance analysis
- Data storage

### Usage

```bash
vint web_fetcher.vint
```

**Note:** Network features require internet connectivity.

## 🧮 Mathematical Algorithms

**File:** `math_showcase.vint`

Computational capabilities including:

- Fibonacci sequence generation
- Prime number detection
- Factorial calculations
- Sorting algorithms
- Statistical analysis

### Usage

```bash
vint math_showcase.vint
```

## 📁 File Management System

**File:** `file_manager.vint`

Advanced file operations:

- Directory management
- File analysis
- Backup functionality
- Logging systems
- Configuration management

### Usage

```bash
vint file_manager.vint
```

## 📝 Task Management Applications

### Simple Task Manager

**File:** `simple_task_manager.vint`

Basic task management with:

- Task creation and completion
- JSON persistence
- Statistics tracking

### Advanced Task Manager

**File:** `showcase_task_manager.vint`

Comprehensive task management with:

- Categories and priorities
- Advanced filtering
- Export functionality
- Interactive menus

## 🛠️ Fixed Examples

The following examples have been fixed to work correctly with current VintLang syntax:

- `examples/json.vint` - Fixed variable declarations
- `examples/os.vint` - Fixed variable declarations  
- `examples/strings.vint` - Fixed variable declarations
- `examples/regex.vint` - Identified syntax issues (needs further work)

## 📚 Key Language Features Demonstrated

### Working Features ✅

- **Variables**: `let` declarations, type inference
- **Data Types**: strings, numbers, booleans, arrays, dictionaries
- **Control Flow**: if/else, loops, functions
- **Modules**: time, os, json, uuid, math
- **String Operations**: split, upper, lower, contains, reverse
- **File I/O**: read, write, directory operations
- **JSON**: encode, decode, manipulation
- **Time**: formatting, timestamps
- **UUID**: generation

### Known Issues ⚠️

- **Regex Module**: Parsing issues with function calls
- **Network Module**: May require internet connectivity
- **CSV Module**: Some syntax parsing issues

## 🚀 Running the Showcases

1. Ensure VintLang is installed and available as `vint`
2. Navigate to the VintLang directory
3. Run any showcase file:

   ```bash
   ./vint comprehensive_showcase.vint
   ```

## 📊 Performance Notes

- All showcases run efficiently on the current VintLang interpreter
- File operations create temporary files for demonstration
- JSON processing handles complex nested structures
- String operations support Unicode text
- Memory usage is optimized for typical business applications

## 🎯 Real-World Applications

These showcases prove VintLang's readiness for:

- **Business Applications**: Data processing, reporting
- **File Management**: Backup, organization, analysis
- **API Integration**: JSON processing, data transformation
- **Automation Scripts**: Task management, file operations
- **Educational Programming**: Clear syntax, comprehensive features
- **Rapid Prototyping**: Quick development cycles

VintLang demonstrates production-ready capabilities for modern software development!

```

## argparse.md

```markdown
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

```

## args.md

```markdown
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

```

## arrays.md

```markdown
# Arrays in vint

Arrays in vint are versatile data structures that can hold multiple items, including different types such as numbers, strings, booleans, functions, and null values. This page covers various aspects of arrays, including how to create, manipulate, and iterate over them using vint's built-in keywords and methods.

## Creating Arrays

To create an array, use square brackets [] and separate items with commas:

```s
list = [1, "second", true]
```

## Accessing and Modifying Array Elements

Arrays in vint are zero-indexed. To access an element, use the element's index in square brackets:

```s
num = [10, 20, 30]
n = num[1]  // n is 20
```

You can reassign an element in an array using its index:

```s
num[1] = 25
```

## Concatenating Arrays

To concatenate two or more arrays, use the + operator:

```s
a = [1, 2, 3]
b = [4, 5, 6]
c = a + b
// c is now [1, 2, 3, 4, 5, 6]
```

## Checking for Array Membership

Use the `in` keyword to check if an item exists in an array:

```s
num = [10, 20, 30]
print(20 in num)  // will print true
```

## Looping Over Arrays

You can use the for and in keywords to loop over array elements. To loop over just the values, use the following syntax:

```
num = [1, 2, 3, 4, 5]

for value in num {
    print(value)
}
```

To loop over both index and value pairs, use this syntax:

```s
man = ["Tach", "ekilie", "Tachera Sasi"]

for idx, n in man {
    print(idx, "-", n)
}
```

## Array Methods

Arrays in vint have several built-in methods:

### length()

length() returns the length of an array:

```s
a = [1, 2, 3]
urefu = a.length()
print(urefu)  // will print 3
```

### push()

push() adds one or more items to the end of an array:

```s
a = [1, 2, 3]
a.push("s", "g")
print(a)  // will print [1, 2, 3, "s", "g"]
```

### last()

last() returns the last item in an array, or null if the array is empty:

```s
a = [1, 2, 3]
last_el = a.last()
print(last_el)  // will print 3

b = []
last_el = b.last()
print(last_el)  // will print tupu
```

### pop()

pop() removes and returns the last item in the array. If the array is empty, it returns null:

```s
a = [1, 2, 3]
last = a.pop()
print(last)  // will print 3
print(a)     // will print [1, 2]
```

### shift()

shift() removes and returns the first item in the array. If the array is empty, it returns null:

```s
a = [1, 2, 3]
first = a.shift()
print(first)  // will print 1
print(a)      // will print [2, 3]
```

### unshift()

unshift() adds one or more items to the beginning of the array:

```s
a = [3, 4]
a.unshift(1, 2)
print(a)  // will print [1, 2, 3, 4]
```

### reverse()

reverse() reverses the array in place:

```s
a = [1, 2, 3]
a.reverse()
print(a)  // will print [3, 2, 1]
```

### sort()

sort() sorts the array in place. It works for arrays of integers, floats, or strings:

```s
a = [3, 1, 2]
a.sort()
print(a)  // will print [1, 2, 3]

b = ["banana", "apple", "cherry"]
b.sort()
print(b)  // will print ["apple", "banana", "cherry"]

c = [3.14, 1.41, 2.71]
c.sort()
print(c)  // will print [1.41, 2.71, 3.14]
```

### map()

map() goes through every element in the array and applies the passed function to each element. It returns a new array with the updated elements:

```s
a = [1, 2, 3]
b = a.map(func(x){ return x * 2 })
print(b) // [2, 4, 6]
```

### filter()

filter() will go through every single element of an array and checks if that element returns true or false when passed into a function. It will return a new array with elements that returned true:

```s
a = [1, 2, 3, 4]

b = a.filter(func(x){
    if (x % 2 == 0) 
        {return true}
    return false
    })

print(b) // [2, 4]
```

### slice()

slice() extracts a section of an array and returns a new array:

```s
a = [1, 2, 3, 4, 5]
sliced = a.slice(1, 3)
print(sliced)  // [2, 3]

// With just start index
sliced2 = a.slice(2)
print(sliced2)  // [3, 4, 5]
```

### concat()

concat() merges two or more arrays into a new array:

```s
a = [1, 2]
b = [3, 4]
c = [5, 6]
combined = a.concat(b, c)
print(combined)  // [1, 2, 3, 4, 5, 6]
```

### includes()

includes() checks if an array contains a specific element:

```s
numbers = [1, 2, 3, 4, 5]
print(numbers.includes(3))  // true
print(numbers.includes(10)) // false
```

### every()

every() tests whether all elements pass a test function:

```s
numbers = [2, 4, 6, 8]
allEven = numbers.every(func(x){ return x % 2 == 0 })
print(allEven)  // true
```

### some()

some() tests whether at least one element passes a test function:

```s
numbers = [1, 3, 5, 8]
hasEven = numbers.some(func(x){ return x % 2 == 0 })
print(hasEven)  // true
```

### reduce()

reduce() reduces the array to a single value using an accumulator function:

```s
numbers = [1, 2, 3, 4]
sum = numbers.reduce(func(acc, val){ return acc + val }, 0)
print(sum)  // 10

// Without initial value
product = numbers.reduce(func(acc, val){ return acc * val })
print(product)  // 24
```

### flatten()

flatten() flattens nested arrays into a single array:

```s
nested = [[1, 2], [3, 4], [5]]
flat = nested.flatten()
print(flat)  // [1, 2, 3, 4, 5]

// With depth limit
deep = [[[1, 2]], [3, 4]]
flatOne = deep.flatten(1)
print(flatOne)  // [[1, 2], 3, 4]
```

### unique()

unique() returns a new array with duplicate elements removed:

```s
numbers = [1, 2, 2, 3, 3, 4]
uniqueNumbers = numbers.unique()
print(uniqueNumbers)  // [1, 2, 3, 4]
```

### fill()

fill() fills all elements of an array with a static value:

```s
arr = [1, 2, 3, 4]
arr.fill(0)
print(arr)  // [0, 0, 0, 0]

// Fill with start and end positions
arr2 = [1, 2, 3, 4, 5]
arr2.fill(9, 1, 3)
print(arr2)  // [1, 9, 9, 4, 5]
```

### lastIndexOf()

lastIndexOf() returns the last index at which a given element can be found:

```s
numbers = [1, 2, 3, 2, 4]
lastIndex = numbers.lastIndexOf(2)
print(lastIndex)  // 3

// Element not found
notFound = numbers.lastIndexOf(10)
print(notFound)  // -1
```

## Mathematical Array Methods

Arrays in vint include several mathematical methods for numeric data analysis:

### sum()

sum() calculates the sum of all numeric elements:

```s
numbers = [1, 2, 3, 4, 5]
total = numbers.sum()
print(total)  // 15

floats = [1.5, 2.5, 3.5]
floatSum = floats.sum()
print(floatSum)  // 7.5
```

### average() / mean()

average() calculates the arithmetic mean of all numeric elements:

```s
numbers = [2, 4, 6, 8]
avg = numbers.average()
print(avg)  // 5

// mean() is an alias for average()
mean = numbers.mean()
print(mean)  // 5
```

### min() / max()

min() and max() find the minimum and maximum values:

```s
numbers = [5, 1, 9, 3, 7]
minimum = numbers.min()
print(minimum)  // 1

maximum = numbers.max()
print(maximum)  // 9
```

### median()

median() calculates the median (middle value) of numeric elements:

```s
oddNumbers = [1, 3, 5, 7, 9]
medianOdd = oddNumbers.median()
print(medianOdd)  // 5

evenNumbers = [2, 4, 6, 8]
medianEven = evenNumbers.median()
print(medianEven)  // 5 (average of 4 and 6)
```

### mode()

mode() returns the most frequently occurring value(s):

```s
numbers = [1, 2, 2, 3, 2, 4]
mostFrequent = numbers.mode()
print(mostFrequent)  // [2]
```

### variance()

variance() calculates the population variance:

```s
data = [2, 4, 4, 4, 5, 5, 7, 9]
var = data.variance()
print(var)  // 4
```

### standardDeviation()

standardDeviation() calculates the population standard deviation:

```s
data = [2, 4, 4, 4, 5, 5, 7, 9]
stdDev = data.standardDeviation()
print(stdDev)  // 2
```

### product()

product() calculates the product of all numeric elements:

```s
numbers = [2, 3, 4]
prod = numbers.product()
print(prod)  // 24
```

## Enhanced Sorting Methods

In addition to the basic sort() method, vint provides enhanced sorting capabilities:

### sortAsc()

sortAsc() is an alias for sort() that explicitly sorts in ascending order:

```s
numbers = [3, 1, 4, 1, 5]
numbers.sortAsc()
print(numbers)  // [1, 1, 3, 4, 5]
```

### sortDesc()

sortDesc() sorts the array in descending order:

```s
numbers = [3, 1, 4, 1, 5]
numbers.sortDesc()
print(numbers)  // [5, 4, 3, 1, 1]
```

### sortBy()

sortBy() sorts the array using a custom comparison function:

```s
// Sort by absolute value
numbers = [-3, 1, -4, 2]
numbers.sortBy(func(x){ return abs(x) })
print(numbers)  // [1, 2, -3, -4]

// Sort strings by length
words = ["hello", "hi", "world", "a"]
words.sortBy(func(x){ return len(x) })
print(words)  // ["a", "hi", "hello", "world"]
```

```

## async.md

```markdown
# VintLang Async Operations Guide

VintLang now supports native async operations and a simple concurrency model inspired by Go's goroutines and channels.

## Async Functions

Create async functions using the `async func` syntax. These functions return promises that can be awaited:

```javascript
// Define an async function
let fetchData = async func(url) {
    // Simulate async work
    return "Data from " + url
}

// Call the function (returns a promise)
let promise = fetchData("https://api.example.com")

// Wait for the result
let result = await promise
print("Result:", result)
```

## Async/Await

Use `await` to wait for promise resolution:

```javascript
let processData = async func(data) {
    return "Processed: " + data
}

let result = await processData("my data")
print(result)  // Output: Processed: my data
```

## Concurrent Execution with `go`

Use the `go` keyword to execute code concurrently:

```javascript
// Execute concurrently
go print("This runs in a goroutine")
go print("This also runs concurrently")

print("This runs in the main thread")
```

## Channels

Channels provide communication between concurrent operations.

### Creating Channels

```javascript
// Unbuffered channel
let ch = chan

// Buffered channel with size 5
let bufferedCh = chan(5)
```

### Channel Operations

```javascript
// Send to channel
send(ch, "Hello")

// Receive from channel
let message = receive(ch)

// Close channel
close(ch)
```

### Producer-Consumer Pattern

```javascript
let dataChan = chan(3)

// Producer goroutine
go func() {
    send(dataChan, "Item 1")
    send(dataChan, "Item 2")
    send(dataChan, "Item 3")
    close(dataChan)
}()

// Consumer
let item1 = receive(dataChan)
let item2 = receive(dataChan)
let item3 = receive(dataChan)

print("Received:", item1, item2, item3)
```

## Complex Example: Async with Channels

Combine async functions with channels for powerful patterns:

```javascript
let processInBackground = async func(input) {
    let resultChan = chan
    
    // Process in background
    go func() {
        let processed = "Processed: " + input
        send(resultChan, processed)
    }()
    
    // Wait for result
    let result = receive(resultChan)
    return result
}

let promise = processInBackground("data")
let result = await promise
print("Final result:", result)
```

## Error Handling

Async functions that encounter errors will reject their promises:

```javascript
let riskyFunction = async func() {
    // If an error occurs, the promise will be rejected
    return "Success!"
}

let result = await riskyFunction()
print("Result:", result)
```

## Multiple Concurrent Operations

Execute multiple async operations concurrently:

```javascript
let task1 = async func() { return "Task 1 complete" }
let task2 = async func() { return "Task 2 complete" }
let task3 = async func() { return "Task 3 complete" }

// Start all tasks
let p1 = task1()
let p2 = task2()
let p3 = task3()

// Wait for all to complete
let r1 = await p1
let r2 = await p2
let r3 = await p3

print("All tasks done:", r1, r2, r3)
```

## Best Practices

1. Use async functions for operations that might take time
2. Use channels for communication between goroutines
3. Always close channels when done sending
4. Use buffered channels to avoid blocking
5. Combine async/await with goroutines for powerful concurrent patterns

The async operations in VintLang provide a simple yet powerful way to handle concurrency and asynchronous operations in your programs.

```

## bool.md

```markdown
# Working with Booleans in vint

Boolean objects in vint are truthy, meaning that any value is true, except tupu and false. They are used to evaluate expressions that return true or false values.

## Evaluating Boolean Expressions

### Evaluating Simple Expressions

In vint, you can evaluate simple expressions that return a boolean value:

```s
print(1 > 2) // Output: `false`

print(1 + 3 < 10) // Output: `true`
```

### Evaluating Complex Expressions

In vint, you can use boolean operators to evaluate complex expressions:

```s
a = 5
b = 10
c = 15

result = (a < b) && (b < c)

if (result) {
    print("Both conditions are true")
} else {
    print("At least one condition is false")
}
// Output: "Both conditions are true"
```

Here, we create three variables a, b, and c. We then evaluate the expression (a < b) && (b < c). Since both conditions are true, the output will be "Both conditions are true".

## Boolean Operators

vint has several boolean operators that you can use to evaluate expressions:

### The && Operator

The && operator evaluates to true only if both operands are true. Here's an example:

```s
print(true && true) // Output: `true`

print(true && false) // Output: `false`
```

### The and() function

```s
print(and(true,true)) // Output: `true`

print(and(true,false)) // Output: `false`
```

### The || Operator

The || operator evaluates to true if at least one of the operands is true. Here's an example:

```s
print(true || false) // Output: `true`

print(false || false) // Output: `false`
```

### The or() Function

```s
print(or(true,false)) // Output: `true`

print(or(false,false)) // Output: `false`
```

### The ! Operator

The ! operator negates the value of the operand. Here's an example:

```s
print(!true) // Output: `false`

print(!false) // Output: `true`
```

### The not() function

```s
print(not(true)) // Output: `false`

print(not(false)) // Output: `true`
```

## Working with Boolean Values in Loops

In vint, you can use boolean expressions in loops to control their behavior. Here's an example:

```s
num = [1, 2, 3, 4, 5]

for v in num {
    if (v % 2 == 0) {
        print(v, "is even")
    } else {
        print(v, "is odd")
    }
}
// Output:
// 1 is odd
// 2 is even
// 3 is odd
// 4 is even
// 5 is odd
```

## Boolean Methods

Boolean values in vint come with several useful built-in methods for conversion and logical operations:

### to_string()

Converts the boolean value to a string representation:

```s
let flag = true
print(flag.to_string())     // "true"

let disabled = false
print(disabled.to_string()) // "false"
```

### to_int()

Converts the boolean value to an integer (1 for true, 0 for false):

```s
let enabled = true
print(enabled.to_int())     // 1

let disabled = false
print(disabled.to_int())    // 0
```

### negate()

Returns the logical negation of the boolean value:

```s
let flag = true
print(flag.negate())        // false

let condition = false
print(condition.negate())   // true
```

### toggle()

Returns the opposite boolean value (same as negate):

```s
let flag = true
print(flag.toggle())        // false

let condition = false
print(condition.toggle())   // true
```

### and()

Performs logical AND operation with another boolean:

```s
let a = true
let b = false
print(a.and(b))            // false
print(a.and(true))         // true
```

### or()

Performs logical OR operation with another boolean:

```s
let a = true
let b = false
print(a.or(b))             // true
print(b.or(false))         // false
```

### xor()

Performs logical XOR (exclusive OR) operation:

```s
let a = true
let b = false
print(a.xor(b))            // true
print(a.xor(true))         // false
```

### implies()

Performs logical implication (if A then B):

```s
let premise = true
let conclusion = false
print(premise.implies(conclusion))  // false
print(false.implies(false))         // true
```

### equivalent()

Checks if two boolean values are logically equivalent:

```s
let a = true
let b = true
print(a.equivalent(b))     // true
print(a.equivalent(false)) // false
```

### nor()

Performs logical NOR operation (NOT OR):

```s
let a = false
let b = false
print(a.nor(b))            // true
print(a.nor(true))         // false
```

### nand()

Performs logical NAND operation (NOT AND):

```s
let a = true
let b = true
print(a.nand(b))           // false
print(a.nand(false))       // true
```

## Practical Boolean Examples

Here are some practical examples using boolean methods:

```s
// Feature flags system
let features = {
    "dark_mode": true,
    "notifications": false,
    "beta_features": true
}

// Convert to configuration strings
for key, value in features {
    config_string = key + "=" + value.to_string()
    print(config_string)
}
// Output:
// dark_mode=true
// notifications=false
// beta_features=true

// Permission system using logical operations
let is_admin = true
let is_owner = false
let can_read = true

// Complex permission checks
let can_write = is_admin.or(is_owner)
let can_delete = is_admin.and(is_owner.negate())
let has_access = can_read.and(can_write.or(is_owner))

print("Can write:", can_write.to_string())     // true
print("Can delete:", can_delete.to_string())   // true
print("Has access:", has_access.to_string())   // true

// State machine logic
let door_open = false
let key_inserted = true
let button_pressed = true

// Door can be opened if key is inserted XOR button is pressed (but not both)
let can_open = key_inserted.xor(button_pressed).and(door_open.negate())
print("Can open door:", can_open.to_string())  // false

// Validation logic using implications
let form_valid = true
let submit_enabled = true

// If form is valid, then submit should be enabled
let validation_check = form_valid.implies(submit_enabled)
print("Validation passes:", validation_check.to_string())  // true
```

## Boolean Method Chaining

Boolean methods support method chaining for complex logical operations:

```s
// Complex boolean logic with chaining
let user_active = true
let subscription_valid = false
let trial_period = true

// Chain multiple operations
let has_access = user_active
    .and(subscription_valid.or(trial_period))
    .and(false.negate())

print("User has access:", has_access.to_string())  // true

// Truth table generation
conditions = [true, false]
for a in conditions {
    for b in conditions {
        print("A:", a.to_string(), "B:", b.to_string())
        print("  AND:", a.and(b).to_string())
        print("  OR:", a.or(b).to_string())
        print("  XOR:", a.xor(b).to_string())
        print("  NAND:", a.nand(b).to_string())
        print("---")
    }
}
```

```

## builtins.md

```markdown
# Built-in Functions in Vint

Vint has a number of built-in functions that are globally available to perform common tasks.

---

## I/O and System Functions

### `print(...)`
Prints messages to the standard output. It can take zero or more arguments, which will be printed with a space between them.
```js
print("Hello,", "world!") // Output: Hello, world!
print(1, 2, 3)         // Output: 1 2 3
```

### `println(...)`

Similar to `print`, but it adds a newline character at the end of the output.

### `input(prompt)`

Reads a line of input from the user from standard input. It can optionally take a string argument to use as a prompt.

```js
let name = input("Enter your name: ")
println("Hello,", name)
```

### `sleep(milliseconds)`

Pauses the program's execution for a specified duration in milliseconds.

```js
println("Waiting for 1 second...")
sleep(1000)
println("Done.")
```

### `exit(code)`

Terminates the program with a specified exit code. An exit code of `0` typically indicates success, while any other number indicates an error.

```js
if (some_error) {
    println("An error occurred!")
    exit(1)
}
```

---

## Type and General Information Functions

### `type(object)`

Returns a string representing the type of the given object.

```js
type(10)      // Output: "INTEGER"
type("hello") // Output: "STRING"
type([])      // Output: "ARRAY"
```

### `len(object)`

Returns the length of a string, array, or dictionary.

```js
len("hello")      // Output: 5
len([1, 2, 3])    // Output: 3
len({"a": 1})   // Output: 1
```

---

## Array Functions

### `append(array, element1, ...)`

Returns a *new* array with the given elements added to the end.

```js
let arr = [1, 2]
let new_arr = append(arr, 3, 4)
println(new_arr) // Output: [1, 2, 3, 4]
```

### `pop(array)`

Removes the last element from an array and returns that element. This function modifies the array in-place.

```js
let arr = [1, 2, 3]
let last = pop(arr)
println(last) // Output: 3
println(arr)  // Output: [1, 2]
```

---

## Dictionary Functions

### `keys(dictionary)`

Returns an array containing all the keys from a dictionary. The order is not guaranteed.

```js
let dict = {"name": "Alex", "age": 30}
println(keys(dict)) // Output: ["name", "age"] (or ["age", "name"])
```

### `values(dictionary)`

Returns an array containing all the values from a dictionary. The order corresponds to the order of the keys returned by `keys()`.

```js
let dict = {"name": "Alex", "age": 30}
println(values(dict)) // Output: ["Alex", 30] (or [30, "Alex"])
```

### `has_key(dictionary, key)`

Returns `true` if the dictionary contains the given key, and `false` otherwise. This is also available as a method on dictionary objects: `my_dict.has_key(key)`.

```js
let dict = {"a": 1}
println(has_key(dict, "a")) // Output: true
println(dict.has_key("b"))  // Output: false
```

---

## String and Character Functions

### `chr(integer)`

Returns a single-character string corresponding to the given integer ASCII code.

```js
println(chr(65)) // Output: "A"
```

### `ord(string)`

Returns the integer ASCII code of the first character of a given string.

```js
println(ord("A")) // Output: 65
```

---

## File Functions

### `open(filepath)`

Opens a file and returns a file object. This is typically used for reading file contents.

```js
let file = open("data.txt")
// You can then use methods on the file object
```

---

## Logical Functions

### `and(boolean1, boolean2)`

Performs a logical AND operation on two boolean values. Returns `true` only if both arguments are `true`.

```js
and(true, true)    // Output: true
and(true, false)   // Output: false
and(false, false)  // Output: false
```

### `or(boolean1, boolean2)`

Performs a logical OR operation on two boolean values. Returns `true` if at least one of the arguments is `true`.

```js
or(true, false)    // Output: true
or(false, false)   // Output: false
or(true, true)     // Output: true
```

### `not(boolean)`

Performs a logical NOT operation on a boolean value. Returns the opposite of the input.

```js
not(true)          // Output: false
not(false)         // Output: true
```

### `xor(boolean1, boolean2)`

Performs a logical XOR (exclusive OR) operation on two boolean values. Returns `true` when exactly one of the arguments is `true`.

```js
xor(true, false)   // Output: true
xor(false, true)   // Output: true
xor(true, true)    // Output: false
xor(false, false)  // Output: false
```

### `nand(boolean1, boolean2)`

Performs a logical NAND (NOT AND) operation on two boolean values. Returns `false` only when both arguments are `true`.

```js
nand(true, true)   // Output: false
nand(true, false)  // Output: true
nand(false, false) // Output: true
```

### `nor(boolean1, boolean2)`

Performs a logical NOR (NOT OR) operation on two boolean values. Returns `true` only when both arguments are `false`.

```js
nor(false, false)  // Output: true
nor(true, false)   // Output: false
nor(true, true)    // Output: false
```

---

## Additional Built-in Functions

### String Functions

#### `startsWith(string, prefix)`

Checks if a string starts with the specified prefix.

```js
startsWith("VintLang", "Vint")    // Output: true
startsWith("hello", "hi")         // Output: false
```

#### `endsWith(string, suffix)`

Checks if a string ends with the specified suffix.

```js
endsWith("VintLang", "Lang")      // Output: true
endsWith("hello", "world")        // Output: false
```

### Array Functions

#### `indexOf(array, element)`

Returns the index of the first occurrence of an element in an array, or -1 if not found.

```js
let arr = [1, 2, 3, 2, 4]
indexOf(arr, 2)    // Output: 1
indexOf(arr, 5)    // Output: -1
```

#### `unique(array)`

Returns a new array containing only the unique elements from the input array, removing duplicates.

```js
let arr = [1, 2, 2, 3, 1, 4]
unique(arr)        // Output: [1, 2, 3, 4]
unique([])         // Output: []
```

### Type Checking Functions

#### `isInt(value)`

Returns true if the value is an integer.

```js
isInt(42)          // Output: true
isInt(3.14)        // Output: false
isInt("hello")     // Output: false
```

#### `isFloat(value)`

Returns true if the value is a float.

```js
isFloat(3.14)      // Output: true
isFloat(42)        // Output: false
isFloat("hello")   // Output: false
```

#### `isString(value)`

Returns true if the value is a string.

```js
isString("hello")  // Output: true
isString(42)       // Output: false
isString(3.14)     // Output: false
```

#### `isBool(value)`

Returns true if the value is a boolean.

```js
isBool(true)       // Output: true
isBool(false)      // Output: true
isBool(42)         // Output: false
```

#### `isArray(value)`

Returns true if the value is an array.

```js
isArray([1, 2, 3]) // Output: true
isArray("hello")   // Output: false
isArray(42)        // Output: false
```

#### `isDict(value)`

Returns true if the value is a dictionary.

```js
isDict({"key": "value"})  // Output: true
isDict([1, 2, 3])         // Output: false
isDict("hello")           // Output: false
```

#### `isNull(value)`

Returns true if the value is null.

```js
isNull(null)       // Output: true
isNull(42)         // Output: false
isNull("")         // Output: false
```

### Parsing Functions

#### `parseInt(string)`

Parses a string and returns an integer.

```js
parseInt("42")     // Output: 42
parseInt("-10")    // Output: -10
parseInt("abc")    // Error: cannot parse 'abc' as integer
```

#### `parseFloat(string)`

Parses a string and returns a float.

```js
parseFloat("3.14")    // Output: 3.14
parseFloat("-2.5")    // Output: -2.5
parseFloat("hello")   // Error: cannot parse 'hello' as float
```

### Utility Functions

#### `debounce(delay, function)`

Creates a debounced version of a function that delays its execution until after `delay` milliseconds have elapsed since the last time the debounced function was invoked. This is useful for rate-limiting function calls, especially in response to user input or events.

The `delay` parameter can be:

- An integer representing milliseconds
- A Duration object

The `function` parameter can be:

- A user-defined function
- A builtin function

```js
// Create a debounced version of print with 500ms delay
let debouncedPrint = debounce(500, print)

// These rapid calls will be debounced - only the last one executes
debouncedPrint("First call")
debouncedPrint("Second call")  
debouncedPrint("Third call")   // Only this prints after 500ms

// Example with user-defined function
let logMessage = func(msg) {
    println("LOG:", msg)
}

let debouncedLog = debounce(1000, logMessage)
debouncedLog("This will be logged after 1 second of inactivity")
```

---

## Note on Existing Modules

VintLang also provides specialized modules for advanced functionality:

- **Math functions** like `abs`, `min`, `max`, `sqrt`, etc. are available in the `math` module
- **String functions** like `toUpper`, `toLower`, `trim`, `contains`, etc. are available in the `string` module  
- **Random functions** like `random.int()` and `random.float()` are available in the `random` module
- **KV functions** like `set`, `get`, `delete`, `increment`, etc. are available in the `kv` module for in-memory key-value storage
- **Array methods** like `reverse()` and `sort()` are available as methods on array objects

Use these modules for more advanced functionality:

```js
import math
import string
import random
import kv

let result = math.abs(-5)        // 5
let upper = string.toUpper("hi") // "HI"
let num = random.int(1, 10)      // Random number 1-10
let arr = [3, 1, 4].sort()       // [1, 3, 4]

// In-memory key-value storage
kv.set("user:123", {"name": "Alice"})
let user = kv.get("user:123")    // {name: Alice}
kv.increment("page_views")       // 1
```

```

## bundler.md

```markdown
# VintLang Bundler

## Overview

The **VintLang Bundler** compiles `.vint` source files into standalone Go binaries.

```sh
vint bundle yourfile.vint
```

This allows you to write code in VintLang, bundle it into an executable, and run it on any system without requiring the VintLang interpreter or Go to be installed on the target machine.

---

## Why Use the Bundler?

- Package and distribute VintLang scripts as self-contained executables
- End-users don’t need to install Go or VintLang
- Ideal for deploying scripts, shipping CLI tools, and automating workflows
- Internally powered by the `vintlang/repl` package for code execution

---

## Installation Requirements

You need the following tools installed on your system:

### 1. Go (version 1.18+)

Download and install from the official site:
[https://go.dev/dl/](https://go.dev/dl/)

Verify installation:

```sh
go version
```

### 2. Git and Go Modules

Ensure Go modules are enabled:

```sh
go env -w GO111MODULE=on
```

### 3. VintLang and the Bundler CLI

Install VintLang globally (includes the bundler):

```sh
go install github.com/vintlang/vintlang@latest
```

This makes the `vint` CLI available, including the `bundle` command.

---

## Multi-File Package Support (NEW!)

The VintLang Bundler now supports bundling multi-file packages with both imports and includes!

### Key Features

- ✅ **Automatic Dependency Discovery**: Finds all imported and included `.vint` files recursively
- ✅ **Package System Integration**: Handles `package` declarations and `import` statements
- ✅ **Include Statement Support**: Handles `include` statements for direct file embedding
- ✅ **Self-Contained Binaries**: No external `.vint` files needed at runtime
- ✅ **Compatible with Built-ins**: Works with all VintLang built-in modules

### Example Multi-File Project (Imports)

**main.vint**:

```js
import my_utils
import os

print("Starting application...")
let result = my_utils.process_data("hello")
print("Result:", result)
```

**my_utils.vint**:

```js
package my_utils {
    let process_data = func(input) {
        return "processed: " + input
    }
}
```

Bundle the entire project:

```sh
vint bundle main.vint
```

The bundler automatically discovers `my_utils.vint`, processes the package structure, and creates a single binary containing both files.

### Example Multi-File Project (Includes)

**main.vint**:

```js
include "config.vint"
include "helpers.vint"

print("Application:", appName)
print("Result:", processData("test"))
```

**config.vint**:

```js
let appName = "My VintLang App"
let version = "1.0.0"
```

**helpers.vint**:

```js
let processData = func(input) {
    return "processed: " + input
}
```

Bundle the entire project:

```sh
vint bundle main.vint
```

The bundler automatically discovers all included files and embeds their content directly into the bundled binary.

### Differences between Import and Include

- **Import statements** (`import module_name`) work with the package system and wrap content in packages
- **Include statements** (`include "file_path"`) directly embed file content without package wrapping
- Both are automatically discovered and bundled into self-contained binaries

---

## Usage

To bundle a `.vint` file into a binary:

```sh
vint bundle hello.vint
```

This creates a standalone executable named `hello` in the same directory.

To run the binary:

```sh
./hello
```

---

## Example

Given a simple `hello.vint` file:

```js
print("Hello, World!")
```

Run:

```sh
vint bundle hello.vint
```

This generates a binary `hello`. Execute it:

```sh
./hello
```

Expected output:

```
Hello, World!
```

---

## How It Works (Current Implementation)

The VintLang bundler has evolved into a sophisticated multi-stage pipeline that handles complex multi-file projects with automatic dependency resolution. Here's how it works:

### 🔍 Phase 1: Dependency Analysis

```
main.vint
    ↓ (parse AST)
    ├── import math_utils → finds math_utils.vint
    ├── include "config.vint" → finds config.vint  
    └── import os → skips (built-in module)
```

**What happens:**

- Parses main file's AST to find `import` and `include` statements
- Sets up search paths (main file directory, current directory, `./modules/`)
- Recursively discovers all dependency files
- Distinguishes between imports (modules) and includes (direct embedding)
- Skips built-in modules (like `os`, `http`, etc.)

### ⚙️ Phase 2: String Processing & Code Combination

```
Files discovered:
├── main.vint (import math_utils; include "config.vint"; ...)
├── math_utils.vint (package math_utils { ... })
└── config.vint (let appName = "App"; ...)

Processing:
├── math_utils.vint → wraps in package if needed
├── config.vint → embeds directly (no package wrapper)
└── main.vint → removes import/include statements for bundled files
```

**What happens:**

- **Import files**: Wrapped in package structure if not already packaged
- **Include files**: Content embedded directly, imports/includes removed  
- **Main file**: Import/include statements removed for bundled dependencies
- All code combined into single VintLang program

### 🏗️ Phase 3: Go Code Generation

```
Combined VintLang Code
    ↓ (escape for Go)
Template → main.go with embedded code
    ↓
package main
import "github.com/vintlang/vintlang/repl"
func main() {
    code := `<embedded VintLang code>`
    repl.Read(code)
}
```

**What happens:**

- Escapes VintLang code for safe embedding in Go string literals
- Generates Go main.go file using template
- Adds metadata (bundler version, build time)
- Creates go.mod file for dependencies

### 🔨 Phase 4: Binary Compilation

```
Temporary Directory
├── main.go (generated)
├── go.mod (generated)
    ↓ (go mod tidy && go build)
Binary Output (self-contained executable)
```

**What happens:**

- Creates temporary build directory
- Runs `go mod tidy` to resolve Go dependencies  
- Compiles with `go build -o binary_name`
- Moves final binary to output location
- Cleans up temporary files

### 🎯 Key Features of Current Implementation

1. **Automatic Dependency Discovery**: Recursively finds all `.vint` files through AST parsing
2. **Dual Processing Modes**:
   - `import module_name` → wraps content in packages
   - `include "file.vint"` → directly embeds content
3. **Smart Module Resolution**: Searches multiple paths, handles built-ins
4. **Self-Contained Output**: No external `.vint` files needed at runtime
5. **Cross-Compilation Support**: Uses GOOS/GOARCH environment variables

The resulting binary is completely portable and self-contained - no VintLang interpreter or external dependencies required!

### 📊 Visual Workflow Diagram

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   main.vint     │    │  math_utils.vint│    │   config.vint   │
│                 │    │                 │    │                 │
│ import math_utils│    │ package math_utils│   │ let appName =   │
│ include "config"│    │ {               │    │   "My App"      │
│ print(appName)  │    │   let add = ... │    │ let version =   │
│ ...             │    │ }               │    │   "1.0"         │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                        ┌────────▼─────────┐
                        │ Dependency       │
                        │ Analyzer         │
                        │ (AST parsing)    │
                        └────────┬─────────┘
                                 │
                        ┌────────▼─────────┐
                        │ String           │
                        │ Processor        │
                        │ (code combining) │
                        └────────┬─────────┘
                                 │
                    ┌────────────▼─────────────┐
                    │ Combined VintLang Code:  │
                    │                          │
                    │ package math_utils {     │
                    │   let add = ...          │
                    │ }                        │
                    │ let appName = "My App"   │
                    │ let version = "1.0"      │
                    │ print(appName)           │
                    │ ...                      │
                    └────────────┬─────────────┘
                                 │
                        ┌────────▼─────────┐
                        │ Bundled          │
                        │ Evaluator        │
                        │ (Go code gen)    │
                        └────────┬─────────┘
                                 │
                    ┌────────────▼─────────────┐
                    │ Generated main.go:       │
                    │                          │
                    │ package main             │
                    │ import "repl"            │
                    │ func main() {            │
                    │   code := `<embedded>`   │
                    │   repl.Read(code)        │
                    │ }                        │
                    └────────────┬─────────────┘
                                 │
                        ┌────────▼─────────┐
                        │ Go Compiler      │
                        │ (go build)       │
                        └────────┬─────────┘
                                 │
                        ┌────────▼─────────┐
                        │ Self-Contained   │
                        │ Binary           │
                        │ (portable exe)   │
                        └──────────────────┘
```

### 🏗️ Bundler Architecture Components

The bundler is built with several specialized components:

| Component | File | Purpose |
|-----------|------|---------|
| **Bundle Controller** | `bundler.go` | Main entry point, coordinates the entire bundling process |
| **Dependency Analyzer** | `dependencies.go` | Discovers and analyzes all imported/included files recursively |
| **String Processor** | `string_processor.go` | Combines files and handles import/include statement processing |
| **Bundled Evaluator** | `bundled_evaluator.go` | Generates the final Go code with embedded VintLang content |
| **Package Processor** | `package_processor.go` | Handles package structure and wrapping for imported modules |

**Flow**: Bundle Controller → Dependency Analyzer → String Processor → Bundled Evaluator → Go Compiler

### 🔄 Evolution from Original Design

The bundler has significantly evolved from the simple design originally described:

| Original Design | Current Implementation |
|----------------|----------------------|
| ✅ Single file bundling | ✅ Multi-file project support with dependency resolution |
| ✅ Simple string embedding | ✅ Advanced AST parsing and code processing |
| ❌ No import support | ✅ Full import/include statement handling |
| ❌ No package system | ✅ Package wrapping and module resolution |
| ❌ Manual dependency management | ✅ Automatic recursive dependency discovery |
| ✅ Basic Go template | ✅ Sophisticated string processing and escaping |

The current implementation handles complex multi-file projects automatically while maintaining the same simple command-line interface.

### 🔍 Step-by-Step Example Transformation

Let's see exactly what happens when bundling a multi-file project:

**Input Files:**

```js
// main.vint
import math_utils
include "config.vint"
print("App:", appName)
print("Result:", math_utils.add(5, 3))

// math_utils.vint  
package math_utils {
    let add = func(a, b) { return a + b }
}

// config.vint
let appName = "Calculator"
```

**After Dependency Analysis:**

```
Found 3 files:
├── main.vint (main file)
├── math_utils.vint (import dependency)
└── config.vint (include dependency)
```

**After String Processing:**

```js
// Combined VintLang code:
package math_utils {
    let add = func(a, b) { return a + b }
}

let appName = "Calculator"

print("App:", appName)
print("Result:", math_utils.add(5, 3))
```

**After Go Code Generation:**

```go
package main
import "github.com/vintlang/vintlang/repl"
func main() {
    code := `package math_utils {
    let add = func(a, b) { return a + b }
}

let appName = "Calculator"

print("App:", appName)
print("Result:", math_utils.add(5, 3))`
    repl.Read(code)
}
```

**Final Result:** Self-contained binary that outputs:

```
App: Calculator
Result: 8
```

---

## Output Structure

The generated Go code looks like this:

```go
package main

import (
 "github.com/vintlang/vintlang/repl"
)

func main() {
 code := ` + "`<your VintLang source code>`" + `
 repl.Read(code)
}
```

---

## Use Cases

- Distribute command-line tools built in VintLang
- Deploy scripts on systems where VintLang is not installed
- Share portable binaries for automation or education
- Build lightweight tools using VintLang and Go’s compiler

---

## Notes for Developers

- Temporary build directories are automatically created and cleaned
- Uses `text/template` for safe source code embedding
- The Go module created during bundling is isolated from your current project
- Spinner and CLI output are available for build feedback

---

## Important Details

- Go is required only during **build time**
- The resulting binary is portable and self-contained
- Cross-compilation is not supported out-of-the-box; build on the target OS/arch

---

## Conclusion

The **VintLang Bundler** lets you turn `.vint` files into standalone executables using a simple command:

```sh
vint bundle yourfile.vint
```

Build once. Run anywhere. No dependencies. No interpreter. Just execution.

```

## cli.md

```markdown
# CLI Module

The `cli` module provides a comprehensive set of tools for building command-line applications in VintLang. It allows you to parse arguments, handle flags, prompt for user input, execute external commands, and more.

## Functions

### `getArgs()`

Returns an array of all command-line arguments passed to the script.

**Usage:**

```js
import cli

let allArgs = cli.getArgs()
println("All arguments:", allArgs)
```

### `getFlags()`

Parses command-line arguments and returns a dictionary of flags (arguments starting with `--`). If a flag is followed by a value that doesn't start with `-`, it's treated as the flag's value. Otherwise, the flag's value is `true`.

**Usage:**

```js
import cli

// Command: vint my_script.vint --verbose --output "file.txt"
let flags = cli.getFlags()
println("Flags:", flags)
// Output: Flags: {"verbose": true, "output": "file.txt"}
```

### `getPositional()`

Returns an array of positional arguments (arguments that are not flags or their values).

**Usage:**

```js
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

```js
import cli

// Command: vint my_script.vint --output="report.txt"
let outputFile = cli.getArgValue("--output")
println("Output file:", outputFile) // "report.txt"
```

### `hasArg(flagName)`

Checks if a named argument (flag) is present in the command-line arguments.

- `flagName` (string): The name of the flag to check for (e.g., `"--verbose"`).

**Usage:**

```js
import cli

// Command: vint my_script.vint --verbose
if (cli.hasArg("--verbose")) {
    println("Verbose mode enabled.")
}
```

### `parse()`

A more advanced argument parser that returns a dictionary containing parsed flags, positional arguments, and helper methods (`has`, `get`, `positional`).

**Usage:**

```js
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

```js
import cli

let name = cli.prompt("Enter your name: ")
println("Hello, " + name)
```

### `confirm(message)`

Asks the user a yes/no question and returns `true` for "yes" and `false` for "no".

- `message` (string): The confirmation message to display.

**Usage:**

```js
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

```js
import cli

let files = cli.execCommand("ls -l")
println(files)
```

### `exit(statusCode)`

Terminates the script with a given status code.

- `statusCode` (integer): The exit status code (0 for success, non-zero for error).

**Usage:**

```js
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

```js
import cli

cli.help("My Awesome App", "This app does awesome things.")
```

### `version(appName, version)`

Prints version information for the CLI application.

- `appName` (string, optional): The name of the application.
- `version` (string, optional): The version number.

**Usage:**

```js
import cli

cli.version("My Awesome App", "1.0.0")
```

```

## clipboard.md

```markdown
# Using the Clipboard Module in Vint

The **Clipboard** module in Vint provides functionality to interact with the system clipboard, allowing you to read, write, and manage clipboard content programmatically.

## Available Functions

### `clipboard.write(text)`
Writes text content to the system clipboard.

**Parameters:**
- `text` - The text to write to clipboard (string, integer, float, or boolean)

**Returns:** 
- `true` if successful
- Error object if operation fails

### `clipboard.read()`
Reads the current text content from the system clipboard.

**Parameters:** None

**Returns:**
- String containing the clipboard content
- Error object if operation fails

### `clipboard.clear()`
Clears the clipboard by writing an empty string to it.

**Parameters:** None

**Returns:**
- `true` if successful
- Error object if operation fails

### `clipboard.hasContent()`
Checks if the clipboard contains any content.

**Parameters:** None

**Returns:**
- `true` if clipboard has content (non-empty)
- `false` if clipboard is empty
- Error object if operation fails

### `clipboard.all()`
Returns all clipboard data as an array. Since the system clipboard typically holds only one piece of text at a time, this returns an array containing the current clipboard content.

**Parameters:** None

**Returns:**
- Array containing clipboard content (empty array if clipboard is empty)
- Error object if operation fails

## Examples

### Basic Usage

```js
import clipboard

// Write text to clipboard
clipboard.write("Hello, World!")

// Read from clipboard
let content = clipboard.read()
print("Clipboard content:", content)

// Check if clipboard has content
if clipboard.hasContent() {
    print("Clipboard has content")
} else {
    print("Clipboard is empty")
}

// Get all clipboard data as array
let allData = clipboard.all()
print("All clipboard data:", allData)

// Clear clipboard
clipboard.clear()
print("Clipboard cleared")

// Check all data after clear
let emptyData = clipboard.all()
print("Clipboard after clear:", emptyData)
```

### Writing Different Data Types

```js
import clipboard

// Write string
clipboard.write("Hello World")

// Write number
clipboard.write(42)

// Write float
clipboard.write(3.14159)

// Write boolean
clipboard.write(true)
```

### Error Handling

```js
import clipboard

// The clipboard functions return error objects if operations fail
let result = clipboard.write("test")
if result.type == "ERROR" {
    print("Failed to write to clipboard:", result.message)
}

let content = clipboard.read()
if content.type == "ERROR" {
    print("Failed to read from clipboard:", content.message)
} else {
    print("Clipboard content:", content)
}
```

### Practical Example: Clipboard Manager

```js
import clipboard

// Save current clipboard content before modifying
let backup = clipboard.read()

// Process some text
let processed_text = "Processed: " + backup

// Write processed text back to clipboard
clipboard.write(processed_text)

print("Original:", backup)
print("Modified clipboard content:", clipboard.read())
```

## Platform Support

The clipboard module works across different platforms:

- **Windows**: Uses Windows Clipboard API
- **macOS**: Uses NSPasteboard
- **Linux**: Uses X11 clipboard (requires X11 environment)

## Notes

- The clipboard module requires appropriate system permissions to access the clipboard
- On some systems, clipboard access might require the application to be running in a graphical environment
- The module handles text content only; binary data is not supported
- Error handling is important as clipboard operations can fail due to system restrictions or lack of permissions

```

## comments.md

```markdown
# Comments in vint

In vint, you can write comments to provide explanations and documentation for your code. Comments are lines of text that are ignored by the vint interpreter, so they will not affect the behavior of your program. There are two types of comments in vint: single-line comments and multi-line comments.

## Single-Line Comments

Single-line comments are used to provide brief explanations or documentation for a single line of code. To write a single-line comment in vint, use two forward slashes (//) followed by your comment text. Here's an example:

```s
// This line will be ignored by the vint interpreter
```

In this example, the comment text "This line will be ignored by the vint interpreter" will be ignored by the interpreter, so it will not affect the behavior of the program.

## Multi-Line Comments

Multi-line comments are used to provide more detailed explanations or documentation for multiple lines of code. To write a multi-line comment in vint, use a forward slash followed by an asterisk ( /*) to start the comment, and an asterisk followed by a forward slash (*/ ) to end the comment. Here's an example:

```s
/*
These lines
Will 
be 
ignored
*/
```

In this example, all the lines between the /*and*/ symbols will be ignored by the vint interpreter, so they will not affect the behavior of the program.

By utilizing single-line and multi-line comments in vint, you can make your code more readable and easier to maintain for yourself and others who may need to work with your code in the future.

```

## const.md

```markdown
# Constants in Vint

Constants are used to declare variables with values that cannot be changed once assigned. This feature helps ensure immutability and prevents accidental reassignments, making your code more robust and predictable.

## Syntax Rules

The `const` keyword is used to declare a constant. It follows the same naming rules as `let`, but with one critical difference: its value is immutable.

- **Must be initialized at declaration.**
- **Cannot be reassigned.**

### Examples of Valid `const` Declarations:

```js
const PI = 3.14159
print(PI)  // Output: 3.14159

const GREETING = "Hello, Vint!"
print(GREETING)  // Output: "Hello, Vint!"
```

In the examples above, `PI` and `GREETING` are declared as constants and can be used throughout the program.

## Immutability

Once a constant is declared, its value cannot be changed. Attempting to reassign a `const` variable will result in an error.

### Example of an Invalid Reassignment

```js
const MAX_CONNECTIONS = 5
print(MAX_CONNECTIONS) // Output: 5

// This will cause an error
MAX_CONNECTIONS = 10 
// Error: Cannot assign to constant 'MAX_CONNECTIONS'
```

This immutability ensures that critical values in your program remain constant, preventing bugs and making your code easier to reason about.

## Best Practices

1. **Use for Unchanging Values:** Use `const` for values that should not change during the execution of your program, such as mathematical constants, configuration settings, or fixed values.

2. **Use Uppercase for Global Constants:** It's a common convention to use `UPPER_SNAKE_CASE` for global constants to make them easily distinguishable from regular variables.

   ```js
   const API_KEY = "your-secret-key"
   ```

3. **Prefer `const` Over `let`:** Whenever possible, prefer `const` over `let` to make your code safer and more predictable. Only use `let` when you know a variable's value needs to change.

## Constants in Packages

Constants work seamlessly with VintLang's package system and support the same access control features:

```js
package Config {
    // Public constants (accessible from outside)
    const VERSION = "2.1.0"
    const MAX_USERS = 1000
    
    // Private constants (internal use only)
    const _SECRET_KEY = "internal-key-abc123"
    const _DEBUG_MODE = true
    
    let getPublicConfig = func() {
        return {
            "version": VERSION,
            "max_users": MAX_USERS
            // Note: private constants are not exposed
        }
    }
}
```

### Package Constant Usage

```js
import "Config"

// Accessing public constants
print("App Version:", Config.VERSION)  // ✅ Works
print("Max Users:", Config.MAX_USERS)  // ✅ Works

// Attempting to access private constants
print(Config._SECRET_KEY)  // ❌ Error: cannot access private property
```

For more information about packages and access control, see the [Packages documentation](packages.md).

```

## crypto.md

```markdown
# Crypto Module

The `crypto` module provides a set of functions for common cryptographic operations, including hashing, symmetric encryption (AES), asymmetric encryption (RSA), and digital signatures.

## Hashing Functions

### `hashMD5(data)`

Computes the MD5 hash of a string.

- `data` (string): The input string to hash.

**Returns:** A string representing the 32-character hexadecimal MD5 hash.

**Usage:**

```js
import crypto

let hashed = crypto.hashMD5("hello world")
println(hashed) // "5eb63bbbe01eeed093cb22bb8f5acdc3"
```

### `hashSHA256(data)`

Computes the SHA-256 hash of a string.

- `data` (string): The input string to hash.

**Returns:** A string representing the 64-character hexadecimal SHA-256 hash.

**Usage:**

```js
import crypto

let hashed = crypto.hashSHA256("hello world")
println(hashed) // "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
```

## Symmetric Encryption Functions (AES)

### `encryptAES(data, key)`

Encrypts a string using AES (Advanced Encryption Standard).

- `data` (string): The plaintext string to encrypt.
- `key` (string): The encryption key. **Must be 16, 24, or 32 bytes long** for AES-128, AES-192, or AES-256 respectively.

**Returns:** A hexadecimal string representing the encrypted data.

**Usage:**

```js
import crypto

let secret = "this is a secret message"
let key = "a_16_byte_secret_key"

let encrypted = crypto.encryptAES(secret, key)
println("Encrypted:", encrypted)
```

### `decryptAES(encryptedData, key)`

Decrypts an AES-encrypted hexadecimal string.

- `encryptedData` (string): The hexadecimal string to decrypt.
- `key` (string): The decryption key. Must be the same key used for encryption.

**Returns:** The original plaintext string.

**Usage:**

```js
import crypto

let key = "a_16_byte_secret_key"
let encrypted = "..." // Result from encryptAES

let decrypted = crypto.decryptAES(encrypted, key)
println("Decrypted:", decrypted)
```

## Asymmetric Encryption Functions (RSA)

### `generateRSA(keySize)`

Generates an RSA key pair with the specified bit size.

- `keySize` (integer, optional): The key size in bits. Must be between 1024 and 4096. Defaults to 2048.

**Returns:** A dictionary containing `"private"` and `"public"` keys in PEM format.

**Usage:**

```js
import crypto

// Generate 2048-bit RSA key pair (default)
let keys = crypto.generateRSA()

// Generate 4096-bit RSA key pair
let strongKeys = crypto.generateRSA(4096)

println("Private key:", keys["private"])
println("Public key:", keys["public"])
```

### `encryptRSA(data, publicKey)`

Encrypts data using RSA public key encryption.

- `data` (string): The plaintext string to encrypt.
- `publicKey` (string): The RSA public key in PEM format.

**Returns:** A hexadecimal string representing the encrypted data.

**Usage:**

```js
import crypto

let keys = crypto.generateRSA(2048)
let message = "Secret message"

let encrypted = crypto.encryptRSA(message, keys["public"])
println("Encrypted:", encrypted)
```

### `decryptRSA(encryptedData, privateKey)`

Decrypts data using RSA private key decryption.

- `encryptedData` (string): The encrypted data in hexadecimal format.
- `privateKey` (string): The RSA private key in PEM format.

**Returns:** The original plaintext string.

**Usage:**

```js
import crypto

let keys = crypto.generateRSA(2048)
let encrypted = "..." // Result from encryptRSA

let decrypted = crypto.decryptRSA(encrypted, keys["private"])
println("Decrypted:", decrypted)
```

## Digital Signature Functions

### `signRSA(data, privateKey)`

Creates a digital signature using RSA private key and SHA-256 hashing.

- `data` (string): The data to sign.
- `privateKey` (string): The RSA private key in PEM format.

**Returns:** A hexadecimal string representing the digital signature.

**Usage:**

```js
import crypto

let keys = crypto.generateRSA(2048)
let document = "Important document content"

let signature = crypto.signRSA(document, keys["private"])
println("Signature:", signature)
```

### `verifyRSA(data, signature, publicKey)`

Verifies a digital signature using RSA public key and SHA-256 hashing.

- `data` (string): The original data that was signed.
- `signature` (string): The signature in hexadecimal format.
- `publicKey` (string): The RSA public key in PEM format.

**Returns:** A boolean value indicating whether the signature is valid.

**Usage:**

```js
import crypto

let keys = crypto.generateRSA(2048)
let document = "Important document content"
let signature = crypto.signRSA(document, keys["private"])

let isValid = crypto.verifyRSA(document, signature, keys["public"])
println("Signature valid:", isValid) // true

let isTampered = crypto.verifyRSA("Tampered content", signature, keys["public"])
println("Tampered signature valid:", isTampered) // false
```

## Complete Example

```js
import crypto

// Generate RSA key pair
let keys = crypto.generateRSA(2048)

// Test message
let message = "Hello, secure world!"

// Test encryption/decryption
let encrypted = crypto.encryptRSA(message, keys["public"])
let decrypted = crypto.decryptRSA(encrypted, keys["private"])
println("Encryption test:", message == decrypted)

// Test digital signatures
let signature = crypto.signRSA(message, keys["private"])
let isValid = crypto.verifyRSA(message, signature, keys["public"])
println("Signature test:", isValid)

// Test hashing
println("MD5:", crypto.hashMD5(message))
println("SHA256:", crypto.hashSHA256(message))
```

## Security Notes

- **RSA Key Size**: Use at least 2048-bit keys for security. 4096-bit keys provide stronger security but slower performance.
- **AES Keys**: Use strong, randomly generated keys. Store keys securely and never hardcode them in source code.
- **Digital Signatures**: Always verify signatures before trusting signed data.
- **MD5**: Consider deprecated for security-critical applications. Use SHA-256 or stronger hash functions instead.

```

## csv.md

```markdown
# CSV Module in VintLang

The `csv` module provides functions to read from and write to CSV (Comma-Separated Values) files.

## Functions

### `csv.read(filePath)`
Reads a CSV file and returns its contents as an array of arrays.

```js
// Assuming 'data.csv' contains:
// name,age
// alice,30
// bob,25

data = csv.read("data.csv")
print(data) 
// Outputs: [["name", "age"], ["alice", "30"], ["bob", "25"]]
```

### `csv.write(filePath, data)`

Writes a 2D array to a CSV file. The `data` argument must be an array of arrays, and all cell values must be strings.

```js
users = [
    ["name", "email"],
    ["John Doe", "john.doe@example.com"],
    ["Jane Smith", "jane.smith@example.com"]
]

csv.write("users.csv", users)
// This will create 'users.csv' with the provided data.
```

```

## datetime.md

```markdown
# DateTime Module in VintLang

The `datetime` module provides comprehensive date and time manipulation capabilities with timezone support, duration handling, and advanced datetime operations.

## Importing DateTime

```js
import datetime
```

## Basic Functions

### `datetime.now([timezone])`

Get the current date and time, optionally in a specific timezone.

```js
let current = datetime.now()
print(current)  // 10:15:32 26-09-2025

let ny_time = datetime.now("America/New_York") 
print(ny_time)  // 06:15:32 26-09-2025
```

### `datetime.utcNow()`

Get the current UTC time.

```js
let utc = datetime.utcNow()
print(utc)  // 10:15:32 26-09-2025
```

### `datetime.parse(datetime_string, [format], [timezone])`

Parse a datetime string into a Time object.

```js
let parsed = datetime.parse("2024-12-25 15:30:00", "2006-01-02 15:04:05")
print(parsed)  // 15:30:00 25-12-2024

let with_tz = datetime.parse("2024-01-01 00:00:00", "2006-01-02 15:04:05", "America/New_York")
```

### `datetime.fromTimestamp(timestamp, [timezone])`

Create a Time object from a Unix timestamp.

```js
let time_from_ts = datetime.fromTimestamp(1704063000)
print(time_from_ts)  // Unix timestamp converted to local time
```

## Duration Functions

### `datetime.duration(string | keyword_args)`

Create a Duration object from a string or keyword arguments.

```js
// From string
let dur1 = datetime.duration("2h30m15s")

// From keyword arguments
let dur2 = datetime.duration(hours=2, minutes=30, seconds=15, days=1, weeks=1)

// Supported units: nanoseconds, microseconds, milliseconds, seconds, minutes, hours, days, weeks
```

### `datetime.sleep(duration)`

Sleep for a specified duration.

```js
datetime.sleep(datetime.duration("2s"))  // Sleep for 2 seconds
datetime.sleep(5)  // Sleep for 5 seconds (integer)
datetime.sleep("1m30s")  // Sleep for 1 minute 30 seconds (string)
```

## Time Utility Functions

### `datetime.since(time)`

Get the duration since a specific time.

```js
let past_time = datetime.parse("2024-01-01 00:00:00", "2006-01-02 15:04:05")
let duration_since = datetime.since(past_time)
print(duration_since)  // Duration since Jan 1, 2024
```

### `datetime.until(time)`

Get the duration until a future time.

```js
let future_time = datetime.parse("2025-12-31 23:59:59", "2006-01-02 15:04:05")
let duration_until = datetime.until(future_time)
print(duration_until)  // Duration until Dec 31, 2025
```

### `datetime.isLeapYear(year)`

Check if a year is a leap year.

```js
print(datetime.isLeapYear(2024))  // true
print(datetime.isLeapYear(2023))  // false
```

### `datetime.daysInMonth(year, month)`

Get the number of days in a specific month.

```js
print(datetime.daysInMonth(2024, 2))  // 29 (February in leap year)
print(datetime.daysInMonth(2023, 2))  // 28 (February in regular year)
```

## Period Boundary Functions

### `datetime.startOfDay(time)`

Get the start of the day (00:00:00) for a given time.

```js
let current = datetime.now()
let start = datetime.startOfDay(current)
print(start)  // 00:00:00 26-09-2025
```

### `datetime.endOfDay(time)`

Get the end of the day (23:59:59) for a given time.

```js
let current = datetime.now()
let end = datetime.endOfDay(current)
print(end)  // 23:59:59 26-09-2025
```

### `datetime.startOfWeek(time)`

Get the start of the week (Sunday 00:00:00) for a given time.

```js
let start_week = datetime.startOfWeek(datetime.now())
```

### `datetime.endOfWeek(time)`

Get the end of the week (Saturday 23:59:59) for a given time.

```js
let end_week = datetime.endOfWeek(datetime.now())
```

### `datetime.startOfMonth(time)`

Get the start of the month for a given time.

```js
let start_month = datetime.startOfMonth(datetime.now())
```

### `datetime.endOfMonth(time)`

Get the end of the month for a given time.

```js
let end_month = datetime.endOfMonth(datetime.now())
```

### `datetime.startOfYear(time)`

Get the start of the year for a given time.

```js
let start_year = datetime.startOfYear(datetime.now())
```

### `datetime.endOfYear(time)`

Get the end of the year for a given time.

```js
let end_year = datetime.endOfYear(datetime.now())
```

## Time Object Methods

Time objects returned by datetime functions have many useful methods:

### Basic Properties

```js
let time = datetime.now()
print(time.year())      // 2025
print(time.month())     // 9
print(time.day())       // 26
print(time.hour())      // 10
print(time.minute())    // 15
print(time.second())    // 32
print(time.nanosecond()) // Nanosecond component
print(time.weekday())   // "Friday"
print(time.yearDay())   // Day of year (1-366)
```

### ISO Week

```js
let iso = time.isoWeek()
print(iso["year"])  // ISO week year
print(iso["week"])  // ISO week number
```

### Time Arithmetic

```js
let time = datetime.now()
let duration = datetime.duration(hours=2, minutes=30)

// Add/subtract durations
let future = time.add(duration)
let past = time.subtract(duration)

// Add/subtract specific units
let tomorrow = time.add(days=1)
let last_week = time.subtract(weeks=1)
```

### Time Comparisons

```js
let time1 = datetime.now()
let time2 = datetime.parse("2025-01-01 00:00:00", "2006-01-02 15:04:05")

print(time1.before(time2))  // true/false
print(time1.after(time2))   // true/false
print(time1.equal(time2))   // true/false
print(time1.compare(time2)) // -1, 0, or 1
```

### Timezone Operations

```js
let time = datetime.now()

// Get current timezone
print(time.timezone())  // "UTC" or local timezone name

// Convert to specific timezone
let ny_time = time.timezone("America/New_York")
let utc_time = time.utc()
let local_time = time.local()
```

### Other Methods

```js
let time = datetime.now()

// Get Unix timestamp
print(time.timestamp())  // Unix timestamp as integer

// Format the time
print(time.format("2006-01-02 15:04:05"))  // Custom formatting

// Truncate/round to duration
let truncated = time.truncate("1h")  // Truncate to hour boundary
let rounded = time.round("15m")      // Round to nearest 15 minutes
```

## Duration Object Methods

Duration objects have methods for accessing different time units:

```js
let duration = datetime.duration(hours=2, minutes=30, seconds=15)

print(duration.hours())        // 2.5041666666666664
print(duration.minutes())      // 150.25
print(duration.seconds())      // 9015
print(duration.milliseconds()) // 9015000
print(duration.nanoseconds())  // 9015000000000
print(duration.string())       // "2h30m15s"
```

### Duration Arithmetic

```js
let dur1 = datetime.duration("1h")
let dur2 = datetime.duration("30m")

let sum = dur1.add(dur2)          // 1h30m
let diff = dur1.subtract(dur2)    // 30m
let product = dur1.multiply(2)    // 2h
let quotient = dur1.divide(2)     // 30m
let ratio = dur1.divide(dur2)     // 2.0 (ratio as float)
```

## Timezone Support

The datetime module supports timezone-aware operations:

### Available Timezones

Common timezone identifiers include:

- `UTC`
- `America/New_York`
- `America/Los_Angeles`
- `Europe/London`
- `Europe/Paris`
- `Asia/Tokyo`
- `Asia/Shanghai`
- And many more standard IANA timezone identifiers

### Examples

```js
// Current time in different timezones
let utc = datetime.now("UTC")
let ny = datetime.now("America/New_York")
let tokyo = datetime.now("Asia/Tokyo")

// Convert between timezones
let local_time = datetime.now()
let ny_time = local_time.timezone("America/New_York")
```

## Practical Examples

### Age Calculator

```js
let calculate_age = func(birth_date_str) {
    let birth = datetime.parse(birth_date_str, "2006-01-02")
    let current = datetime.now()
    let age_duration = current.subtract(birth)
    let age_years = age_duration.hours() / (24 * 365.25)
    return age_years.floor()
}

let age = calculate_age("1990-05-15")
print("Age:", age, "years")
```

### Meeting Scheduler

```js
let schedule_meeting = func(date_str, duration_str, timezone) {
    let start_time = datetime.parse(date_str, "2006-01-02 15:04:05", timezone)
    let duration = datetime.duration(duration_str)
    let end_time = start_time.add(duration)
    
    print("Meeting scheduled:")
    print("Start:", start_time.format("2006-01-02 15:04:05 MST"))
    print("End:", end_time.format("2006-01-02 15:04:05 MST"))
    print("Duration:", duration)
}

schedule_meeting("2024-12-25 14:00:00", "1h30m", "America/New_York")
```

### Time Until Event

```js
let time_until_event = func(event_date_str) {
    let event = datetime.parse(event_date_str, "2006-01-02 15:04:05")
    let now = datetime.now()
    
    if (now.after(event)) {
        print("Event has already passed!")
        return
    }
    
    let duration = datetime.until(event)
    let days = duration.hours() / 24
    print("Time until event:", days.floor(), "days")
}

time_until_event("2024-12-31 23:59:59")
```

## Integration with Time Module

The datetime module works alongside the existing time module. You can use both:

```js
import time
import datetime

// Traditional time module
let time_now = time.now()
print("Time module:", time_now)

// Enhanced datetime module  
let datetime_now = datetime.now()
print("DateTime module:", datetime_now)

// They can work together
let formatted = time.format(time_now, "2006-01-02 15:04:05")
let parsed = datetime.parse(formatted, "2006-01-02 15:04:05")
```

The datetime module provides all the functionality of the time module and much more, making it the recommended choice for complex date and time operations.

```

## debug.md

```markdown
# Debug

The `debug` keyword allows you to print debug messages at runtime for troubleshooting and development.

## Syntax

```js
debug "Your debug message here"
```

When the Vint interpreter encounters a `debug` statement, it prints a magenta-colored debug message to the console and continues execution. This is useful for inspecting variable values or program flow during development.

### Example

```js
let value = 42
debug "Current value is: " + value
println("Done.")
```

Running this script will output:

```
[DEBUG]: Current value is: 42
Done.
```

```

## declaratives.md

```markdown

```

## defer.md

```markdown
# Defer

The `defer` keyword provides a convenient way to schedule a function call to be executed just before the surrounding function returns. This is particularly useful for cleanup tasks, such as closing files or releasing resources, ensuring that they are always executed, regardless of how the function exits.

## Syntax

The `defer` keyword is followed by a function call:

```js
defer functionCall()
```

## Example

Here’s a simple example that demonstrates how `defer` works. The deferred `println` call is executed after the function body has completed but before the function returns.

```js
let my_function = func() {
    defer println("This will be printed last");
    println("This will be printed first");
};

my_function();
// Output:
// This will be printed first
// This will be printed last
```

## Multiple Defer Statements

If a function has multiple `defer` statements, they are pushed onto a stack. When the function returns, the deferred calls are executed in last-in, first-out (LIFO) order.

```js
let another_function = func() {
    defer println("deferred: 1");
    defer println("deferred: 2");
    println("function body");
};

another_function();
// Output:
// function body
// deferred: 2
// deferred: 1
```

This LIFO order is intuitive for managing resources. For example, if you acquire a resource and then lock it, you would want to unlock it first and then release it, which `defer` handles naturally.

```

## dictionaries.md

```markdown
Here’s a detailed explanation of dictionaries in Vint, without the Swahili terms:

### Dictionaries in Vint

In the Vint programming language, dictionaries are key-value data structures that allow you to store and manage data efficiently. These dictionaries can store any type of value (such as strings, integers, booleans, or even functions) and are incredibly useful for organizing and accessing data. 

### Creating Dictionaries

In Vint, dictionaries are created using curly braces `{}`. Each key is followed by a colon `:` and the corresponding value. Here's an example of a dictionary:

```js
dict = {"name": "John", "age": 30}
```

In this dictionary:

- `"name"` is the key, and `"John"` is the value.
- `"age"` is the key, and `30` is the value.

Keys can be of various data types like strings, integers, floats, or booleans, and values can be anything, including strings, integers, booleans, `null`, or even functions.

### Accessing Elements

You can access individual elements in a dictionary by using the key. For example:

```js
print(dict["name"]) // John
```

This will print `"John"`, the value associated with the key `"name"`.

### Updating Elements

To update the value of an existing key, simply assign a new value to the key:

```js
dict["age"] = 35
print(dict["age"]) // 35
```

This updates the `"age"` key to have the value `35`.

### Adding New Elements

To add a new key-value pair to a dictionary, assign a value to a new key:

```js
dict["city"] = "Dar es Salaam"
print(dict["city"]) // Dar es Salaam
```

This adds a new key `"city"` with the value `"Dar es Salaam"`.

### Concatenating Dictionaries

You can combine two dictionaries into one using the `+` operator:

```js
dict1 = {"a": "apple", "b": "banana"}
dict2 = {"c": "cherry", "d": "date"}
combined = dict1 + dict2
print(combined) // {"a": "apple", "b": "banana", "c": "cherry", "d": "date"}
```

In this case, `dict1` and `dict2` are merged into a new dictionary called `combined`.

### Checking If a Key Exists in a Dictionary

To check if a particular key exists in a dictionary, you can use the `in` keyword:

```js
"age" in dict // true
"salary" in dict // false
```

This checks whether the key `"age"` exists in the dictionary, which returns `true`, and checks whether the key `"salary"` exists, which returns `false`.

### Looping Over a Dictionary

You can loop over the keys and values of a dictionary using the `for` keyword:

```js
hobby = {"a": "reading", "b": "cycling", "c": "eating"}
for key, value in hobby {
    print(key, "=>", value)
}
```

This will output:

```
a => reading
b => cycling
c => eating
```

You can also loop over just the values without the keys:

```js
for value in hobby {
    print(value)
}
```

This will output:

```
reading
cycling
eating
```

## Dictionary Methods

Vint dictionaries come with several powerful built-in methods that make data manipulation easy and efficient:

### keys()

Get all keys from the dictionary as an array:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
keyList = contacts.keys()
print(keyList)  // ["Alice", "Bob"]
```

### values()

Get all values from the dictionary as an array:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
valueList = contacts.values()
print(valueList)  // ["alice@email.com", "bob@email.com"]
```

### size()

Get the number of key-value pairs in the dictionary:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
print(contacts.size())  // 2
```

### has()

Check if a key exists in the dictionary:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
print(contacts.has("Alice"))   // true
print(contacts.has("Charlie")) // false
```

### get()

Get a value by key with an optional default value:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
email = contacts.get("Alice", "unknown")        // "alice@email.com"
unknownEmail = contacts.get("Charlie", "unknown") // "unknown"
print(email)        // alice@email.com
print(unknownEmail) // unknown
```

### set()

Set a key-value pair in the dictionary:

```js
contacts = {"Alice": "alice@email.com"}
contacts.set("Bob", "bob@email.com")
print(contacts)  // {"Alice": "alice@email.com", "Bob": "bob@email.com"}

// Method chaining is supported
contacts.set("Charlie", "charlie@email.com").set("Dave", "dave@email.com")
```

### remove()

Remove a key-value pair from the dictionary:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
contacts.remove("Bob")
print(contacts)  // {"Alice": "alice@email.com"}
```

### clear()

Remove all key-value pairs from the dictionary:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
contacts.clear()
print(contacts)  // {}
```

### merge()

Merge another dictionary into this one:

```js
contacts = {"Alice": "alice@email.com"}
newContacts = {"Bob": "bob@email.com", "Charlie": "charlie@email.com"}
contacts.merge(newContacts)
print(contacts)  // {"Alice": "alice@email.com", "Bob": "bob@email.com", "Charlie": "charlie@email.com"}
```

### copy()

Create a shallow copy of the dictionary:

```js
original = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
backup = original.copy()
backup.set("Charlie", "charlie@email.com")
print(original)  // {"Alice": "alice@email.com", "Bob": "bob@email.com"}
print(backup)    // {"Alice": "alice@email.com", "Bob": "bob@email.com", "Charlie": "charlie@email.com"}
```

### filter()

Create a new dictionary with key-value pairs that pass a test function:

```js
scores = {"Alice": 85, "Bob": 92, "Charlie": 78, "Diana": 95}
highScores = scores.filter(func(key, value) { return value >= 90 })
print(highScores)  // {"Bob": 92, "Diana": 95}
```

### map()

Create a new dictionary with transformed values:

```js
prices = {"apple": 1.5, "banana": 0.8, "orange": 2.0}
discountedPrices = prices.map(func(key, value) { return value * 0.9 })
print(discountedPrices)  // {"apple": 1.35, "banana": 0.72, "orange": 1.8}
```

### forEach()

Execute a function for each key-value pair:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
contacts.forEach(func(key, value) { 
    print("Name:", key, "Email:", value) 
})
// Output:
// Name: Alice Email: alice@email.com
// Name: Bob Email: bob@email.com
```

### find()

Find the first key-value pair that satisfies a test function:

```js
users = {"user1": 25, "user2": 17, "user3": 32}
adult = users.find(func(key, value) { return value >= 18 })
print(adult)  // ["user1", 25] or null if not found
```

### some()

Test whether at least one key-value pair passes the test:

```js
scores = {"Alice": 85, "Bob": 72, "Charlie": 95}
hasHighScore = scores.some(func(key, value) { return value >= 90 })
print(hasHighScore)  // true
```

### every()

Test whether all key-value pairs pass the test:

```js
scores = {"Alice": 85, "Bob": 92, "Charlie": 95}
allPassed = scores.every(func(key, value) { return value >= 80 })
print(allPassed)  // true
```

### pick()

Create a new dictionary with only specified keys:

```js
user = {"name": "Alice", "age": 25, "email": "alice@email.com", "password": "secret"}
publicInfo = user.pick("name", "age", "email")
print(publicInfo)  // {"name": "Alice", "age": 25, "email": "alice@email.com"}
```

### omit()

Create a new dictionary excluding specified keys:

```js
user = {"name": "Alice", "age": 25, "email": "alice@email.com", "password": "secret"}
safeInfo = user.omit("password")
print(safeInfo)  // {"name": "Alice", "age": 25, "email": "alice@email.com"}
```

### isEmpty()

Check if the dictionary is empty:

```js
emptyDict = {}
filledDict = {"key": "value"}
print(emptyDict.isEmpty())   // true
print(filledDict.isEmpty())  // false
```

### equals()

Check if two dictionaries are equal:

```js
dict1 = {"name": "Alice", "age": 25}
dict2 = {"name": "Alice", "age": 25}
dict3 = {"name": "Bob", "age": 30}
print(dict1.equals(dict2))   // true
print(dict1.equals(dict3))   // false
```

### entries()

Get an array of [key, value] pairs:

```js
contacts = {"Alice": "alice@email.com", "Bob": "bob@email.com"}
entryList = contacts.entries()
print(entryList)  // [["Alice", "alice@email.com"], ["Bob", "bob@email.com"]]
```

### flatten()

Flatten nested dictionaries (one level deep):

```js
nested = {
    "user": {"name": "Alice", "age": 25},
    "status": "active"
}
flattened = nested.flatten()
print(flattened)  // {"user.name": "Alice", "user.age": 25, "status": "active"}
```

### deepMerge()

Recursively merge dictionaries:

```js
dict1 = {"user": {"name": "Alice"}, "status": "active"}
dict2 = {"user": {"age": 25}, "role": "admin"}
merged = dict1.deepMerge(dict2)
print(merged)  // {"user": {"name": "Alice", "age": 25}, "status": "active", "role": "admin"}
```

## Practical Examples

```

## Advanced Dictionary Usage

Here are some practical examples of using dictionaries with their methods:

```js
// Building a user database
users = {}
users.set("john", {"name": "John Doe", "age": 30, "city": "New York"})
users.set("jane", {"name": "Jane Smith", "age": 25, "city": "Los Angeles"})

// Check if user exists
if (users.has("john")) {
    user = users.get("john")
    print("User found:", user["name"])
}

// Get all usernames
usernames = users.keys()
print("All users:", usernames)

// Create settings with defaults
settings = {"theme": "dark", "notifications": true}
getTheme = settings.get("theme", "light")           // "dark"
getLanguage = settings.get("language", "english")   // "english" (default)

// Configuration management
config = {}
config.set("database", "localhost")
      .set("port", 5432)
      .set("timeout", 30)
print("Config:", config)
```

```

## docs.go

```markdown
package docs

import "embed"

//go:embed *
var Docs embed.FS
```

## dotenv.md

```markdown
# dotenv Module in Vint

The `dotenv` module in Vint is designed to load environment variables from a `.env` file into the application. This allows you to manage sensitive information such as API keys, database credentials, and other configuration values outside your codebase.

---

## Importing the dotenv Module

To use the `dotenv` module, import it as follows:

```js
import dotenv
```

---

## Functions and Examples

### 1. Loading Environment Variables with `load()`

The `load` function loads environment variables from a `.env` file into the application's environment. This function should be called at the start of your application to ensure all the necessary environment variables are available.

**Syntax**:

```js
load(filePath)
```

- `filePath`: The path to the `.env` file (relative or absolute).

**Example**:

```js
import dotenv

dotenv.load(".env")
```

This loads the environment variables from the `.env` file located in the current directory.

---

### 2. Accessing an Environment Variable with `get()`

After loading the environment variables, you can access specific variables using the `get` function. This function retrieves the value of a given environment variable by its name.

**Syntax**:

```js
get(variableName)
```

- `variableName`: The name of the environment variable to retrieve.

**Example**:

```js
import dotenv

dotenv.load(".env")
apiKey = dotenv.get("API_KEY")
print(apiKey)  // Expected output: The value of the "API_KEY" from the .env file
```

In this example, the value of the `API_KEY` environment variable is retrieved and printed.

---

## `.env` File Format

The `.env` file should contain key-value pairs of environment variables. Each line in the file represents a separate environment variable.

**Example `.env` file**:

```
API_KEY=your_api_key_here
DB_HOST=localhost
DB_USER=root
DB_PASS=password123
```

In this case, you would retrieve the value of `API_KEY` as shown in the previous example.

---

## Summary of Functions

| Function           | Description                                             | Example Output                             |
|--------------------|---------------------------------------------------------|--------------------------------------------|
| `load(filePath)`    | Loads environment variables from the specified `.env` file. | No direct output, but environment variables are loaded. |
| `get(variableName)` | Retrieves the value of a specified environment variable.  | The value of the variable (e.g., `"your_api_key_here"`) |

---

The `dotenv` module is an essential tool for securely managing configuration settings in Vint applications. By keeping sensitive data in a `.env` file, you avoid hardcoding secrets into your source code, thus improving security and maintainability.

```

## editor.md

```markdown
# VintLang Editor Module

The `editor` module provides a simple text editor component for terminal applications. It allows you to create and manipulate text editors within your VintLang applications, which is useful for note-taking apps, configuration editors, and other text-based tools.

## Features

- Create and manage text editors
- Set and get text content
- Line-based operations (insert, delete, update, get)
- Vim-like modal editing (normal, insert, command modes)
- Syntax highlighting (basic)
- Line numbering
- File loading and saving
- Customizable dimensions and appearance

## Usage

### Basic Example

```js
import editor
import term

// Create a new editor
let ed = editor.newEditor({
    width: 80,
    height: 20,
    showLineNum: true,
    syntax: "text"
})

// Set some initial text
editor.setText(ed, "Hello, world!\nThis is a simple text editor.\nTry it out!")

// Render the editor
let editorContent = editor.render(ed)
print(editorContent)

// Handle keypresses
while (true) {
    let key = term.getKey()
    
    // Handle special keys
    if (key == "q") {
        break // Exit the loop
    }
    
    // Pass the key to the editor
    editor.handleKeypress(ed, key)
    
    // Re-render the editor
    editorContent = editor.render(ed)
    term.clear()
    print(editorContent)
}
```

## API Reference

### newEditor(options)

Creates a new text editor.

**Parameters:**

- `options` (dict, optional): Options for the editor
  - `filename` (string): Initial file to load
  - `width` (integer): Width of the editor in characters (default: 80)
  - `height` (integer): Height of the editor in lines (default: 24)
  - `showLineNum` (boolean): Whether to show line numbers (default: true)
  - `syntax` (string): Syntax highlighting mode (default: "text")

**Returns:**

- An editor ID string that can be used with other functions

### setText(editorId, text)

Sets the entire text content of the editor.

**Parameters:**

- `editorId` (string): The editor ID
- `text` (string): The text content to set

**Returns:**

- `true` if the text was set successfully

### getText(editorId)

Gets the entire text content of the editor.

**Parameters:**

- `editorId` (string): The editor ID

**Returns:**

- A string containing the editor's text content

### insertLine(editorId, lineNumber, text)

Inserts a new line at the specified position.

**Parameters:**

- `editorId` (string): The editor ID
- `lineNumber` (integer, optional): The line number where to insert the text. If not provided, inserts at the current cursor position.
- `text` (string): The text to insert

**Returns:**

- `true` if the line was inserted successfully

### deleteLine(editorId, lineNumber)

Deletes a line at the specified position.

**Parameters:**

- `editorId` (string): The editor ID
- `lineNumber` (integer, optional): The line number to delete. If not provided, deletes the line at the current cursor position.

**Returns:**

- `true` if the line was deleted successfully

### updateLine(editorId, lineNumber, text)

Updates a line at the specified position.

**Parameters:**

- `editorId` (string): The editor ID
- `lineNumber` (integer, optional): The line number to update. If not provided, updates the line at the current cursor position.
- `text` (string): The new text for the line

**Returns:**

- `true` if the line was updated successfully

### getLine(editorId, lineNumber)

Gets a line at the specified position.

**Parameters:**

- `editorId` (string): The editor ID
- `lineNumber` (integer, optional): The line number to get. If not provided, gets the line at the current cursor position.

**Returns:**

- A string containing the line's text

### getLineCount(editorId)

Gets the number of lines in the editor.

**Parameters:**

- `editorId` (string): The editor ID

**Returns:**

- An integer representing the number of lines

### render(editorId)

Renders the editor content as a string.

**Parameters:**

- `editorId` (string): The editor ID

**Returns:**

- A string containing the rendered editor content

### handleKeypress(editorId, key)

Handles a keypress in the editor.

**Parameters:**

- `editorId` (string): The editor ID
- `key` (string): The key that was pressed

**Returns:**

- `true` if the keypress was handled successfully

### save(editorId, filename)

Saves the editor content to a file.

**Parameters:**

- `editorId` (string): The editor ID
- `filename` (string, optional): The filename to save to. If not provided, uses the editor's current filename.

**Returns:**

- `true` if the file was saved successfully

### load(editorId, filename)

Loads a file into the editor.

**Parameters:**

- `editorId` (string): The editor ID
- `filename` (string): The filename to load

**Returns:**

- `true` if the file was loaded successfully

## Editor Modes

The editor supports three modes of operation, similar to Vim:

1. **Normal Mode**: Default mode for navigation and commands
2. **Insert Mode**: For inserting and editing text
3. **Command Mode**: For executing commands like save and quit

### Normal Mode Keys

- `h`: Move cursor left
- `j`: Move cursor down
- `k`: Move cursor up
- `l`: Move cursor right
- `0`: Move to beginning of line
- `$`: Move to end of line
- `G`: Move to last line
- `gg`: Move to first line
- `x`: Delete character under cursor
- `dd`: Delete current line
- `i`: Enter insert mode
- `o`: Insert new line below and enter insert mode
- `O`: Insert new line above and enter insert mode
- `:`: Enter command mode

### Insert Mode Keys

- `Escape`: Exit insert mode
- `Backspace`: Delete character before cursor
- `Enter`: Split line at cursor
- `ArrowLeft`, `ArrowRight`, `ArrowUp`, `ArrowDown`: Move cursor
- Any other key: Insert character at cursor

### Command Mode Keys

- `Escape`: Exit command mode
- `Enter`: Execute command
- `Backspace`: Delete character before cursor
- Any other key: Add character to command

### Command Mode Commands

- `w`: Write (save) the file
- `q`: Quit (will warn if there are unsaved changes)
- `q!`: Force quit without saving
- `wq`: Write and quit
- `set number`: Show line numbers
- `set nonumber`: Hide line numbers

## Examples

### Simple Notepad Application

```js
import editor
import term
import argparse

// Parse command line arguments
let parser = argparse.newParser("notepad", "A simple notepad application")
argparse.addArgument(parser, "file", {
    description: "File to edit",
    required: false
})
let args = argparse.parse(parser)

// Create a new editor
let ed = editor.newEditor({
    filename: args["file"],
    width: term.getSize().width,
    height: term.getSize().height - 2
})

// Main loop
term.clear()
while (true) {
    // Render the editor
    let content = editor.render(ed)
    term.clear()
    print(content)
    
    // Get keypress
    let key = term.getKey()
    
    // Check for exit key (Ctrl+Q)
    if (key == "Ctrl+Q") {
        // Check if there are unsaved changes
        if (editor.isModified(ed)) {
            term.println("There are unsaved changes. Press Y to quit anyway, or any other key to cancel.")
            let confirm = term.getKey()
            if (confirm.toLowerCase() != "y") {
                continue
            }
        }
        break
    }
    
    // Check for save key (Ctrl+S)
    if (key == "Ctrl+S") {
        let filename = args["file"]
        if (!filename) {
            term.println("Enter filename to save: ")
            filename = term.input()
            args["file"] = filename
        }
        
        editor.save(ed, filename)
        continue
    }
    
    // Handle the keypress in the editor
    editor.handleKeypress(ed, key)
}

term.clear()
term.println("Goodbye!")
```

### Configuration Editor

```js
import editor
import term
import os
import json

// Function to load a JSON configuration file
let loadConfig = func(filename) {
    if (!os.exists(filename)) {
        return {
            "server": {
                "host": "localhost",
                "port": 8080
            },
            "database": {
                "host": "localhost",
                "port": 5432,
                "user": "admin",
                "password": "password"
            },
            "logging": {
                "level": "info",
                "file": "app.log"
            }
        }
    }
    
    let content = open(filename)
    return JSON.parse(content)
}

// Function to save a JSON configuration file
let saveConfig = func(filename, config) {
    let content = JSON.stringify(config, null, 2)
    os.writeFile(filename, content)
}

// Function to edit a configuration section
let editSection = func(section, data) {
    // Convert the section data to a string
    let text = JSON.stringify(data, null, 2)
    
    // Create an editor
    let ed = editor.newEditor({
        width: term.getSize().width,
        height: term.getSize().height - 4,
        syntax: "json"
    })
    
    // Set the text
    editor.setText(ed, text)
    
    // Edit loop
    term.clear()
    while (true) {
        // Render the editor
        let content = editor.render(ed)
        term.clear()
        term.println("Editing " + section + " configuration", "#ffcc00")
        term.println("Press Ctrl+S to save, Ctrl+Q to cancel", "#88ff88")
        print(content)
        
        // Get keypress
        let key = term.getKey()
        
        // Check for exit key (Ctrl+Q)
        if (key == "Ctrl+Q") {
            return null
        }
        
        // Check for save key (Ctrl+S)
        if (key == "Ctrl+S") {
            let newText = editor.getText(ed)
            try {
                return JSON.parse(newText)
            } catch (e) {
                term.println("Error parsing JSON: " + e, "#ff5555")
                term.println("Press any key to continue", "#88ff88")
                term.getKey()
                continue
            }
        }
        
        // Handle the keypress in the editor
        editor.handleKeypress(ed, key)
    }
}

// Main function
let main = func() {
    let configFile = "config.json"
    let config = loadConfig(configFile)
    
    while (true) {
        // Display menu
        term.clear()
        term.println("=== Configuration Editor ===", "#ffcc00")
        term.println("1. Edit Server Configuration", "#88ff88")
        term.println("2. Edit Database Configuration", "#88ff88")
        term.println("3. Edit Logging Configuration", "#88ff88")
        term.println("4. Save Configuration", "#88ff88")
        term.println("5. Exit", "#88ff88")
        term.println("===========================", "#ffcc00")
        term.print("Enter your choice: ")
        
        let choice = term.input()
        
        if (choice == "1") {
            let result = editSection("server", config.server)
            if (result) {
                config.server = result
            }
        } else if (choice == "2") {
            let result = editSection("database", config.database)
            if (result) {
                config.database = result
            }
        } else if (choice == "3") {
            let result = editSection("logging", config.logging)
            if (result) {
                config.logging = result
            }
        } else if (choice == "4") {
            saveConfig(configFile, config)
            term.println("Configuration saved to " + configFile, "#88ff88")
            term.println("Press any key to continue", "#88ff88")
            term.getKey()
        } else if (choice == "5") {
            break
        }
    }
    
    term.clear()
    term.println("Goodbye!", "#88ff88")
}

main()
```

### Markdown Editor with Preview

```js
import editor
import term
import os
import markdown

// Function to convert markdown to HTML
let markdownToHtml = func(text) {
    // This is a simplified markdown converter
    let html = text
    
    // Convert headers
    html = html.replace(/^# (.+)$/gm, "<h1>$1</h1>")
    html = html.replace(/^## (.+)$/gm, "<h2>$1</h2>")
    html = html.replace(/^### (.+)$/gm, "<h3>$1</h3>")
    
    // Convert bold and italic
    html = html.replace(/\*\*(.+?)\*\*/g, "<strong>$1</strong>")
    html = html.replace(/\*(.+?)\*/g, "<em>$1</em>")
    
    // Convert lists
    html = html.replace(/^- (.+)$/gm, "<li>$1</li>")
    
    // Convert links
    html = html.replace(/\[(.+?)\]\((.+?)\)/g, "<a href=\"$2\">$1</a>")
    
    return html
}

// Function to display a preview of the markdown
let previewMarkdown = func(text) {
    let html = markdownToHtml(text)
    
    // Display the HTML preview
    term.clear()
    term.println("=== Markdown Preview ===", "#ffcc00")
    term.println(html)
    term.println("========================", "#ffcc00")
    term.println("Press any key to return to editor", "#88ff88")
    term.getKey()
}

// Main function
let main = func() {
    // Check command line arguments
    if (args.length < 2) {
        term.println("Usage: vint markdown_editor.vint <filename>", "#ff5555")
        exit(1)
    }
    
    let filename = args[1]
    
    // Create an editor
    let ed = editor.newEditor({
        filename: filename,
        width: term.getSize().width,
        height: term.getSize().height - 2,
        syntax: "markdown"
    })
    
    // Main loop
    term.clear()
    while (true) {
        // Render the editor
        let content = editor.render(ed)
        term.clear()
        term.println("Markdown Editor - " + filename, "#ffcc00")
        term.println("Ctrl+S: Save | Ctrl+P: Preview | Ctrl+Q: Quit", "#88ff88")
        print(content)
        
        // Get keypress
        let key = term.getKey()
        
        // Check for exit key (Ctrl+Q)
        if (key == "Ctrl+Q") {
            if (editor.isModified(ed)) {
                term.println("There are unsaved changes. Press Y to quit anyway, or any other key to cancel.")
                let confirm = term.getKey()
                if (confirm.toLowerCase() != "y") {
                    continue
                }
            }
            break
        }
        
        // Check for save key (Ctrl+S)
        if (key == "Ctrl+S") {
            editor.save(ed, filename)
            continue
        }
        
        // Check for preview key (Ctrl+P)
        if (key == "Ctrl+P") {
            let text = editor.getText(ed)
            previewMarkdown(text)
            continue
        }
        
        // Handle the keypress in the editor
        editor.handleKeypress(ed, key)
    }
    
    term.clear()
    term.println("Goodbye!", "#88ff88")
}

main()
```

```

## email.md

```markdown
# Email Module in Vint

The Email module in Vint provides email validation and processing functions. This module helps you validate email addresses, extract components, and normalize email formats for consistent processing.

---

## Importing the Email Module

To use the Email module, simply import it:
```js
import email
```

---

## Functions and Examples

### 1. Validate Email Address (`validate`)

The `validate` function checks if a given string is a valid email address format.

**Syntax**:

```js
validate(emailAddress)
```

**Example**:

```js
import email

print("=== Email Validation Example ===")

// Valid email addresses
valid_emails = [
    "user@example.com",
    "john.doe@company.org",
    "test.email+tag@domain.co.uk"
]

// Invalid email addresses
invalid_emails = [
    "not-an-email",
    "@example.com",
    "user@",
    "user..name@example.com"
]

print("Valid Emails:")
for valid_email in valid_emails {
    result = email.validate(valid_email)
    print("  " + valid_email + " -> " + string(result))
}

print("\nInvalid Emails:")
for invalid_email in invalid_emails {
    result = email.validate(invalid_email)
    print("  " + invalid_email + " -> " + string(result))
}
```

---

### 2. Extract Domain (`extractDomain`)

The `extractDomain` function extracts the domain part from an email address.

**Syntax**:

```js
extractDomain(emailAddress)
```

**Example**:

```js
import email

print("=== Email Domain Extraction Example ===")
email_address = "john.doe@example.com"
domain = email.extractDomain(email_address)
print("Email:  ", email_address)
print("Domain: ", domain)
// Output: Domain: example.com
```

---

### 3. Extract Username (`extractUsername`)

The `extractUsername` function extracts the username part (before @) from an email address.

**Syntax**:

```js
extractUsername(emailAddress)
```

**Example**:

```js
import email

print("=== Email Username Extraction Example ===")
email_address = "john.doe@example.com"
username = email.extractUsername(email_address)
print("Email:    ", email_address)
print("Username: ", username)
// Output: Username: john.doe
```

---

### 4. Normalize Email (`normalize`)

The `normalize` function converts an email address to lowercase and trims whitespace for consistent formatting.

**Syntax**:

```js
normalize(emailAddress)
```

**Example**:

```js
import email

print("=== Email Normalization Example ===")
messy_email = "  John.DOE@EXAMPLE.COM  "
normalized = email.normalize(messy_email)
print("Original:   ", messy_email)
print("Normalized: ", normalized)
// Output: Normalized: john.doe@example.com
```

---

## Complete Usage Example

```js
import email

print("=== Email Module Complete Example ===")

// Process a list of email addresses
email_list = [
    "  Alice@COMPANY.COM  ",
    "bob.smith@example.org",
    "CHARLIE@DOMAIN.CO.UK",
    "invalid-email",
    "diana.jones@website.net"
]

print("Processing email addresses:")
for email_addr in email_list {
    print("\n--- Processing: " + email_addr + " ---")
    
    // Normalize the email first
    normalized = email.normalize(email_addr)
    print("Normalized: " + normalized)
    
    // Validate the email
    is_valid = email.validate(normalized)
    print("Valid: " + string(is_valid))
    
    if is_valid {
        // Extract components
        username = email.extractUsername(normalized)
        domain = email.extractDomain(normalized)
        
        print("Username: " + username)
        print("Domain: " + domain)
        
        // Example: Check for specific domains
        if domain == "company.com" {
            print("✓ Corporate email detected")
        } else {
            print("→ External email")
        }
    } else {
        print("✗ Invalid email format")
    }
}
```

---

## Use Cases

- **User Registration**: Validate email addresses during account creation
- **Email Lists**: Clean and normalize email addresses in mailing lists
- **Domain Analysis**: Extract domains for statistical analysis
- **Email Routing**: Route emails based on domain or username patterns
- **Data Cleaning**: Standardize email formats in databases

---

## Advanced Example: Email Domain Statistics

```js
import email

print("=== Email Domain Statistics ===")

emails = [
    "user1@gmail.com",
    "user2@company.com",
    "user3@gmail.com",
    "user4@yahoo.com",
    "user5@company.com",
    "user6@outlook.com"
]

domain_count = {}

for email_addr in emails {
    if email.validate(email_addr) {
        domain = email.extractDomain(email_addr)
        
        if domain in domain_count {
            domain_count[domain] = domain_count[domain] + 1
        } else {
            domain_count[domain] = 1
        }
    }
}

print("Domain statistics:")
for domain, count in domain_count {
    print("  " + domain + ": " + string(count) + " emails")
}
```

---

## Summary of Functions

| Function           | Description                                    | Return Type |
|--------------------|------------------------------------------------|-------------|
| `validate`         | Validates email address format                 | Boolean     |
| `extractDomain`    | Extracts domain part from email               | String      |
| `extractUsername`  | Extracts username part from email             | String      |
| `normalize`        | Normalizes email to lowercase and trims       | String      |

The Email module provides essential functionality for working with email addresses safely and efficiently in VintLang applications.

```

## encoding.md

```markdown
# encoding Module in Vint

The `encoding` module in Vint provides functions for encoding and decoding data in various formats. This includes commonly used encoding schemes like Base64.

---

## Importing the encoding Module

To use the `encoding` module, import it as follows:

```js
import encoding
```

---

## Functions and Examples

### 1. Base64 Encoding with `base64Encode()`

The `base64Encode` function encodes a string into Base64 format. Base64 encoding is often used for encoding binary data as text, making it suitable for transmission over text-based protocols such as email or HTTP.

**Syntax**:

```js
base64Encode(inputString)
```

- `inputString`: The string you want to encode.

**Example**:

```js
import encoding

encoded = encoding.base64Encode("Hello, World!")
print(encoded)  // Expected output: "SGVsbG8sIFdvcmxkIQ=="
```

In this example, the string `"Hello, World!"` is encoded into Base64 format.

---

### 2. Base64 Decoding with `base64Decode()`

The `base64Decode` function decodes a Base64-encoded string back into its original format.

**Syntax**:

```js
base64Decode(encodedString)
```

- `encodedString`: The Base64-encoded string that you want to decode.

**Example**:

```js
import encoding

encoded = encoding.base64Encode("Hello, World!")
print(encoded)  // Expected output: "SGVsbG8sIFdvcmxkIQ=="

decoded = encoding.base64Decode(encoded)
print(decoded)  // Expected output: "Hello, World!"
```

In this example, the Base64-encoded string is decoded back to its original value.

---

## Summary of Functions

| Function               | Description                                        | Example Output                             |
|------------------------|----------------------------------------------------|--------------------------------------------|
| `base64Encode(input)`   | Encodes a string to Base64 format.                 | `"SGVsbG8sIFdvcmxkIQ=="`                   |
| `base64Decode(encoded)` | Decodes a Base64-encoded string back to its original form. | `"Hello, World!"`                           |

---

The `encoding` module in Vint is essential for working with different encoding schemes such as Base64. It simplifies the process of converting data between text and binary formats, making it easier to handle data transmission or storage in encoded formats.

```

## enhanced-errors.md

```markdown
# Enhanced Error Messages in Vint Core Modules

This document describes the enhanced error handling system implemented across Vint's core modules using the new `ErrorMessage` helper function. The improvements provide consistent, colorized, and highly descriptive error messages that guide users toward correct usage.

## New ErrorMessage Helper Function

All modules now use a centralized `ErrorMessage` function from `module/module.go`:

```go
func ErrorMessage(module, function, expected, received, usage string) *object.Error {
    return &object.Error{
        Message: fmt.Sprintf(
            "\033[1;31m -> %s.%s()\033[0m:\n"+
                "  Expected: %s\n"+
                "  Received: %s\n"+
                "  Usage: %s\n"+
                "  See documentation for details.\n",
            module, function, expected, received, usage,
        ),
    }
}
```

## What Was Enhanced

### 1. **CLI Module** (`import cli`)

- **Functions**: `cli.prompt()`, `cli.confirm()`, `cli.execCommand()`, `cli.exit()`, `cli.hasArg()`, `cli.getArgValue()`, `cli.getPositional()`
- **Improvements**:
  - Consistent error formatting with colors
  - Clear expected vs received information
  - Practical usage examples
  - Descriptive parameter names

**New Format Example**:

```
Error in cli.prompt():
  Expected: 1 string argument (prompt message)
  Received: 2 arguments
  Usage: cli.prompt("Enter your name: ") -> returns user input
  See documentation for details.
```

### 2. **Net Module** (`import net`)

- **Functions**: `net.get()`, `net.post()`, `net.put()`, `net.delete()`, `net.patch()`
- **Improvements**:
  - Parameter-specific error messages
  - Network operation context
  - HTTP method-specific examples

**New Format Example**:

```
Error in net.get():
  Expected: string value for 'url' parameter
  Received: INTEGER
  Usage: net.get(url="https://example.com")
  See documentation for details.
```

### 3. **OS Module** (`import os`)

- **Functions**: `os.run()`, `os.getEnv()`
- **Improvements**:
  - Command execution context
  - System operation guidance
  - Clear parameter descriptions

### 4. **Math Module** (`import math`)

- **Functions**: `math.abs()` and similar numeric functions
- **Improvements**:
  - Mathematical operation context
  - Numeric type specifications
  - Calculation examples

### 5. **Time Module** (`import time`)

- **Functions**: `time.now()`, `time.sleep()`
- **Improvements**:
  - Time operation context
  - Duration specifications
  - Temporal examples

### 6. **Crypto Module** (`import crypto`)

- **Functions**: `crypto.hashMD5()`, `crypto.hashSHA256()`, `crypto.encryptAES()`, `crypto.decryptAES()`
- **Improvements**:
  - Cryptographic operation context
  - Security-related guidance
  - Encryption examples

### 7. **Colors Module** (`import colors`)

- **Functions**: `colors.rgbToHex()`
- **Improvements**:
  - Color value specifications
  - Range validation
  - Visual examples

### 8. **String Module** (`import string`)

- **Functions**: `string.slug()` and others
- **Improvements**:
  - Text processing context
  - String manipulation examples

## Error Message Format

All enhanced error messages follow this consistent pattern:

```
Error in [module].[function]():
  Expected: [clear description of expected input]
  Received: [what was actually provided]
  Usage: [practical example with expected output]
  See documentation for details.
```

### Key Features

1. **🎨 Color Coding**: Red highlighting for error identification
2. **📝 Clear Structure**: Consistent four-line format
3. **🔍 Specific Details**: Exact expected vs received information
4. **💡 Usage Examples**: Practical code examples
5. **📚 Documentation Reference**: Pointer to additional help

## Benefits of the New System

### For Developers

- **Instant Recognition**: Red coloring makes errors immediately visible
- **Clear Guidance**: Know exactly what's expected vs what was provided
- **Learn by Example**: Usage examples teach correct syntax
- **Consistent Experience**: Same error format across all modules
- **Reduced Debugging Time**: Precise error information speeds up fixes

### For the Language

- **Professional Appearance**: Consistent, polished error messages
- **Better Learning Curve**: New users learn faster with clear examples
- **Maintainability**: Centralized error formatting makes updates easier
- **Extensibility**: Easy to add new modules using the same pattern

## Examples by Category

### Argument Count Errors

```
Error in time.sleep():
  Expected: 1 numeric argument (seconds to sleep)
  Received: 2 arguments
  Usage: time.sleep(5) -> sleeps for 5 seconds
  See documentation for details.
```

### Type Errors

```
Error in math.abs():
  Expected: numeric argument (integer or float)
  Received: STRING
  Usage: math.abs(-5) -> 5
  See documentation for details.
```

### Parameter-Specific Errors

```
Error in net.post():
  Expected: dictionary value for 'headers' parameter
  Received: STRING
  Usage: net.post(headers={"Content-Type": "application/json"})
  See documentation for details.
```

### Range Validation Errors

```
Error in colors.rgbToHex():
  RGB values must be in the range 0-255.
  Usage: colors.rgbToHex(255, 0, 128) -> "#FF0080"
```

## Testing Enhanced Errors

Use the provided `new_error_format_test.vint` file to see all the improved error messages in action:

```bash
vint new_error_format_test.vint
```

## Implementation Guidelines

When adding new functions or modules, use the `ErrorMessage` helper:

```go
if len(args) != expectedCount {
    return ErrorMessage(
        "moduleName", "functionName",
        "description of expected arguments",
        fmt.Sprintf("%d arguments", len(args)),
        "usage.example() -> expected output",
    )
}

if args[0].Type() != expectedType {
    return ErrorMessage(
        "moduleName", "functionName", 
        "description of expected type",
        string(args[0].Type()),
        "usage.example() -> expected output",
    )
}
```

## Future Enhancements

The ErrorMessage system enables future improvements:

1. **Error Codes**: Add numeric codes for programmatic handling
2. **Suggestions**: Auto-suggest corrections for common mistakes
3. **Localization**: Multi-language error messages
4. **Context Awareness**: Errors that understand the calling context
5. **Interactive Help**: Links to relevant documentation sections

This enhanced error system represents a significant improvement in Vint's developer experience, making the language more approachable and professional.

```

## error.md

```markdown
# Error

The `error` keyword allows you to raise a fatal runtime error that stops the execution of the script.

## Syntax

`error "Your error message here"`

When the interpreter encounters an `error` statement, it will print a formatted error message to the console and halt execution immediately. This is useful for handling critical problems where the program cannot safely continue.

### Example

```js
let file_path = "data.json"
if !fs.exists(file_path) {
    error "Critical file 'data.json' not found."
}
println("This will not be printed if the file is missing.")
```

If `data.json` does not exist, running this script will output:

```
Error: Critical file 'data.json' not found.
```

And the script will stop.

```

## errors.md

```markdown
# Errors Module

The `errors` module provides a way to create new errors.

## `errors.new(message)`

Creates a new error with the given message. This will stop the execution of the script.

### Parameters

- `message` (string): The error message.

### Example

```js
import "errors"

errors.new("something went wrong")
# The script will stop here and print the error message
```

```

## excel.md

```markdown
# Excel Module in VintLang

The `excel` module provides comprehensive Excel file manipulation capabilities, including reading, writing, formatting, password handling, and advanced features. This module supports both `.xlsx` and `.xls` formats and offers extensive functionality for working with Excel spreadsheets.

---

## Importing the Excel Module

```js
import excel
```

---

## Core Features

- **File Operations**: Create, open, save Excel files with password support
- **Sheet Management**: Add, delete, rename sheets, manage active sheets
- **Cell Operations**: Read/write individual cells, formulas, and styling
- **Range Operations**: Work with cell ranges, copy, clear data
- **Row/Column Management**: Insert, delete, resize rows and columns  
- **Formatting**: Merge cells, set fonts, borders, colors, alignment
- **Advanced Features**: Tables, charts, images, comments
- **Data Exchange**: Convert to/from CSV, JSON formats
- **Password Protection**: Open password-protected files
- **Search & Replace**: Find and replace text across sheets

---

## File Operations

### Create New Excel File

#### `excel.create(filepath?)`

Creates a new Excel workbook. Optionally saves it to the specified path.

```js
import excel

// Create in memory
file_id = excel.create()

// Create and save to file
file_id = excel.create("new_workbook.xlsx")
```

### Open Existing File

#### `excel.open(filepath)`

Opens an existing Excel file.

```js
file_id = excel.open("existing_workbook.xlsx")
```

#### `excel.openWithPassword(filepath, password)`

Opens a password-protected Excel file.

```js
file_id = excel.openWithPassword("protected_workbook.xlsx", "secret123")
```

### Save Operations

#### `excel.save(file_id)`

Saves the current file.

```js
excel.save(file_id)
```

#### `excel.saveAs(file_id, new_filepath)`

Saves the file with a new name or location.

```js
new_file_id = excel.saveAs(file_id, "backup_workbook.xlsx")
```

#### `excel.close(file_id)`

Closes the Excel file and frees memory.

```js
excel.close(file_id)
```

---

## Sheet Management

### Get Sheet Information

#### `excel.getSheets(file_id)`

Returns an array of all sheet names.

```js
sheets = excel.getSheets(file_id)
print("Available sheets:", sheets)
// Output: ["Sheet1", "Data", "Summary"]
```

### Add and Delete Sheets

#### `excel.addSheet(file_id, sheet_name)`

Adds a new worksheet.

```js
sheet_index = excel.addSheet(file_id, "NewSheet")
print("Created sheet at index:", sheet_index)
```

#### `excel.deleteSheet(file_id, sheet_name)`

Deletes a worksheet.

```js
excel.deleteSheet(file_id, "OldSheet")
```

### Rename Sheet

#### `excel.renameSheet(file_id, old_name, new_name)`

Renames an existing worksheet.

```js
excel.renameSheet(file_id, "Sheet1", "DataSheet")
```

### Active Sheet Management

#### `excel.setActiveSheet(file_id, sheet_index)`

Sets the active worksheet by index.

```js
excel.setActiveSheet(file_id, 0)  // Make first sheet active
```

#### `excel.getActiveSheet(file_id)`

Gets the index of the currently active sheet.

```js
active_index = excel.getActiveSheet(file_id)
```

---

## Cell Operations

### Read and Write Cells

#### `excel.getCell(file_id, sheet_name, cell_reference)`

Reads the value from a specific cell.

```js
value = excel.getCell(file_id, "Sheet1", "A1")
print("Cell A1 contains:", value)
```

#### `excel.setCell(file_id, sheet_name, cell_reference, value)`

Writes a value to a specific cell.

```js
excel.setCell(file_id, "Sheet1", "A1", "Hello World")
excel.setCell(file_id, "Sheet1", "B1", 42)
excel.setCell(file_id, "Sheet1", "C1", 3.14)
excel.setCell(file_id, "Sheet1", "D1", true)
```

### Formula Operations

#### `excel.getCellFormula(file_id, sheet_name, cell_reference)`

Gets the formula from a cell.

```js
formula = excel.getCellFormula(file_id, "Sheet1", "E1")
print("Formula:", formula)
```

#### `excel.setCellFormula(file_id, sheet_name, cell_reference, formula)`

Sets a formula in a cell.

```js
excel.setCellFormula(file_id, "Sheet1", "E1", "=SUM(B1:D1)")
excel.setCellFormula(file_id, "Sheet1", "F1", "=AVERAGE(B1:D1)")
excel.setCellFormula(file_id, "Sheet1", "G1", "=IF(B1>0,\"Positive\",\"Zero or Negative\")")
```

---

## Range Operations

### Work with Cell Ranges

#### `excel.getRange(file_id, sheet_name, range_reference)`

Gets data from a cell range as a 2D array.

```js
data = excel.getRange(file_id, "Sheet1", "A1:C3")
// Returns nested arrays: [["A1", "B1", "C1"], ["A2", "B2", "C2"], ["A3", "B3", "C3"]]

// Access specific cell from range
print("First row:", data[0])
print("Cell B2:", data[1][1])
```

#### `excel.setRange(file_id, sheet_name, range_reference, data)`

Sets data for a cell range using a 2D array.

```js
headers = ["Name", "Age", "City"]
data = [
    headers,
    ["John", 30, "New York"],
    ["Jane", 25, "Boston"],
    ["Bob", 35, "Chicago"]
]

excel.setRange(file_id, "Sheet1", "A1:C4", data)
```

---

## Row and Column Management

### Insert and Delete Operations

#### `excel.insertRow(file_id, sheet_name, row_number)`

Inserts a new row at the specified position.

```js
excel.insertRow(file_id, "Sheet1", 2)  // Insert row at position 2
```

#### `excel.insertColumn(file_id, sheet_name, column_letter)`

Inserts a new column at the specified position.

```js
excel.insertColumn(file_id, "Sheet1", "B")  // Insert column B
```

#### `excel.deleteRow(file_id, sheet_name, row_number)`

Deletes a row.

```js
excel.deleteRow(file_id, "Sheet1", 3)  // Delete row 3
```

#### `excel.deleteColumn(file_id, sheet_name, column_letter)`

Deletes a column.

```js
excel.deleteColumn(file_id, "Sheet1", "C")  // Delete column C
```

---

## Cell Formatting

### Merge and Unmerge Cells

#### `excel.mergeCells(file_id, sheet_name, range_reference)`

Merges cells in the specified range.

```js
excel.mergeCells(file_id, "Sheet1", "A1:C1")  // Merge header cells
```

#### `excel.unmergeCells(file_id, sheet_name, range_reference)`

Unmerges previously merged cells.

```js
excel.unmergeCells(file_id, "Sheet1", "A1:C1")
```

---

## File Information and Utilities

### Get File Information

#### `excel.getFileInfo(file_id)`

Returns comprehensive information about the Excel file.

```js
info = excel.getFileInfo(file_id)
print("Number of sheets:", info.sheetCount)
print("Active sheet index:", info.activeSheet)
print("Sheet names:", info.sheets)
```

---

## Complete Usage Example

```js
import excel

print("=== Excel Module Complete Example ===")

// Create a new workbook
file_id = excel.create("employee_data.xlsx")

// Add a new sheet for employee data
excel.addSheet(file_id, "Employees")
excel.renameSheet(file_id, "Sheet1", "Summary")

// Set up employee data
headers = ["ID", "Name", "Department", "Salary", "Bonus"]
employees = [
    [1, "John Doe", "Engineering", 75000, "=D2*0.1"],
    [2, "Jane Smith", "Marketing", 65000, "=D3*0.1"], 
    [3, "Bob Johnson", "Sales", 70000, "=D4*0.15"],
    [4, "Alice Brown", "HR", 60000, "=D5*0.1"]
]

// Write headers
for i = 0; i < headers.length; i++ {
    cell_ref = string.char(65 + i) + "1"  // A1, B1, C1, etc.
    excel.setCell(file_id, "Employees", cell_ref, headers[i])
}

// Write employee data
for row = 0; row < employees.length; row++ {
    for col = 0; col < employees[row].length - 1; col++ {
        cell_ref = string.char(65 + col) + (row + 2)
        excel.setCell(file_id, "Employees", cell_ref, employees[row][col])
    }
    // Set formula for bonus calculation
    bonus_cell = "E" + (row + 2)
    excel.setCellFormula(file_id, "Employees", bonus_cell, employees[row][4])
}

// Merge header row for title
excel.setCell(file_id, "Employees", "A1", "Employee Database")
excel.mergeCells(file_id, "Employees", "A1:E1")

// Add summary in Summary sheet
excel.setCell(file_id, "Summary", "A1", "Summary Report")
excel.setCell(file_id, "Summary", "A3", "Total Employees:")
excel.setCellFormula(file_id, "Summary", "B3", "=COUNTA(Employees.A:A)-1")

excel.setCell(file_id, "Summary", "A4", "Average Salary:")
excel.setCellFormula(file_id, "Summary", "B4", "=AVERAGE(Employees.D:D)")

excel.setCell(file_id, "Summary", "A5", "Total Payroll (with bonuses):")
excel.setCellFormula(file_id, "Summary", "B5", "=SUM(Employees.D:D)+SUM(Employees.E:E)")

// Get file information
info = excel.getFileInfo(file_id)
print("Created workbook with", info.sheetCount, "sheets:")
for sheet in info.sheets {
    print("-", sheet)
}

// Save and close
excel.save(file_id)
excel.close(file_id)

print("Excel file 'employee_data.xlsx' created successfully!")

// Re-open to read data
file_id = excel.open("employee_data.xlsx")

// Read some data back
employee_name = excel.getCell(file_id, "Employees", "B2")
employee_salary = excel.getCell(file_id, "Employees", "D2") 
total_employees = excel.getCell(file_id, "Summary", "B3")

print("First employee:", employee_name, "- Salary:", employee_salary)
print("Total employees:", total_employees)

excel.close(file_id)
```

---

## Advanced Features

### Password Protection

```js
// Open password-protected file
file_id = excel.openWithPassword("secure_data.xlsx", "mypassword")

// Note: Setting passwords on existing files is not yet supported
// in the current version of the excelize library
```

### Working with Multiple Sheets

```js
// Process multiple sheets in a workbook
sheets = excel.getSheets(file_id)
for sheet_name in sheets {
    print("Processing sheet:", sheet_name)
    
    // Get all data from each sheet  
    data = excel.getRange(file_id, sheet_name, "A1:Z100")
    
    // Process data...
    print("Sheet", sheet_name, "has", data.length, "rows")
}
```

### Data Validation and Processing

```js
// Read and validate data
raw_data = excel.getRange(file_id, "Data", "A1:D10")

cleaned_data = []
for row in raw_data {
    if row[0] != "" && row[0] != null {  // Skip empty rows
        cleaned_row = []
        for cell in row {
            // Clean and validate data
            if cell != null {
                cleaned_row.push(string.trim(cell))
            } else {
                cleaned_row.push("")
            }
        }
        cleaned_data.push(cleaned_row)
    }
}

// Write cleaned data back
excel.setRange(file_id, "CleanedData", "A1:D" + cleaned_data.length, cleaned_data)
```

---

## Error Handling

Always use proper error handling when working with Excel files:

```js
// Safe file operations
try {
    file_id = excel.open("data.xlsx")
    data = excel.getCell(file_id, "Sheet1", "A1")
    print("Data:", data)
} catch error {
    print("Error reading Excel file:", error)
} finally {
    if file_id {
        excel.close(file_id)
    }
}
```

---

## Use Cases

- **Data Analysis**: Read Excel reports and perform calculations
- **Report Generation**: Create formatted Excel reports from application data
- **Data Import/Export**: Convert between Excel and other formats
- **Template Processing**: Fill Excel templates with dynamic data
- **Financial Modeling**: Build spreadsheets with complex formulas
- **Batch Processing**: Process multiple Excel files programmatically
- **Data Migration**: Transfer data between different systems via Excel
- **Automated Reporting**: Generate periodic reports in Excel format

---

## Summary of Functions

### File Operations

| Function | Description | Return Type |
|----------|-------------|-------------|
| `create(filepath?)` | Create new Excel file | String (file_id) |
| `open(filepath)` | Open existing Excel file | String (file_id) |
| `openWithPassword(filepath, password)` | Open password-protected file | String (file_id) |
| `save(file_id)` | Save current file | Boolean |
| `saveAs(file_id, filepath)` | Save file with new name | String (new_file_id) |
| `close(file_id)` | Close and cleanup file | Boolean |

### Sheet Management  

| Function | Description | Return Type |
|----------|-------------|-------------|
| `getSheets(file_id)` | Get list of sheet names | Array |
| `addSheet(file_id, name)` | Add new sheet | Integer (index) |
| `deleteSheet(file_id, name)` | Delete sheet | Boolean |
| `renameSheet(file_id, old, new)` | Rename sheet | Boolean |
| `setActiveSheet(file_id, index)` | Set active sheet | Boolean |
| `getActiveSheet(file_id)` | Get active sheet index | Integer |

### Cell Operations

| Function | Description | Return Type |
|----------|-------------|-------------|
| `getCell(file_id, sheet, cell)` | Read cell value | String |
| `setCell(file_id, sheet, cell, value)` | Write cell value | Boolean |
| `getCellFormula(file_id, sheet, cell)` | Get cell formula | String |
| `setCellFormula(file_id, sheet, cell, formula)` | Set cell formula | Boolean |

### Range Operations

| Function | Description | Return Type |
|----------|-------------|-------------|
| `getRange(file_id, sheet, range)` | Get range data | Array (2D) |
| `setRange(file_id, sheet, range, data)` | Set range data | Boolean |

### Row/Column Operations

| Function | Description | Return Type |
|----------|-------------|-------------|
| `insertRow(file_id, sheet, row)` | Insert row | Boolean |
| `insertColumn(file_id, sheet, col)` | Insert column | Boolean |
| `deleteRow(file_id, sheet, row)` | Delete row | Boolean |
| `deleteColumn(file_id, sheet, col)` | Delete column | Boolean |

### Formatting

| Function | Description | Return Type |
|----------|-------------|-------------|
| `mergeCells(file_id, sheet, range)` | Merge cells | Boolean |
| `unmergeCells(file_id, sheet, range)` | Unmerge cells | Boolean |

### Utilities

| Function | Description | Return Type |
|----------|-------------|-------------|
| `getFileInfo(file_id)` | Get file information | Dictionary |

The Excel module provides a powerful and comprehensive interface for working with Excel files in VintLang, supporting both simple data operations and advanced spreadsheet manipulation.

```

## files.md

```markdown
# Files in Vint

The `files` module in Vint provides comprehensive functionality for working with files, including reading, writing, file operations, and metadata access.

---

## Opening a File

You can open a file using the `open` keyword. This will return an object of type `FAILI`, which represents the file.

### Syntax

```js
fileObject = open("filename.txt")
```

### Example

```js
myFile = open("file.txt")

aina(myFile) // Output: FAILI
```

---

## File Methods

File objects in Vint come with powerful built-in methods for comprehensive file operations:

### read()

Read the entire contents of a file as a string:

```js
myFile = open("example.txt")
contents = myFile.read()
print(contents)
```

### write()

Write content to a file (overwrites existing content):

```js
myFile = open("output.txt")
myFile.write("Hello, World!")
```

### append()

Append content to the end of a file:

```js
logFile = open("app.log")
logFile.append("New log entry\n")
```

### exists()

Check if the file exists:

```js
myFile = open("config.txt")
if (myFile.exists()) {
    print("File exists!")
} else {
    print("File not found!")
}
```

### size()

Get the size of the file in bytes:

```js
myFile = open("data.txt")
fileSize = myFile.size()
print("File size:", fileSize, "bytes")
```

### delete()

Delete the file from the filesystem:

```js
tempFile = open("temp.txt")
if (tempFile.exists()) {
    tempFile.delete()
    print("File deleted successfully")
}
```

### copy()

Copy the file to a new location:

```js
sourceFile = open("original.txt")
sourceFile.copy("backup.txt")
print("File copied successfully")
```

### move()

Move or rename the file:

```js
oldFile = open("old_name.txt")
oldFile.move("new_name.txt")
print("File moved/renamed successfully")
```

### lines()

Read the file content as an array of lines:

```js
configFile = open("settings.conf")
lines = configFile.lines()
for line in lines {
    print("Config:", line.trim())
}
```

### extension()

Get the file extension:

```js
documentFile = open("report.pdf")
ext = documentFile.extension()
print("File extension:", ext)  // .pdf

imageFile = open("photo.jpg")
print("Extension:", imageFile.extension())  // .jpg
```

## Practical File Examples

Here are some practical examples using file methods:

```js
// Log file manager
let log_message = func(message) {
    let logFile = open("application.log")
    let timestamp = time.now().format("2006-01-02 15:04:05")
    let entry = "[" + timestamp + "] " + message + "\n"
    logFile.append(entry)
}

log_message("Application started")
log_message("User logged in")

// File backup system
let backup_file = func(filename) {
    let sourceFile = open(filename)
    if (sourceFile.exists()) {
        let backup_name = filename + ".backup"
        sourceFile.copy(backup_name)
        print("Backup created:", backup_name)
        return true
    } else {
        print("Source file not found:", filename)
        return false
    }
}

backup_file("important_data.txt")

// Configuration file processor
let process_config = func(config_file) {
    let file = open(config_file)
    if (!file.exists()) {
        print("Config file not found, creating default...")
        file.write("debug=false\nport=8080\nhost=localhost\n")
        return
    }
    
    let lines = file.lines()
    let settings = {}
    
    for line in lines {
        if (line.contains("=")) {
            let parts = line.split("=")
            if (parts.length() == 2) {
                settings.set(parts[0].trim(), parts[1].trim())
            }
        }
    }
    
    print("Loaded settings:", settings)
    return settings
}

let config = process_config("app.conf")

// File size analyzer
let analyze_files = func(filenames) {
    let total_size = 0
    let file_info = []
    
    for filename in filenames {
        let file = open(filename)
        if (file.exists()) {
            let size = file.size()
            let ext = file.extension()
            total_size += size
            
            file_info.push({
                "name": filename,
                "size": size,
                "extension": ext
            })
        }
    }
    
    print("Total size:", total_size, "bytes")
    print("File details:")
    for info in file_info {
        print("  ", info["name"], "-", info["size"], "bytes", info["extension"])
    }
}

analyze_files(["document.pdf", "image.jpg", "data.csv"])

// Text file processor with method chaining
let process_text_file = func(input_file, output_file) {
    let inputFile = open(input_file)
    
    if (!inputFile.exists()) {
        print("Input file not found")
        return false
    }
    
    // Read and process content
    let content = inputFile.read()
    let processed = content.upper().replace("OLD", "NEW").trim()
    
    // Write to output file
    let outputFile = open(output_file)
    outputFile.write(processed)
    
    print("Processing complete:")
    print("  Input size:", inputFile.size(), "bytes")
    print("  Output size:", outputFile.size(), "bytes")
    print("  Output extension:", outputFile.extension())
    
    return true
}

process_text_file("input.txt", "output.txt")

// File cleanup utility
let cleanup_temp_files = func(directory_pattern) {
    let temp_files = ["temp1.tmp", "cache.tmp", "old_data.bak"]
    let deleted_count = 0
    
    for filename in temp_files {
        let file = open(filename)
        if (file.exists()) {
            let size = file.size()
            file.delete()
            print("Deleted:", filename, "(" + size.to_string() + " bytes)")
            deleted_count++
        }
    }
    
    print("Cleanup complete. Deleted", deleted_count, "files")
}

cleanup_temp_files("*.tmp")
```

## File Error Handling

When working with files, it's important to handle potential errors:

```js
let safe_file_operation = func(filename, operation) {
    let file = open(filename)
    
    // Always check if file exists for read operations
    if (operation == "read" && !file.exists()) {
        print("Error: File", filename, "does not exist")
        return null
    }
    
    // Get file info before operations
    if (file.exists()) {
        print("File info:")
        print("  Size:", file.size(), "bytes")
        print("  Extension:", file.extension())
    }
    
    // Perform the operation
    if (operation == "read") {
        return file.read()
    } else if (operation == "backup") {
        let backup_name = filename + ".backup"
        file.copy(backup_name)
        return backup_name
    }
    
    return null
}

// Safe file reading
let content = safe_file_operation("data.txt", "read")
if (content != null) {
    print("File content loaded successfully")
}
```

---

## Notes

- All file operations are performed relative to the current working directory unless an absolute path is specified
- File methods support method chaining for fluent operations
- Always check if a file exists before performing read operations
- The `lines()` method automatically handles different line ending formats
- File extensions are returned with the leading dot (e.g., ".txt", ".pdf")

---

```

## filewatcher.md

```markdown
# VintLang FileWatcher Module

The `filewatcher` module provides functionality to monitor files and directories for changes. This is useful for building file watchers, auto-reloaders, build tools, and other applications that need to react to file system changes.

## Features

- Watch individual files for modifications
- Watch directories for file creation, modification, and deletion
- Recursive directory watching
- File extension filtering
- Configurable polling intervals
- Event-based callbacks with detailed information

## Usage

### Basic Example

```js
import filewatcher

// Watch a single file
let watcherId = filewatcher.watch("config.json", func(event) {
    print("File changed:", event["path"])
    print("Event type:", event["type"])
    print("Time:", event["time"])
    
    // Read the updated file
    let content = open(event["path"])
    print("New content:", content)
})

// Watch a directory
let dirWatcherId = filewatcher.watchDir("src", func(event) {
    print("File system change detected:")
    print("  Path:", event["path"])
    print("  Type:", event["type"]) // "created", "modified", or "deleted"
    print("  Time:", event["time"])
    
    // React to the change
    if (event["type"] == "created") {
        print("New file created!")
    } else if (event["type"] == "modified") {
        print("File was modified")
    } else if (event["type"] == "deleted") {
        print("File was deleted")
    }
})

// Stop watching after some time
setTimeout(func() {
    filewatcher.stopWatch(watcherId)
    filewatcher.stopWatch(dirWatcherId)
    print("Stopped watching")
}, 60000) // Stop after 60 seconds
```

## API Reference

### watch(path, callback, options)

Watches a file for changes and calls a callback function when changes are detected.

**Parameters:**

- `path` (string): The path to the file to watch
- `callback` (function): The function to call when the file changes. The callback receives an event object with the following properties:
  - `path` (string): The path to the file that changed
  - `type` (string): The type of change (always "modified" for single file watching)
  - `time` (string): The timestamp of the change
- `options` (dict, optional): Options for the watcher
  - `interval` (integer): The polling interval in milliseconds (default: 1000)

**Returns:**

- A watcher ID string that can be used to stop watching

### watchDir(path, callback, options)

Watches a directory for changes and calls a callback function when changes are detected.

**Parameters:**

- `path` (string): The path to the directory to watch
- `callback` (function): The function to call when changes are detected. The callback receives an event object with the following properties:
  - `path` (string): The path to the file that changed
  - `type` (string): The type of change ("created", "modified", or "deleted")
  - `time` (string): The timestamp of the change
- `options` (dict, optional): Options for the watcher
  - `interval` (integer): The polling interval in milliseconds (default: 1000)
  - `recursive` (boolean): Whether to watch subdirectories recursively (default: false)
  - `extensions` (array): Array of file extensions to watch (e.g., [".js", ".vint"]). If not provided, all files are watched.

**Returns:**

- A watcher ID string that can be used to stop watching

### stopWatch(watcherId)

Stops a file or directory watcher.

**Parameters:**

- `watcherId` (string): The watcher ID returned by `watch` or `watchDir`

**Returns:**

- `true` if the watcher was stopped successfully, `false` otherwise

### isWatching(path)

Checks if a file or directory is being watched.

**Parameters:**

- `path` (string): The path to check

**Returns:**

- `true` if the path is being watched, `false` otherwise

## Examples

### Auto-Reloading Development Server

```js
import filewatcher
import http
import os

// Simple HTTP server
let server = http.createServer(func(req, res) {
    if (req.path == "/") {
        // Serve index.html
        let content = open("public/index.html")
        res.writeHead(200, {"Content-Type": "text/html"})
        res.end(content)
    } else if (req.path == "/app.js") {
        // Serve app.js
        let content = open("public/app.js")
        res.writeHead(200, {"Content-Type": "application/javascript"})
        res.end(content)
    } else if (req.path == "/style.css") {
        // Serve style.css
        let content = open("public/style.css")
        res.writeHead(200, {"Content-Type": "text/css"})
        res.end(content)
    } else {
        // 404 Not Found
        res.writeHead(404)
        res.end("Not Found")
    }
})

// Start the server
server.listen(8080)
print("Server running at http://localhost:8080/")

// Set up WebSocket for live reload
let wsServer = http.createWebSocketServer(server)
let clients = []

wsServer.on("connection", func(client) {
    print("New client connected")
    clients.push(client)
    
    client.on("close", func() {
        // Remove client when disconnected
        let index = clients.indexOf(client)
        if (index != -1) {
            clients.splice(index, 1)
        }
    })
})

// Watch the public directory for changes
filewatcher.watchDir("public", func(event) {
    print("File changed:", event["path"])
    
    // Notify all connected clients to reload
    for (let client in clients) {
        client.send(JSON.stringify({
            type: "reload",
            path: event["path"]
        }))
    }
}, {
    recursive: true,
    extensions: [".html", ".js", ".css"]
})

print("Watching public directory for changes...")
```

### Build Tool

```js
import filewatcher
import os
import shell

// Function to build the project
let buildProject = func() {
    print("Building project...")
    
    // Compile all .vint files to .js
    let files = os.listDir("src")
    for (let file in files) {
        if (file.endsWith(".vint")) {
            let inputPath = "src/" + file
            let outputPath = "dist/" + file.replace(".vint", ".js")
            
            print("Compiling", inputPath, "to", outputPath)
            shell.exec("vint compile " + inputPath + " -o " + outputPath)
        }
    }
    
    // Bundle the JavaScript files
    print("Bundling JavaScript...")
    shell.exec("webpack --config webpack.config.js")
    
    print("Build completed!")
}

// Ensure dist directory exists
if (!os.exists("dist")) {
    os.mkdir("dist")
}

// Initial build
buildProject()

// Watch for changes
print("Watching for changes...")
filewatcher.watchDir("src", func(event) {
    if (event["type"] == "modified" || event["type"] == "created") {
        if (event["path"].endsWith(".vint")) {
            print("Source file changed:", event["path"])
            buildProject()
        }
    }
}, {
    recursive: true,
    extensions: [".vint"]
})

print("Build watcher started. Press Ctrl+C to stop.")
```

### Log File Monitor

```js
import filewatcher
import term

// Function to display the last N lines of a file
let tailFile = func(filePath, lines) {
    let content = open(filePath)
    let allLines = content.split("\n")
    let lastLines = allLines.slice(Math.max(0, allLines.length - lines))
    
    // Clear the screen
    term.clear()
    
    // Print header
    term.println("=== Log File Monitor ===", "#ffcc00")
    term.println("File: " + filePath, "#88ff88")
    term.println("Last " + lines + " lines:", "#88ff88")
    term.println("----------------------------", "#ffcc00")
    
    // Print the lines with syntax highlighting
    for (let line in lastLines) {
        if (line.includes("ERROR")) {
            term.println(line, "#ff5555") // Red for errors
        } else if (line.includes("WARNING")) {
            term.println(line, "#ffaa55") // Orange for warnings
        } else if (line.includes("INFO")) {
            term.println(line, "#55aaff") // Blue for info
        } else {
            term.println(line) // Default color for other lines
        }
    }
    
    // Print footer
    term.println("----------------------------", "#ffcc00")
    term.println("Press Ctrl+C to exit", "#88ff88")
}

// Check command line arguments
if (args.length < 2) {
    term.println("Usage: vint logmonitor.vint <log_file> [lines]", "#ff5555")
    exit(1)
}

let logFile = args[1]
let lines = 10 // Default to 10 lines

if (args.length >= 3) {
    lines = parseInt(args[2])
}

// Initial display
tailFile(logFile, lines)

// Watch the log file for changes
filewatcher.watch(logFile, func(event) {
    tailFile(logFile, lines)
}, {
    interval: 500 // Check every 500ms
})

print("Monitoring log file. Press Ctrl+C to exit.")
```

```

## for.md

```markdown
# For Loops in vint

For loops are a fundamental control structure in vint, used for iterating over iterable objects such as strings, arrays, and dictionaries. This page covers the syntax and usage of for loops in vint, including key-value pair iteration, and the use of break and continue statements.

## Basic Syntax

To create a for loop, use the for keyword followed by a temporary identifier (such as i or v) and the iterable object. Enclose the loop body in curly braces {}. Here's an example with a string:

```s
name = "hello"

for i in name {
    print(i)
}
```

Output:

```s
h
e
l
l
o
```

## Iterating Over Key-Value Pairs

### Dictionaries

vint allows you to iterate over both the value or the key-value pair of an iterable. To iterate over just the values, use one temporary identifier:

```s
dict = {"a": "apple", "b": "banana"}

for v in dict {
    print(v)
}
```

Output:

```s
apple
banana
```

To iterate over both the keys and the values, use two temporary identifiers:

```s
for k, v in dict {
    print(k + " is " + v)
}
```

Output:

```s
a is apple
b is banana
```

### Strings

To iterate over just the values in a string, use one temporary identifier:

```s
for v in "mojo" {
    print(v)
}
```

Output:

```s
m
o
j
o
```

To iterate over both the keys and the values in a string, use two temporary identifiers:

```s
for i, v in "mojo" {
    print(i, "->", v)
}
```

Output:

```s
0 -> m
1 -> o
2 -> j
3 -> o
```

### Lists

To iterate over just the values in a list, use one temporary identifier:

```s
names = ["alice", "bob", "charlie"]

for v in names {
    print(v)
}
```

Output:

```s
alice
bob
charlie
```

To iterate over both the keys and the values in a list, use two temporary identifiers:

```s
for i, v in names {
    print(i, "-", v)
}
```

Output:

```s
0 - alice
1 - bob
2 - charlie
```

## Break and Continue

### Break

Use the break keyword to terminate a loop:

```s
for i, v in "hello" {
    if (i == 2) {
        print("breaking loop")
        break
    }
    print(v)
}
```

Output:

```s
h
e
breaking loop
```

### Continue

Use the continue keyword to skip a specific iteration:

```s
for i, v in "hello" {
    if (i == 2) {
        print("skipping iteration")
        continue
    }
    print(v)
}
```

Output:

```s
h
e
skipping iteration
l
o
```

```

## function.md

```markdown
# Functions in Vint

Functions in **Vint** allow you to encapsulate code and execute it when needed. Here's a simple guide to understanding how functions work in **Vint**.

## Immediately Invoked Function

You can define and immediately execute a function:

```js
let go = func() {
    print("this is a function")
}()
```

This function `go` is defined and executed immediately upon declaration.

## Declared but Not Immediately Invoked Function

Functions can also be declared without being executed immediately:

```js
let vint = func() {
    print("This is also a function\nBut not invoked immediately after being declared")
}

vint()  // Executes the function
```

The function `vint` is called later using `vint()`.

## Passing Functions as Arguments

Functions in **Vint** can be passed as arguments to other functions:

```js
let w = func() {
    print("w function")
}

func(w) {
    w()  // Executes the function passed as an argument
    print("func")
}(w)  // Passes `w` as an argument and immediately invokes the outer function
```

In this example, the function `w` is passed to another function and executed within it.

By understanding these basic concepts, you can start creating reusable and flexible code using functions in **Vint**.

```

## hash.md

```markdown
# Hash Module in Vint

The Hash module in Vint provides additional hashing algorithms that complement the existing crypto module. It currently supports SHA1 and SHA512 hashing functions for generating secure hash values from text input.

---

## Importing the Hash Module

To use the Hash module, simply import it:
```js
import hash
```

---

## Functions and Examples

### 1. SHA1 Hash (`sha1`)

The `sha1` function generates a SHA1 hash from the provided string data.

**Syntax**:

```js
sha1(data)
```

**Example**:

```js
import hash

print("=== SHA1 Hash Example ===")
data = "hello world"
sha1_hash = hash.sha1(data)
print("Input:", data)
print("SHA1 Hash:", sha1_hash)
// Output: SHA1 Hash: 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
```

---

### 2. SHA512 Hash (`sha512`)

The `sha512` function generates a SHA512 hash from the provided string data.

**Syntax**:

```js
sha512(data)
```

**Example**:

```js
import hash

print("=== SHA512 Hash Example ===")
data = "hello world"
sha512_hash = hash.sha512(data)
print("Input:", data)
print("SHA512 Hash:", sha512_hash)
// Output: SHA512 Hash: 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
```

---

## Complete Usage Example

```js
import hash

print("=== Hash Module Complete Example ===")

// Test different types of data
test_data = [
    "hello",
    "hello world",
    "VintLang Programming Language",
    "The quick brown fox jumps over the lazy dog"
]

for data in test_data {
    print("\nInput:", data)
    print("SHA1:  ", hash.sha1(data))
    print("SHA512:", hash.sha512(data))
}

// Example with password hashing
password = "mySecretPassword123"
print("\n=== Password Hashing ===")
print("Password SHA1:  ", hash.sha1(password))
print("Password SHA512:", hash.sha512(password))
```

---

## Use Cases

- **Password Storage**: Hash passwords before storing them in databases
- **Data Integrity**: Verify file integrity by comparing hash values
- **Digital Signatures**: Generate unique identifiers for data
- **Checksums**: Create checksums for data validation
- **Security**: Generate secure hash values for authentication

---

## Summary of Functions

| Function | Description                               | Output Length    |
|----------|-------------------------------------------|------------------|
| `sha1`   | Generates SHA1 hash from string data      | 40 characters    |
| `sha512` | Generates SHA512 hash from string data    | 128 characters   |

Both functions return hexadecimal string representations of the hash values, making them easy to store and compare.

```

## http_enhanced.md

```markdown
# Enhanced HTTP Module - Full Backend Development Support

The Vint HTTP module has been significantly enhanced to support full-fledged backend development with enterprise-grade features.

## 🚀 New Features

### 1. **Enhanced Request Object**
- **JSON Body Parsing**: Automatic parsing of `application/json` content
- **Form Data Parsing**: Support for `application/x-www-form-urlencoded` data
- **Cookie Handling**: Easy access to request cookies
- **Query Parameters**: Enhanced query parameter utilities
- **Path Parameters**: Extract parameters from routes like `/users/:id`
- **Headers**: Improved header access methods

```js
http.get("/users/:id", func(req, res) {
    let userId = req.param("id")          // Path parameter
    let sort = req.query("sort")          // Query parameter
    let session = req.cookie("session")   // Cookie value
    let auth = req.get("Authorization")   // Header value
    let body = req.body()                 // Raw body
    let json = req.json()                 // Parsed JSON
    let form = req.form("username")       // Form field
})
```

### 2. **Enhanced Response Object**

- **Method Chaining**: Chain response methods for cleaner code
- **Redirect Support**: Easy redirects with custom status codes
- **Cookie Setting**: Set response cookies with options
- **JSON Responses**: Improved JSON response handling
- **Status Helpers**: Easy status code management

```js
http.post("/login", func(req, res) {
    res.status(200)
       .cookie("session", "abc123")
       .header("X-Custom", "value")
       .json({"success": true})
    
    // Or redirect
    res.redirect("/dashboard", 302)
})
```

### 3. **Interceptors**

Request and response interceptors for cross-cutting concerns:

```js
// Request interceptor - runs before route handlers
http.interceptor("request", func(req) {
    print("Processing request:", req.path())
    // Add request timestamp, validate format, etc.
})

// Response interceptor - runs after route handlers
http.interceptor("response", func(res) {
    print("Processing response")
    // Add security headers, log response time, etc.
})
```

### 4. **Guards**

Security guards for authentication, authorization, and rate limiting:

```js
// Authentication guard
http.guard(func(req) {
    print("Checking authentication")
    // Verify JWT token, session, etc.
})

// Rate limiting guard
http.guard(func(req) {
    print("Checking rate limits")
    // Track requests per IP, enforce limits
})

// Custom validation guard
http.guard(func(req) {
    print("Custom security checks")
    // SQL injection detection, XSS prevention, etc.
})
```

### 5. **Enhanced Middleware**

Built-in middleware for common backend needs:

```js
// CORS support
http.cors()

// Body parsing
http.bodyParser()

// Authentication middleware
http.auth(func(req, res, next) {
    print("Processing authentication")
})

// Global error handler
http.errorHandler(func(err, req, res) {
    print("Handling error:", err)
})
```

### 6. **Route Parameters**

Support for parameterized routes:

```js
// Single parameter
http.get("/users/:id", func(req, res) {
    let id = req.param("id")
})

// Multiple parameters
http.get("/users/:userId/posts/:postId", func(req, res) {
    let userId = req.param("userId")
    let postId = req.param("postId")
})
```

### 7. **Security Features**

- **Automatic CORS**: CORS headers added automatically
- **OPTIONS Handling**: Preflight requests handled automatically
- **Security Headers**: Enhanced security header management
- **Error Sanitization**: Safe error responses

## 📖 Complete Example

```js
import http

// Create application
http.app()

// Add interceptors
http.interceptor("request", func(req) {
    print("Request interceptor")
})

http.interceptor("response", func(res) {
    print("Response interceptor")
})

// Add guards
http.guard(func(req) {
    print("Auth guard")
})

http.guard(func(req) {
    print("Rate limit guard")
})

// Add middleware
http.cors()
http.bodyParser()
http.auth(func(req, res, next) {
    print("Auth middleware")
})

// Set error handler
http.errorHandler(func(err, req, res) {
    print("Error handler")
})

// Define routes
http.get("/", func(req, res) {
    res.send("Welcome to enhanced HTTP server!")
})

http.get("/users/:id", func(req, res) {
    let id = req.param("id")
    res.json({"user": id})
})

http.post("/api/data", func(req, res) {
    let data = req.json()
    res.status(201).json({"created": true})
})

// Start server
http.listen(3000, "Enhanced server running on port 3000!")
```

## 🧪 Testing

Comprehensive test files are included:

- `examples/enhanced_http_test.vint` - Feature demonstrations
- `examples/backend_demo.vint` - Backend application demo
- `examples/complete_backend_app.vint` - Full production-ready example
- `examples/live_server_test.vint` - Live server testing
- `module/http_enhanced_test.go` - Go unit tests
- `object/http_enhanced_test.go` - Object method tests

## 🏗️ Architecture

The enhanced HTTP module follows a layered architecture:

1. **Interceptors** → Process all requests/responses
2. **Guards** → Security and validation checks
3. **Middleware** → Cross-cutting concerns (CORS, auth, etc.)
4. **Route Handlers** → Business logic
5. **Error Handlers** → Error processing

## 🎯 Backend Capabilities

The enhanced HTTP module now supports:

✅ **User Management** - Complete CRUD operations  
✅ **Authentication** - JWT, sessions, cookies  
✅ **Authorization** - Role-based access control  
✅ **File Uploads** - Multi-part form data  
✅ **API Development** - RESTful endpoints  
✅ **Security** - Guards, validation, sanitization  
✅ **Analytics** - Request logging, metrics  
✅ **Admin Panels** - Administrative interfaces  
✅ **Health Monitoring** - System health checks  
✅ **Error Handling** - Comprehensive error management  
✅ **Rate Limiting** - Request throttling  
✅ **CORS Support** - Cross-origin requests  
✅ **Content Types** - JSON, forms, files  

## 🚀 Production Ready

The enhanced HTTP module provides all the essential features needed for production backend applications:

- **Scalability**: Efficient request processing
- **Security**: Multi-layer protection
- **Monitoring**: Built-in health checks and metrics
- **Flexibility**: Extensible middleware and guard system
- **Standards**: RESTful API support with proper HTTP semantics

This makes Vint suitable for building everything from simple APIs to complex enterprise applications.

```

## http_enterprise.md

```markdown
# Enterprise HTTP Module Documentation

The Vint HTTP module has been enhanced with enterprise-level features to support building production-ready backend applications. This document covers the advanced features added to the original HTTP module.

## Table of Contents

1. [Overview](#overview)
2. [Route Grouping & API Versioning](#route-grouping--api-versioning)
3. [Multipart File Uploads](#multipart-file-uploads)
4. [Async Handlers](#async-handlers)
5. [Enhanced Security](#enhanced-security)
6. [Advanced Middleware](#advanced-middleware)
7. [Structured Error Handling](#structured-error-handling)
8. [Performance Monitoring](#performance-monitoring)
9. [Complete Examples](#complete-examples)

## Overview

The enterprise HTTP module extends the basic HTTP functionality with:

- **Route Grouping**: Organize routes with common prefixes and middleware
- **File Uploads**: Complete multipart/form-data support with file handling
- **Async Processing**: Non-blocking handlers for long-running operations
- **Security Features**: CSRF protection, security headers, enhanced CORS
- **Middleware Composition**: Advanced middleware stacking and composition
- **Error Handling**: Structured error responses with consistent format
- **Performance Hooks**: Request timing and metrics for APM integration

## Route Grouping & API Versioning

### Creating Route Groups

Route groups allow you to organize related routes under a common prefix:

```js
import http

http.app()

// Create API v1 group
http.group("/api/v1", func() {
    // All routes in this group will be prefixed with /api/v1
})

// Create API v2 group  
http.group("/api/v2", func() {
    // All routes in this group will be prefixed with /api/v2
})
```

### Nested Route Groups

```js
// Admin routes group
http.group("/admin", func() {
    // User management routes
    http.group("/users", func() {
        // /admin/users/* routes
    })
    
    // System routes
    http.group("/system", func() {
        // /admin/system/* routes
    })
})
```

## Multipart File Uploads

### Basic File Upload

```js
http.post("/upload", func(req, res) {
    // Parse multipart form data
    http.multipart(req)
    
    // Access uploaded file
    let avatar = req.file("avatar")
    
    if avatar {
        // Get file information
        let name = avatar.name()
        let size = avatar.size()
        let type = avatar.type()
        
        // Save the file
        let saved = avatar.save("/uploads/" + name)
        
        res.json({
            "success": true,
            "file": {
                "name": name,
                "size": size,
                "type": type,
                "saved": saved
            }
        })
    } else {
        res.status(400).json({
            "error": "No file uploaded"
        })
    }
})
```

### Multiple File Upload

```js
http.post("/multiple-upload", func(req, res) {
    http.multipart(req)
    
    let uploadedFiles = []
    let files = req.files()
    
    // Process each file
    for file in files {
        let savedPath = file.save("/uploads/" + file.name())
        uploadedFiles.push({
            "name": file.name(),
            "size": file.size(),
            "type": file.type(),
            "path": savedPath
        })
    }
    
    res.json({
        "success": true,
        "count": uploadedFiles.length,
        "files": uploadedFiles
    })
})
```

### Form Data with File Upload

```js
http.post("/profile", func(req, res) {
    http.multipart(req)
    
    // Access form fields
    let username = req.form("username")
    let email = req.form("email")
    
    // Access uploaded file
    let avatar = req.file("avatar")
    
    if avatar {
        avatar.save("/avatars/" + username + "_" + avatar.name())
    }
    
    res.json({
        "message": "Profile updated",
        "user": {
            "username": username,
            "email": email,
            "avatar": avatar ? avatar.name() : null
        }
    })
})
```

## Async Handlers

Async handlers allow long-running operations without blocking other requests:

```js
// Create async handler for heavy processing
let processDataAsync = http.async(func(req, res) {
    // This runs asynchronously and won't block other requests
    let data = req.json()
    
    // Simulate heavy processing
    // In real applications: database operations, API calls, image processing
    processLargeDataset(data)
    
    res.json({
        "message": "Processing started",
        "taskId": generateTaskId()
    })
})

http.post("/process", processDataAsync)

// Immediate response handler
http.post("/quick", func(req, res) {
    res.json({"message": "Quick response"})
})
```

## Enhanced Security

### Security Middleware

```js
// Enable security features
http.security()

// This automatically adds:
// - X-Content-Type-Options: nosniff
// - X-Frame-Options: DENY  
// - X-XSS-Protection: 1; mode=block
// - CSRF protection (when enabled)
```

### Custom Security Headers

```js
http.use(func(req, res, next) {
    res.header("Strict-Transport-Security", "max-age=31536000")
    res.header("Content-Security-Policy", "default-src 'self'")
    next()
})
```

### CORS Configuration

```js
// Basic CORS
http.cors()

// Custom CORS (configuration would be enhanced in future)
http.use(func(req, res, next) {
    res.header("Access-Control-Allow-Origin", "https://myapp.com")
    res.header("Access-Control-Allow-Credentials", "true")
    next()
})
```

## Advanced Middleware

### Middleware Composition

```js
// Authentication middleware
let authMiddleware = func(req, res, next) {
    let token = req.get("Authorization")
    if !token {
        res.status(401).json({"error": "Unauthorized"})
        return
    }
    // Validate token
    next()
}

// Logging middleware
let logMiddleware = func(req, res, next) {
    print("Request: " + req.method() + " " + req.path())
    next()
}

// Rate limiting middleware
let rateLimitMiddleware = func(req, res, next) {
    // Check rate limits
    next()
}

// Apply middleware in order
http.use(logMiddleware)
http.use(rateLimitMiddleware)
http.use(authMiddleware)
```

### Route-Specific Middleware

```js
// Apply multiple middlewares to specific routes
http.post("/protected", [authMiddleware, rateLimitMiddleware], func(req, res) {
    res.json({"message": "Protected resource accessed"})
})
```

## Structured Error Handling

### Global Error Handler

```js
http.errorHandler(func(err, req, res) {
    res.status(500).json({
        "error": {
            "type": "INTERNAL_SERVER_ERROR",
            "message": err.message,
            "code": "ERR_INTERNAL",
            "status": 500,
            "details": {
                "timestamp": Date.now(),
                "path": req.path(),
                "method": req.method(),
                "requestId": generateRequestId()
            }
        }
    })
})
```

### Custom Error Responses

```js
http.get("/users/:id", func(req, res) {
    let userId = req.param("id")
    
    if !isValidId(userId) {
        res.status(400).json({
            "error": {
                "type": "VALIDATION_ERROR",
                "message": "Invalid user ID format",
                "code": "INVALID_USER_ID",
                "status": 400,
                "details": {
                    "field": "id",
                    "value": userId,
                    "expected": "numeric ID"
                }
            }
        })
        return
    }
    
    let user = findUser(userId)
    if !user {
        res.status(404).json({
            "error": {
                "type": "NOT_FOUND",
                "message": "User not found",
                "code": "USER_NOT_FOUND",
                "status": 404,
                "details": {
                    "userId": userId
                }
            }
        })
        return
    }
    
    res.json(user)
})
```

## Performance Monitoring

### Request Timing

```js
// Middleware to track request timing
let timingMiddleware = func(req, res, next) {
    let startTime = Date.now()
    
    // Add custom response method to track timing
    let originalSend = res.send
    res.send = func(data) {
        let duration = Date.now() - startTime
        res.header("X-Response-Time", duration + "ms")
        originalSend(data)
    }
    
    next()
}

http.use(timingMiddleware)
```

### Metrics Endpoint

```js
http.get("/metrics", func(req, res) {
    res.header("Content-Type", "text/plain")
    res.send(`
# HTTP Request Count
http_requests_total{method="GET"} 1234
http_requests_total{method="POST"} 567

# HTTP Request Duration
http_request_duration_seconds{quantile="0.5"} 0.05
http_request_duration_seconds{quantile="0.95"} 0.2

# System Metrics
memory_usage_bytes 104857600
cpu_usage_percent 15.5
`)
})
```

## Complete Examples

### Production-Ready API Server

```js
import http

// Create application
http.app()

// Security setup
http.security()

// Global middleware
http.use(func(req, res, next) {
    print("Request: " + req.method() + " " + req.path())
    res.header("X-Powered-By", "VintLang")
    next()
})

// Authentication middleware
let authMiddleware = func(req, res, next) {
    let token = req.get("Authorization")
    if token && validateJWT(token) {
        next()
    } else {
        res.status(401).json({
            "error": {
                "type": "UNAUTHORIZED",
                "message": "Invalid or missing authentication token",
                "code": "AUTH_REQUIRED"
            }
        })
    }
}

// API v1 routes
http.group("/api/v1", func() {
    // Public routes
    http.post("/auth/login", func(req, res) {
        let credentials = req.json()
        let token = authenticateUser(credentials)
        
        if token {
            res.json({
                "token": token,
                "expires": Date.now() + 3600000
            })
        } else {
            res.status(401).json({
                "error": {
                    "type": "AUTHENTICATION_FAILED",
                    "message": "Invalid credentials"
                }
            })
        }
    })
    
    // Protected routes
    http.use(authMiddleware)
    
    http.get("/users", func(req, res) {
        let page = req.query("page") || "1"
        let users = getUsers(page)
        res.json(users)
    })
    
    http.post("/upload", func(req, res) {
        http.multipart(req)
        let file = req.file("document")
        
        if file {
            let savedPath = file.save("/secure-uploads/" + generateFileName())
            res.json({
                "message": "File uploaded successfully",
                "fileId": generateFileId(savedPath)
            })
        } else {
            res.status(400).json({
                "error": {
                    "type": "VALIDATION_ERROR",
                    "message": "No file provided"
                }
            })
        }
    })
})

// Health check
http.get("/health", func(req, res) {
    res.json({
        "status": "healthy",
        "timestamp": Date.now(),
        "version": "1.0.0"
    })
})

// Error handler
http.errorHandler(func(err, req, res) {
    print("Error: " + err.message)
    res.status(500).json({
        "error": {
            "type": "INTERNAL_SERVER_ERROR",
            "message": "An unexpected error occurred",
            "code": "ERR_INTERNAL"
        }
    })
})

// Start server
http.listen(3000, "Production API server running on port 3000")
```

### File Upload Service

```js
import http

http.app()
http.security()

// File upload with validation
http.post("/files", func(req, res) {
    http.multipart(req)
    
    let files = req.files()
    let results = []
    
    for file in files {
        // Validate file type
        let allowedTypes = ["image/jpeg", "image/png", "application/pdf"]
        if !allowedTypes.includes(file.type()) {
            results.push({
                "name": file.name(),
                "error": "File type not allowed"
            })
            continue
        }
        
        // Validate file size (10MB max)
        if file.size() > 10485760 {
            results.push({
                "name": file.name(),
                "error": "File too large (max 10MB)"
            })
            continue
        }
        
        // Save file
        let fileName = Date.now() + "_" + file.name()
        let savedPath = file.save("/uploads/" + fileName)
        
        results.push({
            "name": file.name(),
            "fileName": fileName,
            "size": file.size(),
            "type": file.type(),
            "url": "/files/" + fileName,
            "status": "uploaded"
        })
    }
    
    res.json({
        "message": "File upload processed",
        "results": results
    })
})

// Serve uploaded files
http.get("/files/:filename", func(req, res) {
    let filename = req.param("filename")
    // In a real implementation, serve the file from storage
    res.send("File: " + filename)
})

http.listen(3000, "File upload service running on port 3000")
```

## Summary

The enterprise HTTP module provides all the features needed to build production-ready backend applications:

- **Scalable**: Route grouping and middleware composition for large applications
- **Secure**: Built-in security features and customizable protection
- **Performance**: Async handlers and monitoring capabilities  
- **Robust**: Structured error handling and comprehensive file upload support
- **Production-Ready**: All features needed for enterprise backend development

This makes VintLang suitable for building everything from simple APIs to complex enterprise applications with the same level of sophistication as modern frameworks like Express.js, FastAPI, or Spring Boot.

```

## identifiers.md

```markdown
# Identifiers in Vint

Identifiers are used to name variables, functions, and other elements in your **Vint** code. This guide explains the rules and best practices for creating effective identifiers.

## Syntax Rules

Identifiers can include letters, numbers, and underscores. However, they must adhere to these rules:

- **Cannot start with a number.**
- **Case-sensitive:** For example, `myVar` and `myvar` are considered different identifiers.

### Examples of Valid Identifiers:

```js
let birth_year = 2020
print(birth_year)  // Output: 2020

let convert_c_to_p = "C to P"
print(convert_c_to_p)  // Output: "C to P"
```

In the examples above, `birth_year` and `convert_c_to_p` follow all syntax rules and are valid identifiers.

## Best Practices

To make your **Vint** code more readable and maintainable, follow these best practices:

1. **Use Descriptive Names:** Choose names that clearly describe the purpose or content of the variable or function.

   ```js
   let total_score = 85
   let calculate_average = func() { /* logic */ }
   ```

2. **Consistent Naming Conventions:** Stick to a single naming style across your codebase:
   - **camelCase**: `myVariableName`
   - **snake_case**: `my_variable_name`

3. **Avoid Ambiguity:** Use meaningful names instead of single letters, except for common cases like loop counters:

   ```js
   for (let i = 0; i < 10; i++) {
       print(i)
   }
   ```

4. **Do Not Use Reserved Keywords:** Avoid using reserved keywords as identifiers (e.g., `let`, `if`, `switch`).

```

## ifStatements.md

```markdown
# Conditional Statements in Vint

Conditional statements in **Vint** allow you to perform different actions based on specific conditions. The `if/else` structure is fundamental for controlling the flow of your code. Here's a simple guide to using conditional statements in **Vint**.

## If Statement (`if`)

The `if` statement checks a condition inside parentheses `()`. If the condition evaluates to true, the code block inside curly braces `{}` will execute:

```js
if (2 > 1) {
    print(true)  // Output: true
}
```

In this example, the condition `2 > 1` is true, so `print(true)` is executed, and the output is `true`.

## Else If and Else Blocks (`else if` and `else`)

You can use `else if` to test additional conditions after an `if` statement. The `else` block specifies code to execute if none of the previous conditions are met:

```js
let a = 10

if (a > 100) {
    print("a is greater than 100")
} else if (a < 10) {
    print("a is less than 10")
} else {
    print("The value of a is", a)
}

// Output: The value of a is 10
```

### Explanation

1. The condition `a > 100` is false.
2. The next condition `a < 10` is also false.
3. Therefore, the `else` block is executed, and the output is `The value of a is 10`.

## Summary

- **`if`**: Executes code if the condition is true.
- **`else if`**: Tests another condition if the previous `if` condition is false.
- **`else`**: Executes code if none of the above conditions are true.

By using `if`, `else if`, and `else`, you can make decisions and control the flow of your **Vint** programs based on dynamic conditions.

# If Statements and If Expressions in Vint

Vint supports both classic if statements and the new if expressions, allowing you to use conditional logic in both statement and expression positions.

---

## Classic If Statement

The classic if statement executes a block of code if a condition is true. You can optionally provide an `else` block.

**Syntax:**

```js
if (condition) {
    // code to run if condition is true
} else {
    // code to run if condition is false
}
```

**Example:**

```js
let x = 0
if (true) {
    x = 42
}
print("Classic if statement result: ", x)
```

---

## If as an Expression (New Feature)

You can now use `if` as an expression, which returns a value. This allows you to assign the result of a conditional directly to a variable, or use it in any expression context.

**Syntax:**

```js
let result = if (condition) { valueIfTrue } else { valueIfFalse }
```

- The `if` expression evaluates to the value of the first block if the condition is true, or the value of the `else` block if provided.
- If the condition is false and there is no `else`, the result is `null`.

**Examples:**

```js
let status = ""
status = if (x > 0) { "Online" } else { "Offline" }
print("If as an expression result: ", status)

let y = if (false) { 123 }
print("If as an expression with no else: ", y) // prints: null
```

---

## Notes

- Parentheses around the condition are required: `if (condition) { ... }`.
- Both the classic statement and the new expression form are fully supported and can be mixed in your code.
- Use `//` for single-line comments and `/* ... */` for multi-line comments in Vint.

---

## See Also

- [Switch Statements](switch.md)
- [Operators](operators.md)

```

## include.md

```markdown
# Include

The `include` keyword is a language construct in Vint that allows you to include and evaluate code from another file into the current file. This is useful for organizing code into reusable modules, separating concerns, and managing larger projects more effectively. When a file is included, its code is executed in the same scope as the `include` statement, meaning any variables, functions, or other constructs defined in the included file become available in the including file.

## Syntax

```js
include "path/to/your/file.vint"
```

The path to the file can be relative or absolute. The file extension is not mandatory but is recommended for clarity.

## Example

Let's say you have a file named `greetings.vint` with the following content:

**greetings.vint**

```js
let greeting = "Hello, Vint!"

func sayHello() {
    println(greeting)
}
```

You can include this file in another file, for instance, `main.vint`, and use the `greeting` variable and the `sayHello` function:

**main.vint**

```js
include "greetings.vint"

sayHello() // Output: Hello, Vint!

let customMessage = greeting + " How are you?"
println(customMessage) // Output: Hello, Vint! How are you?
```

In this example, the `include` statement at the beginning of `main.vint` makes the `greeting` variable and the `sayHello` function from `greetings.vint` available for use. This helps in keeping the code modular and easy to manage.

```

## info.md

```markdown
# Info

The `info` keyword allows you to print informational messages at runtime.

## Syntax

```js
info "Your informational message here"
```

When the Vint interpreter encounters an `info` statement, it prints a cyan-colored informational message to the console and continues execution. This is useful for providing helpful context or status updates to users or developers.

### Example

```js
info "Starting the backup process."
println("Backup in progress...")
```

Running this script will output:

```
[INFO]: Starting the backup process.
Backup in progress...
```

```

## json.md

```markdown
# JSON Module in Vint

The JSON module in Vint provides powerful and straightforward functions for working with JSON data, including decoding, encoding, formatting, merging, and retrieving values. Below is the detailed documentation, along with examples.

---

## Importing the JSON Module

To use the JSON module, simply import it:
```js
import json
```

---

## Functions and Examples

### 1. Decode JSON (`decode`)

The `decode` function parses a JSON string into a Vint dictionary or array.

**Syntax**:

```js
decode(jsonString)
```

**Example**:

```js
import json

print("=== Example 1: Decode ===")
raw_json = '{"name": "John", "age": 30, "isAdmin": false, "friends": ["Jane", "Doe"]}'
decoded = json.decode(raw_json)
print("Decoded Object:", decoded)
// Output: Decoded Object: {"name": "John", "age": 30, "isAdmin": false, "friends": ["Jane", "Doe"]}
```

---

### 2. Encode JSON (`encode`)

The `encode` function converts a Vint dictionary or array into a JSON string. It optionally supports pretty formatting with an `indent` parameter.

**Syntax**:

```js
encode(data, indent = 0)
```

**Example**:

```js
import json

print("\n=== Example 2: Encode ===")
data = {
  "language": "Vint",
  "version": 1.0,
  "features": ["custom modules", "native objects"]
}
encoded_json = json.encode(data, indent=2)
print("Encoded JSON:", encoded_json)
// Output:
// Encoded JSON: {
//   "language": "Vint",
//   "version": 1.0,
//   "features": ["custom modules", "native objects"]
// }
```

---

### 3. Pretty Print JSON (`pretty`)

The `pretty` function reformats a JSON string into a human-readable format with proper indentation.

**Syntax**:

```js
pretty(jsonString)
```

**Example**:

```js
import json

print("\n=== Example 3: Pretty Print ===")
raw_json_pretty = '{"name":"John","age":30,"friends":["Jane","Doe"]}'
pretty_json = json.pretty(raw_json_pretty)
print("Pretty JSON:\n", pretty_json)
// Output:
// Pretty JSON:
// {
//   "name": "John",
//   "age": 30,
//   "friends": ["Jane", "Doe"]
// }
```

---

### 4. Merge JSON Objects (`merge`)

The `merge` function combines two JSON objects. If both objects have the same key, the value from the second object overwrites the first.

**Syntax**:

```js
merge(json1, json2)
```

**Example**:

```js
import json

print("\n=== Example 4: Merge ===")
json1 = {"name": "John", "age": 30}
json2 = {"city": "New York", "age": 35}
merged_json = json.merge(json1, json2)
print("Merged JSON:", merged_json)
// Output: Merged JSON: {"name": "John", "age": 35, "city": "New York"}
```

---

### 5. Get Value by Key (`get`)

The `get` function retrieves a value associated with a key from a JSON object. If the key is not found, it returns `null`.

**Syntax**:

```js
get(jsonObject, key)
```

**Example**:

```js
import json

print("\n=== Example 5: Get Value by Key ===")
json_object = {"name": "John", "age": 30, "city": "New York"}

value = json.get(json_object, "age")
print("Age:", value)
// Output: Age: 30

missing_value = json.get(json_object, "country")
print("Country (missing key):", missing_value)
// Output: Country (missing key): null
```

---

## Summary of Functions

| Function         | Description                                         | Example Output                           |
|------------------|-----------------------------------------------------|------------------------------------------|
| `decode`         | Converts JSON string to a Vint object.             | `{"key": "value"}`                       |
| `encode`         | Converts Vint object to a JSON string.             | `{"key":"value"}`                        |
| `pretty`         | Formats JSON string for better readability.        | `{ "key": "value" }`                     |
| `merge`          | Combines two JSON objects, overwriting duplicates. | `{"key1": "value1", "key2": "value2"}`   |
| `get`            | Retrieves a value by key, returns `null` if absent.| `"value"` or `null`                      |

These functions make working with JSON in Vint easy, flexible, and efficient.

```

## jwt.md

```markdown
# JWT Module

The JWT module provides functions for creating, verifying, and decoding JSON Web Tokens (JWTs). It supports HMAC-SHA256 signing by default.

## Functions

### `jwt.create(payload, secret)`

Creates a JWT token with the provided payload and secret using HS256 signing method.

**Parameters:**

- `payload` (dict): The claims/payload to include in the JWT
- `secret` (string): The secret key used for signing

**Returns:**

- `string`: The JWT token string

**Example:**

```js
let payload = { user: "john", role: "admin" };
let token = jwt.create(payload, "my-secret-key");
print(token); // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### `jwt.createHS256(payload, secret, [expiration_hours])`

Creates a JWT token with HS256 signing method, with optional expiration.

**Parameters:**

- `payload` (dict): The claims/payload to include in the JWT
- `secret` (string): The secret key used for signing
- `expiration_hours` (number, optional): Token expiration time in hours

**Returns:**

- `string`: The JWT token string

**Example:**

```js
let payload = { user: "john", role: "admin" };
let token = jwt.createHS256(payload, "my-secret-key", 24); // Expires in 24 hours
print(token);
```

### `jwt.verify(token, secret)`

Verifies a JWT token and returns the payload if valid.

**Parameters:**

- `token` (string): The JWT token to verify
- `secret` (string): The secret key used for verification

**Returns:**

- `dict`: The token payload if valid
- `error`: Error if token is invalid or verification fails

**Example:**

```js
let token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...";
let result = jwt.verify(token, "my-secret-key");

if (result.type != "ERROR") {
  print("User:", result["user"]);
  print("Role:", result["role"]);
} else {
  print("Invalid token:", result.message);
}
```

### `jwt.verifyHS256(token, secret)`

Verifies a JWT token with explicit HS256 signing method check.

**Parameters:**

- `token` (string): The JWT token to verify
- `secret` (string): The secret key used for verification

**Returns:**

- `dict`: The token payload if valid
- `error`: Error if token is invalid or not HS256 signed

**Example:**

```js
let token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...";
let result = jwt.verifyHS256(token, "my-secret-key");

if (result.type != "ERROR") {
  print("Verified HS256 token:", result);
}
```

### `jwt.decode(token)`

Decodes a JWT token without verification (useful for inspecting headers/payload).

**Parameters:**

- `token` (string): The JWT token to decode

**Returns:**

- `dict`: Contains "header" and "payload" keys with their respective data
- `error`: Error if token format is invalid

**Example:**

```js
let token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...";
let decoded = jwt.decode(token);

print("Header:", decoded["header"]);
print("Payload:", decoded["payload"]);
```

## Error Handling

All JWT functions return error objects when something goes wrong:

```js
let result = jwt.verify("invalid-token", "secret");
if (result.type == "ERROR") {
  print("Error:", result.message);
}
```

## Security Considerations

1. **Keep secrets secure**: Never hardcode JWT secrets in your source code
2. **Use strong secrets**: Use cryptographically strong random strings
3. **Set expiration times**: Always include expiration claims for security
4. **Validate claims**: Always verify the token payload matches your expectations

## Common Use Cases

### User Authentication

```js
// Login endpoint - create token
let payload = {
  user_id: 123,
  username: "john_doe",
  exp: time.now() + 24 * 3600, // 24 hours
};
let token = jwt.create(payload, env.JWT_SECRET);

// Protected endpoint - verify token
let result = jwt.verify(request_token, env.JWT_SECRET);
if (result.type == "ERROR") {
  return { error: "Unauthorized" };
}
let user_id = result["user_id"];
```

### API Rate Limiting

```js
// Create token with rate limit info
let payload = {
  client_id: "app123",
  requests_remaining: 1000,
  exp: time.now() + 3600, // 1 hour
};
let token = jwt.create(payload, "rate-limit-secret");
```

### Session Management

```js
// Create session token
let session = {
  session_id: uuid.generate(),
  user_id: user.id,
  permissions: ["read", "write"],
  exp: time.now() + 8 * 3600, // 8 hours
};
let session_token = jwt.createHS256(session, "session-secret", 8);
```

```

## kv.md

```markdown
# KV Module

The **KV (Key-Value)** module provides a high-performance, thread-safe, in-memory key-value store for VintLang applications. It offers enterprise-grade features including TTL (Time-To-Live), atomic operations, bulk operations, and comprehensive statistics.

## Import

```js
import kv
```

## Features

- 🔒 **Thread-Safe**: Concurrent read/write operations with proper locking
- ⏰ **TTL Support**: Automatic key expiration with customizable time-to-live
- 🚀 **Atomic Operations**: Increment/decrement operations for counters
- 📦 **Bulk Operations**: Efficient multi-key get/set operations
- 📊 **Statistics**: Real-time metrics and store monitoring
- 💾 **Data Dump**: Export all data for debugging and backup
- 🧹 **Memory Management**: Automatic cleanup of expired keys

## Basic Operations

### set(key, value)

Sets a key-value pair in the store.

**Parameters:**

- `key` (string): The key to set
- `value` (any): The value to store

**Returns:** `boolean` - `true` if successful

**Example:**

```js
import kv

kv.set("user:123", {"name": "John", "age": 30})
kv.set("session:456", "active")
kv.set("counter", 42)
```

### get(key)

Retrieves a value by its key.

**Parameters:**

- `key` (string): The key to retrieve

**Returns:** The stored value, or `null` if not found or expired

**Example:**

```js
let user = kv.get("user:123");
println("User:", user); // {name: John, age: 30}

let missing = kv.get("nonexistent");
println("Missing:", missing); // null
```

### delete(key)

Removes a key-value pair from the store.

**Parameters:**

- `key` (string): The key to delete

**Returns:** `boolean` - `true` if key existed and was deleted

**Example:**

```js
let deleted = kv.delete("session:456");
println("Deleted:", deleted); // true
```

### exists(key)

Checks if a key exists in the store (and is not expired).

**Parameters:**

- `key` (string): The key to check

**Returns:** `boolean` - `true` if key exists and is not expired

**Example:**

```js
if (kv.exists("user:123")) {
  println("User exists");
}
```

### clear()

Removes all key-value pairs from the store.

**Returns:** `boolean` - Always `true`

**Example:**

```js
kv.clear();
println("Store cleared");
```

## Store Information

### keys()

Returns an array of all keys in the store (excluding expired keys).

**Returns:** `array` - Array of string keys

**Example:**

```js
kv.set("key1", "value1");
kv.set("key2", "value2");
let allKeys = kv.keys();
println("Keys:", allKeys); // ["key1", "key2"]
```

### values()

Returns an array of all values in the store (excluding expired values).

**Returns:** `array` - Array of stored values

**Example:**

```js
let allValues = kv.values();
println("Values:", allValues); // ["value1", "value2"]
```

### size()

Returns the number of key-value pairs in the store.

**Returns:** `integer` - Number of stored pairs

**Example:**

```js
println("Store size:", kv.size()); // 2
```

### isEmpty()

Checks if the store is empty.

**Returns:** `boolean` - `true` if store has no keys

**Example:**

```js
if (kv.isEmpty()) {
  println("Store is empty");
}
```

## TTL (Time-To-Live) Operations

### setTTL(key, value, ttl_seconds)

Sets a key-value pair with automatic expiration.

**Parameters:**

- `key` (string): The key to set
- `value` (any): The value to store
- `ttl_seconds` (integer): Time-to-live in seconds

**Returns:** `boolean` - `true` if successful

**Example:**

```js
// Set a session that expires in 5 minutes
kv.setTTL("session:temp", "temporary_data", 300);

// Set a cache entry that expires in 1 hour
kv.setTTL("cache:user:123", userData, 3600);
```

### getTTL(key)

Gets the remaining time-to-live for a key.

**Parameters:**

- `key` (string): The key to check

**Returns:** `integer` - Remaining seconds, or `-1` if no TTL set, or `null` if key doesn't exist

**Example:**

```js
let remaining = kv.getTTL("session:temp");
if (remaining != null && remaining > 0) {
  println("Session expires in", remaining, "seconds");
}
```

### expire(key, ttl_seconds)

Sets or updates the TTL for an existing key.

**Parameters:**

- `key` (string): The key to set expiration for
- `ttl_seconds` (integer): Time-to-live in seconds (must be positive)

**Returns:** `boolean` - `true` if key exists and TTL was set

**Example:**

```js
kv.set("temp:data", "some data");
kv.expire("temp:data", 60); // Expire in 1 minute
```

## Bulk Operations

### mget(keys)

Gets multiple values in a single operation.

**Parameters:**

- `keys` (array): Array of string keys to retrieve

**Returns:** `array` - Array of values in same order as keys (`null` for missing/expired keys)

**Example:**

```js
kv.set("user:1", "Alice");
kv.set("user:2", "Bob");

let users = kv.mget(["user:1", "user:2", "user:3"]);
println("Users:", users); // ["Alice", "Bob", null]
```

### mset(pairs)

Sets multiple key-value pairs in a single operation.

**Parameters:**

- `pairs` (dictionary): Dictionary of key-value pairs to set

**Returns:** `boolean` - `true` if all pairs were set successfully

**Example:**

```js
let bulkData = {
  "config:theme": "dark",
  "config:language": "en",
  "config:notifications": true,
};
kv.mset(bulkData);
```

## Atomic Operations

### increment(key, [delta])

Atomically increments a numeric value.

**Parameters:**

- `key` (string): The key to increment
- `delta` (integer, optional): Amount to increment by (default: 1)

**Returns:** `integer` - The new value after increment

**Notes:**

- If key doesn't exist, creates it with the delta value
- If key exists but value is not an integer, returns an error
- Thread-safe for concurrent increments

**Example:**

```js
// Simple counter
let count = kv.increment("page:views");
println("Views:", count); // 1

// Increment by custom amount
let score = kv.increment("user:score", 10);
println("Score:", score); // 10

// Increment existing value
kv.set("counter", 5);
let newCount = kv.increment("counter", 3);
println("New count:", newCount); // 8
```

### decrement(key, [delta])

Atomically decrements a numeric value.

**Parameters:**

- `key` (string): The key to decrement
- `delta` (integer, optional): Amount to decrement by (default: 1)

**Returns:** `integer` - The new value after decrement

**Notes:**

- If key doesn't exist, creates it with the negative delta value
- If key exists but value is not an integer, returns an error
- Thread-safe for concurrent decrements

**Example:**

```js
// Simple countdown
let remaining = kv.decrement("lives");
println("Lives remaining:", remaining); // -1

// Decrement by custom amount
kv.set("inventory", 100);
let newInventory = kv.decrement("inventory", 15);
println("Inventory:", newInventory); // 85
```

## Utility Functions

### dump()

Returns all key-value pairs in the store as a dictionary (excluding expired keys).

**Returns:** `dictionary` - All stored key-value pairs

**Example:**

```js
kv.set("key1", "value1");
kv.set("key2", 42);
let allData = kv.dump();
println("All data:", allData); // {key1: value1, key2: 42}
```

### stats()

Returns statistics about the KV store.

**Returns:** `dictionary` - Statistics including:

- `total_keys`: Total number of keys (including expired)
- `active_keys`: Number of active (non-expired) keys
- `expired_keys`: Number of expired keys
- `keys_with_ttl`: Number of keys that have TTL set

**Example:**

```js
let statistics = kv.stats();
println("Store stats:", statistics);
// {total_keys: 10, active_keys: 8, expired_keys: 2, keys_with_ttl: 5}
```

## Common Use Cases

### Session Management

```js
import kv

// Store user session with 30-minute expiration
func createSession(userId, sessionData) {
    let sessionId = "session:" + userId
    kv.setTTL(sessionId, sessionData, 1800) // 30 minutes
    return sessionId
}

// Check if session is valid
func isSessionValid(sessionId) {
    return kv.exists(sessionId)
}

// Extend session
func extendSession(sessionId) {
    if (kv.exists(sessionId)) {
        kv.expire(sessionId, 1800) // Extend by 30 minutes
        return true
    }
    return false
}
```

### Caching

```js
import kv

// Cache expensive computation results
func getCachedResult(cacheKey, computeFunc) {
    // Check cache first
    let cached = kv.get(cacheKey)
    if (cached != null) {
        return cached
    }

    // Compute and cache result
    let result = computeFunc()
    kv.setTTL(cacheKey, result, 300) // Cache for 5 minutes
    return result
}
```

### Rate Limiting

```js
import kv

// Simple rate limiter
func isRateLimited(userId, limit, windowSeconds) {
    let key = "rate:" + userId
    let current = kv.get(key)

    if (current == null) {
        // First request in window
        kv.setTTL(key, 1, windowSeconds)
        return false
    }

    if (current >= limit) {
        return true // Rate limited
    }

    // Increment counter
    kv.increment(key)
    return false
}

// Usage: Allow 100 requests per minute
if (isRateLimited("user123", 100, 60)) {
    println("Rate limited!")
}
```

### Counters and Metrics

```js
import kv

// Track page views
func trackPageView(page) {
    kv.increment("views:" + page)
    kv.increment("total:views")
}

// Track user actions
func trackUserAction(userId, action) {
    kv.increment("user:" + userId + ":actions")
    kv.increment("action:" + action + ":count")
}

// Get metrics
func getMetrics() {
    return {
        "total_views": kv.get("total:views"),
        "stats": kv.stats(),
        "all_counters": kv.dump()
    }
}
```

### Configuration Management

```js
import kv

// Load configuration
func loadConfig() {
    let defaultConfig = {
        "app:theme": "light",
        "app:language": "en",
        "app:debug": false,
        "cache:ttl": 3600
    }
    kv.mset(defaultConfig)
}

// Update configuration
func updateConfig(key, value) {
    kv.set("config:" + key, value)
}

// Get all configuration
func getConfig() {
    let allData = kv.dump()
    let config = {}

    for (key, value in allData) {
        if (key.startsWith("config:")) {
            config[key] = value
        }
    }

    return config
}
```

## Performance Considerations

### Thread Safety

- All operations are thread-safe using read-write mutexes
- Multiple concurrent reads are allowed
- Writes are exclusive and properly synchronized

### Memory Management

- Expired keys are automatically cleaned up when accessed
- Use `clear()` to free all memory when done
- Monitor using `stats()` to track memory usage

### Bulk Operations

- Use `mget()` and `mset()` for better performance with multiple keys
- Bulk operations are more efficient than multiple single operations

### TTL Best Practices

- Set appropriate TTL values to prevent memory leaks
- Use `getTTL()` to check remaining time before operations
- Consider using `expire()` to extend TTL for active sessions

## Error Handling

All KV functions return appropriate error messages for invalid usage:

```js
// Invalid key type
let result = kv.get(123); // Error: key must be a string

// Invalid TTL
let result = kv.expire("key", -5); // Error: TTL must be positive

// Invalid increment target
kv.set("text", "hello");
let result = kv.increment("text"); // Error: existing value must be an integer
```

## Integration Examples

### With HTTP Server

```js
import kv
import http

// Simple API with caching
let app = http.app()

app.get("/api/user/:id", func(req, res) {
    let userId = req.params.id
    let cacheKey = "user:" + userId

    // Try cache first
    let user = kv.get(cacheKey)
    if (user != null) {
        return res.json({"user": user, "cached": true})
    }

    // Fetch from database (simulated)
    user = fetchUserFromDB(userId)

    // Cache for 10 minutes
    kv.setTTL(cacheKey, user, 600)

    res.json({"user": user, "cached": false})
})

// Track API calls
app.use(func(req, res, next) {
    kv.increment("api:calls")
    kv.increment("api:endpoint:" + req.path)
    next()
})
```

### With Async Operations

```js
import kv

// Async cache implementation
async func getCachedOrFetch(key, fetchFunc) {
    let cached = kv.get(key)
    if (cached != null) {
        return cached
    }

    let result = await fetchFunc()
    kv.setTTL(key, result, 300)
    return result
}
```

The KV module provides a robust foundation for in-memory data storage and caching in VintLang applications, with enterprise-grade features suitable for production use.

```

## llm.md

```markdown
# LLM & OpenAI Module

This module provides access to OpenAI's GPT models for chat and text completion from VintLang scripts.

## Setup

1. **Get an OpenAI API Key:**
   - Sign up at https://platform.openai.com/ and create an API key.
2. **Set the API Key in your environment:**
   - On macOS/Linux: `export OPENAI_API_KEY=sk-...`
   - On Windows: `set OPENAI_API_KEY=sk-...`

## Functions

### `llm.chat(messages, model="gpt-3.5-turbo", max_tokens=128, temperature=0.7)`
- **messages:** List of message objects (`{"role": "user", "content": "..."}`)
- **model:** (optional) Model name (default: gpt-3.5-turbo)
- **max_tokens:** (optional) Max tokens in response
- **temperature:** (optional) Sampling temperature
- **Returns:** (response, error)

### `llm.completion(prompt, model="text-davinci-003", max_tokens=128, temperature=0.7)`
- **prompt:** String prompt
- **model:** (optional) Model name (default: text-davinci-003)
- **max_tokens:** (optional) Max tokens in response
- **temperature:** (optional) Sampling temperature
- **Returns:** (completion, error)

## Example Usage

```js
import llm

messages = [
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "Tell me a joke."}
]
response, err = llm.chat(messages)
if err != null {
    print("Chat error: ", err)
} else {
    print("Chat response: ", response)
}

completion, err = llm.completion("Write a poem about the stars.")
if err != null {
    print("Completion error: ", err)
} else {
    print("Completion: ", completion)
}
```

## Notes

- Requires an internet connection.
- Make sure your API key is kept secret.
- See OpenAI docs for more on models and parameters.

```

## logger.md

```markdown
# Logger Module in Vint

The Logger module in Vint provides structured logging functionality with different severity levels and timestamps. This module helps you log messages to stdout and stderr with proper formatting and timestamps.

---

## Importing the Logger Module

To use the Logger module, simply import it:
```js
import logger
```

---

## Functions and Examples

### 1. Log Info Messages (`info`)

The `info` function logs informational messages to stdout with a timestamp and INFO level indicator.

**Syntax**:

```js
info(message)
```

**Example**:

```js
import logger

logger.info("Application started successfully")
// Output: [2024-01-15 14:30:25] INFO: Application started successfully
```

---

### 2. Log Warning Messages (`warn`)

The `warn` function logs warning messages to stdout with a timestamp and WARN level indicator.

**Syntax**:

```js
warn(message)
```

**Example**:

```js
import logger

logger.warn("This is a warning message")
// Output: [2024-01-15 14:30:25] WARN: This is a warning message
```

---

### 3. Log Error Messages (`error`)

The `error` function logs error messages to stderr with a timestamp and ERROR level indicator.

**Syntax**:

```js
error(message)
```

**Example**:

```js
import logger

logger.error("Something went wrong")
// Output: [2024-01-15 14:30:25] ERROR: Something went wrong
```

---

### 4. Log Debug Messages (`debug`)

The `debug` function logs debug information to stdout with a timestamp and DEBUG level indicator.

**Syntax**:

```js
debug(message)
```

**Example**:

```js
import logger

logger.debug("Debug information for troubleshooting")
// Output: [2024-01-15 14:30:25] DEBUG: Debug information for troubleshooting
```

---

### 5. Log Fatal Messages (`fatal`)

The `fatal` function logs fatal error messages to stderr with a timestamp and FATAL level indicator.

**Syntax**:

```js
fatal(message)
```

**Example**:

```js
import logger

logger.fatal("Critical system failure")
// Output: [2024-01-15 14:30:25] FATAL: Critical system failure
```

---

## Usage Example

```js
import logger

print("=== Logger Module Example ===")

// Log application lifecycle
logger.info("Application starting...")
logger.debug("Loading configuration...")
logger.info("Configuration loaded successfully")

// Simulate some warnings and errors
logger.warn("Low disk space detected")
logger.error("Failed to connect to database")
logger.fatal("Unable to recover from critical error")
```

---

## Summary of Functions

| Function | Description                                           | Output Destination |
|----------|------------------------------------------------------|--------------------|
| `info`   | Logs informational messages with timestamp           | stdout             |
| `warn`   | Logs warning messages with timestamp                  | stdout             |
| `error`  | Logs error messages with timestamp                    | stderr             |
| `debug`  | Logs debug information with timestamp                 | stdout             |
| `fatal`  | Logs fatal error messages with timestamp              | stderr             |

All log messages include timestamps in the format `YYYY-MM-DD HH:MM:SS` for easy tracking and debugging.

```

## main-function.md

```markdown
# Main Function in VintLang

VintLang now supports a **main function** as an entry point for your programs, similar to Go, Zig, C, and C++. This provides a structured way to organize your code while maintaining full backward compatibility.

## How It Works

### Two-Phase Execution Model

1. **Setup Phase**: All top-level statements are executed to define functions, variables, and perform initialization
2. **Main Phase**: If a `main` function is found, it's automatically called as the program's entry point

### Main Function Syntax

Use VintLang's existing function syntax to define a main function:

```javascript
let main = func() {
    println("Hello from main!")
    return 0
}
```

Or with `const`:

```javascript
const main = func() {
    println("Hello from main!")
    return "success"
}
```

## Complete Example

```javascript
// Setup phase - runs first
import time
println("🚀 Program starting...")

// Define helper functions
let greet = func(name) {
    println("Hello,", name, "!")
}

let calculate = func(a, b) {
    return a + b
}

// Main function - entry point
let main = func() {
    println("=== Main Function ===")
    
    greet("Developer")
    let result = calculate(10, 20)
    println("Result:", result)
    println("Time:", time.now())
    
    return result
}

// More setup
println("⚙️ Setup complete")
```

**Output:**

```
🚀 Program starting...
⚙️ Setup complete
=== Main Function ===
Hello, Developer !
Result: 30
Time: 16:42:00 26-09-2025
30
```

## Main Function Features

### Parameters

Main functions currently don't receive command-line arguments, but this could be extended in the future:

```javascript
let main = func() {
    // No parameters for now
    println("Main executed")
}
```

### Return Values

The main function's return value becomes the program's final result:

```javascript
let main = func() {
    return 42  // This will be printed as the program output
}
```

### Error Handling

If the main function returns an error, it will be propagated:

```javascript
let main = func() {
    if (someCondition) {
        return error("Something went wrong")
    }
    return "success"
}
```

## Backward Compatibility

Programs **without** a main function continue to work exactly as before:

```javascript
// This still works - no main function needed
println("Hello World")
let x = 42
println("x =", x)
```

## When to Use Main Functions

### Use main functions when

- Building larger, structured programs
- You want clear separation between setup and execution
- Coming from Go, C, C++, or similar languages
- Building command-line tools or applications

### Stick with the traditional approach when

- Writing simple scripts
- Prototyping or testing small code snippets
- Using VintLang in REPL mode
- Personal preference for simpler structure

## Execution Flow

1. **Parse** the entire program
2. **Setup Phase**: Execute all top-level statements
   - Define variables with `let` and `const`
   - Define functions
   - Run imports
   - Execute setup code
3. **Main Phase**: If `main` function exists, call it
4. **Return** either the main function's result or the last statement's result

## Migration Guide

To convert existing VintLang programs to use main functions:

**Before:**

```javascript
import time
let x = 42
println("Hello World")
println("Time:", time.now())
```

**After:**

```javascript
import time  // Still runs in setup phase

let main = func() {
    let x = 42
    println("Hello World")  
    println("Time:", time.now())
}
```

The behavior remains identical, but the code is now more structured and explicit about the entry point.

```

## math.md

```markdown
# Math Module

## Usage

To use the `math` module, import it into your Vint script:

```js
import math
```

You can then call functions and access constants from the module:

```js
println(math.PI())
println(math.sqrt(16))
```

## Contents

This module provides a wide range of mathematical functions and constants, including:

- Basic Mathematical Functions
- Hyperbolic & Trigonometric Functions
- Exponential & Logarithmic Functions
- Rounding & Comparison Functions

Here is a complete list of the available functions and constants:

### Constants

- **PI**: Represents the mathematical constant `π` (3.14159...).
- **e**: Represents Euler's Number (2.71828...).
- **phi**: Represents the Golden Ratio (1.61803...).
- **ln10**: Represents the natural logarithm of 10.
- **ln2**: Represents the natural logarithm of 2.
- **log10e**: Represents the base-10 logarithm of `e`.
- **log2e**: Represents the base-2 logarithm of `e`.
- **sqrt1_2**: Represents the square root of 1/2.
- **sqrt2**: Represents the square root of 2.
- **sqrt3**: Represents the square root of 3.
- **sqrt5**: Represents the square root of 5.
- **EPSILON**: Represents a very small number, often used for float comparisons.

### Functions

#### `abs(n)`

- **Description**: Calculates the absolute value of a number.
- **Example**: `math.abs(-42)` returns `42`.

#### `acos(n)`

- **Description**: Calculates the arccosine (inverse cosine) of a number in radians.
- **Example**: `math.acos(0.5)` returns `1.047...`.

#### `acosh(n)`

- **Description**: Calculates the inverse hyperbolic cosine of a number.
- **Example**: `math.acosh(2.0)` returns `1.316...`.

#### `asin(n)`

- **Description**: Calculates the arcsine (inverse sine) of a number in radians.
- **Example**: `math.asin(0.5)` returns `0.523...`.

#### `asinh(n)`

- **Description**: Calculates the inverse hyperbolic sine of a number.
- **Example**: `math.asinh(2.0)` returns `1.443...`.

#### `atan(n)`

- **Description**: Calculates the arctangent (inverse tangent) of a number in radians.
- **Example**: `math.atan(1.0)` returns `0.785...`.

#### `atan2(y, x)`

- **Description**: Calculates the arctangent of the quotient of its arguments (`y/x`) in radians.
- **Example**: `math.atan2(1.0, 1.0)` returns `0.785...`.

#### `atanh(n)`

- **Description**: Calculates the inverse hyperbolic tangent of a number.
- **Example**: `math.atanh(0.5)` returns `0.549...`.

#### `cbrt(n)`

- **Description**: Calculates the cube root of a number.
- **Example**: `math.cbrt(8)` returns `2.0`.

#### `ceil(n)`

- **Description**: Rounds a number up to the nearest integer.
- **Example**: `math.ceil(4.3)` returns `5`.

#### `cos(n)`

- **Description**: Calculates the cosine of an angle (in radians).
- **Example**: `math.cos(0.0)` returns `1.0`.

#### `cosh(n)`

- **Description**: Calculates the hyperbolic cosine of a number.
- **Example**: `math.cosh(0.0)` returns `1.0`.

#### `exp(n)`

- **Description**: Calculates `e` raised to the power of `n`.
- **Example**: `math.exp(2.0)` returns `7.389...`.

#### `expm1(n)`

- **Description**: Calculates `e` raised to the power of a number, minus 1.
- **Example**: `math.expm1(1.0)` returns `1.718...`.

#### `factorial(n)`

- **Description**: Calculates the factorial of a non-negative integer.
- **Example**: `math.factorial(5)` returns `120`.

#### `floor(n)`

- **Description**: Rounds a number down to the nearest integer.
- **Example**: `math.floor(4.7)` returns `4`.

#### `hypot(numbers)`

- **Description**: Calculates the square root of the sum of the squares of the numbers in an array.
- **Example**: `math.hypot([3, 4])` returns `5.0`.

#### `log10(n)`

- **Description**: Calculates the base-10 logarithm of a number.
- **Example**: `math.log10(100.0)` returns `2.0`.

#### `log1p(n)`

- **Description**: Calculates the natural logarithm of 1 plus the given number.
- **Example**: `math.log1p(1.0)` returns `0.693...`.

#### `log2(n)`

- **Description**: Calculates the base-2 logarithm of a number.
- **Example**: `math.log2(8)` returns `3.0`.

#### `max(numbers)`

- **Description**: Finds the maximum value in an array of numbers.
- **Example**: `math.max([4, 2, 9, 5])` returns `9.0`.

#### `min(numbers)`

- **Description**: Finds the minimum value in an array of numbers.
- **Example**: `math.min([4, 2, 9, 5])` returns `2.0`.

#### `random()`

- **Description**: Returns a random floating-point number between 0.0 and 1.0.
- **Example**: `math.random()` returns a value like `0.12345...`.

#### `round(n)`

- **Description**: Rounds a floating-point number to the nearest integer.
- **Example**: `math.round(4.6)` returns `5`.

#### `root(x, n)`

- **Description**: Calculates the nth root of a number `x`.
- **Example**: `math.root(27, 3)` returns `3.0`.

#### `sign(n)`

- **Description**: Returns the sign of a number (`-1` for negative, `0` for zero, `1` for positive).
- **Example**: `math.sign(-5)` returns `-1`.

#### `sin(n)`

- **Description**: Calculates the sine of an angle (in radians).
- **Example**: `math.sin(1.0)` returns `0.841...`.

#### `sinh(n)`

- **Description**: Calculates the hyperbolic sine of a number.
- **Example**: `math.sinh(1.0)` returns `1.175...`.

#### `sqrt(n)`

- **Description**: Calculates the square root of a number.
- **Example**: `math.sqrt(4)` returns `2.0`.

#### `tan(n)`

- **Description**: Calculates the tangent of an angle (in radians).
- **Example**: `math.tan(1.0)` returns `1.557...`.

#### `tanh(n)`

- **Description**: Calculates the hyperbolic tangent of a number.
- **Example**: `math.tanh(1.0)` returns `0.761...`.

---

## Statistics Functions

#### `mean(array)`

- **Description**: Calculates the arithmetic mean (average) of an array of numbers.
- **Example**: `math.mean([1, 2, 3, 4, 5])` returns `3.0`.

#### `median(array)`

- **Description**: Calculates the median (middle value) of an array of numbers.
- **Example**: `math.median([1, 2, 3, 4, 5])` returns `3.0`.
- **Note**: For even-length arrays, returns the average of the two middle values.

#### `variance(array)`

- **Description**: Calculates the variance of an array of numbers.
- **Example**: `math.variance([1, 2, 3, 4, 5])` returns `2.0`.

#### `stddev(array)`

- **Description**: Calculates the standard deviation of an array of numbers.
- **Example**: `math.stddev([1, 2, 3, 4, 5])` returns `1.414...`.

---

## Complex Numbers

#### `complex(real, imag)`

- **Description**: Creates a complex number with the given real and imaginary parts.
- **Returns**: A dictionary with `real` and `imag` keys.
- **Example**:

  ```js
  let c = math.complex(3, 4)
  print(c["real"])  // 3
  print(c["imag"])  // 4
  ```

#### `abs(n)` (extended)

- **Description**: Also works with complex numbers to calculate magnitude.
- **Example**:

  ```js
  let c = math.complex(3, 4)
  math.abs(c)  // returns 5.0
  ```

---

## Arbitrary Precision

#### `bigint(value)`

- **Description**: Creates a big integer representation for arbitrary precision arithmetic.
- **Parameters**: A string or integer representing a large number.
- **Returns**: A dictionary with `value` (string) and `type` ("bigint") keys.
- **Example**:

  ```js
  let big = math.bigint("999999999999999999999")
  print(big["value"])  // "999999999999999999999"
  ```

---

## Linear Algebra

#### `dot(array1, array2)`

- **Description**: Calculates the dot product of two vectors (arrays).
- **Example**: `math.dot([1, 2, 3], [4, 5, 6])` returns `32.0`.

#### `cross(array1, array2)`

- **Description**: Calculates the cross product of two 3D vectors.
- **Example**: `math.cross([1, 2, 3], [4, 5, 6])` returns `[-3, 6, -3]`.

#### `magnitude(array)`

- **Description**: Calculates the magnitude (length) of a vector.
- **Example**: `math.magnitude([3, 4])` returns `5.0`.

---

## Numerical Methods

#### `gcd(a, b)`

- **Description**: Calculates the greatest common divisor of two integers.
- **Example**: `math.gcd(48, 18)` returns `6`.

#### `lcm(a, b)`

- **Description**: Calculates the least common multiple of two integers.
- **Example**: `math.lcm(12, 15)` returns `60`.

#### `clamp(value, min, max)`

- **Description**: Clamps a value between a minimum and maximum.
- **Example**: `math.clamp(15, 0, 10)` returns `10.0`.

#### `lerp(start, end, t)`

- **Description**: Linear interpolation between two values.
- **Parameters**: `start` and `end` values, and `t` (0.0 to 1.0) as the interpolation factor.
- **Example**: `math.lerp(0, 10, 0.5)` returns `5.0`.

```

## modules.md

```markdown
# Modules in VintLang

Modules are a way to organize and reuse code in VintLang. This guide explains how to create and use modules.

## Creating a Module

To create a module, use the `module` keyword followed by the module name and a block of code:

```js
module math {
    func add(a, b) {
        return a + b
    }
    
    func subtract(a, b) {
        return a - b
    }
}
```

## Using Modules

To use a module, you need to:

1. Create a file with the `.vint` extension in one of these locations:
   - Current working directory
   - `./modules` directory
   - `./vintLang/modules` directory

2. Import the module using the `import` statement:

```js
import math

result = math.add(5, 3)
print(result)  // Output: 8
```

## Module Structure

A module can contain:

- Functions
- Variables
- Other modules
- Any valid VintLang code

Example of a more complex module:

```js
module utils {
    let version = "1.0.0"
    
    func format(text) {
        return "Formatted: " + text
    }
    
    module helpers {
        func validate(input) {
            return input != null
        }
    }
}
```

## Best Practices

1. Keep modules focused on a single responsibility
2. Use descriptive names for modules
3. Document your modules with comments
4. Place related modules in the same directory
5. Use the `modules` directory for reusable code

## Error Handling

If a module is not found, you'll see an error message like this:

```
Module 'math' not found.

To fix this:
1. Create a file named 'math.vint' in one of these locations:
  1. /current/working/directory
  2. /current/working/directory/modules
2. Make sure the file contains valid VintLang code
3. Try importing again
```

## Module Scope

Variables and functions defined in a module are only accessible within that module unless explicitly exported. This helps prevent naming conflicts and keeps code organized.

## Example: Creating a Custom Module

Here's a complete example of creating and using a custom module:

```js
// file: modules/calculator.vint
module calculator {
    func add(a, b) {
        return a + b
    }
    
    func subtract(a, b) {
        return a - b
    }
    
    func multiply(a, b) {
        return a * b
    }
    
    func divide(a, b) {
        if b == 0 {
            return "Error: Division by zero"
        }
        return a / b
    }
}

// file: main.vint
import calculator

result1 = calculator.add(10, 5)
result2 = calculator.multiply(4, 3)

print("Addition: " + result1)      // Output: Addition: 15
print("Multiplication: " + result2) // Output: Multiplication: 12
```

## Built-in Modules

VintLang provides several built-in modules for common functionality:

### Core Modules

- **`math`** - Mathematical functions (`abs`, `sqrt`, `sin`, `cos`, etc.)
- **`string`** - String manipulation (`toUpper`, `toLower`, `trim`, etc.)
- **`random`** - Random number generation (`int`, `float`, `choice`, etc.)
- **`kv`** - In-memory key-value store with TTL support and atomic operations

### I/O and System

- **`os`** - Comprehensive operating system interface (file operations, process management, environment variables, permissions, links, system info)
- **`time`** - Date and time utilities
- **`datetime`** - Advanced date/time formatting and parsing
- **`path`** - File path manipulation
- **`shell`** - Shell command execution

### Network and Web

- **`net`** - Network utilities
- **`http`** - HTTP client and server functionality
- **`url`** - URL parsing and manipulation
- **`email`** - Email sending capabilities

### Data Processing

- **`json`** - JSON parsing and serialization
- **`csv`** - CSV file processing
- **`xml`** - XML parsing and generation
- **`yaml`** - YAML parsing and serialization
- **`encoding`** - Text encoding utilities

### Security and Crypto

- **`crypto`** - Cryptographic functions
- **`hash`** - Hashing algorithms (MD5, SHA1, SHA256, etc.)

### Database

- **`sqlite`** - SQLite database interface
- **`mysql`** - MySQL database connectivity
- **`postgres`** - PostgreSQL database connectivity
- **`redis`** - Redis client

### Development Tools

- **`regex`** - Regular expression support
- **`term`** - Terminal manipulation
- **`cli`** - Command-line interface utilities
- **`logger`** - Logging functionality
- **`uuid`** - UUID generation
- **`reflect`** - Runtime reflection

### Specialized

- **`schedule`** - Task scheduling
- **`dotenv`** - Environment file loading
- **`sysinfo`** - System information
- **`clipboard`** - System clipboard access
- **`vintSocket`** - WebSocket support
- **`vintChart`** - Chart generation
- **`llm`** - Large Language Model integration
- **`openai`** - OpenAI API client

Example usage of built-in modules:

```js
import kv
import math
import json

// Use KV store for caching
kv.set("pi", math.pi())
let cached_pi = kv.get("pi")

// Store complex data with TTL
let user_data = {"name": "Alice", "score": 100}
kv.setTTL("user:123", user_data, 3600) // 1 hour TTL

// Use atomic operations
kv.increment("page_views")
kv.decrement("inventory", 5)

// Export data
let all_data = kv.dump()
let json_export = json.stringify(all_data)
```

See individual module documentation for detailed API reference.

```

## mysql.md

```markdown
# MySQL Module in VintLang

The `mysql` module in **VintLang** provides a way to interact with MySQL databases. You can connect to a database, execute queries, and fetch data.

## Connecting to a MySQL Database

To connect to a MySQL database, use `mysql.open()`. You need to provide a connection string in the following format: `user:password@tcp(host:port)/dbname`.

```js
conn = mysql.open("user:password@tcp(127.0.0.1:3306)/testdb")
```

## Closing the Connection

Always close the connection when you're done with `mysql.close()`.

```js
mysql.close(conn)
```

## Executing Queries

Use `mysql.execute()` for `INSERT`, `UPDATE`, `DELETE`, or any other queries that don't return rows.

```js
// Inserting data with placeholders
insert_query = "INSERT INTO users (name, age) VALUES (?, ?)"
mysql.execute(conn, insert_query, "Alice", 30)
```

## Fetching Data

### Fetch All Rows

To get all rows from a query result, use `mysql.fetchAll()`.

```js
users = mysql.fetchAll(conn, "SELECT * FROM users")
print(users)
```

### Fetch a Single Row

To get only the first row from a query result, use `mysql.fetchOne()`.

```js
user = mysql.fetchOne(conn, "SELECT * FROM users WHERE id = ?", 1)
print(user)
```

## Full Example

Here's a complete example of how to use the `mysql` module:

```js
import mysql

// Replace with your actual credentials
conn_str = "user:password@tcp(127.0.0.1:3306)/testdb"
conn = mysql.open(conn_str)

if conn.type() == "ERROR" {
    print("Error connecting to MySQL:", conn.message())
} else {
    print("Successfully connected to MySQL")

    // Create a table
    create_query = "CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), age INT)"
    mysql.execute(conn, create_query)

    // Insert data
    mysql.execute(conn, "INSERT INTO users (name, age) VALUES (?, ?)", "Bob", 35)

    // Fetch and print data
    users = mysql.fetchAll(conn, "SELECT * FROM users")
    print("All users:", users)

    // Close the connection
    mysql.close(conn)
    print("Connection closed")
} 
```

## net.md

```markdown
# Net Module

The `net` module provides functions for making HTTP requests (GET, POST, PUT, DELETE, PATCH) from your Vint scripts. Each function supports both simple and advanced usage, allowing you to specify URLs, headers, and request bodies.

## Functions

### 1. `net.get`

**Description:**
Sends an HTTP GET request.

**Usage:**
```js
net.get("https://api.example.com/data")
```

Or with named arguments:

```js
net.get(
  url: "https://api.example.com/data",
  headers: {"Authorization": "Bearer token"},
  body: {"key": "value"}  # Optional, rarely used for GET
)
```

**Arguments:**

- `url` (string, required): The URL to request.
- `headers` (dict, optional): HTTP headers as key-value pairs.
- `body` (dict, optional): Data to send as JSON (rare for GET).

**Returns:**
Response body as a string, or an error object.

---

### 2. `net.post`

**Description:**
Sends an HTTP POST request.

**Usage:**

```js
net.post(
  url: "https://api.example.com/data",
  headers: {"Authorization": "Bearer token"},
  body: {"key": "value"}
)
```

**Arguments:**

- `url` (string, required): The URL to request.
- `headers` (dict, optional): HTTP headers as key-value pairs.
- `body` (dict, optional): Data to send as JSON.

**Returns:**
Response body as a string, or an error object.

---

### 3. `net.put`

**Description:**
Sends an HTTP PUT request.

**Usage:**

```js
net.put(
  url: "https://api.example.com/data/1",
  headers: {"Authorization": "Bearer token"},
  body: {"key": "new value"}
)
```

**Arguments:**

- `url` (string, required): The URL to request.
- `headers` (dict, optional): HTTP headers as key-value pairs.
- `body` (dict, optional): Data to send as JSON.

**Returns:**
Response body as a string, or an error object.

---

### 4. `net.delete`

**Description:**
Sends an HTTP DELETE request.

**Usage:**

```js
net.delete(
  url: "https://api.example.com/data/1",
  headers: {"Authorization": "Bearer token"}
)
```

**Arguments:**

- `url` (string, required): The URL to request.
- `headers` (dict, optional): HTTP headers as key-value pairs.

**Returns:**
Response body as a string, or an error object.

---

### 5. `net.patch`

**Description:**
Sends an HTTP PATCH request.

**Usage:**

```js
net.patch(
  url: "https://api.example.com/data/1",
  headers: {"Authorization": "Bearer token"},
  body: {"key": "patched value"}
)
```

**Arguments:**

- `url` (string, required): The URL to request.
- `headers` (dict, optional): HTTP headers as key-value pairs.
- `body` (dict, optional): Data to send as JSON.

**Returns:**
Response body as a string, or an error object.

---

## Notes

- All functions return the response body as a string, or an error object if something goes wrong.
- Named arguments (`url`, `headers`, `body`) are recommended for clarity.
- Headers and body must be dictionaries.
- For GET requests, the body is rarely used and may not be supported by all servers.

```

## note.md

```markdown
# Note

The `note` keyword allows you to print general notes at runtime.

## Syntax

```js
note "Your note message here"
```

When the Vint interpreter encounters a `note` statement, it prints a blue-colored note message to the console and continues execution. This is useful for providing additional context or reminders in your scripts.

### Example

```js
note "This script was last updated on 2024-06-01."
println("Script running...")
```

Running this script will output:

```
[NOTE]: This script was last updated on 2024-06-01.
Script running...
```

```

## null.md

```markdown
# Null in Vint

The `null` data type in Vint represents the absence of a value or the concept of "nothing" or "empty." This page covers the syntax and usage of the `null` data type in Vint, including its definition and evaluation.

## Definition

A `null` data type is a data type with no value, defined with the `null` keyword:

```js
let a = null
```

## Evaluation

When evaluating a `null` data type in a conditional expression, it will evaluate to `false`:

```js
if (a) {
    print("a is null")
} else {
    print("a has a value")
}

// Output: a has a value
```

## Null Methods

The `null` data type in Vint comes with several utility methods:

### isNull()

Always returns `true` for null values:

```js
let value = null
print(value.isNull())  // true
```

### coalesce()

Returns the first non-null value from the arguments:

```js
let value = null
let result = value.coalesce("default", "backup")
print(result)  // "default"
```

### ifNull()

Returns the provided value if this is null:

```js
let value = null
let result = value.ifNull("default value")
print(result)  // "default value"
```

### toString()

Returns the string representation of null:

```js
let value = null
print(value.toString())  // "null"
```

### equals()

Checks if another value is also null:

```js
let value1 = null
let value2 = null
let value3 = "something"
print(value1.equals(value2))  // true
print(value1.equals(value3))  // false
```

The `null` data type is useful in Vint when you need to represent an uninitialized, missing, or undefined value in your programs. By understanding the `null` data type and its methods, you can create more robust and flexible code.

```

## numbers.md

```markdown
# INTEGERS  AND FLOATS 

Integers and floats are the basic numeric data types in vint, used for representing whole numbers and decimal numbers, respectively. This page covers the syntax and usage of integers and floats in vint, including precedence, unary increments, shorthand assignments, and negative numbers.

## PRECEDENCE

Integers and floats behave as expected in mathematical operations, following the BODMAS rule:
```go
2 + 3 * 5 // 17

let a = 2.5
let b = 3/5

a + b // 2.8
```

## UNARY INCREMENTS

You can perform unary increments (++ and --) on both floats and integers. These will add or subtract 1 from the current value. Note that the float or int have to be assigned to a variable for this operation to work. Here's an example:

```go
let i = 2.4

i++ // 3.4
```

## SHORTHAND ASSIGNMENT

vint supports shorthand assignments with +=, -=, /=, *=, and %=:
You

```go
let i = 2

i *= 3 // 6
i /= 2 // 3
i += 100 // 103
i -= 10 // 93
i %= 90 // 3
```

## NEGATIVE NUMBERS

Negative numbers also behave as expected:

```go
let i = -10

while (i < 0) {
    print(i)
    i++
}

```

Output:

```s
-10
-9
-8
-7
-6
-5
-4
-3
-2
-1
0
1
2
3
4
5
6
7
8
9 
```

## Integer Methods

Integers in vint have several built-in methods:

### abs()

Returns the absolute value of the integer:

```s
let i = -42
print(i.abs())  // 42
```

### is_even()

Returns true if the integer is even, false otherwise:

```s
let i = 4
print(i.is_even())  // true
print((5).is_even())  // false
```

### is_odd()

Returns true if the integer is odd, false otherwise:

```s
let i = 7
print(i.is_odd())  // true
print((8).is_odd())  // false
```

### to_string()

Converts the integer to a string:

```s
let i = 123
print(i.to_string())  // "123"
```

### sign()

Returns 1 if the integer is positive, -1 if negative, or 0 if zero:

```s
print((10).sign())   // 1
print((-5).sign())   // -1
print((0).sign())    // 0
```

### pow()

Raises the integer to the power of another number:

```s
let base = 2
print(base.pow(3))   // 8
print((5).pow(2))    // 25
```

### sqrt()

Returns the square root of the integer as a float:

```s
let num = 16
print(num.sqrt())    // 4.0
print((25).sqrt())   // 5.0
```

### gcd()

Returns the greatest common divisor of two integers:

```s
let a = 24
let b = 18
print(a.gcd(b))      // 6
print((48).gcd(18))  // 6
```

### lcm()

Returns the least common multiple of two integers:

```s
let a = 12
let b = 8
print(a.lcm(b))      // 24
print((15).lcm(20))  // 60
```

### factorial()

Returns the factorial of the integer:

```s
let n = 5
print(n.factorial()) // 120
print((4).factorial()) // 24
print((0).factorial()) // 1
```

### toBinary()

Converts the integer to binary representation:

```s
let num = 255
print(num.toBinary())  // "11111111"
print((5).toBinary())  // "101"
```

### toHex()

Converts the integer to hexadecimal representation:

```s
let num = 255
print(num.toHex())     // "ff"
print((16).toHex())    // "10"
```

### toOctal()

Converts the integer to octal representation:

```s
let num = 64
print(num.toOctal())   // "100"
print((8).toOctal())   // "10"
```

### isPrime()

Checks if the integer is a prime number:

```s
print((17).isPrime())  // true
print((4).isPrime())   // false
print((2).isPrime())   // true
print((1).isPrime())   // false
```

### nthRoot()

Calculates the nth root of the integer:

```s
let num = 8
print(num.nthRoot(3))  // 2.0 (cube root)
print((16).nthRoot(2)) // 4.0 (square root)
```

### mod()

Calculates the modulo (remainder) with another integer:

```s
let num = 10
print(num.mod(3))      // 1
print((15).mod(4))     // 3
```

### clamp()

Restricts the integer to be within specified bounds:

```s
let num = 15
print(num.clamp(1, 10))  // 10 (clamped to max)
print((-5).clamp(1, 10)) // 1 (clamped to min)
print((5).clamp(1, 10))  // 5 (within bounds)
```

### inRange()

Checks if the integer is within the specified range (inclusive):

```s
let num = 5
print(num.inRange(1, 10))  // true
print((15).inRange(1, 10)) // false
print((0).inRange(1, 10))  // false
```

### digits()

Returns an array of individual digits:

```s
let num = 123
print(num.digits())    // [1, 2, 3]
print((456).digits())  // [4, 5, 6]
```

## Float Methods

Floats in vint have powerful built-in methods for mathematical operations and utility functions:

### abs()

Returns the absolute value of the float:

```s
let f = -3.14
print(f.abs())       // 3.14
print((-2.5).abs())  // 2.5
```

### ceil()

Returns the smallest integer greater than or equal to the float:

```s
let price = 29.95
print(price.ceil())  // 30
print((4.1).ceil())  // 5
print((-2.1).ceil()) // -2
```

### floor()

Returns the largest integer less than or equal to the float:

```s
let price = 29.95
print(price.floor()) // 29
print((4.9).floor()) // 4
print((-2.1).floor()) // -3
```

### round()

Rounds the float to a specified number of decimal places:

```s
let pi = 3.14159
print(pi.round(2))   // 3.14
print(pi.round(0))   // 3
print((2.7).round()) // 3
```

### sqrt()

Returns the square root of the float:

```s
let num = 9.0
print(num.sqrt())    // 3.0
print((16.0).sqrt()) // 4.0
```

### pow()

Raises the float to the power of another number:

```s
let base = 2.5
print(base.pow(2))   // 6.25
print((3.0).pow(3))  // 27.0
```

### is_nan()

Checks if the float is NaN (Not a Number):

```s
let valid = 3.14
let invalid = 0.0 / 0.0
print(valid.is_nan())   // false
print(invalid.is_nan()) // true
```

### is_infinite()

Checks if the float is infinite:

```s
let normal = 3.14
let inf = 1.0 / 0.0
print(normal.is_infinite()) // false
print(inf.is_infinite())    // true
```

### to_string()

Converts the float to a string with optional precision:

```s
let price = 29.95
print(price.to_string())   // "29.95"
print(price.to_string(1))  // "30.0"
print((3.14159).to_string(2)) // "3.14"
```

### clamp()

Clamps the float between minimum and maximum values:

```s
let value = 75.5
print(value.clamp(0.0, 50.0))  // 50.0
print((25.3).clamp(30.0, 100.0)) // 30.0
print((45.7).clamp(10.0, 80.0))  // 45.7
```

### toPrecision()

Formats the float to specified precision:

```s
let num = 123.456789
print(num.toPrecision(4))    // "123.5"
print((0.123456).toPrecision(3)) // "0.123"
```

### toFixed()

Formats the float to fixed decimal places:

```s
let num = 123.456
print(num.toFixed(2))        // "123.46"
print((5.0).toFixed(3))      // "5.000"
```

### sign()

Returns the sign of the float:

```s
print((5.5).sign())          // 1.0
print((-3.2).sign())         // -1.0
print((0.0).sign())          // 0.0
```

### truncate()

Removes the fractional part:

```s
print((5.9).truncate())      // 5.0
print((-3.7).truncate())     // -3.0
```

### mod()

Calculates the floating-point remainder:

```s
let num = 5.5
print(num.mod(2.0))          // 1.5
print((10.7).mod(3.0))       // 1.7
```

### degrees()

Converts radians to degrees:

```s
import math
let pi = math.PI
print(pi.degrees())          // 180.0
print((pi / 2).degrees())    // 90.0
```

### radians()

Converts degrees to radians:

```s
print((180.0).radians())     // 3.141592653589793
print((90.0).radians())      // 1.5707963267948966
```

### sin()

Calculates the sine:

```s
print((0.0).sin())           // 0.0
print((math.PI / 2).sin())   // 1.0
```

### cos()

Calculates the cosine:

```s
print((0.0).cos())           // 1.0
print(math.PI.cos())         // -1.0
```

### tan()

Calculates the tangent:

```s
print((0.0).tan())           // 0.0
print((math.PI / 4).tan())   // 1.0
```

### log()

Calculates the natural logarithm:

```s
import math
print(math.E.log())          // 1.0
print((10.0).log())          // 2.302585092994046
```

### exp()

Calculates e raised to the power of the float:

```s
print((0.0).exp())           // 1.0
print((1.0).exp())           // 2.718281828459045
```

## Practical Examples

Here are some practical examples using integer and float methods:

```s
// Calculate compound interest
let principal = 1000.0
let rate = 0.05
let time = 3
let amount = principal * (1.0 + rate).pow(time)
print("Amount after", time, "years:", amount.round(2))

// Check if numbers are perfect squares
numbers = [16, 25, 30, 36]
for num in numbers {
    let sqrt_val = num.sqrt()
    if (sqrt_val.floor() == sqrt_val.ceil()) {
        print(num, "is a perfect square")
    }
}

// Mathematical calculations with bounds
let angle = 1.57079  // approximately π/2
let sin_approx = angle - angle.pow(3) / (3).factorial()
print("sin approximation:", sin_approx.round(6))

// Working with ranges and validation
let score = 87.5
let normalized = score.clamp(0.0, 100.0) / 100.0
print("Normalized score:", normalized.round(3))
```

```

## operators.md

```markdown
# Operators in Vint

Operators are a core feature of any programming language, enabling you to perform various operations on variables and values. This page details the syntax and usage of operators in Vint, including assignment, arithmetic, comparison, membership, and logical operators.

---

## Assignment Operators

Assignment operators are used to assign values to variables. The following are supported in Vint:

- `i = v`: Assigns the value of `v` to `i`.
- `i += v`: Equivalent to `i = i + v`.
- `i -= v`: Equivalent to `i = i - v`.
- `i *= v`: Equivalent to `i = i * v`.
- `i /= v`: Equivalent to `i = i / v`.

For strings, arrays, and dictionaries, the `+=` operator is also valid. For example:

```js
list1 += list2 // Equivalent to list1 = list1 + list2
```

---

## Arithmetic Operators

Vint supports the following arithmetic operations:

| Operator | Description                          | Example          |
|----------|--------------------------------------|------------------|
| `+`      | Addition                             | `2 + 3 = 5`      |
| `-`      | Subtraction                          | `5 - 2 = 3`      |
| `*`      | Multiplication                       | `3 * 4 = 12`     |
| `/`      | Division                             | `10 / 2 = 5`     |
| `%`      | Modulo (remainder of a division)     | `7 % 3 = 1`      |
| `**`     | Exponential power                   | `2 ** 3 = 8`     |

---

## Comparison Operators

Comparison operators evaluate relationships between two values. These return `true` or `false`:

| Operator | Description                     | Example            |
|----------|---------------------------------|--------------------|
| `==`     | Equal to                        | `5 == 5 // true`   |
| `!=`     | Not equal to                    | `5 != 3 // true`   |
| `>`      | Greater than                    | `5 > 3 // true`    |
| `>=`     | Greater than or equal to        | `5 >= 5 // true`   |
| `<`      | Less than                       | `3 < 5 // true`    |
| `<=`     | Less than or equal to           | `3 <= 3 // true`   |

---

## Membership Operator

The membership operator `in` checks if an item exists within a collection:

```js
names = ['juma', 'asha', 'haruna']

"haruna" in names // true
"halima" in names // false
```

---

## Logical Operators

Logical operators allow you to combine or invert conditions:

| Operator | Description              | Example                   |
|----------|--------------------------|---------------------------|
| `&&`     | Logical AND              | `true && false // false` |
| `||`     | Logical OR               | `true || false // true`  |
| `!`      | Logical NOT (negation)   | `!true // false`         |

---

## Precedence of Operators

When multiple operators are used in an expression, operator precedence determines the order of execution. Below is the precedence order, from highest to lowest:

1. `()` : Parentheses
2. `!`  : Logical NOT
3. `%`  : Modulo
4. `**` : Exponential power
5. `/`, `*` : Division and multiplication
6. `+`, `+=`, `-`, `-=` : Addition and subtraction
7. `>`, `>=`, `<`, `<=` : Comparison operators
8. `==`, `!=` : Equality and inequality
9. `=` : Assignment
10. `in` : Membership operator
11. `&&`, `||` : Logical AND and OR

---

```

## os.md

```markdown
# OS Module in VintLang

The **Vint** `os` module provides comprehensive functions to interact with the operating system, file system, processes, and environment. This module closely mirrors Go's standard `os` package functionality, offering powerful system-level operations.

## Table of Contents

- [Process Management](#process-management)
- [Environment Variables](#environment-variables)
- [File Operations](#file-operations)
- [Directory Operations](#directory-operations)
- [File Permissions and Ownership](#file-permissions-and-ownership)
- [File System Links](#file-system-links)
- [System Information](#system-information)
- [Error Checking](#error-checking)
- [User Directories](#user-directories)
- [Legacy Functions](#legacy-functions)

## Process Management

### Exit with Status Code

```js
os.exit(1)  // Exit with status code 1 (error)
os.exit(0)  // Exit with status code 0 (success)
```

### Run Shell Commands

```js
result = os.run("ls -la")
print(result)  // Outputs the directory listing
```

### Process Information

Get detailed information about the current process:

```js
// Process identifiers
pid = os.getpid()        // Current process ID
ppid = os.getppid()      // Parent process ID

// User and group identifiers
uid = os.getuid()        // Real user ID
gid = os.getgid()        // Real group ID
euid = os.geteuid()      // Effective user ID
egid = os.getegid()      // Effective group ID

// Get all groups the user belongs to
groups = os.getgroups()  // Returns array of group IDs
print("User groups:", groups)

// System information
pageSize = os.getpagesize()  // Memory page size
```

## Environment Variables

### Basic Environment Operations

```js
// Set environment variable
os.setEnv("API_KEY", "12345")

// Get environment variable
api_key = os.getEnv("API_KEY")
print(api_key)  // Outputs: "12345"

// Remove environment variable
os.unsetEnv("API_KEY")
```

### Advanced Environment Functions

```js
// Get all environment variables
envVars = os.environ()
for (env in envVars) {
    print(env)  // Each entry is "KEY=value"
}

// Clear all environment variables (use with caution!)
os.clearenv()

// Check if environment variable exists
result = os.lookupEnv("PATH")
if (result["exists"]) {
    print("PATH exists:", result["value"])
} else {
    print("PATH not found")
}

// Expand environment variables in strings
expanded = os.expandEnv("Home is $HOME and user is $USER")
print(expanded)

// Alternative expansion method
expanded2 = os.expand("$HOME/documents", os.getEnv)
```

## File Operations

### Basic File I/O

```js
// Write to a file
os.writeFile("example.txt", "Hello, Vint!")

// Read from a file
content = os.readFile("example.txt")
print(content)  // Outputs: "Hello, Vint!"

// Read file lines as array
lines = os.readLines("example.txt")
print(lines)  // Array of lines

// Check if file exists
exists = os.fileExists("example.txt")
print(exists)  // true or false
```

### Advanced File Operations

```js
// Get detailed file information
fileInfo = os.stat("example.txt")
print("File name:", fileInfo["name"])
print("File size:", fileInfo["size"])
print("Is directory:", fileInfo["isDir"])
print("Mode:", fileInfo["mode"])
print("Modification time:", fileInfo["modTime"])

// Get file info without following symlinks
linkInfo = os.lstat("symlink.txt")

// Truncate file to specific size
os.truncate("example.txt", 10)  // Truncate to 10 bytes

// Rename or move files
os.rename("old_name.txt", "new_name.txt")

// Remove files
os.remove("example.txt")        // Remove single file
os.removeAll("directory/")      // Remove directory and all contents

// Copy and move files (legacy functions)
os.copy("source.txt", "destination.txt")
os.move("old_location.txt", "new_location.txt")

// Create temporary files
tempFile = os.createTemp("", "prefix_*.tmp")
print("Created temp file:", tempFile)

// Check if two files are the same
same = os.sameFile("file1.txt", "file2.txt")
```

## Directory Operations

### Basic Directory Operations

```js
// Get current working directory
currentDir = os.getwd()
print("Current directory:", currentDir)

// Change directory
os.changeDir("/path/to/directory")

// Create directory
os.makeDir("new_folder")

// Create directory with all parent directories
os.mkdirAll("path/to/nested/directory")

// Remove directory (must be empty)
os.removeDir("empty_folder")

// Remove directory and all contents recursively
os.removeAll("folder_with_contents")
```

### Directory Listing

```js
// Get directory contents as comma-separated string
files = os.listDir(".")
print(files)

// Get files only (excluding directories) as array
filesOnly = os.listFiles(".")
print(filesOnly)  // ["file1.txt", "file2.txt", ...]

// Get detailed directory information
dirContents = os.readDir(".")
for (item in dirContents) {
    print("Name:", item["name"])
    print("Is Directory:", item["isDir"])
    print("Size:", item["size"])
}
```

### Temporary Directories

```js
// Create temporary directory
tempDir = os.mkdirTemp("", "myapp_")
print("Created temp dir:", tempDir)
```

## File Permissions and Ownership

```js
// Change file permissions (Unix-style mode)
os.chmod("file.txt", 0o644)  // Read/write for owner, read for group/others

// Change file ownership (Unix only)
os.chown("file.txt", 1000, 100)  // uid=1000, gid=100

// Change ownership without following symlinks
os.lchown("symlink.txt", 1000, 100)

// Change file access and modification times
os.chtimes("file.txt", accessTime, modTime)
```

## File System Links

```js
// Create hard link
os.link("original.txt", "hardlink.txt")

// Create symbolic link
os.symlink("target.txt", "symlink.txt")

// Read symbolic link target
target = os.readlink("symlink.txt")
print("Link points to:", target)
```

## System Information

```js
// Get system information
cpuCount = os.cpuCount()
hostname = os.hostname()
pageSize = os.getpagesize()

print("CPUs:", cpuCount)
print("Hostname:", hostname)
print("Page size:", pageSize)

// Get executable path
execPath = os.executable()
print("Current executable:", execPath)

// Check path separator
isSeparator = os.isPathSeparator("/")  // true on Unix, false on Windows for "\"
```

## Error Checking

The OS module provides functions to check specific types of errors:

```js
// Check if error indicates file exists
if (os.isExist(errorMsg)) {
    print("File already exists")
}

// Check if error indicates file doesn't exist
if (os.isNotExist(errorMsg)) {
    print("File not found")
}

// Check if error is permission-related
if (os.isPermission(errorMsg)) {
    print("Permission denied")
}

// Check if error is timeout-related
if (os.isTimeout(errorMsg)) {
    print("Operation timed out")
}
```

## User Directories

```js
// Get user-specific directories
homeDir = os.userHomeDir()
cacheDir = os.userCacheDir()
configDir = os.userConfigDir()
tempDir = os.tempDir()

print("Home:", homeDir)
print("Cache:", cacheDir)
print("Config:", configDir)
print("Temp:", tempDir)
```

## Legacy Functions

These functions are maintained for backward compatibility:

```js
// Legacy directory functions
home = os.homedir()        // Use os.userHomeDir() instead
temp = os.tmpdir()         // Use os.tempDir() instead

// Legacy file operations
os.copy("src.txt", "dst.txt")    // Copy file
os.move("old.txt", "new.txt")    // Move/rename file
os.deleteFile("file.txt")        // Use os.remove() instead

// Legacy system info
currentDir = os.currentDir()     // Use os.getwd() instead
```

## Complete Example

Here's a comprehensive example demonstrating various OS module functions:

```js
const os = import("os")

// Process and system info
println("=== System Information ===")
println("Process ID:", os.getpid())
println("CPU Count:", os.cpuCount())
println("Hostname:", os.hostname())
println("Home Directory:", os.userHomeDir())

// Environment variables
println("\n=== Environment ===")
os.setEnv("MYVAR", "hello world")
println("MYVAR:", os.getEnv("MYVAR"))
println("PATH exists:", os.lookupEnv("PATH")["exists"])

// File operations
println("\n=== File Operations ===")
os.writeFile("test.txt", "Hello, Vint!")
println("File exists:", os.fileExists("test.txt"))

fileInfo = os.stat("test.txt")
println("File size:", fileInfo["size"])
println("Is directory:", fileInfo["isDir"])

// Cleanup
os.remove("test.txt")
println("File removed")
```

The **Vint** OS module provides comprehensive system-level functionality, enabling powerful file system operations, process management, and environment interaction in your Vint programs.

```

## packages.md

```markdown
# Packages in VintLang

Packages in VintLang provide a powerful way to organize, encapsulate, and reuse your code. They allow you to group related functions, variables, and state into a single, importable unit, similar to modules or libraries in other languages.

---

## Defining a Package

You can define a package using the `package` keyword, followed by the package name and a block of code enclosed in curly braces `{}`.

A single `.vint` file can contain one package definition. The name of the file does not need to match the package name, but it is good practice to keep them related.

**Syntax:**
```js
package MyPackage {
    // ... package members ...
}
```

### Package Members

Inside a package block, you can define:

- **Variables**: To hold the package's state using `let`.
- **Constants**: Immutable values using the `const` keyword.
- **Functions**: To provide the package's functionality.

#### Constants in Packages

VintLang supports package-level constants using the `const` keyword. Constants are immutable and are often used for configuration values, version numbers, or other fixed data:

```js
package Config {
    const VERSION = "1.2.3"
    const MAX_CONNECTIONS = 100
    const API_BASE_URL = "https://api.example.com"
    
    let getConfig = func() {
        return {
            "version": VERSION,
            "max_conn": MAX_CONNECTIONS,
            "api_url": API_BASE_URL
        }
    }
}
```

#### Public vs Private Members

VintLang supports access control for package members using a naming convention:

- **Public members**: Names that do NOT start with an underscore `_` are accessible from outside the package.
- **Private members**: Names that start with an underscore `_` are only accessible within the package itself.

```js
package MyPackage {
    // Public members (accessible from outside)
    let publicVariable = "I'm accessible"
    const PUBLIC_CONSTANT = 42
    let publicFunction = func() { return "Hello!" }
    
    // Private members (internal use only)
    let _privateVariable = "Internal only"
    const _PRIVATE_KEY = "secret-key-123"
    let _privateFunction = func() { return "Internal helper" }
}
```

Attempting to access private members from outside the package will result in an error:

```js
import "MyPackage"

print(MyPackage.publicVariable)  // ✅ Works
print(MyPackage._privateVariable) // ❌ Error: cannot access private property

---

## The Automatic `init` Function

VintLang's package system includes a special feature for initialization. If you define a function named `init` inside your package, the Vint interpreter will **automatically execute it** when the package is first loaded.

This is useful for setting up initial state, connecting to services, or performing any other setup work the package needs before it can be used.

**Example:**
```js
package Counter {
    let count = 0

    // This function will run automatically
    let init = func() {
        print("Counter package has been initialized!")
        @.count = 100 // Set initial state
    }

    let getCount = func() {
        return @.count
    }
}
```

---

## The `@` Operator: Self-Reference

Inside a package, you may need to refer to the package's own members or state. VintLang provides the special `@` operator for this purpose. The `@` operator is a reference to the package's own scope.

This is similar to `this` or `self` in other object-oriented languages.

You use it with dot notation to access other members within the same package.

**Example:**

```js
package Greeter {
    let greeting = "Hello"

    let setGreeting = func(newGreeting) {
        // Use @ to access the 'greeting' variable
        @.greeting = newGreeting
    }

    let sayHello = func(name) {
        // Use @ to access the 'greeting' variable
        print(@.greeting + ", " + name + "!")
    }
}
```

Using `@` is necessary to distinguish between a package-level variable and a local variable with the same name.

---

## Importing and Using Packages

To use a package, you import the file that contains its definition. The package object is then assigned to a variable with the same name as the package.

1. **Create your package file** (e.g., `utils.vint`).
2. **Import it in another file** (e.g., `main.vint`).
3. **Access its members** using dot notation.

If `utils.vint` contains `package utils { ... }`, you would use it like this:

```js
// main.vint

// Import the file containing the package
import "utils"

// Now you can use the 'utils' package
utils.doSomething()
```

---

## Complete Examples

For comprehensive, runnable examples that demonstrate all package features including constants, private members, and initialization, see the files in the `examples/packages_example/` directory:

- **`enhanced_test.vint`**: Showcases constants, private members, auto-initialization, and complex package functionality.
- **`greeter_pkg.vint`**: Simple package with state management and initialization.
- **`enhanced_system_test.vint`**: Demonstrates how to use packages with private member protection.

### Key Features Summary

✅ **Package-level constants** with `const` keyword
✅ **Private member protection** using underscore `_` prefix
✅ **Auto-initialization** with `init()` functions  
✅ **State management** with the `@` operator
✅ **Comprehensive access control** for variables, constants, and functions

```js
// Complete example demonstrating all features
package EnhancedExample {
    // Public constants
    const VERSION = "2.0.0"
    const MAX_ITEMS = 100
    
    // Private constants  
    const _SECRET_KEY = "internal-key-123"
    
    // Public variables
    let counter = 0
    
    // Private variables
    let _internalState = "hidden"
    
    // Auto-initialization
    let init = func() {
        print("Package initialized! Version:", VERSION)
        @.counter = 10
    }
    
    // Public functions
    let increment = func() {
        @.counter = @.counter + 1
        return @.counter
    }
    
    // Private functions
    let _validate = func(value) {
        return value != null && value > 0
    }
    
    let processValue = func(value) {
        if (!_validate(value)) {
            return "Invalid value"
        }
        return "Processed: " + string(value)
    }
}
```

```

## path.md

```markdown
# Path Module in VintLang

The `path` module provides functions for working with file system paths.

## Functions

### `path.join([...paths])`
Joins one or more path components intelligently.

```js
p = path.join("/users", "alice", "docs", "file.txt")
print(p) // Outputs: /users/alice/docs/file.txt
```

### `path.basename(path)`

Returns the last portion of a path.

```js
p = path.basename("/users/alice/docs/file.txt")
print(p) // Outputs: file.txt
```

### `path.dirname(path)`

Returns the directory name of a path.

```js
p = path.dirname("/users/alice/docs/file.txt")
print(p) // Outputs: /users/alice/docs
```

### `path.ext(path)`

Returns the file extension of the path.

```js
p = path.ext("/users/alice/docs/file.txt")
print(p) // Outputs: .txt
```

### `path.isAbs(path)`

Returns `true` if the path is absolute.

```js
p = path.isAbs("/users/alice/docs/file.txt")
print(p) // Outputs: true

p2 = path.isAbs("docs/file.txt")
print(p2) // Outputs: false
```

```

## pointers.md

```markdown
 # Pointers in VintLang

VintLang now supports basic pointer operations, allowing you to reference and dereference values in your programs. This feature enables more advanced data manipulation and can be useful for certain algorithms and data structures.

## Syntax

- **Address-of:** Use `&` to get a pointer to a value.
- **Dereference:** Use `*` to access the value pointed to by a pointer.

## Usage

### Creating a Pointer
```js
let x = 42
let p = &x  # p is now a pointer to the value of x
```

### Dereferencing a Pointer

```js
print(*p)  # prints 42
```

### Printing a Pointer

```js
print(p)  # prints something like Pointer(42) or Pointer(addr=0x..., value=42)
```

## Limitations

- **Pointers in VintLang are pointers to values, not to variables.**
  - If you change the value of `x` after creating `p = &x`, the pointer `p` will still point to the original value, not the updated value of `x`.
- You cannot assign through a pointer (e.g., `*p = 100` is not supported).
- Pointers to literals (e.g., `let p = &42`) are allowed, but they are just pointers to the value at the time of creation.

## Example

```js
let x = 10
let p = &x
print(p)    # Pointer(10)
print(*p)   # 10
x = 20
print(*p)   # Still 10, because p points to the original value
```

## Error Handling

- Dereferencing a non-pointer or a nil pointer will result in a runtime error.

## Summary

- Use `&` to create pointers to values.
- Use `*` to dereference pointers.
- Pointers are useful for referencing values, but do not provide full variable reference semantics.

```

## postgres.md

```markdown
# PostgreSQL Module in VintLang

The `postgres` module in **VintLang** allows you to interact with PostgreSQL databases. This guide will walk you through connecting, executing queries, and fetching data.

## Connecting to a PostgreSQL Database

To connect to a PostgreSQL database, use `postgres.open()`. The connection string should be in the format: `"user=youruser password=yourpassword dbname=yourdbname sslmode=disable"`.

```js
conn = postgres.open("user=postgres password=password dbname=testdb sslmode=disable")
```

## Closing the Connection

Make sure to close the connection with `postgres.close()` when you are finished.

```js
postgres.close(conn)
```

## Executing Queries

Use `postgres.execute()` for `INSERT`, `UPDATE`, `DELETE`, and other statements that do not return data. PostgreSQL uses `$1`, `$2`, etc., as placeholders.

```js
// Inserting data with placeholders
insert_query = "INSERT INTO users (name, age) VALUES ($1, $2)"
postgres.execute(conn, insert_query, "Alice", 30)
```

## Fetching Data

### Fetch All Rows

To retrieve all rows from a query, use `postgres.fetchAll()`.

```js
users = postgres.fetchAll(conn, "SELECT * FROM users")
print(users)
```

### Fetch a Single Row

To retrieve just one row, use `postgres.fetchOne()`.

```js
user = postgres.fetchOne(conn, "SELECT * FROM users WHERE id = $1", 1)
print(user)
```

## Full Example

Here is a complete example demonstrating the use of the `postgres` module:

```js
import postgres

// Replace with your actual credentials
conn_str = "user=postgres password=password dbname=testdb sslmode=disable"
conn = postgres.open(conn_str)

if conn.type() == "ERROR" {
    print("Error connecting to PostgreSQL:", conn.message())
} else {
    print("Successfully connected to PostgreSQL")

    // Create a table
    create_query = "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(255), age INT)"
    postgres.execute(conn, create_query)

    // Insert data
    postgres.execute(conn, "INSERT INTO users (name, age) VALUES ($1, $2)", "Bob", 35)

    // Fetch and print data
    users = postgres.fetchAll(conn, "SELECT * FROM users")
    print("All users:", users)

    // Close the connection
    postgres.close(conn)
    print("Connection closed")
} 
```

## random.md

```markdown
# Random Module in VintLang

The `random` module provides functions for generating random numbers and data.

## Functions

### `random.int(min, max)`
Returns a random integer in the range `[min, max]`, inclusive.

```js
num = random.int(1, 100)
print(num) // Outputs a random number between 1 and 100
```

### `random.float()`

Returns a random float in the range `[0.0, 1.0)`.

```js
f = random.float()
print(f)
```

### `random.string(length)`

Returns a random string of a given length.

```js
s = random.string(12)
print(s) // Outputs a random 12-character string
```

### `random.choice(array)`

Returns a random element from an array.

```js
items = ["apple", "banana", "cherry"]
item = random.choice(items)
print(item) // Outputs one of the fruits
```

```

## range.md

```markdown
# Range Function (`range`)

The `range` function in Vint generates a sequence of numbers and is commonly used in loops or for creating arrays of sequential values.

---

## Syntax

```js
range(end)
range(start, end)
range(start, end, step)
```

---

## Parameters

- **`end`**: The upper limit of the sequence (exclusive).
- **`start`** (optional): The starting value of the sequence. Default is `0`.
- **`step`** (optional): The increment or decrement between each number in the sequence. Default is `1`.

---

## Return Value

The function returns an array of integers.

---

## Examples

### Basic Usage

```js
// Generate numbers from 0 to 4
for i in range(5) {
    print(i)
}
// Output: 0 1 2 3 4
```

### Specifying a Start and End

```js
// Generate numbers from 1 to 9
for i in range(1, 10) {
    print(i)
}
// Output: 1 2 3 4 5 6 7 8 9
```

### Using a Step Value

```js
// Generate even numbers from 0 to 8
for i in range(0, 10, 2) {
    print(i)
}
// Output: 0 2 4 6 8
```

### Generating a Reverse Sequence

```js
// Generate numbers in reverse order
for i in range(10, 0, -1) {
    print(i)
}
// Output: 10 9 8 7 6 5 4 3 2 1
```

---

## Notes

1. **Exclusive End**: The `end` value is not included in the sequence; the range stops before reaching it.
2. **Negative Steps**: If a negative `step` is provided, ensure `start` is greater than `end` to create a valid reverse sequence.
3. **Non-Zero Step**: The `step` value cannot be `0`, as it would result in an infinite loop or an error.

---

```

## redis.md

```markdown
# Redis Module

The Redis module provides comprehensive Redis database functionality for VintLang, allowing you to interact with Redis servers for caching, data storage, and more.

## Connection Management

### `redis.connect(address, [password], [db])`

Establishes a connection to a Redis server.

- `address`: Redis server address (e.g., "localhost:6379")
- `password`: Optional password for authentication
- `db`: Optional database number (default: 0)

**Returns**: Connection object

**Example**:

```js
conn = redis.connect("localhost:6379");
auth_conn = redis.connect("localhost:6379", "mypassword", 1);
```

### `redis.close(connection)`

Closes the Redis connection.

**Example**:

```js
redis.close(conn);
```

### `redis.ping(connection)`

Tests the connection to Redis.

**Returns**: "PONG" if successful

**Example**:

```js
result = redis.ping(conn); // Returns "PONG"
```

## String Operations

### `redis.set(connection, key, value)`

Sets a string value.

**Example**:

```js
redis.set(conn, "mykey", "myvalue");
```

### `redis.get(connection, key)`

Gets a string value.

**Returns**: String value or null if key doesn't exist

**Example**:

```js
value = redis.get(conn, "mykey");
```

### `redis.setex(connection, key, value, seconds)`

Sets a string value with expiration time.

**Example**:

```js
redis.setex(conn, "session:123", "userdata", 3600); // Expires in 1 hour
```

### `redis.mset(connection, key1, value1, key2, value2, ...)`

Sets multiple key-value pairs.

**Example**:

```js
redis.mset(conn, "key1", "value1", "key2", "value2");
```

### `redis.mget(connection, key1, key2, ...)`

Gets multiple values.

**Returns**: Array of values (null for non-existent keys)

**Example**:

```js
values = redis.mget(conn, "key1", "key2");
```

## Numeric Operations

### `redis.incr(connection, key)`

Increments the integer value of a key by 1.

**Example**:

```js
counter = redis.incr(conn, "visits"); // Increments and returns new value
```

### `redis.decr(connection, key)`

Decrements the integer value of a key by 1.

### `redis.incrby(connection, key, increment)`

Increments the integer value of a key by the given amount.

### `redis.decrby(connection, key, decrement)`

Decrements the integer value of a key by the given amount.

## Key Operations

### `redis.exists(connection, key)`

Checks if a key exists.

**Returns**: true/false

### `redis.del(connection, key1, [key2, ...])`

Deletes one or more keys.

**Returns**: Number of keys that were deleted

### `redis.expire(connection, key, seconds)`

Sets expiration time for a key.

**Returns**: true if timeout was set, false if key doesn't exist

### `redis.ttl(connection, key)`

Gets the time to live of a key in seconds.

**Returns**: TTL in seconds (-1 if no timeout, -2 if key doesn't exist)

### `redis.keys(connection, pattern)`

Returns all keys matching a pattern.

**Example**:

```js
keys = redis.keys(conn, "user:*"); // All keys starting with "user:"
```

## Hash Operations

### `redis.hset(connection, key, field, value)`

Sets a field in a hash.

### `redis.hget(connection, key, field)`

Gets a field from a hash.

### `redis.hgetall(connection, key)`

Gets all fields and values from a hash.

**Returns**: Dictionary with field-value pairs

### `redis.hdel(connection, key, field1, [field2, ...])`

Deletes one or more hash fields.

### `redis.hexists(connection, key, field)`

Determines if a hash field exists.

### `redis.hkeys(connection, key)`

Gets all field names in a hash.

### `redis.hvals(connection, key)`

Gets all values in a hash.

## List Operations

### `redis.lpush(connection, key, value1, [value2, ...])`

Prepends one or more values to a list.

### `redis.rpush(connection, key, value1, [value2, ...])`

Appends one or more values to a list.

### `redis.lpop(connection, key)`

Removes and returns the first element of a list.

### `redis.rpop(connection, key)`

Removes and returns the last element of a list.

### `redis.llen(connection, key)`

Returns the length of a list.

### `redis.lrange(connection, key, start, stop)`

Returns a range of elements from a list.

**Example**:

```js
elements = redis.lrange(conn, "mylist", 0, -1); // All elements
```

## Set Operations

### `redis.sadd(connection, key, member1, [member2, ...])`

Adds one or more members to a set.

### `redis.srem(connection, key, member1, [member2, ...])`

Removes one or more members from a set.

### `redis.smembers(connection, key)`

Returns all members of a set.

### `redis.scard(connection, key)`

Returns the number of members in a set.

### `redis.sismember(connection, key, member)`

Determines if a member is in a set.

## Sorted Set Operations

### `redis.zadd(connection, key, score1, member1, [score2, member2, ...])`

Adds one or more members to a sorted set with scores.

**Example**:

```js
redis.zadd(conn, "leaderboard", 100, "player1", 85, "player2");
```

### `redis.zrem(connection, key, member1, [member2, ...])`

Removes one or more members from a sorted set.

### `redis.zrange(connection, key, start, stop)`

Returns a range of members in a sorted set by index.

### `redis.zcard(connection, key)`

Returns the number of members in a sorted set.

### `redis.zscore(connection, key, member)`

Returns the score of a member in a sorted set.

## Usage Example

```js
// Connect to Redis
conn = redis.connect("localhost:6379");

// Basic string operations
redis.set(conn, "greeting", "Hello, World!");
message = redis.get(conn, "greeting");
print(message); // Output: Hello, World!

// Working with hashes
redis.hset(conn, "user:1", "name", "John Doe");
redis.hset(conn, "user:1", "email", "john@example.com");
user = redis.hgetall(conn, "user:1");
print(user); // Output: {"name": "John Doe", "email": "john@example.com"}

// Working with lists
redis.rpush(conn, "tasks", "task1", "task2", "task3");
task = redis.lpop(conn, "tasks");
print(task); // Output: task1

// Close connection
redis.close(conn);
```

```

## reflect.md

```markdown
# Reflect Module

The `reflect` module provides runtime type inspection and reflection utilities for VintLang. It allows you to examine the type and structure of values, check for null, and determine if a value is an array, object, or function.

## Importing

```js
import reflect
```

## Functions

### reflect.typeOf(value)

Returns the type name of the given value as a string.

- **Arguments:**
  - `value`: Any value
- **Returns:** String (e.g., "STRING", "ARRAY", "DICT", "NULL", "FUNCTION", etc.)
- **Example:**

  ```js
  reflect.typeOf("hello")        // "STRING"
  reflect.typeOf([1,2,3])        // "ARRAY"
  reflect.typeOf({"a": 1})      // "DICT"
  reflect.typeOf(null)           // "NULL"
  reflect.typeOf(func() {})      // "FUNCTION"
  ```

### reflect.valueOf(value)

Returns the raw value passed in (identity function).

- **Arguments:**
  - `value`: Any value
- **Returns:** The same value
- **Example:**

  ```js
  reflect.valueOf(42); // 42
  reflect.valueOf("foo"); // "foo"
  ```

### reflect.isNil(value)

Checks if the value is `null`.

- **Arguments:**
  - `value`: Any value
- **Returns:** Boolean
- **Example:**

  ```js
  reflect.isNil(null); // true
  reflect.isNil(123); // false
  ```

### reflect.isArray(value)

Checks if the value is an array.

- **Arguments:**
  - `value`: Any value
- **Returns:** Boolean
- **Example:**

  ```js
  reflect.isArray([1, 2, 3]); // true
  reflect.isArray("not array"); // false
  ```

### reflect.isObject(value)

Checks if the value is a dictionary/object.

- **Arguments:**
  - `value`: Any value
- **Returns:** Boolean
- **Example:**

  ```js
  reflect.isObject({ a: 1 }); // true
  reflect.isObject([1, 2, 3]); // false
  ```

### reflect.isFunction(value)

Checks if the value is a function.

- **Arguments:**
  - `value`: Any value
- **Returns:** Boolean
- **Example:**

  ```js
  let f = func(x) { x * 2 }
  reflect.isFunction(f)          // true
  reflect.isFunction(123)        // false
  ```

## Example Usage

See `examples/reflect.vint` for a full demonstration of all reflect module functions.

```

## regex.md

```markdown
# Regex Module in Vint

The `regex` module in Vint provides powerful functionality to perform regular expression operations like matching patterns, replacing strings, and splitting strings. Regular expressions are useful for text processing and pattern matching.

---

## Importing the Regex Module

To use the Regex module, import it as follows:

```js
import regex
```

---

## Functions and Examples

### 1. Matching a Pattern with `match`

The `match` function checks if a string matches a specified pattern. It returns `true` if the string matches the pattern and `false` if it does not.

**Syntax**:

```js
match(pattern, string)
```

- `pattern`: The regular expression pattern.
- `string`: The string to match against the pattern.

**Example**:

```js
import regex

result = regex.match("^Hello", "Hello World")
print(result)  // Expected output: true
```

In this case, the string starts with `"Hello"`, so it matches the pattern.

---

### 2. Using `match` to Check Non-Matches

You can use `match` to check if a string does *not* match a given pattern. If the pattern is not found at the beginning of the string, it will return `false`.

**Example**:

```js
import regex

result = regex.match("^World", "Hello World")
print(result)  // Expected output: false
```

Since the string does not start with `"World"`, the result is `false`.

---

### 3. Replacing Part of a String with `replaceString`

The `replaceString` function replaces parts of a string that match a pattern with a new value.

**Syntax**:

```js
replaceString(pattern, replacement, string)
```

- `pattern`: The regular expression pattern.
- `replacement`: The string to replace the matched part with.
- `string`: The input string.

**Example**:

```js
import regex

newString = regex.replaceString("World", "VintLang", "Hello World")
print(newString)  // Expected output: "Hello VintLang"
```

In this case, `"World"` is replaced by `"VintLang"`, resulting in `"Hello VintLang"`.

---

### 4. Splitting a String with `splitString`

The `splitString` function splits a string into a list of substrings based on a regular expression pattern.

**Syntax**:

```js
splitString(pattern, string)
```

- `pattern`: The regular expression pattern used as a delimiter.
- `string`: The string to be split.

**Example**:

```js
import regex

words = regex.splitString("\\s+", "Hello World VintLang")
print(words)  // Expected output: ["Hello", "World", "VintLang"]
```

Here, `\\s+` matches one or more whitespace characters, so the string is split into words.

---

### 5. Splitting a String by a Comma

You can also split a string by a specific delimiter, such as a comma, using `splitString`.

**Example**:

```js
import regex

csv = regex.splitString(",", "apple,banana,orange")
print(csv)  // Expected output: ["apple", "banana", "orange"]
```

The string `"apple,banana,orange"` is split at each comma.

---

### 6. Matching a Complex Pattern

You can match more complex patterns, such as an email address, using `match` with a regex pattern.

**Example**:

```js
import regex

emailMatch = regex.match("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", "test@example.com")
print(emailMatch)  // Expected output: true
```

This pattern matches valid email addresses, so `"test@example.com"` matches successfully.

---

### 7. Replacing Digits in a String

You can use `replaceString` to replace parts of a string that match a pattern, such as replacing digits with asterisks.

**Example**:

```js
import regex

maskedString = regex.replaceString("\\d", "*", "My phone number is 123456789")
print(maskedString)  // Expected output: "My phone number is *********"
```

Here, `\\d` matches any digit, so the digits in the phone number are replaced with asterisks.

---

## Summary of Functions

| Function           | Description                                             | Example Output                             |
|--------------------|---------------------------------------------------------|--------------------------------------------|
| `match(pattern, string)`  | Checks if the string matches the given pattern.         | `true` or `false`                          |
| `replaceString(pattern, replacement, string)`  | Replaces parts of a string matching a pattern with a new value.  | A modified string                          |
| `splitString(pattern, string)` | Splits a string into a list of substrings based on the pattern.   | List of substrings                         |

---

The `regex` module is extremely useful for text manipulation, pattern matching, and string processing tasks in Vint. Whether you need to validate input, split strings, or replace parts of a string, the regex module provides powerful tools to handle these tasks efficiently.

```

## schedule.md

```markdown
# Schedule Module

The `schedule` module provides powerful scheduling capabilities for VintLang, similar to Go's ticker and NestJS's schedule decorators. It allows you to execute functions at regular intervals or at specific times using cron-like expressions.

## Functions

### `ticker(intervalSeconds, callback)`

Creates a ticker that executes a callback function at regular intervals.

**Parameters:**
- `intervalSeconds` (integer): The interval in seconds between executions
- `callback` (function): The function to execute at each interval

**Returns:** A ticker object that can be stopped with `stopTicker()`

**Example:**
```javascript
import schedule

// Execute every 5 seconds
let ticker = schedule.ticker(5, func() {
    print("Tick at", time.now())
})

// Stop the ticker after some time
time.sleep(20)
schedule.stopTicker(ticker)
```

### `stopTicker(tickerObj)`

Stops a running ticker.

**Parameters:**

- `tickerObj`: The ticker object returned by `ticker()`

**Returns:** Boolean indicating if the ticker was successfully stopped

### `schedule(cronExpr, callback)`

Schedules a function to execute at specific times using cron-like expressions.

**Parameters:**

- `cronExpr` (string): A cron expression in the format "second minute hour day month weekday"
- `callback` (function): The function to execute when the schedule triggers

**Returns:** A schedule object that can be stopped with `stopSchedule()`

**Cron Expression Format:**

```
second minute hour day month weekday
  |      |     |    |    |      |
  |      |     |    |    |      +-- Day of week (0-6, 0=Sunday)
  |      |     |    |    +--------- Month (1-12)
  |      |     |    +-------------- Day of month (1-31)
  |      |     +------------------- Hour (0-23)
  |      +------------------------- Minute (0-59)
  +-------------------------------- Second (0-59)
```

Use `*` for wildcards (any value).

**Examples:**

```javascript
import schedule

// Every minute at second 0
schedule.schedule("0 * * * * *", func() {
    print("Top of the minute!")
})

// Daily at 9:30 AM
schedule.schedule("0 30 9 * * *", func() {
    print("Good morning!")
})

// Every 30 seconds
schedule.schedule("*/30 * * * * *", func() {
    print("Every 30 seconds")
})

// Every Friday at 5:00 PM
schedule.schedule("0 0 17 * * 5", func() {
    print("TGIF!")
})
```

### `stopSchedule(scheduleObj)`

Stops a running scheduled task.

**Parameters:**

- `scheduleObj`: The schedule object returned by `schedule()`

**Returns:** Boolean indicating if the schedule was successfully stopped

## Helper Functions

The module provides convenient helper functions for common scheduling patterns:

### `everySecond(callback)`

Executes a callback every second. Equivalent to `ticker(1, callback)`.

**Example:**

```javascript
let job = schedule.everySecond(func() {
    print("Ping!")
})
```

### `everyMinute(callback)`

Executes a callback every minute at second 0. Equivalent to `schedule("0 * * * * *", callback)`.

**Example:**

```javascript
let job = schedule.everyMinute(func() {
    print("Another minute passed")
})
```

### `everyHour(callback)`

Executes a callback every hour at minute 0, second 0. Equivalent to `schedule("0 0 * * * *", callback)`.

**Example:**

```javascript
let job = schedule.everyHour(func() {
    print("It's a new hour!")
})
```

### `daily(hour, minute, callback)`

Executes a callback daily at the specified time.

**Parameters:**

- `hour` (integer): Hour of the day (0-23)
- `minute` (integer): Minute of the hour (0-59)
- `callback` (function): The function to execute

**Example:**

```javascript
// Daily reminder at 2:30 PM
let job = schedule.daily(14, 30, func() {
    print("Time for afternoon coffee!")
})
```

## Complete Example

```javascript
import schedule
import time

print("Starting scheduling demo...")

// Ticker example - every 3 seconds
let ticker = schedule.ticker(3, func() {
    print("[TICKER] Current time:", time.now())
})

// Schedule example - every 10 seconds
let job1 = schedule.schedule("*/10 * * * * *", func() {
    print("[SCHEDULE] Every 10 seconds")
})

// Helper function example
let job2 = schedule.everySecond(func() {
    print("[HELPER] Ping!")
})

// Daily example (would execute at specified time)
let job3 = schedule.daily(9, 0, func() {
    print("[DAILY] Good morning! Time to start the day.")
})

// Let everything run for 30 seconds
print("Running for 30 seconds...")
time.sleep(30)

// Clean up
print("Stopping all jobs...")
schedule.stopTicker(ticker)
schedule.stopSchedule(job1)
schedule.stopTicker(job2)
schedule.stopSchedule(job3)

print("Demo completed!")
```

## Notes

- All scheduling is done using Go's `time.Ticker` and `time.Timer` internally
- Callbacks currently log execution messages to demonstrate functionality
- The module uses goroutines for non-blocking execution
- Proper cleanup is important - always stop tickers and schedules when done
- Cron expressions support basic patterns; advanced features like ranges (1-5) or lists (1,3,5) are not yet implemented

## Error Handling

The module provides comprehensive error messages for common mistakes:

- Invalid argument types
- Invalid time ranges (e.g., hour > 23, minute > 59)
- Malformed cron expressions
- Negative intervals for tickers

All errors include usage examples to help with correct implementation.

```

## shell.md

```markdown
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

```

## sqlite.md

```markdown
# SQLite Module in VintLang

In **VintLang**, the `sqlite` module allows interaction with SQLite databases. You can open a database, execute queries, fetch data, and manage tables. This guide covers basic database operations.

## Open a Database

Use `sqlite.open()` to open a connection to an SQLite database.

```js
db = sqlite.open("example.db")
```

## Close a Database

To close the database connection, use `sqlite.close()`.

```js
sqlite.close(db)
```

## Execute a Query

You can execute `INSERT`, `UPDATE`, `DELETE`, and other queries using `sqlite.execute()`.

### Insert Data

```js
sqlite.execute(db, "INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 25)
```

### Update Data

```js
sqlite.execute(db, "UPDATE users SET age = ? WHERE name = ?", 26, "Alice")
```

## Fetch Data

Use `sqlite.fetchAll()` to retrieve all rows from a query.

```js
users = sqlite.fetchAll(db, "SELECT * FROM users")
print(users)
```

You can also fetch a single row with `sqlite.fetchOne()`.

```js
first_user = sqlite.fetchOne(db, "SELECT * FROM users LIMIT 1")
print(first_user)
```

## Create a Table

To create a new table, use `sqlite.createTable()`.

```js
sqlite.createTable(db, "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
```

## Drop a Table

Use `sqlite.dropTable()` to delete a table from the database.

```js
sqlite.dropTable(db, "users")
```

## Example Usage

```js
import sqlite

// Open a database
db = sqlite.open("example.db")

// Create a table
sqlite.createTable(db, "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")

// Insert data
sqlite.execute(db, "INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 25)

// Fetch all rows
users = sqlite.fetchAll(db, "SELECT * FROM users")
print(users)

// Fetch a single row
first_user = sqlite.fetchOne(db, "SELECT * FROM users LIMIT 1")
print(first_user)

// Close the database connection
sqlite.close(db)
```

By using the **VintLang** `sqlite` module, you can easily manage SQLite databases in your programs.

```

## strings.md

```markdown
# Strings in Vint

Strings are a sequence of characters used to represent text in the Vint programming language. Here’s a detailed explanation of how to work with strings, including syntax, manipulation, and useful built-in methods.

## Basic Syntax

In Vint, strings can be enclosed in either single quotes (`''`) or double quotes (`""`):

```js
print("Hello")  // Output: Hello

let name = 'Tachera'

print("Hello", name)  // Output: Hello Tachera
```

## Concatenating Strings

Strings can be concatenated using the `+` operator:

```js
let greeting = "Hello" + " " + "World"
print(greeting)  // Output: Hello World

let message = "Hello"
message += " World"
// Output: Hello World
```

You can also repeat a string a specific number of times using the `*` operator:

```js
print("Hello " * 3)  // Output: Hello Hello Hello

let repeated = "World"
repeated *= 2
// Output: WorldWorld
```

## Looping Over a String

You can loop through each character of a string using the `for` keyword:

```js
let name = "Avicenna"

for char in name {
    print(char)
}
// Output:
// A
// v
// i
// c
// e
// n
// n
// a
```

You can also loop through the string using its index and character:

```js
for i, char in name {
    print(i, "=>", char)
}
// Output:
// 0 => A
// 1 => v
// 2 => i
// 3 => c
// 4 => e
// 5 => n
// 6 => n
// 7 => a
```

## Comparing Strings

You can compare two strings using the `==` operator:

```js
let a = "Vint"
print(a == "Vint")  // Output: true
print(a == "vint")  // Output: false
```

## String Methods

### Length of a String (`length`)

You can find the length of a string using the `length` method. It does not accept any parameters:

```js
let message = "Vint"
print(message.length())  // Output: 4
```

### Convert to Uppercase (`upper`)

This method converts the string to uppercase:

```js
let text = "vint"
print(text.upper())  // Output: VINT
```

### Convert to Lowercase (`lower`)

This method converts the string to lowercase:

```js
let text = "VINT"
print(text.lower())  // Output: vint
```

### Split a String (`split`)

The `split` method splits a string into an array based on a specified delimiter. If no delimiter is provided, it splits by whitespace.

Example without a delimiter:

```js
let sentence = "Vint programming language"
let words = sentence.split()
print(words)  // Output: ["Vint", "programming", "language"]
```

Example with a delimiter:

```js
let sentence = "Vint,programming,language"
let words = sentence.split(",")
print(words)  // Output: ["Vint", "programming", "language"]
```

### Replace Substrings (`replace`)

You can replace a substring with another string using the `replace` method:

```js
let greeting = "Hello World"
let newGreeting = greeting.replace("World", "Vint")
print(newGreeting)  // Output: Hello Vint
```

### Trim Whitespace (`trim`)

You can remove whitespace from the start and end of a string using the `trim` method:

```js
let text = "  Hello World  "
print(text.trim())  // Output: "Hello World"
```

### Check String Start (`startsWith`)

The `startsWith` method checks if a string starts with a specified prefix:

```js
let text = "Hello World"
print(text.startsWith("Hello"))  // Output: true
print(text.startsWith("World"))  // Output: false
```

### Check String End (`endsWith`)

The `endsWith` method checks if a string ends with a specified suffix:

```js
let text = "Hello World"
print(text.endsWith("World"))   // Output: true
print(text.endsWith("Hello"))   // Output: false
```

### Check String Contains (`includes`)

The `includes` method checks if a string contains a specified substring:

```js
let text = "Hello World"
print(text.includes("lo Wo"))   // Output: true
print(text.includes("xyz"))     // Output: false
```

### Repeat String (`repeat`)

The `repeat` method repeats a string a specified number of times:

```js
let text = "Ha"
print(text.repeat(3))  // Output: "HaHaHa"
print("X".repeat(5))   // Output: "XXXXX"
```

### Capitalize String (`capitalize`)

The `capitalize` method capitalizes the first letter of a string:

```js
let text = "hello world"
print(text.capitalize())  // Output: "Hello world"
```

### Check if Numeric (`isNumeric`)

The `isNumeric` method checks if a string contains only numeric characters:

```js
print("123".isNumeric())      // Output: true
print("12.34".isNumeric())    // Output: true
print("abc".isNumeric())      // Output: false
print("12a".isNumeric())      // Output: false
```

### Check if Alphabetic (`isAlpha`)

The `isAlpha` method checks if a string contains only alphabetic characters:

```js
print("hello".isAlpha())      // Output: true
print("Hello".isAlpha())      // Output: true
print("hello123".isAlpha())   // Output: false
print("hello world".isAlpha()) // Output: false (space is not alphabetic)
```

### Case-Insensitive Compare (`compareIgnoreCase`)

The `compareIgnoreCase` method compares strings ignoring case differences:

```js
let text = "Hello"
print(text.compareIgnoreCase("hello"))   // Output: 0 (equal)
print(text.compareIgnoreCase("apple"))   // Output: 1 (greater)
print(text.compareIgnoreCase("zebra"))   // Output: -1 (less)
```

### Format String (`format`)

The `format` method applies simple formatting to strings using placeholders:

```js
let template = "Hello {0}, you are {1} years old"
print(template.format("John", 25))  // Output: "Hello John, you are 25 years old"

let message = "The result is {0}"
print(message.format(42))  // Output: "The result is 42"
```

### Remove Accents (`removeAccents`)

The `removeAccents` method removes accent characters from a string:

```js
print("café".removeAccents())     // Output: "cafe"
print("naïve".removeAccents())    // Output: "naive"
print("résumé".removeAccents())   // Output: "resume"
```

### Convert to Integer (`toInt`)

The `toInt` method converts a string to an integer:

```js
print("123".toInt())    // Output: 123
print("-456".toInt())   // Output: -456
// Note: Returns an error if the string is not a valid integer
```

```js
let message = "  Hello World  "
print(message.trim())  // Output: Hello World
```

### Get a Substring (`substring`)

You can extract a substring from a string by specifying the starting and ending indices:

```js
let sentence = "Vint programming"
print(sentence.substring(0, 4))  // Output: Vint
```

### Find the Index of a Substring (`indexOf`)

You can find the index of a substring within a string using the `indexOf` method:

```js
let sentence = "Vint programming"
print(sentence.indexOf("programming"))  // Output: 5
```

### Slugify a String (`slug`)

You can convert a string into a URL-friendly format (slug) using the `slug` method:

```js
let title = "Creating a Slug String"
print(title.slug())  // Output: creating-a-slug-string
```

### Checking Substring Presence (`contains`)

Check if a string contains a specific substring:

```js
let name = "Tachera Sasi"
print(name.contains("Sasi"))  // Output: true
```

### Get Character at Index (`charAt`)

Get the character at a specific index:

```js
let word = "Hello"
print(word.charAt(1))  // Output: e
print(word.charAt(10)) // Output: "" (empty string for out of bounds)
```

### Repeat String (`times`)

Repeat a string a specified number of times:

```js
let pattern = "Ha"
print(pattern.times(3))  // Output: HaHaHa
```

### Pad String Start (`padStart`)

Pad the string to a target length from the beginning:

```js
let num = "5"
print(num.padStart(3, "0"))  // Output: 005

let word = "hi"
print(word.padStart(5, "*"))  // Output: ***hi
```

### Pad String End (`padEnd`)

Pad the string to a target length from the end:

```js
let num = "5"
print(num.padEnd(3, "0"))  // Output: 500

let word = "hi"
print(word.padEnd(5, "*"))  // Output: hi***
```

### Check String Start (`startsWith`)

Check if a string starts with a specified prefix:

```js
let message = "Hello World"
print(message.startsWith("Hello"))  // Output: true
print(message.startsWith("World"))  // Output: false
```

### Check String End (`endsWith`)

Check if a string ends with a specified suffix:

```js
let filename = "document.pdf"
print(filename.endsWith(".pdf"))  // Output: true
print(filename.endsWith(".txt"))  // Output: false
```

### Extract Slice (`slice`)

Extract a section of the string:

```js
let text = "Hello World"
print(text.slice(0, 5))   // Output: Hello
print(text.slice(6))      // Output: World
print(text.slice(-5))     // Output: World
```

## Example Usage

Here’s an example of how you might use these string operations in Vint:

```js
import "string"

// Example: Trim whitespace
let trimmed = string.trim("  Hello, World!  ")
print(trimmed)  // Output: "Hello, World!"

// Example: Check if a string contains a substring
let containsResult = string.contains("Hello, World!", "World")
print(containsResult)  // Output: true

// Example: Convert to uppercase
let upperResult = string.toUpper("hello")
print(upperResult)  // Output: "HELLO"

// Example: Convert to lowercase
let lowerResult = string.toLower("HELLO")
print(lowerResult)  // Output: "hello"

// Example: Replace a substring
let replaceResult = string.replace("Hello, World!", "World", "Vint")
print(replaceResult)  // Output: "Hello, Vint!"

// Example: Split a string into parts
let splitResult = string.split("a,b,c,d", ",")
print(splitResult)  // Output: ["a", "b", "c", "d"]

// Example: Join string parts
let joinResult = string.join(["a", "b", "c"], "-")
print(joinResult)  // Output: "a-b-c"

// Example: Get the length of a string
let lengthResult = string.length("Hello")
print(lengthResult)  // Output: 5
```

## Example with Vint Data

Here's an example using Vint-specific strings:

```js
let name = "Tachera Sasi"
let reversed = name.reverse()
print(reversed)  // Output: "isaS arehcaT"

let upperName = name.upper()
print(upperName)  // Output: "TACHERA SASI"

let trimmedName = name.trim("T")
print(trimmedName)  // Output: "achera Sasi"
```

Understanding how to manipulate and work with strings in Vint allows you to efficiently handle text data in your programs.

```

## success.md

```markdown
# Success

The `success` keyword allows you to print success messages at runtime.

## Syntax

```js
success "Your success message here"
```

When the Vint interpreter encounters a `success` statement, it prints a green-colored success message to the console and continues execution. This is useful for indicating when an operation has completed successfully.

### Example

```js
success "Backup completed successfully!"
println("All done.")
```

Running this script will output:

```
[SUCCESS]: Backup completed successfully!
All done.
```

```

## switch.md

```markdown
# Switch Statements in Vint

Switch statements in **Vint** allow you to execute different code blocks based on the value of a given expression. This guide covers the basics of switch statements and their usage.

## Basic Syntax

A switch statement starts with the `switch` keyword, followed by the expression inside parentheses `()`, and all cases enclosed within curly braces `{}`.

Each case uses the keyword `case`, followed by a value to check. Multiple values in a case can be separated by commas `,`. The block of code to execute if the condition is met is placed within curly braces `{}`.

### Example:
```js
let a = 2

switch (a) {
 case 3 {
  print("a is three")
 }
 case 2 {
  print("a is two")
 }
}
```

## Multiple Values in a Case

A single `case` can handle multiple possible values. These values are separated by commas `,`.

### Example

```js
switch (a) {
 case 1, 2, 3 {
  print("a is one, two, or three")
 }
 case 4 {
  print("a is four")
 }
}
```

## Default Case (`default`)

The `default` statement is executed when none of the specified cases match. It is represented by the `default` keyword.

### Example

```js
let z = 20

switch(z) {
 case 10 {
  print("ten")
 }
 case 30 {
  print("thirty")
 }
 default {
  print("twenty")
 }
}
```

## Nested Switch Statements

Switch statements can be nested to handle more complex conditions.

### Example

```js
let x = 1
let y = 2

switch (x) {
 case 1 {
  switch (y) {
   case 2 {
    print("x is one and y is two")
   }
   case 3 {
    print("x is one and y is three")
   }
  }
 }
 case 2 {
  print("x is two")
 }
}
```

## Logical Conditions in Cases

Cases can also be used with logical conditions.

### Example

```js
let isTrue = true
let isFalse = false

switch (isTrue) {
 case true {
  print("isTrue is true")
 }
 case isFalse {
  print("isFalse is true")
 }
 default {
  print("Neither condition is true")
 }
}
```

By mastering switch statements in **Vint**, you can write clean, structured, and efficient code that efficiently handles complex branching logic.

```

## sysinfo.md

```markdown
# Sysinfo Module in Vint

The `sysinfo` module in Vint provides information about the system, such as the operating system and architecture. This can be useful for system diagnostics, logging, or adapting your program based on the system it’s running on.

---

## Importing the Sysinfo Module

To use the Sysinfo module, import it as follows:

```js
import sysinfo
```

---

## Functions and Examples

### 1. Get the Operating System (`os`)

The `os` function returns the name of the operating system on which the Vint program is running.

**Syntax**:

```js
os()
```

- Returns a string representing the operating system (e.g., `"Linux"`, `"Windows"`, `"macOS"`).

**Example**:

```js
import sysinfo

os_name = sysinfo.os()
print("Operating System:", os_name)
// Output: "Operating System: Linux" (or whatever the actual OS is)
```

---

### 2. Get the System Architecture (`arch`)

The `arch` function returns the architecture of the system (e.g., 32-bit or 64-bit).

**Syntax**:

```js
arch()
```

- Returns a string representing the architecture (e.g., `"x86_64"` for 64-bit or `"i386"` for 32-bit).

**Example**:

```js
import sysinfo

architecture = sysinfo.arch()
print("System Architecture:", architecture)
// Output: "System Architecture: x86_64" (or whatever the actual architecture is)
```

---

### 3. Example Combining `os` and `arch`

You can combine both the `os()` and `arch()` functions to display comprehensive system information.

**Example**:

```js
import sysinfo

os_name = sysinfo.os()
architecture = sysinfo.arch()

print("OS:", os_name, "Arch:", architecture)
// Output: "OS: Linux Arch: x86_64" (or whatever the actual OS and architecture are)
```

This will print out both the operating system and the architecture in a single output.

---

### 4. Get Memory Information (`memInfo`)

The `memInfo` function returns detailed information about system memory usage.

**Syntax**:

```js
memInfo()
```

- Returns a dictionary containing memory statistics in GB and usage percentage.

**Example**:

```js
import sysinfo

let memory = sysinfo.memInfo()
print("Total Memory:", memory["total"])
print("Available Memory:", memory["available"])
print("Used Memory:", memory["used"])
print("Free Memory:", memory["free"])
print("Usage Percentage:", memory["percent"] + "%")
// Output example:
// Total Memory: 15.62 GB
// Available Memory: 14.18 GB
// Used Memory: 1.08 GB
// Free Memory: 11.75 GB
// Usage Percentage: 6.89%
```

---

### 5. Get CPU Information (`cpuInfo`)

The `cpuInfo` function returns detailed information about the CPU.

**Syntax**:

```js
cpuInfo()
```

- Returns a dictionary containing CPU model, cores, frequency, and current usage.

**Example**:

```js
import sysinfo

let cpu = sysinfo.cpuInfo()
print("CPU Model:", cpu["model"])
print("CPU Cores:", cpu["cores"])
print("CPU Frequency:", cpu["frequency"])
print("CPU Usage:", cpu["usage"] + "%")
// Output example:
// CPU Model: AMD EPYC 7763 64-Core Processor
// CPU Cores: 1
// CPU Frequency: 3244.00 MHz
// CPU Usage: 33.33%
```

---

### 6. Get Disk Information (`diskInfo`)

The `diskInfo` function returns information about disk usage for the root filesystem.

**Syntax**:

```js
diskInfo()
```

- Returns a dictionary containing disk space information in GB and usage percentage.

**Example**:

```js
import sysinfo

let disk = sysinfo.diskInfo()
print("Total Disk Space:", disk["total"])
print("Used Disk Space:", disk["used"])
print("Free Disk Space:", disk["free"])
print("Disk Usage:", disk["percent"] + "%")
// Output example:
// Total Disk Space: 71.61 GB
// Used Disk Space: 49.81 GB
// Free Disk Space: 21.78 GB
// Disk Usage: 69.58%
```

---

### 7. Get Network Information (`netInfo`)

The `netInfo` function returns information about all network interfaces with addresses.

**Syntax**:

```js
netInfo()
```

- Returns an array of dictionaries, each containing interface name and addresses.

**Example**:

```js
import sysinfo

let interfaces = sysinfo.netInfo()
print("Number of interfaces:", len(interfaces))
for i = 0; i < len(interfaces); i = i + 1 {
    let iface = interfaces[i]
    print("Interface:", iface["name"])
    print("Addresses:", iface["addrs"])
}
// Output example:
// Number of interfaces: 3
// Interface: lo
// Addresses: [127.0.0.1/8, ::1/128]
// Interface: eth0
// Addresses: [10.1.0.75/20, fe80::7eed:8dff:fe4d:223/64]
// Interface: docker0
// Addresses: [172.17.0.1/16]
```

---

### 8. Comprehensive Example

Here's an example that demonstrates all available sysinfo functions:

**Example**:

```js
import sysinfo

print("=== System Information ===")

// Basic system info
print("OS:", sysinfo.os())
print("Architecture:", sysinfo.arch())

// Memory information
let mem = sysinfo.memInfo()
print("Memory - Total:", mem["total"], "Used:", mem["used"], "Usage:", mem["percent"] + "%")

// CPU information
let cpu = sysinfo.cpuInfo()
print("CPU:", cpu["model"], "(" + cpu["cores"] + " cores)")

// Disk information
let disk = sysinfo.diskInfo()
print("Disk - Total:", disk["total"], "Used:", disk["used"], "Usage:", disk["percent"] + "%")

// Network interfaces
let net = sysinfo.netInfo()
print("Network interfaces:", len(net))
```

---

### Summary of Functions

| Function          | Description                                    | Example Output                              |
|-------------------|------------------------------------------------|---------------------------------------------|
| `os`              | Returns the operating system name.             | `"Linux"`, `"Windows"`, `"macOS"`           |
| `arch`            | Returns the system architecture.               | `"x86_64"`, `"i386"`                        |
| `memInfo`         | Returns memory usage information.              | `{"total": "15.62 GB", "used": "1.08 GB", "percent": "6.89"}` |
| `cpuInfo`         | Returns CPU information and usage.             | `{"model": "AMD EPYC", "cores": "1", "usage": "33.33"}` |
| `diskInfo`        | Returns disk usage information.                | `{"total": "71.61 GB", "used": "49.81 GB", "percent": "69.58"}` |
| `netInfo`         | Returns network interface information.         | `[{"name": "eth0", "addrs": ["10.1.0.75/20"]}]` |

---

The `sysinfo` module is useful for comprehensive system monitoring and diagnostics in your Vint programs. It provides detailed information about memory, CPU, disk, and network resources, making it perfect for system administration tools, monitoring applications, and resource-aware programs.

```

## term.md

```markdown

# Terminal UI (`term`) Module

The `term` module provides a rich set of tools for building beautiful and interactive terminal user interfaces (TUIs). It is built on top of the powerful `lipgloss` library and offers a wide range of components, from simple text styling to complex layouts and widgets.

## Basic I/O

### `print(message, [color])`

Prints a message to the terminal.

- `message` (string): The text to print.
- `color` (string, optional): The color to apply to the text.

### `println(message, [color])`

Prints a message to the terminal, followed by a newline.

- `message` (string): The text to print.
- `color` (string, optional): The color to apply to the text.

### `input(prompt)`

Prompts the user for input and returns the entered text.

- `prompt` (string): The message to display to the user.

### `password(prompt)`

Prompts the user for a password without displaying the entered characters.

- `prompt` (string): The message to display to the user.

### `confirm(prompt)`

Asks a yes/no question and returns `true` or `false`.

- `prompt` (string): The question to ask the user.

## Styling & Formatting

### `style(text, options)`

Applies a set of styles to a string.

- `text` (string): The text to style.
- `options` (dict): A dictionary of style options, such as `color`, `background`, `bold`, `italic`, and `underline`.

### `banner(text)`

Creates a large, stylized banner.

- `text` (string): The text to display in the banner.

### `box(text)`

Draws a box around a piece of text.

- `text` (string): The content of the box.

### `badge(text)`

Creates a small, colored badge with text.

- `text` (string): The text for the badge.

### `avatar(text)`

Creates a circular avatar with the first letter of the given text.

- `text` (string): The text to create the avatar from.

## Interactive Components

### `select(options)`

Displays a list of options and allows the user to select one.

- `options` (array): An array of strings representing the choices.

### `checkbox(options)`

Displays a list of options and allows the user to select multiple.

- `options` (array): An array of strings representing the choices.

### `radio(options)`

Displays a list of options and allows the user to select one (similar to `select`).

- `options` (array): An array of strings representing the choices.

### `menu(items)`

Creates a numbered menu from a list of items.

- `items` (array): An array of strings for the menu.

### `form(config)`

Creates a simple form with labeled fields.

- `config` (dict): A dictionary where keys are the field labels.

### `wizard(steps)`

Creates a multi-step wizard or form.

- `steps` (array): An array of strings representing the steps.

## Notifications & Alerts

### `alert(message)`

Displays a prominent alert message.

- `message` (string): The alert message.

### `notify(message)`

Shows a notification message.

- `message` (string): The notification text.

### `error(message)`

Displays a formatted error message.

- `message` (string): The error text.

### `success(message)`

Displays a formatted success message.

- `message` (string): The success text.

### `info(message)`

Displays a formatted informational message.

- `message` (string): The info text.

### `warning(message)`

Displays a formatted warning message.

- `message` (string): The warning text.

## Layouts & Widgets

### `layout(config)`

Creates a flexible layout with configurable direction, padding, and borders.

- `config` (dict): A dictionary of layout options.

### `grid(items, config)`

Creates a grid layout for a list of items.

- `items` (array): The items to display in the grid.
- `config` (dict): A dictionary with grid options, such as `columns`.

### `tabs(config)`

Creates a tabbed interface.

- `config` (dict): A dictionary where keys are the tab titles.

### `accordion(sections)`

Creates a collapsible accordion view.

- `sections` (dict): A dictionary where keys are the section titles and values are the content.

### `tree(config)`

Creates a tree view from a nested dictionary.

- `config` (dict): The nested dictionary representing the tree structure.

### `table(rows)`

Creates a formatted table from an array of rows.

- `rows` (array): An array of arrays, where each inner array is a row.

### `card(config)`

Creates a card with a title and content.

- `config` (dict): A dictionary with `title` and `content` keys.

### `list(items)`

Creates a bulleted list.

- `items` (array): An array of strings to display in the list.

### `split(left, right)`

Creates a split view with two panels.

- `left` (string): The content for the left panel.
- `right` (string): The content for the right panel.

### `modal(config)`

Creates a modal dialog.

- `config` (dict): A dictionary with `title` and `content` for the modal.

### `tooltip(text, message)`

Adds a tooltip to a piece of text.

- `text` (string): The text to add the tooltip to.
- `message` (string): The tooltip message.

## Visualizations

### `chart(data)`

Creates a simple bar chart from an array of numbers.

- `data` (array): An array of integers or floats.

### `gauge(value)`

Creates a gauge or progress indicator.

- `value` (integer): A value between 0 and 100.

### `heatmap(data)`

Creates a heatmap from an array of numbers.

- `data` (array): An array of integers or floats.

### `calendar(config)`

Displays a calendar for a given month and year.

- `config` (dict): A dictionary with `year` and `month` keys.

### `timeline(events)`

Creates a timeline from a list of events.

- `events` (array): An array of dictionaries, where each dictionary has `title` and `time` keys.

### `kanban(columns)`

Creates a Kanban board.

- `columns` (dict): A dictionary where keys are the column titles and values are arrays of tasks.

## Other Utilities

### `clear()`

Clears the terminal screen.

### `spinner(message)`

Displays a loading spinner with a message.

- `message` (string): The message to display next to the spinner.

### `progress(total)`

Initializes a progress bar.

- `total` (integer): The total value for the progress bar.

### `cursor(visible)`

Shows or hides the terminal cursor.

- `visible` (boolean): `true` to show the cursor, `false` to hide it.

### `beep()`

Plays a terminal beep sound.

### `moveCursor(x, y)`

Moves the cursor to a specific position.

- `x` (integer): The column to move to.
- `y` (integer): The row to move to.

### `getSize()`

Returns the size of the terminal as a dictionary with `width` and `height` keys.

### `countdown(duration)`

Creates a countdown timer.

- `duration` (integer): The duration of the countdown in seconds.

### `loading(message)`

Displays a loading message.

- `message` (string): The message to display.

### `dashboard(widgets)`

Creates a dashboard layout with multiple widgets.

- `widgets` (dict): A dictionary where keys are the widget titles and values are their content.

```

## time.md

```markdown
# Time in Vint

## Importing Time

To use time-related functionalities in Vint, you first need to import the `time` module as follows:
```js
import time
```

## Time Methods

### `now()`

To get the current time, use the `time.now()` method. This will return the current time as a `time` object:

```js
import time

current_time = time.now()
```

### `since()`

Use this method to get the total time since in seconds. It accepts a time object or a string in the format `HH:mm:ss dd-MM-YYYY`:

```js
import time

now = time.now()

time.since(now) // returns the since time

// alternatively:

now.since("00:00:00 01-01-1900") // returns the since time in seconds since that date
```

### `sleep()`

Use `sleep()` if you want your program to pause or "sleep." It accepts one argument, which is the total time to sleep in seconds:

```js
time.sleep(10) // will pause the program for ten seconds
```

### `add()`

Use the `add()` method to add to the current time, explained with an example:

```js
import time

now = time.now()

tomorrow = now.add(days=1)
next_hour = now.add(hours=24)
next_year = now.add(years=1)
three_months_later = now.add(months=3)
next_week = now.add(days=7)
custom_time = now.add(days=3, hours=4, minutes=50, seconds=3)
```

It will return a `time` object with the specified time added.

## Example Usage

### Print the current timestamp

```js
print(time.now())
```

### Function to greet a user based on the time of the day

```js
let greet = func(name) {
    let current_time = time.now()  // Get the current time
    print(current_time)            // Print the current time
    if (current_time.hour < 12) {  // Check if it's before noon
        print("Good morning, " + name + "!")
    } else {
        print("Good evening, " + name + "!")
    }
}
```

### Time-related operations

```js
year = 2024
print("Is", year, "Leap year:", time.isLeapYear(year))
print(time.format(time.now(), "02-01-2006 15:04:05"))
print(time.add(time.now(), "1h"))
print(time.subtract(time.now(), "2h30m45s"))
```

## Time Object Methods

Time objects in Vint have several powerful built-in methods for manipulation and extraction of time components:

### format()

Format the time object using a custom format string:

```js
import time

now = time.now()
formatted = now.format("2006-01-02 15:04:05")  // Standard format
print(formatted)  // 2024-08-11 15:30:45

// Custom formats
print(now.format("02-01-2006"))           // 11-08-2024
print(now.format("15:04"))                // 15:30
print(now.format("Monday, January 2, 2006"))  // Sunday, August 11, 2024
```

### year()

Get the year component of the time:

```js
import time

now = time.now()
current_year = now.year()
print("Current year:", current_year)  // Current year: 2024
```

### month()

Get the month component of the time (1-12):

```js
import time

now = time.now()
current_month = now.month()
print("Current month:", current_month)  // Current month: 8
```

### day()

Get the day component of the time (1-31):

```js
import time

now = time.now()
current_day = now.day()
print("Current day:", current_day)  // Current day: 11
```

### hour()

Get the hour component of the time (0-23):

```js
import time

now = time.now()
current_hour = now.hour()
print("Current hour:", current_hour)  // Current hour: 15
```

### minute()

Get the minute component of the time (0-59):

```js
import time

now = time.now()
current_minute = now.minute()
print("Current minute:", current_minute)  // Current minute: 30
```

### second()

Get the second component of the time (0-59):

```js
import time

now = time.now()
current_second = now.second()
print("Current second:", current_second)  // Current second: 45
```

### weekday()

Get the weekday name of the time:

```js
import time

now = time.now()
day_name = now.weekday()
print("Today is:", day_name)  // Today is: Sunday
```

## Practical Time Examples

Here are some practical examples using time methods:

```js
import time

// Create a timestamp logger
let log_with_timestamp = func(message) {
    let now = time.now()
    let timestamp = now.format("2006-01-02 15:04:05")
    print("[" + timestamp + "] " + message)
}

log_with_timestamp("Application started")
// Output: [2024-08-11 15:30:45] Application started

// Build a custom date display
let display_date = func() {
    let now = time.now()
    let weekday = now.weekday()
    let day = now.day()
    let month = now.month()
    let year = now.year()
    
    let months = ["", "January", "February", "March", "April", "May", "June",
                  "July", "August", "September", "October", "November", "December"]
    
    let formatted = weekday + ", " + months[month] + " " + day.to_string() + ", " + year.to_string()
    print(formatted)
}

display_date()
// Output: Sunday, August 11, 2024

// Time-based conditional logic
let get_greeting = func() {
    let now = time.now()
    let hour = now.hour()
    
    if (hour < 12) {
        return "Good morning!"
    } else if (hour < 18) {
        return "Good afternoon!"
    } else {
        return "Good evening!"
    }
}

print(get_greeting())

// Schedule checker
let is_business_hours = func() {
    let now = time.now()
    let hour = now.hour()
    let weekday = now.weekday()
    
    // Check if it's a weekday (Monday-Friday) and between 9 AM and 5 PM
    let is_weekday = weekday != "Saturday" && weekday != "Sunday"
    let is_work_time = hour >= 9 && hour < 17
    
    return is_weekday && is_work_time
}

if (is_business_hours()) {
    print("Office is open!")
} else {
    print("Office is closed!")
}

// Age calculator
let calculate_age = func(birth_year) {
    let now = time.now()
    let current_year = now.year()
    return current_year - birth_year
}

let age = calculate_age(1990)
print("Age:", age)

// Deadline checker
let check_deadline = func(deadline_date) {
    let now = time.now()
    let deadline = time.parse(deadline_date)  // Assuming we have a parse method
    
    let days_left = deadline.since(now) / (24 * 60 * 60)  // Convert seconds to days
    
    if (days_left > 0) {
        print("Deadline in", days_left.floor(), "days")
    } else {
        print("Deadline has passed!")
    }
}
```

## Method Chaining with Time

Time methods can be used in combination for complex operations:

```js
import time

// Get a formatted timestamp for a specific time
let birthday = time.now().add(days=30)
let birthday_info = "Birthday: " + birthday.weekday() + ", " + 
                   birthday.format("January 2, 2006") + " at " +
                   birthday.format("15:04")

print(birthday_info)
// Output: Birthday: Tuesday, September 10, 2024 at 15:30

// Create time-based file naming
let create_backup_filename = func(base_name) {
    let now = time.now()
    let timestamp = now.year().to_string() + 
                   now.month().to_string().padStart(2, "0") +
                   now.day().to_string().padStart(2, "0") + "_" +
                   now.hour().to_string().padStart(2, "0") +
                   now.minute().to_string().padStart(2, "0")
    
    return base_name + "_" + timestamp + ".backup"
}

let filename = create_backup_filename("database")
print(filename)  // database_20240811_1530.backup
```

```

## todo.md

```markdown
# TODOs

The `todo` keyword allows you to leave compiler-visible TODOs that warn at runtime.

## Syntax

```js
todo "Your todo message here"
```

When the Vint interpreter encounters a `todo` statement, it will print a warning to the console with your message, and then continue execution.

### Example

```js
todo "Implement user authentication"

let x = 10
println(x)
```

Running this script will output:

```
TODO: "Implement user authentication"
10
```

```

## tooling.md

```markdown
# Vint CLI Tooling

Vint provides several CLI tools to help you manage, format, and scaffold your projects:

---

## Formatter

Automatically formats your Vint code for consistency.

**Usage:**
```sh
vint fmt <file.vint>
```

This will overwrite the file with a pretty-printed version.

---

## Linter (Planned)

A linter will analyze your code for common mistakes and style issues.

**Planned Usage:**

```sh
vint lint <file.vint>
```

---

## Project Scaffolding

Quickly create a new Vint project with the recommended structure and sample files.

**Usage:**

```sh
vint init <project-name>
vint new <project-name>
```

This creates a new directory with a `main.vint`, `greetings_module.vint`, and a `vintconfig.json`.

---

## Package Manager

Install and manage Vint packages (currently supports installing `vintpm`).

**Usage:**

```sh
vint get <package>
```

---

For more information, run `vint help` or see the README.

```

## url.md

```markdown
# URL Module

The `url` module provides a set of functions for working with URLs. You can use it to parse, encode, decode, join, build, and validate URLs.

## Functions

### `parse(urlString)`

Parses a URL string and returns its components.

- `urlString` (string): The URL to parse.

**Returns:** A string containing the URL components (scheme, host, path, query, fragment).

**Usage:**

```js
import url

let components = url.parse("https://example.com/path?query=value#fragment")
println(components)
// Output: scheme:https host:example.com path:/path query:query=value fragment:fragment
```

### `build(components)`

Builds a URL from a dictionary of components.

- `components` (dict): A dictionary containing URL components.

**Valid components:**

- `scheme` (string): The URL scheme (e.g., "https", "http", "ftp")
- `host` (string): The hostname (e.g., "example.com", "localhost")
- `path` (string): The path component (e.g., "/api/v1")
- `query` (string): The query string (e.g., "limit=10&offset=0")
- `fragment` (string): The fragment identifier (e.g., "section1")
- `port` (string): The port number (e.g., "8080")
- `user` (string): User information for the URL

**Returns:** The constructed URL string.

**Usage:**

```js
import url

let components = {"scheme": "https", "host": "api.example.com", "path": "/v1/users", "query": "limit=10"}
let built_url = url.build(components)
println(built_url) // "https://api.example.com/v1/users?limit=10"

// Minimal example
let minimal = {"scheme": "http", "host": "localhost"}
println(url.build(minimal)) // "http://localhost"
```

### `encode(text)`

URL-encodes a string.

- `text` (string): The string to encode.

**Returns:** The URL-encoded string.

**Usage:**

```js
import url

let encoded = url.encode("hello world!")
println(encoded) // "hello%20world%21"
```

### `decode(encodedText)`

Decodes a URL-encoded string.

- `encodedText` (string): The URL-encoded string to decode.

**Returns:** The decoded string.

**Usage:**

```js
import url

let decoded = url.decode("hello%20world%21")
println(decoded) // "hello world!"
```

### `join(baseURL, path)`

Joins a base URL and a relative path to create a full URL.

- `baseURL` (string): The base URL.
- `path` (string): The relative path to join.

**Returns:** The full URL.

**Usage:**

```js
import url

let fullURL = url.join("https://example.com/", "/path/to/resource")
println(fullURL) // "https://example.com/path/to/resource"
```

### `isValid(urlString)`

Checks if a string is a valid URL.

- `urlString` (string): The URL to validate.

**Returns:** `true` if the URL is valid, `false` otherwise.

**Usage:**

```js
import url

println(url.isValid("https://example.com")) // true
println(url.isValid("not a url")) // false
```

```

## uuid.md

```markdown
# Using the UUID Module in Vint

The **UUID** (Universal Unique Identifier) module in Vint allows you to generate unique identifiers that are globally unique. These identifiers can be used for creating unique keys for database records, user sessions, or other purposes where a unique value is needed.

## Generating a UUID

You can generate a new UUID using the `uuid.generate()` function. It returns a unique identifier each time it is called.

### Example:

```js
import uuid

// Generate and print a new UUID
print(uuid.generate())
```

Each time `uuid.generate()` is called, it generates a new, unique UUID value. This is useful for ensuring that each identifier is distinct across systems.

The generated UUID can be used for various purposes such as tracking sessions, unique IDs for objects, or database keys.

```

## vintChart.md

```markdown
# VintChart Module (Experimental)

The `vintChart` module provides functions for creating various types of charts and saving them as HTML files. This module is experimental and its API may change in the future.

## Functions

### `barChart(labels, values, outputFile)`

Creates a bar chart.

- `labels` (array): An array of strings for the x-axis labels.
- `values` (array): An array of numbers for the y-axis values.
- `outputFile` (string): The path to save the HTML file (e.g., `"bar_chart.html"`).

**Usage:**

```js
import vintChart

let labels = ["A", "B", "C"]
let values = [10, 20, 15]
vintChart.barChart(labels, values, "my_bar_chart.html")
```

### `pieChart(labels, values, outputFile)`

Creates a pie chart.

- `labels` (array): An array of strings for the pie slices.
- `values` (array): An array of numbers for the values of the slices.
- `outputFile` (string): The path to save the HTML file (e.g., `"pie_chart.html"`).

**Usage:**

```js
import vintChart

let labels = ["Work", "Sleep", "Play"]
let values = [8, 8, 8]
vintChart.pieChart(labels, values, "my_pie_chart.html")
```

### `lineGraph(labels, values, outputFile)`

Creates a line graph.

- `labels` (array): An array of strings for the x-axis labels.
- `values` (array): An array of numbers for the y-axis values.
- `outputFile` (string): The path to save the HTML file (e.g., `"line_graph.html"`).

**Usage:**

```js
import vintChart

let labels = ["Jan", "Feb", "Mar"]
let values = [100, 120, 110]
vintChart.lineGraph(labels, values, "my_line_graph.html")
```

```

## vintSocket.md

```markdown
# VintSocket Module (Experimental)

The `vintSocket` module provides functions for working with WebSockets. It allows you to create WebSocket servers and connect to WebSocket servers as a client. This module is experimental and its API may change in the future.

## Functions

### `createServer(port)`

Creates a WebSocket server on the specified port.

- `port` (string): The port number to listen on.

**Usage:**

```js
import vintSocket

vintSocket.createServer("8080")
println("WebSocket server started on port 8080")
```

### `connect(url)`

Connects to a WebSocket server.

- `url` (string): The URL of the WebSocket server (e.g., `"ws://localhost:8080"`).

**Usage:**

```js
import vintSocket

vintSocket.connect("ws://localhost:8080")
```

### `sendMessage(clientIndex, message)`

Sends a message to a specific connected client.

- `clientIndex` (integer): The index of the client in the list of connections.
- `message` (string): The message to send.

**Usage:**

```js
import vintSocket

// Assuming a client is connected at index 0
vintSocket.sendMessage(0, "Hello, client!")
```

### `broadcast(message)`

Sends a message to all connected clients.

- `message` (string): The message to send.

**Usage:**

```js
import vintSocket

vintSocket.broadcast("Hello, everyone!")
```

```

## warn.md

```markdown
# Warn

The `warn` keyword allows you to emit non-fatal warnings at runtime.

## Syntax

`warn "Your warning message here"`

When the interpreter encounters a `warn` statement, it will print a formatted warning to the console and then continue execution. This is useful for alerting developers to potential issues that don't need to stop the program, such as using a deprecated feature or missing a configuration file.

### Example

```js
warn "Configuration file not found, using default settings."
println("Program is running with default configuration.")
```

Running this will output:

```

[WARN]: Configuration file not found, using default settings.

Program is running with default configuration.
```

```

## while.md

```markdown
# While Loops in Vint

While loops in Vint are used to execute a block of code repeatedly, as long as a given condition is true. This page covers the basics of while loops, including how to use the `break` and `continue` keywords within them.

## Basic Syntax

A while loop is executed when a specified condition is true. You initialize a while loop with the `while` keyword followed by the condition in parentheses `()`. The consequence of the loop should be enclosed in curly braces `{}`.

```js
let i = 1

while (i <= 5) {
    print(i)
    i++
}
```

### Output

```js
1
2
3
4
5
```

## Break and Continue

### Break

Use the `break` keyword to terminate a loop:

```js
let i = 1

while (i < 5) {
    if (i == 3) {
        print("broken")
        break
    }
    print(i)
    i++
}
```

### Output

```js
1
2
broken
```

### Continue

Use the `continue` keyword to skip a specific iteration:

```js
let i = 0

while (i < 5) {
    i++
    if (i == 3) {
        print("skipped")
        continue
    }
    print(i)
}
```

### Output

```js
1
2
skipped
4
5
```

By understanding while loops in Vint, you can create code that repeats a specific action or checks for certain conditions, offering more flexibility and control over your code execution.

## Repeat Loops

The `repeat` keyword allows you to execute a block of code a specific number of times. The default loop variable `i` is available inside the block, representing the current iteration (starting from 0).

### Syntax

```js
repeat 5 {
    println("Iteration:", i)
}
```

This will print:

```
Iteration: 0
Iteration: 1
Iteration: 2
Iteration: 3
Iteration: 4
```

You can also use an expression for the count:

```js
let n = 3
repeat n {
    println(i)
}
```

The variable `i` is always available in the scope of the repeat block.

```

## xml.md

```markdown
# XML Module in Vint

The XML module in Vint provides basic XML processing capabilities including validation, value extraction, and character escaping/unescaping. This module helps you work with XML data safely and efficiently.

---

## Importing the XML Module

To use the XML module, simply import it:
```js
import xml
```

---

## Functions and Examples

### 1. Validate XML (`validate`)

The `validate` function checks if a given XML string is well-formed and valid.

**Syntax**:

```js
validate(xmlString)
```

**Example**:

```js
import xml

print("=== XML Validation Example ===")

// Valid XML
valid_xml = "<root><name>John</name><age>30</age></root>"
is_valid = xml.validate(valid_xml)
print("Valid XML:", is_valid)
// Output: Valid XML: true

// Invalid XML
invalid_xml = "<root><name>John</age></root>"
is_invalid = xml.validate(invalid_xml)
print("Invalid XML:", is_invalid)
// Output: Invalid XML: false
```

---

### 2. Extract Value from XML Tag (`extract`)

The `extract` function extracts the value from a specific XML tag.

**Syntax**:

```js
extract(xmlString, tagName)
```

**Example**:

```js
import xml

print("=== XML Value Extraction Example ===")
xml_data = "<user><name>John Doe</name><email>john@example.com</email></user>"

name = xml.extract(xml_data, "name")
email = xml.extract(xml_data, "email")

print("Name:", name)
print("Email:", email)
// Output: 
// Name: John Doe
// Email: john@example.com
```

---

### 3. Escape XML Characters (`escape`)

The `escape` function escapes special XML characters to make text safe for XML content.

**Syntax**:

```js
escape(text)
```

**Example**:

```js
import xml

print("=== XML Escape Example ===")
unsafe_text = "<script>alert('Hello & Goodbye');</script>"
safe_text = xml.escape(unsafe_text)

print("Original:", unsafe_text)
print("Escaped: ", safe_text)
// Output: Escaped: &lt;script&gt;alert(&#39;Hello &amp; Goodbye&#39;);&lt;/script&gt;
```

---

### 4. Unescape XML Entities (`unescape`)

The `unescape` function converts XML entities back to their original characters.

**Syntax**:

```js
unescape(escapedText)
```

**Example**:

```js
import xml

print("=== XML Unescape Example ===")
escaped_text = "&lt;tag&gt;Hello &amp; Goodbye&lt;/tag&gt;"
unescaped_text = xml.unescape(escaped_text)

print("Escaped:  ", escaped_text)
print("Unescaped:", unescaped_text)
// Output: Unescaped: <tag>Hello & Goodbye</tag>
```

---

## Complete Usage Example

```js
import xml

print("=== XML Module Complete Example ===")

// Create XML data
user_name = "John & Jane"
user_email = "user@example.com"

// Escape data for safe XML insertion
safe_name = xml.escape(user_name)
safe_email = xml.escape(user_email)

// Build XML string
xml_string = "<user><name>" + safe_name + "</name><email>" + safe_email + "</email></user>"
print("Generated XML:", xml_string)

// Validate the XML
if xml.validate(xml_string) {
    print("XML is valid!")
    
    // Extract values
    extracted_name = xml.extract(xml_string, "name")
    extracted_email = xml.extract(xml_string, "email")
    
    // Unescape extracted values
    final_name = xml.unescape(extracted_name)
    final_email = xml.unescape(extracted_email)
    
    print("Final Name:", final_name)
    print("Final Email:", final_email)
} else {
    print("Generated XML is invalid!")
}
```

---

## Use Cases

- **XML Document Processing**: Parse and extract data from XML files
- **Web Scraping**: Extract information from XML responses
- **Configuration Files**: Read XML configuration data
- **Data Exchange**: Safely prepare data for XML transmission
- **Template Processing**: Build XML documents dynamically

---

## Summary of Functions

| Function    | Description                                         | Return Type |
|-------------|-----------------------------------------------------|-------------|
| `validate`  | Validates XML structure and syntax                  | Boolean     |
| `extract`   | Extracts value from a specific XML tag             | String      |
| `escape`    | Escapes special characters for safe XML content    | String      |
| `unescape`  | Converts XML entities back to original characters  | String      |

The XML module provides essential functionality for working with XML data safely and efficiently in VintLang applications.

```

## yaml.md

```markdown
# YAML Module

The YAML module provides functionality to work with YAML (YAML Ain't Markup Language) data in Vint. It allows you to parse YAML strings, convert Vint objects to YAML format, and manipulate YAML data structures.

## Functions

### yaml.decode(yamlString)

Parses a YAML string and converts it to Vint objects.

**Parameters:**

- `yamlString` (string): A valid YAML string to parse

**Returns:** Vint object (dict, array, string, number, boolean, or null)

**Example:**

```js
import yaml

let yamlData = yaml.decode("name: John\nage: 30\nactive: true")
print(yamlData) // {"name": "John", "age": 30, "active": true}
```

### yaml.encode(object)

Converts a Vint object to YAML format string.

**Parameters:**

- `object`: Any Vint object (dict, array, string, number, boolean, or null)

**Returns:** String containing YAML representation

**Example:**

```js
import yaml

let data = {
    "name": "Alice",
    "skills": ["Python", "Go", "Vint"],
    "active": true
}
let yamlString = yaml.encode(data)
print(yamlString)
```

### yaml.merge(object1, object2)

Merges two YAML-compatible objects into one. Properties from the second object will overwrite properties from the first object with the same key.

**Parameters:**

- `object1`: First object to merge
- `object2`: Second object to merge

**Returns:** New merged object

**Example:**

```js
import yaml

let obj1 = {"name": "John", "age": 30}
let obj2 = {"city": "NYC", "age": 35}
let merged = yaml.merge(obj1, obj2)
print(merged) // {"name": "John", "age": 35, "city": "NYC"}
```

### yaml.get(object, key)

Retrieves a value from a YAML-compatible object by key.

**Parameters:**

- `object`: The object to search in
- `key` (string): The key to look for

**Returns:** The value associated with the key, or null if not found

**Example:**

```js
import yaml

let data = yaml.decode("person:\n  name: Jane\n  age: 25")
let name = yaml.get(data.person, "name")
print(name) // "Jane"
```

## Supported YAML Features

The YAML module supports:

- Scalar values (strings, numbers, booleans, null)
- Sequences (arrays)
- Mappings (dictionaries/objects)
- Nested structures
- Multi-line strings

## Error Handling

All YAML functions return error objects when:

- Invalid YAML syntax is provided to `decode()`
- Incorrect argument types are passed
- Wrong number of arguments are provided

## Notes

- YAML keys are converted to strings in Vint objects
- The module handles YAML's flexible key types by converting them to strings
- Complex YAML features like anchors, references, and custom tags are not supported
- The implementation uses the gopkg.in/yaml.v3 library for robust YAML processing

## Usage with Files

You can combine the YAML module with file operations to work with YAML configuration files:

```js
import yaml
import os

// Read YAML config file
let configContent = os.read("config.yaml")
let config = yaml.decode(configContent)

// Modify and save back
config.updated = true
let newYaml = yaml.encode(config)
os.write("config.yaml", newYaml)
```

```

