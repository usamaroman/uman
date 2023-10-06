package ast

import "github.com/usamaroman/uman/token"

type ExpressionStatement struct {
	Token      token.Token // first token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
