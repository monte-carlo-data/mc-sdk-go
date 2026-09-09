# ProblemOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Machine-readable code identifying the category of error. Branch on this rather than on &#x60;status&#x60;, because several codes can share one status. | 
**Detail** | **string** | Human-readable explanation of this particular occurrence of the error. | 
**Errors** | Pointer to [**[]FieldErrorOut**](FieldErrorOut.md) | The fields that failed validation. Empty unless this is a validation error. | [optional] 
**Extra** | Pointer to **map[string]string** | Values that parameterize &#x60;message&#x60;, for a client that renders its own text instead of showing &#x60;message&#x60; directly. Absent when there is nothing to report. | [optional] 
**RequestId** | **string** | Identifier for this request, also returned in the &#x60;X-Request-Id&#x60; header. Quote it when reporting the error to support. | 
**Status** | **int32** | HTTP status code, repeated here for clients that only see the body. | 
**Title** | **string** | Short, human-readable summary of the category of error. | 
**Type** | **string** | URI identifying the category of error. The same category always uses the same URI. | 

## Methods

### NewProblemOut

`func NewProblemOut(code string, detail string, requestId string, status int32, title string, type_ string, ) *ProblemOut`

NewProblemOut instantiates a new ProblemOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemOutWithDefaults

`func NewProblemOutWithDefaults() *ProblemOut`

NewProblemOutWithDefaults instantiates a new ProblemOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ProblemOut) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProblemOut) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProblemOut) SetCode(v string)`

SetCode sets Code field to given value.


### GetDetail

`func (o *ProblemOut) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ProblemOut) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ProblemOut) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetErrors

`func (o *ProblemOut) GetErrors() []FieldErrorOut`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ProblemOut) GetErrorsOk() (*[]FieldErrorOut, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ProblemOut) SetErrors(v []FieldErrorOut)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ProblemOut) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetExtra

`func (o *ProblemOut) GetExtra() map[string]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *ProblemOut) GetExtraOk() (*map[string]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *ProblemOut) SetExtra(v map[string]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *ProblemOut) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### SetExtraNil

`func (o *ProblemOut) SetExtraNil(b bool)`

 SetExtraNil sets the value for Extra to be an explicit nil

### UnsetExtra
`func (o *ProblemOut) UnsetExtra()`

UnsetExtra ensures that no value is present for Extra, not even an explicit nil
### GetRequestId

`func (o *ProblemOut) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProblemOut) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProblemOut) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetStatus

`func (o *ProblemOut) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProblemOut) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProblemOut) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *ProblemOut) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProblemOut) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProblemOut) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *ProblemOut) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProblemOut) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProblemOut) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


