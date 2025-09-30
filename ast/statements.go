package ast

import (
	"bytes"
	"strings"
	"github.com/vintlang/vintlang/token"
)

// LetStatement represents variable declarations like: let x = 5;
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

// ConstStatement represents constant declarations like: const PI = 3.14;
type ConstStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (cs *ConstStatement) statementNode()       {}
func (cs *ConstStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ConstStatement) String() string {
	var out bytes.Buffer

	out.WriteString(cs.TokenLiteral() + " ")
	out.WriteString(cs.Name.String())
	out.WriteString(" = ")

	if cs.Value != nil {
		out.WriteString(cs.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

// ReturnStatement represents return statements like: return x + y;
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement represents expressions used as statements
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// BlockStatement represents block statements like { ... }
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// IncludeStatement represents include statements like: include "file.vint"
type IncludeStatement struct {
	Token token.Token // the 'include' token
	Path  Expression
}

func (is *IncludeStatement) statementNode()       {}
func (is *IncludeStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IncludeStatement) String() string {
	var out bytes.Buffer
	out.WriteString(is.TokenLiteral() + " ")
	out.WriteString(is.Path.String())
	return out.String()
}

// GoStatement represents a go statement for concurrent execution
type GoStatement struct {
	Token      token.Token
	Expression Expression
}

func (gs *GoStatement) statementNode()       {}
func (gs *GoStatement) TokenLiteral() string { return gs.Token.Literal }
func (gs *GoStatement) String() string {
	var out bytes.Buffer
	out.WriteString("go ")
	out.WriteString(gs.Expression.String())
	return out.String()
}



// PackageBlock represents package block statements
type PackageBlock struct {
	Token      token.Token
	Statements []Statement
}

func (pb *PackageBlock) statementNode()       {}
func (pb *PackageBlock) TokenLiteral() string { return pb.Token.Literal }
func (pb *PackageBlock) String() string {
	var out bytes.Buffer

	for _, s := range pb.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// ErrorDeclaration represents custom error type definitions like: error FileNotFound(path)
type ErrorDeclaration struct {
	Token      token.Token    // the 'error' token
	Name       *Identifier    // error type name
	Parameters []*Identifier  // parameter names
}

func (ed *ErrorDeclaration) statementNode()       {}
func (ed *ErrorDeclaration) expressionNode()      {}
func (ed *ErrorDeclaration) TokenLiteral() string { return ed.Token.Literal }
func (ed *ErrorDeclaration) String() string {
	var out bytes.Buffer
	
	params := []string{}
	for _, p := range ed.Parameters {
		params = append(params, p.String())
	}
	
	out.WriteString("error ")
	out.WriteString(ed.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	
	return out.String()
}

// ThrowStatement represents throw statements like: throw FileNotFound("/missing.txt")
type ThrowStatement struct {
	Token     token.Token  // the 'throw' token
	ErrorExpr Expression   // the error expression to throw
}

func (ts *ThrowStatement) statementNode()       {}
func (ts *ThrowStatement) expressionNode()      {}
func (ts *ThrowStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *ThrowStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString("throw ")
	if ts.ErrorExpr != nil {
		out.WriteString(ts.ErrorExpr.String())
	}
	
	return out.String()
}