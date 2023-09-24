package parser

import (
	"fmt"
	"log"
	"strconv"

	"uman/ast"
	"uman/lexer"
	"uman/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X !X
	CALL        // func(X)
)

var precedences = map[token.TokenType]int{
	token.EQUALS:   EQUALS,
	token.NEQ:      EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.EGT:      LESSGREATER,
	token.ELT:      LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

func (p *Parser) peekPrecedences() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) currPrecedences() int {
	if p, ok := precedences[p.currToken.Type]; ok {
		return p
	}

	return LOWEST
}

type Parser struct {
	l      *lexer.Lexer
	errors []string

	currToken token.Token
	peekToken token.Token

	prefixParserFns map[token.TokenType]prefixParseFn // !test
	infixParserFns  map[token.TokenType]infixParseFn  // test + test
}

func New(input string) *Parser {
	l := lexer.New(input)

	p := &Parser{
		l:               l,
		errors:          make([]string, 0),
		prefixParserFns: make(map[token.TokenType]prefixParseFn),
		infixParserFns:  make(map[token.TokenType]infixParseFn),
	}

	p.registerPrefixFn(token.IDENT, p.parseIdent)
	p.registerPrefixFn(token.INT_VAL, p.parseIntegerLiteral)
	p.registerPrefixFn(token.STRING_VAL, p.parseStringLiteral)
	p.registerPrefixFn(token.TRUE, p.parseBoolean)
	p.registerPrefixFn(token.FALSE, p.parseBoolean)
	p.registerPrefixFn(token.BANG, p.parsePrefixExpression)
	p.registerPrefixFn(token.MINUS, p.parsePrefixExpression)
	p.registerPrefixFn(token.LPAREN, p.parseGroupedExpression)

	p.registerInfixFn(token.PLUS, p.parseInfixExpression)
	p.registerInfixFn(token.MINUS, p.parseInfixExpression)
	p.registerInfixFn(token.ASTERISK, p.parseInfixExpression)
	p.registerInfixFn(token.SLASH, p.parseInfixExpression)
	p.registerInfixFn(token.LT, p.parseInfixExpression)
	p.registerInfixFn(token.ELT, p.parseInfixExpression)
	p.registerInfixFn(token.GT, p.parseInfixExpression)
	p.registerInfixFn(token.EGT, p.parseInfixExpression)
	p.registerInfixFn(token.EQUALS, p.parseInfixExpression)
	p.registerInfixFn(token.NEQ, p.parseInfixExpression)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) registerPrefixFn(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParserFns[tokenType] = fn
}

func (p *Parser) registerInfixFn(tokenType token.TokenType, fn infixParseFn) {
	p.infixParserFns[tokenType] = fn
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
	case token.LET:
		return p.parseVariableStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseVariableStatement() *ast.VariableStatement {
	stmt := &ast.VariableStatement{
		Token: p.currToken,
	}

	p.nextToken()

	stmt.Ident = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	if !p.expectPeek(token.COLON) {
		return nil
	}

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
	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	log.Println(stmt.TokenLiteral(), stmt.Ident.Value, stmt.DataType, "=", stmt.Value)
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{
		Token: p.currToken,
	}
	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

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

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token: p.currToken,
	}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParserFns[p.currToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currToken.Type)
		return nil
	}
	leftExp := prefix()
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedences() {
		infix := p.infixParserFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) parseIdent() ast.Expression {
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{
		Token: p.currToken,
	}

	i, err := strconv.ParseInt(p.currToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = i

	return lit
}

func (p *Parser) parseStringLiteral() ast.Expression {
	expression := &ast.StringLiteral{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	return expression
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	stmt := &ast.PrefixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
	}

	p.nextToken()

	stmt.Right = p.parseExpression(PREFIX)

	return stmt
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	stmt := &ast.InfixExpression{
		Token:    p.currToken,
		Left:     left,
		Operator: p.currToken.Literal,
	}

	precedence := p.currPrecedences()
	p.nextToken()
	stmt.Right = p.parseExpression(precedence)

	return stmt
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseBoolean() ast.Expression {
	exp := &ast.BooleanLiteral{
		Token: p.currToken,
		Value: p.currTokenIs(token.TRUE),
	}
	return exp
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}
