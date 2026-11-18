# DeploymentSummaryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the deployment. | 
**Name** | **string** | Display name of the deployment. | 
**Type** | Pointer to [**NullableDeploymentType**](DeploymentType.md) | What the deployment hosts. Null when nothing is provisioned on it, in which case it has to be provisioned before it can be used. | [optional] 
**RuntimePlatform** | Pointer to [**NullableRuntimePlatform**](RuntimePlatform.md) | Where the deployment&#39;s collection agent or data store runs. Null for a Monte Carlo hosted deployment, &#x60;AGENT_OBSERVABILITY&#x60; included, which runs neither. Also null for one with nothing provisioned on it, and for one whose platform Monte Carlo has not recorded. | [optional] 
**Enabled** | **bool** | Whether the deployment can serve connections. A deployment still waiting for its collection agent or data store to be registered, or with nothing provisioned on it, is not enabled. An &#x60;AGENT_OBSERVABILITY&#x60; deployment is enabled only while its &#x60;status&#x60; is &#x60;READY&#x60;. | 
**CreatedTime** | Pointer to **NullableTime** | When the deployment was assigned to your account, or for an &#x60;AGENT_OBSERVABILITY&#x60; deployment, when it was created. Null when Monte Carlo has no record of that. | [optional] 
**LastUpdatedTime** | Pointer to **NullableTime** | When Monte Carlo last updated the infrastructure behind the deployment. Null when Monte Carlo has no record of an update. | [optional] 
**Status** | [**DeploymentStatus**](DeploymentStatus.md) | How far Monte Carlo has got with setting the deployment up. An &#x60;AGENT_OBSERVABILITY&#x60; deployment is &#x60;PROVISIONING&#x60; while Monte Carlo sets it up, then &#x60;READY&#x60; or &#x60;FAILED&#x60;. Every other deployment is &#x60;READY&#x60;. &#x60;enabled&#x60; says whether the deployment can serve connections. | 
**StatusDetail** | Pointer to **NullableString** | More about the status, such as why setting up the deployment failed. Null when there is nothing to add. | [optional] 
**DeletionProtection** | **bool** | Whether deleting the deployment is refused once it has been &#x60;READY&#x60;. Only an &#x60;AGENT_OBSERVABILITY&#x60; deployment can be protected. Every other deployment is &#x60;false&#x60;. Turn it off with an update before deleting a protected deployment. | 

## Methods

### NewDeploymentSummaryOut

`func NewDeploymentSummaryOut(id string, name string, enabled bool, status DeploymentStatus, deletionProtection bool, ) *DeploymentSummaryOut`

NewDeploymentSummaryOut instantiates a new DeploymentSummaryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentSummaryOutWithDefaults

`func NewDeploymentSummaryOutWithDefaults() *DeploymentSummaryOut`

NewDeploymentSummaryOutWithDefaults instantiates a new DeploymentSummaryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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
### GetStatus

`func (o *DeploymentSummaryOut) GetStatus() DeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentSummaryOut) GetStatusOk() (*DeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentSummaryOut) SetStatus(v DeploymentStatus)`

SetStatus sets Status field to given value.


### GetStatusDetail

`func (o *DeploymentSummaryOut) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *DeploymentSummaryOut) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *DeploymentSummaryOut) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *DeploymentSummaryOut) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### SetStatusDetailNil

`func (o *DeploymentSummaryOut) SetStatusDetailNil(b bool)`

 SetStatusDetailNil sets the value for StatusDetail to be an explicit nil

### UnsetStatusDetail
`func (o *DeploymentSummaryOut) UnsetStatusDetail()`

UnsetStatusDetail ensures that no value is present for StatusDetail, not even an explicit nil
### GetDeletionProtection

`func (o *DeploymentSummaryOut) GetDeletionProtection() bool`

GetDeletionProtection returns the DeletionProtection field if non-nil, zero value otherwise.

### GetDeletionProtectionOk

`func (o *DeploymentSummaryOut) GetDeletionProtectionOk() (*bool, bool)`

GetDeletionProtectionOk returns a tuple with the DeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionProtection

`func (o *DeploymentSummaryOut) SetDeletionProtection(v bool)`

SetDeletionProtection sets DeletionProtection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


