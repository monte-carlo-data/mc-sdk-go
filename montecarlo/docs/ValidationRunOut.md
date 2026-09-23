# ValidationRunOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the run. Poll &#x60;GET /validations/{run_id}&#x60; with it. | 
**Status** | [**RunStatus**](RunStatus.md) | Whether the run is still going. Every validation is final once it is not. | 
**TargetType** | [**TargetType**](TargetType.md) | What the run validates. | 
**TargetId** | **NullableString** | Identifier of what is being validated. Null for candidate values, which are not stored anywhere. | 
**ValidationsPassed** | **int32** | How many validations reached a passing verdict. One that was skipped or never reached a verdict is not counted here, but is still in &#x60;validations_total&#x60;. | 
**ValidationsTotal** | **int32** | How many validations the run covers. | 
**StartedAt** | **time.Time** | When the run started. | 
**FinishedAt** | **NullableTime** | When the run finished. Null while it is still going. | 
**ExpiresAt** | **time.Time** | When the run stops being readable. Measured from the start, not the finish, and never extended, so a slow run is readable for less time after it ends. | 
**Validations** | [**[]ValidationOut**](ValidationOut.md) | Every validation the run covers, in the order they are declared. Validations waiting on a prerequisite are listed before they start. | 

## Methods

### NewValidationRunOut

`func NewValidationRunOut(id string, status RunStatus, targetType TargetType, targetId NullableString, validationsPassed int32, validationsTotal int32, startedAt time.Time, finishedAt NullableTime, expiresAt time.Time, validations []ValidationOut, ) *ValidationRunOut`

NewValidationRunOut instantiates a new ValidationRunOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationRunOutWithDefaults

`func NewValidationRunOutWithDefaults() *ValidationRunOut`

NewValidationRunOutWithDefaults instantiates a new ValidationRunOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ValidationRunOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ValidationRunOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ValidationRunOut) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *ValidationRunOut) GetStatus() RunStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidationRunOut) GetStatusOk() (*RunStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidationRunOut) SetStatus(v RunStatus)`

SetStatus sets Status field to given value.


### GetTargetType

`func (o *ValidationRunOut) GetTargetType() TargetType`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ValidationRunOut) GetTargetTypeOk() (*TargetType, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ValidationRunOut) SetTargetType(v TargetType)`

SetTargetType sets TargetType field to given value.


### GetTargetId

`func (o *ValidationRunOut) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ValidationRunOut) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ValidationRunOut) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### SetTargetIdNil

`func (o *ValidationRunOut) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *ValidationRunOut) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetValidationsPassed

`func (o *ValidationRunOut) GetValidationsPassed() int32`

GetValidationsPassed returns the ValidationsPassed field if non-nil, zero value otherwise.

### GetValidationsPassedOk

`func (o *ValidationRunOut) GetValidationsPassedOk() (*int32, bool)`

GetValidationsPassedOk returns a tuple with the ValidationsPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationsPassed

`func (o *ValidationRunOut) SetValidationsPassed(v int32)`

SetValidationsPassed sets ValidationsPassed field to given value.


### GetValidationsTotal

`func (o *ValidationRunOut) GetValidationsTotal() int32`

GetValidationsTotal returns the ValidationsTotal field if non-nil, zero value otherwise.

### GetValidationsTotalOk

`func (o *ValidationRunOut) GetValidationsTotalOk() (*int32, bool)`

GetValidationsTotalOk returns a tuple with the ValidationsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationsTotal

`func (o *ValidationRunOut) SetValidationsTotal(v int32)`

SetValidationsTotal sets ValidationsTotal field to given value.


### GetStartedAt

`func (o *ValidationRunOut) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ValidationRunOut) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ValidationRunOut) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetFinishedAt

`func (o *ValidationRunOut) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ValidationRunOut) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ValidationRunOut) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.


### SetFinishedAtNil

`func (o *ValidationRunOut) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *ValidationRunOut) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetExpiresAt

`func (o *ValidationRunOut) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ValidationRunOut) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ValidationRunOut) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetValidations

`func (o *ValidationRunOut) GetValidations() []ValidationOut`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *ValidationRunOut) GetValidationsOk() (*[]ValidationOut, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *ValidationRunOut) SetValidations(v []ValidationOut)`

SetValidations sets Validations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


