package parser

import (
	"testing"
)

func TestReturnStatement(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue string
	}{
		{"вернуть 5;", "5"},
		{"вернуть 5;", "zoo"},
		{"вернуть 5;", "fasfdaf"},
	}

	for _, tt := range tests {
		p := New(tt.input)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if program == nil {
			t.Fatalf("returned nil")
		}

	}

	//
	//if len(program.Statements) != 3 {
	//	t.Fatalf("wrong len, got length=%d", len(program.Statements))
	//}
	//
	//for _, stmt := range program.Statements {
	//	returnStmt, ok := stmt.(*ast.ReturnStatement)
	//	if !ok {
	//		t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
	//		continue
	//	}
	//
	//	if returnStmt.TokenLiteral() != "вернуть" {
	//		t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
	//			returnStmt.TokenLiteral())
	//	}
	//}

}
