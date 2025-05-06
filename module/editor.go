package module

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/vintlang/vintlang/object"
// )

// var EditorFunctions = map[string]object.ModuleFunction{}

// func init() {
// 	EditorFunctions["newEditor"] = newEditor
// 	EditorFunctions["setText"] = setText
// 	EditorFunctions["getText"] = getText
// 	EditorFunctions["insertLine"] = insertLine
// 	EditorFunctions["deleteLine"] = deleteLine
// 	EditorFunctions["updateLine"] = updateLine
// 	EditorFunctions["getLine"] = getLine
// 	EditorFunctions["getLineCount"] = getLineCount
// 	EditorFunctions["render"] = renderEditor
// 	EditorFunctions["handleKeypress"] = handleKeypress
// 	EditorFunctions["save"] = saveEditor
// 	EditorFunctions["load"] = loadEditor
// }

// // Editor structure to store editor state
// type editor struct {
// 	lines       []string
// 	cursorRow   int
// 	cursorCol   int
// 	scrollRow   int
// 	filename    string
// 	modified    bool
// 	mode        string // "normal", "insert", "command"
// 	statusMsg   string
// 	showLineNum bool
// 	syntax      string
// 	width       int
// 	height      int
// }

// // Map to store editors by ID
// var editors = make(map[string]*editor)

// // newEditor creates a new text editor
// func newEditor(args []object.Object, defs map[string]object.Object) object.Object {
// 	// Create a unique ID for the editor
// 	editorID := fmt.Sprintf("editor_%d", len(editors))

// 	// Create a new editor
// 	ed := &editor{
// 		lines:       []string{""},
// 		cursorRow:   0,
// 		cursorCol:   0,
// 		scrollRow:   0,
// 		filename:    "",
// 		modified:    false,
// 		mode:        "normal",
// 		statusMsg:   "Welcome to VintLang Editor",
// 		showLineNum: true,
// 		syntax:      "text",
// 		width:       80,
// 		height:      24,
// 	}

// 	// Process optional parameters
// 	if filename, ok := defs["filename"]; ok {
// 		if filenameStr, ok := filename.(*object.String); ok {
// 			ed.filename = filenameStr.Value
// 			// Try to load the file
// 			if content, err := loadFile(ed.filename); err == nil {
// 				ed.lines = strings.Split(content, "\n")
// 				if len(ed.lines) == 0 {
// 					ed.lines = []string{""}
// 				}
// 			}
// 		}
// 	}

// 	if width, ok := defs["width"]; ok {
// 		if widthInt, ok := width.(*object.Integer); ok {
// 			ed.width = int(widthInt.Value)
// 		}
// 	}

// 	if height, ok := defs["height"]; ok {
// 		if heightInt, ok := height.(*object.Integer); ok {
// 			ed.height = int(heightInt.Value)
// 		}
// 	}

// 	if showLineNum, ok := defs["showLineNum"]; ok {
// 		if showLineNumBool, ok := showLineNum.(*object.Boolean); ok {
// 			ed.showLineNum = showLineNumBool.Value
// 		}
// 	}

// 	if syntax, ok := defs["syntax"]; ok {
// 		if syntaxStr, ok := syntax.(*object.String); ok {
// 			ed.syntax = syntaxStr.Value
// 		}
// 	}

// 	// Store the editor
// 	editors[editorID] = ed

// 	return &object.String{Value: editorID}
// }

// // setText sets the entire text content of the editor
// func setText(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 2 {
// 		return &object.Error{Message: "setText requires exactly 2 arguments: editor ID and text"}
// 	}

// 	// Get editor ID
// 	editorID, ok := args[0].(*object.String)
// 	if !ok {
// 		return &object.Error{Message: "editor ID must be a string"}
// 	}

// 	// Get text
// 	text, ok := args[1].(*object.String)
// 	if !ok {
// 		return &object.Error{Message: "text must be a string"}
// 	}

// 	// Check if editor exists
// 	ed, exists := editors[editorID.Value]
// 	if !exists {
// 		return &object.Error{Message: fmt.Sprintf("editor '%s' not found", editorID.Value)}
// 	}

// 	// Set the text
// 	ed.lines = strings.Split(text.Value, "\n")
// 	if len(ed.lines) == 0 {
// 		ed.lines = []string{""}
// 	}

// 	// Reset cursor position
// 	ed.cursorRow = 0
// 	ed.cursorCol = 0
// 	ed.scrollRow = 0
// 	ed.modified = true

