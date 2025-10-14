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
