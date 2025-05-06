# VintLang Editor Module

The `editor` module provides a simple text editor component for terminal applications. It allows you to create and manipulate text editors within your VintLang applications, which is useful for note-taking apps, configuration editors, and other text-based tools.

## Features

- Create and manage text editors
- Set and get text content
- Line-based operations (insert, delete, update, get)
- Vim-like modal editing (normal, insert, command modes)
- Syntax highlighting (basic)
- Line numbering
- File loading and saving
- Customizable dimensions and appearance

## Usage

### Basic Example

```js
import editor
import term

// Create a new editor
let ed = editor.newEditor({
    width: 80,
    height: 20,
    showLineNum: true,
    syntax: "text"
})

// Set some initial text
editor.setText(ed, "Hello, world!\nThis is a simple text editor.\nTry it out!")

// Render the editor
let editorContent = editor.render(ed)
print(editorContent)

// Handle keypresses
while (true) {
    let key = term.getKey()
    
    // Handle special keys
    if (key == "q") {
        break // Exit the loop
    }
    
    // Pass the key to the editor
    editor.handleKeypress(ed, key)
    
    // Re-render the editor
    editorContent = editor.render(ed)
    term.clear()
    print(editorContent)
}
```

## API Reference

### newEditor(options)

Creates a new text editor.

**Parameters:**
- `options` (dict, optional): Options for the editor
  - `filename` (string): Initial file to load
  - `width` (integer): Width of the editor in characters (default: 80)
  - `height` (integer): Height of the editor in lines (default: 24)
  - `showLineNum` (boolean): Whether to show line numbers (default: true)
  - `syntax` (string): Syntax highlighting mode (default: "text")

**Returns:**
- An editor ID string that can be used with other functions

### setText(editorId, text)

Sets the entire text content of the editor.

**Parameters:**
- `editorId` (string): The editor ID
- `text` (string): The text content to set

**Returns:**
- `true` if the text was set successfully

### getText(editorId)

Gets the entire text content of the editor.

**Parameters:**
- `editorId` (string): The editor ID

**Returns:**
- A string containing the editor's text content

### insertLine(editorId, lineNumber, text)

Inserts a new line at the specified position.

**Parameters:**
- `editorId` (string): The editor ID
- `lineNumber` (integer, optional): The line number where to insert the text. If not provided, inserts at the current cursor position.
- `text` (string): The text to insert

**Returns:**
- `true` if the line was inserted successfully

### deleteLine(editorId, lineNumber)

Deletes a line at the specified position.

**Parameters:**
- `editorId` (string): The editor ID
- `lineNumber` (integer, optional): The line number to delete. If not provided, deletes the line at the current cursor position.

**Returns:**
- `true` if the line was deleted successfully

### updateLine(editorId, lineNumber, text)

Updates a line at the specified position.

**Parameters:**
- `editorId` (string): The editor ID
- `lineNumber` (integer, optional): The line number to update. If not provided, updates the line at the current cursor position.
- `text` (string): The new text for the line

**Returns:**
- `true` if the line was updated successfully

### getLine(editorId, lineNumber)

Gets a line at the specified position.

**Parameters:**
- `editorId` (string): The editor ID
- `lineNumber` (integer, optional): The line number to get. If not provided, gets the line at the current cursor position.

**Returns:**
- A string containing the line's text

### getLineCount(editorId)

Gets the number of lines in the editor.

**Parameters:**
- `editorId` (string): The editor ID

**Returns:**
- An integer representing the number of lines

### render(editorId)

Renders the editor content as a string.

**Parameters:**
- `editorId` (string): The editor ID

**Returns:**
- A string containing the rendered editor content

### handleKeypress(editorId, key)

Handles a keypress in the editor.

**Parameters:**
- `editorId` (string): The editor ID
- `key` (string): The key that was pressed

**Returns:**
- `true` if the keypress was handled successfully

### save(editorId, filename)

Saves the editor content to a file.

**Parameters:**
- `editorId` (string): The editor ID
- `filename` (string, optional): The filename to save to. If not provided, uses the editor's current filename.

**Returns:**
- `true` if the file was saved successfully

### load(editorId, filename)

Loads a file into the editor.

**Parameters:**
- `editorId` (string): The editor ID
- `filename` (string): The filename to load

**Returns:**
- `true` if the file was loaded successfully

## Editor Modes

The editor supports three modes of operation, similar to Vim:

1. **Normal Mode**: Default mode for navigation and commands
2. **Insert Mode**: For inserting and editing text
3. **Command Mode**: For executing commands like save and quit

