package evaluator

import (
	"testing"

	"uman/object"
	"uman/parser"
)

func testEval(input string) object.Object {
	p := parser.New(input)
	program := p.ParseProgram()
	return Eval(program)
}

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
		return false
	}
	return true
}

func TestEvalStringExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"test"`, "test"},
		{`"тест"`, "тест"},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObject(t, evaluated, tt.expected)
	}
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%s, want=%s",
			result.Value, expected)
		return false
	}
	return true
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"истина", true},
		{"ложь", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"1 != -1", true},
		{"1 <= 2", true},
		{"1 >= 0", true},
		{"1 >= 2", false},
		{"0 >= 0", true},
		{"(1 < 2) == истина", true},
		{"(истина != ложь) == ложь", false},
		{"истина == ложь", false},
		{"истина != ложь", true},
		{"ложь != истина", true},
		{"(1 < 2) == истина", true},
		{"(1 < 2) == ложь", false},
		{"(1 > 2) == истина", false},
		{"(1 > 2) == ложь", true},
		{"(5 > 5 == истина) != ложь", false},
		{"(1 - 1) == 0", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t",
			result.Value, expected)
		return false
	}
	return true
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!истина", false},
		{"!ложь", true},
		{"!5", false},
		{"!!истина", true},
		{"!!ложь", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"если ( истина ) { 10 }", 10},
		{"если ( ложь ) { 10 }", "ничего"},
		{"если (1) { 10 }", 10},
		{"если (1 < 2) { 10 }", 10},
		{"если (1 > 2) { 10 }", "ничего"},
		{"если (1 > 2) { 10 } иначе { 20 }", 20},
		{"если (1 < 2) { 10 } иначе { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		num, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(num))
		} else {
			testNullObject(t, evaluated)
		}
	}

}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"вернуть 10;", 10},
		{"вернуть 10; 9;", 10},
		{"вернуть 2 * 5; 9;", 10},
		{"9; вернуть 2 * 5; 9;", 10},
		{
			`
			if (10 > 1) {
				if (10 > 1) {
				    вернуть 10;
				}
				вернуть 1;
			}
			`, 10,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + истина;",
			"разные типы: INTEGER + BOOLEAN",
		},
		{
			"5 + истина; 5;",
			"разные типы: INTEGER + BOOLEAN",
		},
		{
			"-истина",
			"неизвестный оператор: -BOOLEAN",
		},
		{
			"истина + ложь;",
			"неизвестный оператор: BOOLEAN + BOOLEAN",
		},
		{
			"5; истина + ложь; 5",
			"неизвестный оператор: BOOLEAN + BOOLEAN",
		},
		{
			"если (10 > 1) { истина + ложь; }",
			"неизвестный оператор: BOOLEAN + BOOLEAN",
		},
		{
			`если (10 > 1) {
					  если (10 > 1) {
						return истина + ложь;
					  }
					  вернуть 1;
					}`,
			"неизвестный оператор: BOOLEAN + BOOLEAN",
		},
		{
			"тест",
			"нет переменной: тест",
		},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)

		if !ok {
			t.Errorf("%d, no error object returned. got=%T(%+v)", i,
				evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("%d, wrong error message. expected=%q, got=%q", i,
				tt.expectedMessage, errObj.Message)
		}

	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}
