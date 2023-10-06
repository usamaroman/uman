package ast

import "github.com/usamaroman/uman/token"

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (b *BooleanLiteral) expressionNode() {}
func (b *BooleanLiteral) TokenLiteral() string {
	return b.Token.Literal
}
func (b *BooleanLiteral) String() string {
	return b.Token.Literal
}
