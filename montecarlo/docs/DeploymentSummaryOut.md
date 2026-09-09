# DeploymentSummaryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **NullableTime** | When the deployment was assigned to your account. Null when Monte Carlo has no record of that. | [optional] 
**Enabled** | **bool** | Whether the deployment can serve connections. A deployment still waiting for its collection agent or data store to be registered, or with nothing provisioned on it, is not enabled. | 
**Id** | **string** | Unique identifier of the deployment. | 
**LastUpdatedTime** | Pointer to **NullableTime** | When Monte Carlo last updated the infrastructure behind the deployment. Null when Monte Carlo has no record of an update. | [optional] 
**Name** | **string** | Display name of the deployment. | 
**RuntimePlatform** | Pointer to [**NullableRuntimePlatform**](RuntimePlatform.md) | Where the deployment&#39;s collection agent or data store runs. Null for a Monte Carlo hosted deployment, which runs neither, for one with nothing provisioned on it, and for one whose platform Monte Carlo has not recorded. | [optional] 
**Type** | Pointer to [**NullableDeploymentType**](DeploymentType.md) | What the deployment hosts. Null when nothing is provisioned on it, in which case it has to be provisioned before it can be used. | [optional] 

## Methods

### NewDeploymentSummaryOut

`func NewDeploymentSummaryOut(enabled bool, id string, name string, ) *DeploymentSummaryOut`

NewDeploymentSummaryOut instantiates a new DeploymentSummaryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentSummaryOutWithDefaults

`func NewDeploymentSummaryOutWithDefaults() *DeploymentSummaryOut`

NewDeploymentSummaryOutWithDefaults instantiates a new DeploymentSummaryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *DeploymentSummaryOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DeploymentSummaryOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DeploymentSummaryOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DeploymentSummaryOut) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *DeploymentSummaryOut) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *DeploymentSummaryOut) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetEnabled

`func (o *DeploymentSummaryOut) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeploymentSummaryOut) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeploymentSummaryOut) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *DeploymentSummaryOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentSummaryOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentSummaryOut) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *DeploymentSummaryOut) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *DeploymentSummaryOut) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *DeploymentSummaryOut) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *DeploymentSummaryOut) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### SetLastUpdatedTimeNil

`func (o *DeploymentSummaryOut) SetLastUpdatedTimeNil(b bool)`

 SetLastUpdatedTimeNil sets the value for LastUpdatedTime to be an explicit nil

### UnsetLastUpdatedTime
`func (o *DeploymentSummaryOut) UnsetLastUpdatedTime()`

UnsetLastUpdatedTime ensures that no value is present for LastUpdatedTime, not even an explicit nil
### GetName

`func (o *DeploymentSummaryOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentSummaryOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentSummaryOut) SetName(v string)`

SetName sets Name field to given value.


### GetRuntimePlatform

`func (o *DeploymentSummaryOut) GetRuntimePlatform() RuntimePlatform`

GetRuntimePlatform returns the RuntimePlatform field if non-nil, zero value otherwise.

### GetRuntimePlatformOk

`func (o *DeploymentSummaryOut) GetRuntimePlatformOk() (*RuntimePlatform, bool)`

GetRuntimePlatformOk returns a tuple with the RuntimePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimePlatform

`func (o *DeploymentSummaryOut) SetRuntimePlatform(v RuntimePlatform)`

SetRuntimePlatform sets RuntimePlatform field to given value.

### HasRuntimePlatform

`func (o *DeploymentSummaryOut) HasRuntimePlatform() bool`

HasRuntimePlatform returns a boolean if a field has been set.

### SetRuntimePlatformNil

`func (o *DeploymentSummaryOut) SetRuntimePlatformNil(b bool)`

 SetRuntimePlatformNil sets the value for RuntimePlatform to be an explicit nil

### UnsetRuntimePlatform
`func (o *DeploymentSummaryOut) UnsetRuntimePlatform()`

UnsetRuntimePlatform ensures that no value is present for RuntimePlatform, not even an explicit nil
### GetType

`func (o *DeploymentSummaryOut) GetType() DeploymentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentSummaryOut) GetTypeOk() (*DeploymentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentSummaryOut) SetType(v DeploymentType)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentSummaryOut) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DeploymentSummaryOut) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DeploymentSummaryOut) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


