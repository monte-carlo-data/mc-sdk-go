# \AccountAnalyticsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountAnalyticsCatalog**](AccountAnalyticsAPI.md#GetAccountAnalyticsCatalog) | **Get** /api/v2/account-analytics/catalog | List metrics, or the dimensions that work with chosen metrics
[**GetAccountAnalyticsDocs**](AccountAnalyticsAPI.md#GetAccountAnalyticsDocs) | **Get** /api/v2/account-analytics/docs | List the Analytics API docs, or read pages
[**QueryAccountAnalytics**](AccountAnalyticsAPI.md#QueryAccountAnalytics) | **Post** /api/v2/account-analytics/query | Query metrics



## GetAccountAnalyticsCatalog

> CatalogOut GetAccountAnalyticsCatalog(ctx).Metrics(metrics).Search(search).Execute()

List metrics, or the dimensions that work with chosen metrics



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
	metrics := "metrics_example" // string | Comma-separated metric names, such as `coverage_rate,covered_assets`. With them, the answer lists the dimensions all of them can be grouped and filtered by. Without them, it lists every metric. Send the names in one parameter: only the last of a repeated parameter is read. (optional)
	search := "search_example" // string | Keeps the metrics, or the dimensions, whose name, label or description contains this text, ignoring case, such as `incident`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAnalyticsAPI.GetAccountAnalyticsCatalog(context.Background()).Metrics(metrics).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAnalyticsAPI.GetAccountAnalyticsCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAnalyticsCatalog`: CatalogOut
	fmt.Fprintf(os.Stdout, "Response from `AccountAnalyticsAPI.GetAccountAnalyticsCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAnalyticsCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metrics** | **string** | Comma-separated metric names, such as &#x60;coverage_rate,covered_assets&#x60;. With them, the answer lists the dimensions all of them can be grouped and filtered by. Without them, it lists every metric. Send the names in one parameter: only the last of a repeated parameter is read. | 
 **search** | **string** | Keeps the metrics, or the dimensions, whose name, label or description contains this text, ignoring case, such as &#x60;incident&#x60;. | 

### Return type

[**CatalogOut**](CatalogOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAnalyticsDocs

> DocsOut GetAccountAnalyticsDocs(ctx).Slug(slug).Execute()

List the Analytics API docs, or read pages



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
	slug := "slug_example" // string | The pages to read: one slug, or up to 10 comma-separated, such as `analytics-api-dashboard-style,analytics-api-dimensions`. The answer has them in that order. Without it, the answer lists every page without its body. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAnalyticsAPI.GetAccountAnalyticsDocs(context.Background()).Slug(slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAnalyticsAPI.GetAccountAnalyticsDocs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAnalyticsDocs`: DocsOut
	fmt.Fprintf(os.Stdout, "Response from `AccountAnalyticsAPI.GetAccountAnalyticsDocs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAnalyticsDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slug** | **string** | The pages to read: one slug, or up to 10 comma-separated, such as &#x60;analytics-api-dashboard-style,analytics-api-dimensions&#x60;. The answer has them in that order. Without it, the answer lists every page without its body. | 

### Return type

[**DocsOut**](DocsOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryAccountAnalytics

> MetricsQueryOut QueryAccountAnalytics(ctx).MetricsQuery(metricsQuery).Execute()

Query metrics



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
	metricsQuery := *openapiclient.NewMetricsQuery([]string{"Metrics_example"}) // MetricsQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAnalyticsAPI.QueryAccountAnalytics(context.Background()).MetricsQuery(metricsQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAnalyticsAPI.QueryAccountAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryAccountAnalytics`: MetricsQueryOut
	fmt.Fprintf(os.Stdout, "Response from `AccountAnalyticsAPI.QueryAccountAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryAccountAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsQuery** | [**MetricsQuery**](MetricsQuery.md) |  | 

### Return type

[**MetricsQueryOut**](MetricsQueryOut.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

