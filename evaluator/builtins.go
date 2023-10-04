package evaluator

import (
	"fmt"

	"uman/object"
)

var builtins = map[string]*object.Builtin{
	"длина": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("неверное количество аргументов получено %d, надо 1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("нельзя передавать в длина(), получено %s",
					args[0].Type())
			}
		},
	},

	"вывести": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Print(arg.Inspect())
				fmt.Print(" ")
			}
			fmt.Println()

			return NULL
		},
	},

	"первый": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("неверное количество аргументов получено %d, надо 1",
					len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("первый аргумент должен быть массивом, получено %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return NULL
		},
	},

	"последний": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("неверное количество аргументов получено %d, надо 1",
					len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("первый аргумент должен быть массивом, получено %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return NULL
		},
	},

	"добавить": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("получено неверное количество аргументов %d, надо 2",
					len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("первый аргумент должен быть массивом, получено %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array)
			arr.Elements = append(arr.Elements, args[1])
			return &object.Array{Elements: arr.Elements}
		},
	},
}
