# \WarehousesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWarehouse**](WarehousesAPI.md#CreateWarehouse) | **Post** /api/v2/warehouses | Create a warehouse
[**DeleteWarehouse**](WarehousesAPI.md#DeleteWarehouse) | **Delete** /api/v2/warehouses/{warehouse_id} | Delete a warehouse
[**GetWarehouse**](WarehousesAPI.md#GetWarehouse) | **Get** /api/v2/warehouses/{warehouse_id} | Get a warehouse
[**ListWarehouses**](WarehousesAPI.md#ListWarehouses) | **Get** /api/v2/warehouses | List warehouses
[**UpdateWarehouse**](WarehousesAPI.md#UpdateWarehouse) | **Patch** /api/v2/warehouses/{warehouse_id} | Update a warehouse



## CreateWarehouse

> WarehouseOut CreateWarehouse(ctx).WarehouseIn(warehouseIn).Execute()

Create a warehouse



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
	warehouseIn := *openapiclient.NewWarehouseIn("Name_example", openapiclient.WarehouseType("bigquery"), "DeploymentId_example") // WarehouseIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehousesAPI.CreateWarehouse(context.Background()).WarehouseIn(warehouseIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.CreateWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWarehouse`: WarehouseOut
	fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.CreateWarehouse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **warehouseIn** | [**WarehouseIn**](WarehouseIn.md) |  | 

### Return type

[**WarehouseOut**](WarehouseOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWarehouse

> DeleteWarehouse(ctx, warehouseId).Execute()

Delete a warehouse



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
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WarehousesAPI.DeleteWarehouse(context.Background(), warehouseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.DeleteWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWarehouseRequest struct via the builder pattern


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


## GetWarehouse

> WarehouseOut GetWarehouse(ctx, warehouseId).Execute()

Get a warehouse



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
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehousesAPI.GetWarehouse(context.Background(), warehouseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.GetWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarehouse`: WarehouseOut
	fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.GetWarehouse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WarehouseOut**](WarehouseOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWarehouses

> PagedWarehouseOut ListWarehouses(ctx).Cursor(cursor).Limit(limit).WithCount(withCount).Execute()

List warehouses



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
	cursor := "cursor_example" // string | Position to continue from, as returned in `next_cursor` by the previous page. Omit it to start from the first page. The value is opaque; do not build or modify one. (optional)
	limit := int32(56) // int32 | Maximum number of items to return, between 1 and 100. (optional) (default to 50)
	withCount := true // bool | Whether to also return the total number of items across every page, in `count`. Off by default: counting costs an extra query. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehousesAPI.ListWarehouses(context.Background()).Cursor(cursor).Limit(limit).WithCount(withCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.ListWarehouses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWarehouses`: PagedWarehouseOut
	fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.ListWarehouses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWarehousesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Position to continue from, as returned in &#x60;next_cursor&#x60; by the previous page. Omit it to start from the first page. The value is opaque; do not build or modify one. | 
 **limit** | **int32** | Maximum number of items to return, between 1 and 100. | [default to 50]
 **withCount** | **bool** | Whether to also return the total number of items across every page, in &#x60;count&#x60;. Off by default: counting costs an extra query. | [default to false]

### Return type

[**PagedWarehouseOut**](PagedWarehouseOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWarehouse

> WarehouseOut UpdateWarehouse(ctx, warehouseId).WarehousePatch(warehousePatch).Execute()

Update a warehouse



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
	warehouseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warehousePatch := *openapiclient.NewWarehousePatch() // WarehousePatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarehousesAPI.UpdateWarehouse(context.Background(), warehouseId).WarehousePatch(warehousePatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarehousesAPI.UpdateWarehouse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWarehouse`: WarehouseOut
	fmt.Fprintf(os.Stdout, "Response from `WarehousesAPI.UpdateWarehouse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWarehouseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **warehousePatch** | [**WarehousePatch**](WarehousePatch.md) |  | 

### Return type

[**WarehouseOut**](WarehouseOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

