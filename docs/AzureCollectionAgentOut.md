# AzureCollectionAgentOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | [**AuthenticationType**](AuthenticationType.md) | How Monte Carlo authenticates when it calls the collection agent. | 
**CreatedTime** | Pointer to **NullableTime** | When the collection agent was created. That is when its deployment was provisioned, which is before you register the agent. | [optional] 
**DeploymentId** | **string** | Identifier of the deployment this collection agent runs on. | 
**Enabled** | **bool** | Whether Monte Carlo is using this collection agent. An agent Monte Carlo has not validated is not enabled, either because it has not been registered yet or because validation failed. | 
**FunctionAppUrl** | **string** | URL of the function app Monte Carlo calls. Empty until the agent has been registered. | 
**Id** | **string** | Unique identifier of the collection agent. | 
**ImageBuild** | Pointer to **NullableString** | Build of the image the collection agent is running. Null until Monte Carlo has contacted the agent. | [optional] 
**ImageVersion** | Pointer to **NullableString** | Version of the image the collection agent is running. Null until Monte Carlo has contacted the agent. | [optional] 
**IsRemoteUpgradeable** | **bool** | Whether Monte Carlo can update the collection agent&#39;s image for you. | 
**LastUpdatedTime** | Pointer to **NullableTime** | When the collection agent was last changed. Registering it, renaming it, changing how Monte Carlo reaches it, and Monte Carlo picking up a new image version all update this. Null until any of those has happened. | [optional] 
**Name** | Pointer to **NullableString** | Display name of the collection agent. Null when it has no name. | [optional] 

## Methods

### NewAzureCollectionAgentOut

`func NewAzureCollectionAgentOut(authenticationType AuthenticationType, deploymentId string, enabled bool, functionAppUrl string, id string, isRemoteUpgradeable bool, ) *AzureCollectionAgentOut`

NewAzureCollectionAgentOut instantiates a new AzureCollectionAgentOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCollectionAgentOutWithDefaults

`func NewAzureCollectionAgentOutWithDefaults() *AzureCollectionAgentOut`

NewAzureCollectionAgentOutWithDefaults instantiates a new AzureCollectionAgentOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *AzureCollectionAgentOut) GetAuthenticationType() AuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AzureCollectionAgentOut) GetAuthenticationTypeOk() (*AuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AzureCollectionAgentOut) SetAuthenticationType(v AuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetCreatedTime

`func (o *AzureCollectionAgentOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AzureCollectionAgentOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AzureCollectionAgentOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AzureCollectionAgentOut) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *AzureCollectionAgentOut) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *AzureCollectionAgentOut) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetDeploymentId

`func (o *AzureCollectionAgentOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AzureCollectionAgentOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AzureCollectionAgentOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetEnabled

`func (o *AzureCollectionAgentOut) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AzureCollectionAgentOut) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AzureCollectionAgentOut) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFunctionAppUrl

`func (o *AzureCollectionAgentOut) GetFunctionAppUrl() string`

GetFunctionAppUrl returns the FunctionAppUrl field if non-nil, zero value otherwise.

### GetFunctionAppUrlOk

`func (o *AzureCollectionAgentOut) GetFunctionAppUrlOk() (*string, bool)`

GetFunctionAppUrlOk returns a tuple with the FunctionAppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionAppUrl

`func (o *AzureCollectionAgentOut) SetFunctionAppUrl(v string)`

SetFunctionAppUrl sets FunctionAppUrl field to given value.


### GetId

`func (o *AzureCollectionAgentOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureCollectionAgentOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureCollectionAgentOut) SetId(v string)`

SetId sets Id field to given value.


### GetImageBuild

`func (o *AzureCollectionAgentOut) GetImageBuild() string`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *AzureCollectionAgentOut) GetImageBuildOk() (*string, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *AzureCollectionAgentOut) SetImageBuild(v string)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *AzureCollectionAgentOut) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### SetImageBuildNil

`func (o *AzureCollectionAgentOut) SetImageBuildNil(b bool)`

 SetImageBuildNil sets the value for ImageBuild to be an explicit nil

### UnsetImageBuild
`func (o *AzureCollectionAgentOut) UnsetImageBuild()`

UnsetImageBuild ensures that no value is present for ImageBuild, not even an explicit nil
### GetImageVersion

`func (o *AzureCollectionAgentOut) GetImageVersion() string`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *AzureCollectionAgentOut) GetImageVersionOk() (*string, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *AzureCollectionAgentOut) SetImageVersion(v string)`

SetImageVersion sets ImageVersion field to given value.

### HasImageVersion

`func (o *AzureCollectionAgentOut) HasImageVersion() bool`

HasImageVersion returns a boolean if a field has been set.

### SetImageVersionNil

`func (o *AzureCollectionAgentOut) SetImageVersionNil(b bool)`

 SetImageVersionNil sets the value for ImageVersion to be an explicit nil

### UnsetImageVersion
`func (o *AzureCollectionAgentOut) UnsetImageVersion()`

UnsetImageVersion ensures that no value is present for ImageVersion, not even an explicit nil
### GetIsRemoteUpgradeable

`func (o *AzureCollectionAgentOut) GetIsRemoteUpgradeable() bool`

GetIsRemoteUpgradeable returns the IsRemoteUpgradeable field if non-nil, zero value otherwise.

### GetIsRemoteUpgradeableOk

`func (o *AzureCollectionAgentOut) GetIsRemoteUpgradeableOk() (*bool, bool)`

GetIsRemoteUpgradeableOk returns a tuple with the IsRemoteUpgradeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteUpgradeable

`func (o *AzureCollectionAgentOut) SetIsRemoteUpgradeable(v bool)`

SetIsRemoteUpgradeable sets IsRemoteUpgradeable field to given value.


### GetLastUpdatedTime

`func (o *AzureCollectionAgentOut) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AzureCollectionAgentOut) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AzureCollectionAgentOut) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AzureCollectionAgentOut) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### SetLastUpdatedTimeNil

`func (o *AzureCollectionAgentOut) SetLastUpdatedTimeNil(b bool)`

 SetLastUpdatedTimeNil sets the value for LastUpdatedTime to be an explicit nil

### UnsetLastUpdatedTime
`func (o *AzureCollectionAgentOut) UnsetLastUpdatedTime()`

UnsetLastUpdatedTime ensures that no value is present for LastUpdatedTime, not even an explicit nil
### GetName

`func (o *AzureCollectionAgentOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureCollectionAgentOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureCollectionAgentOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureCollectionAgentOut) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureCollectionAgentOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureCollectionAgentOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


