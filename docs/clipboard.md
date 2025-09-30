# Using the Clipboard Module in Vint

The **Clipboard** module in Vint provides functionality to interact with the system clipboard, allowing you to read, write, and manage clipboard content programmatically.

## Available Functions

### `clipboard.write(text)`
Writes text content to the system clipboard.

**Parameters:**
- `text` - The text to write to clipboard (string, integer, float, or boolean)

**Returns:** 
- `true` if successful
- Error object if operation fails

### `clipboard.read()`
Reads the current text content from the system clipboard.

**Parameters:** None

**Returns:**
- String containing the clipboard content
- Error object if operation fails

### `clipboard.clear()`
Clears the clipboard by writing an empty string to it.

**Parameters:** None

**Returns:**
- `true` if successful
- Error object if operation fails

### `clipboard.hasContent()`
Checks if the clipboard contains any content.

**Parameters:** None

**Returns:**
- `true` if clipboard has content (non-empty)
- `false` if clipboard is empty
- Error object if operation fails

### `clipboard.all()`
Returns all clipboard data as an array. Since the system clipboard typically holds only one piece of text at a time, this returns an array containing the current clipboard content.

**Parameters:** None

**Returns:**
- Array containing clipboard content (empty array if clipboard is empty)
- Error object if operation fails

## Examples

### Basic Usage

```js
import clipboard

// Write text to clipboard
clipboard.write("Hello, World!")

// Read from clipboard
let content = clipboard.read()
print("Clipboard content:", content)

// Check if clipboard has content
if clipboard.hasContent() {
    print("Clipboard has content")
} else {
    print("Clipboard is empty")
}

// Get all clipboard data as array
let allData = clipboard.all()
print("All clipboard data:", allData)

// Clear clipboard
clipboard.clear()
print("Clipboard cleared")

// Check all data after clear
let emptyData = clipboard.all()
print("Clipboard after clear:", emptyData)
```

### Writing Different Data Types

```js
import clipboard

// Write string
clipboard.write("Hello World")

// Write number
clipboard.write(42)

// Write float
clipboard.write(3.14159)

// Write boolean
clipboard.write(true)
```

### Error Handling

```js
import clipboard

// The clipboard functions return error objects if operations fail
let result = clipboard.write("test")
if result.type == "ERROR" {
    print("Failed to write to clipboard:", result.message)
}

let content = clipboard.read()
if content.type == "ERROR" {
    print("Failed to read from clipboard:", content.message)
} else {
    print("Clipboard content:", content)
}
```

### Practical Example: Clipboard Manager

```js
import clipboard

// Save current clipboard content before modifying
let backup = clipboard.read()

// Process some text
let processed_text = "Processed: " + backup

// Write processed text back to clipboard
clipboard.write(processed_text)

print("Original:", backup)
print("Modified clipboard content:", clipboard.read())
```

## Platform Support

The clipboard module works across different platforms:
- **Windows**: Uses Windows Clipboard API
- **macOS**: Uses NSPasteboard
- **Linux**: Uses X11 clipboard (requires X11 environment)

## Notes

- The clipboard module requires appropriate system permissions to access the clipboard
- On some systems, clipboard access might require the application to be running in a graphical environment
- The module handles text content only; binary data is not supported
- Error handling is important as clipboard operations can fail due to system restrictions or lack of permissions