package ast

import "uman/token"

// ReturnStatement returns variable from function
// implements Statement interface
type ReturnStatement struct {
	Token token.Token // token.RETURN
	Value Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
func (rs *ReturnStatement) statementNode() {}
