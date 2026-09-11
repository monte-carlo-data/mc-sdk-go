# \CollectionAgentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGenericCollectionAgentOauthClient**](CollectionAgentsAPI.md#CreateGenericCollectionAgentOauthClient) | **Post** /api/v2/collection-agents/generic/credentials/oauth | Create an OAuth client for a generic collection agent
[**CreateGenericCollectionAgentToken**](CollectionAgentsAPI.md#CreateGenericCollectionAgentToken) | **Post** /api/v2/collection-agents/generic/credentials/token | Create a token for a generic collection agent
[**DeleteAwsCollectionAgent**](CollectionAgentsAPI.md#DeleteAwsCollectionAgent) | **Delete** /api/v2/collection-agents/aws/{collection_agent_id} | Delete an AWS collection agent
[**DeleteAzureCollectionAgent**](CollectionAgentsAPI.md#DeleteAzureCollectionAgent) | **Delete** /api/v2/collection-agents/azure/{collection_agent_id} | Delete an Azure collection agent
[**DeleteGcpCollectionAgent**](CollectionAgentsAPI.md#DeleteGcpCollectionAgent) | **Delete** /api/v2/collection-agents/gcp/{collection_agent_id} | Delete a GCP collection agent
[**DeleteGenericCollectionAgent**](CollectionAgentsAPI.md#DeleteGenericCollectionAgent) | **Delete** /api/v2/collection-agents/generic/{collection_agent_id} | Delete a generic collection agent
[**DeleteGenericCollectionAgentOauthClient**](CollectionAgentsAPI.md#DeleteGenericCollectionAgentOauthClient) | **Delete** /api/v2/collection-agents/generic/credentials/oauth/{credential_id} | Delete a generic collection agent OAuth client
[**DeleteGenericCollectionAgentToken**](CollectionAgentsAPI.md#DeleteGenericCollectionAgentToken) | **Delete** /api/v2/collection-agents/generic/credentials/token/{credential_id} | Delete a generic collection agent token
[**GetAwsCollectionAgent**](CollectionAgentsAPI.md#GetAwsCollectionAgent) | **Get** /api/v2/collection-agents/aws/{collection_agent_id} | Get an AWS collection agent
[**GetAzureCollectionAgent**](CollectionAgentsAPI.md#GetAzureCollectionAgent) | **Get** /api/v2/collection-agents/azure/{collection_agent_id} | Get an Azure collection agent
[**GetGcpCollectionAgent**](CollectionAgentsAPI.md#GetGcpCollectionAgent) | **Get** /api/v2/collection-agents/gcp/{collection_agent_id} | Get a GCP collection agent
[**GetGenericCollectionAgent**](CollectionAgentsAPI.md#GetGenericCollectionAgent) | **Get** /api/v2/collection-agents/generic/{collection_agent_id} | Get a generic collection agent
[**GetGenericCollectionAgentOauthClient**](CollectionAgentsAPI.md#GetGenericCollectionAgentOauthClient) | **Get** /api/v2/collection-agents/generic/credentials/oauth/{credential_id} | Get a generic collection agent OAuth client
[**GetGenericCollectionAgentToken**](CollectionAgentsAPI.md#GetGenericCollectionAgentToken) | **Get** /api/v2/collection-agents/generic/credentials/token/{credential_id} | Get a generic collection agent token
[**ListCollectionAgents**](CollectionAgentsAPI.md#ListCollectionAgents) | **Get** /api/v2/collection-agents | List collection agents
[**ListGenericCollectionAgentCredentials**](CollectionAgentsAPI.md#ListGenericCollectionAgentCredentials) | **Get** /api/v2/collection-agents/generic/credentials | List generic collection agent credentials
[**RegisterAwsCollectionAgent**](CollectionAgentsAPI.md#RegisterAwsCollectionAgent) | **Post** /api/v2/collection-agents/aws | Register an AWS collection agent
[**RegisterAzureCollectionAgent**](CollectionAgentsAPI.md#RegisterAzureCollectionAgent) | **Post** /api/v2/collection-agents/azure | Register an Azure collection agent
[**RegisterGcpCollectionAgent**](CollectionAgentsAPI.md#RegisterGcpCollectionAgent) | **Post** /api/v2/collection-agents/gcp | Register a GCP collection agent
[**RegisterGenericCollectionAgent**](CollectionAgentsAPI.md#RegisterGenericCollectionAgent) | **Post** /api/v2/collection-agents/generic | Register a generic collection agent
[**UpdateAwsCollectionAgent**](CollectionAgentsAPI.md#UpdateAwsCollectionAgent) | **Patch** /api/v2/collection-agents/aws/{collection_agent_id} | Update an AWS collection agent
[**UpdateAzureCollectionAgent**](CollectionAgentsAPI.md#UpdateAzureCollectionAgent) | **Patch** /api/v2/collection-agents/azure/{collection_agent_id} | Update an Azure collection agent
[**UpdateGcpCollectionAgent**](CollectionAgentsAPI.md#UpdateGcpCollectionAgent) | **Patch** /api/v2/collection-agents/gcp/{collection_agent_id} | Update a GCP collection agent
[**UpdateGenericCollectionAgent**](CollectionAgentsAPI.md#UpdateGenericCollectionAgent) | **Patch** /api/v2/collection-agents/generic/{collection_agent_id} | Update a generic collection agent



## CreateGenericCollectionAgentOauthClient

> GenericCollectionAgentOAuthClientCreatedOut CreateGenericCollectionAgentOauthClient(ctx).GenericCollectionAgentOAuthClientIn(genericCollectionAgentOAuthClientIn).Execute()

Create an OAuth client for a generic collection agent



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
	genericCollectionAgentOAuthClientIn := *openapiclient.NewGenericCollectionAgentOAuthClientIn("DeploymentId_example") // GenericCollectionAgentOAuthClientIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.CreateGenericCollectionAgentOauthClient(context.Background()).GenericCollectionAgentOAuthClientIn(genericCollectionAgentOAuthClientIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.CreateGenericCollectionAgentOauthClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGenericCollectionAgentOauthClient`: GenericCollectionAgentOAuthClientCreatedOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.CreateGenericCollectionAgentOauthClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGenericCollectionAgentOauthClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genericCollectionAgentOAuthClientIn** | [**GenericCollectionAgentOAuthClientIn**](GenericCollectionAgentOAuthClientIn.md) |  | 

### Return type

[**GenericCollectionAgentOAuthClientCreatedOut**](GenericCollectionAgentOAuthClientCreatedOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGenericCollectionAgentToken

> GenericCollectionAgentTokenCreatedOut CreateGenericCollectionAgentToken(ctx).GenericCollectionAgentTokenIn(genericCollectionAgentTokenIn).Execute()

Create a token for a generic collection agent



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
	genericCollectionAgentTokenIn := *openapiclient.NewGenericCollectionAgentTokenIn("DeploymentId_example") // GenericCollectionAgentTokenIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.CreateGenericCollectionAgentToken(context.Background()).GenericCollectionAgentTokenIn(genericCollectionAgentTokenIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.CreateGenericCollectionAgentToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGenericCollectionAgentToken`: GenericCollectionAgentTokenCreatedOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.CreateGenericCollectionAgentToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGenericCollectionAgentTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genericCollectionAgentTokenIn** | [**GenericCollectionAgentTokenIn**](GenericCollectionAgentTokenIn.md) |  | 

### Return type

[**GenericCollectionAgentTokenCreatedOut**](GenericCollectionAgentTokenCreatedOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
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


## DeleteGcpCollectionAgent

> DeleteGcpCollectionAgent(ctx, collectionAgentId).Execute()

Delete a GCP collection agent



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
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionAgentsAPI.DeleteGcpCollectionAgent(context.Background(), collectionAgentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.DeleteGcpCollectionAgent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGcpCollectionAgentRequest struct via the builder pattern


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


## DeleteGenericCollectionAgent

> DeleteGenericCollectionAgent(ctx, collectionAgentId).Execute()

Delete a generic collection agent



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
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionAgentsAPI.DeleteGenericCollectionAgent(context.Background(), collectionAgentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.DeleteGenericCollectionAgent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGenericCollectionAgentRequest struct via the builder pattern


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


## DeleteGenericCollectionAgentOauthClient

> DeleteGenericCollectionAgentOauthClient(ctx, credentialId).Execute()

Delete a generic collection agent OAuth client



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
	credentialId := "credentialId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionAgentsAPI.DeleteGenericCollectionAgentOauthClient(context.Background(), credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.DeleteGenericCollectionAgentOauthClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGenericCollectionAgentOauthClientRequest struct via the builder pattern


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


## DeleteGenericCollectionAgentToken

> DeleteGenericCollectionAgentToken(ctx, credentialId).Execute()

Delete a generic collection agent token



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
	credentialId := "credentialId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionAgentsAPI.DeleteGenericCollectionAgentToken(context.Background(), credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.DeleteGenericCollectionAgentToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGenericCollectionAgentTokenRequest struct via the builder pattern


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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
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


## GetGcpCollectionAgent

> GcpCollectionAgentOut GetGcpCollectionAgent(ctx, collectionAgentId).Execute()

Get a GCP collection agent



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
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.GetGcpCollectionAgent(context.Background(), collectionAgentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.GetGcpCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGcpCollectionAgent`: GcpCollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.GetGcpCollectionAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionAgentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpCollectionAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GcpCollectionAgentOut**](GcpCollectionAgentOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenericCollectionAgent

> GenericCollectionAgentOut GetGenericCollectionAgent(ctx, collectionAgentId).Execute()

Get a generic collection agent



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
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.GetGenericCollectionAgent(context.Background(), collectionAgentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.GetGenericCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenericCollectionAgent`: GenericCollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.GetGenericCollectionAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionAgentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenericCollectionAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenericCollectionAgentOut**](GenericCollectionAgentOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenericCollectionAgentOauthClient

> GenericCollectionAgentOAuthClientOut GetGenericCollectionAgentOauthClient(ctx, credentialId).Execute()

Get a generic collection agent OAuth client



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
	credentialId := "credentialId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.GetGenericCollectionAgentOauthClient(context.Background(), credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.GetGenericCollectionAgentOauthClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenericCollectionAgentOauthClient`: GenericCollectionAgentOAuthClientOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.GetGenericCollectionAgentOauthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenericCollectionAgentOauthClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenericCollectionAgentOAuthClientOut**](GenericCollectionAgentOAuthClientOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenericCollectionAgentToken

> GenericCollectionAgentTokenOut GetGenericCollectionAgentToken(ctx, credentialId).Execute()

Get a generic collection agent token



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
	credentialId := "credentialId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.GetGenericCollectionAgentToken(context.Background(), credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.GetGenericCollectionAgentToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenericCollectionAgentToken`: GenericCollectionAgentTokenOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.GetGenericCollectionAgentToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenericCollectionAgentTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenericCollectionAgentTokenOut**](GenericCollectionAgentTokenOut.md)

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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
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


## ListGenericCollectionAgentCredentials

> []CollectionAgentCredentialOut ListGenericCollectionAgentCredentials(ctx).DeploymentId(deploymentId).Execute()

List generic collection agent credentials



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only the credentials of this deployment's agent. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.ListGenericCollectionAgentCredentials(context.Background()).DeploymentId(deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.ListGenericCollectionAgentCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGenericCollectionAgentCredentials`: []CollectionAgentCredentialOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.ListGenericCollectionAgentCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGenericCollectionAgentCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentId** | **string** | Only the credentials of this deployment&#39;s agent. | 

### Return type

[**[]CollectionAgentCredentialOut**](CollectionAgentCredentialOut.md)

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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
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


## RegisterGcpCollectionAgent

> GcpCollectionAgentOut RegisterGcpCollectionAgent(ctx).GcpCollectionAgentIn(gcpCollectionAgentIn).Execute()

Register a GCP collection agent



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
	gcpCollectionAgentIn := *openapiclient.NewGcpCollectionAgentIn(openapiclient.GcpAgentAuthenticationType("GCP_JSON_SERVICE_ACCOUNT_KEY"), "DeploymentId_example", "CloudRunUrl_example") // GcpCollectionAgentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.RegisterGcpCollectionAgent(context.Background()).GcpCollectionAgentIn(gcpCollectionAgentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.RegisterGcpCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterGcpCollectionAgent`: GcpCollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.RegisterGcpCollectionAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterGcpCollectionAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gcpCollectionAgentIn** | [**GcpCollectionAgentIn**](GcpCollectionAgentIn.md) |  | 

### Return type

[**GcpCollectionAgentOut**](GcpCollectionAgentOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterGenericCollectionAgent

> GenericCollectionAgentOut RegisterGenericCollectionAgent(ctx).GenericCollectionAgentIn(genericCollectionAgentIn).Execute()

Register a generic collection agent



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
	genericCollectionAgentIn := *openapiclient.NewGenericCollectionAgentIn("DeploymentId_example") // GenericCollectionAgentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.RegisterGenericCollectionAgent(context.Background()).GenericCollectionAgentIn(genericCollectionAgentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.RegisterGenericCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterGenericCollectionAgent`: GenericCollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.RegisterGenericCollectionAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterGenericCollectionAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genericCollectionAgentIn** | [**GenericCollectionAgentIn**](GenericCollectionAgentIn.md) |  | 

### Return type

[**GenericCollectionAgentOut**](GenericCollectionAgentOut.md)

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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
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
	openapiclient "github.com/monte-carlo-data/mc-sdk-go/montecarlo"
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


## UpdateGcpCollectionAgent

> GcpCollectionAgentOut UpdateGcpCollectionAgent(ctx, collectionAgentId).GcpCollectionAgentPatch(gcpCollectionAgentPatch).Execute()

Update a GCP collection agent



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
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	gcpCollectionAgentPatch := *openapiclient.NewGcpCollectionAgentPatch() // GcpCollectionAgentPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.UpdateGcpCollectionAgent(context.Background(), collectionAgentId).GcpCollectionAgentPatch(gcpCollectionAgentPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.UpdateGcpCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGcpCollectionAgent`: GcpCollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.UpdateGcpCollectionAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionAgentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGcpCollectionAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpCollectionAgentPatch** | [**GcpCollectionAgentPatch**](GcpCollectionAgentPatch.md) |  | 

### Return type

[**GcpCollectionAgentOut**](GcpCollectionAgentOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGenericCollectionAgent

> GenericCollectionAgentOut UpdateGenericCollectionAgent(ctx, collectionAgentId).GenericCollectionAgentPatch(genericCollectionAgentPatch).Execute()

Update a generic collection agent



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
	collectionAgentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	genericCollectionAgentPatch := *openapiclient.NewGenericCollectionAgentPatch() // GenericCollectionAgentPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAgentsAPI.UpdateGenericCollectionAgent(context.Background(), collectionAgentId).GenericCollectionAgentPatch(genericCollectionAgentPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAgentsAPI.UpdateGenericCollectionAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGenericCollectionAgent`: GenericCollectionAgentOut
	fmt.Fprintf(os.Stdout, "Response from `CollectionAgentsAPI.UpdateGenericCollectionAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionAgentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGenericCollectionAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **genericCollectionAgentPatch** | [**GenericCollectionAgentPatch**](GenericCollectionAgentPatch.md) |  | 

### Return type

[**GenericCollectionAgentOut**](GenericCollectionAgentOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

