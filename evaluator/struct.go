package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

// evalStructStatement evaluates struct declarations and registers them in the environment
func evalStructStatement(node *ast.StructStatement, env *object.Environment) object.VintObject {
	structDef := &object.Struct{
		Name:    node.Name.Value,
		Fields:  make([]object.StructField, 0, len(node.Fields)),
		Methods: make(map[string]*object.StructMethod),
		Env:     env,
	}

	// Register fields
	for _, f := range node.Fields {
		field := object.StructField{
			Name:    f.Name.Value,
			Default: f.Default,
		}
		structDef.Fields = append(structDef.Fields, field)
	}

	// Register methods
	for _, m := range node.Methods {
		method := &object.StructMethod{
			Name:       m.Name.Value,
			Parameters: m.Parameters,
			Defaults:   m.Defaults,
			Body:       m.Body,
		}
		structDef.Methods[m.Name.Value] = method
	}

	// Define the struct type in the current environment
	return env.Define(node.Name.Value, structDef)
}

// instantiateStruct creates a new instance of a struct with the given field values
func instantiateStruct(structDef *object.Struct, fieldArgs map[string]object.VintObject, line int) object.VintObject {
	instanceEnv := object.NewEnvironment()

	// Initialize all fields with defaults first, then override with provided values
	for _, field := range structDef.Fields {
		if val, ok := fieldArgs[field.Name]; ok {
			// User provided a value for this field
			instanceEnv.Define(field.Name, val)
		} else if field.Default != nil {
			// Use the default value
			defaultVal := Eval(field.Default, structDef.Env)
			if isError(defaultVal) {
				return defaultVal
			}
			instanceEnv.Define(field.Name, defaultVal)
		} else {
			// No value provided and no default — error
			return newError("Line %d: Missing value for field '%s' in struct '%s'",
				line, field.Name, structDef.Name)
		}
	}

	// Verify no unknown fields were provided
	for name := range fieldArgs {
		if !structDef.HasField(name) {
			return newError("Line %d: Struct '%s' has no field '%s'",
				line, structDef.Name, name)
		}
	}

	instance := &object.StructInstance{
		Struct: structDef,
		Fields: instanceEnv,
	}

	return instance
}

// callStructMethod calls a method on a struct instance
func callStructMethod(instance *object.StructInstance, methodName string, args []object.VintObject, defs map[string]object.VintObject, line int) object.VintObject {
	method, ok := instance.GetMethod(methodName)
	if !ok {
		return newError("Line %d: Struct '%s' has no method '%s'",
			line, instance.Struct.Name, methodName)
	}

	// Create a new environment for the method execution
	// The method's environment encloses the struct definition's environment
	methodEnv := object.NewEnclosedEnvironment(instance.Struct.Env)

	// Bind 'this' to the struct instance
	methodEnv.Define("this", instance)

	// Bind parameters
	for i, param := range method.Parameters {
		if i < len(args) {
			methodEnv.Define(param.Value, args[i])
		} else if defVal, ok := method.Defaults[param.Value]; ok {
			evaluated := Eval(defVal, methodEnv)
			if isError(evaluated) {
				return evaluated
			}
			methodEnv.Define(param.Value, evaluated)
		} else {
			return newError("Line %d: Missing argument '%s' for method '%s' in struct '%s'",
				line, param.Value, methodName, instance.Struct.Name)
		}
	}

	// Execute the method body
	result := Eval(method.Body, methodEnv)

	// Copy back any field changes made through 'this'
	// This ensures mutations via 'this.field = value' persist on the instance
	return unwrapReturnValue(result)
}

// evalStructCall handles struct instantiation via call syntax:
// User(name = "Alice", age = 30) or User("Alice", 30)
func evalStructCall(node *ast.CallExpression, structDef *object.Struct, env *object.Environment) object.VintObject {
	fieldArgs := make(map[string]object.VintObject)

	positionalIndex := 0

	for _, exprr := range node.Arguments {
		switch exp := exprr.(type) {
		case *ast.Assign:
			// Keyword argument: name = "Alice"
			val := Eval(exp.Value, env)
			if isError(val) {
				return val
			}
			fieldArgs[exp.Name.Value] = val
		default:
			// Positional argument: matched to fields by order
			evaluated := Eval(exp, env)
			if isError(evaluated) {
				return evaluated
			}
			if positionalIndex < len(structDef.Fields) {
				fieldArgs[structDef.Fields[positionalIndex].Name] = evaluated
				positionalIndex++
			} else {
				return newError("Line %d: Too many arguments for struct '%s'",
					node.Token.Line, structDef.Name)
			}
		}
	}

	return instantiateStruct(structDef, fieldArgs, node.Token.Line)
}

// evalStructLiteral handles struct instantiation via brace syntax:
// User{name: "Alice", age: 30}
func evalStructLiteral(node *ast.StructLiteral, env *object.Environment) object.VintObject {
	nameObj := Eval(node.Name, env)
	if isError(nameObj) {
		return nameObj
	}

	structDef, ok := nameObj.(*object.Struct)
	if !ok {
		return newError("'%s' is not a struct type", node.Name.String())
	}

	fieldArgs := make(map[string]object.VintObject)
	for name, expr := range node.Fields {
		val := Eval(expr, env)
		if isError(val) {
			return val
		}
		fieldArgs[name] = val
	}

	return instantiateStruct(structDef, fieldArgs, node.Token.Line)
}
