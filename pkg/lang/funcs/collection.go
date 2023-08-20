package funcs

import (
	"errors"
	"fmt"
	"github.com/edwin-Marrima/salada/pkg/lang/types"
	"reflect"
)

var LengthOfFunc = types.New(types.FuncSpecification{
	Name: "LengthOf",
	Description: "lengthOf determines the length of a given list, map, or string.\n If given a list or map, the " +
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
		switch reflect.TypeOf(arg).Kind() {
		case reflect.String:
			length = len(arg.(string))
		case reflect.Slice, reflect.Array:
			length = reflect.ValueOf(arg).Len()
		case reflect.Map:
			length = reflect.ValueOf(arg).Len()
		}
		return length, nil
	},
})

var AllTrueFunc = types.New(types.FuncSpecification{
	Name: "allTrue",
	Description: "allTrue returns true if all elements in a given collection are true or \"true\". " +
		"It also returns true if the collection is empty.",
	Params: []types.Parameter{
		{
			Name: "list",
			Type: []reflect.Kind{reflect.Slice},
		},
	},
	Implementation: func(args ...any) (any, error) {
		arg := args[0]
		for _, v := range arg.([]interface{}) {
			if !(v == true || v == "true") {
				return false, nil
			}
		}
		return true, nil
	},
})

var IndexFunc = types.New(types.FuncSpecification{
	Name:        "index",
	Description: "index finds the element index for a given value in a list. This function produces an error if the given value is not present in the list",
	Params: []types.Parameter{
		{
			Name: "list",
			Type: []reflect.Kind{reflect.Slice},
		},
		{
			Name: "value",
			Type: []reflect.Kind{reflect.String, reflect.Int, reflect.Float64},
		},
	},
	Implementation: func(args ...any) (any, error) {
		list, value := args[0], args[1]
		for index, v := range list.([]interface{}) {
			if v == value {
				return index, nil
			}
		}
		return nil, errors.New(fmt.Sprintf(`call to function "Index" failed: item "%s" not found`, fmt.Sprint(value)))
	},
})

func LengthOf(args ...interface{}) (any, error) {
	return LengthOfFunc.Call(args...)
}

func AllTrue(args ...interface{}) (any, error) {
	return AllTrueFunc.Call(args...)
}
func Index(args ...interface{}) (any, error) {
	return IndexFunc.Call(args...)
}
