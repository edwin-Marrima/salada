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
		return FuncReturn{}, fmt.Errorf(errWrongFunctionParameters, f.spec.Name, f.convertExpectedParametersToString(), f.convertSentParametersToString(args))
	}
	for i, arg := range args {
		if !f.typeIsEqual(reflect.TypeOf(arg).Kind(), f.spec.Params[i].Type) {
			return FuncReturn{}, fmt.Errorf(errWrongFunctionParameters, f.spec.Name, f.convertExpectedParametersToString(), f.convertSentParametersToString(args))
		}
	}
	return f.spec.Implementation(args)
}

func (f Function) typeIsEqual(t reflect.Kind, expectedType []reflect.Kind) bool {
	for _, v := range expectedType {
		if v != t {
			return false
		}
	}
	return true
}

// convertSentParametersToString transforms a list of sent parameters into a formatted string in order to use in error messages.
func (f Function) convertSentParametersToString(parameters []any) string {
	argsString := fmt.Sprintf("%s(", f.spec.Name)
	for i, arg := range parameters {
		if i > 0 {
			argsString += ", "
		}
		argsString += fmt.Sprintf(`%s %s`, f.spec.Params[i].Name, reflect.TypeOf(arg).Kind().String())
	}
	return argsString + ")"
}

// convertExpectedParametersToString transforms a list of expected parameters into a formatted string in order to use in error messages.
func (f Function) convertExpectedParametersToString() string {
	paramsString := fmt.Sprintf("%s(", f.spec.Name)
	for i, p := range f.spec.Params {
		if i > 0 {
			paramsString += ", "
		}
		parametersSupportedTypes := p.Type[0].String()
		for _, v := range p.Type[1:] {
			parametersSupportedTypes += fmt.Sprintf(" | %s", v.String())
		}
		paramsString += fmt.Sprintf(`%s %s`, p.Name, parametersSupportedTypes)
	}
	return paramsString + ")"
}
