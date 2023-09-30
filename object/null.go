package object

type Null struct{}

func (n Null) Inspect() string {
	return ""
}

func (n Null) Type() ObjectType {
	return NullObj
}
