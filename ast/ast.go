package ast

import (
	"bytes"
)

// Node represents any node in the AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents any statement node
type Statement interface {
	Node
	statementNode()
}

// Expression represents any expression node
type Expression interface {
	Node
	expressionNode()
}

// Program represents the root of every AST our parser produces
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
