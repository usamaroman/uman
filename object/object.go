package object

type ObjectType string

type BuiltinFunction func(args ...Object) Object

const (
	IntegerObj     = "INTEGER"
	BooleanObj     = "BOOLEAN"
	StringObj      = "STRING"
	NullObj        = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj       = "ERROR"
	FunctionObj    = "FUNCTION"
	BuiltinObj     = "BUILTIN"
	ArrayObj       = "ARRAY"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
