package tf_plan

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePlanFile(t *testing.T) {
	testCase := []struct {
		inputFilePath   string
		expectedOutcome PlanFile
		description     string
	}{
		{
			description:   "-",
			inputFilePath: "testFile/1.json",
			expectedOutcome: PlanFile{
				FormatVersion: "1.2",
				PlannedValues: PlannedValues{
					RootModule: RootModule{
						Resources: []Resource{
							{
								Address:      "aws_cloudwatch_log_group.yada",
								Type:         "aws_cloudwatch_log_group",
								Name:         "yada",
								ProviderName: "registry.terraform.io/hashicorp/aws",
								Values: map[string]interface{}{
									"kms_key_id":        "arn:sadsdsdsdsdkl",
									"name":              "Yada",
									"retention_in_days": float64(5),
									"skip_destroy":      false,
									"tags": map[string]interface{}{
										"Application": "serviceA",
										"Environment": "production",
									},
									"tags_all": map[string]interface{}{
										"Application": "serviceA",
										"Environment": "production",
									},
								},
								SensitiveValues: map[string]interface{}{
									"tags":     map[string]interface{}{},
									"tags_all": map[string]interface{}{},
								},
							},
							{
								Address:      "aws_s3_bucket.example",
								Type:         "aws_s3_bucket",
								Name:         "example",
								ProviderName: "registry.terraform.io/hashicorp/aws",
								Values: map[string]interface{}{
									"bucket":        "my-tf-test-bucket",
									"force_destroy": false,
									"tags": map[string]interface{}{
										"Environment": "Dev",
										"Name":        "My bucket",
									},
									"tags_all": map[string]interface{}{
										"Environment": "Dev",
										"Name":        "My bucket",
									},
									"timeouts": nil,
								},
								SensitiveValues: map[string]interface{}{
									"cors_rule":                            []any{},
									"grant":                                []any{},
									"lifecycle_rule":                       []any{},
									"logging":                              []any{},
									"object_lock_configuration":            []any{},
									"replication_configuration":            []any{},
									"server_side_encryption_configuration": []any{},
									"tags":                                 map[string]interface{}{},
									"tags_all":                             map[string]interface{}{},
									"versioning":                           []any{},
									"website":                              []any{},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range testCase {
		t.Run(tt.description, func(t *testing.T) {
			outcome, err := ParsePlanFile(tt.inputFilePath)
			assert.Equal(t, &tt.expectedOutcome, outcome)
			fmt.Println(outcome, err)
		})
	}
}
