package ast

import "github.com/vintlang/vintlang/token"

// Type system foundation for VintLang
type Type interface {
	Node
	typeNode()
}

// Basic type identifiers
type BasicType struct {
	Token token.Token // int, string, bool, float, etc.
	Name  string
}

func (bt *BasicType) expressionNode() {}
func (bt *BasicType) typeNode()       {}
func (bt *BasicType) TokenLiteral() string { return bt.Token.Literal }
func (bt *BasicType) String() string       { return bt.Name }

// Array type: []int, []string, etc.
type ArrayType struct {
	Token       token.Token // the '[' token
	ElementType Type
}

func (at *ArrayType) expressionNode() {}
func (at *ArrayType) typeNode()       {}
func (at *ArrayType) TokenLiteral() string { return at.Token.Literal }
func (at *ArrayType) String() string {
	return "[]" + at.ElementType.String()
}

// Function type: func(int, string) bool
type FunctionType struct {
	Token      token.Token // the 'func' token
	Parameters []Type
	ReturnType Type
}

func (ft *FunctionType) expressionNode() {}
func (ft *FunctionType) typeNode()       {}
func (ft *FunctionType) TokenLiteral() string { return ft.Token.Literal }
func (ft *FunctionType) String() string {
	params := ""
	for i, p := range ft.Parameters {
		if i > 0 {
			params += ", "
		}
		params += p.String()
	}
	return "func(" + params + ") " + ft.ReturnType.String()
}

// Optional type: int?, string?, etc.
type OptionalType struct {
	Token    token.Token // the '?' token
	BaseType Type
}

func (ot *OptionalType) expressionNode() {}
func (ot *OptionalType) typeNode()       {}
func (ot *OptionalType) TokenLiteral() string { return ot.Token.Literal }
func (ot *OptionalType) String() string {
	return ot.BaseType.String() + "?"
}

// Dict type: {string: int}
type DictType struct {
	Token     token.Token // the '{' token
	KeyType   Type
	ValueType Type
}

func (dt *DictType) expressionNode() {}
func (dt *DictType) typeNode()       {}
func (dt *DictType) TokenLiteral() string { return dt.Token.Literal }
func (dt *DictType) String() string {
	return "{" + dt.KeyType.String() + ": " + dt.ValueType.String() + "}"
}

// Union type: int | string | bool
type UnionType struct {
	Token token.Token // the first type token
	Types []Type
}

func (ut *UnionType) expressionNode() {}
func (ut *UnionType) typeNode()       {}
func (ut *UnionType) TokenLiteral() string { return ut.Token.Literal }
func (ut *UnionType) String() string {
	result := ""
	for i, t := range ut.Types {
		if i > 0 {
			result += " | "
		}
		result += t.String()
	}
	return result
}

// Typed parameter for functions
type TypedParameter struct {
	Token      token.Token
	Identifier *Identifier
	Type       Type
	Default    Expression // optional default value
}

func (tp *TypedParameter) expressionNode() {}
func (tp *TypedParameter) TokenLiteral() string { return tp.Token.Literal }
func (tp *TypedParameter) String() string {
	result := tp.Identifier.String() + ": " + tp.Type.String()
	if tp.Default != nil {
		result += " = " + tp.Default.String()
	}
	return result
}

// Type annotation for variables
type TypeAnnotation struct {
	Token token.Token // the ':' token
	Type  Type
}

func (ta *TypeAnnotation) expressionNode() {}
func (ta *TypeAnnotation) TokenLiteral() string { return ta.Token.Literal }
func (ta *TypeAnnotation) String() string {
	return ": " + ta.Type.String()
}

// Enhanced let statement with type annotation
type TypedLetStatement struct {
	Token          token.Token // the 'let' token
	Name           *Identifier
	TypeAnnotation *TypeAnnotation // optional: let x: int = 5
	Value          Expression
}

func (tls *TypedLetStatement) statementNode() {}
func (tls *TypedLetStatement) TokenLiteral() string { return tls.Token.Literal }
func (tls *TypedLetStatement) String() string {
	result := tls.TokenLiteral() + " " + tls.Name.String()
	if tls.TypeAnnotation != nil {
		result += tls.TypeAnnotation.String()
	}
	result += " = " + tls.Value.String() + ";"
	return result
}

// Enhanced function literal with typed parameters and return type
type TypedFunctionLiteral struct {
	Token      token.Token // the 'func' token
	Parameters []*TypedParameter
	ReturnType Type // optional return type annotation
	Body       *BlockStatement
}

func (tfl *TypedFunctionLiteral) expressionNode() {}
func (tfl *TypedFunctionLiteral) TokenLiteral() string { return tfl.Token.Literal }
func (tfl *TypedFunctionLiteral) String() string {
	params := ""
	for i, p := range tfl.Parameters {
		if i > 0 {
			params += ", "
		}
		params += p.String()
	}
	result := tfl.TokenLiteral() + "(" + params + ")"
	if tfl.ReturnType != nil {
		result += ": " + tfl.ReturnType.String()
	}
	result += " " + tfl.Body.String()
	return result
}

// Type casting expression: x as int
type TypeCastExpression struct {
	Token      token.Token // the 'as' token
	Expression Expression
	TargetType Type
}

func (tce *TypeCastExpression) expressionNode() {}
func (tce *TypeCastExpression) TokenLiteral() string { return tce.Token.Literal }
func (tce *TypeCastExpression) String() string {
	return "(" + tce.Expression.String() + " as " + tce.TargetType.String() + ")"
}

// Type checking expression: x is int
type TypeCheckExpression struct {
	Token      token.Token // the 'is' token
	Expression Expression
	CheckType  Type
}

func (tce *TypeCheckExpression) expressionNode() {}
func (tce *TypeCheckExpression) TokenLiteral() string { return tce.Token.Literal }
func (tce *TypeCheckExpression) String() string {
	return "(" + tce.Expression.String() + " is " + tce.CheckType.String() + ")"
}