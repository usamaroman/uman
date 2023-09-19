package ast

import "uman/token"

// VariableStatement variable creates using :
// implements Statement interface
type VariableStatement struct {
	Token token.Token // token.Colon
	Ident *Identifier
	Value Expression
}

func (cs *VariableStatement) TokenLiteral() string {
	return cs.Token.Literal
}
func (cs *VariableStatement) statementNode() {}
