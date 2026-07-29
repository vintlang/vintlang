package object

import "github.com/vintlang/vintlang/internal/ast"

type BuiltinFunction func(args ...VintObject) VintObject

type Builtin struct {
	Fn         BuiltinFunction
	ParamTypes []ast.Type // nil for untyped params (no enforcement)
	ReturnType ast.Type   // nil for unknown/void return
}

func (b *Builtin) Inspect() string      { return "builtin function" }
func (b *Builtin) Type() VintObjectType { return BUILTIN_OBJ }
