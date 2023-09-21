package ast

import (
	"bytes"
	"uman/token"
)

// ReturnStatement returns variable from function
// implements Statement interface
type ReturnStatement struct {
	Token token.Token // token.RETURN
	Value Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}

	out.WriteString(";")
	return out.String()
}
func (rs *ReturnStatement) statementNode() {}
