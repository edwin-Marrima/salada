package funcs

import (
	"errors"
	"github.com/edwin-Marrima/salada/pkg/lang/types"
	"os"
	"path"
	"path/filepath"
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
			Name: "path",
			Type: []reflect.Kind{reflect.String},
		},
	},
	Implementation: func(args ...any) (any, error) {
		filePath := args[0]
		ext := path.Ext(filePath.(string))
		return ext, nil
	},
})

var AbspathFunc = types.New(types.FuncSpecification{
	Name: "abspath",
	Description: "abspath takes a string containing a filesystem path and converts it to an absolute path. That is, if the path is not absolute," +
		" it will be joined with the current working directory.",
	Params: []types.Parameter{
		{
			Name: "path",
			Type: []reflect.Kind{reflect.String},
		},
	},
	Implementation: func(args ...any) (any, error) {
		filePath := args[0]
		p, _ := filepath.Abs(filePath.(string))
		return p, nil
	},
})

func FileExists(args ...any) (any, error) {
	return FileExistsFunc.Call(args...)
}
func FileExtension(args ...any) (any, error) {
	return FileExtensionFunc.Call(args...)
}
func Abspath(args ...any) (any, error) {
	return AbspathFunc.Call(args...)
}
