package ast

import (
	"bytes"
	"strings"
	"github.com/vintlang/vintlang/token"
)

// PrefixExpression represents prefix expressions like !x, -y, ++z
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression represents infix expressions like x + y, a == b, c > d
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode()      {}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

// PostfixExpression represents postfix expressions like x++, y--
type PostfixExpression struct {
	Token    token.Token
	Operator string
}

func (pe *PostfixExpression) expressionNode()      {}
func (pe *PostfixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PostfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Token.Literal)
	out.WriteString(pe.Operator)
	out.WriteString(")")
	return out.String()
}

// IfExpression represents if-else expressions
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// CallExpression represents function calls like func(arg1, arg2)
type CallExpression struct {
	Token     token.Token
	Function  Expression // can be Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// IndexExpression represents array/dict indexing like arr[0], dict["key"]
type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

// SliceExpression represents array slicing like arr[1:3]
type SliceExpression struct {
	Token token.Token
	Left  Expression
	Start Expression
	End   Expression
}

func (se *SliceExpression) expressionNode()      {}
func (se *SliceExpression) TokenLiteral() string { return se.Token.Literal }
func (se *SliceExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(se.Left.String())
	out.WriteString("[")
	if se.Start != nil {
		out.WriteString(se.Start.String())
	}
	out.WriteString(":")
	if se.End != nil {
		out.WriteString(se.End.String())
	}
	out.WriteString("])")

	return out.String()
}

// Assignment expressions
type Assign struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ae *Assign) expressionNode()      {}
func (ae *Assign) TokenLiteral() string { return ae.Token.Literal }
func (ae *Assign) String() string {
	var out bytes.Buffer

	out.WriteString(ae.Name.String())
	out.WriteString(ae.TokenLiteral())
	out.WriteString(ae.Value.String())

	return out.String()
}

type AssignEqual struct {
	Token token.Token
	Left  *Identifier
	Value Expression
}

func (ae *AssignEqual) expressionNode()      {}
func (ae *AssignEqual) TokenLiteral() string { return ae.Token.Literal }
func (ae *AssignEqual) String() string {
	var out bytes.Buffer

	out.WriteString(ae.Left.String())
	out.WriteString(ae.TokenLiteral())
	out.WriteString(ae.Value.String())

	return out.String()
}

type AssignmentExpression struct {
	Token token.Token
	Left  Expression
	Value Expression
}

func (ae *AssignmentExpression) expressionNode()      {}
func (ae *AssignmentExpression) TokenLiteral() string { return ae.Token.Literal }
func (ae *AssignmentExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ae.Left.String())
	out.WriteString(ae.TokenLiteral())
	out.WriteString(ae.Value.String())

	return out.String()
}

// WhileExpression represents while loops
type WhileExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
}

func (we *WhileExpression) expressionNode()      {}
func (we *WhileExpression) TokenLiteral() string { return we.Token.Literal }
func (we *WhileExpression) String() string {
	var out bytes.Buffer

	out.WriteString("while")
	out.WriteString(we.Condition.String())
	out.WriteString(" ")
	out.WriteString(we.Consequence.String())

	return out.String()
}

// For loop expressions
type For struct {
	Token        token.Token
	Identifier   string      // "i"
	StarterName  *Identifier // i = 0
	StarterValue Expression
	Closer       Expression // i++
	Condition    Expression // i < 1
	Block        *BlockStatement
}

type ForIn struct {
	Token    token.Token
	Key      string
	Value    string
	Iterable Expression
	Block    *BlockStatement
}

func (fi *ForIn) expressionNode()      {}
func (fi *ForIn) TokenLiteral() string { return fi.Token.Literal }
func (fi *ForIn) String() string {
	var out bytes.Buffer

	out.WriteString("for ")
	if fi.Key != "" {
		out.WriteString(fi.Key + ", ")
	}
	out.WriteString(fi.Value + " ")
	out.WriteString("in ")
	out.WriteString(fi.Iterable.String() + " {\n")
	out.WriteString("\t" + fi.Block.String())
	out.WriteString("\n}")

	return out.String()
}

// RepeatStatement represents repeat loops
type RepeatStatement struct {
	Token   token.Token // the 'repeat' token
	VarName string      // loop variable name, default: "i"
	Count   Expression
	Block   *BlockStatement
}

func (rs *RepeatStatement) statementNode()       {}
func (rs *RepeatStatement) expressionNode()      {}
func (rs *RepeatStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *RepeatStatement) String() string {
	var out bytes.Buffer
	out.WriteString("repeat ")
	if rs.VarName != "" {
		out.WriteString("(" + rs.VarName + ") ")
	}
	out.WriteString(rs.Count.String())
	out.WriteString(" ")
	out.WriteString(rs.Block.String())
	return out.String()
}

// Switch expressions
type CaseExpression struct {
	Token   token.Token
	Default bool
	Expr    []Expression
	Block   *BlockStatement
}

func (ce *CaseExpression) expressionNode()      {}
func (ce *CaseExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CaseExpression) String() string {
	var out bytes.Buffer

	if ce.Default {
		out.WriteString("default ")
	} else {
		out.WriteString("case ")

		tmp := []string{}
		for _, exp := range ce.Expr {
			tmp = append(tmp, exp.String())
		}
		out.WriteString(strings.Join(tmp, ","))
	}
	out.WriteString(ce.Block.String())
	return out.String()
}

