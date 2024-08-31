package main

import (
	"fmt"
	"os"

	"github.com/pb33f/libopenapi"
)

func main() {

	// load in the petstore sample OpenAPI specification
	// into a byte slice.
	petstore, _ := os.ReadFile("petstore-simple.json")

	// create a new Document from from the byte slice.
	document, err := libopenapi.NewDocument(petstore)

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

}
