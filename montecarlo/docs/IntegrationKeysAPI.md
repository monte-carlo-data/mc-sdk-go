# \IntegrationKeysAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAirflowCallbacksKey**](IntegrationKeysAPI.md#CreateAirflowCallbacksKey) | **Post** /api/v2/integration-keys/airflow-callbacks | Create an Airflow callbacks key
[**CreateIngestionKey**](IntegrationKeysAPI.md#CreateIngestionKey) | **Post** /api/v2/integration-keys/ingestion | Create an ingestion key
[**DeleteAirflowCallbacksKey**](IntegrationKeysAPI.md#DeleteAirflowCallbacksKey) | **Delete** /api/v2/integration-keys/airflow-callbacks/{key_id} | Delete an Airflow callbacks key
[**DeleteIngestionKey**](IntegrationKeysAPI.md#DeleteIngestionKey) | **Delete** /api/v2/integration-keys/ingestion/{key_id} | Delete an ingestion key
[**GetAirflowCallbacksKey**](IntegrationKeysAPI.md#GetAirflowCallbacksKey) | **Get** /api/v2/integration-keys/airflow-callbacks/{key_id} | Get an Airflow callbacks key
[**GetIngestionKey**](IntegrationKeysAPI.md#GetIngestionKey) | **Get** /api/v2/integration-keys/ingestion/{key_id} | Get an ingestion key
[**ListIntegrationKeys**](IntegrationKeysAPI.md#ListIntegrationKeys) | **Get** /api/v2/integration-keys | List integration keys



## CreateAirflowCallbacksKey

> AirflowCallbacksKeyCreatedOut CreateAirflowCallbacksKey(ctx).AirflowCallbacksKeyIn(airflowCallbacksKeyIn).Execute()

Create an Airflow callbacks key



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
	airflowCallbacksKeyIn := *openapiclient.NewAirflowCallbacksKeyIn("EtlContainerId_example") // AirflowCallbacksKeyIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKeysAPI.CreateAirflowCallbacksKey(context.Background()).AirflowCallbacksKeyIn(airflowCallbacksKeyIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKeysAPI.CreateAirflowCallbacksKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAirflowCallbacksKey`: AirflowCallbacksKeyCreatedOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKeysAPI.CreateAirflowCallbacksKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAirflowCallbacksKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **airflowCallbacksKeyIn** | [**AirflowCallbacksKeyIn**](AirflowCallbacksKeyIn.md) |  | 

### Return type

[**AirflowCallbacksKeyCreatedOut**](AirflowCallbacksKeyCreatedOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIngestionKey

> IngestionKeyCreatedOut CreateIngestionKey(ctx).IngestionKeyIn(ingestionKeyIn).Execute()

Create an ingestion key



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
	ingestionKeyIn := *openapiclient.NewIngestionKeyIn() // IngestionKeyIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKeysAPI.CreateIngestionKey(context.Background()).IngestionKeyIn(ingestionKeyIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKeysAPI.CreateIngestionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIngestionKey`: IngestionKeyCreatedOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKeysAPI.CreateIngestionKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIngestionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestionKeyIn** | [**IngestionKeyIn**](IngestionKeyIn.md) |  | 

### Return type

[**IngestionKeyCreatedOut**](IngestionKeyCreatedOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAirflowCallbacksKey

> DeleteAirflowCallbacksKey(ctx, keyId).Execute()

Delete an Airflow callbacks key



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
	keyId := "keyId_example" // string | Id of the integration key, as returned when it is created or listed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationKeysAPI.DeleteAirflowCallbacksKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKeysAPI.DeleteAirflowCallbacksKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | Id of the integration key, as returned when it is created or listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAirflowCallbacksKeyRequest struct via the builder pattern


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


## DeleteIngestionKey

> DeleteIngestionKey(ctx, keyId).Execute()

Delete an ingestion key



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
	keyId := "keyId_example" // string | Id of the integration key, as returned when it is created or listed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationKeysAPI.DeleteIngestionKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKeysAPI.DeleteIngestionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | Id of the integration key, as returned when it is created or listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngestionKeyRequest struct via the builder pattern


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


## GetAirflowCallbacksKey

> AirflowCallbacksKeyOut GetAirflowCallbacksKey(ctx, keyId).Execute()

Get an Airflow callbacks key



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
	keyId := "keyId_example" // string | Id of the integration key, as returned when it is created or listed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKeysAPI.GetAirflowCallbacksKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKeysAPI.GetAirflowCallbacksKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAirflowCallbacksKey`: AirflowCallbacksKeyOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKeysAPI.GetAirflowCallbacksKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | Id of the integration key, as returned when it is created or listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAirflowCallbacksKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AirflowCallbacksKeyOut**](AirflowCallbacksKeyOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngestionKey

> IngestionKeyOut GetIngestionKey(ctx, keyId).Execute()

Get an ingestion key



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
	keyId := "keyId_example" // string | Id of the integration key, as returned when it is created or listed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKeysAPI.GetIngestionKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKeysAPI.GetIngestionKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngestionKey`: IngestionKeyOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKeysAPI.GetIngestionKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | Id of the integration key, as returned when it is created or listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngestionKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IngestionKeyOut**](IngestionKeyOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationKeys

> []IntegrationKeyOut ListIntegrationKeys(ctx).Scope(scope).ParentId(parentId).Execute()

List integration keys



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
	scope := openapiclient.IntegrationKeyScope("INGESTION") // IntegrationKeyScope | Only keys of this scope. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only keys scoped to this warehouse, BI container or ETL container. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationKeysAPI.ListIntegrationKeys(context.Background()).Scope(scope).ParentId(parentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationKeysAPI.ListIntegrationKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationKeys`: []IntegrationKeyOut
	fmt.Fprintf(os.Stdout, "Response from `IntegrationKeysAPI.ListIntegrationKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | [**IntegrationKeyScope**](IntegrationKeyScope.md) | Only keys of this scope. | 
 **parentId** | **string** | Only keys scoped to this warehouse, BI container or ETL container. | 

### Return type

[**[]IntegrationKeyOut**](IntegrationKeyOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

