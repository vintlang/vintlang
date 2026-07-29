package object

type ReturnValue struct {
	Value VintObject
}

func (rv *ReturnValue) Inspect() string      { return rv.Value.Inspect() }
func (rv *ReturnValue) Type() VintObjectType { return RETURN_VALUE_OBJ }
