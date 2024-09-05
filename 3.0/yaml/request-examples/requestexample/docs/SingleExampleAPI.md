# \SingleExampleAPI

All URIs are relative to *https://httpbin.org/anything*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchRequestBody**](SingleExampleAPI.md#PatchRequestBody) | **Patch** /requestBody | Within &#x60;example&#x60;
[**PostRequestBodyPrimitiveExample**](SingleExampleAPI.md#PostRequestBodyPrimitiveExample) | **Post** /requestBody-primitive-example | Primitive &#x60;example&#x60;
[**PutParameterExamplesParam1Param2**](SingleExampleAPI.md#PutParameterExamplesParam1Param2) | **Put** /parameterExamples/{param1}/{param2} | Within &#x60;example&#x60;
[**PutRequestBody**](SingleExampleAPI.md#PutRequestBody) | **Put** /requestBody | &#x60;$ref&#x60; within &#x60;example&#x60;



## PatchRequestBody

> User PatchRequestBody(ctx).User(user).Execute()

Within `example`



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
	user := *openapiclient.NewUser() // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SingleExampleAPI.PatchRequestBody(context.Background()).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SingleExampleAPI.PatchRequestBody``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRequestBody`: User
	fmt.Fprintf(os.Stdout, "Response from `SingleExampleAPI.PatchRequestBody`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchRequestBodyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRequestBodyPrimitiveExample

> string PostRequestBodyPrimitiveExample(ctx).Body(body).Execute()

Primitive `example`



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
	body := "column1,column2,column3,column4" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SingleExampleAPI.PostRequestBodyPrimitiveExample(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SingleExampleAPI.PostRequestBodyPrimitiveExample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRequestBodyPrimitiveExample`: string
	fmt.Fprintf(os.Stdout, "Response from `SingleExampleAPI.PostRequestBodyPrimitiveExample`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRequestBodyPrimitiveExampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutParameterExamplesParam1Param2

> User PutParameterExamplesParam1Param2(ctx, param1, param2).User(user).Execute()

Within `example`



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
	param1 := "param1-example" // string | 
	param2 := "param2-example" // string | 
	user := *openapiclient.NewUser() // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SingleExampleAPI.PutParameterExamplesParam1Param2(context.Background(), param1, param2).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SingleExampleAPI.PutParameterExamplesParam1Param2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutParameterExamplesParam1Param2`: User
	fmt.Fprintf(os.Stdout, "Response from `SingleExampleAPI.PutParameterExamplesParam1Param2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**param1** | **string** |  | 
**param2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutParameterExamplesParam1Param2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRequestBody

> User PutRequestBody(ctx).User(user).Execute()

`$ref` within `example`



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
	user := *openapiclient.NewUser() // User | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SingleExampleAPI.PutRequestBody(context.Background()).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SingleExampleAPI.PutRequestBody``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRequestBody`: User
	fmt.Fprintf(os.Stdout, "Response from `SingleExampleAPI.PutRequestBody`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutRequestBodyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

