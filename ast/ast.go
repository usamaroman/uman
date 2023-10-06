package ast

import (
	"bytes"

	"github.com/usamaroman/uman/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
func (p *Program) String() string {
	var out bytes.Buffer

	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}

	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}
func (id *Identifier) String() string {
	return id.Value
}
func (id *Identifier) expressionNode() {}
