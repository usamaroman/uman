package lexer

import (
	"testing"

	"uman/token"
)

func TestNextToken(t *testing.T) {
	input := `:= == != ;,+-*/
() { } 
<=	> < >=
true = "true" 
тест: строка = "тест"
номер: число = 5
5 6
истина
цикл (;)
[1, 2]
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.COLON, ":"},
		{token.ASSIGN, "="},
		{token.EQUALS, "=="},
		{token.NEQ, "!="},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.ELT, "<="},
		{token.GT, ">"},
		{token.LT, "<"},
		{token.EGT, ">="},
		{token.IDENT, "true"},
		{token.ASSIGN, "="},
		{token.STRING_VAL, "true"},
		{token.IDENT, "тест"},
		{token.COLON, ":"},
		{token.STRING, "строка"},
		{token.ASSIGN, "="},
		{token.STRING_VAL, "тест"},
		{token.IDENT, "номер"},
		{token.COLON, ":"},
		{token.INT, "число"},
		{token.ASSIGN, "="},
		{token.INT_VAL, "5"},
		{token.INT_VAL, "5"},
		{token.INT_VAL, "6"},
		{token.TRUE, "истина"},
		{token.FOR, "цикл"},
		{token.LPAREN, "("},
		{token.SEMICOLON, ";"},
		{token.RPAREN, ")"},
		{token.LBRACKET, "["},
		{token.INT_VAL, "1"},
		{token.COMMA, ","},
		{token.INT_VAL, "2"},
		{token.RBRACKET, "]"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
