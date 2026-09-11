# \CollectionDataStoresAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAwsCollectionDataStore**](CollectionDataStoresAPI.md#DeleteAwsCollectionDataStore) | **Delete** /api/v2/collection-data-stores/aws/{collection_data_store_id} | Delete an AWS collection data store
[**DeleteAzureCollectionDataStore**](CollectionDataStoresAPI.md#DeleteAzureCollectionDataStore) | **Delete** /api/v2/collection-data-stores/azure/{collection_data_store_id} | Delete an Azure collection data store
[**DeleteGcpCollectionDataStore**](CollectionDataStoresAPI.md#DeleteGcpCollectionDataStore) | **Delete** /api/v2/collection-data-stores/gcp/{collection_data_store_id} | Delete a GCP collection data store
[**GetAwsCollectionDataStore**](CollectionDataStoresAPI.md#GetAwsCollectionDataStore) | **Get** /api/v2/collection-data-stores/aws/{collection_data_store_id} | Get an AWS collection data store
[**GetAzureCollectionDataStore**](CollectionDataStoresAPI.md#GetAzureCollectionDataStore) | **Get** /api/v2/collection-data-stores/azure/{collection_data_store_id} | Get an Azure collection data store
[**GetGcpCollectionDataStore**](CollectionDataStoresAPI.md#GetGcpCollectionDataStore) | **Get** /api/v2/collection-data-stores/gcp/{collection_data_store_id} | Get a GCP collection data store
[**ListCollectionDataStores**](CollectionDataStoresAPI.md#ListCollectionDataStores) | **Get** /api/v2/collection-data-stores | List collection data stores
[**RegisterAwsCollectionDataStore**](CollectionDataStoresAPI.md#RegisterAwsCollectionDataStore) | **Post** /api/v2/collection-data-stores/aws | Register an AWS collection data store
[**RegisterAzureCollectionDataStore**](CollectionDataStoresAPI.md#RegisterAzureCollectionDataStore) | **Post** /api/v2/collection-data-stores/azure | Register an Azure collection data store
[**RegisterGcpCollectionDataStore**](CollectionDataStoresAPI.md#RegisterGcpCollectionDataStore) | **Post** /api/v2/collection-data-stores/gcp | Register a GCP collection data store
[**UpdateAwsCollectionDataStore**](CollectionDataStoresAPI.md#UpdateAwsCollectionDataStore) | **Patch** /api/v2/collection-data-stores/aws/{collection_data_store_id} | Update an AWS collection data store
[**UpdateAzureCollectionDataStore**](CollectionDataStoresAPI.md#UpdateAzureCollectionDataStore) | **Patch** /api/v2/collection-data-stores/azure/{collection_data_store_id} | Update an Azure collection data store
[**UpdateGcpCollectionDataStore**](CollectionDataStoresAPI.md#UpdateGcpCollectionDataStore) | **Patch** /api/v2/collection-data-stores/gcp/{collection_data_store_id} | Update a GCP collection data store



## DeleteAwsCollectionDataStore

> DeleteAwsCollectionDataStore(ctx, collectionDataStoreId).Execute()

Delete an AWS collection data store



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
	collectionDataStoreId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionDataStoresAPI.DeleteAwsCollectionDataStore(context.Background(), collectionDataStoreId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.DeleteAwsCollectionDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionDataStoreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAwsCollectionDataStoreRequest struct via the builder pattern


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


## DeleteAzureCollectionDataStore

> DeleteAzureCollectionDataStore(ctx, collectionDataStoreId).Execute()

Delete an Azure collection data store



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
	collectionDataStoreId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionDataStoresAPI.DeleteAzureCollectionDataStore(context.Background(), collectionDataStoreId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.DeleteAzureCollectionDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionDataStoreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAzureCollectionDataStoreRequest struct via the builder pattern


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


## DeleteGcpCollectionDataStore

> DeleteGcpCollectionDataStore(ctx, collectionDataStoreId).Execute()

Delete a GCP collection data store



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
	collectionDataStoreId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionDataStoresAPI.DeleteGcpCollectionDataStore(context.Background(), collectionDataStoreId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.DeleteGcpCollectionDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionDataStoreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGcpCollectionDataStoreRequest struct via the builder pattern


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


## GetAwsCollectionDataStore

> AwsCollectionDataStoreOut GetAwsCollectionDataStore(ctx, collectionDataStoreId).Execute()

Get an AWS collection data store



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
	collectionDataStoreId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionDataStoresAPI.GetAwsCollectionDataStore(context.Background(), collectionDataStoreId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.GetAwsCollectionDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAwsCollectionDataStore`: AwsCollectionDataStoreOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionDataStoresAPI.GetAwsCollectionDataStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionDataStoreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsCollectionDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwsCollectionDataStoreOut**](AwsCollectionDataStoreOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAzureCollectionDataStore

> AzureCollectionDataStoreOut GetAzureCollectionDataStore(ctx, collectionDataStoreId).Execute()

Get an Azure collection data store



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
	collectionDataStoreId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionDataStoresAPI.GetAzureCollectionDataStore(context.Background(), collectionDataStoreId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.GetAzureCollectionDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAzureCollectionDataStore`: AzureCollectionDataStoreOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionDataStoresAPI.GetAzureCollectionDataStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionDataStoreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureCollectionDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AzureCollectionDataStoreOut**](AzureCollectionDataStoreOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpCollectionDataStore

> GcpCollectionDataStoreOut GetGcpCollectionDataStore(ctx, collectionDataStoreId).Execute()

Get a GCP collection data store



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
	collectionDataStoreId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionDataStoresAPI.GetGcpCollectionDataStore(context.Background(), collectionDataStoreId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.GetGcpCollectionDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGcpCollectionDataStore`: GcpCollectionDataStoreOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionDataStoresAPI.GetGcpCollectionDataStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionDataStoreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpCollectionDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GcpCollectionDataStoreOut**](GcpCollectionDataStoreOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionDataStores

> []CollectionDataStoreOut ListCollectionDataStores(ctx).Execute()

List collection data stores



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
	resp, r, err := apiClient.CollectionDataStoresAPI.ListCollectionDataStores(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.ListCollectionDataStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollectionDataStores`: []CollectionDataStoreOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionDataStoresAPI.ListCollectionDataStores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionDataStoresRequest struct via the builder pattern


### Return type

[**[]CollectionDataStoreOut**](CollectionDataStoreOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAwsCollectionDataStore

> AwsCollectionDataStoreOut RegisterAwsCollectionDataStore(ctx).AwsCollectionDataStoreIn(awsCollectionDataStoreIn).Execute()

Register an AWS collection data store



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
	awsCollectionDataStoreIn := *openapiclient.NewAwsCollectionDataStoreIn("DeploymentId_example", "BucketName_example", "RoleArn_example") // AwsCollectionDataStoreIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionDataStoresAPI.RegisterAwsCollectionDataStore(context.Background()).AwsCollectionDataStoreIn(awsCollectionDataStoreIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.RegisterAwsCollectionDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAwsCollectionDataStore`: AwsCollectionDataStoreOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionDataStoresAPI.RegisterAwsCollectionDataStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAwsCollectionDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsCollectionDataStoreIn** | [**AwsCollectionDataStoreIn**](AwsCollectionDataStoreIn.md) |  | 

### Return type

[**AwsCollectionDataStoreOut**](AwsCollectionDataStoreOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAzureCollectionDataStore

> AzureCollectionDataStoreOut RegisterAzureCollectionDataStore(ctx).AzureCollectionDataStoreIn(azureCollectionDataStoreIn).Execute()

Register an Azure collection data store



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
	azureCollectionDataStoreIn := *openapiclient.NewAzureCollectionDataStoreIn(openapiclient.AzureDataStoreAuthenticationType("AZURE_STORAGE_ACCOUNT_KEYS"), "DeploymentId_example", "ContainerName_example") // AzureCollectionDataStoreIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionDataStoresAPI.RegisterAzureCollectionDataStore(context.Background()).AzureCollectionDataStoreIn(azureCollectionDataStoreIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.RegisterAzureCollectionDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAzureCollectionDataStore`: AzureCollectionDataStoreOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionDataStoresAPI.RegisterAzureCollectionDataStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAzureCollectionDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureCollectionDataStoreIn** | [**AzureCollectionDataStoreIn**](AzureCollectionDataStoreIn.md) |  | 

### Return type

[**AzureCollectionDataStoreOut**](AzureCollectionDataStoreOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterGcpCollectionDataStore

> GcpCollectionDataStoreOut RegisterGcpCollectionDataStore(ctx).GcpCollectionDataStoreIn(gcpCollectionDataStoreIn).Execute()

Register a GCP collection data store



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
	gcpCollectionDataStoreIn := *openapiclient.NewGcpCollectionDataStoreIn("DeploymentId_example", "BucketName_example", "ServiceAccountKey_example") // GcpCollectionDataStoreIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionDataStoresAPI.RegisterGcpCollectionDataStore(context.Background()).GcpCollectionDataStoreIn(gcpCollectionDataStoreIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.RegisterGcpCollectionDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterGcpCollectionDataStore`: GcpCollectionDataStoreOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionDataStoresAPI.RegisterGcpCollectionDataStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterGcpCollectionDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gcpCollectionDataStoreIn** | [**GcpCollectionDataStoreIn**](GcpCollectionDataStoreIn.md) |  | 

### Return type

[**GcpCollectionDataStoreOut**](GcpCollectionDataStoreOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAwsCollectionDataStore

> AwsCollectionDataStoreOut UpdateAwsCollectionDataStore(ctx, collectionDataStoreId).AwsCollectionDataStorePatch(awsCollectionDataStorePatch).Execute()

Update an AWS collection data store



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
	collectionDataStoreId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	awsCollectionDataStorePatch := *openapiclient.NewAwsCollectionDataStorePatch() // AwsCollectionDataStorePatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionDataStoresAPI.UpdateAwsCollectionDataStore(context.Background(), collectionDataStoreId).AwsCollectionDataStorePatch(awsCollectionDataStorePatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.UpdateAwsCollectionDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAwsCollectionDataStore`: AwsCollectionDataStoreOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionDataStoresAPI.UpdateAwsCollectionDataStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionDataStoreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAwsCollectionDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsCollectionDataStorePatch** | [**AwsCollectionDataStorePatch**](AwsCollectionDataStorePatch.md) |  | 

### Return type

[**AwsCollectionDataStoreOut**](AwsCollectionDataStoreOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAzureCollectionDataStore

> AzureCollectionDataStoreOut UpdateAzureCollectionDataStore(ctx, collectionDataStoreId).AzureCollectionDataStorePatch(azureCollectionDataStorePatch).Execute()

Update an Azure collection data store



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
	collectionDataStoreId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	azureCollectionDataStorePatch := *openapiclient.NewAzureCollectionDataStorePatch() // AzureCollectionDataStorePatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionDataStoresAPI.UpdateAzureCollectionDataStore(context.Background(), collectionDataStoreId).AzureCollectionDataStorePatch(azureCollectionDataStorePatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.UpdateAzureCollectionDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAzureCollectionDataStore`: AzureCollectionDataStoreOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionDataStoresAPI.UpdateAzureCollectionDataStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionDataStoreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAzureCollectionDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureCollectionDataStorePatch** | [**AzureCollectionDataStorePatch**](AzureCollectionDataStorePatch.md) |  | 

### Return type

[**AzureCollectionDataStoreOut**](AzureCollectionDataStoreOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGcpCollectionDataStore

> GcpCollectionDataStoreOut UpdateGcpCollectionDataStore(ctx, collectionDataStoreId).GcpCollectionDataStorePatch(gcpCollectionDataStorePatch).Execute()

Update a GCP collection data store



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
	collectionDataStoreId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	gcpCollectionDataStorePatch := *openapiclient.NewGcpCollectionDataStorePatch() // GcpCollectionDataStorePatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionDataStoresAPI.UpdateGcpCollectionDataStore(context.Background(), collectionDataStoreId).GcpCollectionDataStorePatch(gcpCollectionDataStorePatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionDataStoresAPI.UpdateGcpCollectionDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGcpCollectionDataStore`: GcpCollectionDataStoreOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionDataStoresAPI.UpdateGcpCollectionDataStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionDataStoreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGcpCollectionDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpCollectionDataStorePatch** | [**GcpCollectionDataStorePatch**](GcpCollectionDataStorePatch.md) |  | 

### Return type

[**GcpCollectionDataStoreOut**](GcpCollectionDataStoreOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

