package types

import (
	"fmt"
	"reflect"
)

type Function struct {
	spec *FuncSpecification
}

// New creates a new function with the given specification.
func New(spec FuncSpecification) Function {
	return Function{
		spec: &spec,
	}
}

// Call actually calls the function with the given arguments, which must
// conform to the function's parameter specification or an error will be returned.
func (f Function) Call(args ...any) (FuncReturn, error) {
	if len(args) != len(f.spec.Params) {
		return FuncReturn{}, NewWrongFunctionParametersError(f.spec.Name, f.expectedParametersToString(), f.sentParametersToString(args))
	}
	return f.spec.Implementation(args)
}

func (f Function) sentParametersToString(parameters []any) string {
	argsString := fmt.Sprintf("%s(", f.spec.Name)
	for i, arg := range parameters {
		if i > 0 {
			argsString += ", "
		}
		argsString += fmt.Sprintf(`%s`, reflect.TypeOf(arg).Kind().String())

	}
	return argsString + ")"
}
func (f Function) expectedParametersToString() string {
	paramsString := fmt.Sprintf("%s(", f.spec.Name)
	for i, p := range f.spec.Params {
		if i > 0 {
			paramsString += ", "
		}
		paramsString += fmt.Sprintf(`%s`, p.Type.String())
	}
	return paramsString + ")"
}
