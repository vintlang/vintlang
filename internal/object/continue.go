package object

type Continue struct{}

func (c *Continue) Type() VintObjectType { return CONTINUE_OBJ }
func (c *Continue) Inspect() string      { return "continue" }
