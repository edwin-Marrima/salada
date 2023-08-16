package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFunction_Call(t *testing.T) {
	t.Run("Must return error when args are not compliant with Parameters specified on FuncSpecification", func(t *testing.T) {
		tests := []struct {
			expectedError error
			spec          FuncSpecification
		}{
			{
				spec: FuncSpecification{
					Name: "max",
					Params: []Parameter{
						{
							Name: "x",
							Type: reflect.Int64,
						},
						{
							Name: "y",
							Type: reflect.Int64,
						},
					},
				},
				expectedError: NewWrongFunctionParametersError("max", "max(int64, int64)", "max(string)"),
			},
		}
		for _, testCase := range tests {
			funcT := New(testCase.spec)
			_, err := funcT.Call("a")
			fmt.Println(err.Error())
			if !reflect.DeepEqual(err, testCase.expectedError) {
				t.Errorf("Result was incorrect.\ngot: %s\nwant: %s.", err.Error(), testCase.expectedError.Error())
			}
		}
	})

}
