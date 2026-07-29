package parser

import (
	"testing"

	"github.com/vintlang/vintlang/internal/ast"
	"github.com/vintlang/vintlang/internal/lexer"
)

func TestTypedLetStatement(t *testing.T) {
	input := `let x: int = 5`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.TypedLetStatement)
	if !ok {
		t.Fatalf("stmt not *ast.TypedLetStatement. got=%T", program.Statements[0])
	}

	if stmt.Name.Value != "x" {
		t.Errorf("stmt.Name.Value not 'x'. got=%s", stmt.Name.Value)
	}

	if stmt.TypeAnnotation == nil {
		t.Fatal("stmt.TypeAnnotation is nil")
	}

	basicType, ok := stmt.TypeAnnotation.Type.(*ast.BasicType)
	if !ok {
		t.Fatalf("type not *ast.BasicType. got=%T", stmt.TypeAnnotation.Type)
	}
	if basicType.Name != "int" {
		t.Errorf("basicType.Name not 'int'. got=%s", basicType.Name)
	}

	if stmt.Value == nil {
		t.Fatal("stmt.Value is nil")
	}
	if !testIntegerLiteral(t, stmt.Value, 5) {
		return
	}
}

func TestTypedLetZeroValue(t *testing.T) {
	input := `let x: int`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.TypedLetStatement)
	if !ok {
		t.Fatalf("stmt not *ast.TypedLetStatement. got=%T", program.Statements[0])
	}

	if stmt.Name.Value != "x" {
		t.Errorf("stmt.Name.Value not 'x'. got=%s", stmt.Name.Value)
	}

	if stmt.TypeAnnotation == nil {
		t.Fatal("stmt.TypeAnnotation is nil")
	}

	basicType, ok := stmt.TypeAnnotation.Type.(*ast.BasicType)
	if !ok {
		t.Fatalf("type not *ast.BasicType. got=%T", stmt.TypeAnnotation.Type)
	}
	if basicType.Name != "int" {
		t.Errorf("basicType.Name not 'int'. got=%s", basicType.Name)
	}

	if stmt.Value != nil {
		t.Errorf("stmt.Value should be nil for zero value. got=%s", stmt.Value)
	}
}

func TestTypedLetMultipleTypes(t *testing.T) {
	tests := []struct {
		input     string
		name      string
		typeName  string
		hasValue  bool
		value     int64
	}{
		{"let a: string = \"hello\"", "a", "string", true, 0},
		{"let b: bool = true", "b", "bool", true, 0},
		{"let c: float64 = 3.14", "c", "float64", true, 0},
		{"let d: int8", "d", "int8", false, 0},
		{"let e: []int = [1, 2, 3]", "e", "[]int", true, 0},
		{"let f: {string: int}", "f", "{string: int}", false, 0},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			p.Errors()
		}
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.TypedLetStatement)
		if stmt.Name.Value != tt.name {
			t.Errorf("name wrong. expected=%q, got=%q", tt.name, stmt.Name.Value)
		}
	}
}

func TestUntypedLetStillWorks(t *testing.T) {
	input := `let x = 5`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	_, ok := program.Statements[0].(*ast.LetStatement)
	if !ok {
		t.Fatalf("stmt not *ast.LetStatement. got=%T", program.Statements[0])
	}
}

func TestMissingTypeError(t *testing.T) {
	input := `let x`
	l := lexer.New(input)
	p := New(l)
	p.ParseProgram()

	errors := p.Errors()
	if len(errors) == 0 {
		t.Errorf("expected error for 'let x' with no type or initializer, got none")
	}
}