### Normal Mode Keys

- `h`: Move cursor left
- `j`: Move cursor down
- `k`: Move cursor up
- `l`: Move cursor right
- `0`: Move to beginning of line
- `$`: Move to end of line
- `G`: Move to last line
- `gg`: Move to first line
- `x`: Delete character under cursor
- `dd`: Delete current line
- `i`: Enter insert mode
- `o`: Insert new line below and enter insert mode
- `O`: Insert new line above and enter insert mode
- `:`: Enter command mode

### Insert Mode Keys

- `Escape`: Exit insert mode
- `Backspace`: Delete character before cursor
- `Enter`: Split line at cursor
- `ArrowLeft`, `ArrowRight`, `ArrowUp`, `ArrowDown`: Move cursor
- Any other key: Insert character at cursor

### Command Mode Keys

- `Escape`: Exit command mode
- `Enter`: Execute command
- `Backspace`: Delete character before cursor
- Any other key: Add character to command

### Command Mode Commands

- `w`: Write (save) the file
- `q`: Quit (will warn if there are unsaved changes)
- `q!`: Force quit without saving
- `wq`: Write and quit
- `set number`: Show line numbers
- `set nonumber`: Hide line numbers

## Examples

### Simple Notepad Application

```js
import editor
import term
import argparse

// Parse command line arguments
let parser = argparse.newParser("notepad", "A simple notepad application")
argparse.addArgument(parser, "file", {
    description: "File to edit",
    required: false
})
let args = argparse.parse(parser)

// Create a new editor
let ed = editor.newEditor({
    filename: args["file"],
    width: term.getSize().width,
    height: term.getSize().height - 2
})

// Main loop
term.clear()
while (true) {
    // Render the editor
    let content = editor.render(ed)
    term.clear()
    print(content)
    
    // Get keypress
    let key = term.getKey()
    
    // Check for exit key (Ctrl+Q)
    if (key == "Ctrl+Q") {
        // Check if there are unsaved changes
        if (editor.isModified(ed)) {
            term.println("There are unsaved changes. Press Y to quit anyway, or any other key to cancel.")
            let confirm = term.getKey()
            if (confirm.toLowerCase() != "y") {
                continue
            }
        }
        break
    }
    
    // Check for save key (Ctrl+S)
    if (key == "Ctrl+S") {
        let filename = args["file"]
        if (!filename) {
            term.println("Enter filename to save: ")
            filename = term.input()
            args["file"] = filename
        }
        
        editor.save(ed, filename)
        continue
    }
    
    // Handle the keypress in the editor
    editor.handleKeypress(ed, key)
}

term.clear()
term.println("Goodbye!")
```

### Configuration Editor

```js
import editor
import term
import os
import json

// Function to load a JSON configuration file
let loadConfig = func(filename) {
    if (!os.exists(filename)) {
        return {
            "server": {
                "host": "localhost",
                "port": 8080
            },
            "database": {
                "host": "localhost",
                "port": 5432,
                "user": "admin",
                "password": "password"
            },
            "logging": {
                "level": "info",
                "file": "app.log"
            }
        }
    }
    
    let content = open(filename)
    return JSON.parse(content)
}

// Function to save a JSON configuration file
let saveConfig = func(filename, config) {
    let content = JSON.stringify(config, null, 2)
    os.writeFile(filename, content)
}

// Function to edit a configuration section
let editSection = func(section, data) {
    // Convert the section data to a string
    let text = JSON.stringify(data, null, 2)
    
    // Create an editor
    let ed = editor.newEditor({
        width: term.getSize().width,
        height: term.getSize().height - 4,
        syntax: "json"
    })
    
    // Set the text
    editor.setText(ed, text)
    
    // Edit loop
    term.clear()
    while (true) {
        // Render the editor
        let content = editor.render(ed)
        term.clear()
        term.println("Editing " + section + " configuration", "#ffcc00")
        term.println("Press Ctrl+S to save, Ctrl+Q to cancel", "#88ff88")
        print(content)
        
        // Get keypress
        let key = term.getKey()
        
        // Check for exit key (Ctrl+Q)
        if (key == "Ctrl+Q") {
            return null
        }
        
        // Check for save key (Ctrl+S)
        if (key == "Ctrl+S") {
            let newText = editor.getText(ed)
            try {
                return JSON.parse(newText)
            } catch (e) {
                term.println("Error parsing JSON: " + e, "#ff5555")
                term.println("Press any key to continue", "#88ff88")
                term.getKey()
                continue
            }
        }
        
        // Handle the keypress in the editor
        editor.handleKeypress(ed, key)
    }
}

// Main function
let main = func() {
    let configFile = "config.json"
    let config = loadConfig(configFile)
    
    while (true) {
        // Display menu
        term.clear()
        term.println("=== Configuration Editor ===", "#ffcc00")
        term.println("1. Edit Server Configuration", "#88ff88")
        term.println("2. Edit Database Configuration", "#88ff88")
        term.println("3. Edit Logging Configuration", "#88ff88")
        term.println("4. Save Configuration", "#88ff88")
        term.println("5. Exit", "#88ff88")
        term.println("===========================", "#ffcc00")
        term.print("Enter your choice: ")
        
        let choice = term.input()
        
        if (choice == "1") {
            let result = editSection("server", config.server)
            if (result) {
                config.server = result
            }
        } else if (choice == "2") {
            let result = editSection("database", config.database)
            if (result) {
                config.database = result
            }
        } else if (choice == "3") {
            let result = editSection("logging", config.logging)
            if (result) {
                config.logging = result
            }
        } else if (choice == "4") {
            saveConfig(configFile, config)
            term.println("Configuration saved to " + configFile, "#88ff88")
            term.println("Press any key to continue", "#88ff88")
            term.getKey()
        } else if (choice == "5") {
            break
        }
    }
    
    term.clear()
    term.println("Goodbye!", "#88ff88")
}

main()
```

