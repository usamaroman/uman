package object

type ObjectType string

const (
	IntegerObj     = "INTEGER"
	BooleanObj     = "BOOLEAN"
	StringObj      = "STRING"
	NullObj        = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj       = "ERROR"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