func TestTypedConstStatement(t *testing.T) {
	input := `const APP_NAME: string = "VintLang"`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.TypedLetStatement)
	if !ok {
		t.Fatalf("stmt not *ast.TypedLetStatement. got=%T", program.Statements[0])
	}

	if stmt.Name.Value != "APP_NAME" {
		t.Errorf("name wrong. expected='APP_NAME', got=%q", stmt.Name.Value)
	}

	basicType, ok := stmt.TypeAnnotation.Type.(*ast.BasicType)
	if !ok {
		t.Fatalf("type not *ast.BasicType. got=%T", stmt.TypeAnnotation.Type)
	}
	if basicType.Name != "string" {
		t.Errorf("type name wrong. expected='string', got=%s", basicType.Name)
	}
}

func TestTypedFunctionLiteral(t *testing.T) {
	input := `let add = func(a: int, b: int): int { return a + b }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	// 'let add = func...' has no type annotation on 'add' itself
	// so it's a LetStatement with a TypedFunctionLiteral value
	letStmt, ok := program.Statements[0].(*ast.LetStatement)
	if !ok {
		t.Fatalf("stmt not *ast.LetStatement. got=%T", program.Statements[0])
	}

	funcLit, ok := letStmt.Value.(*ast.TypedFunctionLiteral)
	if !ok {
		t.Fatalf("value not *ast.TypedFunctionLiteral. got=%T", letStmt.Value)
	}

	if len(funcLit.Parameters) != 2 {
		t.Fatalf("expected 2 parameters. got=%d", len(funcLit.Parameters))
	}

	if funcLit.Parameters[0].Identifier.Value != "a" {
		t.Errorf("param[0] name wrong. expected='a', got=%q", funcLit.Parameters[0].Identifier.Value)
	}
	if funcLit.Parameters[0].Type.String() != "int" {
		t.Errorf("param[0] type wrong. expected='int', got=%s", funcLit.Parameters[0].Type.String())
	}

	if funcLit.Parameters[1].Identifier.Value != "b" {
		t.Errorf("param[1] name wrong. expected='b', got=%q", funcLit.Parameters[1].Identifier.Value)
	}
	if funcLit.Parameters[1].Type.String() != "int" {
		t.Errorf("param[1] type wrong. expected='int', got=%s", funcLit.Parameters[1].Type.String())
	}

	if funcLit.ReturnType == nil {
		t.Fatal("ReturnType is nil")
	}
	if funcLit.ReturnType.String() != "int" {
		t.Errorf("ReturnType wrong. expected='int', got=%s", funcLit.ReturnType.String())
	}
}

func TestUntypedFunctionLiteral(t *testing.T) {
	input := `let add = func(a, b) { return a + b }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	letStmt, ok := program.Statements[0].(*ast.LetStatement)
	if !ok {
		t.Fatalf("stmt not *ast.LetStatement. got=%T", program.Statements[0])
	}

	_, ok = letStmt.Value.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("value not *ast.FunctionLiteral. got=%T", letStmt.Value)
	}
}

