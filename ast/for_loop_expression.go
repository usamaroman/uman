package ast

import (
	"bytes"

	"uman/token"
)

type ForLoopExpression struct {
	Token     token.Token // token.FOR
	Condition Expression
	Statement *BlockStatement
}

func (f *ForLoopExpression) TokenLiteral() string {
	return f.Token.Literal
}

func (f *ForLoopExpression) String() string {
	var out bytes.Buffer

	out.WriteString("цикл ")
	out.WriteString("(")
	out.WriteString(f.Condition.String())
	out.WriteString(")")
	out.WriteString(f.Statement.String())

	return out.String()
}

func (f *ForLoopExpression) expressionNode() {}
