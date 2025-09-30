package object

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	if b.Value {
		return "true"
	} else {
		return "false"
	}
}
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

func (b *Boolean) Method(method string, args []VintObject) VintObject {
	switch method {
	case "to_string":
		return b.toString(args)
	case "to_int":
		return b.toInt(args)
	case "negate":
		return b.negate(args)
	case "and":
		return b.and(args)
	case "or":
		return b.or(args)
	case "xor":
		return b.xor(args)
	case "implies":
		return b.implies(args)
	case "equivalent":
		return b.equivalent(args)
	case "nor":
		return b.nor(args)
	case "nand":
		return b.nand(args)
	default:
		return newError("Method '%s' is not supported for Boolean objects", method)
	}
}

func (b *Boolean) toString(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("to_string() expects 0 arguments, got %d", len(args))
	}
	return &String{Value: b.Inspect()}
}

func (b *Boolean) toInt(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("to_int() expects 0 arguments, got %d", len(args))
	}

	if b.Value {
		return &Integer{Value: 1}
	}
	return &Integer{Value: 0}
}

func (b *Boolean) negate(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("negate() expects 0 arguments, got %d", len(args))
	}
	return &Boolean{Value: !b.Value}
}

func (b *Boolean) and(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("and() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Boolean)
	if !ok {
		return newError("Argument must be a Boolean")
	}

	return &Boolean{Value: b.Value && other.Value}
}

func (b *Boolean) or(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("or() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Boolean)
	if !ok {
		return newError("Argument must be a Boolean")
	}

	return &Boolean{Value: b.Value || other.Value}
}

func (b *Boolean) xor(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("xor() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Boolean)
	if !ok {
		return newError("Argument must be a Boolean")
	}

	return &Boolean{Value: b.Value != other.Value}
}

func (b *Boolean) implies(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("implies() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Boolean)
	if !ok {
		return newError("Argument must be a Boolean")
	}

	// p -> q is equivalent to !p || q
	return &Boolean{Value: !b.Value || other.Value}
}

func (b *Boolean) equivalent(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("equivalent() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Boolean)
	if !ok {
		return newError("Argument must be a Boolean")
	}

	// p <-> q is equivalent to (p && q) || (!p && !q)
	return &Boolean{Value: b.Value == other.Value}
}

func (b *Boolean) nor(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("nor() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Boolean)
	if !ok {
		return newError("Argument must be a Boolean")
	}

	// NOR is !(p || q)
	return &Boolean{Value: !(b.Value || other.Value)}
}

func (b *Boolean) nand(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("nand() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Boolean)
	if !ok {
		return newError("Argument must be a Boolean")
	}

	// NAND is !(p && q)
	return &Boolean{Value: !(b.Value && other.Value)}
}
