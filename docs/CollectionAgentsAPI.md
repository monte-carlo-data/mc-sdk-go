# \CollectionAgentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAwsCollectionAgent**](CollectionAgentsAPI.md#DeleteAwsCollectionAgent) | **Delete** /api/v2/collection-agents/aws/{collection_agent_id} | Delete an AWS collection agent
[**DeleteAzureCollectionAgent**](CollectionAgentsAPI.md#DeleteAzureCollectionAgent) | **Delete** /api/v2/collection-agents/azure/{collection_agent_id} | Delete an Azure collection agent
[**GetAwsCollectionAgent**](CollectionAgentsAPI.md#GetAwsCollectionAgent) | **Get** /api/v2/collection-agents/aws/{collection_agent_id} | Get an AWS collection agent
[**GetAzureCollectionAgent**](CollectionAgentsAPI.md#GetAzureCollectionAgent) | **Get** /api/v2/collection-agents/azure/{collection_agent_id} | Get an Azure collection agent
[**ListCollectionAgents**](CollectionAgentsAPI.md#ListCollectionAgents) | **Get** /api/v2/collection-agents | List collection agents
[**RegisterAwsCollectionAgent**](CollectionAgentsAPI.md#RegisterAwsCollectionAgent) | **Post** /api/v2/collection-agents/aws | Register an AWS collection agent
[**RegisterAzureCollectionAgent**](CollectionAgentsAPI.md#RegisterAzureCollectionAgent) | **Post** /api/v2/collection-agents/azure | Register an Azure collection agent
[**UpdateAwsCollectionAgent**](CollectionAgentsAPI.md#UpdateAwsCollectionAgent) | **Patch** /api/v2/collection-agents/aws/{collection_agent_id} | Update an AWS collection agent
[**UpdateAzureCollectionAgent**](CollectionAgentsAPI.md#UpdateAzureCollectionAgent) | **Patch** /api/v2/collection-agents/azure/{collection_agent_id} | Update an Azure collection agent



## DeleteAwsCollectionAgent

> DeleteAwsCollectionAgent(ctx, collectionAgentId).Execute()

Delete an AWS collection agent



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
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionAgentsAPI.DeleteAwsCollectionAgent(context.Background(), collectionAgentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.DeleteAwsCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionAgentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAwsCollectionAgentRequest struct via the builder pattern


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


## DeleteAzureCollectionAgent

> DeleteAzureCollectionAgent(ctx, collectionAgentId).Execute()

Delete an Azure collection agent



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
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionAgentsAPI.DeleteAzureCollectionAgent(context.Background(), collectionAgentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.DeleteAzureCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionAgentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAzureCollectionAgentRequest struct via the builder pattern


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


## GetAwsCollectionAgent

> AwsCollectionAgentOut GetAwsCollectionAgent(ctx, collectionAgentId).Execute()

Get an AWS collection agent



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
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.GetAwsCollectionAgent(context.Background(), collectionAgentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.GetAwsCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAwsCollectionAgent`: AwsCollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.GetAwsCollectionAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionAgentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsCollectionAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwsCollectionAgentOut**](AwsCollectionAgentOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAzureCollectionAgent

> AzureCollectionAgentOut GetAzureCollectionAgent(ctx, collectionAgentId).Execute()

Get an Azure collection agent



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
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.GetAzureCollectionAgent(context.Background(), collectionAgentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.GetAzureCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAzureCollectionAgent`: AzureCollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.GetAzureCollectionAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionAgentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureCollectionAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AzureCollectionAgentOut**](AzureCollectionAgentOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionAgents

> []CollectionAgentOut ListCollectionAgents(ctx).Execute()

List collection agents



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
	resp, r, err := apiClient.CollectionAgentsAPI.ListCollectionAgents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.ListCollectionAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollectionAgents`: []CollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.ListCollectionAgents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionAgentsRequest struct via the builder pattern


### Return type

[**[]CollectionAgentOut**](CollectionAgentOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAwsCollectionAgent

> AwsCollectionAgentOut RegisterAwsCollectionAgent(ctx).AwsCollectionAgentIn(awsCollectionAgentIn).Execute()

Register an AWS collection agent



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
	awsCollectionAgentIn := *openapiclient.NewAwsCollectionAgentIn("DeploymentId_example", "LambdaFunctionArn_example", "RoleArn_example") // AwsCollectionAgentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.RegisterAwsCollectionAgent(context.Background()).AwsCollectionAgentIn(awsCollectionAgentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.RegisterAwsCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAwsCollectionAgent`: AwsCollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.RegisterAwsCollectionAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAwsCollectionAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsCollectionAgentIn** | [**AwsCollectionAgentIn**](AwsCollectionAgentIn.md) |  | 

### Return type

[**AwsCollectionAgentOut**](AwsCollectionAgentOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAzureCollectionAgent

> AzureCollectionAgentOut RegisterAzureCollectionAgent(ctx).AzureCollectionAgentIn(azureCollectionAgentIn).Execute()

Register an Azure collection agent



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
	azureCollectionAgentIn := *openapiclient.NewAzureCollectionAgentIn(openapiclient.AzureAgentAuthenticationType("AZURE_FUNCTION_APP_KEY"), "DeploymentId_example", "FunctionAppUrl_example") // AzureCollectionAgentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.RegisterAzureCollectionAgent(context.Background()).AzureCollectionAgentIn(azureCollectionAgentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.RegisterAzureCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAzureCollectionAgent`: AzureCollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.RegisterAzureCollectionAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAzureCollectionAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureCollectionAgentIn** | [**AzureCollectionAgentIn**](AzureCollectionAgentIn.md) |  | 

### Return type

[**AzureCollectionAgentOut**](AzureCollectionAgentOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAwsCollectionAgent

> AwsCollectionAgentOut UpdateAwsCollectionAgent(ctx, collectionAgentId).AwsCollectionAgentPatch(awsCollectionAgentPatch).Execute()

Update an AWS collection agent



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
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	awsCollectionAgentPatch := *openapiclient.NewAwsCollectionAgentPatch() // AwsCollectionAgentPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.UpdateAwsCollectionAgent(context.Background(), collectionAgentId).AwsCollectionAgentPatch(awsCollectionAgentPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.UpdateAwsCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAwsCollectionAgent`: AwsCollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.UpdateAwsCollectionAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionAgentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAwsCollectionAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsCollectionAgentPatch** | [**AwsCollectionAgentPatch**](AwsCollectionAgentPatch.md) |  | 

### Return type

[**AwsCollectionAgentOut**](AwsCollectionAgentOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAzureCollectionAgent

> AzureCollectionAgentOut UpdateAzureCollectionAgent(ctx, collectionAgentId).AzureCollectionAgentPatch(azureCollectionAgentPatch).Execute()

Update an Azure collection agent



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
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	azureCollectionAgentPatch := *openapiclient.NewAzureCollectionAgentPatch() // AzureCollectionAgentPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.UpdateAzureCollectionAgent(context.Background(), collectionAgentId).AzureCollectionAgentPatch(azureCollectionAgentPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.UpdateAzureCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAzureCollectionAgent`: AzureCollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.UpdateAzureCollectionAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionAgentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAzureCollectionAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureCollectionAgentPatch** | [**AzureCollectionAgentPatch**](AzureCollectionAgentPatch.md) |  | 

### Return type

[**AzureCollectionAgentOut**](AzureCollectionAgentOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

