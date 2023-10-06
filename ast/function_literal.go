package ast

import (
	"bytes"
	"strings"

	"github.com/usamaroman/uman/token"
)

type FunctionLiteral struct {
	Token     token.Token
	Arguments []*Identifier
	Body      *BlockStatement
}

func (f *FunctionLiteral) expressionNode() {}
func (f *FunctionLiteral) TokenLiteral() string {
	return f.Token.Literal
}

func (f *FunctionLiteral) String() string {
	var out bytes.Buffer

	args := make([]string, 0)
	for _, arg := range f.Arguments {
		args = append(args, arg.String())
	}

	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())

	return out.String()
}
