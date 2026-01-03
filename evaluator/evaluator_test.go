package evaluator

import (
	"testing"

	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
)

func TestNullCoalescingOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		// Basic null coalescing
		{"null ?? \"default\"", "default"},
		{"\"value\" ?? \"default\"", "value"},

		// Different data types
		{"null ?? 42", 42},
		{"null ?? true", true},
		{"null ?? false", false},

		// Chained null coalescing
		{"null ?? null ?? \"final\"", "final"},
		{"null ?? \"first\" ?? \"second\"", "first"},

		// Edge cases with non-null "falsy" values
		{"\"\" ?? \"default\"", ""},
		{"0 ?? 999", 0},
		{"false ?? true", false},

		// Complex expressions
		{"let a = null; a ?? \"default\"", "default"},
		{"let a = \"value\"; a ?? \"default\"", "value"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		evaluated := Eval(program, env)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			testStringObject(t, evaluated, expected)
		case bool:
			if evaluated.Type() != object.BOOLEAN_OBJ {
				t.Errorf("input: %q - object is not Boolean. got=%T (%+v)", tt.input, evaluated, evaluated)
				continue
			}
			boolObj := evaluated.(*object.Boolean)
			if boolObj.Value != expected {
				t.Errorf("input: %q - object has wrong value. got=%t, want=%t", tt.input, boolObj.Value, expected)
			}
		}
	}
}

func TestIfExpressionAndStatement(t *testing.T) {
	cases := []struct {
		input    string
		expected object.VintObject
	}{
		// If as an expression
		{"let status = \"\"; status = if (true) { \"Online\" } else { \"Offline\" }; status", &object.String{Value: "Online"}},
		{"let status = \"\"; status = if (false) { \"Online\" } else { \"Offline\" }; status", &object.String{Value: "Offline"}},
		{"let status = \"\"; status = if (false) { \"Online\" }; status", &object.Null{}},
		// If as a statement (side effect)
		{"let x = 0; if (true) { x = 5 }; x", &object.Integer{Value: 5}},
		{"let x = 0; if (false) { x = 5 }; x", &object.Integer{Value: 0}},
	}

	for _, tt := range cases {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)
		if result.Type() != tt.expected.Type() || result.Inspect() != tt.expected.Inspect() {
			t.Errorf("input: %q\nexpected: %s\ngot: %s", tt.input, tt.expected.Inspect(), result.Inspect())
		}
	}
}

func TestAsyncFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{
			"let f = async func() { return 42 }; let p = f(); await p",
			42,
		},
		{
			"let f = async func(x) { return x + 1 }; let p = f(5); await p",
			6,
		},
		{
			"let f = async func() { return \"hello\" }; let p = f(); await p",
			"hello",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, result, int64(expected))
		case string:
			testStringObject(t, result, expected)
		default:
			t.Errorf("unexpected expected type: %T", expected)
		}
	}
}

func TestChannels(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{
			"let ch = chan; ch",
			"chan(unbuffered)",
		},
		{
			"let ch = chan(5); ch",
			"chan(buffered:5)",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		if result.Inspect() != tt.expected {
			t.Errorf("wrong result. expected=%q, got=%q", tt.expected, result.Inspect())
		}
	}
}

func testIntegerObject(t *testing.T, obj object.VintObject, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
		return false
	}
	return true
}

func testStringObject(t *testing.T, obj object.VintObject, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%q, want=%q",
			result.Value, expected)
		return false
	}
	return true
}

func TestRangeExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1..5", "1..5"},
		{"0..3", "0..3"},
		{"10..15", "10..15"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		rangeObj, ok := result.(*object.Range)
		if !ok {
			t.Errorf("object is not Range. got=%T (%+v)", result, result)
			continue
		}

		if rangeObj.Inspect() != tt.expected {
			t.Errorf("range has wrong string representation. got=%q, want=%q",
				rangeObj.Inspect(), tt.expected)
		}
	}
}

