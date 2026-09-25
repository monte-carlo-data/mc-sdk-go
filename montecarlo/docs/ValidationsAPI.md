# \ValidationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetValidationRun**](ValidationsAPI.md#GetValidationRun) | **Get** /api/v2/validations/{run_id} | Get a validation run



## GetValidationRun

> ValidationRunOut GetValidationRun(ctx, runId).Since(since).IfNoneMatch(ifNoneMatch).Execute()

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
	since := int32(56) // int32 | The `revision` from the previous response. Only validations that changed after it are returned. Omit it to get every validation. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | The `ETag` from a previous response. The read answers 304 with no body while the run's `revision` is unchanged. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationsAPI.GetValidationRun(context.Background(), runId).Since(since).IfNoneMatch(ifNoneMatch).Execute()
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

 **since** | **int32** | The &#x60;revision&#x60; from the previous response. Only validations that changed after it are returned. Omit it to get every validation. | 
 **ifNoneMatch** | **string** | The &#x60;ETag&#x60; from a previous response. The read answers 304 with no body while the run&#39;s &#x60;revision&#x60; is unchanged. | 

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