func TestTypedFunctionNoReturn(t *testing.T) {
	input := `let logger = func(msg: string) { println(msg) }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	letStmt := program.Statements[0].(*ast.LetStatement)
	funcLit := letStmt.Value.(*ast.TypedFunctionLiteral)

	if len(funcLit.Parameters) != 1 {
		t.Fatalf("expected 1 parameter. got=%d", len(funcLit.Parameters))
	}
	if funcLit.Parameters[0].Type.String() != "string" {
		t.Errorf("param type wrong. expected='string', got=%s", funcLit.Parameters[0].Type.String())
	}
	if funcLit.ReturnType != nil {
		t.Errorf("ReturnType should be nil for void function. got=%s", funcLit.ReturnType.String())
	}
}

func TestTypeAliasStatement(t *testing.T) {
	input := `type UserID = int`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.TypeAliasStatement)
	if !ok {
		t.Fatalf("stmt not *ast.TypeAliasStatement. got=%T", program.Statements[0])
	}

	if stmt.Name.Value != "UserID" {
		t.Errorf("alias name wrong. expected='UserID', got=%q", stmt.Name.Value)
	}

	basicType, ok := stmt.Target.(*ast.BasicType)
	if !ok {
		t.Fatalf("target not *ast.BasicType. got=%T", stmt.Target)
	}
	if basicType.Name != "int" {
		t.Errorf("target type wrong. expected='int', got=%s", basicType.Name)
	}
}

func TestTypeAliasComplex(t *testing.T) {
	tests := []struct {
		input        string
		aliasName    string
		targetString string
	}{
		{"type Handler = func(int) bool", "Handler", "func(int) bool"},
		{"type JSON = {string: any}", "JSON", "{string: any}"},
		{"type IntSlice = []int", "IntSlice", "[]int"},
		{"type Matrix = [][]float64", "Matrix", "[][]float64"},
		{"type PointPtr = *Point", "PointPtr", "*Point"},
		{"type IntChan = chan int", "IntChan", "chan int"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.TypeAliasStatement)
		if stmt.Name.Value != tt.aliasName {
			t.Errorf("alias name wrong. expected=%q, got=%q", tt.aliasName, stmt.Name.Value)
		}
		if stmt.Target.String() != tt.targetString {
			t.Errorf("target string wrong. expected=%q, got=%q", tt.targetString, stmt.Target.String())
		}
	}
}

func TestArrayTypeAnnotation(t *testing.T) {
	input := `let scores: []int = [1, 2, 3]`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.TypedLetStatement)
	arrType, ok := stmt.TypeAnnotation.Type.(*ast.ArrayType)
	if !ok {
		t.Fatalf("type not *ast.ArrayType. got=%T", stmt.TypeAnnotation.Type)
	}
	elemType, ok := arrType.ElementType.(*ast.BasicType)
	if !ok {
		t.Fatalf("element type not *ast.BasicType. got=%T", arrType.ElementType)
	}
	if elemType.Name != "int" {
		t.Errorf("element type name wrong. expected='int', got=%s", elemType.Name)
	}
}

func TestFixedArrayTypeAnnotation(t *testing.T) {
	input := `let rgb: [3]int = [255, 0, 128]`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.TypedLetStatement)
	fixedArrType, ok := stmt.TypeAnnotation.Type.(*ast.FixedArrayType)
	if !ok {
		t.Fatalf("type not *ast.FixedArrayType. got=%T", stmt.TypeAnnotation.Type)
	}
	if fixedArrType.Size != 3 {
		t.Errorf("array size wrong. expected=3, got=%d", fixedArrType.Size)
	}
	elemType, ok := fixedArrType.ElementType.(*ast.BasicType)
	if !ok {
		t.Fatalf("element type not *ast.BasicType. got=%T", fixedArrType.ElementType)
	}
	if elemType.Name != "int" {
		t.Errorf("element type name wrong. expected='int', got=%s", elemType.Name)
	}
}

func TestDictTypeAnnotation(t *testing.T) {
	input := `let config: {string: any} = {"host": "localhost"}`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.TypedLetStatement)
	dictType, ok := stmt.TypeAnnotation.Type.(*ast.DictType)
	if !ok {
		t.Fatalf("type not *ast.DictType. got=%T", stmt.TypeAnnotation.Type)
	}

	keyType, ok := dictType.KeyType.(*ast.BasicType)
	if !ok {
		t.Fatalf("key type not *ast.BasicType. got=%T", dictType.KeyType)
	}
	if keyType.Name != "string" {
		t.Errorf("key type wrong. expected='string', got=%s", keyType.Name)
	}

	valType, ok := dictType.ValueType.(*ast.BasicType)
	if !ok {
		t.Fatalf("value type not *ast.BasicType. got=%T", dictType.ValueType)
	}
	if valType.Name != "any" {
		t.Errorf("value type wrong. expected='any', got=%s", valType.Name)
	}
}

func TestPointerTypeAnnotation(t *testing.T) {
	input := `let ptr: *int = &x`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.TypedLetStatement)
	ptrType, ok := stmt.TypeAnnotation.Type.(*ast.PointerType)
	if !ok {
		t.Fatalf("type not *ast.PointerType. got=%T", stmt.TypeAnnotation.Type)
	}
	baseType, ok := ptrType.BaseType.(*ast.BasicType)
	if !ok {
		t.Fatalf("base type not *ast.BasicType. got=%T", ptrType.BaseType)
	}
	if baseType.Name != "int" {
		t.Errorf("base type name wrong. expected='int', got=%s", baseType.Name)
	}
}

func TestChannelTypeAnnotation(t *testing.T) {
	input := `let ch: chan string = chan(string)`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.TypedLetStatement)
	chanType, ok := stmt.TypeAnnotation.Type.(*ast.ChannelType)
	if !ok {
		t.Fatalf("type not *ast.ChannelType. got=%T", stmt.TypeAnnotation.Type)
	}
	elemType, ok := chanType.ElementType.(*ast.BasicType)
	if !ok {
		t.Fatalf("element type not *ast.BasicType. got=%T", chanType.ElementType)
	}
	if elemType.Name != "string" {
		t.Errorf("element type name wrong. expected='string', got=%s", elemType.Name)
	}
}

func TestFunctionTypeAnnotation(t *testing.T) {
	input := `let op: func(int, int) int = add`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.TypedLetStatement)
	funcType, ok := stmt.TypeAnnotation.Type.(*ast.FunctionType)
	if !ok {
		t.Fatalf("type not *ast.FunctionType. got=%T", stmt.TypeAnnotation.Type)
	}
	if len(funcType.Parameters) != 2 {
		t.Fatalf("expected 2 parameter types. got=%d", len(funcType.Parameters))
	}
	if funcType.Parameters[0].String() != "int" {
		t.Errorf("param[0] type wrong. expected='int', got=%s", funcType.Parameters[0].String())
	}
	if funcType.Parameters[1].String() != "int" {
		t.Errorf("param[1] type wrong. expected='int', got=%s", funcType.Parameters[1].String())
	}
	if funcType.ReturnType.String() != "int" {
		t.Errorf("return type wrong. expected='int', got=%s", funcType.ReturnType.String())
	}
}

func TestMultiReturnTypeAnnotation(t *testing.T) {
	input := `let divide = func(a: int, b: int): (int, error) { return 0 }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	letStmt := program.Statements[0].(*ast.LetStatement)
	funcLit := letStmt.Value.(*ast.TypedFunctionLiteral)

	if funcLit.ReturnType == nil {
		t.Fatal("ReturnType is nil")
	}

	multiRet, ok := funcLit.ReturnType.(*ast.MultiReturnType)
	if !ok {
		t.Fatalf("return type not *ast.MultiReturnType. got=%T", funcLit.ReturnType)
	}
	if len(multiRet.Types) != 2 {
		t.Fatalf("expected 2 return types. got=%d", len(multiRet.Types))
	}
	if multiRet.Types[0].String() != "int" {
		t.Errorf("return type[0] wrong. expected='int', got=%s", multiRet.Types[0].String())
	}
	if multiRet.Types[1].String() != "error" {
		t.Errorf("return type[1] wrong. expected='error', got=%s", multiRet.Types[1].String())
	}
}

