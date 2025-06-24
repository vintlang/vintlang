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
