# VintLang Installation Guide (Linux & macOS)

Follow the steps below to easily install **VintLang** on your Linux or macOS system.

---

## For Linux

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

## Sample Code — Rich example

The simple examples below have been replaced with a single, richer example that demonstrates several recent and powerful VintLang features: packages, dynamic `import()` at runtime, YAML and JSON handling, file I/O, and reusable functions.
<!-- using ```js so that we get syntax highlighting-->
```js
import os
import json

// If you have a local package (see `examples/packages_example`) you can import it
// using a package name. The example package defines `greeter_pkg.greet()`.
// import greeter_pkg

// Dynamic import() lets you load modules at runtime (useful for plugins)
let yaml = import("yaml")
let time = import("time")
let math = import("math")

// Load a YAML configuration if present; otherwise write a sensible default
let cfgStr = ""
if (os.fileExists("config.yaml")) {
   cfgStr = os.readFile("config.yaml")
   print("Loaded config.yaml")
} else {
   let defaultCfg = {
      "app": {"name": "vint-app", "port": 8080},
      "features": ["web", "logging", "yaml"]
   }
   cfgStr = yaml.encode(defaultCfg)
   os.writeFile("config.yaml", cfgStr)
   print("Wrote default config.yaml")
}

let cfg = yaml.decode(cfgStr)
print("App name:", yaml.get(cfg, "app.name"))

// Use dynamic math module
let pow = math.pow(2, 10)  // 2^10
print("2^10 =", pow)

// Optional: call into a package if available (see examples/packages_example)
// Example package API (greeter_pkg): sayHello, setGreeting, getPackageInfo
// To use it uncomment and run:
// import greeter_pkg
// greeter_pkg.sayHello("Vint User")
// print(greeter_pkg.getPackageInfo())

// Produce a JSON summary
let summary = {
   "generated_at": time.format(time.now(), "2006-01-02T15:04:05"),
   "app": yaml.get(cfg, "app.name"),
   "value": pow
}

os.writeFile("summary.json", json.encode(summary))
print("Wrote summary.json")

// Small reusable function
let report = func(path) {
   let content = os.readFile(path)
   print("Report (" + path + ") size:", string(len(content)))
}

report("summary.json")

// End of example
```

How to run this example locally:

```bash
vint examples/comprehensive_showcase.vint
```

Notes:

- The example above intentionally mixes static and dynamic imports to show both workflows.
- Some examples in `examples/` (LLM, HTTP, enterprise integrations) require network access or API keys — they are safe to read but may need extra setup to run.

---

## Contributing

We welcome contributions to VintLang! Whether you're fixing bugs, adding features, or improving documentation, your help is appreciated.

---

## Author

Created by **Tachera Sasi**

---
