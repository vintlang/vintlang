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

type TraceStatement struct {
	Token token.Token // the 'trace' token
	Value Expression
}

func (ts *TraceStatement) statementNode()       {}
func (ts *TraceStatement) expressionNode()      {}
func (ts *TraceStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *TraceStatement) String() string {
	return fmt.Sprintf("trace %s", ts.Value.String())
}

type FatalStatement struct {
	Token token.Token // the 'fatal' token
	Value Expression
}

func (fs *FatalStatement) statementNode()       {}
func (fs *FatalStatement) expressionNode()      {}
func (fs *FatalStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *FatalStatement) String() string {
	return fmt.Sprintf("fatal %s", fs.Value.String())
}

type CriticalStatement struct {
	Token token.Token // the 'critical' token
	Value Expression
}

func (cs *CriticalStatement) statementNode()       {}
func (cs *CriticalStatement) expressionNode()      {}
func (cs *CriticalStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *CriticalStatement) String() string {
	return fmt.Sprintf("critical %s", cs.Value.String())
}

type LogStatement struct {
	Token token.Token // the 'log' token
	Value Expression
}

func (ls *LogStatement) statementNode()       {}
func (ls *LogStatement) expressionNode()      {}
func (ls *LogStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LogStatement) String() string {
	return fmt.Sprintf("log %s", ls.Value.String())
}


