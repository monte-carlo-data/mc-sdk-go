# \CustomConnectorTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomConnectorType**](CustomConnectorTypesAPI.md#GetCustomConnectorType) | **Get** /api/v2/custom-connector-types/{custom_connector_type_id} | Get a custom connector type
[**ListCustomConnectorTypes**](CustomConnectorTypesAPI.md#ListCustomConnectorTypes) | **Get** /api/v2/custom-connector-types | List custom connector types



## GetCustomConnectorType

> CustomConnectorTypeOut GetCustomConnectorType(ctx, customConnectorTypeId).Execute()

Get a custom connector type



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
	customConnectorTypeId := "customConnectorTypeId_example" // string | Id of the custom connector type, as returned when it is listed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomConnectorTypesAPI.GetCustomConnectorType(context.Background(), customConnectorTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorTypesAPI.GetCustomConnectorType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomConnectorType`: CustomConnectorTypeOut
	fmt.Fprintf(os.Stdout, "Response from `CustomConnectorTypesAPI.GetCustomConnectorType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customConnectorTypeId** | **string** | Id of the custom connector type, as returned when it is listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomConnectorTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomConnectorTypeOut**](CustomConnectorTypeOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomConnectorTypes

> []CustomConnectorTypeOut ListCustomConnectorTypes(ctx).AssetClass(assetClass).CollectionAgentId(collectionAgentId).Execute()

List custom connector types



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
	assetClass := openapiclient.CustomConnectorAssetClass("warehouse") // CustomConnectorAssetClass | Only connector types of this asset class. (optional)
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only connector types this collection agent registered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomConnectorTypesAPI.ListCustomConnectorTypes(context.Background()).AssetClass(assetClass).CollectionAgentId(collectionAgentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomConnectorTypesAPI.ListCustomConnectorTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomConnectorTypes`: []CustomConnectorTypeOut
	fmt.Fprintf(os.Stdout, "Response from `CustomConnectorTypesAPI.ListCustomConnectorTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomConnectorTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetClass** | [**CustomConnectorAssetClass**](CustomConnectorAssetClass.md) | Only connector types of this asset class. | 
 **collectionAgentId** | **string** | Only connector types this collection agent registered. | 

### Return type

[**[]CustomConnectorTypeOut**](CustomConnectorTypeOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

