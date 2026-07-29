package object

import "fmt"

type Instance struct {
	Package *Package
	Env     *Environment
}

func (i *Instance) Type() VintObjectType { return INSTANCE }
func (i *Instance) Inspect() string {
	return fmt.Sprintf("Package: %s", i.Package.Name.Value)
}
