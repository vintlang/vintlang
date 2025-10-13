package object

import (
	"fmt"
	"strings"

	"github.com/vintlang/vintlang/ast"
)

type Package struct {
	Name         *ast.Identifier
	Env          *Environment
	Scope        *Environment
	PrivateNames map[string]bool // Track private identifiers (those starting with _)
}

func (p *Package) Type() ObjectType { return PACKAGE_OBJ }
func (p *Package) Inspect() string {
	return fmt.Sprintf("package: %s", p.Name.Value)
}

// IsPrivate checks if an identifier is private (starts with underscore)
func (p *Package) IsPrivate(name string) bool {
	return strings.HasPrefix(name, "_")
}

// GetPublic returns a public member of the package (not starting with _)
func (p *Package) GetPublic(name string) (VintObject, bool) {
	if p.IsPrivate(name) {
		return nil, false // Private members are not accessible from outside
	}
	return p.Scope.Get(name)
}

// GetPrivate returns any member (used internally within package)
func (p *Package) GetPrivate(name string) (VintObject, bool) {
	return p.Scope.Get(name)
}

// DefinePrivate marks a name as private
func (p *Package) DefinePrivate(name string) {
	if p.PrivateNames == nil {
		p.PrivateNames = make(map[string]bool)
	}
	p.PrivateNames[name] = true
}
