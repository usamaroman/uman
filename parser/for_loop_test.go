package parser

import (
	"testing"

	"github.com/usamaroman/uman/ast"
)

func TestForLoopExpression(t *testing.T) {
	input := `
цикл (x < y) {
    вывести(y);
    создать с: число = x + y;
}
`

	p := New(input)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.ForLoopExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.ForLoopExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.Statement.Statements) != 2 {
		t.Errorf("consequence is not 2 statements. got=%d\n", len(exp.Statement.Statements))
	}

	stm, ok := exp.Statement.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T",
			exp.Statement.Statements[0])
	}

	if !testStmt(t, stm.Expression, "вывести", "y") {
		return
	}
}

func TestForLoopExpressionEvaluator(t *testing.T) {
	input := `
создать i: число = 0;
цикл (i >= 0) {
    вывести(123);
    i = i - 1;
}
`

	p := New(input)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 2 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[1].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.ForLoopExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.ForLoopExpression. got=%T", stmt.Expression)
	}

	if len(exp.Statement.Statements) != 2 {
		t.Errorf("consequence is not 2 statements. got=%d\n", len(exp.Statement.Statements))
	}
}

func testStmt(t *testing.T, exp ast.Expression, value string, arg string) bool {
	fn, ok := exp.(*ast.CallExpression)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if fn.Function.TokenLiteral() != value {
		t.Errorf("function name not %s. got=%s", value, fn.Function.TokenLiteral())
		return false
	}

	if fn.Arguments[0].TokenLiteral() != arg {
		t.Errorf("argument not %s. got=%s", arg, fn.Arguments[0].TokenLiteral())
		return false
	}

	return true
}