// 	return &object.Boolean{Value: true}
// }

// // getText gets the entire text content of the editor
// func getText(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 1 {
// 		return &object.Error{Message: "getText requires exactly 1 argument: editor ID"}
// 	}

// 	// Get editor ID
// 	editorID, ok := args[0].(*object.String)
// 	if !ok {
// 		return &object.Error{Message: "editor ID must be a string"}
// 	}

// 	// Check if editor exists
// 	ed, exists := editors[editorID.Value]
// 	if !exists {
// 		return &object.Error{Message: fmt.Sprintf("editor '%s' not found", editorID.Value)}
// 	}

// 	// Get the text
// 	text := strings.Join(ed.lines, "\n")

// 	return &object.String{Value: text}
// }

// // insertLine inserts a new line at the specified position
// func insertLine(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) < 2 || len(args) > 3 {
// 		return &object.Error{Message: "insertLine requires 2-3 arguments: editor ID, line number (optional), and text"}
// 	}

// 	// Get editor ID
// 	editorID, ok := args[0].(*object.String)
// 	if !ok {
// 		return &object.Error{Message: "editor ID must be a string"}
// 	}

// 	// Check if editor exists
// 	ed, exists := editors[editorID.Value]
// 	if !exists {
// 		return &object.Error{Message: fmt.Sprintf("editor '%s' not found", editorID.Value)}
// 	}

// 	var lineNum int
// 	var text string

// 	if len(args) == 2 {
// 		// Insert at current cursor position
// 		lineNum = ed.cursorRow
// 		textObj, ok := args[1].(*object.String)
// 		if !ok {
// 			return &object.Error{Message: "text must be a string"}
// 		}
// 		text = textObj.Value
// 	} else {
// 		// Insert at specified line number
// 		lineNumObj, ok := args[1].(*object.Integer)
// 		if !ok {
// 			return &object.Error{Message: "line number must be an integer"}
// 		}
// 		lineNum = int(lineNumObj.Value)

// 		textObj, ok := args[2].(*object.String)
// 		if !ok {
// 			return &object.Error{Message: "text must be a string"}
// 		}
// 		text = textObj.Value
// 	}

// 	// Validate line number
// 	if lineNum < 0 {
// 		lineNum = 0
// 	}
// 	if lineNum > len(ed.lines) {
// 		lineNum = len(ed.lines)
// 	}

// 	// Insert the line
// 	if lineNum == len(ed.lines) {
// 		ed.lines = append(ed.lines, text)
// 	} else {
// 		ed.lines = append(ed.lines[:lineNum+1], ed.lines[lineNum:]...)
// 		ed.lines[lineNum] = text
// 	}

// 	ed.modified = true

// 	return &object.Boolean{Value: true}
// }

// // deleteLine deletes a line at the specified position
// func deleteLine(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) < 1 || len(args) > 2 {
// 		return &object.Error{Message: "deleteLine requires 1-2 arguments: editor ID and optional line number"}
// 	}

// 	// Get editor ID
// 	editorID, ok := args[0].(*object.String)
// 	if !ok {
// 		return &object.Error{Message: "editor ID must be a string"}
// 	}

// 	// Check if editor exists
// 	ed, exists := editors[editorID.Value]
// 	if !exists {
// 		return &object.Error{Message: fmt.Sprintf("editor '%s' not found", editorID.Value)}
// 	}

// 	var lineNum int

// 	if len(args) == 1 {
// 		// Delete at current cursor position
// 		lineNum = ed.cursorRow
// 	} else {
// 		// Delete at specified line number
// 		lineNumObj, ok := args[1].(*object.Integer)
// 		if !ok {
// 			return &object.Error{Message: "line number must be an integer"}
// 		}
// 		lineNum = int(lineNumObj.Value)
// 	}

// 	// Validate line number
// 	if lineNum < 0 || lineNum >= len(ed.lines) {
// 		return &object.Error{Message: fmt.Sprintf("invalid line number: %d", lineNum)}
// 	}

// 	// Delete the line
// 	if len(ed.lines) == 1 {
// 		// If it's the only line, just clear it
// 		ed.lines[0] = ""
// 	} else {
// 		ed.lines = append(ed.lines[:lineNum], ed.lines[lineNum+1:]...)
// 	}

