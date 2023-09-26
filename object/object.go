package object

type ObjectType string

const (
	IntegerObj = "INTEGER"
	BooleanObj = "BOOLEAN"
	StringObj  = "STRING"
	NullObj    = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
