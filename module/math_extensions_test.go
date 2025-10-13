package module

import (
	"testing"

	"github.com/vintlang/vintlang/object"
)

// Test statistics functions

func TestMean(t *testing.T) {
	tests := []struct {
		name     string
		input    []object.VintObject
		expected float64
		hasError bool
	}{
		{
			name: "basic mean",
			input: []object.VintObject{
				&object.Array{Elements: []object.VintObject{
					&object.Integer{Value: 1},
					&object.Integer{Value: 2},
					&object.Integer{Value: 3},
					&object.Integer{Value: 4},
					&object.Integer{Value: 5},
				}},
			},
			expected: 3.0,
			hasError: false,
		},
		{
			name: "mean with floats",
			input: []object.VintObject{
				&object.Array{Elements: []object.VintObject{
					&object.Float{Value: 1.5},
					&object.Float{Value: 2.5},
					&object.Float{Value: 3.5},
				}},
			},
			expected: 2.5,
			hasError: false,
		},
		{
			name: "empty array error",
			input: []object.VintObject{
				&object.Array{Elements: []object.VintObject{}},
			},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mean(tt.input, map[string]object.VintObject{})
			if tt.hasError {
				if result.Type() != object.ERROR_OBJ {
					t.Errorf("expected error, got %s", result.Type())
				}
			} else {
				if result.Type() != object.FLOAT_OBJ {
					t.Errorf("expected float, got %s", result.Type())
				}
				val := result.(*object.Float).Value
				if val != tt.expected {
					t.Errorf("expected %f, got %f", tt.expected, val)
				}
			}
		})
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name     string
		input    []object.VintObject
		expected float64
		hasError bool
	}{
		{
			name: "odd length array",
			input: []object.VintObject{
				&object.Array{Elements: []object.VintObject{
					&object.Integer{Value: 1},
					&object.Integer{Value: 2},
					&object.Integer{Value: 3},
					&object.Integer{Value: 4},
					&object.Integer{Value: 5},
				}},
			},
			expected: 3.0,
			hasError: false,
		},
		{
			name: "even length array",
			input: []object.VintObject{
				&object.Array{Elements: []object.VintObject{
					&object.Integer{Value: 1},
					&object.Integer{Value: 2},
					&object.Integer{Value: 3},
					&object.Integer{Value: 4},
				}},
			},
			expected: 2.5,
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := median(tt.input, map[string]object.VintObject{})
			if tt.hasError {
				if result.Type() != object.ERROR_OBJ {
					t.Errorf("expected error, got %s", result.Type())
				}
			} else {
				if result.Type() != object.FLOAT_OBJ {
					t.Errorf("expected float, got %s", result.Type())
				}
				val := result.(*object.Float).Value
				if val != tt.expected {
					t.Errorf("expected %f, got %f", tt.expected, val)
				}
			}
		})
	}
}

func TestStddev(t *testing.T) {
	input := []object.VintObject{
		&object.Array{Elements: []object.VintObject{
			&object.Integer{Value: 1},
			&object.Integer{Value: 2},
			&object.Integer{Value: 3},
			&object.Integer{Value: 4},
			&object.Integer{Value: 5},
		}},
	}

	result := stddev(input, map[string]object.VintObject{})
	if result.Type() != object.FLOAT_OBJ {
		t.Errorf("expected float, got %s", result.Type())
	}
	// For [1,2,3,4,5], variance is 2, stddev is sqrt(2) ≈ 1.414
	expected := 1.4142135623730951
	val := result.(*object.Float).Value
	if val != expected {
		t.Errorf("expected %f, got %f", expected, val)
	}
}

// Test complex numbers

func TestComplexNum(t *testing.T) {
	input := []object.VintObject{
		&object.Integer{Value: 3},
		&object.Integer{Value: 4},
	}

	result := complexNum(input, map[string]object.VintObject{})
	if result.Type() != object.DICT_OBJ {
		t.Errorf("expected dict, got %s", result.Type())
	}

	dict := result.(*object.Dict)
	realKey := &object.String{Value: "real"}
	imagKey := &object.String{Value: "imag"}

	realPair, hasReal := dict.Pairs[realKey.HashKey()]
	imagPair, hasImag := dict.Pairs[imagKey.HashKey()]

	if !hasReal || !hasImag {
		t.Error("complex number should have real and imag keys")
	}

	realVal := realPair.Value.(*object.Float).Value
	imagVal := imagPair.Value.(*object.Float).Value

	if realVal != 3.0 || imagVal != 4.0 {
		t.Errorf("expected 3+4i, got %f+%fi", realVal, imagVal)
	}
}

