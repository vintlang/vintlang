package evaluator

import (
	"strings"
	"testing"

	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
)

func testEval(input string) object.VintObject {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return Eval(program, env)
}

func testParse(input string) *parser.Parser {
	l := lexer.New(input)
	p := parser.New(l)
	p.ParseProgram()
	return p
}

// ================================
// Parsing Tests
// ================================

func TestStructParsing(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		// Basic struct declaration
		{"struct Point { x: 0, y: 0 }", false},
		// Struct with methods
		{"struct User { name: \"\", func greet() { return \"hi\" } }", false},
		// Struct with method parameters
		{"struct Calc { val: 0, func add(n) { return this.val + n } }", false},
		// Empty struct
		{"struct Empty { }", false},
		// Struct with multiple methods
		{"struct Multi { x: 0, func a() { return 1 }, func b() { return 2 } }", false},
	}

	for _, tt := range tests {
		p := testParse(tt.input)
		errors := p.Errors()
		if tt.expectError && len(errors) == 0 {
			t.Errorf("expected parse error for input: %q", tt.input)
		}
		if !tt.expectError && len(errors) > 0 {
			t.Errorf("unexpected parse errors for input: %q - %v", tt.input, errors)
		}
	}
}

// ================================
// Struct Declaration Tests
// ================================

func TestStructDeclaration(t *testing.T) {
	input := `struct Point { x: 0, y: 0 }; Point`

	result := testEval(input)
	structObj, ok := result.(*object.Struct)
	if !ok {
		t.Fatalf("result is not Struct. got=%T (%+v)", result, result)
	}

	if structObj.Name != "Point" {
		t.Errorf("struct name wrong. got=%q, want=%q", structObj.Name, "Point")
	}

	if len(structObj.Fields) != 2 {
		t.Errorf("struct has wrong number of fields. got=%d, want=%d", len(structObj.Fields), 2)
	}

	if structObj.Fields[0].Name != "x" {
		t.Errorf("first field name wrong. got=%q, want=%q", structObj.Fields[0].Name, "x")
	}

	if structObj.Fields[1].Name != "y" {
		t.Errorf("second field name wrong. got=%q, want=%q", structObj.Fields[1].Name, "y")
	}
}

func TestStructWithMethods(t *testing.T) {
	input := `
	struct Greeter {
		name: "World"
		func greet() {
			return "Hello, " + this.name
		}
	}
	Greeter
	`

	result := testEval(input)
	structObj, ok := result.(*object.Struct)
	if !ok {
		t.Fatalf("result is not Struct. got=%T (%+v)", result, result)
	}

	if len(structObj.Methods) != 1 {
		t.Errorf("struct has wrong number of methods. got=%d, want=%d", len(structObj.Methods), 1)
	}

	if _, ok := structObj.Methods["greet"]; !ok {
		t.Error("struct does not have 'greet' method")
	}
}

// ================================
// Struct Instantiation Tests
// ================================

func TestStructInstantiationWithNamedArgs(t *testing.T) {
	input := `
	struct Point { x: 0, y: 0 }
	let p = Point(x = 10, y = 20)
	p.x + p.y
	`

	result := testEval(input)
	testStructIntegerObject(t, result, 30)
}

func TestStructInstantiationWithPositionalArgs(t *testing.T) {
	input := `
	struct Point { x: 0, y: 0 }
	let p = Point(5, 15)
	p.x + p.y
	`

	result := testEval(input)
	testStructIntegerObject(t, result, 20)
}

func TestStructInstantiationWithDefaults(t *testing.T) {
	input := `
	struct Config {
		host: "localhost"
		port: 8080
	}
	let c = Config()
	c.host
	`

	result := testEval(input)
	testStructStringObject(t, result, "localhost")
}

func TestStructPartialDefaults(t *testing.T) {
	input := `
	struct Config {
		host: "localhost"
		port: 8080
	}
	let c = Config(host = "example.com")
	c.host
	`

	result := testEval(input)
	testStructStringObject(t, result, "example.com")
}

func TestStructPartialDefaultsPort(t *testing.T) {
	input := `
	struct Config {
		host: "localhost"
		port: 8080
	}
	let c = Config(host = "example.com")
	c.port
	`

	result := testEval(input)
	testStructIntegerObject(t, result, 8080)
}

// ================================
// Struct Property Access Tests
// ================================

func TestStructPropertyAccess(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`struct P { x: 10 }; let p = P(); p.x`, int64(10)},
		{`struct P { name: "hello" }; let p = P(); p.name`, "hello"},
		{`struct P { ok: true }; let p = P(); p.ok`, true},
		{`struct P { x: 0, y: 0 }; let p = P(x = 3, y = 4); p.x`, int64(3)},
	}

	for _, tt := range tests {
		result := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int64:
			testStructIntegerObject(t, result, expected)
		case string:
			testStructStringObject(t, result, expected)
		case bool:
			testStructBooleanObject(t, result, expected)
		}
	}
}

