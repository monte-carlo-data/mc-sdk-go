# DeploymentOut

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
**AwsExternalId** | Pointer to **NullableString** | Value to supply when you register an AWS collection agent or data store on this deployment. It goes in the trust policy of the role Monte Carlo assumes. Null until Monte Carlo has generated one, for a deployment on another platform, for a caller who is not permitted to register one, and if the value could not be read just now. Retry the request in that last case. | [optional] 

## Methods

### NewDeploymentOut

`func NewDeploymentOut(id string, name string, enabled bool, status DeploymentStatus, deletionProtection bool, ) *DeploymentOut`

NewDeploymentOut instantiates a new DeploymentOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentOutWithDefaults

`func NewDeploymentOutWithDefaults() *DeploymentOut`

NewDeploymentOutWithDefaults instantiates a new DeploymentOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeploymentOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentOut) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DeploymentOut) GetType() DeploymentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentOut) GetTypeOk() (*DeploymentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentOut) SetType(v DeploymentType)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentOut) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DeploymentOut) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DeploymentOut) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetRuntimePlatform

`func (o *DeploymentOut) GetRuntimePlatform() RuntimePlatform`

GetRuntimePlatform returns the RuntimePlatform field if non-nil, zero value otherwise.

### GetRuntimePlatformOk

`func (o *DeploymentOut) GetRuntimePlatformOk() (*RuntimePlatform, bool)`

GetRuntimePlatformOk returns a tuple with the RuntimePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimePlatform

`func (o *DeploymentOut) SetRuntimePlatform(v RuntimePlatform)`

SetRuntimePlatform sets RuntimePlatform field to given value.

### HasRuntimePlatform

`func (o *DeploymentOut) HasRuntimePlatform() bool`

HasRuntimePlatform returns a boolean if a field has been set.

### SetRuntimePlatformNil

`func (o *DeploymentOut) SetRuntimePlatformNil(b bool)`

 SetRuntimePlatformNil sets the value for RuntimePlatform to be an explicit nil

### UnsetRuntimePlatform
`func (o *DeploymentOut) UnsetRuntimePlatform()`

UnsetRuntimePlatform ensures that no value is present for RuntimePlatform, not even an explicit nil
### GetEnabled

`func (o *DeploymentOut) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeploymentOut) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeploymentOut) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedTime

`func (o *DeploymentOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DeploymentOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DeploymentOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DeploymentOut) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *DeploymentOut) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *DeploymentOut) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetLastUpdatedTime

`func (o *DeploymentOut) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *DeploymentOut) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *DeploymentOut) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *DeploymentOut) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### SetLastUpdatedTimeNil

`func (o *DeploymentOut) SetLastUpdatedTimeNil(b bool)`

 SetLastUpdatedTimeNil sets the value for LastUpdatedTime to be an explicit nil

### UnsetLastUpdatedTime
`func (o *DeploymentOut) UnsetLastUpdatedTime()`

UnsetLastUpdatedTime ensures that no value is present for LastUpdatedTime, not even an explicit nil
### GetStatus

`func (o *DeploymentOut) GetStatus() DeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentOut) GetStatusOk() (*DeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentOut) SetStatus(v DeploymentStatus)`

SetStatus sets Status field to given value.


### GetStatusDetail

`func (o *DeploymentOut) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *DeploymentOut) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *DeploymentOut) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *DeploymentOut) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### SetStatusDetailNil

`func (o *DeploymentOut) SetStatusDetailNil(b bool)`

 SetStatusDetailNil sets the value for StatusDetail to be an explicit nil

### UnsetStatusDetail
`func (o *DeploymentOut) UnsetStatusDetail()`

UnsetStatusDetail ensures that no value is present for StatusDetail, not even an explicit nil
### GetDeletionProtection

`func (o *DeploymentOut) GetDeletionProtection() bool`

GetDeletionProtection returns the DeletionProtection field if non-nil, zero value otherwise.

### GetDeletionProtectionOk

`func (o *DeploymentOut) GetDeletionProtectionOk() (*bool, bool)`

GetDeletionProtectionOk returns a tuple with the DeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionProtection

`func (o *DeploymentOut) SetDeletionProtection(v bool)`

SetDeletionProtection sets DeletionProtection field to given value.


### GetAwsExternalId

`func (o *DeploymentOut) GetAwsExternalId() string`

GetAwsExternalId returns the AwsExternalId field if non-nil, zero value otherwise.

### GetAwsExternalIdOk

`func (o *DeploymentOut) GetAwsExternalIdOk() (*string, bool)`

GetAwsExternalIdOk returns a tuple with the AwsExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalId

`func (o *DeploymentOut) SetAwsExternalId(v string)`

SetAwsExternalId sets AwsExternalId field to given value.

### HasAwsExternalId

`func (o *DeploymentOut) HasAwsExternalId() bool`

HasAwsExternalId returns a boolean if a field has been set.

### SetAwsExternalIdNil

`func (o *DeploymentOut) SetAwsExternalIdNil(b bool)`

 SetAwsExternalIdNil sets the value for AwsExternalId to be an explicit nil

### UnsetAwsExternalId
`func (o *DeploymentOut) UnsetAwsExternalId()`

UnsetAwsExternalId ensures that no value is present for AwsExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


