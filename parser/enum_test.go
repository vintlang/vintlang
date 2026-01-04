package parser

import (
	"testing"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/lexer"
)

func TestEnumStatement(t *testing.T) {
	input := `
	enum Status {
		PENDING = 0,
		ACTIVE = 1,
		COMPLETED = 2
	}
	`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.EnumStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.EnumStatement. got=%T",
			program.Statements[0])
	}

	if stmt.Name.Value != "Status" {
		t.Errorf("enum name is not 'Status'. got=%s", stmt.Name.Value)
	}

	if len(stmt.Values) != 3 {
		t.Errorf("enum does not have 3 members. got=%d", len(stmt.Values))
	}

	expectedMembers := map[string]int64{
		"PENDING":   0,
		"ACTIVE":    1,
		"COMPLETED": 2,
	}

	for name, expectedValue := range expectedMembers {
		expr, ok := stmt.Values[name]
		if !ok {
			t.Errorf("enum member '%s' not found", name)
			continue
		}

		intLiteral, ok := expr.(*ast.IntegerLiteral)
		if !ok {
			t.Errorf("enum member value is not IntegerLiteral. got=%T", expr)
			continue
		}

		if intLiteral.Value != expectedValue {
			t.Errorf("enum member '%s' has wrong value. expected=%d, got=%d",
				name, expectedValue, intLiteral.Value)
		}
	}
}

func TestEnumStatementWithStrings(t *testing.T) {
	input := `
	enum Color {
		RED = "red",
		GREEN = "green",
		BLUE = "blue"
	}
	`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.EnumStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.EnumStatement. got=%T",
			program.Statements[0])
	}

	if stmt.Name.Value != "Color" {
		t.Errorf("enum name is not 'Color'. got=%s", stmt.Name.Value)
	}

	if len(stmt.Values) != 3 {
		t.Errorf("enum does not have 3 members. got=%d", len(stmt.Values))
	}

	expectedMembers := map[string]string{
		"RED":   "red",
		"GREEN": "green",
		"BLUE":  "blue",
	}

	for name, expectedValue := range expectedMembers {
		expr, ok := stmt.Values[name]
		if !ok {
			t.Errorf("enum member '%s' not found", name)
			continue
		}

		strLiteral, ok := expr.(*ast.StringLiteral)
		if !ok {
			t.Errorf("enum member value is not StringLiteral. got=%T", expr)
			continue
		}

		if strLiteral.Value != expectedValue {
			t.Errorf("enum member '%s' has wrong value. expected=%s, got=%s",
				name, expectedValue, strLiteral.Value)
		}
	}
}

