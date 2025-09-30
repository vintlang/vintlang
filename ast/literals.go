package ast

import (
	"bytes"
	"strings"

	"github.com/vintlang/vintlang/token"
)

// Identifier represents identifiers like variable names
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// IntegerLiteral represents integer values like 5, 10, 342343
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// FloatLiteral represents floating point values like 3.14, 2.718
type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) String() string       { return fl.Token.Literal }

// StringLiteral represents string values like "hello", "world"
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

// Boolean represents boolean values like true, false
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// ArrayLiteral represents array literals like [1, 2, 3]
type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// DictLiteral represents dictionary/hash literals like {key: value}
type DictLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (dl *DictLiteral) expressionNode()      {}
func (dl *DictLiteral) TokenLiteral() string { return dl.Token.Literal }
func (dl *DictLiteral) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range dl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("(")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// FunctionLiteral represents function literals like fn(x, y) { return x + y; }
type FunctionLiteral struct {
	Token      token.Token
	Name       string
	Parameters []*Identifier
	Defaults   map[string]Expression
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// AsyncFunctionLiteral represents an async function
type AsyncFunctionLiteral struct {
	Token      token.Token
	Name       string
	Parameters []*Identifier
	Defaults   map[string]Expression
	Body       *BlockStatement
}

func (afl *AsyncFunctionLiteral) expressionNode()      {}
func (afl *AsyncFunctionLiteral) TokenLiteral() string { return afl.Token.Literal }
func (afl *AsyncFunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range afl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("async func")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(afl.Body.String())

	return out.String()
}

// Null represents null values
type Null struct {
	Token token.Token
}

func (n *Null) expressionNode()      {}
func (n *Null) TokenLiteral() string { return n.Token.Literal }
func (n *Null) String() string       { return n.Token.Literal }

// Break represents break statements in loops
type Break struct {
	Statement
	Token token.Token // the 'break' token
}

func (b *Break) expressionNode()      {}
func (b *Break) TokenLiteral() string { return b.Token.Literal }
func (b *Break) String() string       { return b.Token.Literal }

// Continue represents continue statements in loops
type Continue struct {
	Statement
	Token token.Token // the 'continue' token
}

func (c *Continue) expressionNode()      {}
func (c *Continue) TokenLiteral() string { return c.Token.Literal }
func (c *Continue) String() string       { return c.Token.Literal }

// At represents the @ symbol
type At struct {
	Token token.Token
}

func (a *At) expressionNode()      {}
func (a *At) TokenLiteral() string { return a.Token.Literal }
func (a *At) String() string       { return "@" }
