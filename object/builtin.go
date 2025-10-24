package object

type BuiltinFunction func(args ...VintObject) VintObject

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Inspect() string      { return "builtin function" }
func (b *Builtin) Type() VintObjectType { return BUILTIN_OBJ }
