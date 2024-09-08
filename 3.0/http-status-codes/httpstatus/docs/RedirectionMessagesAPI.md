# \RedirectionMessagesAPI

All URIs are relative to *https://httpbin.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatus300**](RedirectionMessagesAPI.md#GetStatus300) | **Get** /status/300 | Returns a \&quot;300 Multiple Choice\&quot;
[**GetStatus301**](RedirectionMessagesAPI.md#GetStatus301) | **Get** /status/301 | Returns a \&quot;301 Moved Permanently\&quot;
[**GetStatus302**](RedirectionMessagesAPI.md#GetStatus302) | **Get** /status/302 | Returns a \&quot;302 Found\&quot;
[**GetStatus303**](RedirectionMessagesAPI.md#GetStatus303) | **Get** /status/303 | Returns a \&quot;303 See Other\&quot;
[**GetStatus304**](RedirectionMessagesAPI.md#GetStatus304) | **Get** /status/304 | Returns a \&quot;304 Not Modified\&quot;
[**GetStatus305**](RedirectionMessagesAPI.md#GetStatus305) | **Get** /status/305 | Returns a \&quot;305 Use Proxy\&quot;
[**GetStatus306**](RedirectionMessagesAPI.md#GetStatus306) | **Get** /status/306 | Returns a \&quot;306 Switch Proxy\&quot;
[**GetStatus307**](RedirectionMessagesAPI.md#GetStatus307) | **Get** /status/307 | Returns a \&quot;307 Temporary Redirect\&quot;
[**GetStatus308**](RedirectionMessagesAPI.md#GetStatus308) | **Get** /status/308 | Returns a \&quot;308 Permanent Redirect\&quot;



## GetStatus300

> GetStatus300(ctx).Execute()

Returns a \"300 Multiple Choice\"



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
	r, err := apiClient.RedirectionMessagesAPI.GetStatus300(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionMessagesAPI.GetStatus300``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatus300Request struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus301

> GetStatus301(ctx).Execute()

Returns a \"301 Moved Permanently\"



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
	r, err := apiClient.RedirectionMessagesAPI.GetStatus301(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionMessagesAPI.GetStatus301``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatus301Request struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus302

> GetStatus302(ctx).Execute()

Returns a \"302 Found\"



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
	r, err := apiClient.RedirectionMessagesAPI.GetStatus302(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionMessagesAPI.GetStatus302``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatus302Request struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus303

> GetStatus303(ctx).Execute()

Returns a \"303 See Other\"



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
	r, err := apiClient.RedirectionMessagesAPI.GetStatus303(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionMessagesAPI.GetStatus303``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatus303Request struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus304

> GetStatus304(ctx).Execute()

Returns a \"304 Not Modified\"



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
	r, err := apiClient.RedirectionMessagesAPI.GetStatus304(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionMessagesAPI.GetStatus304``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatus304Request struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus305

> GetStatus305(ctx).Execute()

Returns a \"305 Use Proxy\"



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
	r, err := apiClient.RedirectionMessagesAPI.GetStatus305(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionMessagesAPI.GetStatus305``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatus305Request struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus306

> GetStatus306(ctx).Execute()

Returns a \"306 Switch Proxy\"



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
	r, err := apiClient.RedirectionMessagesAPI.GetStatus306(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionMessagesAPI.GetStatus306``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatus306Request struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus307

> GetStatus307(ctx).Execute()

Returns a \"307 Temporary Redirect\"



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
	r, err := apiClient.RedirectionMessagesAPI.GetStatus307(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionMessagesAPI.GetStatus307``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatus307Request struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus308

> GetStatus308(ctx).Execute()

Returns a \"308 Permanent Redirect\"



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
	r, err := apiClient.RedirectionMessagesAPI.GetStatus308(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RedirectionMessagesAPI.GetStatus308``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatus308Request struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

