# Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **NullableString** | Position to continue from, as returned in &#x60;next_cursor&#x60; by the previous page. Omit it to start from the first page. The value is opaque; do not build or modify one. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of items to return, between 1 and 100. | [optional] [default to 50]
**WithCount** | Pointer to **bool** | Whether to also return the total number of items across every page, in &#x60;count&#x60;. Off by default: counting costs an extra query. | [optional] [default to false]

## Methods

### NewInput

`func NewInput() *Input`

NewInput instantiates a new Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputWithDefaults

`func NewInputWithDefaults() *Input`

NewInputWithDefaults instantiates a new Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *Input) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *Input) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *Input) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *Input) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *Input) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *Input) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
### GetLimit

`func (o *Input) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Input) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Input) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Input) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetWithCount

`func (o *Input) GetWithCount() bool`

GetWithCount returns the WithCount field if non-nil, zero value otherwise.

### GetWithCountOk

`func (o *Input) GetWithCountOk() (*bool, bool)`

GetWithCountOk returns a tuple with the WithCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCount

`func (o *Input) SetWithCount(v bool)`

SetWithCount sets WithCount field to given value.

### HasWithCount

`func (o *Input) HasWithCount() bool`

HasWithCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


