# MetricsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | **[]string** | Metric names from the catalog. | 
**GroupBy** | Pointer to **[]string** | Dimension names from the catalog, such as &#x60;metric_time__week&#x60;. | [optional] 
**Filters** | Pointer to [**[]MetricsFilter**](MetricsFilter.md) | Conditions every counted row must meet, all of them. | [optional] 
**OrderBy** | Pointer to [**[]MetricsOrderBy**](MetricsOrderBy.md) | How to sort the rows, by requested names. | [optional] 
**Limit** | Pointer to **int32** | The most rows to return. The default is 1,000 and the most is 10,000. | [optional] [default to 1000]

## Methods

### NewMetricsQuery

`func NewMetricsQuery(metrics []string, ) *MetricsQuery`

NewMetricsQuery instantiates a new MetricsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQueryWithDefaults

`func NewMetricsQueryWithDefaults() *MetricsQuery`

NewMetricsQueryWithDefaults instantiates a new MetricsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *MetricsQuery) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricsQuery) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricsQuery) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.


### GetGroupBy

`func (o *MetricsQuery) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MetricsQuery) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MetricsQuery) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *MetricsQuery) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetFilters

`func (o *MetricsQuery) GetFilters() []MetricsFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MetricsQuery) GetFiltersOk() (*[]MetricsFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MetricsQuery) SetFilters(v []MetricsFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *MetricsQuery) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrderBy

`func (o *MetricsQuery) GetOrderBy() []MetricsOrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *MetricsQuery) GetOrderByOk() (*[]MetricsOrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *MetricsQuery) SetOrderBy(v []MetricsOrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *MetricsQuery) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetLimit

`func (o *MetricsQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MetricsQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MetricsQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *MetricsQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


