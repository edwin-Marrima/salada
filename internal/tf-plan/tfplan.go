package tf_plan

import (
	"encoding/json"
	"os"
)

type PlanFile struct {
	FormatVersion string              `json:"format_version"`
	Resources     map[string]Resource `json:"resources"`
	// ModuleRemovedResources contains resources whose definition was deleted from the mods
	ModuleRemovedResources map[string]Resource
	// ProviderConfiguration holds provider configuration
	ProviderConfiguration map[string]ProviderConfiguration
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
}

func ParsePlanFile(planFilepath string) (*PlanFile, error) {
	data, err := os.ReadFile(planFilepath)
	if err != nil {
		return nil, err
	}
	var tfplan PlanFile

	var planFile map[string]interface{}
	if err = json.Unmarshal(data, &planFile); err != nil {
		return nil, err
	}
	tfplan.FormatVersion = planFile["format_version"].(string)
	tfplan.Resources, tfplan.ModuleRemovedResources = extractModuleResources(planFile)
	tfplan.ProviderConfiguration = extractProvider(planFile)
	return &tfplan, nil
}
func extractProvider(planFile map[string]interface{}) map[string]ProviderConfiguration {
	provider := make(map[string]ProviderConfiguration)

	config := planFile["configuration"].(map[string]interface{})
	providerConfiguration := config["provider_config"].(map[string]interface{})
	for k, v := range providerConfiguration {
		prv := v.(map[string]interface{})
		provider[k] = ProviderConfiguration{
			Name:              prv["name"].(string),
			FullName:          prv["full_name"].(string),
			VersionConstraint: prv["version_constraint"].(string),
		}
	}
	return provider
}
func extractModuleResources(planFile map[string]interface{}) (map[string]Resource, map[string]Resource) {
	plannedResources := make(map[string]Resource)
	deletedResources := make(map[string]Resource)
	plannedValues := planFile["planned_values"].(map[string]interface{})
	rootModule := plannedValues["root_module"].(map[string]interface{})

	//read root module resources
	for _, resource := range rootModule["resources"].([]interface{}) {
		r := resource.(map[string]interface{})
		plannedResources[r["address"].(string)] = Resource{
			Address:         r["address"].(string),
			Type:            r["type"].(string),
			Name:            r["name"].(string),
			ProviderName:    r["provider_name"].(string),
			Values:          r["values"].(map[string]interface{}),
			SensitiveValues: r["sensitive_values"].(map[string]interface{}),
		}
	}

	//read sub_modules resources
	childModules := rootModule["child_modules"]
	for _, childModule := range childModules.([]interface{}) {
		ch := childModule.(map[string]interface{})["resources"].([]interface{})
		for _, resource := range ch {
			r := resource.(map[string]interface{})
			plannedResources[r["address"].(string)] = Resource{
				Address:         r["address"].(string),
				Type:            r["type"].(string),
				Name:            r["name"].(string),
				ProviderName:    r["provider_name"].(string),
				Values:          r["values"].(map[string]interface{}),
				SensitiveValues: r["sensitive_values"].(map[string]interface{}),
			}
		}
	}

	//read actions
	for _, resourceChanges := range planFile["resource_changes"].([]interface{}) {
		rc := resourceChanges.(map[string]interface{})
		change := rc["change"].(map[string]interface{})
		actions := change["actions"].([]interface{})

		address := rc["address"].(string)
		if v, found := plannedResources[address]; found {
			for _, action := range actions {
				v.Actions = append(v.Actions, action.(string))
			}
			plannedResources[address] = v
		}
	}
	// retrieve module erased resources
	priorState := planFile["prior_state"].(map[string]interface{})
	values := priorState["values"].(map[string]interface{})
	rootModule = values["root_module"].(map[string]interface{})
	for _, resource := range rootModule["resources"].([]interface{}) {
		rc := resource.(map[string]interface{})
		if _, found := plannedResources[rc["address"].(string)]; !found {
			deletedResources[rc["address"].(string)] = Resource{
				Address:         rc["address"].(string),
				Type:            rc["type"].(string),
				Name:            rc["name"].(string),
				Actions:         []string{"delete"},
				ProviderName:    rc["provider_name"].(string),
				Values:          rc["values"].(map[string]interface{}),
				SensitiveValues: rc["sensitive_values"].(map[string]interface{}),
			}
		}
	}
	return plannedResources, deletedResources
}
