# PagedWarehouseOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]WarehouseOut**](WarehouseOut.md) | The items on this page, in the list&#39;s sort order. | 
**NextCursor** | **NullableString** | Pass this as &#x60;cursor&#x60; to fetch the next page. Null when this is the last page. | 
**HasMore** | **bool** | Whether there are more items after this page. | 
**Count** | **NullableInt32** | Total number of items across every page. Only returned when the request set &#x60;with_count&#x3D;true&#x60;; null otherwise. | 

## Methods

### NewPagedWarehouseOut

`func NewPagedWarehouseOut(items []WarehouseOut, nextCursor NullableString, hasMore bool, count NullableInt32, ) *PagedWarehouseOut`

NewPagedWarehouseOut instantiates a new PagedWarehouseOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedWarehouseOutWithDefaults

`func NewPagedWarehouseOutWithDefaults() *PagedWarehouseOut`

NewPagedWarehouseOutWithDefaults instantiates a new PagedWarehouseOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PagedWarehouseOut) GetItems() []WarehouseOut`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PagedWarehouseOut) GetItemsOk() (*[]WarehouseOut, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PagedWarehouseOut) SetItems(v []WarehouseOut)`

SetItems sets Items field to given value.


### GetNextCursor

`func (o *PagedWarehouseOut) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PagedWarehouseOut) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PagedWarehouseOut) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *PagedWarehouseOut) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *PagedWarehouseOut) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetHasMore

`func (o *PagedWarehouseOut) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PagedWarehouseOut) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PagedWarehouseOut) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetCount

`func (o *PagedWarehouseOut) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PagedWarehouseOut) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PagedWarehouseOut) SetCount(v int32)`

SetCount sets Count field to given value.


### SetCountNil

`func (o *PagedWarehouseOut) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PagedWarehouseOut) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


