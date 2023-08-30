package schema

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

const documentLoaderPath = "file://json_schema/main.json"

func Validate(schemaPath string) (bool, error) {
	sl := gojsonschema.NewSchemaLoader()

	documentLoader := gojsonschema.NewReferenceLoader(schemaPath)

	schemaLoader := gojsonschema.NewReferenceLoader("file://json_schema/main.json")
	providerLoader := gojsonschema.NewReferenceLoader("file://json_schema/provider.json")
	variableLoader := gojsonschema.NewReferenceLoader("file://json_schema/variable.json")
	err := sl.AddSchemas(schemaLoader, providerLoader, variableLoader)
	if err != nil {
		return false, err
	}
	schema, err := sl.Compile(schemaLoader)
	if err != nil {
		return false, err
	}
	result, err := schema.Validate(documentLoader)
	if err != nil {
		return false, err
	}
	if result.Valid() {
		return true, err
	}
	for _, e := range result.Errors() {
		fmt.Println("=>", e)
	}
	return false, nil
}
