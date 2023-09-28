package object

type Builtin struct {
	Fn BuiltinFunction
}

func (b Builtin) Type() ObjectType {
	return BuiltinObj
}

func (b Builtin) Inspect() string {
	return "встроенная функция"
}
