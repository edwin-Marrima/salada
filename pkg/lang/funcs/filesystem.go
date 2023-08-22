package funcs

import (
	"errors"
	"github.com/edwin-Marrima/salada/pkg/lang/types"
	"os"
	"path"
	"reflect"
)

var FileExistsFunc = types.New(types.FuncSpecification{
	Name:        "fileexists",
	Description: "fileexists determines whether a file exists at a given path.",
	Params: []types.Parameter{
		{
			Name: "path",
			Type: []reflect.Kind{reflect.String},
		},
	},
	Implementation: func(args ...any) (any, error) {
		path := args[0]
		if _, err := os.Stat(path.(string)); errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return true, nil
	},
})

var FileExtensionFunc = types.New(types.FuncSpecification{
	Name: "fileExtension",
	Description: "returns the file name extension used by path. The extension is the suffix beginning " +
		"at the final dot in the final slash-separated element of path; it is empty if there is no dot.",
	Params: []types.Parameter{
		{
			Name: "fileExtension",
			Type: []reflect.Kind{reflect.String},
		},
	},
	Implementation: func(args ...any) (any, error) {
		filePath := args[0]
		ext := path.Ext(filePath.(string))
		return ext, nil
	},
})

func FileExists(args ...any) (any, error) {
	return FileExistsFunc.Call(args...)
}
func FileExtension(args ...any) (any, error) {
	return FileExtensionFunc.Call(args...)
}
