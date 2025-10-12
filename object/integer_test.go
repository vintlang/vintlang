package object

import (
	"testing"
)

func TestIntegerNewMethods(t *testing.T) {
	tests := []struct {
		method   string
		value    int64
		args     []VintObject
		expected string
		isError  bool
	}{
		// toBinary tests
		{"toBinary", 5, []VintObject{}, "101", false},
		{"toBinary", 0, []VintObject{}, "0", false},
		{"toBinary", 255, []VintObject{}, "11111111", false},
		
		// toHex tests
		{"toHex", 255, []VintObject{}, "ff", false},
		{"toHex", 16, []VintObject{}, "10", false},
		{"toHex", 0, []VintObject{}, "0", false},
		
		// toOctal tests
		{"toOctal", 8, []VintObject{}, "10", false},
		{"toOctal", 64, []VintObject{}, "100", false},
		{"toOctal", 0, []VintObject{}, "0", false},
		
		// isPrime tests
		{"isPrime", 2, []VintObject{}, "true", false},
		{"isPrime", 3, []VintObject{}, "true", false},
		{"isPrime", 4, []VintObject{}, "false", false},
		{"isPrime", 17, []VintObject{}, "true", false},
		{"isPrime", 1, []VintObject{}, "false", false},
		{"isPrime", 0, []VintObject{}, "false", false},
		{"isPrime", -5, []VintObject{}, "false", false},
		
		// mod tests
		{"mod", 10, []VintObject{&Integer{Value: 3}}, "1", false},
		{"mod", 15, []VintObject{&Integer{Value: 4}}, "3", false},
		{"mod", 7, []VintObject{&Integer{Value: 7}}, "0", false},
		
		// clamp tests
		{"clamp", 5, []VintObject{&Integer{Value: 1}, &Integer{Value: 10}}, "5", false},
		{"clamp", -5, []VintObject{&Integer{Value: 1}, &Integer{Value: 10}}, "1", false},
		{"clamp", 15, []VintObject{&Integer{Value: 1}, &Integer{Value: 10}}, "10", false},
		
		// inRange tests
		{"inRange", 5, []VintObject{&Integer{Value: 1}, &Integer{Value: 10}}, "true", false},
		{"inRange", 0, []VintObject{&Integer{Value: 1}, &Integer{Value: 10}}, "false", false},
		{"inRange", 11, []VintObject{&Integer{Value: 1}, &Integer{Value: 10}}, "false", false},
	}

	for _, test := range tests {
		integer := &Integer{Value: test.value}
		result := integer.Method(test.method, test.args)
		
		if test.isError {
			if _, ok := result.(*Error); !ok {
				t.Errorf("Expected error for %s(%d), got %s", test.method, test.value, result.Inspect())
			}
		} else {
			if result.Inspect() != test.expected {
				t.Errorf("Expected %s(%d) = %s, got %s", test.method, test.value, test.expected, result.Inspect())
			}
		}
	}
}

func TestIntegerDigits(t *testing.T) {
	tests := []struct {
		value    int64
		expected []int64
	}{
		{123, []int64{1, 2, 3}},
		{0, []int64{0}},
		{987, []int64{9, 8, 7}},
		{-456, []int64{4, 5, 6}}, // negative should return absolute value digits
	}

	for _, test := range tests {
		integer := &Integer{Value: test.value}
		result := integer.digits([]VintObject{})
		
		arr, ok := result.(*Array)
		if !ok {
			t.Errorf("Expected Array for digits(%d), got %T", test.value, result)
			continue
		}
		
		if len(arr.Elements) != len(test.expected) {
			t.Errorf("Expected %d digits for %d, got %d", len(test.expected), test.value, len(arr.Elements))
			continue
		}
		
		for i, elem := range arr.Elements {
			digit, ok := elem.(*Integer)
			if !ok {
				t.Errorf("Expected Integer digit, got %T", elem)
				continue
			}
			
			if digit.Value != test.expected[i] {
				t.Errorf("Expected digit %d at position %d for %d, got %d", test.expected[i], i, test.value, digit.Value)
			}
		}
	}
}

func TestIntegerNthRoot(t *testing.T) {
	tests := []struct {
		value    int64
		root     int64
		expected float64
		isError  bool
	}{
		{8, 3, 2.0, false},
		{16, 2, 4.0, false},
		{27, 3, 3.0, false},
		{-8, 3, -2.0, false},
		{-8, 2, 0, true}, // even root of negative
		{8, 0, 0, true},  // zero root
		{8, -1, 0, true}, // negative root
	}

	for _, test := range tests {
		integer := &Integer{Value: test.value}
		result := integer.nthRoot([]VintObject{&Integer{Value: test.root}})
		
		if test.isError {
			if _, ok := result.(*Error); !ok {
				t.Errorf("Expected error for nthRoot(%d, %d), got %s", test.value, test.root, result.Inspect())
			}
		} else {
			float, ok := result.(*Float)
			if !ok {
				t.Errorf("Expected Float for nthRoot(%d, %d), got %T", test.value, test.root, result)
				continue
			}
			
			// Allow small floating point differences
			if abs(float.Value-test.expected) > 0.0001 {
				t.Errorf("Expected nthRoot(%d, %d) ≈ %f, got %f", test.value, test.root, test.expected, float.Value)
			}
		}
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func TestIntegerErrorCases(t *testing.T) {
	integer := &Integer{Value: 10}
	
	// Test division by zero in mod
	result := integer.mod([]VintObject{&Integer{Value: 0}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for mod with zero divisor")
	}
	
	// Test invalid clamp bounds
	result = integer.clamp([]VintObject{&Integer{Value: 10}, &Integer{Value: 1}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for invalid clamp bounds")
	}
}