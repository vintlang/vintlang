# VintLang FileWatcher Module

The `filewatcher` module provides functionality to monitor files and directories for changes. This is useful for building file watchers, auto-reloaders, build tools, and other applications that need to react to file system changes.

## Features

- Watch individual files for modifications
- Watch directories for file creation, modification, and deletion
- Recursive directory watching
- File extension filtering
- Configurable polling intervals
- Event-based callbacks with detailed information

## Usage

### Basic Example

```js
import filewatcher

// Watch a single file
let watcherId = filewatcher.watch("config.json", func(event) {
    print("File changed:", event["path"])
    print("Event type:", event["type"])
    print("Time:", event["time"])
    
    // Read the updated file
    let content = open(event["path"])
    print("New content:", content)
})

// Watch a directory
let dirWatcherId = filewatcher.watchDir("src", func(event) {
    print("File system change detected:")
    print("  Path:", event["path"])
    print("  Type:", event["type"]) // "created", "modified", or "deleted"
    print("  Time:", event["time"])
    
    // React to the change
    if (event["type"] == "created") {
        print("New file created!")
    } else if (event["type"] == "modified") {
        print("File was modified")
    } else if (event["type"] == "deleted") {
        print("File was deleted")
    }
})

// Stop watching after some time
setTimeout(func() {
    filewatcher.stopWatch(watcherId)
    filewatcher.stopWatch(dirWatcherId)
    print("Stopped watching")
}, 60000) // Stop after 60 seconds
```

## API Reference

### watch(path, callback, options)

Watches a file for changes and calls a callback function when changes are detected.

**Parameters:**
- `path` (string): The path to the file to watch
- `callback` (function): The function to call when the file changes. The callback receives an event object with the following properties:
  - `path` (string): The path to the file that changed
  - `type` (string): The type of change (always "modified" for single file watching)
  - `time` (string): The timestamp of the change
- `options` (dict, optional): Options for the watcher
  - `interval` (integer): The polling interval in milliseconds (default: 1000)

**Returns:**
- A watcher ID string that can be used to stop watching

### watchDir(path, callback, options)

Watches a directory for changes and calls a callback function when changes are detected.

**Parameters:**
- `path` (string): The path to the directory to watch
- `callback` (function): The function to call when changes are detected. The callback receives an event object with the following properties:
  - `path` (string): The path to the file that changed
  - `type` (string): The type of change ("created", "modified", or "deleted")
  - `time` (string): The timestamp of the change
- `options` (dict, optional): Options for the watcher
  - `interval` (integer): The polling interval in milliseconds (default: 1000)
  - `recursive` (boolean): Whether to watch subdirectories recursively (default: false)
  - `extensions` (array): Array of file extensions to watch (e.g., [".js", ".vint"]). If not provided, all files are watched.

**Returns:**
- A watcher ID string that can be used to stop watching

### stopWatch(watcherId)

Stops a file or directory watcher.

**Parameters:**
- `watcherId` (string): The watcher ID returned by `watch` or `watchDir`

**Returns:**
- `true` if the watcher was stopped successfully, `false` otherwise

### isWatching(path)

Checks if a file or directory is being watched.

**Parameters:**
- `path` (string): The path to check

**Returns:**
- `true` if the path is being watched, `false` otherwise

## Examples

### Auto-Reloading Development Server

