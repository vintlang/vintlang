# Files in Vint

The `files` module in Vint provides comprehensive functionality for working with files, including reading, writing, file operations, and metadata access.

---

## Opening a File

You can open a file using the `open` keyword. This will return an object of type `FAILI`, which represents the file.

### Syntax

```js
fileObject = open("filename.txt")
```

### Example

```js
myFile = open("file.txt")

aina(myFile) // Output: FAILI
```

---

## File Methods

File objects in Vint come with powerful built-in methods for comprehensive file operations:

### read()

Read the entire contents of a file as a string:

```js
myFile = open("example.txt")
contents = myFile.read()
print(contents)
```

### write()

Write content to a file (overwrites existing content):

```js
myFile = open("output.txt")
myFile.write("Hello, World!")
```

### append()

Append content to the end of a file:

```js
logFile = open("app.log")
logFile.append("New log entry\n")
```

### exists()

Check if the file exists:

```js
myFile = open("config.txt")
if (myFile.exists()) {
    print("File exists!")
} else {
    print("File not found!")
}
```

### size()

Get the size of the file in bytes:

```js
myFile = open("data.txt")
fileSize = myFile.size()
print("File size:", fileSize, "bytes")
```

### delete()

Delete the file from the filesystem:

```js
tempFile = open("temp.txt")
if (tempFile.exists()) {
    tempFile.delete()
    print("File deleted successfully")
}
```

### copy()

Copy the file to a new location:

```js
sourceFile = open("original.txt")
sourceFile.copy("backup.txt")
print("File copied successfully")
```

### move()

Move or rename the file:

```js
oldFile = open("old_name.txt")
oldFile.move("new_name.txt")
print("File moved/renamed successfully")
```

### lines()

Read the file content as an array of lines:

```js
configFile = open("settings.conf")
lines = configFile.lines()
for line in lines {
    print("Config:", line.trim())
}
```

### extension()

Get the file extension:

```js
documentFile = open("report.pdf")
ext = documentFile.extension()
print("File extension:", ext)  // .pdf

imageFile = open("photo.jpg")
print("Extension:", imageFile.extension())  // .jpg
```

## Practical File Examples

Here are some practical examples using file methods:

```js
// Log file manager
let log_message = func(message) {
    let logFile = open("application.log")
    let timestamp = time.now().format("2006-01-02 15:04:05")
    let entry = "[" + timestamp + "] " + message + "\n"
    logFile.append(entry)
}

log_message("Application started")
log_message("User logged in")

// File backup system
let backup_file = func(filename) {
    let sourceFile = open(filename)
    if (sourceFile.exists()) {
        let backup_name = filename + ".backup"
        sourceFile.copy(backup_name)
        print("Backup created:", backup_name)
        return true
    } else {
        print("Source file not found:", filename)
        return false
    }
}

backup_file("important_data.txt")

// Configuration file processor
let process_config = func(config_file) {
    let file = open(config_file)
    if (!file.exists()) {
        print("Config file not found, creating default...")
        file.write("debug=false\nport=8080\nhost=localhost\n")
        return
    }
    
    let lines = file.lines()
    let settings = {}
    
    for line in lines {
        if (line.contains("=")) {
            let parts = line.split("=")
            if (parts.length() == 2) {
                settings.set(parts[0].trim(), parts[1].trim())
            }
        }
    }
    
    print("Loaded settings:", settings)
    return settings
}

let config = process_config("app.conf")

// File size analyzer
let analyze_files = func(filenames) {
    let total_size = 0
    let file_info = []
    
    for filename in filenames {
        let file = open(filename)
        if (file.exists()) {
            let size = file.size()
            let ext = file.extension()
            total_size += size
            
            file_info.push({
                "name": filename,
                "size": size,
                "extension": ext
            })
        }
    }
    
    print("Total size:", total_size, "bytes")
    print("File details:")
    for info in file_info {
        print("  ", info["name"], "-", info["size"], "bytes", info["extension"])
    }
}

analyze_files(["document.pdf", "image.jpg", "data.csv"])

// Text file processor with method chaining
let process_text_file = func(input_file, output_file) {
    let inputFile = open(input_file)
    
    if (!inputFile.exists()) {
        print("Input file not found")
        return false
    }
    
    // Read and process content
    let content = inputFile.read()
    let processed = content.upper().replace("OLD", "NEW").trim()
    
    // Write to output file
    let outputFile = open(output_file)
    outputFile.write(processed)
    
    print("Processing complete:")
    print("  Input size:", inputFile.size(), "bytes")
    print("  Output size:", outputFile.size(), "bytes")
    print("  Output extension:", outputFile.extension())
    
    return true
}

process_text_file("input.txt", "output.txt")

// File cleanup utility
let cleanup_temp_files = func(directory_pattern) {
    let temp_files = ["temp1.tmp", "cache.tmp", "old_data.bak"]
    let deleted_count = 0
    
    for filename in temp_files {
        let file = open(filename)
        if (file.exists()) {
            let size = file.size()
            file.delete()
            print("Deleted:", filename, "(" + size.to_string() + " bytes)")
            deleted_count++
        }
    }
    
    print("Cleanup complete. Deleted", deleted_count, "files")
}

cleanup_temp_files("*.tmp")
```

## File Error Handling

When working with files, it's important to handle potential errors:

```js
let safe_file_operation = func(filename, operation) {
    let file = open(filename)
    
    // Always check if file exists for read operations
    if (operation == "read" && !file.exists()) {
        print("Error: File", filename, "does not exist")
        return null
    }
    
    // Get file info before operations
    if (file.exists()) {
        print("File info:")
        print("  Size:", file.size(), "bytes")
        print("  Extension:", file.extension())
    }
    
    // Perform the operation
    if (operation == "read") {
        return file.read()
    } else if (operation == "backup") {
        let backup_name = filename + ".backup"
        file.copy(backup_name)
        return backup_name
    }
    
    return null
}

// Safe file reading
let content = safe_file_operation("data.txt", "read")
if (content != null) {
    print("File content loaded successfully")
}
```

---

## Notes

- All file operations are performed relative to the current working directory unless an absolute path is specified
- File methods support method chaining for fluent operations
- Always check if a file exists before performing read operations
- The `lines()` method automatically handles different line ending formats
- File extensions are returned with the leading dot (e.g., ".txt", ".pdf")

---