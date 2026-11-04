# CatalogOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**[]CatalogItemOut**](CatalogItemOut.md) | Every metric you can query. Left out when metrics were chosen. | [optional] 
**Dimensions** | Pointer to [**[]CatalogItemOut**](CatalogItemOut.md) | The dimensions you can group and filter all of the chosen metrics by. Left out when no metrics were chosen. | [optional] 

## Methods

### NewCatalogOut

`func NewCatalogOut() *CatalogOut`

NewCatalogOut instantiates a new CatalogOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogOutWithDefaults

`func NewCatalogOutWithDefaults() *CatalogOut`

NewCatalogOutWithDefaults instantiates a new CatalogOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *CatalogOut) GetMetrics() []CatalogItemOut`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *CatalogOut) GetMetricsOk() (*[]CatalogItemOut, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *CatalogOut) SetMetrics(v []CatalogItemOut)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *CatalogOut) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### SetMetricsNil

`func (o *CatalogOut) SetMetricsNil(b bool)`

 SetMetricsNil sets the value for Metrics to be an explicit nil

### UnsetMetrics
`func (o *CatalogOut) UnsetMetrics()`

UnsetMetrics ensures that no value is present for Metrics, not even an explicit nil
### GetDimensions

`func (o *CatalogOut) GetDimensions() []CatalogItemOut`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *CatalogOut) GetDimensionsOk() (*[]CatalogItemOut, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *CatalogOut) SetDimensions(v []CatalogItemOut)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *CatalogOut) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### SetDimensionsNil

`func (o *CatalogOut) SetDimensionsNil(b bool)`

 SetDimensionsNil sets the value for Dimensions to be an explicit nil

### UnsetDimensions
`func (o *CatalogOut) UnsetDimensions()`

UnsetDimensions ensures that no value is present for Dimensions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