func TestRangeInForLoop(t *testing.T) {
	tests := []struct {
		input    string
		expected []int64
	}{
		{
			`let result = []; for i in 1..3 { result = result.push(i) }; result`,
			[]int64{1, 2, 3},
		},
		{
			`let result = []; for i in 0..2 { result = result.push(i) }; result`,
			[]int64{0, 1, 2},
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

		for i, expectedVal := range tt.expected {
			intObj, ok := array.Elements[i].(*object.Integer)
			if !ok {
				t.Errorf("array element %d is not Integer. got=%T (%+v)",
					i, array.Elements[i], array.Elements[i])
				continue
			}
			if intObj.Value != expectedVal {
				t.Errorf("array element %d has wrong value. got=%d, want=%d",
					i, intObj.Value, expectedVal)
			}
		}
	}
}

// TestWhileLoopReturnNull tests that while loops return NULL instead of nil
// when the condition is false from the start (bug fix)
func TestWhileLoopReturnNull(t *testing.T) {
	tests := []struct {
		input    string
		expected object.VintObject
	}{
		// While loop that never executes should return NULL
		{"while (false) { \"value\" }", &object.Null{}},
		// While loop with false condition and variable
		{"let x = while (false) { 42 }; x", &object.Null{}},
		// While loop that executes should return NULL (not last iteration value)
		{"let count = 0; let result = while (count < 3) { count = count + 1; count }; result", &object.Null{}},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		if result.Type() != tt.expected.Type() {
			t.Errorf("input: %q - wrong type. got=%s, want=%s", tt.input, result.Type(), tt.expected.Type())
		}
	}
}

// TestEmptyBlocksReturnNull tests that empty blocks return NULL instead of nil (bug fix)
func TestEmptyBlocksReturnNull(t *testing.T) {
	tests := []struct {
		input    string
		expected object.VintObject
	}{
		// Empty if block
		{"if (true) {}", &object.Null{}},
		// Empty else block
		{"if (false) {} else {}", &object.Null{}},
		// Empty function body
		{"let f = func() {}; f()", &object.Null{}},
		// Empty while block (never executes)
		{"while (false) {}", &object.Null{}},
		// Empty switch case
		{"switch (1) { case 1 {} }", &object.Null{}},
		// Empty match case
		{"match 1 { 1 => {} }", &object.Null{}},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		if result.Type() != tt.expected.Type() {
			t.Errorf("input: %q - wrong type. got=%s, want=%s", tt.input, result.Type(), tt.expected.Type())
		}
	}
}

// TestControlFlowReturnValues tests that control flow structures return correct values
func TestControlFlowReturnValues(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		// If expression with value
		{"if (true) { 42 }", 42},
		{"if (false) { 10 } else { 20 }", 20},
		{"if (1 < 2) { \"yes\" } else { \"no\" }", "yes"},
		
		// If expression last statement
		{"if (true) { let a = 1; let b = 2; a + b }", 3},
		
		// Switch expression with value
		{"switch (1) { case 1 { \"one\" } case 2 { \"two\" } }", "one"},
		{"switch (2) { case 1 { 10 } default { 99 } }", 99},
		
		// Match expression with value
		{"match 1 { 1 => { \"matched\" } _ => { \"default\" } }", "matched"},
		{"match 3 { 1 => { 10 } 2 => { 20 } _ => { 30 } }", 30},
		
		// For-in should return NULL
		{"for i in [1, 2, 3] { i }", nil},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, result, int64(expected))
		case string:
			testStringObject(t, result, expected)
		case nil:
			if result.Type() != object.NULL_OBJ {
				t.Errorf("input: %q - expected NULL, got=%s", tt.input, result.Type())
			}
		}
	}
}

// TestNestedControlFlow tests nested control structures
func TestNestedControlFlow(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		// Nested if
		{"if (true) { if (true) { 42 } }", 42},
		{"if (true) { if (false) { 10 } else { 20 } }", 20},
		
		// If inside while - x starts at 0, loops while x < 3
		// Iteration 1: x=0, x<3 true, x != 2, x becomes 1
		// Iteration 2: x=1, x<3 true, x != 2, x becomes 2
		// Iteration 3: x=2, x<3 true, x == 2, x becomes 10, then 11
		// Iteration 4: x=11, x<3 false, loop exits
		{"let x = 0; while (x < 3) { if (x == 2) { x = 10 }; x = x + 1 }; x", 11},
		
		// Switch inside if
		{"if (true) { switch (1) { case 1 { \"yes\" } } }", "yes"},
		
		// Match inside for-in - last iteration (i=3) matches the _ pattern
		{"let result = \"\"; for i in [1, 2, 3] { result = match i { 2 => { \"two\" } _ => { \"other\" } } }; result", "other"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		result := Eval(program, env)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, result, int64(expected))
		case string:
			testStringObject(t, result, expected)
		}
	}
}
