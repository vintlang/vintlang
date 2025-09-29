package evaluator

import (
	"testing"

	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
)

func TestForInLoopWithArrays(t *testing.T) {
	tests := []struct {
		input    string
		expected []int64
	}{
		{
			"let result = []; for i in [1, 2, 3] { result = result.push(i) }; result",
			[]int64{1, 2, 3},
		},
		{
			"let result = []; for idx, val in [10, 20, 30] { result = result.push(idx) }; result",
			[]int64{0, 1, 2}, // indices
		},
		{
			"let result = []; for idx, val in [10, 20, 30] { result = result.push(val) }; result",
			[]int64{10, 20, 30}, // values
		},
		{
			"let result = []; for i in [] { result = result.push(i) }; result",
			[]int64{}, // empty array
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		array, ok := result.(*object.Array)
		if !ok {
			t.Errorf("object is not Array. got=%T (%+v)", result, result)
			continue
		}

		if len(array.Elements) != len(tt.expected) {
			t.Errorf("array has wrong length. got=%d, want=%d",
				len(array.Elements), len(tt.expected))
			continue
		}

		for i, elem := range array.Elements {
			integer, ok := elem.(*object.Integer)
			if !ok {
				t.Errorf("array element is not Integer. got=%T (%+v)", elem, elem)
				continue
			}

			if integer.Value != tt.expected[i] {
				t.Errorf("array element has wrong value. got=%d, want=%d",
					integer.Value, tt.expected[i])
			}
		}
	}
}

func TestForInLoopWithStrings(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			`let result = []; for char in "abc" { result = result.push(char) }; result`,
			[]string{"a", "b", "c"},
		},
		{
			`let result = []; for idx, char in "hi" { result = result.push(idx) }; result`,
			[]string{"0", "1"}, // indices as strings when converted to inspect
		},
		{
			`let result = []; for char in "" { result = result.push(char) }; result`,
			[]string{}, // empty string
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		array, ok := result.(*object.Array)
		if !ok {
			t.Errorf("object is not Array. got=%T (%+v)", result, result)
			continue
		}

		if len(array.Elements) != len(tt.expected) {
			t.Errorf("array has wrong length. got=%d, want=%d",
				len(array.Elements), len(tt.expected))
			continue
		}

		for i, elem := range array.Elements {
			str, ok := elem.(*object.String)
			if !ok {
				// For indices, they might be integers
				if integer, intOk := elem.(*object.Integer); intOk {
					if integer.Inspect() != tt.expected[i] {
						t.Errorf("array element has wrong value. got=%s, want=%s",
							integer.Inspect(), tt.expected[i])
					}
					continue
				}
				t.Errorf("array element is not String. got=%T (%+v)", elem, elem)
				continue
			}

			if str.Value != tt.expected[i] {
				t.Errorf("array element has wrong value. got=%s, want=%s",
					str.Value, tt.expected[i])
			}
		}
	}
}

