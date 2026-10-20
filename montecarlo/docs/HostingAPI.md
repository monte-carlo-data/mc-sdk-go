# \HostingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHostingDomains**](HostingAPI.md#GetHostingDomains) | **Get** /api/v2/hosting/domains | Get the hosting domains
[**GetHostingNetwork**](HostingAPI.md#GetHostingNetwork) | **Get** /api/v2/hosting/network | Get the hosting network
[**GetHostingPrivateLink**](HostingAPI.md#GetHostingPrivateLink) | **Get** /api/v2/hosting/private-link | Get the hosting private link details



## GetHostingDomains

> HostingDomainsOut GetHostingDomains(ctx).Execute()

Get the hosting domains



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
	resp, r, err := apiClient.HostingAPI.GetHostingDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostingAPI.GetHostingDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostingDomains`: HostingDomainsOut
	fmt.Fprintf(os.Stdout, "Response from `HostingAPI.GetHostingDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostingDomainsRequest struct via the builder pattern


### Return type

[**HostingDomainsOut**](HostingDomainsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostingNetwork

> HostingNetworkOut GetHostingNetwork(ctx).InstanceRegion(instanceRegion).Execute()

Get the hosting network



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
	instanceRegion := openapiclient.InstanceRegion("primary") // InstanceRegion | The region of your Monte Carlo instance to read. `primary`, the default, is the region serving your instance, which is the recovery region during a failover. `secondary` is the disaster recovery region, and only an instance with disaster recovery has one. With disaster recovery, read both and allow both, so access keeps working after a failover. During a failover both read the recovery region, so applying your policies then drops the original primary region until you apply them again after failing back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostingAPI.GetHostingNetwork(context.Background()).InstanceRegion(instanceRegion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostingAPI.GetHostingNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostingNetwork`: HostingNetworkOut
	fmt.Fprintf(os.Stdout, "Response from `HostingAPI.GetHostingNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostingNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceRegion** | [**InstanceRegion**](InstanceRegion.md) | The region of your Monte Carlo instance to read. &#x60;primary&#x60;, the default, is the region serving your instance, which is the recovery region during a failover. &#x60;secondary&#x60; is the disaster recovery region, and only an instance with disaster recovery has one. With disaster recovery, read both and allow both, so access keeps working after a failover. During a failover both read the recovery region, so applying your policies then drops the original primary region until you apply them again after failing back. | 

### Return type

[**HostingNetworkOut**](HostingNetworkOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostingPrivateLink

> HostingPrivateLinkOut GetHostingPrivateLink(ctx).InstanceRegion(instanceRegion).Execute()

Get the hosting private link details



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
	instanceRegion := openapiclient.InstanceRegion("primary") // InstanceRegion | The region of your Monte Carlo instance to read. `primary`, the default, is the region serving your instance, which is the recovery region during a failover. `secondary` is the disaster recovery region, and only an instance with disaster recovery has one. With disaster recovery, read both and allow both, so access keeps working after a failover. During a failover both read the recovery region, so applying your policies then drops the original primary region until you apply them again after failing back. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostingAPI.GetHostingPrivateLink(context.Background()).InstanceRegion(instanceRegion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostingAPI.GetHostingPrivateLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostingPrivateLink`: HostingPrivateLinkOut
	fmt.Fprintf(os.Stdout, "Response from `HostingAPI.GetHostingPrivateLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostingPrivateLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceRegion** | [**InstanceRegion**](InstanceRegion.md) | The region of your Monte Carlo instance to read. &#x60;primary&#x60;, the default, is the region serving your instance, which is the recovery region during a failover. &#x60;secondary&#x60; is the disaster recovery region, and only an instance with disaster recovery has one. With disaster recovery, read both and allow both, so access keeps working after a failover. During a failover both read the recovery region, so applying your policies then drops the original primary region until you apply them again after failing back. | 

### Return type

[**HostingPrivateLinkOut**](HostingPrivateLinkOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

