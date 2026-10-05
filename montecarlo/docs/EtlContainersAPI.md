# \EtlContainersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEtlContainer**](EtlContainersAPI.md#CreateEtlContainer) | **Post** /api/v2/etl-containers | Create an ETL container
[**DeleteEtlContainer**](EtlContainersAPI.md#DeleteEtlContainer) | **Delete** /api/v2/etl-containers/{etl_container_id} | Delete an ETL container
[**GetEtlContainer**](EtlContainersAPI.md#GetEtlContainer) | **Get** /api/v2/etl-containers/{etl_container_id} | Get an ETL container
[**ListEtlContainers**](EtlContainersAPI.md#ListEtlContainers) | **Get** /api/v2/etl-containers | List ETL containers
[**UpdateEtlContainer**](EtlContainersAPI.md#UpdateEtlContainer) | **Patch** /api/v2/etl-containers/{etl_container_id} | Update an ETL container



## CreateEtlContainer

> EtlContainerOut CreateEtlContainer(ctx).EtlContainerIn(etlContainerIn).Execute()

Create an ETL container



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
	etlContainerIn := *openapiclient.NewEtlContainerIn(openapiclient.NewEtlContainerType("airflow"), "Name_example") // EtlContainerIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtlContainersAPI.CreateEtlContainer(context.Background()).EtlContainerIn(etlContainerIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlContainersAPI.CreateEtlContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEtlContainer`: EtlContainerOut
	fmt.Fprintf(os.Stdout, "Response from `EtlContainersAPI.CreateEtlContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEtlContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **etlContainerIn** | [**EtlContainerIn**](EtlContainerIn.md) |  | 

### Return type

[**EtlContainerOut**](EtlContainerOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEtlContainer

> DeleteEtlContainer(ctx, etlContainerId).Execute()

Delete an ETL container



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
	etlContainerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EtlContainersAPI.DeleteEtlContainer(context.Background(), etlContainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlContainersAPI.DeleteEtlContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**etlContainerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEtlContainerRequest struct via the builder pattern


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


## GetEtlContainer

> EtlContainerOut GetEtlContainer(ctx, etlContainerId).Execute()

Get an ETL container



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
	etlContainerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtlContainersAPI.GetEtlContainer(context.Background(), etlContainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlContainersAPI.GetEtlContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEtlContainer`: EtlContainerOut
	fmt.Fprintf(os.Stdout, "Response from `EtlContainersAPI.GetEtlContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**etlContainerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEtlContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EtlContainerOut**](EtlContainerOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEtlContainers

> []EtlContainerOut ListEtlContainers(ctx).Execute()

List ETL containers



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
	resp, r, err := apiClient.EtlContainersAPI.ListEtlContainers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlContainersAPI.ListEtlContainers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEtlContainers`: []EtlContainerOut
	fmt.Fprintf(os.Stdout, "Response from `EtlContainersAPI.ListEtlContainers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEtlContainersRequest struct via the builder pattern


### Return type

[**[]EtlContainerOut**](EtlContainerOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEtlContainer

> EtlContainerOut UpdateEtlContainer(ctx, etlContainerId).EtlContainerPatch(etlContainerPatch).Execute()

Update an ETL container



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
	etlContainerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	etlContainerPatch := *openapiclient.NewEtlContainerPatch() // EtlContainerPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EtlContainersAPI.UpdateEtlContainer(context.Background(), etlContainerId).EtlContainerPatch(etlContainerPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlContainersAPI.UpdateEtlContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEtlContainer`: EtlContainerOut
	fmt.Fprintf(os.Stdout, "Response from `EtlContainersAPI.UpdateEtlContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**etlContainerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEtlContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **etlContainerPatch** | [**EtlContainerPatch**](EtlContainerPatch.md) |  | 

### Return type

[**EtlContainerOut**](EtlContainerOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

