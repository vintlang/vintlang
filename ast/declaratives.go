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

type DeferStatement struct {
	Token token.Token // the 'defer' token
	Call  Expression
}

func (ds *DeferStatement) statementNode()       {}
func (ds *DeferStatement) expressionNode()      {}
func (ds *DeferStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DeferStatement) String() string {
	return fmt.Sprintf("defer %s", ds.Call.String())
}

type InfoStatement struct {
	Token token.Token // the 'info' token
	Value Expression
}

func (is *InfoStatement) statementNode()       {}
func (is *InfoStatement) expressionNode()      {}
func (is *InfoStatement) TokenLiteral() string { return is.Token.Literal }
func (is *InfoStatement) String() string {
	return fmt.Sprintf("info %s", is.Value.String())
}

type DebugStatement struct {
	Token token.Token // the 'debug' token
	Value Expression
}

func (ds *DebugStatement) statementNode()       {}
func (ds *DebugStatement) expressionNode()      {}
func (ds *DebugStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DebugStatement) String() string {
	return fmt.Sprintf("debug %s", ds.Value.String())
}

type NoteStatement struct {
	Token token.Token // the 'note' token
	Value Expression
}

func (ns *NoteStatement) statementNode()       {}
func (ns *NoteStatement) expressionNode()      {}
func (ns *NoteStatement) TokenLiteral() string { return ns.Token.Literal }
func (ns *NoteStatement) String() string {
	return fmt.Sprintf("note %s", ns.Value.String())
}

type SuccessStatement struct {
	Token token.Token // the 'success' token
	Value Expression
}

func (ss *SuccessStatement) statementNode()       {}
func (ss *SuccessStatement) expressionNode()      {}
func (ss *SuccessStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SuccessStatement) String() string {
	return fmt.Sprintf("success %s", ss.Value.String())
}
