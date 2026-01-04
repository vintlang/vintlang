package evaluator

import (
	"testing"

	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
)

func TestEnumEvaluation(t *testing.T) {
	input := `
	enum Status {
		PENDING = 0,
		ACTIVE = 1,
		COMPLETED = 2
	}
	
	let current = Status.ACTIVE
	current
	`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	result := Eval(program, env)

	integer, ok := result.(*object.Integer)
	if !ok {
		t.Fatalf("result is not Integer. got=%T (%+v)", result, result)
	}

	if integer.Value != 1 {
		t.Errorf("integer has wrong value. expected=1, got=%d", integer.Value)
	}
}

func TestEnumWithStrings(t *testing.T) {
	input := `
	enum Environment {
		DEV = "development",
		STAGING = "staging",
		PROD = "production"
	}
	
	let env = Environment.PROD
	env
	`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	result := Eval(program, env)

	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("result is not String. got=%T (%+v)", result, result)
	}

	if str.Value != "production" {
		t.Errorf("string has wrong value. expected='production', got=%s", str.Value)
	}
}

func TestEnumInConditional(t *testing.T) {
	input := `
	enum OrderStatus {
		PENDING = 0,
		CONFIRMED = 1,
		SHIPPED = 2,
		DELIVERED = 3
	}
	
	let myOrder = OrderStatus.CONFIRMED
	let result = ""
	
	if myOrder == OrderStatus.CONFIRMED {
		result = "confirmed"
	} else {
		result = "not confirmed"
	}
	
	result
	`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	result := Eval(program, env)

	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("result is not String. got=%T (%+v)", result, result)
	}

	if str.Value != "confirmed" {
		t.Errorf("string has wrong value. expected='confirmed', got=%s", str.Value)
	}
}

func TestEnumMemberNotFound(t *testing.T) {
	input := `
	enum Status {
		PENDING = 0,
		ACTIVE = 1
	}
	
	Status.INVALID
	`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	result := Eval(program, env)

	errObj, ok := result.(*object.Error)
	if !ok {
		t.Fatalf("result is not Error. got=%T (%+v)", result, result)
	}

	expectedMsg := "Enum 'Status' has no member 'INVALID'"
	if errObj.Message != expectedMsg {
		t.Errorf("wrong error message. expected=%q, got=%q", expectedMsg, errObj.Message)
	}
}

func TestEnumImmutability(t *testing.T) {
	input := `
	enum Status {
		PENDING = 0,
		ACTIVE = 1
	}
	
	Status = 5
	`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	result := Eval(program, env)

	errObj, ok := result.(*object.Error)
	if !ok {
		t.Fatalf("result is not Error. got=%T (%+v)", result, result)
	}

	// The error message should indicate that it cannot assign to a constant
	// The exact message might vary, so we just check it's an error
	if errObj.Type() != object.ERROR_OBJ {
		t.Errorf("expected error for attempting to reassign enum")
	}
}

