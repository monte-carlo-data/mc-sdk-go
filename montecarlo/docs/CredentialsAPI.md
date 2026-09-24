# \CredentialsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAwsSecretsManagerCredentials**](CredentialsAPI.md#CreateAwsSecretsManagerCredentials) | **Post** /api/v2/credentials/self-hosted/aws | Create AWS Secrets Manager credentials
[**CreateAzureKeyVaultCredentials**](CredentialsAPI.md#CreateAzureKeyVaultCredentials) | **Post** /api/v2/credentials/self-hosted/azure | Create Azure Key Vault credentials
[**CreateEnvVarCredentials**](CredentialsAPI.md#CreateEnvVarCredentials) | **Post** /api/v2/credentials/self-hosted/env-var | Create environment variable credentials
[**CreateFileCredentials**](CredentialsAPI.md#CreateFileCredentials) | **Post** /api/v2/credentials/self-hosted/file | Create file credentials
[**CreateGcpSecretManagerCredentials**](CredentialsAPI.md#CreateGcpSecretManagerCredentials) | **Post** /api/v2/credentials/self-hosted/gcp | Create GCP Secret Manager credentials
[**CreateSnowflakeCredentials**](CredentialsAPI.md#CreateSnowflakeCredentials) | **Post** /api/v2/credentials/snowflake | Create Snowflake credentials
[**DeleteAwsSecretsManagerCredentials**](CredentialsAPI.md#DeleteAwsSecretsManagerCredentials) | **Delete** /api/v2/credentials/self-hosted/aws/{credentials_id} | Delete AWS Secrets Manager credentials
[**DeleteAzureKeyVaultCredentials**](CredentialsAPI.md#DeleteAzureKeyVaultCredentials) | **Delete** /api/v2/credentials/self-hosted/azure/{credentials_id} | Delete Azure Key Vault credentials
[**DeleteCredentials**](CredentialsAPI.md#DeleteCredentials) | **Delete** /api/v2/credentials/{credentials_id} | Delete credentials
[**DeleteEnvVarCredentials**](CredentialsAPI.md#DeleteEnvVarCredentials) | **Delete** /api/v2/credentials/self-hosted/env-var/{credentials_id} | Delete environment variable credentials
[**DeleteFileCredentials**](CredentialsAPI.md#DeleteFileCredentials) | **Delete** /api/v2/credentials/self-hosted/file/{credentials_id} | Delete file credentials
[**DeleteGcpSecretManagerCredentials**](CredentialsAPI.md#DeleteGcpSecretManagerCredentials) | **Delete** /api/v2/credentials/self-hosted/gcp/{credentials_id} | Delete GCP Secret Manager credentials
[**DeleteSnowflakeCredentials**](CredentialsAPI.md#DeleteSnowflakeCredentials) | **Delete** /api/v2/credentials/snowflake/{credentials_id} | Delete Snowflake credentials
[**GetAwsSecretsManagerCredentials**](CredentialsAPI.md#GetAwsSecretsManagerCredentials) | **Get** /api/v2/credentials/self-hosted/aws/{credentials_id} | Get AWS Secrets Manager credentials
[**GetAzureKeyVaultCredentials**](CredentialsAPI.md#GetAzureKeyVaultCredentials) | **Get** /api/v2/credentials/self-hosted/azure/{credentials_id} | Get Azure Key Vault credentials
[**GetEnvVarCredentials**](CredentialsAPI.md#GetEnvVarCredentials) | **Get** /api/v2/credentials/self-hosted/env-var/{credentials_id} | Get environment variable credentials
[**GetFileCredentials**](CredentialsAPI.md#GetFileCredentials) | **Get** /api/v2/credentials/self-hosted/file/{credentials_id} | Get file credentials
[**GetGcpSecretManagerCredentials**](CredentialsAPI.md#GetGcpSecretManagerCredentials) | **Get** /api/v2/credentials/self-hosted/gcp/{credentials_id} | Get GCP Secret Manager credentials
[**GetSnowflakeCredentials**](CredentialsAPI.md#GetSnowflakeCredentials) | **Get** /api/v2/credentials/snowflake/{credentials_id} | Get Snowflake credentials
[**ListCredentials**](CredentialsAPI.md#ListCredentials) | **Get** /api/v2/credentials | List credentials
[**UpdateAwsSecretsManagerCredentials**](CredentialsAPI.md#UpdateAwsSecretsManagerCredentials) | **Patch** /api/v2/credentials/self-hosted/aws/{credentials_id} | Update AWS Secrets Manager credentials
[**UpdateAzureKeyVaultCredentials**](CredentialsAPI.md#UpdateAzureKeyVaultCredentials) | **Patch** /api/v2/credentials/self-hosted/azure/{credentials_id} | Update Azure Key Vault credentials
[**UpdateEnvVarCredentials**](CredentialsAPI.md#UpdateEnvVarCredentials) | **Patch** /api/v2/credentials/self-hosted/env-var/{credentials_id} | Update environment variable credentials
[**UpdateFileCredentials**](CredentialsAPI.md#UpdateFileCredentials) | **Patch** /api/v2/credentials/self-hosted/file/{credentials_id} | Update file credentials
[**UpdateGcpSecretManagerCredentials**](CredentialsAPI.md#UpdateGcpSecretManagerCredentials) | **Patch** /api/v2/credentials/self-hosted/gcp/{credentials_id} | Update GCP Secret Manager credentials
[**UpdateSnowflakeCredentials**](CredentialsAPI.md#UpdateSnowflakeCredentials) | **Patch** /api/v2/credentials/snowflake/{credentials_id} | Update Snowflake credentials
[**ValidateAwsSecretsManagerCredentials**](CredentialsAPI.md#ValidateAwsSecretsManagerCredentials) | **Post** /api/v2/credentials/self-hosted/aws/validate | Validate AWS Secrets Manager credentials
[**ValidateAzureKeyVaultCredentials**](CredentialsAPI.md#ValidateAzureKeyVaultCredentials) | **Post** /api/v2/credentials/self-hosted/azure/validate | Validate Azure Key Vault credentials
[**ValidateEnvVarCredentials**](CredentialsAPI.md#ValidateEnvVarCredentials) | **Post** /api/v2/credentials/self-hosted/env-var/validate | Validate environment variable credentials
[**ValidateFileCredentials**](CredentialsAPI.md#ValidateFileCredentials) | **Post** /api/v2/credentials/self-hosted/file/validate | Validate file credentials
[**ValidateGcpSecretManagerCredentials**](CredentialsAPI.md#ValidateGcpSecretManagerCredentials) | **Post** /api/v2/credentials/self-hosted/gcp/validate | Validate GCP Secret Manager credentials
[**ValidateSnowflakeCredentials**](CredentialsAPI.md#ValidateSnowflakeCredentials) | **Post** /api/v2/credentials/snowflake/validate | Validate Snowflake credentials



## CreateAwsSecretsManagerCredentials

> AwsSecretsManagerCredentialsOut CreateAwsSecretsManagerCredentials(ctx).AwsSecretsManagerCredentialsIn(awsSecretsManagerCredentialsIn).Execute()

Create AWS Secrets Manager credentials



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
	awsSecretsManagerCredentialsIn := *openapiclient.NewAwsSecretsManagerCredentialsIn("ConnectionType_example", "AwsSecret_example") // AwsSecretsManagerCredentialsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.CreateAwsSecretsManagerCredentials(context.Background()).AwsSecretsManagerCredentialsIn(awsSecretsManagerCredentialsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.CreateAwsSecretsManagerCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAwsSecretsManagerCredentials`: AwsSecretsManagerCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.CreateAwsSecretsManagerCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAwsSecretsManagerCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsSecretsManagerCredentialsIn** | [**AwsSecretsManagerCredentialsIn**](AwsSecretsManagerCredentialsIn.md) |  | 

### Return type

[**AwsSecretsManagerCredentialsOut**](AwsSecretsManagerCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAzureKeyVaultCredentials

> AzureKeyVaultCredentialsOut CreateAzureKeyVaultCredentials(ctx).AzureKeyVaultCredentialsIn(azureKeyVaultCredentialsIn).Execute()

Create Azure Key Vault credentials



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
	azureKeyVaultCredentialsIn := *openapiclient.NewAzureKeyVaultCredentialsIn("ConnectionType_example", "AkvSecret_example") // AzureKeyVaultCredentialsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.CreateAzureKeyVaultCredentials(context.Background()).AzureKeyVaultCredentialsIn(azureKeyVaultCredentialsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.CreateAzureKeyVaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAzureKeyVaultCredentials`: AzureKeyVaultCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.CreateAzureKeyVaultCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAzureKeyVaultCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureKeyVaultCredentialsIn** | [**AzureKeyVaultCredentialsIn**](AzureKeyVaultCredentialsIn.md) |  | 

### Return type

[**AzureKeyVaultCredentialsOut**](AzureKeyVaultCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvVarCredentials

> EnvVarCredentialsOut CreateEnvVarCredentials(ctx).EnvVarCredentialsIn(envVarCredentialsIn).Execute()

Create environment variable credentials



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
	envVarCredentialsIn := *openapiclient.NewEnvVarCredentialsIn("ConnectionType_example", "EnvVarName_example") // EnvVarCredentialsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.CreateEnvVarCredentials(context.Background()).EnvVarCredentialsIn(envVarCredentialsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.CreateEnvVarCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnvVarCredentials`: EnvVarCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.CreateEnvVarCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvVarCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **envVarCredentialsIn** | [**EnvVarCredentialsIn**](EnvVarCredentialsIn.md) |  | 

### Return type

[**EnvVarCredentialsOut**](EnvVarCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFileCredentials

> FileCredentialsOut CreateFileCredentials(ctx).FileCredentialsIn(fileCredentialsIn).Execute()

Create file credentials



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
	fileCredentialsIn := *openapiclient.NewFileCredentialsIn("ConnectionType_example", "FilePath_example") // FileCredentialsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.CreateFileCredentials(context.Background()).FileCredentialsIn(fileCredentialsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.CreateFileCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFileCredentials`: FileCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.CreateFileCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileCredentialsIn** | [**FileCredentialsIn**](FileCredentialsIn.md) |  | 

### Return type

[**FileCredentialsOut**](FileCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGcpSecretManagerCredentials

> GcpSecretManagerCredentialsOut CreateGcpSecretManagerCredentials(ctx).GcpSecretManagerCredentialsIn(gcpSecretManagerCredentialsIn).Execute()

Create GCP Secret Manager credentials



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
	gcpSecretManagerCredentialsIn := *openapiclient.NewGcpSecretManagerCredentialsIn("ConnectionType_example", "GcpSecret_example") // GcpSecretManagerCredentialsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.CreateGcpSecretManagerCredentials(context.Background()).GcpSecretManagerCredentialsIn(gcpSecretManagerCredentialsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.CreateGcpSecretManagerCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGcpSecretManagerCredentials`: GcpSecretManagerCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.CreateGcpSecretManagerCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGcpSecretManagerCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gcpSecretManagerCredentialsIn** | [**GcpSecretManagerCredentialsIn**](GcpSecretManagerCredentialsIn.md) |  | 

### Return type

[**GcpSecretManagerCredentialsOut**](GcpSecretManagerCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnowflakeCredentials

> SnowflakeCredentialsOut CreateSnowflakeCredentials(ctx).SnowflakeCredentialsIn(snowflakeCredentialsIn).Execute()

Create Snowflake credentials



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
	snowflakeCredentialsIn := *openapiclient.NewSnowflakeCredentialsIn("Account_example", "User_example", "PrivateKey_example") // SnowflakeCredentialsIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.CreateSnowflakeCredentials(context.Background()).SnowflakeCredentialsIn(snowflakeCredentialsIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.CreateSnowflakeCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnowflakeCredentials`: SnowflakeCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.CreateSnowflakeCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnowflakeCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **snowflakeCredentialsIn** | [**SnowflakeCredentialsIn**](SnowflakeCredentialsIn.md) |  | 

### Return type

[**SnowflakeCredentialsOut**](SnowflakeCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAwsSecretsManagerCredentials

> DeleteAwsSecretsManagerCredentials(ctx, credentialsId).Execute()

Delete AWS Secrets Manager credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CredentialsAPI.DeleteAwsSecretsManagerCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.DeleteAwsSecretsManagerCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAwsSecretsManagerCredentialsRequest struct via the builder pattern


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


## DeleteAzureKeyVaultCredentials

> DeleteAzureKeyVaultCredentials(ctx, credentialsId).Execute()

Delete Azure Key Vault credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CredentialsAPI.DeleteAzureKeyVaultCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.DeleteAzureKeyVaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAzureKeyVaultCredentialsRequest struct via the builder pattern


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


## DeleteCredentials

> DeleteCredentials(ctx, credentialsId).Execute()

Delete credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CredentialsAPI.DeleteCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.DeleteCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredentialsRequest struct via the builder pattern


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


## DeleteEnvVarCredentials

> DeleteEnvVarCredentials(ctx, credentialsId).Execute()

Delete environment variable credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CredentialsAPI.DeleteEnvVarCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.DeleteEnvVarCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvVarCredentialsRequest struct via the builder pattern


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


## DeleteFileCredentials

> DeleteFileCredentials(ctx, credentialsId).Execute()

Delete file credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CredentialsAPI.DeleteFileCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.DeleteFileCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileCredentialsRequest struct via the builder pattern


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


## DeleteGcpSecretManagerCredentials

> DeleteGcpSecretManagerCredentials(ctx, credentialsId).Execute()

Delete GCP Secret Manager credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CredentialsAPI.DeleteGcpSecretManagerCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.DeleteGcpSecretManagerCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGcpSecretManagerCredentialsRequest struct via the builder pattern


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


## DeleteSnowflakeCredentials

> DeleteSnowflakeCredentials(ctx, credentialsId).Execute()

Delete Snowflake credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CredentialsAPI.DeleteSnowflakeCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.DeleteSnowflakeCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnowflakeCredentialsRequest struct via the builder pattern


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


## GetAwsSecretsManagerCredentials

> AwsSecretsManagerCredentialsOut GetAwsSecretsManagerCredentials(ctx, credentialsId).Execute()

Get AWS Secrets Manager credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.GetAwsSecretsManagerCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.GetAwsSecretsManagerCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAwsSecretsManagerCredentials`: AwsSecretsManagerCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.GetAwsSecretsManagerCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsSecretsManagerCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwsSecretsManagerCredentialsOut**](AwsSecretsManagerCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAzureKeyVaultCredentials

> AzureKeyVaultCredentialsOut GetAzureKeyVaultCredentials(ctx, credentialsId).Execute()

Get Azure Key Vault credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.GetAzureKeyVaultCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.GetAzureKeyVaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAzureKeyVaultCredentials`: AzureKeyVaultCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.GetAzureKeyVaultCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureKeyVaultCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AzureKeyVaultCredentialsOut**](AzureKeyVaultCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvVarCredentials

> EnvVarCredentialsOut GetEnvVarCredentials(ctx, credentialsId).Execute()

Get environment variable credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.GetEnvVarCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.GetEnvVarCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvVarCredentials`: EnvVarCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.GetEnvVarCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvVarCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvVarCredentialsOut**](EnvVarCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileCredentials

> FileCredentialsOut GetFileCredentials(ctx, credentialsId).Execute()

Get file credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.GetFileCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.GetFileCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileCredentials`: FileCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.GetFileCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileCredentialsOut**](FileCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGcpSecretManagerCredentials

> GcpSecretManagerCredentialsOut GetGcpSecretManagerCredentials(ctx, credentialsId).Execute()

Get GCP Secret Manager credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.GetGcpSecretManagerCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.GetGcpSecretManagerCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGcpSecretManagerCredentials`: GcpSecretManagerCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.GetGcpSecretManagerCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGcpSecretManagerCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GcpSecretManagerCredentialsOut**](GcpSecretManagerCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnowflakeCredentials

> SnowflakeCredentialsOut GetSnowflakeCredentials(ctx, credentialsId).Execute()

Get Snowflake credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.GetSnowflakeCredentials(context.Background(), credentialsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.GetSnowflakeCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnowflakeCredentials`: SnowflakeCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.GetSnowflakeCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnowflakeCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnowflakeCredentialsOut**](SnowflakeCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentials

> PagedCredentialsSummaryOut ListCredentials(ctx).Cursor(cursor).Limit(limit).WithCount(withCount).Execute()

List credentials



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
	resp, r, err := apiClient.CredentialsAPI.ListCredentials(context.Background()).Cursor(cursor).Limit(limit).WithCount(withCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.ListCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCredentials`: PagedCredentialsSummaryOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.ListCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Position to continue from, as returned in &#x60;next_cursor&#x60; by the previous page. Omit it to start from the first page. The value is opaque; do not build or modify one. | 
 **limit** | **int32** | Maximum number of items to return, between 1 and 100. | [default to 50]
 **withCount** | **bool** | Whether to also return the total number of items across every page, in &#x60;count&#x60;. Off by default: counting costs an extra query. | [default to false]

### Return type

[**PagedCredentialsSummaryOut**](PagedCredentialsSummaryOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAwsSecretsManagerCredentials

> AwsSecretsManagerCredentialsOut UpdateAwsSecretsManagerCredentials(ctx, credentialsId).AwsSecretsManagerCredentialsPatch(awsSecretsManagerCredentialsPatch).Execute()

Update AWS Secrets Manager credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	awsSecretsManagerCredentialsPatch := *openapiclient.NewAwsSecretsManagerCredentialsPatch() // AwsSecretsManagerCredentialsPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.UpdateAwsSecretsManagerCredentials(context.Background(), credentialsId).AwsSecretsManagerCredentialsPatch(awsSecretsManagerCredentialsPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.UpdateAwsSecretsManagerCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAwsSecretsManagerCredentials`: AwsSecretsManagerCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.UpdateAwsSecretsManagerCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAwsSecretsManagerCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsSecretsManagerCredentialsPatch** | [**AwsSecretsManagerCredentialsPatch**](AwsSecretsManagerCredentialsPatch.md) |  | 

### Return type

[**AwsSecretsManagerCredentialsOut**](AwsSecretsManagerCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAzureKeyVaultCredentials

> AzureKeyVaultCredentialsOut UpdateAzureKeyVaultCredentials(ctx, credentialsId).AzureKeyVaultCredentialsPatch(azureKeyVaultCredentialsPatch).Execute()

Update Azure Key Vault credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	azureKeyVaultCredentialsPatch := *openapiclient.NewAzureKeyVaultCredentialsPatch() // AzureKeyVaultCredentialsPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.UpdateAzureKeyVaultCredentials(context.Background(), credentialsId).AzureKeyVaultCredentialsPatch(azureKeyVaultCredentialsPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.UpdateAzureKeyVaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAzureKeyVaultCredentials`: AzureKeyVaultCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.UpdateAzureKeyVaultCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAzureKeyVaultCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureKeyVaultCredentialsPatch** | [**AzureKeyVaultCredentialsPatch**](AzureKeyVaultCredentialsPatch.md) |  | 

### Return type

[**AzureKeyVaultCredentialsOut**](AzureKeyVaultCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvVarCredentials

> EnvVarCredentialsOut UpdateEnvVarCredentials(ctx, credentialsId).EnvVarCredentialsPatch(envVarCredentialsPatch).Execute()

Update environment variable credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	envVarCredentialsPatch := *openapiclient.NewEnvVarCredentialsPatch() // EnvVarCredentialsPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.UpdateEnvVarCredentials(context.Background(), credentialsId).EnvVarCredentialsPatch(envVarCredentialsPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.UpdateEnvVarCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEnvVarCredentials`: EnvVarCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.UpdateEnvVarCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvVarCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envVarCredentialsPatch** | [**EnvVarCredentialsPatch**](EnvVarCredentialsPatch.md) |  | 

### Return type

[**EnvVarCredentialsOut**](EnvVarCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFileCredentials

> FileCredentialsOut UpdateFileCredentials(ctx, credentialsId).FileCredentialsPatch(fileCredentialsPatch).Execute()

Update file credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	fileCredentialsPatch := *openapiclient.NewFileCredentialsPatch() // FileCredentialsPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.UpdateFileCredentials(context.Background(), credentialsId).FileCredentialsPatch(fileCredentialsPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.UpdateFileCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFileCredentials`: FileCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.UpdateFileCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileCredentialsPatch** | [**FileCredentialsPatch**](FileCredentialsPatch.md) |  | 

### Return type

[**FileCredentialsOut**](FileCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGcpSecretManagerCredentials

> GcpSecretManagerCredentialsOut UpdateGcpSecretManagerCredentials(ctx, credentialsId).GcpSecretManagerCredentialsPatch(gcpSecretManagerCredentialsPatch).Execute()

Update GCP Secret Manager credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	gcpSecretManagerCredentialsPatch := *openapiclient.NewGcpSecretManagerCredentialsPatch() // GcpSecretManagerCredentialsPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.UpdateGcpSecretManagerCredentials(context.Background(), credentialsId).GcpSecretManagerCredentialsPatch(gcpSecretManagerCredentialsPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.UpdateGcpSecretManagerCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGcpSecretManagerCredentials`: GcpSecretManagerCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.UpdateGcpSecretManagerCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGcpSecretManagerCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gcpSecretManagerCredentialsPatch** | [**GcpSecretManagerCredentialsPatch**](GcpSecretManagerCredentialsPatch.md) |  | 

### Return type

[**GcpSecretManagerCredentialsOut**](GcpSecretManagerCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnowflakeCredentials

> SnowflakeCredentialsOut UpdateSnowflakeCredentials(ctx, credentialsId).SnowflakeCredentialsPatch(snowflakeCredentialsPatch).Execute()

Update Snowflake credentials



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
	credentialsId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	snowflakeCredentialsPatch := *openapiclient.NewSnowflakeCredentialsPatch() // SnowflakeCredentialsPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.UpdateSnowflakeCredentials(context.Background(), credentialsId).SnowflakeCredentialsPatch(snowflakeCredentialsPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.UpdateSnowflakeCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSnowflakeCredentials`: SnowflakeCredentialsOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.UpdateSnowflakeCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnowflakeCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snowflakeCredentialsPatch** | [**SnowflakeCredentialsPatch**](SnowflakeCredentialsPatch.md) |  | 

### Return type

[**SnowflakeCredentialsOut**](SnowflakeCredentialsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAwsSecretsManagerCredentials

> ValidationRunOut ValidateAwsSecretsManagerCredentials(ctx).AwsSecretsManagerCredentialsValidateIn(awsSecretsManagerCredentialsValidateIn).Execute()

Validate AWS Secrets Manager credentials



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
	awsSecretsManagerCredentialsValidateIn := *openapiclient.NewAwsSecretsManagerCredentialsValidateIn("DeploymentId_example", "ConnectionType_example", "AwsSecret_example") // AwsSecretsManagerCredentialsValidateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.ValidateAwsSecretsManagerCredentials(context.Background()).AwsSecretsManagerCredentialsValidateIn(awsSecretsManagerCredentialsValidateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.ValidateAwsSecretsManagerCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateAwsSecretsManagerCredentials`: ValidationRunOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.ValidateAwsSecretsManagerCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAwsSecretsManagerCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsSecretsManagerCredentialsValidateIn** | [**AwsSecretsManagerCredentialsValidateIn**](AwsSecretsManagerCredentialsValidateIn.md) |  | 

### Return type

[**ValidationRunOut**](ValidationRunOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAzureKeyVaultCredentials

> ValidationRunOut ValidateAzureKeyVaultCredentials(ctx).AzureKeyVaultCredentialsValidateIn(azureKeyVaultCredentialsValidateIn).Execute()

Validate Azure Key Vault credentials



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
	azureKeyVaultCredentialsValidateIn := *openapiclient.NewAzureKeyVaultCredentialsValidateIn("DeploymentId_example", "ConnectionType_example", "AkvSecret_example") // AzureKeyVaultCredentialsValidateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.ValidateAzureKeyVaultCredentials(context.Background()).AzureKeyVaultCredentialsValidateIn(azureKeyVaultCredentialsValidateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.ValidateAzureKeyVaultCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateAzureKeyVaultCredentials`: ValidationRunOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.ValidateAzureKeyVaultCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAzureKeyVaultCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureKeyVaultCredentialsValidateIn** | [**AzureKeyVaultCredentialsValidateIn**](AzureKeyVaultCredentialsValidateIn.md) |  | 

### Return type

[**ValidationRunOut**](ValidationRunOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateEnvVarCredentials

> ValidationRunOut ValidateEnvVarCredentials(ctx).EnvVarCredentialsValidateIn(envVarCredentialsValidateIn).Execute()

Validate environment variable credentials



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
	envVarCredentialsValidateIn := *openapiclient.NewEnvVarCredentialsValidateIn("DeploymentId_example", "ConnectionType_example", "EnvVarName_example") // EnvVarCredentialsValidateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.ValidateEnvVarCredentials(context.Background()).EnvVarCredentialsValidateIn(envVarCredentialsValidateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.ValidateEnvVarCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateEnvVarCredentials`: ValidationRunOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.ValidateEnvVarCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateEnvVarCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **envVarCredentialsValidateIn** | [**EnvVarCredentialsValidateIn**](EnvVarCredentialsValidateIn.md) |  | 

### Return type

[**ValidationRunOut**](ValidationRunOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateFileCredentials

> ValidationRunOut ValidateFileCredentials(ctx).FileCredentialsValidateIn(fileCredentialsValidateIn).Execute()

Validate file credentials



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
	fileCredentialsValidateIn := *openapiclient.NewFileCredentialsValidateIn("DeploymentId_example", "ConnectionType_example", "FilePath_example") // FileCredentialsValidateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.ValidateFileCredentials(context.Background()).FileCredentialsValidateIn(fileCredentialsValidateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.ValidateFileCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateFileCredentials`: ValidationRunOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.ValidateFileCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateFileCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileCredentialsValidateIn** | [**FileCredentialsValidateIn**](FileCredentialsValidateIn.md) |  | 

### Return type

[**ValidationRunOut**](ValidationRunOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateGcpSecretManagerCredentials

> ValidationRunOut ValidateGcpSecretManagerCredentials(ctx).GcpSecretManagerCredentialsValidateIn(gcpSecretManagerCredentialsValidateIn).Execute()

Validate GCP Secret Manager credentials



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
	gcpSecretManagerCredentialsValidateIn := *openapiclient.NewGcpSecretManagerCredentialsValidateIn("DeploymentId_example", "ConnectionType_example", "GcpSecret_example") // GcpSecretManagerCredentialsValidateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.ValidateGcpSecretManagerCredentials(context.Background()).GcpSecretManagerCredentialsValidateIn(gcpSecretManagerCredentialsValidateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.ValidateGcpSecretManagerCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateGcpSecretManagerCredentials`: ValidationRunOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.ValidateGcpSecretManagerCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateGcpSecretManagerCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gcpSecretManagerCredentialsValidateIn** | [**GcpSecretManagerCredentialsValidateIn**](GcpSecretManagerCredentialsValidateIn.md) |  | 

### Return type

[**ValidationRunOut**](ValidationRunOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSnowflakeCredentials

> ValidationRunOut ValidateSnowflakeCredentials(ctx).SnowflakeCredentialsValidateIn(snowflakeCredentialsValidateIn).Execute()

Validate Snowflake credentials



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
	snowflakeCredentialsValidateIn := *openapiclient.NewSnowflakeCredentialsValidateIn("DeploymentId_example", "Account_example", "User_example", "PrivateKey_example") // SnowflakeCredentialsValidateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialsAPI.ValidateSnowflakeCredentials(context.Background()).SnowflakeCredentialsValidateIn(snowflakeCredentialsValidateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialsAPI.ValidateSnowflakeCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateSnowflakeCredentials`: ValidationRunOut
	fmt.Fprintf(os.Stdout, "Response from `CredentialsAPI.ValidateSnowflakeCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateSnowflakeCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **snowflakeCredentialsValidateIn** | [**SnowflakeCredentialsValidateIn**](SnowflakeCredentialsValidateIn.md) |  | 

### Return type

[**ValidationRunOut**](ValidationRunOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

