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
    "admin" if user.permissions > 10 => {
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
let content = os.readFile("config.yaml")
os.writeFile("output.json", json.encode(data))

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

### For Linux

1. **Download the Binary:**

   First, download the **VintLang** binary for Linux. You can do this using the `curl` command. This will download the `tar.gz` file containing the binary to your current directory.

   ```bash
   curl -O -L https://github.com/vintlang/vintlang/releases/download/v0.2.0/vintLang_linux_amd64.tar.gz
   ```

2. **Extract the Binary to a Global Location:**

   After downloading the binary, you need to extract it into a directory that is globally accessible. `/usr/local/bin` is a commonly used directory for this purpose. The `tar` command will extract the contents of the `tar.gz` file and place them in `/usr/local/bin`.

   ```bash
   sudo tar -C /usr/local/bin -xzvf vintLang_linux_amd64.tar.gz
   ```

   This step ensures that the **VintLang** command can be used from anywhere on your system.

3. **Verify the Installation:**

   Once the extraction is complete, confirm that **VintLang** was installed successfully by checking its version. If the installation was successful, it will display the installed version of **VintLang**.

   ```bash
   vint -v
   ```

4. **Initialize a vint project**

   Create a simple boilerplate vint project

   ```bash
   vint <optional:project-name>
   ```

---

Note: Install the `vintlang` VSCode extension for language support (syntax highlighting, snippets, and tooling).

---

## For macOS

1. **Download the Binary:**

   Begin by downloading the **VintLang** binary for macOS using the following `curl` command. This will download the `tar.gz` file for macOS to your current directory.

   ```bash
   curl -O -L https://github.com/vintlang/vintlang/releases/download/v0.2.0/vintLang_mac_amd64.tar.gz
   ```

2. **Extract the Binary to a Global Location:**

   Next, extract the downloaded binary to a globally accessible location. As with Linux, the standard directory for this on macOS is `/usr/local/bin`. Use the following command to extract the binary:

   ```bash
   sudo tar -C /usr/local/bin -xzvf vintLang_mac_amd64.tar.gz
   ```

   This allows you to run **VintLang** from any terminal window.

3. **Verify the Installation:**

   To check that the installation was successful, run the following command. It will output the version of **VintLang** that was installed:

   ```bash
   vint -v
   ```

4. **Initialize a vint project**

   Create a simple boilerplate vint project

   ```bash
   vint init <optional:project-name>
   ```

---

Note: Install the `vintlang` VSCode extension for language support (syntax highlighting, snippets, and tooling).

---

## Summary of Installation Steps

1. **Download the Binary** using `curl` for your system (Linux or macOS).
2. **Extract the Binary** to `/usr/local/bin` (or another globally accessible directory).
3. **Verify the Installation** by checking the version with `vint -v`.
4. **Initialize a vintlang project** by running `vint init <projectname>`.
5. **Install the vintlang extension from vscode** install vintlang extension in vscode


## Contributing

We welcome contributions to VintLang! Whether you're fixing bugs, adding features, or improving documentation, your help is appreciated.

---

## Author

Created by **Tachera Sasi**

---