func TestAsExpression(t *testing.T) {
	input := `let y: float64 = x as float64`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.TypedLetStatement)
	asExpr, ok := stmt.Value.(*ast.TypeCastExpression)
	if !ok {
		t.Fatalf("value not *ast.TypeCastExpression. got=%T", stmt.Value)
	}

	ident, ok := asExpr.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression not *ast.Identifier. got=%T", asExpr.Expression)
	}
	if ident.Value != "x" {
		t.Errorf("identifier wrong. expected='x', got=%s", ident.Value)
	}

	targetType, ok := asExpr.TargetType.(*ast.BasicType)
	if !ok {
		t.Fatalf("target type not *ast.BasicType. got=%T", asExpr.TargetType)
	}
	if targetType.Name != "float64" {
		t.Errorf("target type wrong. expected='float64', got=%s", targetType.Name)
	}
}

func TestIsExpression(t *testing.T) {
	input := `let result: bool = x is int`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.TypedLetStatement)
	isExpr, ok := stmt.Value.(*ast.TypeCheckExpression)
	if !ok {
		t.Fatalf("value not *ast.TypeCheckExpression. got=%T", stmt.Value)
	}

	checkType, ok := isExpr.CheckType.(*ast.BasicType)
	if !ok {
		t.Fatalf("check type not *ast.BasicType. got=%T", isExpr.CheckType)
	}
	if checkType.Name != "int" {
		t.Errorf("check type wrong. expected='int', got=%s", checkType.Name)
	}
}

