# ValidationOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Identifier of the validation, unique within the run. | 
**Description** | **NullableString** | What the validation checks. | 
**Status** | [**ValidationStatus**](ValidationStatus.md) | Whether the validation ran, not whether what it checked is healthy. | 
**IsPrerequisite** | **bool** | Whether the rest of the run depends on this one. A prerequisite that does not pass skips every validation waiting on it. | 
**Passed** | **NullableBool** | Whether the validation found no blocking problem. Null until it reaches a verdict, and for one that never did. | 
**Errors** | [**[]ValidationFailureOut**](ValidationFailureOut.md) | Blocking problems. Empty when the validation passed. | 
**Warnings** | [**[]ValidationFailureOut**](ValidationFailureOut.md) | Problems that did not stop the validation from passing. | 
**AdditionalData** | **map[string]interface{}** | Extra detail the validation reported, if it reported any. | 
**Truncated** | **bool** | Whether some detail was too large to return and was clipped. The problems that are here are still accurate; there may have been more of them. | 

## Methods

### NewValidationOut

`func NewValidationOut(name string, description NullableString, status ValidationStatus, isPrerequisite bool, passed NullableBool, errors []ValidationFailureOut, warnings []ValidationFailureOut, additionalData map[string]interface{}, truncated bool, ) *ValidationOut`

NewValidationOut instantiates a new ValidationOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationOutWithDefaults

`func NewValidationOutWithDefaults() *ValidationOut`

NewValidationOutWithDefaults instantiates a new ValidationOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ValidationOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidationOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidationOut) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ValidationOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ValidationOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ValidationOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ValidationOut) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ValidationOut) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *ValidationOut) GetStatus() ValidationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidationOut) GetStatusOk() (*ValidationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidationOut) SetStatus(v ValidationStatus)`

SetStatus sets Status field to given value.


### GetIsPrerequisite

`func (o *ValidationOut) GetIsPrerequisite() bool`

GetIsPrerequisite returns the IsPrerequisite field if non-nil, zero value otherwise.

### GetIsPrerequisiteOk

`func (o *ValidationOut) GetIsPrerequisiteOk() (*bool, bool)`

GetIsPrerequisiteOk returns a tuple with the IsPrerequisite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrerequisite

`func (o *ValidationOut) SetIsPrerequisite(v bool)`

SetIsPrerequisite sets IsPrerequisite field to given value.


### GetPassed

`func (o *ValidationOut) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *ValidationOut) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *ValidationOut) SetPassed(v bool)`

SetPassed sets Passed field to given value.


### SetPassedNil

`func (o *ValidationOut) SetPassedNil(b bool)`

 SetPassedNil sets the value for Passed to be an explicit nil

### UnsetPassed
`func (o *ValidationOut) UnsetPassed()`

UnsetPassed ensures that no value is present for Passed, not even an explicit nil
### GetErrors

`func (o *ValidationOut) GetErrors() []ValidationFailureOut`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidationOut) GetErrorsOk() (*[]ValidationFailureOut, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidationOut) SetErrors(v []ValidationFailureOut)`

SetErrors sets Errors field to given value.


### GetWarnings

`func (o *ValidationOut) GetWarnings() []ValidationFailureOut`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ValidationOut) GetWarningsOk() (*[]ValidationFailureOut, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ValidationOut) SetWarnings(v []ValidationFailureOut)`

SetWarnings sets Warnings field to given value.


### GetAdditionalData

`func (o *ValidationOut) GetAdditionalData() map[string]interface{}`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *ValidationOut) GetAdditionalDataOk() (*map[string]interface{}, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *ValidationOut) SetAdditionalData(v map[string]interface{})`

SetAdditionalData sets AdditionalData field to given value.


### SetAdditionalDataNil

`func (o *ValidationOut) SetAdditionalDataNil(b bool)`

 SetAdditionalDataNil sets the value for AdditionalData to be an explicit nil

### UnsetAdditionalData
`func (o *ValidationOut) UnsetAdditionalData()`

UnsetAdditionalData ensures that no value is present for AdditionalData, not even an explicit nil
### GetTruncated

`func (o *ValidationOut) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *ValidationOut) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *ValidationOut) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


