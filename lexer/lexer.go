package lexer

import (
	"unicode"

	"github.com/usamaroman/uman/token"
)

type Lexer struct {
	input        []rune
	position     int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch           rune
}

func New(input string) *Lexer {
	in := []rune(input)
	l := &Lexer{input: in}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case ':':
		tok = token.New(token.COLON, l.ch)
	case '!':
		if l.peekRune() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:    token.NEQ,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = token.New(token.BANG, l.ch)
		}
	case '=':
		if l.peekRune() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:    token.EQUALS,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = token.New(token.ASSIGN, l.ch)
		}
	case ';':
		tok = token.New(token.SEMICOLON, l.ch)
	case ',':
		tok = token.New(token.COMMA, l.ch)
	case '(':
		tok = token.New(token.LPAREN, l.ch)
	case ')':
		tok = token.New(token.RPAREN, l.ch)
	case '{':
		tok = token.New(token.LBRACE, l.ch)
	case '}':
		tok = token.New(token.RBRACE, l.ch)
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case '-':
		tok = token.New(token.MINUS, l.ch)
	case '*':
		tok = token.New(token.ASTERISK, l.ch)
	case '/':
		tok = token.New(token.SLASH, l.ch)
	case '>':
		if l.peekRune() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EGT, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.New(token.GT, l.ch)
		}
	case '<':
		if l.peekRune() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.ELT, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.New(token.LT, l.ch)
		}
	case '"':
		tok.Type = token.STRING_VAL
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '[':
		tok = token.New(token.LBRACKET, '[')
	case ']':
		tok = token.New(token.RBRACKET, ']')
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readDigit()
			tok.Type = token.INT_VAL
			return tok
		} else {
			tok = token.New(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return string(l.input[position:l.position])
}

func (l *Lexer) readDigit() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return string(l.input[position:l.position])
}

func isLetter(ch rune) bool {
	return unicode.Is(unicode.Cyrillic, ch) || unicode.Is(unicode.Latin, ch)
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return string(l.input[position:l.position])
}

func (l *Lexer) peekRune() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return byte(l.input[l.readPosition])
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}
