package schema

import (
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		description     string
		document        string
		expectedOutcome interface{}
	}{
		{
			description:     "Must return true when provided configuration File is valid",
			document:        "file://test_artifacts/provider_validation_001.json",
			expectedOutcome: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			b, _ := Validate(tt.document)
			if !reflect.DeepEqual(b, tt.expectedOutcome) {
				t.Errorf("Result was incorrect.\ngot : %t\nwant: %t", b, tt.expectedOutcome)
			}
		})
	}
}
