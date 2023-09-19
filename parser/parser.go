package parser

import (
	"log"

	"uman/ast"
	"uman/lexer"
	"uman/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func New(input string) *Parser {
	l := lexer.New(input)

	p := &Parser{
		l: l,
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
		//case token.COLON:
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

	if !p.expectPeek(token.STRING) && !p.expectPeek(token.INT) {
		return nil
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	log.Println(stmt.Token, stmt.Ident, stmt.Value)
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
		return false
	}
}
