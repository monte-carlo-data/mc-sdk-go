# PagedConnectionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ConnectionOut**](ConnectionOut.md) | The items on this page, in the list&#39;s sort order. | 
**NextCursor** | **NullableString** | Pass this as &#x60;cursor&#x60; to fetch the next page. Null when this is the last page. | 
**HasMore** | **bool** | Whether there are more items after this page. | 
**Count** | **NullableInt32** | Total number of items across every page. Only returned when the request set &#x60;with_count&#x3D;true&#x60;; null otherwise. | 

## Methods

### NewPagedConnectionOut

`func NewPagedConnectionOut(items []ConnectionOut, nextCursor NullableString, hasMore bool, count NullableInt32, ) *PagedConnectionOut`

NewPagedConnectionOut instantiates a new PagedConnectionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedConnectionOutWithDefaults

`func NewPagedConnectionOutWithDefaults() *PagedConnectionOut`

NewPagedConnectionOutWithDefaults instantiates a new PagedConnectionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PagedConnectionOut) GetItems() []ConnectionOut`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PagedConnectionOut) GetItemsOk() (*[]ConnectionOut, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PagedConnectionOut) SetItems(v []ConnectionOut)`

SetItems sets Items field to given value.


### GetNextCursor

`func (o *PagedConnectionOut) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PagedConnectionOut) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PagedConnectionOut) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *PagedConnectionOut) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *PagedConnectionOut) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetHasMore

`func (o *PagedConnectionOut) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PagedConnectionOut) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PagedConnectionOut) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetCount

`func (o *PagedConnectionOut) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PagedConnectionOut) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PagedConnectionOut) SetCount(v int32)`

SetCount sets Count field to given value.


### SetCountNil

`func (o *PagedConnectionOut) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PagedConnectionOut) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


