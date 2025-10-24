package object

import "fmt"

type NativeObject struct {
	Value any
}

func (n *NativeObject) Type() VintObjectType { return "NATIVE_OBJ" }
func (n *NativeObject) Inspect() string {
	return fmt.Sprintf("NativeObject(%v)", n.Value)
}
