package funcs

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestFileExists(t *testing.T) {
	testCases := []struct {
		description    string
		args           []interface{}
		expectedResult interface{}
	}{
		{
			description:    "Must return true when file exists",
			args:           []interface{}{"./test-artifacts/filesystem_file-exists-001.txt"},
			expectedResult: true,
		},
		{
			description:    "Must return false when file does not exist",
			args:           []interface{}{"./test-artifacts/filesystem_file-exists-xxx.txt"},
			expectedResult: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			result, _ := FileExists(tt.args...)
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("Result was incorrect.\ngot : %d\nwant: %d", result, tt.expectedResult)
			}
		})
	}
}
func TestFileExtension(t *testing.T) {
	testCases := []struct {
		description    string
		args           []interface{}
		expectedResult interface{}
	}{
		{
			description:    "Must return file extension",
			args:           []interface{}{"./test-artifacts/filesystem_file-extension-001.txt"},
			expectedResult: ".txt",
		},
		{
			description:    "Must return file extension",
			args:           []interface{}{"./test-artifacts/filesystem_file-extension-xxx.abc"},
			expectedResult: ".abc",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			result, _ := FileExtension(tt.args...)
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("Result was incorrect.\ngot : %d\nwant: %d", result, tt.expectedResult)
			}
		})
	}
}

func TestAbspath(t *testing.T) {
	testCases := []struct {
		description    string
		args           []interface{}
		expectedResult interface{}
	}{
		{
			description:    "Must return an absolute representation of path",
			args:           []any{"./test-artifacts/filesystem_file-extension-001.txt"},
			expectedResult: "",
		},
		{
			description:    "Must return an absolute representation of path",
			args:           []any{"./test-artifacts/filesystem_file-extension-xxx.txt"},
			expectedResult: "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			tt.expectedResult, _ = filepath.Abs(tt.args[0].(string))
			result, _ := Abspath(tt.args...)
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("Result was incorrect.\ngot : %d\nwant: %d", result, tt.expectedResult)
			}
		})
	}
}
