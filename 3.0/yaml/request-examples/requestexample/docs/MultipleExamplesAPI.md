# \MultipleExamplesAPI

All URIs are relative to *https://httpbin.org/anything*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetParameterExamplesParam1Param2**](MultipleExamplesAPI.md#GetParameterExamplesParam1Param2) | **Get** /parameterExamples/{param1}/{param2} | Within &#x60;examples&#x60; (parameters)
[**PatchParameterExamplesParam1Param2**](MultipleExamplesAPI.md#PatchParameterExamplesParam1Param2) | **Patch** /parameterExamples/{param1}/{param2} | Within &#x60;examples&#x60; (mixed)
[**PatchRequestBodyPrimitiveExample**](MultipleExamplesAPI.md#PatchRequestBodyPrimitiveExample) | **Patch** /requestBody-primitive-example | Stringified JSON object in an &#x60;examples&#x60; value
[**PostRequestBody**](MultipleExamplesAPI.md#PostRequestBody) | **Post** /requestBody | Within &#x60;examples&#x60; (body)
[**PutRequestBodyPrimitiveExample**](MultipleExamplesAPI.md#PutRequestBodyPrimitiveExample) | **Put** /requestBody-primitive-example | Stringified JSON arrays in an &#x60;examples&#x60; value



## GetParameterExamplesParam1Param2

> User GetParameterExamplesParam1Param2(ctx, param1, param2).Execute()

Within `examples` (parameters)



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipleExamplesAPI.GetParameterExamplesParam1Param2(context.Background(), param1, param2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipleExamplesAPI.GetParameterExamplesParam1Param2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParameterExamplesParam1Param2`: User
	fmt.Fprintf(os.Stdout, "Response from `MultipleExamplesAPI.GetParameterExamplesParam1Param2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**param1** | **string** |  | 
**param2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParameterExamplesParam1Param2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchParameterExamplesParam1Param2

> User PatchParameterExamplesParam1Param2(ctx, param1, param2).Param5(param5).User(user).Param3(param3).Param4(param4).Param6(param6).Execute()

Within `examples` (mixed)



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
	param5 := "param5-example" // string | 
	user := *openapiclient.NewUser() // User | 
	param3 := "param3-example" // string |  (optional)
	param4 := "param4-example" // string |  (optional)
	param6 := "param6-example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipleExamplesAPI.PatchParameterExamplesParam1Param2(context.Background(), param1, param2).Param5(param5).User(user).Param3(param3).Param4(param4).Param6(param6).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipleExamplesAPI.PatchParameterExamplesParam1Param2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchParameterExamplesParam1Param2`: User
	fmt.Fprintf(os.Stdout, "Response from `MultipleExamplesAPI.PatchParameterExamplesParam1Param2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**param1** | **string** |  | 
**param2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchParameterExamplesParam1Param2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **param5** | **string** |  | 
 **user** | [**User**](User.md) |  | 
 **param3** | **string** |  | 
 **param4** | **string** |  | 
 **param6** | **string** |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRequestBodyPrimitiveExample

> Pet PatchRequestBodyPrimitiveExample(ctx).Pet(pet).Execute()

Stringified JSON object in an `examples` value



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
	pet := *openapiclient.NewPet() // Pet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipleExamplesAPI.PatchRequestBodyPrimitiveExample(context.Background()).Pet(pet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipleExamplesAPI.PatchRequestBodyPrimitiveExample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRequestBodyPrimitiveExample`: Pet
	fmt.Fprintf(os.Stdout, "Response from `MultipleExamplesAPI.PatchRequestBodyPrimitiveExample`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchRequestBodyPrimitiveExampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pet** | [**Pet**](Pet.md) |  | 

### Return type

[**Pet**](Pet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRequestBody

> User PostRequestBody(ctx).User(user).Execute()

Within `examples` (body)



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
	resp, r, err := apiClient.MultipleExamplesAPI.PostRequestBody(context.Background()).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipleExamplesAPI.PostRequestBody``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRequestBody`: User
	fmt.Fprintf(os.Stdout, "Response from `MultipleExamplesAPI.PostRequestBody`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRequestBodyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRequestBodyPrimitiveExample

> []Pet PutRequestBodyPrimitiveExample(ctx).Pet(pet).Execute()

Stringified JSON arrays in an `examples` value



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
	pet := []openapiclient.Pet{*openapiclient.NewPet()} // []Pet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipleExamplesAPI.PutRequestBodyPrimitiveExample(context.Background()).Pet(pet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipleExamplesAPI.PutRequestBodyPrimitiveExample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRequestBodyPrimitiveExample`: []Pet
	fmt.Fprintf(os.Stdout, "Response from `MultipleExamplesAPI.PutRequestBodyPrimitiveExample`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutRequestBodyPrimitiveExampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pet** | [**[]Pet**](Pet.md) |  | 

### Return type

[**[]Pet**](Pet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

