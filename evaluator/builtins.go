package evaluator

import (
	"fmt"
	"uman/object"
)

var builtins = map[string]*object.Builtin{
	"длина": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("неверное количество аргументов %d, должен быть 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len([]rune(arg.Value)))}
			default:
				return newError("не подходящий тип данных %s", args[0].Type())
			}
		},
	},

	"вывести": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return &object.String{Value: ""}
		},
	},
}