func TestTypedStructField(t *testing.T) {
	input := `struct Point { x: int, y: int }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.StructStatement)
	if len(stmt.Fields) != 2 {
		t.Fatalf("expected 2 fields. got=%d", len(stmt.Fields))
	}

	for i, expectedName := range []string{"x", "y"} {
		if stmt.Fields[i].Name.Value != expectedName {
			t.Errorf("field[%d] name wrong. expected=%q, got=%q", i, expectedName, stmt.Fields[i].Name.Value)
		}
		if stmt.Fields[i].Type == nil {
			t.Fatalf("field[%d] type is nil", i)
		}
		basicType, ok := stmt.Fields[i].Type.(*ast.BasicType)
		if !ok {
			t.Fatalf("field[%d] type not *ast.BasicType. got=%T", i, stmt.Fields[i].Type)
		}
		if basicType.Name != "int" {
			t.Errorf("field[%d] type name wrong. expected='int', got=%s", i, basicType.Name)
		}
	}
}

func TestTypedStructFieldWithDefault(t *testing.T) {
	input := `struct Config { host: string = "localhost", port: int = 8080 }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.StructStatement)
	if len(stmt.Fields) != 2 {
		t.Fatalf("expected 2 fields. got=%d", len(stmt.Fields))
	}

	if stmt.Fields[0].Type == nil {
		t.Fatal("field[0] type is nil")
	}
	if stmt.Fields[0].Default == nil {
		t.Fatal("field[0] default is nil")
	}
	if stmt.Fields[0].Name.Value != "host" {
		t.Errorf("field[0] name wrong. expected='host', got=%q", stmt.Fields[0].Name.Value)
	}
}

func TestUntypedStructFieldBackwardCompat(t *testing.T) {
	input := `struct Point { x: 0, y: 0 }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.StructStatement)
	if len(stmt.Fields) != 2 {
		t.Fatalf("expected 2 fields. got=%d", len(stmt.Fields))
	}

	// Old-style struct literals should have Type == nil
	if stmt.Fields[0].Type != nil {
		t.Errorf("field[0] should have nil Type for old-style struct. got=%T", stmt.Fields[0].Type)
	}
	if stmt.Fields[0].Default == nil {
		t.Fatal("field[0] default is nil")
	}
}

func TestStructWithUntypedMethods(t *testing.T) {
	input := `struct Point { x: int, y: int, func area() { return 0 } }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.StructStatement)
	if len(stmt.Fields) != 2 {
		t.Errorf("expected 2 fields. got=%d", len(stmt.Fields))
	}
	if len(stmt.Methods) != 1 {
		t.Errorf("expected 1 method. got=%d", len(stmt.Methods))
	}
}

