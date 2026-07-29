package evaluator

import (
	"github.com/vintlang/vintlang/internal/ast"
	"github.com/vintlang/vintlang/internal/object"
)

var typeNameToRuntime = map[string]object.VintObjectType{
	"int":     object.INTEGER_OBJ,
	"int8":    object.INTEGER_OBJ,
	"int16":   object.INTEGER_OBJ,
	"int32":   object.INTEGER_OBJ,
	"int64":   object.INTEGER_OBJ,
	"uint":    object.INTEGER_OBJ,
	"uint8":   object.INTEGER_OBJ,
	"uint16":  object.INTEGER_OBJ,
	"uint32":  object.INTEGER_OBJ,
	"uint64":  object.INTEGER_OBJ,
	"byte":    object.INTEGER_OBJ,
	"float32": object.FLOAT_OBJ,
	"float64": object.FLOAT_OBJ,
	"string":  object.STRING_OBJ,
	"bool":    object.BOOLEAN_OBJ,
	"nil":     object.NULL_OBJ,
	"error":   object.ERROR_OBJ,
}

// compatible checks if a runtime value matches a declared type.
func compatible(declared ast.Type, obj object.VintObject) bool {
	if obj == nil {
		return false
	}
	switch t := declared.(type) {
	case *ast.BasicType:
		if t.Name == "any" {
			return true
		}
		// Struct instances match by struct name
		if si, ok := obj.(*object.StructInstance); ok {
			return si.Struct.Name == t.Name
		}
		// Enum instances match by enum name
		if e, ok := obj.(*object.Enum); ok {
			return e.Name == t.Name
		}
		// Check built-in type map
		expected, ok := typeNameToRuntime[t.Name]
		if ok && obj.Type() == expected {
			return true
		}
		// error type also matches CUSTOM_ERROR
		if t.Name == "error" && obj.Type() == object.CUSTOM_ERROR_OBJ {
			return true
		}
		return false
	case *ast.ArrayType:
		return obj.Type() == object.ARRAY_OBJ
	case *ast.FixedArrayType:
		return obj.Type() == object.ARRAY_OBJ
	case *ast.DictType:
		return obj.Type() == object.DICT_OBJ
	case *ast.PointerType:
		return obj.Type() == object.POINTER_OBJ
	case *ast.ChannelType:
		return obj.Type() == object.CHANNEL_OBJ
	case *ast.FunctionType:
		return obj.Type() == object.FUNCTION_OBJ || obj.Type() == object.BUILTIN_OBJ
	case *ast.MultiReturnType:
		// Multi-return checks each component
		return true
	default:
		return false
	}
}

// inferType returns the ast.Type that matches a runtime object.
func inferType(obj object.VintObject) ast.Type {
	switch obj := obj.(type) {
	case *object.Integer:
		return &ast.BasicType{Name: "int"}
	case *object.Float:
		return &ast.BasicType{Name: "float64"}
	case *object.String:
		return &ast.BasicType{Name: "string"}
	case *object.Boolean:
		return &ast.BasicType{Name: "bool"}
	case *object.Null:
		return &ast.BasicType{Name: "nil"}
	case *object.Error:
		return &ast.BasicType{Name: "error"}
	case *object.CustomError:
		return &ast.BasicType{Name: "error"}
	case *object.StructInstance:
		return &ast.BasicType{Name: obj.Struct.Name}
	case *object.Enum:
		return &ast.BasicType{Name: obj.Name}
	case *object.Array:
		return &ast.BasicType{Name: "[]any"}
	case *object.Dict:
		return &ast.BasicType{Name: "{any: any}"}
	case *object.Pointer:
		return &ast.BasicType{Name: "*any"}
	default:
		return &ast.BasicType{Name: "any"}
	}
}

func evalTypeCast(node *ast.TypeCastExpression, env *object.Environment) object.VintObject {
	val := Eval(node.Expression, env)
	if isError(val) {
		return val
	}

	switch t := node.TargetType.(type) {
	case *ast.BasicType:
		switch t.Name {
		case "int", "int8", "int16", "int32", "int64",
			"uint", "uint8", "uint16", "uint32", "uint64", "byte":
			return convertToInteger(val)
		case "float32", "float64":
			return convertToFloat(val)
		case "string":
			return convertToString(val)
		case "bool":
			return convertToBoolean(val)
		default:
			return newError("cannot cast %s to %s", val.Type(), t.Name)
		}
	default:
		return newError("cannot cast %s to %s", val.Type(), node.TargetType.String())
	}
}

func evalTypeCheck(node *ast.TypeCheckExpression, env *object.Environment) object.VintObject {
	val := Eval(node.Expression, env)
	if isError(val) {
		return val
	}

	return nativeBoolToBooleanObject(compatible(node.CheckType, val))
}

// zeroValueFromType returns the zero value for a given type annotation.
func zeroValueFromType(t ast.Type) object.VintObject {
	switch t := t.(type) {
	case *ast.BasicType:
		switch t.Name {
		case "int", "int8", "int16", "int32", "int64",
			"uint", "uint8", "uint16", "uint32", "uint64", "byte":
			return &object.Integer{Value: 0}
		case "float32", "float64":
			return &object.Float{Value: 0.0}
		case "string":
			return &object.String{Value: ""}
		case "bool":
			return &object.Boolean{Value: false}
		default:
			return NULL
		}
	case *ast.ArrayType, *ast.FixedArrayType, *ast.DictType,
		*ast.PointerType, *ast.ChannelType, *ast.FunctionType:
		return NULL
	default:
		return NULL
	}
}
