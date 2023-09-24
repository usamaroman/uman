package ast

import (
	"bytes"
	"uman/token"
)

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BlockStatement) String() string {
	var out bytes.Buffer

	for _, statement := range b.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}
func (b *BlockStatement) statementNode() {}
