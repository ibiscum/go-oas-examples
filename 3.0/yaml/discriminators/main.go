package main

import (
	"fmt"
	"os"

	"github.com/pb33f/libopenapi"
)

func main() {

	// load the discriminators sample OpenAPI specification into a byte slice.
	discriminators, _ := os.ReadFile("discriminators.yaml")

	// create a new Document from from the byte slice.
	document, err := libopenapi.NewDocument(discriminators)

	// if anything went wrong, an error is thrown
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	docModel, errors := document.BuildV3Model()
	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("error: %e\n", errors[i])
		}
		panic(fmt.Sprintf("cannot create v3 model from document: %d errors reported", len(errors)))
	}

	for pathPairs := docModel.Model.Paths.PathItems.First(); pathPairs != nil; pathPairs = pathPairs.Next() {
		pathName := pathPairs.Key()
		pathItem := pathPairs.Value()
		fmt.Printf("Path %s has %d operations\n", pathName, pathItem.GetOperations().Len())
	}

	for schemaPairs := docModel.Model.Components.Schemas.First(); schemaPairs != nil; schemaPairs = schemaPairs.Next() {
		// get the name of the schema
		schemaName := schemaPairs.Key()

		// get the schema object from the map
		schemaValue := schemaPairs.Value()

		// build the schema
		schema := schemaValue.Schema()

		// if the schema has properties, print the number of properties
		if schema != nil && schema.Properties != nil {
			fmt.Printf("Schema '%s' has %d properties\n", schemaName, schema.Properties.Len())
		}
	}
}
