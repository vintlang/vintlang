package module

import (
	"testing"

	"github.com/vintlang/vintlang/object"
)

func TestClipboardWrite(t *testing.T) {
	// Test writing a string
	args := []object.VintObject{
		&object.String{Value: "test content"},
	}
	result := clipboardWrite(args, map[string]object.VintObject{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if boolResult, ok := result.(*object.Boolean); !ok || !boolResult.Value {
		t.Errorf("Expected true boolean result, got %T with value %v", result, boolResult.Value)
	}
}

func TestClipboardWriteInteger(t *testing.T) {
	// Test writing an integer
	args := []object.VintObject{
		&object.Integer{Value: 42},
	}
	result := clipboardWrite(args, map[string]object.VintObject{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if boolResult, ok := result.(*object.Boolean); !ok || !boolResult.Value {
		t.Errorf("Expected true boolean result, got %T", result)
	}
}

func TestClipboardWriteFloat(t *testing.T) {
	// Test writing a float
	args := []object.VintObject{
		&object.Float{Value: 3.14159},
	}
	result := clipboardWrite(args, map[string]object.VintObject{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if boolResult, ok := result.(*object.Boolean); !ok || !boolResult.Value {
		t.Errorf("Expected true boolean result, got %T", result)
	}
}

func TestClipboardWriteBoolean(t *testing.T) {
	// Test writing a boolean
	args := []object.VintObject{
		&object.Boolean{Value: true},
	}
	result := clipboardWrite(args, map[string]object.VintObject{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if boolResult, ok := result.(*object.Boolean); !ok || !boolResult.Value {
		t.Errorf("Expected true boolean result, got %T", result)
	}
}

func TestClipboardWriteInvalidArgs(t *testing.T) {
	// Test with no arguments
	result := clipboardWrite([]object.VintObject{}, map[string]object.VintObject{})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for no arguments, got %T", result)
	}

	// Test with too many arguments
	result = clipboardWrite([]object.VintObject{
		&object.String{Value: "test"},
		&object.String{Value: "extra"},
	}, map[string]object.VintObject{})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for too many arguments, got %T", result)
	}
}

func TestClipboardRead(t *testing.T) {
	// First write something to clipboard
	writeArgs := []object.VintObject{
		&object.String{Value: "test read content"},
	}
	clipboardWrite(writeArgs, map[string]object.VintObject{})

	// Now read it back
	result := clipboardRead([]object.VintObject{}, map[string]object.VintObject{})

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
	result := clipboardRead([]object.VintObject{
		&object.String{Value: "unexpected"},
	}, map[string]object.VintObject{})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for arguments, got %T", result)
	}
}

func TestClipboardHasContent(t *testing.T) {
	// First write something
	writeArgs := []object.VintObject{
		&object.String{Value: "content check"},
	}
	clipboardWrite(writeArgs, map[string]object.VintObject{})

	// Check if has content
	result := clipboardHasContent([]object.VintObject{}, map[string]object.VintObject{})

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
	writeArgs := []object.VintObject{
		&object.String{Value: "content to clear"},
	}
	clipboardWrite(writeArgs, map[string]object.VintObject{})

	// Clear clipboard
	result := clipboardClear([]object.VintObject{}, map[string]object.VintObject{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if boolResult, ok := result.(*object.Boolean); !ok || !boolResult.Value {
		t.Errorf("Expected true boolean result, got %T", result)
	}

	// Check that clipboard is now empty
	hasContentResult := clipboardHasContent([]object.VintObject{}, map[string]object.VintObject{})
	if boolResult, ok := hasContentResult.(*object.Boolean); !ok {
		t.Errorf("Expected boolean result from hasContent, got %T", hasContentResult)
	} else if boolResult.Value {
		t.Errorf("Expected false (clipboard should be empty after clear), got true")
	}
}

func TestClipboardClearInvalidArgs(t *testing.T) {
	// Test with arguments (should accept none)
	result := clipboardClear([]object.VintObject{
		&object.String{Value: "unexpected"},
	}, map[string]object.VintObject{})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for arguments, got %T", result)
	}
}

func TestClipboardHasContentInvalidArgs(t *testing.T) {
	// Test with arguments (should accept none)
	result := clipboardHasContent([]object.VintObject{
		&object.String{Value: "unexpected"},
	}, map[string]object.VintObject{})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for arguments, got %T", result)
	}
}

func TestClipboardAll(t *testing.T) {
	// First write something to clipboard
	writeArgs := []object.VintObject{
		&object.String{Value: "test all content"},
	}
	clipboardWrite(writeArgs, map[string]object.VintObject{})

	// Test the all method
	result := clipboardAll([]object.VintObject{}, map[string]object.VintObject{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if arrayResult, ok := result.(*object.Array); !ok {
		t.Errorf("Expected array result, got %T", result)
	} else if len(arrayResult.Elements) != 1 {
		t.Errorf("Expected array with 1 element, got %d elements", len(arrayResult.Elements))
	} else if stringElement, ok := arrayResult.Elements[0].(*object.String); !ok {
		t.Errorf("Expected string element, got %T", arrayResult.Elements[0])
	} else if stringElement.Value != "test all content" {
		t.Errorf("Expected 'test all content', got '%s'", stringElement.Value)
	}
}

func TestClipboardAllEmpty(t *testing.T) {
	// Clear clipboard first
	clipboardClear([]object.VintObject{}, map[string]object.VintObject{})

	// Test all method with empty clipboard
	result := clipboardAll([]object.VintObject{}, map[string]object.VintObject{})

	if result.Type() == object.ERROR_OBJ {
		t.Errorf("Expected success, got error: %s", result.(*object.Error).Message)
	}

	if arrayResult, ok := result.(*object.Array); !ok {
		t.Errorf("Expected array result, got %T", result)
	} else if len(arrayResult.Elements) != 0 {
		t.Errorf("Expected empty array, got %d elements", len(arrayResult.Elements))
	}
}

func TestClipboardAllInvalidArgs(t *testing.T) {
	// Test with arguments (should accept none)
	result := clipboardAll([]object.VintObject{
		&object.String{Value: "unexpected"},
	}, map[string]object.VintObject{})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("Expected error for arguments, got %T", result)
	}
}
