# \FormDataAPI

All URIs are relative to *https://httpbin.org/anything*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostRequestBodyFormDataExample**](FormDataAPI.md#PostRequestBodyFormDataExample) | **Post** /requestBody-form-data-example | Demo handling of formData



## PostRequestBodyFormDataExample

> Token PostRequestBodyFormDataExample(ctx).ClientId(clientId).ClientSecret(clientSecret).Scope(scope).Execute()

Demo handling of formData



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
	clientId := "clientId_example" // string | 
	clientSecret := "clientSecret_example" // string | 
	scope := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormDataAPI.PostRequestBodyFormDataExample(context.Background()).ClientId(clientId).ClientSecret(clientSecret).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormDataAPI.PostRequestBodyFormDataExample``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRequestBodyFormDataExample`: Token
	fmt.Fprintf(os.Stdout, "Response from `FormDataAPI.PostRequestBodyFormDataExample`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRequestBodyFormDataExampleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 
 **scope** | **int32** |  | 

### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

