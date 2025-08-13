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

## How It Works

The bundler performs the following steps internally:

1. Reads the `.vint` source file
2. Escapes characters as needed for embedding in Go
3. Generates a temporary `main.go` that runs the embedded code via `repl.Read(...)`
4. Initializes a temporary Go module and compiles the binary using `go build`
5. Outputs a binary named after the original `.vint` file

No external dependencies are required to run the resulting binary.

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

