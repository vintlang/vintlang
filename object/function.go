package object

import (
	"bytes"
	"strings"

	"github.com/vintlang/vintlang/ast"
)

type Function struct {
	Name        string
	Parameters  []*ast.Identifier
	Defaults    map[string]ast.Expression
	Body        *ast.BlockStatement
	Env         *Environment
	IsAsync     bool // Support for async handlers
	IsStreaming bool // Support for streaming responses
}

func (f *Function) Type() VintObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("func")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
