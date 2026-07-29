package object

type Byte struct {
	Value  []byte
	String string
}

func (b *Byte) Inspect() string      { return "b" + b.String }
func (b *Byte) Type() VintObjectType { return BYTE_OBJ }
