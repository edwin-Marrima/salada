package funcs

import (
	"testing"
)

func TestLengthOf(t *testing.T) {
	testCases := []struct {
		description    string
		args           interface{}
		expectedResult interface{}
	}{
		{
			description:    "Must return string length",
			args:           "abc",
			expectedResult: 3,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			result, _ := LengthOf(tt.args)
			if result != tt.expectedResult {
				t.Errorf("Result was incorrect.\ngot: %d\nwant: %d.", result, tt.expectedResult)
			}
		})
	}
}
