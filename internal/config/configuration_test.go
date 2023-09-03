package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		description     string
		filePath        string
		expectedOutcome Configuration
	}{
		{
			description: "MUST return configuration object",
			filePath:    "test_artifacts/parse_001.json",
			expectedOutcome: Configuration{
				Provider: &Provider{
					List: []ProviderList{
						{
							Name:    "aws",
							Source:  "hashicorp/aws",
							Version: "greaterThan(5.9)",
							Allowed: true,
							ConfigurationParameter: []Property{
								{
									Name:  "access_key",
									Value: nil,
								},
								{
									Name:  "region",
									Value: toPointer("af-south-1"),
								},
							},
						},
					},
				},
				Variable: Variable{
					Name:        Content{Value: "match_regex(`^*`)"},
					Description: Content{Value: "match_regex(`^*`)"},
				},
				Resource: []Resource{
					{
						Type: "aws_s3_bucket",
						ChangeActions: ChangeActions{
							Update: &CronExpression{Expression: "* * * * *"},
						},
						Allowed:    Allowed{When: When{Expression: "false"}},
						Attributes: []Property{{Name: "bucket_prefix", Value: toPointer("x")}},
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			c, err := Parse(tt.filePath)
			fmt.Println("===>", err)
			assert.Equal(t, tt.expectedOutcome, *c)
		})
	}
}

func toPointer[T any](value T) *T {
	return &value
}