func TestAbsComplex(t *testing.T) {
	// Create complex number 3+4i
	dict := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	realKey := &object.String{Value: "real"}
	imagKey := &object.String{Value: "imag"}
	
	dict.Pairs[realKey.HashKey()] = object.DictPair{
		Key:   realKey,
		Value: &object.Float{Value: 3.0},
	}
	dict.Pairs[imagKey.HashKey()] = object.DictPair{
		Key:   imagKey,
		Value: &object.Float{Value: 4.0},
	}

	result := abs([]object.VintObject{dict}, map[string]object.VintObject{})
	if result.Type() != object.FLOAT_OBJ {
		t.Errorf("expected float, got %s", result.Type())
	}

	// |3+4i| = 5
	expected := 5.0
	val := result.(*object.Float).Value
	if val != expected {
		t.Errorf("expected %f, got %f", expected, val)
	}
}

// Test linear algebra functions

func TestDot(t *testing.T) {
	input := []object.VintObject{
		&object.Array{Elements: []object.VintObject{
			&object.Integer{Value: 1},
			&object.Integer{Value: 2},
			&object.Integer{Value: 3},
		}},
		&object.Array{Elements: []object.VintObject{
			&object.Integer{Value: 4},
			&object.Integer{Value: 5},
			&object.Integer{Value: 6},
		}},
	}

	result := dot(input, map[string]object.VintObject{})
	if result.Type() != object.FLOAT_OBJ {
		t.Errorf("expected float, got %s", result.Type())
	}

	// 1*4 + 2*5 + 3*6 = 4 + 10 + 18 = 32
	expected := 32.0
	val := result.(*object.Float).Value
	if val != expected {
		t.Errorf("expected %f, got %f", expected, val)
	}
}

func TestCross(t *testing.T) {
	input := []object.VintObject{
		&object.Array{Elements: []object.VintObject{
			&object.Integer{Value: 1},
			&object.Integer{Value: 2},
			&object.Integer{Value: 3},
		}},
		&object.Array{Elements: []object.VintObject{
			&object.Integer{Value: 4},
			&object.Integer{Value: 5},
			&object.Integer{Value: 6},
		}},
	}

	result := cross(input, map[string]object.VintObject{})
	if result.Type() != object.ARRAY_OBJ {
		t.Errorf("expected array, got %s", result.Type())
	}

	arr := result.(*object.Array)
	if len(arr.Elements) != 3 {
		t.Errorf("expected 3 elements, got %d", len(arr.Elements))
	}

	// Cross product should be [-3, 6, -3]
	expected := []float64{-3, 6, -3}
	for i, elem := range arr.Elements {
		val := elem.(*object.Float).Value
		if val != expected[i] {
			t.Errorf("element %d: expected %f, got %f", i, expected[i], val)
		}
	}
}

func TestMagnitude(t *testing.T) {
	input := []object.VintObject{
		&object.Array{Elements: []object.VintObject{
			&object.Integer{Value: 3},
			&object.Integer{Value: 4},
		}},
	}

	result := magnitude(input, map[string]object.VintObject{})
	if result.Type() != object.FLOAT_OBJ {
		t.Errorf("expected float, got %s", result.Type())
	}

	// sqrt(3^2 + 4^2) = sqrt(9 + 16) = sqrt(25) = 5
	expected := 5.0
	val := result.(*object.Float).Value
	if val != expected {
		t.Errorf("expected %f, got %f", expected, val)
	}
}

// Test numerical methods

