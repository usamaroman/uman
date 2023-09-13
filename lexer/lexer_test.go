package lexer

import (
	"testing"

	"uman/token"
)

func TestNextToken(t *testing.T) {
	input := `:=;,+-*/  true истина`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.COLON, ":"},
		{token.ASSIGN, "="},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLACH, "/"},
		{token.IDENT, "true"},
		{token.TRUE, "истина"},
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
