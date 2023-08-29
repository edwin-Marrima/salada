package schema

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

const documentLoaderPath = "file://json_schema/main.json"

func Validate(schemaPath string) (bool, error) {
	sl := gojsonschema.NewSchemaLoader()

	documentLoader := gojsonschema.NewReferenceLoader(schemaPath)

	schemaLoader := gojsonschema.NewReferenceLoader(documentLoaderPath)
	schemaLoader2 := gojsonschema.NewReferenceLoader("file://json_schema/provider.json")
	schemaLoader3 := gojsonschema.NewReferenceLoader("file://json_schema/variable.json")
	err := sl.AddSchemas(schemaLoader, schemaLoader2, schemaLoader3)
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
