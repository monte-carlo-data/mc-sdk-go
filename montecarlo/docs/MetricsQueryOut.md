# MetricsQueryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | **[]string** | The &#x60;group_by&#x60; names, then the &#x60;metrics&#x60; names, in request order. | 
**Rows** | **[][]interface{}** | One list of values per row, in &#x60;columns&#x60; order. Metric values are numbers or null. | 
**DataAsOf** | **NullableTime** | The time the data in this answer reflects. The data is refreshed periodically, not in real time, so later changes are not in it. Null when there is no data to answer from. | 
**Truncated** | **bool** | True when more rows matched than &#x60;limit&#x60;. Add filters, narrow the time range or group by fewer dimensions. Raise &#x60;limit&#x60; only up to 10,000. | 
**Limit** | **int32** | The most rows this answer could hold, as asked for. | 
**RowsReturned** | **int32** | The number of rows in &#x60;rows&#x60;. | 
**TotalRows** | **NullableInt32** | The number of rows the query matched, which is more than &#x60;rows_returned&#x60; when &#x60;truncated&#x60;. Null when it is not known. Use it instead of counting &#x60;rows&#x60;. | 
**Warnings** | **[]string** | Short notes on how to read the answer, such as how to narrow a truncated one. Empty when there is nothing to note. | 

## Methods

### NewMetricsQueryOut

`func NewMetricsQueryOut(columns []string, rows [][]interface{}, dataAsOf NullableTime, truncated bool, limit int32, rowsReturned int32, totalRows NullableInt32, warnings []string, ) *MetricsQueryOut`

NewMetricsQueryOut instantiates a new MetricsQueryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQueryOutWithDefaults

`func NewMetricsQueryOutWithDefaults() *MetricsQueryOut`

NewMetricsQueryOutWithDefaults instantiates a new MetricsQueryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *MetricsQueryOut) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *MetricsQueryOut) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *MetricsQueryOut) SetColumns(v []string)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *MetricsQueryOut) GetRows() [][]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *MetricsQueryOut) GetRowsOk() (*[][]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *MetricsQueryOut) SetRows(v [][]interface{})`

SetRows sets Rows field to given value.


### GetDataAsOf

`func (o *MetricsQueryOut) GetDataAsOf() time.Time`

GetDataAsOf returns the DataAsOf field if non-nil, zero value otherwise.

### GetDataAsOfOk

`func (o *MetricsQueryOut) GetDataAsOfOk() (*time.Time, bool)`

GetDataAsOfOk returns a tuple with the DataAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAsOf

`func (o *MetricsQueryOut) SetDataAsOf(v time.Time)`

SetDataAsOf sets DataAsOf field to given value.


### SetDataAsOfNil

`func (o *MetricsQueryOut) SetDataAsOfNil(b bool)`

 SetDataAsOfNil sets the value for DataAsOf to be an explicit nil

### UnsetDataAsOf
`func (o *MetricsQueryOut) UnsetDataAsOf()`

UnsetDataAsOf ensures that no value is present for DataAsOf, not even an explicit nil
### GetTruncated

`func (o *MetricsQueryOut) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *MetricsQueryOut) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *MetricsQueryOut) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.


### GetLimit

`func (o *MetricsQueryOut) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MetricsQueryOut) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MetricsQueryOut) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetRowsReturned

`func (o *MetricsQueryOut) GetRowsReturned() int32`

GetRowsReturned returns the RowsReturned field if non-nil, zero value otherwise.

### GetRowsReturnedOk

`func (o *MetricsQueryOut) GetRowsReturnedOk() (*int32, bool)`

GetRowsReturnedOk returns a tuple with the RowsReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsReturned

`func (o *MetricsQueryOut) SetRowsReturned(v int32)`

SetRowsReturned sets RowsReturned field to given value.


### GetTotalRows

`func (o *MetricsQueryOut) GetTotalRows() int32`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *MetricsQueryOut) GetTotalRowsOk() (*int32, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *MetricsQueryOut) SetTotalRows(v int32)`

SetTotalRows sets TotalRows field to given value.


### SetTotalRowsNil

`func (o *MetricsQueryOut) SetTotalRowsNil(b bool)`

 SetTotalRowsNil sets the value for TotalRows to be an explicit nil

### UnsetTotalRows
`func (o *MetricsQueryOut) UnsetTotalRows()`

UnsetTotalRows ensures that no value is present for TotalRows, not even an explicit nil
### GetWarnings

`func (o *MetricsQueryOut) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *MetricsQueryOut) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *MetricsQueryOut) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


