package evaluator

import (
	"math"
	"strings"

	"github.com/vintlang/vintlang/object"
)

func evalInfixExpression(operator string, left, right object.Object, line int) object.Object {
	if right == nil {
		return newError("Line %d: Right operand is missing or invalid for '%s' operator", line, operator)
	}
	if left == nil {
		return newError("Line %d: Left operand is missing or invalid for '%s' operator", line, operator)
	}
	switch {

	case operator == "in":
		return evalInExpression(left, right, line)

	case operator == "??":
		// Null coalescing: return left if not null, otherwise return right
		if left.Type() != object.NULL_OBJ {
			return left
		}
		return right

	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right, line)

	case operator == "+" && left.Type() == object.DICT_OBJ && right.Type() == object.DICT_OBJ:
		leftVal := left.(*object.Dict).Pairs
		rightVal := right.(*object.Dict).Pairs
		pairs := make(map[object.HashKey]object.DictPair)
		for k, v := range leftVal {
			pairs[k] = v
		}
		for k, v := range rightVal {
			pairs[k] = v
		}
		return &object.Dict{Pairs: pairs}

	case operator == "+" && left.Type() == object.ARRAY_OBJ && right.Type() == object.ARRAY_OBJ:
		leftVal := left.(*object.Array).Elements
		rightVal := right.(*object.Array).Elements
		elements := append(leftVal, rightVal...)
		return &object.Array{Elements: elements}

	case operator == "*" && left.Type() == object.ARRAY_OBJ && right.Type() == object.INTEGER_OBJ:
		leftVal := left.(*object.Array).Elements
		rightVal := int(right.(*object.Integer).Value)
		elements := leftVal
		for i := rightVal; i > 1; i-- {
			elements = append(elements, leftVal...)
		}
		return &object.Array{Elements: elements}

	case operator == "*" && left.Type() == object.INTEGER_OBJ && right.Type() == object.ARRAY_OBJ:
		leftVal := int(left.(*object.Integer).Value)
		rightVal := right.(*object.Array).Elements
		elements := rightVal
		for i := leftVal; i > 1; i-- {
			elements = append(elements, rightVal...)
		}
		return &object.Array{Elements: elements}

	case operator == "*" && left.Type() == object.STRING_OBJ && right.Type() == object.INTEGER_OBJ:
		leftVal := left.(*object.String).Value
		rightVal := right.(*object.Integer).Value
		return &object.String{Value: strings.Repeat(leftVal, int(rightVal))}

	case operator == "*" && left.Type() == object.INTEGER_OBJ && right.Type() == object.STRING_OBJ:
		leftVal := left.(*object.Integer).Value
		rightVal := right.(*object.String).Value
		return &object.String{Value: strings.Repeat(rightVal, int(leftVal))}

	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right, line)

	case left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ:
		return evalFloatInfixExpression(operator, left, right, line)

	case left.Type() == object.INTEGER_OBJ && right.Type() == object.FLOAT_OBJ:
		return evalFloatIntegerInfixExpression(operator, left, right, line)

	case left.Type() == object.FLOAT_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalFloatIntegerInfixExpression(operator, left, right, line)

	case operator == "==":
		return nativeBoolToBooleanObject(left == right)

	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ:
		return evalBooleanInfixExpression(operator, left, right, line)

	case left.Type() != right.Type():
		return newError("Line %d: Type mismatch: cannot use '%s' operator between %s and %s. Consider type conversion", line, operator, left.Type(), right.Type())

	default:
		return newError("Line %d: Unsupported operation: '%s' operator is not supported for %s and %s types",
			line, operator, left.Type(), right.Type())
	}
}

func evalFloatIntegerInfixExpression(operator string, left, right object.Object, line int) object.Object {
	var leftVal, rightVal float64
	if left.Type() == object.FLOAT_OBJ {
		leftVal = left.(*object.Float).Value
		rightVal = float64(right.(*object.Integer).Value)
	} else {
		leftVal = float64(left.(*object.Integer).Value)
		rightVal = right.(*object.Float).Value
	}

	var val float64
	switch operator {
	case "+":
		val = leftVal + rightVal
	case "-":
		val = leftVal - rightVal
	case "*":
		val = leftVal * rightVal
	case "**":
		val = math.Pow(float64(leftVal), float64(rightVal))
	case "/":
		val = leftVal / rightVal
	case "%":
		val = math.Mod(leftVal, rightVal)
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("Line %d: Unsupported numeric operation: '%s' operator cannot be used with %s and %s",
			line, operator, left.Type(), right.Type())
	}

	if math.Mod(val, 1) == 0 {
		return &object.Integer{Value: int64(val)}
	} else {
		return &object.Float{Value: val}
	}
}

func evalStringInfixExpression(operator string, left, right object.Object, line int) object.Object {

	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("Line %d: Unsupported string operation: '%s' operator cannot be used with strings", line, operator)
	}
}

func evalBooleanInfixExpression(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(*object.Boolean).Value
	rightVal := right.(*object.Boolean).Value

	switch operator {
	case "&&":
		return nativeBoolToBooleanObject(leftVal && rightVal)
	case "||":
		return nativeBoolToBooleanObject(leftVal || rightVal)
	default:
		return newError("Line %d: Unsupported boolean operation: '%s' operator cannot be used with boolean values", line, operator)
	}
}

func evalFloatInfixExpression(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "**":
		return &object.Float{Value: math.Pow(float64(leftVal), float64(rightVal))}
	case "/":
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("Line %d: Unsupported float operation: '%s' operator cannot be used with floating point numbers",
			line, operator)
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "**":
		return &object.Float{Value: float64(math.Pow(float64(leftVal), float64(rightVal)))}
	case "/":
		x := float64(leftVal) / float64(rightVal)
		if math.Mod(x, 1) == 0 {
			return &object.Integer{Value: int64(x)}
		} else {
			return &object.Float{Value: x}
		}
	case "%":
		return &object.Integer{Value: leftVal % rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("Line %d: Unsupported integer operation: '%s' operator cannot be used with integer values",
			line, operator)
	}
}
