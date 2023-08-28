package schema

import (
	"github.com/xeipuuv/gojsonschema"
)

const documentLoaderPath = "file://json_schema/main.json"

func Validate(schemaPath string) (bool, error) {
	sl := gojsonschema.NewSchemaLoader()

	documentLoader := gojsonschema.NewReferenceLoader(schemaPath)

	schemaLoader := gojsonschema.NewReferenceLoader(documentLoaderPath)
	schemaLoader2 := gojsonschema.NewReferenceLoader("file://json_schema/provider.json")
	err := sl.AddSchemas(schemaLoader, schemaLoader2)
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

	return false, nil
}
