# ValidationFailureOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyMessage** | **NullableString** | What went wrong, in terms a reader can act on. Null when none was given. | 
**Resolution** | **NullableString** | What to change to fix it. Null when there is no specific step to suggest. | 
**Cause** | **NullableString** | The underlying error the check ran into, as the system reported it. Values that look like passwords or tokens are masked, and long text is shortened. Null when none was reported. | 
**StackTrace** | **NullableString** | The call stack the system reported with the error, ending at the most recent call. Values that look like passwords or tokens are masked, and a long trace keeps only its end. Null when none was reported. | 

## Methods

### NewValidationFailureOut

`func NewValidationFailureOut(friendlyMessage NullableString, resolution NullableString, cause NullableString, stackTrace NullableString, ) *ValidationFailureOut`

NewValidationFailureOut instantiates a new ValidationFailureOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationFailureOutWithDefaults

`func NewValidationFailureOutWithDefaults() *ValidationFailureOut`

NewValidationFailureOutWithDefaults instantiates a new ValidationFailureOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFriendlyMessage

`func (o *ValidationFailureOut) GetFriendlyMessage() string`

GetFriendlyMessage returns the FriendlyMessage field if non-nil, zero value otherwise.

### GetFriendlyMessageOk

`func (o *ValidationFailureOut) GetFriendlyMessageOk() (*string, bool)`

GetFriendlyMessageOk returns a tuple with the FriendlyMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyMessage

`func (o *ValidationFailureOut) SetFriendlyMessage(v string)`

SetFriendlyMessage sets FriendlyMessage field to given value.


### SetFriendlyMessageNil

`func (o *ValidationFailureOut) SetFriendlyMessageNil(b bool)`

 SetFriendlyMessageNil sets the value for FriendlyMessage to be an explicit nil

### UnsetFriendlyMessage
`func (o *ValidationFailureOut) UnsetFriendlyMessage()`

UnsetFriendlyMessage ensures that no value is present for FriendlyMessage, not even an explicit nil
### GetResolution

`func (o *ValidationFailureOut) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ValidationFailureOut) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ValidationFailureOut) SetResolution(v string)`

SetResolution sets Resolution field to given value.


### SetResolutionNil

`func (o *ValidationFailureOut) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *ValidationFailureOut) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetCause

`func (o *ValidationFailureOut) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *ValidationFailureOut) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *ValidationFailureOut) SetCause(v string)`

SetCause sets Cause field to given value.


### SetCauseNil

`func (o *ValidationFailureOut) SetCauseNil(b bool)`

 SetCauseNil sets the value for Cause to be an explicit nil

### UnsetCause
`func (o *ValidationFailureOut) UnsetCause()`

UnsetCause ensures that no value is present for Cause, not even an explicit nil
### GetStackTrace

`func (o *ValidationFailureOut) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *ValidationFailureOut) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *ValidationFailureOut) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.


### SetStackTraceNil

`func (o *ValidationFailureOut) SetStackTraceNil(b bool)`

 SetStackTraceNil sets the value for StackTrace to be an explicit nil

### UnsetStackTrace
`func (o *ValidationFailureOut) UnsetStackTrace()`

UnsetStackTrace ensures that no value is present for StackTrace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


