package funcs

import (
	"fmt"
	"github.com/edwin-Marrima/salada/pkg/lang/types"
	"reflect"
)

var LengthOfFunc = types.New(types.FuncSpecification{
	Name: "List length",
	Description: "length determines the length of a given list, map, or string.\n If given a list or map, the " +
		"result is the number of elements in that collection. If given a string, the result is the number of characters in the string.",
	Params: []types.Parameter{
		{
			Name: "x",
			Type: []reflect.Kind{reflect.Map, reflect.String, reflect.Slice},
		},
	},
	Implementation: func(args ...any) (any, error) {
		arg := args[0]
		var length int
		fmt.Println("===>", reflect.TypeOf(arg).Kind())
		switch reflect.TypeOf(arg).Kind() {
		case reflect.String:
			length = len(arg.(string))
		}
		return length, nil
	},
})

func LengthOf(args ...interface{}) (any, error) {
	return LengthOfFunc.Call(args...)
}
