package module

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/vintlang/vintlang/object"
)

var ClipboardFunctions = map[string]object.ModuleFunction{}

func init() {
	ClipboardFunctions["write"] = clipboardWrite
	ClipboardFunctions["read"] = clipboardRead
	ClipboardFunctions["clear"] = clipboardClear
	ClipboardFunctions["hasContent"] = clipboardHasContent
	ClipboardFunctions["all"] = clipboardAll
}

func clipboardWrite(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"clipboard", "write",
			"1 argument (text to write)",
			fmt.Sprintf("%d arguments", len(args)),
			`clipboard.write("Hello World") -> writes text to clipboard`,
		)
	}

	var text string
	switch arg := args[0].(type) {
	case *object.String:
		text = arg.Value
	case *object.Integer:
		text = fmt.Sprintf("%d", arg.Value)
	case *object.Float:
		text = fmt.Sprintf("%f", arg.Value)
	case *object.Boolean:
		if arg.Value {
			text = "true"
		} else {
			text = "false"
		}
	default:
		return ErrorMessage(
			"clipboard", "write",
			"string, integer, float, or boolean argument",
			string(arg.Type()),
			`clipboard.write("text") or clipboard.write(123)`,
		)
	}

	err := clipboard.WriteAll(text)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to write to clipboard: %s", err.Error())}
	}

	return &object.Boolean{Value: true}
}

func clipboardRead(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 {
		return ErrorMessage(
			"clipboard", "read",
			"No arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`clipboard.read() -> returns clipboard text content`,
		)
	}

	text, err := clipboard.ReadAll()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to read from clipboard: %s", err.Error())}
	}

	return &object.String{Value: text}
}

func clipboardClear(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 {
		return ErrorMessage(
			"clipboard", "clear",
			"No arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`clipboard.clear() -> clears clipboard content`,
		)
	}

	err := clipboard.WriteAll("")
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to clear clipboard: %s", err.Error())}
	}

	return &object.Boolean{Value: true}
}

func clipboardHasContent(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 {
		return ErrorMessage(
			"clipboard", "hasContent",
			"No arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`clipboard.hasContent() -> returns true if clipboard has content`,
		)
	}

	text, err := clipboard.ReadAll()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to read from clipboard: %s", err.Error())}
	}

	return &object.Boolean{Value: len(text) > 0}
}

func clipboardAll(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 {
		return ErrorMessage(
			"clipboard", "all",
			"No arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`clipboard.all() -> returns array with current clipboard content`,
		)
	}

	text, err := clipboard.ReadAll()
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to read from clipboard: %s", err.Error())}
	}

	// We create an array with the current clipboard content
	// Since system clipboard only holds one item at a time, we return array with single item
	elements := []object.Object{}
	if len(text) > 0 {
		elements = append(elements, &object.String{Value: text})
	}

	return &object.Array{Elements: elements}
}