func TestForInLoopWithDictionaries(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int64
	}{
		{
			`let dict = {"a": 1, "b": 2}; let result = {}; for key, val in dict { result[key] = val }; result`,
			map[string]int64{"a": 1, "b": 2},
		},
		{
			`let dict = {"x": 10}; let result = []; for key in dict { result = result.push(key) }; result`,
			map[string]int64{"x": 0}, // just checking key exists
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		// Handle both dict and array results
		switch res := result.(type) {
		case *object.Dict:
			if len(res.Pairs) != len(tt.expected) {
				t.Errorf("dict has wrong length. got=%d, want=%d",
					len(res.Pairs), len(tt.expected))
				continue
			}

			for _, pair := range res.Pairs {
				key := pair.Key.Inspect()
				expectedVal, exists := tt.expected[key]
				if !exists {
					t.Errorf("unexpected key in result dict: %s", key)
					continue
				}

				integer, ok := pair.Value.(*object.Integer)
				if !ok {
					t.Errorf("dict value is not Integer. got=%T (%+v)", pair.Value, pair.Value)
					continue
				}

				if integer.Value != expectedVal {
					t.Errorf("dict value has wrong value. got=%d, want=%d",
						integer.Value, expectedVal)
				}
			}
		case *object.Array:
			// For key-only tests, just verify we got the keys
			for key := range tt.expected {
				found := false
				for _, elem := range res.Elements {
					if str, ok := elem.(*object.String); ok && str.Value == key {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected key %s not found in result array", key)
				}
			}
		default:
			t.Errorf("unexpected result type. got=%T (%+v)", result, result)
		}
	}
}

func TestForInLoopWithRanges(t *testing.T) {
	tests := []struct {
		input    string
		expected []int64
	}{
		{
			`let result = []; for i in 1..3 { result = result.push(i) }; result`,
			[]int64{1, 2, 3},
		},
		{
			`let result = []; for i in 0..0 { result = result.push(i) }; result`,
			[]int64{0},
		},
		{
			`let result = []; for i in 5..7 { result = result.push(i) }; result`,
			[]int64{5, 6, 7},
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		array, ok := result.(*object.Array)
		if !ok {
			t.Errorf("object is not Array. got=%T (%+v)", result, result)
			continue
		}

		if len(array.Elements) != len(tt.expected) {
			t.Errorf("array has wrong length. got=%d, want=%d",
				len(array.Elements), len(tt.expected))
			continue
		}

		for i, elem := range array.Elements {
			integer, ok := elem.(*object.Integer)
			if !ok {
				t.Errorf("array element is not Integer. got=%T (%+v)", elem, elem)
				continue
			}

			if integer.Value != tt.expected[i] {
				t.Errorf("array element has wrong value. got=%d, want=%d",
					integer.Value, tt.expected[i])
			}
		}
	}
}

func TestForInLoopWithBreakAndContinue(t *testing.T) {
	tests := []struct {
		input    string
		expected []int64
	}{
		{
			`let result = []; for i in [1, 2, 3, 4, 5] { if (i == 3) { break }; result = result.push(i) }; result`,
			[]int64{1, 2},
		},
		{
			`let result = []; for i in [1, 2, 3, 4, 5] { if (i == 3) { continue }; result = result.push(i) }; result`,
			[]int64{1, 2, 4, 5},
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		array, ok := result.(*object.Array)
		if !ok {
			t.Errorf("object is not Array. got=%T (%+v)", result, result)
			continue
		}

		if len(array.Elements) != len(tt.expected) {
			t.Errorf("array has wrong length. got=%d, want=%d",
				len(array.Elements), len(tt.expected))
			continue
		}

		for i, elem := range array.Elements {
			integer, ok := elem.(*object.Integer)
			if !ok {
				t.Errorf("array element is not Integer. got=%T (%+v)", elem, elem)
				continue
			}

			if integer.Value != tt.expected[i] {
				t.Errorf("array element has wrong value. got=%d, want=%d",
					integer.Value, tt.expected[i])
			}
		}
	}
}

func TestForInLoopErrorCases(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{
			"for i in 42 { print(i) }",
			true, // integer is not iterable
		},
		{
			"for i in true { print(i) }",
			true, // boolean is not iterable
		},
		{
			"for i in null { print(i) }",
			true, // null is not iterable
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		if tt.expectError {
			if _, ok := result.(*object.Error); !ok {
				t.Errorf("expected error for input %q, but got %T (%+v)", tt.input, result, result)
			}
		} else {
			if _, ok := result.(*object.Error); ok {
				t.Errorf("unexpected error for input %q: %s", tt.input, result.(*object.Error).Message)
			}
		}
	}
}

func TestNestedForInLoops(t *testing.T) {
	input := `
		let result = [];
		for i in [1, 2] {
			for j in [10, 20] {
				result = result.push(i * j);
			}
		};
		result
	`

	expected := []int64{10, 20, 20, 40} // 1*10, 1*20, 2*10, 2*20

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	result := Eval(program, env)

	array, ok := result.(*object.Array)
	if !ok {
		t.Errorf("object is not Array. got=%T (%+v)", result, result)
		return
	}

	if len(array.Elements) != len(expected) {
		t.Errorf("array has wrong length. got=%d, want=%d",
			len(array.Elements), len(expected))
		return
	}

	for i, elem := range array.Elements {
		integer, ok := elem.(*object.Integer)
		if !ok {
			t.Errorf("array element is not Integer. got=%T (%+v)", elem, elem)
			continue
		}

		if integer.Value != expected[i] {
			t.Errorf("array element has wrong value. got=%d, want=%d",
				integer.Value, expected[i])
		}
	}
}

func TestForInLoopWithReturn(t *testing.T) {
	input := `
		let f = fun() {
			for i in [1, 2, 3] {
				if (i == 2) {
					return i;
				}
			}
			return -1;
		};
		f()
	`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	result := Eval(program, env)

	integer, ok := result.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", result, result)
		return
	}

	if integer.Value != 2 {
		t.Errorf("wrong value returned. got=%d, want=%d", integer.Value, 2)
	}
}