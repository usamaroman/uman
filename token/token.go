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

var keywords = map[string]TokenType{
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
	if tok, ok := keywords[literal]; ok {
		return tok
	}
	return IDENT
}
