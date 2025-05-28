// environment.go
// Implementa los entornos léxicos (ámbitos de variables) con un puntero opcional al entorno externo para modelar scopes anidados.
//
// Comentarios añadidos automáticamente el 24/05/2025 para proporcionar
// una explicación detallada en español de cada parte del código.
// ----------------------------------------------------------------------------

package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// -----------------------------------------------------------------------------
// Declaración de tipo: Environment
// -----------------------------------------------------------------------------
// Environment modela una estructura
// dentro del AST o del sistema de objetos. Consulte los campos y sus
// métodos asociados para más detalles.

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
