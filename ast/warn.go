package ast

import (
	"fmt"

	"github.com/vintlang/vintlang/token"
)

type WarnStatement struct {
	Token token.Token // the 'warn' token
	Value Expression
}

func (ts *WarnStatement) statementNode()       {}
func (ts *WarnStatement) expressionNode()      {}
func (ts *WarnStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *WarnStatement) String() string {
	return fmt.Sprintf("WARN %s", ts.Value.String())
}