# FieldErrorOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **[]string** | Where the invalid field sits, as the path to it from the top of the request body. For example &#x60;[&#39;port&#39;]&#x60;, or &#x60;[&#39;settings&#39;, &#39;timeout&#39;]&#x60; for a field nested inside an object. | 
**Code** | **string** | Machine-readable code identifying the kind of validation failure. | 
**Message** | **string** | Human-readable explanation of why this field failed validation. | 
**Extra** | Pointer to **map[string]string** | Values that parameterize &#x60;message&#x60;, for a client that renders its own text instead of showing &#x60;message&#x60; directly. Absent when there is nothing to report. | [optional] 

## Methods

### NewFieldErrorOut

`func NewFieldErrorOut(field []string, code string, message string, ) *FieldErrorOut`

NewFieldErrorOut instantiates a new FieldErrorOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldErrorOutWithDefaults

`func NewFieldErrorOutWithDefaults() *FieldErrorOut`

NewFieldErrorOutWithDefaults instantiates a new FieldErrorOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldErrorOut) GetField() []string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldErrorOut) GetFieldOk() (*[]string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldErrorOut) SetField(v []string)`

SetField sets Field field to given value.


### GetCode

`func (o *FieldErrorOut) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FieldErrorOut) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FieldErrorOut) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *FieldErrorOut) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FieldErrorOut) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FieldErrorOut) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetExtra

`func (o *FieldErrorOut) GetExtra() map[string]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *FieldErrorOut) GetExtraOk() (*map[string]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *FieldErrorOut) SetExtra(v map[string]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *FieldErrorOut) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### SetExtraNil

`func (o *FieldErrorOut) SetExtraNil(b bool)`

 SetExtraNil sets the value for Extra to be an explicit nil

### UnsetExtra
`func (o *FieldErrorOut) UnsetExtra()`

UnsetExtra ensures that no value is present for Extra, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


