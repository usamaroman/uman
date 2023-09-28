package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{
		store: s,
		outer: nil,
	}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	environment := NewEnvironment()
	environment.outer = outer
	return environment
}

func (e *Environment) Set(name string, obj Object) Object {
	e.store[name] = obj
	return obj
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}