```js
import filewatcher
import http
import os

// Simple HTTP server
let server = http.createServer(func(req, res) {
    if (req.path == "/") {
        // Serve index.html
        let content = open("public/index.html")
        res.writeHead(200, {"Content-Type": "text/html"})
        res.end(content)
    } else if (req.path == "/app.js") {
        // Serve app.js
        let content = open("public/app.js")
        res.writeHead(200, {"Content-Type": "application/javascript"})
        res.end(content)
    } else if (req.path == "/style.css") {
        // Serve style.css
        let content = open("public/style.css")
        res.writeHead(200, {"Content-Type": "text/css"})
        res.end(content)
    } else {
        // 404 Not Found
        res.writeHead(404)
        res.end("Not Found")
    }
})

// Start the server
server.listen(8080)
print("Server running at http://localhost:8080/")

// Set up WebSocket for live reload
let wsServer = http.createWebSocketServer(server)
let clients = []

wsServer.on("connection", func(client) {
    print("New client connected")
    clients.push(client)
    
    client.on("close", func() {
        // Remove client when disconnected
        let index = clients.indexOf(client)
        if (index != -1) {
            clients.splice(index, 1)
        }
    })
})

// Watch the public directory for changes
filewatcher.watchDir("public", func(event) {
    print("File changed:", event["path"])
    
    // Notify all connected clients to reload
    for (let client in clients) {
        client.send(JSON.stringify({
            type: "reload",
            path: event["path"]
        }))
    }
}, {
    recursive: true,
    extensions: [".html", ".js", ".css"]
})

print("Watching public directory for changes...")
```

### Build Tool

```js
import filewatcher
import os
import shell

// Function to build the project
let buildProject = func() {
    print("Building project...")
    
    // Compile all .vint files to .js
    let files = os.listDir("src")
    for (let file in files) {
        if (file.endsWith(".vint")) {
            let inputPath = "src/" + file
            let outputPath = "dist/" + file.replace(".vint", ".js")
            
            print("Compiling", inputPath, "to", outputPath)
            shell.exec("vint compile " + inputPath + " -o " + outputPath)
        }
    }
    
    // Bundle the JavaScript files
    print("Bundling JavaScript...")
    shell.exec("webpack --config webpack.config.js")
    
    print("Build completed!")
}

// Ensure dist directory exists
if (!os.exists("dist")) {
    os.mkdir("dist")
}

// Initial build
buildProject()

// Watch for changes
print("Watching for changes...")
filewatcher.watchDir("src", func(event) {
    if (event["type"] == "modified" || event["type"] == "created") {
        if (event["path"].endsWith(".vint")) {
            print("Source file changed:", event["path"])
            buildProject()
        }
    }
}, {
    recursive: true,
    extensions: [".vint"]
})

print("Build watcher started. Press Ctrl+C to stop.")
```

### Log File Monitor

```js
import filewatcher
import term

// Function to display the last N lines of a file
let tailFile = func(filePath, lines) {
    let content = open(filePath)
    let allLines = content.split("\n")
    let lastLines = allLines.slice(Math.max(0, allLines.length - lines))
    
    // Clear the screen
    term.clear()
    
    // Print header
    term.println("=== Log File Monitor ===", "#ffcc00")
    term.println("File: " + filePath, "#88ff88")
    term.println("Last " + lines + " lines:", "#88ff88")
    term.println("----------------------------", "#ffcc00")
    
    // Print the lines with syntax highlighting
    for (let line in lastLines) {
        if (line.includes("ERROR")) {
            term.println(line, "#ff5555") // Red for errors
        } else if (line.includes("WARNING")) {
            term.println(line, "#ffaa55") // Orange for warnings
        } else if (line.includes("INFO")) {
            term.println(line, "#55aaff") // Blue for info
        } else {
            term.println(line) // Default color for other lines
        }
    }
    
    // Print footer
    term.println("----------------------------", "#ffcc00")
    term.println("Press Ctrl+C to exit", "#88ff88")
}

// Check command line arguments
if (args.length < 2) {
    term.println("Usage: vint logmonitor.vint <log_file> [lines]", "#ff5555")
    exit(1)
}

let logFile = args[1]
let lines = 10 // Default to 10 lines

if (args.length >= 3) {
    lines = parseInt(args[2])
}

// Initial display
tailFile(logFile, lines)

// Watch the log file for changes
filewatcher.watch(logFile, func(event) {
    tailFile(logFile, lines)
}, {
    interval: 500 // Check every 500ms
})

print("Monitoring log file. Press Ctrl+C to exit.")
```