# CustomConnectorTerminology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | **string** | Noun for the job tier. | 
**Task** | **NullableString** | Noun for the task tier. Null if the tool has none. | 
**Group** | **NullableString** | Noun for the group tier above jobs. Null if the tool has none. | 

## Methods

### NewCustomConnectorTerminology

`func NewCustomConnectorTerminology(job string, task NullableString, group NullableString, ) *CustomConnectorTerminology`

NewCustomConnectorTerminology instantiates a new CustomConnectorTerminology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomConnectorTerminologyWithDefaults

`func NewCustomConnectorTerminologyWithDefaults() *CustomConnectorTerminology`

NewCustomConnectorTerminologyWithDefaults instantiates a new CustomConnectorTerminology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *CustomConnectorTerminology) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *CustomConnectorTerminology) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *CustomConnectorTerminology) SetJob(v string)`

SetJob sets Job field to given value.


### GetTask

`func (o *CustomConnectorTerminology) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *CustomConnectorTerminology) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *CustomConnectorTerminology) SetTask(v string)`

SetTask sets Task field to given value.


### SetTaskNil

`func (o *CustomConnectorTerminology) SetTaskNil(b bool)`

 SetTaskNil sets the value for Task to be an explicit nil

### UnsetTask
`func (o *CustomConnectorTerminology) UnsetTask()`

UnsetTask ensures that no value is present for Task, not even an explicit nil
### GetGroup

`func (o *CustomConnectorTerminology) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CustomConnectorTerminology) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CustomConnectorTerminology) SetGroup(v string)`

SetGroup sets Group field to given value.


### SetGroupNil

`func (o *CustomConnectorTerminology) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *CustomConnectorTerminology) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


