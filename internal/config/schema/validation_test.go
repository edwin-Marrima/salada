package schema

import (
	"fmt"
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
		{
			description:     "Must return true when `configuration_parameters` contains objects with oneOf: [allowed,value] sub_properties",
			document:        "file://test_artifacts/provider_validation_002.json",
			expectedOutcome: true,
		},
		{
			description:     "Must return false when  `configuration_parameters` contains objects with both: [allowed,value] sub_properties",
			document:        "file://test_artifacts/provider_validation_003.json",
			expectedOutcome: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			b, err := Validate(tt.document)
			fmt.Println("====>", err)
			if !reflect.DeepEqual(b, tt.expectedOutcome) {
				t.Errorf("Result was incorrect.\ngot : %t\nwant: %t", b, tt.expectedOutcome)
			}
		})
	}
}