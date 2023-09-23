package tf_plan

import (
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
				Variables: []Variable{
					{
						Name:        "aaa",
						Description: "value",
						HasDefault:  true,
						Sensitive:   true,
					},
					{
						Name:        "environment",
						Description: "value",
						HasDefault:  true,
						Sensitive:   false,
					},
				},
				ProviderConfiguration: map[string]ProviderConfiguration{
					"aws": {
						Name:              "aws",
						FullName:          "registry.terraform.io/hashicorp/aws",
						VersionConstraint: "~> 4.0",
					},
				},
				FormatVersion: "1.2",
				ModuleRemovedResources: map[string]Resource{
					"null_resource.cluster_dummy": {
						Address:      "null_resource.cluster_dummy",
						Type:         "null_resource",
						Name:         "cluster_dummy",
						ProviderName: "registry.terraform.io/hashicorp/null",
						Actions:      []string{"delete"},
						Values: map[string]interface{}{
							"id": "1580278013551376455",
							"triggers": map[string]interface{}{
								"cluster_instance_ids": "aaaaa",
							},
						},
						SensitiveValues: map[string]interface{}{
							"triggers": map[string]interface{}{},
						},
					},
				},
				Resources: map[string]Resource{
					"aws_cloudwatch_log_group.yada": {
						Address:      "aws_cloudwatch_log_group.yada",
						Type:         "aws_cloudwatch_log_group",
						Name:         "yada",
						ProviderName: "registry.terraform.io/hashicorp/aws",
						Actions:      []string{"create"},
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
					"aws_s3_bucket.example": {
						Address:      "aws_s3_bucket.example",
						Type:         "aws_s3_bucket",
						Name:         "example",
						ProviderName: "registry.terraform.io/hashicorp/aws",
						Actions:      []string{"create"},
						Values: map[string]interface{}{
							"bucket":        "aaaaa-dev",
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
					"module.bucket_provisioner.aws_s3_bucket.example": {
						Address:      "module.bucket_provisioner.aws_s3_bucket.example",
						Type:         "aws_s3_bucket",
						Name:         "example",
						Actions:      []string{"create"},
						ProviderName: "registry.terraform.io/hashicorp/aws",
						Values: map[string]interface{}{
							"bucket":        "bucket-test",
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
	}
	for _, tt := range testCase {
		t.Run(tt.description, func(t *testing.T) {
			outcome, _ := ParsePlanFile(tt.inputFilePath)
			assert.Equal(t, &tt.expectedOutcome, outcome)

		})
	}
}
