package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/vintlang/vintlang/ast"
)

// StructField defines a field in a struct definition
type StructField struct {
	Name    string
	Default ast.Expression // optional default value expression
}

// StructMethod defines a method in a struct definition
type StructMethod struct {
	Name       string
	Parameters []*ast.Identifier
	Defaults   map[string]ast.Expression
	Body       *ast.BlockStatement
}

// Struct represents a struct type definition (the blueprint)
type Struct struct {
	Name    string
	Fields  []StructField
	Methods map[string]*StructMethod
	Env     *Environment // the environment where the struct was defined
}

func (s *Struct) Type() VintObjectType { return STRUCT_OBJ }
func (s *Struct) Inspect() string {
	var out bytes.Buffer

	out.WriteString("struct ")
	out.WriteString(s.Name)
	out.WriteString(" { ")

	fields := []string{}
	for _, f := range s.Fields {
		fields = append(fields, f.Name)
	}
	out.WriteString(strings.Join(fields, ", "))

	if len(s.Methods) > 0 {
		methods := []string{}
		for name := range s.Methods {
			methods = append(methods, name+"()")
		}
		out.WriteString("; ")
		out.WriteString(strings.Join(methods, ", "))
	}

	out.WriteString(" }")

	return out.String()
}

// HasField checks if the struct definition has a field with the given name
func (s *Struct) HasField(name string) bool {
	for _, f := range s.Fields {
		if f.Name == name {
			return true
		}
	}
	return false
}

// GetMethod returns a method by name
func (s *Struct) GetMethod(name string) (*StructMethod, bool) {
	m, ok := s.Methods[name]
	return m, ok
}

// StructInstance represents an instantiated struct (an instance)
type StructInstance struct {
	Struct *Struct      // reference to the struct definition
	Fields *Environment // instance fields with their values
}

func (si *StructInstance) Type() VintObjectType { return STRUCT_INSTANCE_OBJ }
func (si *StructInstance) Inspect() string {
	var out bytes.Buffer

	out.WriteString(si.Struct.Name)
	out.WriteString("{")

	pairs := []string{}
	for _, f := range si.Struct.Fields {
		if val, ok := si.Fields.Get(f.Name); ok {
			pairs = append(pairs, fmt.Sprintf("%s: %s", f.Name, val.Inspect()))
		}
	}

	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// GetField returns a field value from the instance
func (si *StructInstance) GetField(name string) (VintObject, bool) {
	return si.Fields.Get(name)
}

// SetField sets a field value on the instance
func (si *StructInstance) SetField(name string, val VintObject) error {
	// Verify the field exists in the struct definition
	if !si.Struct.HasField(name) {
		return fmt.Errorf("struct '%s' has no field '%s'", si.Struct.Name, name)
	}
	si.Fields.SetScoped(name, val)
	return nil
}

// GetMethod returns a method from the struct definition
func (si *StructInstance) GetMethod(name string) (*StructMethod, bool) {
	return si.Struct.GetMethod(name)
}
