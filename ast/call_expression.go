package ast

import (
	"bytes"
	"strings"

	"github.com/usamaroman/uman/token"
)

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (c *CallExpression) TokenLiteral() string {
	return c.Token.Literal
}

func (c *CallExpression) String() string {
	var out bytes.Buffer

	args := make([]string, 0)
	for _, a := range c.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(c.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

func (c *CallExpression) expressionNode() {}