// 	// Adjust cursor if needed
// 	if ed.cursorRow >= len(ed.lines) {
// 		ed.cursorRow = len(ed.lines) - 1
// 	}
// 	if ed.cursorCol > len(ed.lines[ed.cursorRow]) {
// 		ed.cursorCol = len(ed.lines[ed.cursorRow])
// 	}

// 	ed.modified = true

// 	return &object.Boolean{Value: true}
// }

// // updateLine updates a line at the specified position
// func updateLine(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) < 2 || len(args) > 3 {
// 		return &object.Error{Message: "updateLine requires 2-3 arguments: editor ID, line number (optional), and text"}
// 	}

// 	// Get editor ID
// 	editorID, ok := args[0].(*object.String)
// 	if !ok {
// 		return &object.Error{Message: "editor ID must be a string"}
// 	}

// 	// Check if editor exists
// 	ed, exists := editors[editorID.Value]
// 	if !exists {
// 		return &object.Error{Message: fmt.Sprintf("editor '%s' not found", editorID.Value)}
// 	}

// 	var lineNum int
// 	var text string

// 	if len(args) == 2 {
// 		// Update at current cursor position
// 		lineNum = ed.cursorRow
// 		textObj, ok := args[1].(*object.String)
// 		if !ok {
// 			return &object.Error{Message: "text must be a string"}
// 		}
// 		text = textObj.Value
// 	} else {
// 		// Update at specified line number
// 		lineNumObj, ok := args[1].(*object.Integer)
// 		if !ok {
// 			return &object.Error{Message: "line number must be an integer"}
// 		}
// 		lineNum = int(lineNumObj.Value)

// 		textObj, ok := args[2].(*object.String)
// 		if !ok {
// 			return &object.Error{Message: "text must be a string"}
// 		}
// 		text = textObj.Value
// 	}

// 	// Validate line number
// 	if lineNum < 0 || lineNum >= len(ed.lines) {
// 		return &object.Error{Message: fmt.Sprintf("invalid line number: %d", lineNum)}
// 	}

// 	// Update the line
// 	ed.lines[lineNum] = text
// 	ed.modified = true

// 	return &object.Boolean{Value: true}
// }

// // getLine gets a line at the specified position
// func getLine(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) < 1 || len(args) > 2 {
// 		return &object.Error{Message: "getLine requires 1-2 arguments: editor ID and optional line number"}
// 	}

// 	// Get editor ID
// 	editorID, ok := args[0].(*object.String)
// 	if !ok {
// 		return &object.Error{Message: "editor ID must be a string"}
// 	}

// 	// Check if editor exists
// 	ed, exists := editors[editorID.Value]
// 	if !exists {
// 		return &object.Error{Message: fmt.Sprintf("editor '%s' not found", editorID.Value)}
// 	}

// 	var lineNum int

// 	if len(args) == 1 {
// 		// Get current cursor position
// 		lineNum = ed.cursorRow
// 	} else {
// 		// Get specified line number
// 		lineNumObj, ok := args[1].(*object.Integer)
// 		if !ok {
// 			return &object.Error{Message: "line number must be an integer"}
// 		}
// 		lineNum = int(lineNumObj.Value)
// 	}

// 	// Validate line number
// 	if lineNum < 0 || lineNum >= len(ed.lines) {
// 		return &object.Error{Message: fmt.Sprintf("invalid line number: %d", lineNum)}
// 	}

// 	// Get the line
// 	return &object.String{Value: ed.lines[lineNum]}
// }

// // getLineCount gets the number of lines in the editor
// func getLineCount(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 1 {
// 		return &object.Error{Message: "getLineCount requires exactly 1 argument: editor ID"}
// 	}

// 	// Get editor ID
// 	editorID, ok := args[0].(*object.String)
// 	if !ok {
// 		return &object.Error{Message: "editor ID must be a string"}
// 	}

// 	// Check if editor exists
// 	ed, exists := editors[editorID.Value]
// 	if !exists {
// 		return &object.Error{Message: fmt.Sprintf("editor '%s' not found", editorID.Value)}
// 	}

// 	// Get the line count
// 	return &object.Integer{Value: int64(len(ed.lines))}
// }

// // renderEditor renders the editor content as a string
// func renderEditor(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 1 {
// 		return &object.Error{Message: "render requires exactly 1 argument: editor ID"}
// 	}

// 	// Get editor ID
// 	editorID, ok := args[0].(*object.String)
// 	if !ok {
// 		return &object.Error{Message: "editor ID must be a string"}
// 	}

