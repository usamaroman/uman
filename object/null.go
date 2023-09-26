package object

type Null struct{}

func (n Null) Inspect() string {
	return "ничего"
}

func (n Null) Type() ObjectType {
	return NullObj
}
