package evaluator

import (
	"fmt"

	"uman/ast"
	"uman/object"
	"uman/token"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

var dataTypes = map[token.TokenType]object.ObjectType{
	token.INT:      object.IntegerObj,
	token.STRING:   object.StringObj,
	token.BOOL:     object.BooleanObj,
	token.FUNCTION: object.FunctionObj,
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ErrorObj
	}
	return false
}

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right, env, node.Left.TokenLiteral())

	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.ForLoopExpression:
		return evalForLoopExpression(node, env)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.ReturnStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.VariableStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}

		if !checkDataType(node, val) {
			return newError("неверная инициализация типа данных %s %s", node.DataType, val.Type())
		}

		if obj, ok := env.Get(node.Ident.Value); ok {
			return newError("переменная %s уже существует = %s", node.Ident.Value, obj.Inspect())
		}

		env.Set(node.Ident.Value, val)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.FunctionLiteral:
		args := node.Arguments
		body := node.Body
		return &object.Function{
			Arguments: args,
			Body:      body,
			Env:       env,
		}
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)

	// expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.BooleanLiteral:
		return nativeBoolToBooleanObj(node.Value)
	default:
		return nil
	}
	return nil
}

func checkDataType(node *ast.VariableStatement, obj object.Object) bool {
	val, ok := dataTypes[node.DataType]
	if !ok {
		return false
	}
	return val == obj.Type()
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newError("нет функции %s", fn.Type())
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Arguments {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("нет переменной: %s", node.Value)
}

func evalIfExpression(node *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(node.Condition, env)
	if isError(condition) {
		return condition
	}
	if isTrue(condition) {
		return Eval(node.Consequence, env)
	} else if node.Alternative != nil {
		return Eval(node.Alternative, env)
	} else {
		return NULL
	}
}

func evalForLoopExpression(node *ast.ForLoopExpression, env *object.Environment) object.Object {
	condition := Eval(node.Condition, env)
	if isError(condition) {
		return condition
	}

	if condition.Type() != object.BooleanObj {
		return newError("условие должно быть булевого типа, получено %s", condition.Type())
	}

	for isTrue(Eval(node.Condition, env)) {
		Eval(node.Statement, env)
	}

	return NULL
}

func isTrue(obj object.Object) bool {
	switch obj {
	case TRUE:
		return true
	case FALSE:
		return false
	case NULL:
		return false
	default:
		return false
	}
}

func evalInfixExpression(operator string, left object.Object, right object.Object, env *object.Environment, ident string) object.Object {
	switch {
	case operator == "=":
		return env.Set(ident, right)
	case left.Type() == object.IntegerObj && right.Type() == object.IntegerObj:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.StringObj && right.Type() == object.StringObj:
		return evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObj(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObj(left != right)
	case left.Type() != right.Type():
		return newError("разные типы: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("неизвестный оператор: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	if operator != "+" {
		return newError("неизвестный оператор: %s %s %s", left.Type(), operator, right.Type())
	}

	leftValue := left.(*object.String).Value
	rightValue := right.(*object.String).Value

	return &object.String{Value: leftValue + rightValue}
}

func evalIntegerInfixExpression(operator string, left object.Object, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case ">":
		return nativeBoolToBooleanObj(leftVal > rightVal)
	case "<":
		return nativeBoolToBooleanObj(leftVal < rightVal)
	case ">=":
		return nativeBoolToBooleanObj(leftVal >= rightVal)
	case "<=":
		return nativeBoolToBooleanObj(leftVal <= rightVal)
	case "==":
		return nativeBoolToBooleanObj(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObj(leftVal != rightVal)
	default:
		return newError("неизвестный оператор: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusOperatorExpression(right)
	default:
		return newError("неизвестный оператор: %s%s", operator, right.Type())
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.IntegerObj {
		return newError("неизвестный оператор: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func nativeBoolToBooleanObj(value bool) *object.Boolean {
	if value {
		return TRUE
	}
	return FALSE
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.ReturnValueObj || rt == object.ErrorObj {
				return result
			}
		}
	}

	return result
}
