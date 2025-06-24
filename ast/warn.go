package ast

import (
	"fmt"

	"github.com/vintlang/vintlang/token"
)

type WarnStatement struct {
	Token token.Token // the 'warn' token
	Value Expression
}

func (ws *WarnStatement) statementNode()       {}
func (ws *WarnStatement) expressionNode()      {}
func (ws *WarnStatement) TokenLiteral() string { return ws.Token.Literal }
func (ws *WarnStatement) String() string {
	return fmt.Sprintf("warn %s", ws.Value.String())
}
