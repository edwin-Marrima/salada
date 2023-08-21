package funcs

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/edwin-Marrima/salada/pkg/lang/types"
	"reflect"
)

var Base64DecodeFunc = types.New(types.FuncSpecification{
	Name: "base64decode",
	Description: "base64decode takes a string containing a Base64 character sequence and returns the original string." +
		" If the bytes after Base64 decoding are not valid UTF-8, this function produces an error.",
	Params: []types.Parameter{
		{
			Name: "base64Content",
			Type: []reflect.Kind{reflect.String},
		},
	},
	Implementation: func(args ...any) (any, error) {
		base64Text := args[0]
		s, err := base64.StdEncoding.DecodeString(base64Text.(string))
		if err != nil {
			fmt.Println(err.Error())
			return nil, errors.New(fmt.Sprintf(`call to function "base64decode" failed: "%s" is an illegal base64 data`, base64Text.(string)))
		}
		return string(s), nil
	},
})

var Base64EncodeFunc = types.New(types.FuncSpecification{
	Name:        "urlencode",
	Description: "urlencode applies URL encoding to a given string.",
	Params: []types.Parameter{
		{
			Name: "text",
			Type: []reflect.Kind{reflect.String},
		},
	},
	Implementation: func(args ...any) (any, error) {
		s := args[0]
		base64 := base64.StdEncoding.EncodeToString([]byte(s.(string)))
		return base64, nil
	},
})

func Base64Decode(args ...any) (any, error) {
	return Base64DecodeFunc.Call(args...)
}
func Base64Encode(args ...any) (any, error) {
	return Base64EncodeFunc.Call(args...)
}
