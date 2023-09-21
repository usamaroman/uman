package token

type TokenType string

const (
	IDENT   = "IDENT"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	ASSIGN   = "="
	EQUALS   = "=="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	GT       = ">"
	EGT      = ">="
	LT       = "<"
	ELT      = "<="
	BANG     = "!"
	NEQ      = "!="

	LET       = "LET"
	COLON     = ":"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	STRING_VAL = "STRING_VAL"
	INT_VAL    = "INT_VAL"

	// Keywords
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	INT      = "INT"
	STRING   = "STRING"
)

var Keywords = map[string]TokenType{
	"создать": LET,
	"функция": FUNCTION,
	"истина":  TRUE,
	"ложь":    FALSE,
	"если":    IF,
	"иначе":   ELSE,
	"вернуть": RETURN,
	"число":   INT,
	"строка":  STRING,
}

type Token struct {
	Type    TokenType
	Literal string
}

func New(tokenType TokenType, literal rune) Token {
	return Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}

func LookupIdent(literal string) TokenType {
	if tok, ok := Keywords[literal]; ok {
		return tok
	}
	return IDENT
}