func TestStructMethodTypedParams(t *testing.T) {
	input := `struct Calc { func add(x: int, y: int): int { return x + y } }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.StructStatement)
	if len(stmt.Methods) != 1 {
		t.Fatalf("expected 1 method, got %d", len(stmt.Methods))
	}

	m := stmt.Methods[0]
	if m.Name.Value != "add" {
		t.Errorf("method name wrong, expected 'add', got %s", m.Name.Value)
	}
	if len(m.Parameters) != 2 {
		t.Fatalf("expected 2 params, got %d", len(m.Parameters))
	}
	if len(m.ParamTypes) != 2 {
		t.Fatalf("expected 2 param types, got %d", len(m.ParamTypes))
	}
	if m.ParamTypes[0].String() != "int" {
		t.Errorf("param[0] type wrong, expected 'int', got %s", m.ParamTypes[0].String())
	}
	if m.ReturnType == nil || m.ReturnType.String() != "int" {
		t.Errorf("return type wrong, expected 'int', got %v", m.ReturnType)
	}
}

func TestStructMethodUntypedBackwardCompat(t *testing.T) {
	input := `struct Point { x: int, y: int, func area() { return 0 } }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.StructStatement)
	if len(stmt.Methods) != 1 {
		t.Fatalf("expected 1 method, got %d", len(stmt.Methods))
	}

	m := stmt.Methods[0]
	if len(m.ParamTypes) != 0 {
		t.Fatalf("expected 0 param types for untyped method, got %d", len(m.ParamTypes))
	}
}

func TestTypeAliasResolution(t *testing.T) {
	input := `type UserID = int; let x: UserID = 5`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 2 {
		t.Fatalf("expected 2 statements, got %d", len(program.Statements))
	}

	aliasStmt, ok := program.Statements[0].(*ast.TypeAliasStatement)
	if !ok {
		t.Fatalf("stmt[0] not TypeAliasStatement, got %T", program.Statements[0])
	}
	if aliasStmt.Name.Value != "UserID" {
		t.Errorf("alias name wrong, expected 'UserID', got %s", aliasStmt.Name.Value)
	}

	typedStmt, ok := program.Statements[1].(*ast.TypedLetStatement)
	if !ok {
		t.Fatalf("stmt[1] not TypedLetStatement, got %T", program.Statements[1])
	}

	// The alias should have been resolved to BasicType("int")
	basicType, ok := typedStmt.TypeAnnotation.Type.(*ast.BasicType)
	if !ok {
		t.Fatalf("resolved type not BasicType, got %T", typedStmt.TypeAnnotation.Type)
	}
	if basicType.Name != "int" {
		t.Errorf("resolved type name wrong, expected 'int', got %s", basicType.Name)
	}
}

func TestTypeAliasResolutionInStruct(t *testing.T) {
	input := `type UserID = int; struct User { id: UserID }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[1].(*ast.StructStatement)
	if stmt.Fields[0].Type == nil {
		t.Fatal("field type is nil")
	}
	basicType, ok := stmt.Fields[0].Type.(*ast.BasicType)
	if !ok {
		t.Fatalf("field type not BasicType, got %T", stmt.Fields[0].Type)
	}
	if basicType.Name != "int" {
		t.Errorf("field type wrong, expected 'int', got %s", basicType.Name)
	}
}

func TestTypeCastInInfixPosition(t *testing.T) {
	input := `let ratio: float64 = (a as float64) / (b as float64)`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	_ = program
}

func TestTypedFunctionLiteralNoName(t *testing.T) {
	input := `let add = func(a: int, b: int): int { return a + b }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	letStmt := program.Statements[0].(*ast.LetStatement)
	funcLit := letStmt.Value.(*ast.TypedFunctionLiteral)

	if funcLit.Name != "" {
		t.Errorf("anonymous function should have empty name. got=%q", funcLit.Name)
	}
}

func TestTypedFunctionLiteralWithName(t *testing.T) {
	input := `let add = func add(a: int, b: int): int { return a + b }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	letStmt := program.Statements[0].(*ast.LetStatement)
	funcLit := letStmt.Value.(*ast.TypedFunctionLiteral)

	if funcLit.Name != "add" {
		t.Errorf("named function should have name 'add'. got=%q", funcLit.Name)
	}
}
