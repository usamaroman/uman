package object

type ReturnValue struct {
	Value Object
}

func (r *ReturnValue) Type() ObjectType {
	return ReturnValueObj
}

func (r *ReturnValue) Inspect() string {
	return r.Value.Inspect()
}
