package parser

import (
	"testing"
	"uman/ast"
)

func TestReturnStatement(t *testing.T) {
	input := `
вернуть 5;
вернуть "zoo";
вернуть fasfdaf;
`

	p := New(input)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("wrong len, got length=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "вернуть" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}

}
