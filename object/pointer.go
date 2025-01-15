package object

import "fmt"

type Pointer struct {
	Ref Object 
}

func (p *Pointer) Type() ObjectType {
	return POINTER_OBJ
}

func (p *Pointer) Inspect() string {
	return fmt.Sprintf("Pointer(%s)", p.Ref.Inspect())
}

const POINTER_OBJ = "POINTER"
