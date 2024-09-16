package api

import (
	"context"
	"sync"
)

type CustomServer struct {
	Lock sync.Mutex
}

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ StrictServerInterface = (*CustomServer)(nil)

func NewCustomServer() *CustomServer {
	return &CustomServer{}
}

// GET /anything
func (cs *CustomServer) OneOfWithTopLevelDiscriminatorAndMapping(ctx context.Context, request OneOfWithTopLevelDiscriminatorAndMappingRequestObject) (OneOfWithTopLevelDiscriminatorAndMappingResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()

	return OneOfWithTopLevelDiscriminatorAndMapping200Response{}, nil
}

func (cs *CustomServer) MappingOfSchemaNames(ctx context.Context, request MappingOfSchemaNamesRequestObject) (MappingOfSchemaNamesResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()

	return MappingOfSchemaNames200Response{}, nil
}

func (cs *CustomServer) MappingWithDuplicateSchemas(ctx context.Context, request MappingWithDuplicateSchemasRequestObject) (MappingWithDuplicateSchemasResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()

	return MappingWithDuplicateSchemas201Response{}, nil
}

func (cs *CustomServer) OneOfWithImproperlyPlacedDiscriminator(ctx context.Context, request OneOfWithImproperlyPlacedDiscriminatorRequestObject) (OneOfWithImproperlyPlacedDiscriminatorResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()

	return OneOfWithImproperlyPlacedDiscriminator200Response{}, nil
}
