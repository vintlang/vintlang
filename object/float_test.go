package object

import (
	"math"
	"testing"
)

func TestFloatNewMethods(t *testing.T) {
	tests := []struct {
		method   string
		value    float64
		args     []VintObject
		expected string
		isError  bool
	}{
		// toPrecision tests
		{"toPrecision", 123.456, []VintObject{&Integer{Value: 4}}, "123.5", false},
		{"toPrecision", 0.123456, []VintObject{&Integer{Value: 3}}, "0.123", false},
		
		// toFixed tests
		{"toFixed", 123.456, []VintObject{&Integer{Value: 2}}, "123.46", false},
		{"toFixed", 123.0, []VintObject{&Integer{Value: 2}}, "123.00", false},
		
		// sign tests
		{"sign", 5.5, []VintObject{}, "1", false},
		{"sign", -3.2, []VintObject{}, "-1", false},
		{"sign", 0.0, []VintObject{}, "0", false},
		
		// truncate tests
		{"truncate", 5.9, []VintObject{}, "5", false},
		{"truncate", -3.7, []VintObject{}, "-3", false},
		{"truncate", 0.0, []VintObject{}, "0", false},
		
		// mod tests
		{"mod", 5.5, []VintObject{&Float{Value: 2.0}}, "1.5", false},
		{"mod", 10.0, []VintObject{&Integer{Value: 3}}, "1", false},
		
		// degrees tests (π radians = 180 degrees)
		{"degrees", math.Pi, []VintObject{}, "180", false},
		{"degrees", math.Pi / 2, []VintObject{}, "90", false},
		
		// radians tests (180 degrees = π radians)
		{"radians", 180.0, []VintObject{}, "3.141592653589793", false},
		{"radians", 90.0, []VintObject{}, "1.5707963267948966", false},
		
		// trigonometric tests
		{"sin", 0.0, []VintObject{}, "0", false},
		{"cos", 0.0, []VintObject{}, "1", false},
		{"tan", 0.0, []VintObject{}, "0", false},
		
		// logarithmic tests
		{"exp", 0.0, []VintObject{}, "1", false},
		{"log", math.E, []VintObject{}, "1", false},
	}

	for _, test := range tests {
		float := &Float{Value: test.value}
		result := float.Method(test.method, test.args)
		
		if test.isError {
			if _, ok := result.(*Error); !ok {
				t.Errorf("Expected error for %s(%f), got %s", test.method, test.value, result.Inspect())
			}
		} else {
			// For floating point comparisons, we need to be more lenient
			if test.method == "degrees" || test.method == "radians" || test.method == "sin" || test.method == "cos" || test.method == "tan" || test.method == "log" || test.method == "exp" {
				resultFloat, ok := result.(*Float)
				if !ok {
					t.Errorf("Expected Float for %s(%f), got %T", test.method, test.value, result)
					continue
				}
				expectedFloat := parseFloat(test.expected)
				if math.Abs(resultFloat.Value-expectedFloat) > 1e-10 {
					t.Errorf("Expected %s(%f) ≈ %s, got %f", test.method, test.value, test.expected, resultFloat.Value)
				}
			} else {
				if result.Inspect() != test.expected {
					t.Errorf("Expected %s(%f) = %s, got %s", test.method, test.value, test.expected, result.Inspect())
				}
			}
		}
	}
}

func parseFloat(s string) float64 {
	switch s {
	case "0":
		return 0.0
	case "1":
		return 1.0
	case "-1":
		return -1.0
	case "180":
		return 180.0
	case "90":
		return 90.0
	case "3.141592653589793":
		return math.Pi
	case "1.5707963267948966":
		return math.Pi / 2
	default:
		return 0.0
	}
}

func TestFloatTrigonometry(t *testing.T) {
	tests := []struct {
		method   string
		value    float64
		expected float64
	}{
		{"sin", math.Pi / 2, 1.0},
		{"cos", math.Pi, -1.0},
		{"tan", math.Pi / 4, 1.0},
	}

	for _, test := range tests {
		float := &Float{Value: test.value}
		result := float.Method(test.method, []VintObject{})
		
		resultFloat, ok := result.(*Float)
		if !ok {
			t.Errorf("Expected Float for %s(%f), got %T", test.method, test.value, result)
			continue
		}
		
		if math.Abs(resultFloat.Value-test.expected) > 1e-10 {
			t.Errorf("Expected %s(%f) ≈ %f, got %f", test.method, test.value, test.expected, resultFloat.Value)
		}
	}
}

func TestFloatErrorCases(t *testing.T) {
	float := &Float{Value: 5.5}
	
	// Test toPrecision with invalid precision
	result := float.toPrecision([]VintObject{&Integer{Value: 0}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for toPrecision with precision 0")
	}
	
	result = float.toPrecision([]VintObject{&Integer{Value: 25}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for toPrecision with precision > 21")
	}
	
	// Test toFixed with invalid decimal places
	result = float.toFixed([]VintObject{&Integer{Value: -1}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for toFixed with negative decimal places")
	}
	
	result = float.toFixed([]VintObject{&Integer{Value: 25}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for toFixed with decimal places > 20")
	}
	
	// Test mod with zero divisor
	result = float.mod([]VintObject{&Float{Value: 0.0}})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for mod with zero divisor")
	}
	
	// Test log with non-positive number
	negFloat := &Float{Value: -1.0}
	result = negFloat.log([]VintObject{})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for log of negative number")
	}
	
	zeroFloat := &Float{Value: 0.0}
	result = zeroFloat.log([]VintObject{})
	if _, ok := result.(*Error); !ok {
		t.Errorf("Expected error for log of zero")
	}
}

func TestFloatNaN(t *testing.T) {
	nanFloat := &Float{Value: math.NaN()}
	
	// Test sign with NaN
	result := nanFloat.sign([]VintObject{})
	resultFloat, ok := result.(*Float)
	if !ok {
		t.Errorf("Expected Float for sign(NaN), got %T", result)
		return
	}
	
	if !math.IsNaN(resultFloat.Value) {
		t.Errorf("Expected sign(NaN) to return NaN, got %f", resultFloat.Value)
	}
}