// 	// Check if editor exists
// 	ed, exists := editors[editorID.Value]
// 	if !exists {
// 		return &object.Error{Message: fmt.Sprintf("editor '%s' not found", editorID.Value)}
// 	}

// 	// Render the editor
// 	var sb strings.Builder

// 	// Calculate visible lines
// 	startRow := ed.scrollRow
// 	endRow := ed.scrollRow + ed.height - 2 // -2 for status bar and message line
// 	if endRow >= len(ed.lines) {
// 		endRow = len(ed.lines)
// 	}

// 	// Render visible lines
// 	for i := startRow; i < endRow; i++ {
// 		// Add line number if enabled
// 		if ed.showLineNum {
// 			sb.WriteString(fmt.Sprintf("%3d │ ", i+1))
// 		}

// 		// Add line content
// 		line := ed.lines[i]
// 		if len(line) > ed.width {
// 			line = line[:ed.width-3] + "..."
// 		}
// 		sb.WriteString(line)
// 		sb.WriteString("\n")
// 	}

// 	// Add status line
// 	sb.WriteString("\n")
// 	statusLine := fmt.Sprintf(" %s | %s | Line %d/%d | Col %d | %s", 
// 		ed.filename, 
// 		ed.mode, 
// 		ed.cursorRow+1, 
// 		len(ed.lines), 
// 		ed.cursorCol+1,
// 		ed.modified ? "modified" : "saved")

// 	if len(statusLine) > ed.width {
// 		statusLine = statusLine[:ed.width-3] + "..."
// 	}
// 	sb.WriteString(statusLine)

// 	// Add message line
// 	if ed.statusMsg != "" {
// 		sb.WriteString("\n")
// 		sb.WriteString(" " + ed.statusMsg)
// 	}

// 	return &object.String{Value: sb.String()}
// }

// // handleKeypress handles a keypress in the editor
// func handleKeypress(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) != 2 {
// 		return &object.Error{Message: "handleKeypress requires exactly 2 arguments: editor ID and key"}
// 	}

// 	// Get editor ID
// 	editorID, ok := args[0].(*object.String)
// 	if !ok {
// 		return &object.Error{Message: "editor ID must be a string"}
// 	}

// 	// Get key
// 	key, ok := args[1].(*object.String)
// 	if !ok {
// 		return &object.Error{Message: "key must be a string"}
// 	}

// 	// Check if editor exists
// 	ed, exists := editors[editorID.Value]
// 	if !exists {
// 		return &object.Error{Message: fmt.Sprintf("editor '%s' not found", editorID.Value)}
// 	}

// 	// Handle the keypress based on the mode
// 	switch ed.mode {
// 	case "normal":
// 		return handleNormalModeKey(ed, key.Value)
// 	case "insert":
// 		return handleInsertModeKey(ed, key.Value)
// 	case "command":
// 		return handleCommandModeKey(ed, key.Value)
// 	default:
// 		return &object.Error{Message: fmt.Sprintf("unknown mode: %s", ed.mode)}
// 	}
// }

