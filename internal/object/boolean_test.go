package object

import (
	"testing"
)

func TestBooleanToggle(t *testing.T) {
	tests := []struct {
		value    bool
		expected bool
	}{
		{true, false},
		{false, true},
	}

	for _, test := range tests {
		boolean := &Boolean{Value: test.value}
		result := boolean.toggle([]VintObject{})

		resultBool, ok := result.(*Boolean)
		if !ok {
			t.Errorf("Expected Boolean for toggle(%t), got %T", test.value, result)
			continue
		}

		if resultBool.Value != test.expected {
			t.Errorf("Expected toggle(%t) = %t, got %t", test.value, test.expected, resultBool.Value)
		}
	}
}

func TestBooleanToggleWithArgs(t *testing.T) {
	boolean := &Boolean{Value: true}

	// Test toggle with arguments (should error)
	result := boolean.toggle([]VintObject{&String{Value: "test"}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for toggle with arguments")
	}
}

func TestBooleanExistingMethods(t *testing.T) {
	trueBool := &Boolean{Value: true}
	falseBool := &Boolean{Value: false}

	// Test toString
	result := trueBool.toString([]VintObject{})
	if result.Inspect() != "true" {
		t.Errorf("Expected true.toString() = 'true', got %s", result.Inspect())
	}

	result = falseBool.toString([]VintObject{})
	if result.Inspect() != "false" {
		t.Errorf("Expected false.toString() = 'false', got %s", result.Inspect())
	}

	// Test toInt
	result = trueBool.toInt([]VintObject{})
	if result.Inspect() != "1" {
		t.Errorf("Expected true.toInt() = '1', got %s", result.Inspect())
	}

	result = falseBool.toInt([]VintObject{})
	if result.Inspect() != "0" {
		t.Errorf("Expected false.toInt() = '0', got %s", result.Inspect())
	}

	// Test negate (should be same as toggle)
	result = trueBool.negate([]VintObject{})
	if result.Inspect() != "false" {
		t.Errorf("Expected true.negate() = 'false', got %s", result.Inspect())
	}

	result = falseBool.negate([]VintObject{})
	if result.Inspect() != "true" {
		t.Errorf("Expected false.negate() = 'true', got %s", result.Inspect())
	}
}

func TestBooleanLogicalOperations(t *testing.T) {
	trueBool := &Boolean{Value: true}
	falseBool := &Boolean{Value: false}

	// Test AND
	result := trueBool.and([]VintObject{trueBool})
	if result.Inspect() != "true" {
		t.Errorf("Expected true AND true = true, got %s", result.Inspect())
	}

	result = trueBool.and([]VintObject{falseBool})
	if result.Inspect() != "false" {
		t.Errorf("Expected true AND false = false, got %s", result.Inspect())
	}

	// Test OR
	result = falseBool.or([]VintObject{falseBool})
	if result.Inspect() != "false" {
		t.Errorf("Expected false OR false = false, got %s", result.Inspect())
	}

	result = falseBool.or([]VintObject{trueBool})
	if result.Inspect() != "true" {
		t.Errorf("Expected false OR true = true, got %s", result.Inspect())
	}

	// Test XOR
	result = trueBool.xor([]VintObject{trueBool})
	if result.Inspect() != "false" {
		t.Errorf("Expected true XOR true = false, got %s", result.Inspect())
	}

	result = trueBool.xor([]VintObject{falseBool})
	if result.Inspect() != "true" {
		t.Errorf("Expected true XOR false = true, got %s", result.Inspect())
	}

	// Test IMPLIES
	result = trueBool.implies([]VintObject{falseBool})
	if result.Inspect() != "false" {
		t.Errorf("Expected true IMPLIES false = false, got %s", result.Inspect())
	}

	result = falseBool.implies([]VintObject{trueBool})
	if result.Inspect() != "true" {
		t.Errorf("Expected false IMPLIES true = true, got %s", result.Inspect())
	}

	// Test EQUIVALENT
	result = trueBool.equivalent([]VintObject{trueBool})
	if result.Inspect() != "true" {
		t.Errorf("Expected true EQUIVALENT true = true, got %s", result.Inspect())
	}

	result = trueBool.equivalent([]VintObject{falseBool})
	if result.Inspect() != "false" {
		t.Errorf("Expected true EQUIVALENT false = false, got %s", result.Inspect())
	}

	// Test NOR
	result = falseBool.nor([]VintObject{falseBool})
	if result.Inspect() != "true" {
		t.Errorf("Expected false NOR false = true, got %s", result.Inspect())
	}

	result = trueBool.nor([]VintObject{falseBool})
	if result.Inspect() != "false" {
		t.Errorf("Expected true NOR false = false, got %s", result.Inspect())
	}

	// Test NAND
	result = trueBool.nand([]VintObject{trueBool})
	if result.Inspect() != "false" {
		t.Errorf("Expected true NAND true = false, got %s", result.Inspect())
	}

	result = trueBool.nand([]VintObject{falseBool})
	if result.Inspect() != "true" {
		t.Errorf("Expected true NAND false = true, got %s", result.Inspect())
	}
}

func TestBooleanErrorCases(t *testing.T) {
	boolean := &Boolean{Value: true}

	// Test methods with wrong argument types
	result := boolean.and([]VintObject{&String{Value: "test"}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for and() with non-Boolean argument")
	}

	result = boolean.or([]VintObject{&Integer{Value: 1}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for or() with non-Boolean argument")
	}

	// Test methods with wrong number of arguments
	result = boolean.and([]VintObject{})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for and() with no arguments")
	}

	result = boolean.or([]VintObject{&Boolean{Value: true}, &Boolean{Value: false}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for or() with too many arguments")
	}

	// Test unsupported method
	result = boolean.Method("unsupported", []VintObject{})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for unsupported method")
	}
}
