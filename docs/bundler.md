# VintLang Bundler

## Overview

The **VintLang Bundler** compiles `.vint` source files into standalone Go binaries.

```sh
vint bundle yourfile.vint
```

This allows you to write code in VintLang, bundle it into an executable, and run it on any system without requiring the VintLang interpreter or Go to be installed on the target machine.

---

## Why Use the Bundler?

* Package and distribute VintLang scripts as self-contained executables
* End-users don’t need to install Go or VintLang
* Ideal for deploying scripts, shipping CLI tools, and automating workflows
* Internally powered by the `vintlang/repl` package for code execution

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

* ✅ **Automatic Dependency Discovery**: Finds all imported and included `.vint` files recursively
* ✅ **Package System Integration**: Handles `package` declarations and `import` statements
* ✅ **Include Statement Support**: Handles `include` statements for direct file embedding
* ✅ **Self-Contained Binaries**: No external `.vint` files needed at runtime
* ✅ **Compatible with Built-ins**: Works with all VintLang built-in modules

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

* **Import statements** (`import module_name`) work with the package system and wrap content in packages
* **Include statements** (`include "file_path"`) directly embed file content without package wrapping
* Both are automatically discovered and bundled into self-contained binaries

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

* Distribute command-line tools built in VintLang
* Deploy scripts on systems where VintLang is not installed
* Share portable binaries for automation or education
* Build lightweight tools using VintLang and Go’s compiler

---

## Notes for Developers

* Temporary build directories are automatically created and cleaned
* Uses `text/template` for safe source code embedding
* The Go module created during bundling is isolated from your current project
* Spinner and CLI output are available for build feedback

---

## Important Details

* Go is required only during **build time**
* The resulting binary is portable and self-contained
* Cross-compilation is not supported out-of-the-box; build on the target OS/arch

---

## Conclusion

The **VintLang Bundler** lets you turn `.vint` files into standalone executables using a simple command:

```sh
vint bundle yourfile.vint
```

Build once. Run anywhere. No dependencies. No interpreter. Just execution.

