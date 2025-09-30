package module

import (
	"testing"

	"github.com/vintlang/vintlang/object"
)

func TestClipboardWrite(t *testing.T) {
	// Test writing a string
	args := []object.Object{
		&object.String{Value: "test content"},
	}
	result := clipboardWrite(args, map[string]object.Object{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if boolResult, ok := result.(*object.Boolean); !ok || !boolResult.Value {
		t.Errorf("Expected true boolean result, got %T with value %v", result, boolResult.Value)
	}
}

func TestClipboardWriteInteger(t *testing.T) {
	// Test writing an integer
	args := []object.Object{
		&object.Integer{Value: 42},
	}
	result := clipboardWrite(args, map[string]object.Object{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if boolResult, ok := result.(*object.Boolean); !ok || !boolResult.Value {
		t.Errorf("Expected true boolean result, got %T", result)
	}
}

func TestClipboardWriteFloat(t *testing.T) {
	// Test writing a float
	args := []object.Object{
		&object.Float{Value: 3.14159},
	}
	result := clipboardWrite(args, map[string]object.Object{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if boolResult, ok := result.(*object.Boolean); !ok || !boolResult.Value {
		t.Errorf("Expected true boolean result, got %T", result)
	}
}

func TestClipboardWriteBoolean(t *testing.T) {
	// Test writing a boolean
	args := []object.Object{
		&object.Boolean{Value: true},
	}
	result := clipboardWrite(args, map[string]object.Object{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if boolResult, ok := result.(*object.Boolean); !ok || !boolResult.Value {
		t.Errorf("Expected true boolean result, got %T", result)
	}
}

func TestClipboardWriteInvalidArgs(t *testing.T) {
	// Test with no arguments
	result := clipboardWrite([]object.Object{}, map[string]object.Object{})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for no arguments, got %T", result)
	}

	// Test with too many arguments
	result = clipboardWrite([]object.Object{
		&object.String{Value: "test"},
		&object.String{Value: "extra"},
	}, map[string]object.Object{})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for too many arguments, got %T", result)
	}
}

func TestClipboardRead(t *testing.T) {
	// First write something to clipboard
	writeArgs := []object.Object{
		&object.String{Value: "test read content"},
	}
	clipboardWrite(writeArgs, map[string]object.Object{})

	// Now read it back
	result := clipboardRead([]object.Object{}, map[string]object.Object{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if stringResult, ok := result.(*object.String); !ok {
		t.Errorf("Expected string result, got %T", result)
	} else if stringResult.Value != "test read content" {
		t.Errorf("Expected 'test read content', got '%s'", stringResult.Value)
	}
}

func TestClipboardReadInvalidArgs(t *testing.T) {
	// Test with arguments (should accept none)
	result := clipboardRead([]object.Object{
		&object.String{Value: "unexpected"},
	}, map[string]object.Object{})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for arguments, got %T", result)
	}
}

func TestClipboardHasContent(t *testing.T) {
	// First write something
	writeArgs := []object.Object{
		&object.String{Value: "content check"},
	}
	clipboardWrite(writeArgs, map[string]object.Object{})

	// Check if has content
	result := clipboardHasContent([]object.Object{}, map[string]object.Object{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if boolResult, ok := result.(*object.Boolean); !ok {
		t.Errorf("Expected boolean result, got %T", result)
	} else if !boolResult.Value {
		t.Errorf("Expected true (clipboard has content), got false")
	}
}

func TestClipboardClear(t *testing.T) {
	// First write something
	writeArgs := []object.Object{
		&object.String{Value: "content to clear"},
	}
	clipboardWrite(writeArgs, map[string]object.Object{})

	// Clear clipboard
	result := clipboardClear([]object.Object{}, map[string]object.Object{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if boolResult, ok := result.(*object.Boolean); !ok || !boolResult.Value {
		t.Errorf("Expected true boolean result, got %T", result)
	}

	// Check that clipboard is now empty
	hasContentResult := clipboardHasContent([]object.Object{}, map[string]object.Object{})
	if boolResult, ok := hasContentResult.(*object.Boolean); !ok {
		t.Errorf("Expected boolean result from hasContent, got %T", hasContentResult)
	} else if boolResult.Value {
		t.Errorf("Expected false (clipboard should be empty after clear), got true")
	}
}

func TestClipboardClearInvalidArgs(t *testing.T) {
	// Test with arguments (should accept none)
	result := clipboardClear([]object.Object{
		&object.String{Value: "unexpected"},
	}, map[string]object.Object{})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for arguments, got %T", result)
	}
}

func TestClipboardHasContentInvalidArgs(t *testing.T) {
	// Test with arguments (should accept none)
	result := clipboardHasContent([]object.Object{
		&object.String{Value: "unexpected"},
	}, map[string]object.Object{})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for arguments, got %T", result)
	}
}