// // handleNormalModeKey handles a keypress in normal mode
// func handleNormalModeKey(ed *editor, key string) object.Object {
// 	switch key {
// 	case "i":
// 		// Enter insert mode
// 		ed.mode = "insert"
// 		ed.statusMsg = "-- INSERT --"
// 	case ":":
// 		// Enter command mode
// 		ed.mode = "command"
// 		ed.statusMsg = ":"
// 	case "h":
// 		// Move cursor left
// 		if ed.cursorCol > 0 {
// 			ed.cursorCol--
// 		}
// 	case "j":
// 		// Move cursor down
// 		if ed.cursorRow < len(ed.lines)-1 {
// 			ed.cursorRow++
// 			if ed.cursorCol > len(ed.lines[ed.cursorRow]) {
// 				ed.cursorCol = len(ed.lines[ed.cursorRow])
// 			}
// 		}
// 	case "k":
// 		// Move cursor up
// 		if ed.cursorRow > 0 {
// 			ed.cursorRow--
// 			if ed.cursorCol > len(ed.lines[ed.cursorRow]) {
// 				ed.cursorCol = len(ed.lines[ed.cursorRow])
// 			}
// 		}
// 	case "l":
// 		// Move cursor right
// 		if ed.cursorCol < len(ed.lines[ed.cursorRow]) {
// 			ed.cursorCol++
// 		}
// 	case "0":
// 		// Move to beginning of line
// 		ed.cursorCol = 0
// 	case "$":
// 		// Move to end of line
// 		ed.cursorCol = len(ed.lines[ed.cursorRow])
// 	case "G":
// 		// Move to last line
// 		ed.cursorRow = len(ed.lines) - 1
// 		if ed.cursorCol > len(ed.lines[ed.cursorRow]) {
// 			ed.cursorCol = len(ed.lines[ed.cursorRow])
// 		}
// 	case "gg":
// 		// Move to first line
// 		ed.cursorRow = 0
// 		if ed.cursorCol > len(ed.lines[ed.cursorRow]) {
// 			ed.cursorCol = len(ed.lines[ed.cursorRow])
// 		}
// 	case "x":
// 		// Delete character under cursor
// 		if ed.cursorCol < len(ed.lines[ed.cursorRow]) {
// 			line := ed.lines[ed.cursorRow]
// 			ed.lines[ed.cursorRow] = line[:ed.cursorCol] + line[ed.cursorCol+1:]
// 			ed.modified = true
// 		}
// 	case "dd":
// 		// Delete current line
// 		if len(ed.lines) > 1 {
// 			ed.lines = append(ed.lines[:ed.cursorRow], ed.lines[ed.cursorRow+1:]...)
// 			if ed.cursorRow >= len(ed.lines) {
// 				ed.cursorRow = len(ed.lines) - 1
// 			}
// 			if ed.cursorCol > len(ed.lines[ed.cursorRow]) {
// 				ed.cursorCol = len(ed.lines[ed.cursorRow])
// 			}
// 		} else {
// 			// If it's the only line, just clear it
// 			ed.lines[0] = ""
// 			ed.cursorCol = 0
// 		}
// 		ed.modified = true
// 	case "o":
// 		// Insert new line below and enter insert mode
// 		ed.lines = append(ed.lines[:ed.cursorRow+1], ed.lines[ed.cursorRow:]...)
// 		ed.lines[ed.cursorRow+1] = ""
// 		ed.cursorRow++
// 		ed.cursorCol = 0
// 		ed.mode = "insert"
// 		ed.statusMsg = "-- INSERT --"
// 		ed.modified = true
// 	case "O":
// 		// Insert new line above and enter insert mode
// 		ed.lines = append(ed.lines[:ed.cursorRow], append([]string{""}, ed.lines[ed.cursorRow:]...)...)
// 		ed.cursorCol = 0
// 		ed.mode = "insert"
// 		ed.statusMsg = "-- INSERT --"
// 		ed.modified = true
// 	default:
// 		// Unknown key
// 		ed.statusMsg = fmt.Sprintf("Unknown key: %s", key)
// 	}

// 	// Adjust scroll if needed
// 	if ed.cursorRow < ed.scrollRow {
// 		ed.scrollRow = ed.cursorRow
// 	} else if ed.cursorRow >= ed.scrollRow+ed.height-2 {
// 		ed.scrollRow = ed.cursorRow - (ed.height - 3)
// 	}

// 	return &object.Boolean{Value: true}
// }

