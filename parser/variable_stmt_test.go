package parser

import (
	"testing"

	"uman/ast"
	"uman/token"
)

func TestVariableStatements(t *testing.T) {
	tests := []struct {
		input            string
		expectedIdent    string
		expectedValue    any
		expectedDataType token.TokenType
	}{
		{"создать вернуть: число = 5;", "вернуть", 5, token.INT},
		{"создать вернуть: число = a;", "вернуть", "a", token.INT},
		{`создать текст: строка = asd;`, "текст", "asd", token.STRING},
	}

	for _, tt := range tests {
		p := New(tt.input)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if program == nil {
			t.Fatalf("returned nil")
		}

		if len(program.Statements) != 1 {
			t.Fatalf("wrong len, got length=%d", len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testVariableStatement(t, stmt, tt.expectedIdent) {
			return
		}

		val := stmt.(*ast.VariableStatement).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}

func testVariableStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "создать" {
		t.Errorf("s.TokenLiteral not 'создать'. got=%q", stmt.TokenLiteral())
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

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	return true
}

func testStringVar(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.StringLiteral)
	if !ok {
		t.Errorf("exp not *ast.StringLiteral. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	return true
}
