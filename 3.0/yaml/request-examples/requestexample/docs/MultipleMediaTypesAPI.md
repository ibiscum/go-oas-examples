# \MultipleMediaTypesAPI

All URIs are relative to *https://httpbin.org/anything*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostRequestBodyMultiMediaTypes**](MultipleMediaTypesAPI.md#PostRequestBodyMultiMediaTypes) | **Post** /requestBody-multi-media-types | Within &#x60;examples&#x60;



## PostRequestBodyMultiMediaTypes

> PostRequestBodyMultiMediaTypes(ctx).Execute()

Within `examples`



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ibiscum/go-oas-examples"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MultipleMediaTypesAPI.PostRequestBodyMultiMediaTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipleMediaTypesAPI.PostRequestBodyMultiMediaTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostRequestBodyMultiMediaTypesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain, application/json
- **Accept**: text/plain, application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

