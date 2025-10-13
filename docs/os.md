# OS module in VintLang

In **Vint**, the `os` module provides functions to interact with the operating system. You can execute commands, manage files, and interact with the environment. This guide covers common operations using the `os` module.

## Exit with a Status Code

You can exit a program and return a status code using `os.exit()`. A non-zero status code generally indicates an error.

```js
os.exit(1)  // Exit with status code 1
```

## Run a Shell Command

Use `os.run()` to execute a shell command. It returns the result of the command as a string.

```js
result = os.run("ls -la")
print(result)  // Outputs the directory listing
```

You can also run other commands, like:

```js
// print(os.run("go run . vintLang/main.vint"))
```

## Get the working dir
use `os.getwd()`

## Get and Set Environment Variables

Environment variables can be set, retrieved, and removed with `os.setEnv()`, `os.getEnv()`, and `os.unsetEnv()`.

### Set Environment Variable:
```js
os.setEnv("API_KEY", "12345")
```

### Get Environment Variable:
```js
api_key = os.getEnv("API_KEY")
print(api_key)  // Outputs: "12345"
```

### Unset Environment Variable:
```js
os.unsetEnv("API_KEY")
api_key = os.getEnv("API_KEY")
print(api_key)  // Outputs: "" (empty string)
```

## Read and Write Files

### Write to a File:
```js
os.writeFile("example.txt", "Hello, Vint!")
```

### Read from a File:
```js
content = os.readFile("example.txt")
print(content)  // Outputs: "Hello, Vint!"
```

## List Directory Contents

### Get Directory Contents as String

You can list the files in a directory using `os.listDir()`:

```js
files = os.listDir(".")
print(files)  // Outputs a comma-separated string of files in the current directory
```

### Get Directory Contents as Array

For more convenient file manipulation, use `os.listFiles()` which returns an array:

```js
files = os.listFiles(".")
print(files)  // Outputs: ["file1.txt", "file2.txt", "subdirectory", ...]

// Iterate through files
for (file in files) {
    print("Found:", file)
}

// Check for specific files
if ("README.md" in files) {
    print("README found!")
}
```

The `listFiles()` function is particularly useful when you need to:

- Iterate through files individually
- Filter files based on conditions
- Count the number of files
- Process files programmatically

## Create a Directory

Use `os.makeDir()` to create a new directory:

```js
os.makeDir("new_folder")
```

## Check if a File Exists

To check if a file exists, use `os.fileExists()`:

```js
exists = os.fileExists("example.txt")
print(exists)  // Outputs: false (if the file doesn't exist)
```

## Read a File Line by Line

To read a file and get its lines in a list, use `os.readLines()`:

```js
os.writeFile("example.txt", "Hello\nWorld")
lines = os.readLines("example.txt")
print(lines)  // Outputs: ["Hello", "World"]
```

## Delete a File

To delete a file, use `os.deleteFile()`:

```js
// os.deleteFile("example.txt")
```

## System Information

The `os` module provides several functions to get system information:

### Get Home Directory:
```js
home = os.homedir()
print(home)  // Outputs: "/home/username"
```

### Get Temporary Directory:
```js
temp = os.tmpdir()
print(temp)  // Outputs: "/tmp" (on Unix-like systems)
```

### Get CPU Count:
```js
cpus = os.cpuCount()
print(cpus)  // Outputs: 4 (number of logical CPUs)
```

### Get Hostname:
```js
hostname = os.hostname()
print(hostname)  // Outputs: "my-computer"
```

## File Operations

### Copy a File:
```js
os.copy("source.txt", "destination.txt")
print("File copied successfully")
```

### Move or Rename a File:
```js
os.move("old_name.txt", "new_name.txt")
print("File moved/renamed successfully")
```

By utilizing the **Vint** `os` module, you can effectively manage files, directories, environment variables, and system information within your programs.