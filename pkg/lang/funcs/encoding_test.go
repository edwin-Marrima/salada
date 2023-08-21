package funcs

import (
	"errors"
	"testing"
)

func TestBase64Decode(t *testing.T) {
	testCases := []struct {
		description    string
		args           []interface{}
		expectedResult interface{}
		expectedError  bool
	}{
		{
			description:    "Must convert string containing base64 to the original string",
			args:           []any{"SGVsbG8gV29ybGQ="},
			expectedResult: "Hello World",
		},
		{
			description:    "Must return error when provided parameter is not valid base64 text",
			args:           []any{"abc"},
			expectedResult: errors.New(`call to function "base64decode" failed: "abc" is an illegal base64 data`),
			expectedError:  true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			result, err := Base64Decode(tt.args...)
			if tt.expectedError {
				if err.Error() != tt.expectedResult.(error).Error() {
					t.Errorf("Result was incorrect.\ngot : %s\nwant: %s", err.Error(), tt.expectedResult.(error).Error())
				}
				return
			}
			if result != tt.expectedResult {
				t.Errorf("Result was incorrect.\ngot : %s\nwant: %s", result, tt.expectedResult)
			}
		})
	}
}

func TestBase64Encode(t *testing.T) {
	testCases := []struct {
		description    string
		args           []interface{}
		expectedResult interface{}
	}{
		{
			description:    "Must convert string containing base64 to the original string",
			args:           []any{"Hello World"},
			expectedResult: "SGVsbG8gV29ybGQ=",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			result, _ := Base64Encode(tt.args...)
			if result != tt.expectedResult {
				t.Errorf("Result was incorrect.\ngot : %s\nwant: %s", result, tt.expectedResult)
			}
		})
	}
}