// // handleInsertModeKey handles a keypress in insert mode
// func handleInsertModeKey(ed *editor, key string) object.Object {
// 	switch key {
// 	case "Escape":
// 		// Exit insert mode
// 		ed.mode = "normal"
// 		ed.statusMsg = ""
// 	case "Backspace":
// 		// Delete character before cursor
// 		if ed.cursorCol > 0 {
// 			line := ed.lines[ed.cursorRow]
// 			ed.lines[ed.cursorRow] = line[:ed.cursorCol-1] + line[ed.cursorCol:]
// 			ed.cursorCol--
// 			ed.modified = true
// 		} else if ed.cursorRow > 0 {
// 			// At beginning of line, join with previous line
// 			prevLine := ed.lines[ed.cursorRow-1]
// 			currLine := ed.lines[ed.cursorRow]
// 			ed.cursorCol = len(prevLine)
// 			ed.lines[ed.cursorRow-1] = prevLine + currLine
// 			ed.lines = append(ed.lines[:ed.cursorRow], ed.lines[ed.cursorRow+1:]...)
// 			ed.cursorRow--
// 			ed.modified = true
// 		}
// 	case "Enter":
// 		// Split line at cursor
// 		line := ed.lines[ed.cursorRow]
// 		ed.lines[ed.cursorRow] = line[:ed.cursorCol]
// 		ed.lines = append(ed.lines[:ed.cursorRow+1], ed.lines[ed.cursorRow:]...)
// 		ed.lines[ed.cursorRow+1] = line[ed.cursorCol:]
// 		ed.cursorRow++
// 		ed.cursorCol = 0
// 		ed.modified = true
// 	case "ArrowLeft":
// 		// Move cursor left
// 		if ed.cursorCol > 0 {
// 			ed.cursorCol--
// 		}
// 	case "ArrowRight":
// 		// Move cursor right
// 		if ed.cursorCol < len(ed.lines[ed.cursorRow]) {
// 			ed.cursorCol++
// 		}
// 	case "ArrowUp":
// 		// Move cursor up
// 		if ed.cursorRow > 0 {
// 			ed.cursorRow--
// 			if ed.cursorCol > len(ed.lines[ed.cursorRow]) {
// 				ed.cursorCol = len(ed.lines[ed.cursorRow])
// 			}
// 		}
// 	case "ArrowDown":
// 		// Move cursor down
// 		if ed.cursorRow < len(ed.lines)-1 {
// 			ed.cursorRow++
// 			if ed.cursorCol > len(ed.lines[ed.cursorRow]) {
// 				ed.cursorCol = len(ed.lines[ed.cursorRow])
// 			}
// 		}
// 	default:
// 		// Insert character at cursor
// 		if len(key) == 1 {
// 			line := ed.lines[ed.cursorRow]
// 			ed.lines[ed.cursorRow] = line[:ed.cursorCol] + key + line[ed.cursorCol:]
// 			ed.cursorCol++
// 			ed.modified = true
// 		}
// 	}

// 	// Adjust scroll if needed
// 	if ed.cursorRow < ed.scrollRow {
// 		ed.scrollRow = ed.cursorRow
// 	} else if ed.cursorRow >= ed.scrollRow+ed.height-2 {
// 		ed.scrollRow = ed.cursorRow - (ed.height - 3)
// 	}

// 	return &object.Boolean{Value: true}
// }

// // handleCommandModeKey handles a keypress in command mode
// func handleCommandModeKey(ed *editor, key string) object.Object {
// 	switch key {
// 	case "Escape":
// 		// Exit command mode
// 		ed.mode = "normal"
// 		ed.statusMsg = ""
// 	case "Enter":
// 		// Execute command
// 		command := ed.statusMsg[1:] // Remove the leading ":"
// 		executeCommand(ed, command)
// 		ed.mode = "normal"
// 		ed.statusMsg = ""
// 	case "Backspace":
// 		// Delete character before cursor
// 		if len(ed.statusMsg) > 1 {
// 			ed.statusMsg = ed.statusMsg[:len(ed.statusMsg)-1]
// 		}
// 	default:
// 		// Add character to command
// 		if len(key) == 1 {
// 			ed.statusMsg += key
// 		}
// 	}

// 	return &object.Boolean{Value: true}
// }

// // executeCommand executes a command in the editor
// func executeCommand(ed *editor, command string) {
// 	command = strings.TrimSpace(command)

// 	if command == "q" {
// 		// Quit
// 		if ed.modified {
// 			ed.statusMsg = "No write since last change (add ! to override)"
// 		} else {
// 			// TODO: Implement actual quitting
// 			ed.statusMsg = "Quit"
// 		}
// 	} else if command == "q!" {
// 		// Force quit
// 		// TODO: Implement actual quitting
// 		ed.statusMsg = "Force quit"
// 	} else if command == "w" {
// 		// Write
// 		if ed.filename == "" {
// 			ed.statusMsg = "No file name"
// 		} else {
// 			saveFile(ed.filename, strings.Join(ed.lines, "\n"))
// 			ed.modified = false
// 			ed.statusMsg = fmt.Sprintf("\"%s\" written", ed.filename)
// 		}
// 	} else if strings.HasPrefix(command, "w ") {
// 		// Write to file
// 		filename := strings.TrimSpace(command[2:])
// 		if filename != "" {
// 			saveFile(filename, strings.Join(ed.lines, "\n"))
// 			ed.filename = filename
// 			ed.modified = false
// 			ed.statusMsg = fmt.Sprintf("\"%s\" written", ed.filename)
// 		}
// 	} else if command == "wq" {
// 		// Write and quit
// 		if ed.filename == "" {
// 			ed.statusMsg = "No file name"
// 		} else {
// 			saveFile(ed.filename, strings.Join(ed.lines, "\n"))
// 			ed.modified = false
// 			ed.statusMsg = fmt.Sprintf("\"%s\" written", ed.filename)
// 			// TODO: Implement actual quitting
// 		}
// 	} else if strings.HasPrefix(command, "set ") {
// 		// Set option
// 		option := strings.TrimSpace(command[4:])
// 		if option == "number" || option == "nu" {
// 			ed.showLineNum = true
// 			ed.statusMsg = "show line numbers"
// 		} else if option == "nonumber" || option == "nonu" {
// 			ed.showLineNum = false
// 			ed.statusMsg = "hide line numbers"
// 		} else {
// 			ed.statusMsg = fmt.Sprintf("Unknown option: %s", option)
// 		}
// 	} else {
// 		// Unknown command
// 		ed.statusMsg = fmt.Sprintf("Unknown command: %s", command)
// 	}
// }

