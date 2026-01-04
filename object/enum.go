package object

import (
	"bytes"
	"fmt"
	"strings"
)

// Enum represents an enum type with named constants
type Enum struct {
	Name    string
	Members map[string]VintObject // member name -> value
}

func (e *Enum) Type() VintObjectType { return ENUM_OBJ }
func (e *Enum) Inspect() string {
	var out bytes.Buffer

	out.WriteString("enum ")
	out.WriteString(e.Name)
	out.WriteString(" { ")

	members := []string{}
	for name, value := range e.Members {
		members = append(members, fmt.Sprintf("%s = %s", name, value.Inspect()))
	}

	out.WriteString(strings.Join(members, ", "))
	out.WriteString(" }")

	return out.String()
}

// GetMember returns a specific enum member value
func (e *Enum) GetMember(name string) (VintObject, bool) {
	val, ok := e.Members[name]
	return val, ok
}