// ================================
// Struct Property Assignment Tests
// ================================

func TestStructPropertyAssignment(t *testing.T) {
	input := `
	struct User { name: "", age: 0 }
	let u = User(name = "Alice", age = 25)
	u.name = "Bob"
	u.age = 30
	u.name
	`

	result := testEval(input)
	testStructStringObject(t, result, "Bob")
}

func TestStructPropertyAssignmentAge(t *testing.T) {
	input := `
	struct User { name: "", age: 0 }
	let u = User(name = "Alice", age = 25)
	u.age = 30
	u.age
	`

	result := testEval(input)
	testStructIntegerObject(t, result, 30)
}

// ================================
// Struct Method Tests
// ================================

func TestStructMethodCall(t *testing.T) {
	input := `
	struct User {
		name: ""
		func greet() {
			return "Hello, " + this.name
		}
	}
	let u = User(name = "Alice")
	u.greet()
	`

	result := testEval(input)
	testStructStringObject(t, result, "Hello, Alice")
}

func TestStructMethodWithParams(t *testing.T) {
	input := `
	struct Calc {
		value: 0
		func add(n) {
			return this.value + n
		}
	}
	let c = Calc(value = 10)
	c.add(5)
	`

	result := testEval(input)
	testStructIntegerObject(t, result, 15)
}

func TestStructMethodWithDefaultParams(t *testing.T) {
	input := `
	struct Greeter {
		name: "World"
		func greet(greeting = "Hello") {
			return greeting + ", " + this.name + "!"
		}
	}
	let g = Greeter(name = "VintLang")
	g.greet()
	`

	result := testEval(input)
	testStructStringObject(t, result, "Hello, VintLang!")
}

func TestStructMethodWithCustomParams(t *testing.T) {
	input := `
	struct Greeter {
		name: "World"
		func greet(greeting = "Hello") {
			return greeting + ", " + this.name + "!"
		}
	}
	let g = Greeter(name = "VintLang")
	g.greet("Hi")
	`

	result := testEval(input)
	testStructStringObject(t, result, "Hi, VintLang!")
}

func TestStructMethodMutatesFields(t *testing.T) {
	input := `
	struct Counter {
		count: 0
		func increment() {
			this.count = this.count + 1
		}
		func value() {
			return this.count
		}
	}
	let c = Counter()
	c.increment()
	c.increment()
	c.increment()
	c.value()
	`

	result := testEval(input)
	testStructIntegerObject(t, result, 3)
}

func TestStructMethodReturnsValue(t *testing.T) {
	input := `
	struct Rect {
		width: 0
		height: 0
		func area() {
			return this.width * this.height
		}
	}
	let r = Rect(width = 10, height = 5)
	r.area()
	`

	result := testEval(input)
	testStructIntegerObject(t, result, 50)
}

func TestStructMethodCallsOtherMethod(t *testing.T) {
	input := `
	struct Rect {
		width: 0
		height: 0
		func area() {
			return this.width * this.height
		}
		func perimeter() {
			return 2 * (this.width + this.height)
		}
		func describe() {
			return string(this.area()) + ":" + string(this.perimeter())
		}
	}
	let r = Rect(width = 10, height = 5)
	r.describe()
	`

	result := testEval(input)
	testStructStringObject(t, result, "50:30")
}

// ================================
// Struct Instance Independence Tests
// ================================

func TestStructInstanceIndependence(t *testing.T) {
	input := `
	struct Counter {
		count: 0
		func inc() { this.count = this.count + 1 }
		func val() { return this.count }
	}
	let c1 = Counter()
	let c2 = Counter()
	c1.inc()
	c1.inc()
	c2.inc()
	c1.val() * 10 + c2.val()
	`

	result := testEval(input)
	// c1 incremented 2 times, c2 incremented 1 time
	// c1.val() = 2, c2.val() = 1 => 2*10+1 = 21
	testStructIntegerObject(t, result, 21)
}

// ================================
// Struct Type Tests
// ================================

func TestStructType(t *testing.T) {
	input := `
	struct User { name: "" }
	let u = User(name = "Alice")
	type(u)
	`

	result := testEval(input)
	testStructStringObject(t, result, "User")
}

func TestStructDefinitionType(t *testing.T) {
	input := `
	struct User { name: "" }
	type(User)
	`

	result := testEval(input)
	testStructStringObject(t, result, "struct:User")
}

// ================================
// Struct Inspect Tests
// ================================

func TestStructInstanceInspect(t *testing.T) {
	input := `
	struct Point { x: 0, y: 0 }
	let p = Point(x = 3, y = 4)
	string(p)
	`

	result := testEval(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("result is not String. got=%T (%+v)", result, result)
	}
	if !strings.Contains(str.Value, "Point") {
		t.Errorf("inspect should contain struct name. got=%q", str.Value)
	}
	if !strings.Contains(str.Value, "x:") || !strings.Contains(str.Value, "y:") {
		t.Errorf("inspect should contain field names. got=%q", str.Value)
	}
}

