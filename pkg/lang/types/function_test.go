package types

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestFunction_Call(t *testing.T) {

	tests := []struct {
		expectedError   error
		spec            FuncSpecification
		funcArgs        []interface{}
		testDescription string
	}{
		{
			testDescription: "Must return error when number os arguments does not match",
			funcArgs:        []interface{}{"2"},
			spec: FuncSpecification{
				Name: "max",
				Params: []Parameter{
					{
						Name: "x",
						Type: []reflect.Kind{reflect.Int64},
					},
					{
						Name: "y",
						Type: []reflect.Kind{reflect.Int64},
					},
				},
			},
			expectedError: fmt.Errorf(errWrongFunctionParameters, "max", "max(x int64, y int64)", "max(x string)"),
		},
		{
			testDescription: "Must return error when parameters type does not match",
			funcArgs:        []interface{}{"2", 4},
			spec: FuncSpecification{
				Name: "max",
				Params: []Parameter{
					{
						Name: "x",
						Type: []reflect.Kind{reflect.Int64, reflect.String},
					},
					{
						Name: "y",
						Type: []reflect.Kind{reflect.Int64},
					},
				},
			},
			expectedError: fmt.Errorf(errWrongFunctionParameters, "max", "max(x int64 | string, y int64)", "max(x string, y int)"),
		},
		{
			testDescription: "Must run customValidations after type validation and the customValidation Func error must be contained on the returned error",
			funcArgs:        []interface{}{0},
			spec: FuncSpecification{
				Name: "div",
				Params: []Parameter{
					{
						Name: "x",
						Type: []reflect.Kind{reflect.Int},
						ExtendedValidation: func(args any) error {
							if args.(int) == 0 {
								return errors.New("division by zero is not allowed")
							}
							return nil
						},
					},
				},
			},
			expectedError: errors.New("division by zero is not allowed"),
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.testDescription, func(t *testing.T) {
			funcT := New(testCase.spec)
			_, err := funcT.Call(testCase.funcArgs...)
			if !reflect.DeepEqual(err, testCase.expectedError) {
				t.Errorf("Result was incorrect.\ngot: %s\nwant: %s.", err.Error(), testCase.expectedError.Error())
			}
		})
	}

}
