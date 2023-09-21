package ast

import (
	"bytes"
	"uman/token"
)

// VariableStatement variable creates using :
// implements Statement interface
type VariableStatement struct {
	Token    token.Token // token.LET
	Ident    *Identifier
	DataType token.TokenType
	Value    Expression
}

func (vs *VariableStatement) TokenLiteral() string {
	return vs.Token.Literal
}
func (vs *VariableStatement) String() string {
	var out bytes.Buffer

	out.WriteString(vs.TokenLiteral())
	out.WriteString(": ")
	out.WriteString(getDataTypeFromKeywords(vs.DataType))
	out.WriteString(" = ")

	if vs.Value != nil {
		out.WriteString(vs.Value.String())
	}

	out.WriteString(";")
	return out.String()
}
func (vs *VariableStatement) statementNode() {}

func getDataTypeFromKeywords(tokenType token.TokenType) string {
	for k, v := range token.Keywords {
		if v == tokenType {
			return k
		}
	}
	return ""
}
