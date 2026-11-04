# MetricsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | A dimension name from the catalog. | 
**Op** | [**MetricsFilterOperator**](MetricsFilterOperator.md) | &#x60;in&#x60;, &#x60;not_in&#x60;: any of the values. &#x60;not_in&#x60; also keeps rows where the dimension is empty. &#x60;gte&#x60;, &#x60;lte&#x60;, &#x60;gt&#x60;, &#x60;lt&#x60;: compare with one value and never match an empty dimension. &#x60;contains&#x60;: one text, matched anywhere in the value, ignoring case. &#x60;is_null&#x60;: one boolean, &#x60;true&#x60; for rows where the dimension is empty and &#x60;false&#x60; for rows where it is not. The text &#x60;true&#x60; or &#x60;false&#x60; works too. | 
**Values** | **[]interface{}** | The values to compare with. A date is an ISO string, such as &#x60;2026-10-01&#x60;. | 

## Methods

### NewMetricsFilter

`func NewMetricsFilter(field string, op MetricsFilterOperator, values []interface{}, ) *MetricsFilter`

NewMetricsFilter instantiates a new MetricsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsFilterWithDefaults

`func NewMetricsFilterWithDefaults() *MetricsFilter`

NewMetricsFilterWithDefaults instantiates a new MetricsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *MetricsFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MetricsFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MetricsFilter) SetField(v string)`

SetField sets Field field to given value.


### GetOp

`func (o *MetricsFilter) GetOp() MetricsFilterOperator`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *MetricsFilter) GetOpOk() (*MetricsFilterOperator, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *MetricsFilter) SetOp(v MetricsFilterOperator)`

SetOp sets Op field to given value.


### GetValues

`func (o *MetricsFilter) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MetricsFilter) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MetricsFilter) SetValues(v []interface{})`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


