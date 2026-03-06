# VintLang

**VintLang** is a modern, expressive programming language designed for clarity and productivity. It combines familiar syntax with powerful features for building robust applications.

---

## Language Features

### Modern Syntax & Control Flow

```js
// If expressions - use if as both statement and expression
let status = if (user.isActive) { "Online" } else { "Offline" }

// Pattern matching with guards
match user.role {
    "admin" => {
        print("Full access granted")
    }
    "user" => {
        print("Standard access")
    }
    _ => {
        print("Guest access")
    }
}

// Switch with variable binding
switch (response.status) {
    case code if code >= 200 && code < 300 {
        print("Success:", code)
    }
    default {
        print("Error occurred")
    }
}
```

### Flexible Loops & Iteration

```js
// For-in loops with arrays, strings, dictionaries
for item in items {
    process(item)
}

for key, value in data {
    print(key + ":", value)
}

// While loops with break/continue
while (condition) {
    if (shouldSkip) {
        continue
    }
    process()
}

// Repeat loops
repeat 5 {
    print("Iteration:", i)
}
```

### Rich Data Structures

```js
// Arrays with built-in methods
let numbers = [1, 2, 3, 4, 5]
let doubled = map(numbers, func(x) { return x * 2 })
let evens = filter(numbers, func(x) { return x % 2 == 0 })

// Dictionaries
let config = {
    "app": {"name": "myapp", "port": 8080},
    "features": ["web", "logging"]
}
```

### Modules & Packages

```js
// Static imports
import os
import json
import time

// Dynamic imports at runtime
let yaml = import("yaml")
let math = import("math")

// Custom packages
import my_package
my_package.doSomething()
```

### File I/O & System Operations

```js
// Read and write files
let content = os.readFile("config.yaml");
os.writeFile("output.json", json.encode(data));

// Check file existence
if (os.fileExists("config.json")) {
  // load config
}
```

### Comprehensive Example

```js
import os
import json
import time

// Dynamic imports
let yaml = import("yaml")
let math = import("math")

// Load or create configuration
let cfgStr = ""
if (os.fileExists("config.yaml")) {
    cfgStr = os.readFile("config.yaml")
} else {
    let defaultCfg = {
        "app": {"name": "vint-app", "port": 8080},
        "features": ["web", "logging"]
    }
    cfgStr = yaml.encode(defaultCfg)
    os.writeFile("config.yaml", cfgStr)
}

let cfg = yaml.decode(cfgStr)
let pow = math.pow(2, 10)

let summary = {
    "generated_at": time.format(time.now(), "2006-01-02T15:04:05"),
    "app": yaml.get(cfg, "app.name"),
    "value": pow
}

os.writeFile("summary.json", json.encode(summary))
```

Run this example:

```bash
vint examples/comprehensive_showcase.vint
```

---

## Installation

Follow the steps below to easily install **VintLang** on your Linux or macOS system.

### Quick Install (Linux / macOS)

Download and install the latest release automatically:

```bash
# Linux (amd64)
curl -sL https://github.com/vintlang/vintlang/releases/latest/download/vint-linux-amd64.tar.gz | sudo tar xz -C /usr/local/bin

# Linux (arm64)
curl -sL https://github.com/vintlang/vintlang/releases/latest/download/vint-linux-arm64.tar.gz | sudo tar xz -C /usr/local/bin

# macOS Apple Silicon (M1/M2/M3/M4)
curl -sL https://github.com/vintlang/vintlang/releases/latest/download/vint-darwin-arm64.tar.gz | sudo tar xz -C /usr/local/bin

# macOS Intel
curl -sL https://github.com/vintlang/vintlang/releases/latest/download/vint-darwin-amd64.tar.gz | sudo tar xz -C /usr/local/bin
```

### Windows

Download the latest `.zip` from the [releases page](https://github.com/vintlang/vintlang/releases/latest), extract it, and add the folder to your `PATH`.

### Android (Termux)

```bash
curl -sL https://github.com/vintlang/vintlang/releases/latest/download/vint-android-arm64.tar.gz | tar xz -C $PREFIX/bin
```

### Verify & Initialize

```bash
# Check version
vint -v

# Create a new project
vint init <project-name>
```

> **Tip:** Install the `vintlang` VSCode extension for syntax highlighting, snippets, and tooling.

---

## Summary of Installation Steps

1. **Download the latest binary** for your platform from the [releases page](https://github.com/vintlang/vintlang/releases/latest).
2. **Extract** to a directory on your `PATH` (e.g. `/usr/local/bin`).
3. **Verify** by running `vint -v`.
4. **Initialize a project** with `vint init <project-name>`.
5. **Install the vintlang VSCode extension** for language support.

## Contributing

We welcome contributions to VintLang! Whether you're fixing bugs, adding features, or improving documentation, your help is appreciated.

---

## Author

Created by **Tachera Sasi**

---
