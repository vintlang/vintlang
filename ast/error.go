package ast

import (
	"fmt"

	"github.com/vintlang/vintlang/token"
)

type ErrorStatement struct {
	Token token.Token // the 'error' token
	Value Expression
}

func (es *ErrorStatement) statementNode()       {}
func (es *ErrorStatement) expressionNode()      {}
func (es *ErrorStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ErrorStatement) String() string {
	return fmt.Sprintf("error %s", es.Value.String())
}