func TestGCD(t *testing.T) {
	tests := []struct {
		a        int64
		b        int64
		expected int64
	}{
		{48, 18, 6},
		{12, 15, 3},
		{100, 50, 50},
		{17, 19, 1}, // primes
	}

	for _, tt := range tests {
		input := []object.VintObject{
			&object.Integer{Value: tt.a},
			&object.Integer{Value: tt.b},
		}

		result := gcd(input, map[string]object.VintObject{})
		if result.Type() != object.INTEGER_OBJ {
			t.Errorf("expected integer, got %s", result.Type())
		}

		val := result.(*object.Integer).Value
		if val != tt.expected {
			t.Errorf("gcd(%d, %d): expected %d, got %d", tt.a, tt.b, tt.expected, val)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		a        int64
		b        int64
		expected int64
	}{
		{12, 15, 60},
		{4, 6, 12},
		{10, 5, 10},
	}

	for _, tt := range tests {
		input := []object.VintObject{
			&object.Integer{Value: tt.a},
			&object.Integer{Value: tt.b},
		}

		result := lcm(input, map[string]object.VintObject{})
		if result.Type() != object.INTEGER_OBJ {
			t.Errorf("expected integer, got %s", result.Type())
		}

		val := result.(*object.Integer).Value
		if val != tt.expected {
			t.Errorf("lcm(%d, %d): expected %d, got %d", tt.a, tt.b, tt.expected, val)
		}
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		value    float64
		min      float64
		max      float64
		expected float64
	}{
		{5, 0, 10, 5},
		{-5, 0, 10, 0},
		{15, 0, 10, 10},
		{7.5, 5, 10, 7.5},
	}

	for _, tt := range tests {
		input := []object.VintObject{
			&object.Float{Value: tt.value},
			&object.Float{Value: tt.min},
			&object.Float{Value: tt.max},
		}

		result := clamp(input, map[string]object.VintObject{})
		if result.Type() != object.FLOAT_OBJ {
			t.Errorf("expected float, got %s", result.Type())
		}

		val := result.(*object.Float).Value
		if val != tt.expected {
			t.Errorf("clamp(%f, %f, %f): expected %f, got %f", tt.value, tt.min, tt.max, tt.expected, val)
		}
	}
}

func TestLerp(t *testing.T) {
	tests := []struct {
		start    float64
		end      float64
		t        float64
		expected float64
	}{
		{0, 10, 0.5, 5},
		{0, 10, 0.25, 2.5},
		{0, 10, 0, 0},
		{0, 10, 1, 10},
		{5, 15, 0.5, 10},
	}

	for _, tt := range tests {
		input := []object.VintObject{
			&object.Float{Value: tt.start},
			&object.Float{Value: tt.end},
			&object.Float{Value: tt.t},
		}

		result := lerp(input, map[string]object.VintObject{})
		if result.Type() != object.FLOAT_OBJ {
			t.Errorf("expected float, got %s", result.Type())
		}

		val := result.(*object.Float).Value
		if val != tt.expected {
			t.Errorf("lerp(%f, %f, %f): expected %f, got %f", tt.start, tt.end, tt.t, tt.expected, val)
		}
	}
}

func TestBigint(t *testing.T) {
	tests := []struct {
		name  string
		input object.VintObject
	}{
		{
			name:  "from string",
			input: &object.String{Value: "999999999999999999999"},
		},
		{
			name:  "from integer",
			input: &object.Integer{Value: 12345},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := bigint([]object.VintObject{tt.input}, map[string]object.VintObject{})
			if result.Type() != object.DICT_OBJ {
				t.Errorf("expected dict, got %s", result.Type())
			}

			dict := result.(*object.Dict)
			valueKey := &object.String{Value: "value"}
			typeKey := &object.String{Value: "type"}

			valuePair, hasValue := dict.Pairs[valueKey.HashKey()]
			typePair, hasType := dict.Pairs[typeKey.HashKey()]

			if !hasValue || !hasType {
				t.Error("bigint should have value and type keys")
			}

			typeVal := typePair.Value.(*object.String).Value
			if typeVal != "bigint" {
				t.Errorf("expected type 'bigint', got '%s'", typeVal)
			}

			// Just check that value exists and is a string
			if valuePair.Value.Type() != object.STRING_OBJ {
				t.Errorf("expected string value, got %s", valuePair.Value.Type())
			}
		})
	}
}
