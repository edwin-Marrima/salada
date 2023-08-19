package types

import "reflect"

type Parameter struct {
	// Name is an optional name for the argument. This package ignores this
	// value, but callers may use it for documentation, etc.
	Name string
	// Description is an optional description for the argument.
	Description string
	// Type is type that any argument for this parameter must conform to.
	// The parameter must be compatible with at least one element of the type list, elements are evaluated using `or`
	Type []reflect.Kind
}

type FuncSpecification struct {
	// Name is a function name
	Name string
	// Description is an optional description for the function specification.
	Description string
	// Params is a description of the positional parameters for the function.
	// The standard checking logic rejects any calls that do not provide
	// arguments conforming to this definition, freeing the function
	// implementer from dealing with such inconsistencies.
	Params []Parameter
	// Implementation implements the function's behavior.
	Implementation func(args ...any) (FuncReturn, error)
}

type FuncReturn struct {
	Value any
	Type  reflect.Type
}
