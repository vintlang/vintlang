package evaluator

import (
	"testing"

	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
)

func TestCustomErrorTypes(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			// Test error type declaration
			"error TestError(param); type(TestError)",
			"ERROR_TYPE",
		},
		{
			// Test error instance creation
			"error TestError(param); let err = TestError('test'); type(err)",
			"CUSTOM_ERROR",
		},
		{
			// Test error with multiple parameters
			"error MultiError(a, b, c); let err = MultiError(1, 2, 3); type(err)",
			"CUSTOM_ERROR",
		},
		{
			// Test error instance inspection
			"error FileNotFound(path); FileNotFound('/missing.txt')",
			"\x1b[1;31mFileNotFound:\x1b[0m FileNotFound(/missing.txt)",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		
		evaluated := Eval(program, env)
		
		switch expected := tt.expected.(type) {
		case string:
			if evaluated.Type() == object.ERROR_OBJ {
				errObj := evaluated.(*object.Error)
				if errObj.Message != expected && evaluated.Inspect() != expected {
					t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
				}
			} else if evaluated.Inspect() != expected {
				t.Errorf("object has wrong value. expected=%q, got=%q (%T)", expected, evaluated.Inspect(), evaluated)
			}
		}
	}
}

func TestThrowStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"error TestError(msg); throw TestError('something went wrong')",
			"thrown: \x1b[1;31mTestError:\x1b[0m TestError(something went wrong)",
		},
		{
			"throw 'simple error message'",
			"thrown: simple error message",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		
		evaluated := Eval(program, env)
		
		if evaluated.Type() != object.ERROR_OBJ {
			t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
			continue
		}
		
		errObj := evaluated.(*object.Error)
		if errObj.Message != tt.expected {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expected, errObj.Message)
		}
	}
}

func TestErrorParameterValidation(t *testing.T) {
	tests := []struct {
		input        string
		expectedError string
	}{
		{
			"error TestError(param); TestError()",
			"error TestError expects 1 arguments, got 0",
		},
		{
			"error TestError(a, b); TestError('only one')",
			"error TestError expects 2 arguments, got 1",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		
		evaluated := Eval(program, env)
		
		if evaluated.Type() != object.ERROR_OBJ {
			t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
			continue
		}
		
		errObj := evaluated.(*object.Error)
		if errObj.Message != tt.expectedError {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedError, errObj.Message)
		}
	}
}