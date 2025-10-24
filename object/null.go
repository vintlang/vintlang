package object

type Null struct{}

func (n *Null) Inspect() string      { return "null" }
func (n *Null) Type() VintObjectType { return NULL_OBJ }

// Method dynamically dispatches null-related methods
func (n *Null) Method(method string, args []VintObject) VintObject {
	switch method {
	case "isNull":
		return n.isNull(args)
	case "coalesce":
		return n.coalesce(args)
	case "ifNull":
		return n.ifNull(args)
	case "toString":
		return n.toString(args)
	case "equals":
		return n.equals(args)
	default:
		return newError("Method '%s' is not supported for Null objects", method)
	}
}

// isNull always returns true for null objects
func (n *Null) isNull(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("isNull() expects 0 arguments, got %d", len(args))
	}
	return &Boolean{Value: true}
}

// coalesce returns the first non-null value from arguments
func (n *Null) coalesce(args []VintObject) VintObject {
	for _, arg := range args {
		if arg.Type() != NULL_OBJ {
			return arg
		}
	}
	return n // return null if all arguments are null
}

// ifNull returns the provided value if this is null, otherwise returns this (which is null)
func (n *Null) ifNull(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("ifNull() expects 1 argument, got %d", len(args))
	}
	return args[0] // since this is null, always return the provided value
}

// toString returns string representation of null
func (n *Null) toString(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("toString() expects 0 arguments, got %d", len(args))
	}
	return &String{Value: "null"}
}

// equals checks if another value is also null
func (n *Null) equals(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("equals() expects 1 argument, got %d", len(args))
	}
	return &Boolean{Value: args[0].Type() == NULL_OBJ}
}
