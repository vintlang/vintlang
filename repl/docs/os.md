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

## Get and Set Environment Variables

Environment variables can be set and retrieved with `os.setEnv()` and `os.getEnv()`.

### Set Environment Variable:
```js
os.setEnv("API_KEY", "12345")
```

### Get Environment Variable:
```js
api_key = os.getEnv("API_KEY")
print(api_key)  // Outputs: "12345"
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

You can list the files in a directory using `os.listDir()`:

```js
files = os.listDir(".")
print(files)  // Outputs a list of files in the current directory
```

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

By utilizing the **Vint** `os` module, you can effectively manage files, directories, and environment variables within your programs.