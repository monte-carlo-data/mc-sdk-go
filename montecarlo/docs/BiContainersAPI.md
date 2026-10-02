# \BiContainersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBiContainer**](BiContainersAPI.md#CreateBiContainer) | **Post** /api/v2/bi-containers | Create a BI container
[**DeleteBiContainer**](BiContainersAPI.md#DeleteBiContainer) | **Delete** /api/v2/bi-containers/{bi_container_id} | Delete a BI container
[**GetBiContainer**](BiContainersAPI.md#GetBiContainer) | **Get** /api/v2/bi-containers/{bi_container_id} | Get a BI container
[**ListBiContainers**](BiContainersAPI.md#ListBiContainers) | **Get** /api/v2/bi-containers | List BI containers
[**UpdateBiContainer**](BiContainersAPI.md#UpdateBiContainer) | **Patch** /api/v2/bi-containers/{bi_container_id} | Update a BI container



## CreateBiContainer

> BiContainerOut CreateBiContainer(ctx).BiContainerIn(biContainerIn).Execute()

Create a BI container



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
	biContainerIn := *openapiclient.NewBiContainerIn(openapiclient.NewBiContainerType("looker"), "Name_example", "DeploymentId_example") // BiContainerIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiContainersAPI.CreateBiContainer(context.Background()).BiContainerIn(biContainerIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiContainersAPI.CreateBiContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBiContainer`: BiContainerOut
	fmt.Fprintf(os.Stdout, "Response from `BiContainersAPI.CreateBiContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBiContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **biContainerIn** | [**BiContainerIn**](BiContainerIn.md) |  | 

### Return type

[**BiContainerOut**](BiContainerOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBiContainer

> DeleteBiContainer(ctx, biContainerId).Execute()

Delete a BI container



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
	biContainerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BiContainersAPI.DeleteBiContainer(context.Background(), biContainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiContainersAPI.DeleteBiContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**biContainerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBiContainerRequest struct via the builder pattern


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


## GetBiContainer

> BiContainerOut GetBiContainer(ctx, biContainerId).Execute()

Get a BI container



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
	biContainerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiContainersAPI.GetBiContainer(context.Background(), biContainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiContainersAPI.GetBiContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBiContainer`: BiContainerOut
	fmt.Fprintf(os.Stdout, "Response from `BiContainersAPI.GetBiContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**biContainerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBiContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BiContainerOut**](BiContainerOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBiContainers

> []BiContainerOut ListBiContainers(ctx).Execute()

List BI containers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiContainersAPI.ListBiContainers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiContainersAPI.ListBiContainers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBiContainers`: []BiContainerOut
	fmt.Fprintf(os.Stdout, "Response from `BiContainersAPI.ListBiContainers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBiContainersRequest struct via the builder pattern


### Return type

[**[]BiContainerOut**](BiContainerOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBiContainer

> BiContainerOut UpdateBiContainer(ctx, biContainerId).BiContainerPatch(biContainerPatch).Execute()

Update a BI container



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
	biContainerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	biContainerPatch := *openapiclient.NewBiContainerPatch() // BiContainerPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BiContainersAPI.UpdateBiContainer(context.Background(), biContainerId).BiContainerPatch(biContainerPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BiContainersAPI.UpdateBiContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBiContainer`: BiContainerOut
	fmt.Fprintf(os.Stdout, "Response from `BiContainersAPI.UpdateBiContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**biContainerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBiContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **biContainerPatch** | [**BiContainerPatch**](BiContainerPatch.md) |  | 

### Return type

[**BiContainerOut**](BiContainerOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

