package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidateProviderSchema(t *testing.T) {
	tests := []struct {
		description     string
		document        string
		expectedOutcome interface{}
	}{
		{
			description:     "Must return true when provided configuration File is valid",
			document:        "file://test_artifacts/provider/provider_validation_001.json",
			expectedOutcome: true,
		},
		{
			description:     "Must return true when `configuration_parameters` contains array of objects with name and value property",
			document:        "file://test_artifacts/provider/provider_validation_002.json",
			expectedOutcome: true,
		},
		{
			description:     "Must return false when  `configuration_parameters` does not contain both: [name and value]",
			document:        "file://test_artifacts/provider/provider_validation_003.json",
			expectedOutcome: false,
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

func TestValidateVariableSchema(t *testing.T) {
	tests := []struct {
		description     string
		document        string
		expectedOutcome interface{}
	}{
		{
			description:     "Must return true when provided configuration File is valid",
			document:        "file://test_artifacts/variable/validation_001.json",
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
func TestValidateResourcesSchema(t *testing.T) {
	tests := []struct {
		description     string
		document        string
		expectedOutcome interface{}
	}{
		{
			description:     "Must return true when resources configuration File is valid",
			document:        "file://test_artifacts/resource/validation_001.json",
			expectedOutcome: true,
		},
		{
			description:     "Attributes property MUST be an array of objects",
			document:        "file://test_artifacts/resource/validation_002.json",
			expectedOutcome: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			b, err := Validate(tt.document)
			fmt.Println("=======>X", err)
			if !reflect.DeepEqual(b, tt.expectedOutcome) {
				t.Errorf("Result was incorrect.\ngot : %t\nwant: %t", b, tt.expectedOutcome)
			}
		})
	}
}
