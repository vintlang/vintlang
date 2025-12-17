# Make Module Documentation

The `make` module provides a programmable build system for VintLang, designed to replace traditional Makefiles with a more flexible and powerful scripting approach.

## Overview

The `make` module allows you to:
- Execute shell commands
- Set environment variables
- Check for command availability
- Create complex build workflows using VintLang's full programming capabilities

## Installation

The `make` module is built-in to VintLang. Simply import it in your vint script:

```js
import make
```

## Functions

### make.check(command: string) -> boolean

Check if a command exists in the system PATH.

**Parameters:**
- `command`: Name of the command to check

**Returns:** `true` if the command exists, `false` otherwise

**Example:**
```js
if (make.check("upx")) {
    print("UPX is available")
} else {
    print("UPX not found. Please install it.")
}
```

### make.env(key: string, value: string) -> boolean

Set an environment variable for the current process.

**Parameters:**
- `key`: Environment variable name
- `value`: Environment variable value

**Returns:** `true` on success, error on failure

**Example:**
```js
make.env("GOOS", "linux")
make.env("GOARCH", "amd64")
```

### make.exec(command: string) -> string

Execute a shell command and return its output.

**Parameters:**
- `command`: Shell command to execute

**Returns:** Command output as a string, or error object on failure

**Example:**
```js
let output = make.exec("go build -o myapp")
print(output)
```

### make.echo(message: string) -> null

Print a message to stdout (similar to Makefile's `@echo`).

**Parameters:**
- `message`: Message to print

**Example:**
```js
make.echo("Building Linux binary...")
```

## Complete Example: Build System

Here's a complete example of a build system using the `make` module:

```js
// build.vint - A programmable build system
import make
import os
import cli

const VERSION = "1.0.0"
const LDFLAGS = "-s -w"

let echo = func(msg) {
    print("🔨 " + msg)
}

// Define build tasks
let tasks = {
    "build": func() {
        echo("Building application...")
        make.env("CGO_ENABLED", "0")
        let result = make.exec("go build -ldflags=\"" + LDFLAGS + "\" -o app")
        if (result.type != "error") {
            echo("✅ Build successful!")
        } else {
            print("❌ Build failed")
        }
    },
    
    "build_linux": func() {
        echo("Building for Linux...")
        make.env("GOOS", "linux")
        make.env("GOARCH", "amd64")
        make.exec("go build -ldflags=\"" + LDFLAGS + "\" -o app-linux")
        echo("✅ Linux build complete!")
    },
    
    "test": func() {
        echo("Running tests...")
        make.exec("go test ./...")
        echo("✅ Tests complete!")
    },
    
    "clean": func() {
        echo("Cleaning build artifacts...")
        make.exec("rm -f app app-linux app-windows.exe")
        echo("✅ Clean complete!")
    },
    
    "install_deps": func() {
        echo("Installing dependencies...")
        if (!make.check("go")) {
            print("❌ Go not found. Please install Go first.")
            os.exit(1)
        }
        make.exec("go mod download")
        echo("✅ Dependencies installed!")
    }
}

// Main execution
let args = cli.getArgs()
if (len(args) < 1) {
    print("Usage: vint build.vint <task>")
    print("Available tasks:", tasks.keys())
    os.exit(1)
}

let taskName = args[0]
if (tasks[taskName] != null) {
    tasks[taskName]()
} else {
    print("Unknown task:", taskName)
    os.exit(1)
}
```

**Usage:**
```bash
# Build the application
vint build.vint build

# Build for Linux
vint build.vint build_linux

# Run tests
vint build.vint test

# Clean build artifacts
vint build.vint clean
```

## Advantages Over Makefile

1. **Full Programming Language**: Use loops, conditionals, functions, and all VintLang features
2. **Cross-Platform**: Works consistently across Windows, Linux, and macOS
3. **Better Error Handling**: Proper error handling with try-catch patterns
4. **Modularity**: Import other modules (json, yaml, http, etc.) for advanced build scripts
5. **Type Safety**: Benefit from VintLang's type system
6. **Debugging**: Easier to debug than shell scripts

## Migration from Makefile

Here's how common Makefile patterns translate to the `make` module:

### Setting Variables
**Makefile:**
```makefile
VERSION=1.0.0
LDFLAGS=-s -w
```

**VintLang:**
```js
const VERSION = "1.0.0"
const LDFLAGS = "-s -w"
```

### Environment Variables
**Makefile:**
```makefile
build:
	GOOS=linux GOARCH=amd64 go build
```

**VintLang:**
```js
make.env("GOOS", "linux")
make.env("GOARCH", "amd64")
make.exec("go build")
```

### Command Execution
**Makefile:**
```makefile
build:
	@echo "Building..."
	go build -o app
```

**VintLang:**
```js
make.echo("Building...")
make.exec("go build -o app")
```

### Conditional Execution
**Makefile:**
```makefile
build:
	@which upx || echo "UPX not found"
```

**VintLang:**
```js
if (!make.check("upx")) {
    print("UPX not found")
}
```

## Best Practices

1. **Organize Tasks**: Use dictionaries or functions to organize your build tasks
2. **Error Handling**: Check return values from `make.exec()` for errors
3. **Logging**: Use `make.echo()` for user-friendly progress messages
4. **Modularity**: Split complex build scripts into multiple files
5. **Documentation**: Add comments explaining what each task does

## See Also

- [OS Module](./os_module.md) - File system operations
- [Shell Module](./shell_module.md) - Additional shell utilities
- [Examples](../examples/) - More build script examples
