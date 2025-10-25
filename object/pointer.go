package object

import "fmt"

type Pointer struct {
	Ref VintObject
}

func (p *Pointer) Type() VintObjectType {
	return POINTER_OBJ
}

func (p *Pointer) Inspect() string {
	return fmt.Sprintf("Pointer(addr=%p, value=%s)", p.Ref, p.Ref.Inspect())
}