// // saveEditor saves the editor content to a file
// func saveEditor(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(args) < 1 || len(args) > 2 {
// 		return &object.Error{Message: "save requires 1-2 arguments: editor ID and optional filename"}
// 	}

// 	// Get editor ID
// 	editorID, ok := args[0].(*object.String)
// 	if !ok {
// 		return &object.Error{Message: "editor ID must be a string"}
// 	}

//  // Check if editor exists
//  	ed, exists := editors[editorID.Value]
//  	if !exists {
//  		return &object.Error{Message: fmt.Sprintf("editor '%s' not found", editorID.Value)}
//  	}

//  	// Get filename
//  	var filename string
//  	if len(args) == 2 {
//  		filenameObj, ok := args[1].(*object.String)
//  		if !ok {
//  			return &object.Error{Message: "filename must be a string"}
//  		}
//  		filename = filenameObj.Value
//  		ed.filename = filename
//  	} else if ed.filename != "" {
//  		filename = ed.filename
//  	} else {
//  		return &object.Error{Message: "no filename specified"}
//  	}

//  	// Save the file
//  	err := saveFile(filename, strings.Join(ed.lines, "\n"))
//  	if err != nil {
//  		return &object.Error{Message: fmt.Sprintf("failed to save file: %v", err)}
//  	}

//  	ed.modified = false
//  	ed.statusMsg = fmt.Sprintf("\"%s\" written", filename)

//  	return &object.Boolean{Value: true}
//  }

//  // loadEditor loads a file into the editor
//  func loadEditor(args []object.Object, defs map[string]object.Object) object.Object {
//  	if len(args) != 2 {
//  		return &object.Error{Message: "load requires exactly 2 arguments: editor ID and filename"}
//  	}

//  	// Get editor ID
//  	editorID, ok := args[0].(*object.String)
//  	if !ok {
//  		return &object.Error{Message: "editor ID must be a string"}
//  	}

//  	// Get filename
//  	filename, ok := args[1].(*object.String)
//  	if !ok {
//  		return &object.Error{Message: "filename must be a string"}
//  	}

//  	// Check if editor exists
//  	ed, exists := editors[editorID.Value]
//  	if !exists {
//  		return &object.Error{Message: fmt.Sprintf("editor '%s' not found", editorID.Value)}
//  	}

//  	// Load the file
//  	content, err := loadFile(filename.Value)
//  	if err != nil {
//  		return &object.Error{Message: fmt.Sprintf("failed to load file: %v", err)}
//  	}

//  	// Set the content
//  	ed.lines = strings.Split(content, "\n")
//  	if len(ed.lines) == 0 {
//  		ed.lines = []string{""}
//  	}

//  	// Reset cursor position
//  	ed.cursorRow = 0
//  	ed.cursorCol = 0
//  	ed.scrollRow = 0
//  	ed.modified = false
//  	ed.filename = filename.Value
//  	ed.statusMsg = fmt.Sprintf("\"%s\" loaded", filename.Value)

//  	return &object.Boolean{Value: true}
//  }

//  // saveFile saves content to a file
//  func saveFile(filename string, content string) error {
//  	return os.WriteFile(filename, []byte(content), 0644)
//  }

//  // loadFile loads content from a file
//  func loadFile(filename string) (string, error) {
//  	data, err := os.ReadFile(filename)
//  	if err != nil {
//  		return "", err
//  	}
//  	return string(data), nil
//  }
