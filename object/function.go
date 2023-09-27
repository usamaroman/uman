package object

import (
	"bytes"
	"strings"

	"uman/ast"
)

type Function struct {
	Arguments []*ast.Identifier
	Body      *ast.BlockStatement
	Env       *Environment
}

func (f *Function) Type() ObjectType {
	return FunctionObj
}

func (f *Function) Inspect() string {
	var out bytes.Buffer

	args := make([]string, 0)
	for _, arg := range f.Arguments {
		args = append(args, arg.String())
	}

	out.WriteString("функция")
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
