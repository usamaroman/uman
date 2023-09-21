package ast

import (
	"log"
	"testing"
	"uman/token"
)

func TestAst(t *testing.T) {
	p := &Program{
		Statements: []Statement{
			&VariableStatement{
				Token: token.Token{
					Type:    token.IDENT,
					Literal: "test",
				},
				Ident: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "test",
					},
					Value: "test",
				},
				DataType: token.STRING,
				Value: &Identifier{
					Token: token.Token{
						Type:    token.STRING_VAL,
						Literal: "тест",
					},
					Value: `"тест"`,
				},
			},
		},
	}

	log.Println(p.String())
	if p.String() != `test: строка = "тест";` {
		t.Fatalf("got %q", p.String())
	}

}
