package ast

import (
	"fmt"

	"github.com/vintlang/vintlang/token"
)

type TodoStatement struct {
	Token token.Token // the 'todo' token
	Value Expression
}

func (ts *TodoStatement) statementNode()       {}
func (ts *TodoStatement) expressionNode()      {}
func (ts *TodoStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *TodoStatement) String() string {
	return fmt.Sprintf("todo %s", ts.Value.String())
}

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
