# \ValidationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetValidationRun**](ValidationsAPI.md#GetValidationRun) | **Get** /api/v2/validations/{run_id} | Get a validation run



## GetValidationRun

> ValidationRunOut GetValidationRun(ctx, runId).Execute()

Get a validation run



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
	runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of the validation run, as returned by the operation that started it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationsAPI.GetValidationRun(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationsAPI.GetValidationRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValidationRun`: ValidationRunOut
	fmt.Fprintf(os.Stdout, "Response from `ValidationsAPI.GetValidationRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Id of the validation run, as returned by the operation that started it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidationRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ValidationRunOut**](ValidationRunOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

