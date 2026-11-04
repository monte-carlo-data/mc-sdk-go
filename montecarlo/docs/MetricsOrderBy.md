# MetricsOrderBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | A requested metric or dimension. | 
**Descending** | Pointer to **bool** | Sort from the largest value down. | [optional] [default to false]

## Methods

### NewMetricsOrderBy

`func NewMetricsOrderBy(field string, ) *MetricsOrderBy`

NewMetricsOrderBy instantiates a new MetricsOrderBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsOrderByWithDefaults

`func NewMetricsOrderByWithDefaults() *MetricsOrderBy`

NewMetricsOrderByWithDefaults instantiates a new MetricsOrderBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *MetricsOrderBy) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MetricsOrderBy) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MetricsOrderBy) SetField(v string)`

SetField sets Field field to given value.


### GetDescending

`func (o *MetricsOrderBy) GetDescending() bool`

GetDescending returns the Descending field if non-nil, zero value otherwise.

### GetDescendingOk

`func (o *MetricsOrderBy) GetDescendingOk() (*bool, bool)`

GetDescendingOk returns a tuple with the Descending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescending

`func (o *MetricsOrderBy) SetDescending(v bool)`

SetDescending sets Descending field to given value.

### HasDescending

`func (o *MetricsOrderBy) HasDescending() bool`

HasDescending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


