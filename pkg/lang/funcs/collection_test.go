package funcs

import (
	"errors"
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
		{
			description:    "Must return slice length",
			args:           []string{"a", "b"},
			expectedResult: 2,
		},
		{
			description:    "Must return slice length",
			args:           map[string]string{"a": "b", "c": "d"},
			expectedResult: 2,
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

func TestAllTrue(t *testing.T) {
	testCases := []struct {
		description    string
		args           interface{}
		expectedResult interface{}
	}{
		{
			description:    `Must Return true when "list" only contains true or "true"`,
			args:           []interface{}{true, "true", true},
			expectedResult: true,
		},
		{
			description:    `Must Return false when "list" when collection contains element other than true or "true"`,
			args:           []interface{}{true, "uy", true},
			expectedResult: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			result, _ := AllTrue(tt.args)
			if result != tt.expectedResult {
				t.Errorf("Result was incorrect.\ngot : %d\nwant: %d", result, tt.expectedResult)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	testCases := []struct {
		description    string
		args           []interface{}
		expectedResult interface{}
		expectErr      bool
	}{
		{
			description:    "Must return element index in given list",
			args:           []interface{}{[]any{"ew", 1.34, 7, 9}, 7},
			expectedResult: 2,
			expectErr:      false,
		},
		{
			description:    "Must return element index in given list",
			args:           []interface{}{[]any{1.34, 7, 9}, 3},
			expectedResult: errors.New(`call to function "Index" failed: item "3" not found`),
			expectErr:      true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			result, err := Index(tt.args...)
			if tt.expectErr {
				if err.Error() != tt.expectedResult.(error).Error() {
					t.Errorf("Result was incorrect.\ngot : %s\nwant: %s", err.Error(), tt.expectedResult.(error).Error())
				}
				return
			}
			if result != tt.expectedResult {
				t.Errorf("Result was incorrect.\ngot : %d\nwant: %d", result, tt.expectedResult)

			}

		})
	}
}
