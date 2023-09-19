package parser

import (
	"fmt"
	"log"

	"uman/ast"
	"uman/lexer"
	"uman/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token

	errors []string
}

func New(input string) *Parser {
	l := lexer.New(input)

	p := &Parser{
		l:      l,
		errors: make([]string, 0),
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.IDENT:
		return p.parseVariableStatement()
	default:
		return nil
	}
}

func (p *Parser) parseVariableStatement() ast.Statement {
	stmt := &ast.VariableStatement{}

	stmt.Ident = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	if !p.expectPeek(token.COLON) {
		return nil
	}

	stmt.Token = p.currToken

	if p.getDataType() != token.STRING && p.getDataType() != token.INT {
		p.addError("missing data type")
		return nil
	} else {
		p.nextToken()
	}

	stmt.DataType = p.currToken.Type

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	log.Println(stmt.Ident.Value, stmt.Token, stmt.DataType, "=", stmt.Value)
	return stmt
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) getDataType() token.TokenType {
	return p.peekToken.Type
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) addError(msg string) {
	p.errors = append(p.errors, msg)
}