type SwitchExpression struct {
	Token   token.Token
	Value   Expression
	Choices []*CaseExpression
}

func (se *SwitchExpression) expressionNode()      {}
func (se *SwitchExpression) TokenLiteral() string { return se.Token.Literal }
func (se *SwitchExpression) String() string {
	var out bytes.Buffer
	out.WriteString("\nswitch (")
	out.WriteString(se.Value.String())
	out.WriteString(")\n{\n")

	for _, tmp := range se.Choices {
		if tmp != nil {
			out.WriteString(tmp.String())
		}
	}
	out.WriteString("}\n")

	return out.String()
}

// Method and property expressions
type MethodExpression struct {
	Token     token.Token
	Object    Expression
	Method    Expression
	Arguments []Expression
	Defaults  map[string]Expression
}

func (me *MethodExpression) expressionNode()      {}
func (me *MethodExpression) TokenLiteral() string { return me.Token.Literal }
func (me *MethodExpression) String() string {
	var out bytes.Buffer
	out.WriteString(me.Object.String())
	out.WriteString(".")
	out.WriteString(me.Method.String())

	return out.String()
}

type PropertyAssignment struct {
	Token token.Token // the '=' token
	Name  *PropertyExpression
	Value Expression
}

func (pa *PropertyAssignment) expressionNode()      {}
func (pa *PropertyAssignment) TokenLiteral() string { return pa.Token.Literal }
func (pa *PropertyAssignment) String() string       { return "tach" }

type PropertyExpression struct {
	Expression
	Token    token.Token // The . token
	Object   Expression
	Property Expression
}

func (pe *PropertyExpression) expressionNode()      {}
func (pe *PropertyExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PropertyExpression) String() string       { return "Tach two" }

// Import and Package expressions
type Import struct {
	Token       token.Token
	Identifiers map[string]*Identifier
}

func (i *Import) expressionNode()      {}
func (i *Import) TokenLiteral() string { return i.Token.Literal }
func (i *Import) String() string {
	var out bytes.Buffer
	out.WriteString("import ")
	for k := range i.Identifiers {
		out.WriteString(k + " ")
	}
	return out.String()
}

type Package struct {
	Token token.Token
	Name  *Identifier
	Block *BlockStatement
}

func (p *Package) expressionNode()      {}
func (p *Package) TokenLiteral() string { return p.Token.Literal }
func (p *Package) String() string {
	var out bytes.Buffer

	out.WriteString("package " + p.Name.Value + "\n")
	out.WriteString("::\n")
	for _, s := range p.Block.Statements {
		out.WriteString(s.String())
	}
	out.WriteString("\n::")

	return out.String()
}

// Async expressions
type AwaitExpression struct {
	Token token.Token
	Value Expression
}

func (ae *AwaitExpression) expressionNode()      {}
func (ae *AwaitExpression) TokenLiteral() string { return ae.Token.Literal }
func (ae *AwaitExpression) String() string {
	var out bytes.Buffer
	out.WriteString("await ")
	out.WriteString(ae.Value.String())
	return out.String()
}

type ChannelExpression struct {
	Token  token.Token
	Buffer Expression // optional buffer size
}

func (ce *ChannelExpression) expressionNode()      {}
func (ce *ChannelExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *ChannelExpression) String() string {
	var out bytes.Buffer
	out.WriteString("chan")
	if ce.Buffer != nil {
		out.WriteString("(")
		out.WriteString(ce.Buffer.String())
		out.WriteString(")")
	}
	return out.String()
}

type RangeExpression struct {
	Token token.Token // the '..' token
	Start Expression
	End   Expression
}

func (re *RangeExpression) expressionNode()      {}
func (re *RangeExpression) TokenLiteral() string { return re.Token.Literal }
func (re *RangeExpression) String() string {
	var out bytes.Buffer
	out.WriteString(re.Start.String())
	out.WriteString("..")
	out.WriteString(re.End.String())
	return out.String()
}

// Match expressions
type MatchCase struct {
	Token   token.Token
	Pattern Expression  // Dict pattern or "_" for wildcard
	Block   *BlockStatement
}

func (mc *MatchCase) expressionNode()      {}
func (mc *MatchCase) TokenLiteral() string { return mc.Token.Literal }
func (mc *MatchCase) String() string {
	var out bytes.Buffer
	
	out.WriteString(mc.Pattern.String())
	out.WriteString(" => ")
	out.WriteString(mc.Block.String())
	
	return out.String()
}

type MatchExpression struct {
	Token token.Token
	Value Expression
	Cases []*MatchCase
}

func (me *MatchExpression) expressionNode()      {}
func (me *MatchExpression) TokenLiteral() string { return me.Token.Literal }
func (me *MatchExpression) String() string {
	var out bytes.Buffer
	out.WriteString("match ")
	out.WriteString(me.Value.String())
	out.WriteString(" {\n")
	
	for _, c := range me.Cases {
		if c != nil {
			out.WriteString(c.String())
			out.WriteString("\n")
		}
	}
	
	out.WriteString("}")
	return out.String()
}

