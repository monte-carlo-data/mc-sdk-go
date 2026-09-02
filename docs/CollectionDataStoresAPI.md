# \CollectionDataStoresAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAwsCollectionDataStore**](CollectionDataStoresAPI.md#DeleteAwsCollectionDataStore) | **Delete** /api/v2/collection-data-stores/aws/{collection_data_store_id} | Delete an AWS collection data store
[**GetAwsCollectionDataStore**](CollectionDataStoresAPI.md#GetAwsCollectionDataStore) | **Get** /api/v2/collection-data-stores/aws/{collection_data_store_id} | Get an AWS collection data store
[**ListCollectionDataStores**](CollectionDataStoresAPI.md#ListCollectionDataStores) | **Get** /api/v2/collection-data-stores | List collection data stores
[**RegisterAwsCollectionDataStore**](CollectionDataStoresAPI.md#RegisterAwsCollectionDataStore) | **Post** /api/v2/collection-data-stores/aws | Register an AWS collection data store
[**UpdateAwsCollectionDataStore**](CollectionDataStoresAPI.md#UpdateAwsCollectionDataStore) | **Patch** /api/v2/collection-data-stores/aws/{collection_data_store_id} | Update an AWS collection data store



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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go"
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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go"
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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go"
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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go"
)

func main() {
	awsCollectionDataStoreIn := *openapiclient.NewAwsCollectionDataStoreIn("BucketName_example", "DeploymentId_example", "RoleArn_example") // AwsCollectionDataStoreIn | 

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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go"
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

