package evaluator

import (
	"testing"
	"uman/object"
	"uman/parser"
)

func testEval(input string) object.Object {
	p := parser.New(input)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return Eval(program, env)
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

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`
	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
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
						вернуть истина + ложь;
					  }
					  вернуть 1;
					}`,
			"неизвестный оператор: BOOLEAN + BOOLEAN",
		},
		{
			"тест",
			"нет переменной: тест",
		},
		{
			`"Hello" - "World"`,
			"неизвестный оператор: STRING - STRING",
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
		{"создать a: число = 5; a;", 5},
		{"создать a: число = 5 * 5; a;", 25},
		{"создать a: число = 5; создать b: число = a; b;", 5},
		{"создать a: число = 5; создать b: число = a; создать c: число = a + b + 5; c;", 15},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "функция(x) { x + 2; };"
	evaluated := testEval(input)

	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Arguments) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Arguments)
	}

	if fn.Arguments[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Arguments[0])
	}

	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"создать identity: число = функция(x) { x; }; identity(5);", 5},
		{"создать identity: число = функция(x) { вернуть x; }; identity(5);", 5},
		{"создать double: число = функция(x) { x * 2; }; double(5);", 10},
		{"создать add: число = функция(x, y) { x + y; }; add(5, 5);", 10},
		{"создать add: число = функция(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"функция(x) { x; }(5)", 5},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
создать фиб: число = функция(x) {
	если ( x == 0 ) { 
		вернуть 0;
	}

	если ( x == 1 ) { 
		вернуть 1;
	}
	
	вернуть фиб(x - 2) + фиб(x - 1);
};

создать рез: число = фиб(6);
рез;
`
	testIntegerObject(t, testEval(input), 8)
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`длина("")`, 0},
		{`длина("four")`, 4},
		{`длина("qwerty")`, 6},
		{`длина("hello world")`, 11},
		{`длина("йй")`, 2},
		{`длина(1)`, "не подходящий тип данных INTEGER"},
		{`длина("one", "two")`, "неверное количество аргументов 2, должен быть 1"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)",
					evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q",
					expected, errObj.Message)
			}
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"
	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}
	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d",
			len(result.Elements))
	}
	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"let i = 0; [1][i];",
			1,
		},
		{
			"[1, 2, 3][1 + 1];",
			3,
		},
		{
			"let myArray = [1, 2, 3]; myArray[2];",
			3,
		},
		{
			"let myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];",
			6,
		},
		{
			"let myArray = [1, 2, 3]; let i = myArray[0]; myArray[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			nil,
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}
