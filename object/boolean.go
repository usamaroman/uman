package object

import "fmt"

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	if b.Value == true {
		return fmt.Sprintf("%s", "истина")
	}
	return fmt.Sprintf("%s", "ложь")
}

func (b *Boolean) Type() ObjectType {
	return BooleanObj
}
