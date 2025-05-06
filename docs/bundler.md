# VintLang Bundler

## Overview

The **VintLang Bundler** compiles `.vint` source code files into standalone Go binaries.

This means you can **write code in VintLang**, bundle it into an executable, and run it anywhere — **without needing the Vint interpreter installed** on the target machine.

---

## ✨ Why Use the Bundler?

* ✅ **Distribute VintLang programs as standalone executables**
* ✅ **No need for the end-user to install Go or VintLang**
* ✅ Useful for **deployment, automation, and sharing your VintLang tools**
* ✅ Integrates seamlessly with `vintlang/repl` under the hood

---

## 📥 Installation Requirements

To use the bundler, you must have the following installed:

### 1. 🛠️ Go (version 1.18 or later)

Install Go from the official site:
[https://go.dev/dl/](https://go.dev/dl/)

Once installed, confirm with:

```sh
go version
```

### 2. 🌐 Git & Go Modules Support

Make sure Go modules are enabled:

```sh
go env -w GO111MODULE=on
```

### 3. 🧠 Install VintLang

Install VintLang globally (if you haven’t already):

```sh
go install github.com/vintlang/vintlang@latest
```

Or add it to your Go project:

```sh
go get github.com/vintlang/vintlang
```

---

## 🏗️ How the Bundler Works

The bundler does the following under the hood:

1. **Reads** your `.vint` source file.
2. **Escapes** any backticks in your source to safely embed it in Go code.
3. **Generates a `main.go`** file that embeds the source and calls `repl.Read(...)`.
4. **Creates a temporary Go module**, initializes it with `go.mod`, and builds a binary.
5. **Outputs a self-contained executable** with the same name as your `.vint` file.

You get a file like `hello` or `myapp` — which you can run directly:

```sh
./hello
```

No external dependencies. Just run it.

---

## 📂 Example

Assume you have a VintLang file called `hello.vint`:

```vint
print("Hello, World!")
```

Now bundle it:

```go
err := bundler.Bundle("hello.vint")
if err != nil {
	fmt.Println("Bundle failed:", err)
}
```

You’ll get an executable named `hello`. Run it:

```sh
./hello
```

✅ Output:

```
Hello, World!
```

---

## 🔍 What’s Inside the Generated Code?

A temporary `main.go` is generated that looks like this:

```go
package main

import (
	"github.com/vintlang/vintlang/repl"
)

func main() {
	code := ` + "`<your original Vint code>`" + `
	repl.Read(code)
}
```

It is compiled using Go’s `go build`, resulting in a binary that includes your Vint source embedded at compile time.

---

## 🚀 Use Cases

* Ship **CLI tools** written in VintLang.
* Create **portable binaries** for VintLang scripts.
* Deploy VintLang logic on servers **without installing VintLang there**.
* Experiment with building **VintLang-based applications** while keeping your Go toolchain.

---

## 🧰 Developer Notes

* The bundler uses `text/template` to safely embed Vint code.
* It uses a loading spinner for nicer CLI feedback.
* It auto-generates and cleans up temp build folders.
* The embedded Go module is local and isolated to avoid polluting your workspace.

---

## ❗ Important

* The bundled binary **still depends on Go at build time**, but **not at runtime**.
* You must run `go mod tidy` and `go build` during bundling — make sure your system Go installation is working.

---

## 📌 Conclusion

The **VintLang Bundler** brings the power of binary distribution to VintLang.

Whether you're building tools, scripts, or micro-apps — bundle them once, run them anywhere.

