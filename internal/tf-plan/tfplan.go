package tf_plan

import (
	"encoding/json"
	"os"
)

type PlanFile struct {
	FormatVersion string        `json:"format_version"`
	PlannedValues PlannedValues `json:"planned_values"`
}
type PlannedValues struct {
	RootModule `json:"root_module"`
}
type RootModule struct {
	Resources []Resource `json:"resources"`
}
type Resource struct {
	Address         string                 `json:"address"`
	Type            string                 `json:"type"`
	Name            string                 `json:"name"`
	ProviderName    string                 `json:"provider_name"`
	Values          map[string]interface{} `json:"values"`
	Actions         []string
	SensitiveValues map[string]interface{} `json:"sensitive_values"`
}

// ProviderConfiguration holds terraform provider configuration
type ProviderConfiguration struct {
	Name              string
	FullName          string
	VersionConstraint string
	Expressions       map[string][]ExpressionConfig
}

type ExpressionConfig map[string]References

type References struct {
	References []interface{}
}

func ParsePlanFile(planFilepath string) (*PlanFile, error) {
	data, err := os.ReadFile(planFilepath)
	if err != nil {
		return nil, err
	}
	var tfplan PlanFile
	if err = json.Unmarshal(data, &tfplan); err != nil {
		return nil, err
	}

	return &tfplan, nil
}