### Markdown Editor with Preview

```js
import editor
import term
import os
import markdown

// Function to convert markdown to HTML
let markdownToHtml = func(text) {
    // This is a simplified markdown converter
    let html = text
    
    // Convert headers
    html = html.replace(/^# (.+)$/gm, "<h1>$1</h1>")
    html = html.replace(/^## (.+)$/gm, "<h2>$1</h2>")
    html = html.replace(/^### (.+)$/gm, "<h3>$1</h3>")
    
    // Convert bold and italic
    html = html.replace(/\*\*(.+?)\*\*/g, "<strong>$1</strong>")
    html = html.replace(/\*(.+?)\*/g, "<em>$1</em>")
    
    // Convert lists
    html = html.replace(/^- (.+)$/gm, "<li>$1</li>")
    
    // Convert links
    html = html.replace(/\[(.+?)\]\((.+?)\)/g, "<a href=\"$2\">$1</a>")
    
    return html
}

// Function to display a preview of the markdown
let previewMarkdown = func(text) {
    let html = markdownToHtml(text)
    
    // Display the HTML preview
    term.clear()
    term.println("=== Markdown Preview ===", "#ffcc00")
    term.println(html)
    term.println("========================", "#ffcc00")
    term.println("Press any key to return to editor", "#88ff88")
    term.getKey()
}

// Main function
let main = func() {
    // Check command line arguments
    if (args.length < 2) {
        term.println("Usage: vint markdown_editor.vint <filename>", "#ff5555")
        exit(1)
    }
    
    let filename = args[1]
    
    // Create an editor
    let ed = editor.newEditor({
        filename: filename,
        width: term.getSize().width,
        height: term.getSize().height - 2,
        syntax: "markdown"
    })
    
    // Main loop
    term.clear()
    while (true) {
        // Render the editor
        let content = editor.render(ed)
        term.clear()
        term.println("Markdown Editor - " + filename, "#ffcc00")
        term.println("Ctrl+S: Save | Ctrl+P: Preview | Ctrl+Q: Quit", "#88ff88")
        print(content)
        
        // Get keypress
        let key = term.getKey()
        
        // Check for exit key (Ctrl+Q)
        if (key == "Ctrl+Q") {
            if (editor.isModified(ed)) {
                term.println("There are unsaved changes. Press Y to quit anyway, or any other key to cancel.")
                let confirm = term.getKey()
                if (confirm.toLowerCase() != "y") {
                    continue
                }
            }
            break
        }
        
        // Check for save key (Ctrl+S)
        if (key == "Ctrl+S") {
            editor.save(ed, filename)
            continue
        }
        
        // Check for preview key (Ctrl+P)
        if (key == "Ctrl+P") {
            let text = editor.getText(ed)
            previewMarkdown(text)
            continue
        }
        
        // Handle the keypress in the editor
        editor.handleKeypress(ed, key)
    }
    
    term.clear()
    term.println("Goodbye!", "#88ff88")
}

main()
```