// ================================
// Struct with Arrays and Dicts
// ================================

func TestStructInArray(t *testing.T) {
	input := `
	struct Item { name: "", price: 0, func label() { return this.name + ": $" + string(this.price) } }
	let items = [Item(name = "Apple", price = 1), Item(name = "Banana", price = 2)]
	items[0].label() + " | " + items[1].label()
	`

	result := testEval(input)
	testStructStringObject(t, result, "Apple: $1 | Banana: $2")
}

// ================================
// Struct Return from Method
// ================================

func TestStructReturnNewInstance(t *testing.T) {
	input := `
	struct Vector {
		x: 0
		y: 0
		func scale(factor) {
			return Vector(x = this.x * factor, y = this.y * factor)
		}
		func display() {
			return string(this.x) + "," + string(this.y)
		}
	}
	let v = Vector(x = 3, y = 4)
	let v2 = v.scale(2)
	v2.display()
	`

	result := testEval(input)
	testStructStringObject(t, result, "6,8")
}

// ================================
// Error Cases
// ================================

func TestStructUnknownFieldError(t *testing.T) {
	input := `
	struct Point { x: 0, y: 0 }
	let p = Point(x = 1, y = 2)
	p.z
	`

	result := testEval(input)
	if result.Type() != object.ERROR_OBJ {
		t.Fatalf("expected error for accessing unknown field, got=%T (%+v)", result, result)
	}
	errObj := result.(*object.Error)
	if !strings.Contains(errObj.Message, "no field") {
		t.Errorf("error message should mention unknown field. got=%q", errObj.Message)
	}
}

func TestStructUnknownFieldInConstructor(t *testing.T) {
	input := `
	struct Point { x: 0, y: 0 }
	let p = Point(x = 1, z = 2)
	`

	result := testEval(input)
	if result.Type() != object.ERROR_OBJ {
		t.Fatalf("expected error for unknown field in constructor, got=%T (%+v)", result, result)
	}
}

func TestStructMissingRequiredField(t *testing.T) {
	// A field with no default and no value should error
	input := `
	struct Required { x }
	let p = Required()
	`

	result := testEval(input)
	if result.Type() != object.ERROR_OBJ {
		t.Fatalf("expected error for missing required field, got=%T (%+v)", result, result)
	}
}

func TestStructAssignToUnknownField(t *testing.T) {
	input := `
	struct Point { x: 0, y: 0 }
	let p = Point()
	p.z = 10
	`

	result := testEval(input)
	if result.Type() != object.ERROR_OBJ {
		t.Fatalf("expected error for assigning to unknown field, got=%T (%+v)", result, result)
	}
}

func TestStructUnknownMethodError(t *testing.T) {
	input := `
	struct Point { x: 0, y: 0 }
	let p = Point()
	p.foo()
	`

	result := testEval(input)
	if result.Type() != object.ERROR_OBJ {
		t.Fatalf("expected error for calling unknown method, got=%T (%+v)", result, result)
	}
}

func TestStructTooManyPositionalArgs(t *testing.T) {
	input := `
	struct Point { x: 0, y: 0 }
	let p = Point(1, 2, 3)
	`

	result := testEval(input)
	if result.Type() != object.ERROR_OBJ {
		t.Fatalf("expected error for too many args, got=%T (%+v)", result, result)
	}
}

// ================================
// Struct with main() function
// ================================

func TestStructWithMainFunction(t *testing.T) {
	input := `
	struct User {
		name: ""
		age: 0
		func greet() {
			return "Hello, " + this.name
		}
	}

	func main() {
		let u = User(name = "Alice", age = 30)
		return u.greet()
	}
	`

	result := testEval(input)
	testStructStringObject(t, result, "Hello, Alice")
}

// ================================
// Helper functions
// ================================

func testStructIntegerObject(t *testing.T, obj object.VintObject, expected int64) {
	t.Helper()
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Fatalf("object is not Integer. got=%T (%+v)", obj, obj)
	}
	if result.Value != expected {
		t.Errorf("integer value wrong. got=%d, want=%d", result.Value, expected)
	}
}

func testStructStringObject(t *testing.T, obj object.VintObject, expected string) {
	t.Helper()
	result, ok := obj.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", obj, obj)
	}
	if result.Value != expected {
		t.Errorf("string value wrong. got=%q, want=%q", result.Value, expected)
	}
}

func testStructBooleanObject(t *testing.T, obj object.VintObject, expected bool) {
	t.Helper()
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Fatalf("object is not Boolean. got=%T (%+v)", obj, obj)
	}
	if result.Value != expected {
		t.Errorf("boolean value wrong. got=%t, want=%t", result.Value, expected)
	}
}
