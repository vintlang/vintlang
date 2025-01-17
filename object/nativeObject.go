package object

import "fmt"

type NativeObject struct {
	Value interface{}
}

func (n *NativeObject) Type() ObjectType { return "NATIVE_OBJ" }
func (n *NativeObject) Inspect() string {
	return fmt.Sprintf("NativeObject(%v)", n.Value)
}
