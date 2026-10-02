# \ConnectionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnection**](ConnectionsAPI.md#CreateConnection) | **Post** /api/v2/connections | Create a connection
[**DeleteConnection**](ConnectionsAPI.md#DeleteConnection) | **Delete** /api/v2/connections/{connection_id} | Delete a connection
[**GetConnection**](ConnectionsAPI.md#GetConnection) | **Get** /api/v2/connections/{connection_id} | Get a connection
[**ListConnections**](ConnectionsAPI.md#ListConnections) | **Get** /api/v2/connections | List connections
[**UpdateConnection**](ConnectionsAPI.md#UpdateConnection) | **Patch** /api/v2/connections/{connection_id} | Update a connection
[**ValidateConnection**](ConnectionsAPI.md#ValidateConnection) | **Post** /api/v2/connections/{connection_id}/validate | Validate a connection



## CreateConnection

> ConnectionOut CreateConnection(ctx).ConnectionIn(connectionIn).Execute()

Create a connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
)

func main() {
	connectionIn := *openapiclient.NewConnectionIn("Name_example", "CredentialsId_example") // ConnectionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.CreateConnection(context.Background()).ConnectionIn(connectionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.CreateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnection`: ConnectionOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.CreateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionIn** | [**ConnectionIn**](ConnectionIn.md) |  | 

### Return type

[**ConnectionOut**](ConnectionOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnection

> DeleteConnection(ctx, connectionId).Execute()

Delete a connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
)

func main() {
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectionsAPI.DeleteConnection(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.DeleteConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnection

> ConnectionOut GetConnection(ctx, connectionId).Execute()

Get a connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
)

func main() {
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.GetConnection(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.GetConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnection`: ConnectionOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.GetConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectionOut**](ConnectionOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnections

> PagedConnectionOut ListConnections(ctx).WarehouseId(warehouseId).BiContainerId(biContainerId).Cursor(cursor).Limit(limit).WithCount(withCount).Execute()

List connections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
)

func main() {
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only list connections on this warehouse. Omit it to list every connection in your account. (optional)
	biContainerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only list connections on this BI container. Omit it to list every connection in your account. (optional)
	cursor := "cursor_example" // string | Position to continue from, as returned in `next_cursor` by the previous page. Omit it to start from the first page. The value is opaque; do not build or modify one. (optional)
	limit := int32(56) // int32 | Maximum number of items to return, between 1 and 100. (optional) (default to 50)
	withCount := true // bool | Whether to also return the total number of items across every page, in `count`. Off by default: counting costs an extra query. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ListConnections(context.Background()).WarehouseId(warehouseId).BiContainerId(biContainerId).Cursor(cursor).Limit(limit).WithCount(withCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ListConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnections`: PagedConnectionOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ListConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **warehouseId** | **string** | Only list connections on this warehouse. Omit it to list every connection in your account. | 
 **biContainerId** | **string** | Only list connections on this BI container. Omit it to list every connection in your account. | 
 **cursor** | **string** | Position to continue from, as returned in &#x60;next_cursor&#x60; by the previous page. Omit it to start from the first page. The value is opaque; do not build or modify one. | 
 **limit** | **int32** | Maximum number of items to return, between 1 and 100. | [default to 50]
 **withCount** | **bool** | Whether to also return the total number of items across every page, in &#x60;count&#x60;. Off by default: counting costs an extra query. | [default to false]

### Return type

[**PagedConnectionOut**](PagedConnectionOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnection

> ConnectionOut UpdateConnection(ctx, connectionId).ConnectionPatch(connectionPatch).Execute()

Update a connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
)

func main() {
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	connectionPatch := *openapiclient.NewConnectionPatch() // ConnectionPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.UpdateConnection(context.Background(), connectionId).ConnectionPatch(connectionPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.UpdateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConnection`: ConnectionOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.UpdateConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionPatch** | [**ConnectionPatch**](ConnectionPatch.md) |  | 

### Return type

[**ConnectionOut**](ConnectionOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConnection

> ValidationRunOut ValidateConnection(ctx, connectionId).Execute()

Validate a connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
)

func main() {
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectionsAPI.ValidateConnection(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionsAPI.ValidateConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateConnection`: ValidationRunOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectionsAPI.ValidateConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ValidationRunOut**](ValidationRunOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

