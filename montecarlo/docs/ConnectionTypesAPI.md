# \ConnectionTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListConnectionTypes**](ConnectionTypesAPI.md#ListConnectionTypes) | **Get** /api/v2/connection-types | List connection types



## ListConnectionTypes

> []ConnectionTypeOut ListConnectionTypes(ctx).Execute()

List connection types



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
	resp, r, err := apiClient.ConnectionTypesAPI.ListConnectionTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectionTypesAPI.ListConnectionTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnectionTypes`: []ConnectionTypeOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectionTypesAPI.ListConnectionTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionTypesRequest struct via the builder pattern


### Return type

[**[]ConnectionTypeOut**](ConnectionTypeOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

