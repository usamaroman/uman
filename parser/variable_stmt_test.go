package parser

import (
	"testing"
	"uman/ast"
)

func TestVariableStatements(t *testing.T) {
	// верные выражения 1 3 и последнее, остальные неверно проинициализированы
	input := `
текст: строка = "тест"; 
строка: строка = "wasd";
number: число = 5;
номер: число ;
: число ;
номер строка = 5;
номер: число = 5;
`

	p := New(input)
	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("wrong len, got length=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{expectedIdentifier: "текст"},
		{expectedIdentifier: "number"},
		{expectedIdentifier: "номер"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testVariableStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func testVariableStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != ":" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	varStmt, ok := stmt.(*ast.VariableStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", stmt)
		return false
	}

	if varStmt.Ident.Value != name {
		t.Errorf("varStmt.Ident.Value not '%s'. got=%s", name, varStmt.Ident.Value)
		return false
	}

	if varStmt.Ident.TokenLiteral() != name {
		t.Errorf("varStmt.Ident not '%s'. got=%s", name, varStmt.Ident)
		return false
	}

	return true
}