# CollectionAgentOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | [**AuthenticationType**](AuthenticationType.md) | How Monte Carlo authenticates when it calls the collection agent. | 
**CreatedTime** | Pointer to **NullableTime** | When the collection agent was created. That is when its deployment was provisioned, which is before you register the agent. | [optional] 
**DeploymentId** | **string** | Identifier of the deployment this collection agent runs on. | 
**Enabled** | **bool** | Whether Monte Carlo is using this collection agent. An agent Monte Carlo has not validated is not enabled, either because it has not been registered yet or because validation failed. | 
**Endpoint** | **string** | Address Monte Carlo reaches the collection agent at, in whatever form its platform uses. On AWS that is the ARN of a Lambda function. Empty until the agent has been registered. | 
**Id** | **string** | Unique identifier of the collection agent. | 
**ImageBuild** | Pointer to **NullableString** | Build of the image the collection agent is running. Null until Monte Carlo has contacted the agent. | [optional] 
**ImageVersion** | Pointer to **NullableString** | Version of the image the collection agent is running. Null until Monte Carlo has contacted the agent. | [optional] 
**IsRemoteUpgradeable** | **bool** | Whether Monte Carlo can update the collection agent&#39;s image for you. | 
**LastUpdatedTime** | Pointer to **NullableTime** | When the collection agent was last changed. Registering it, changing its function or role, and Monte Carlo picking up a new image version all update this. Null until any of those has happened. | [optional] 
**Name** | Pointer to **NullableString** | Display name of the collection agent. Null when it has no name. | [optional] 
**Platform** | Pointer to [**NullableRuntimePlatform**](RuntimePlatform.md) | Where the collection agent runs. Use it to build the platform-specific path for any other operation on this agent. Null for an agent whose platform Monte Carlo has not recorded, and no platform-specific path can address one of those. | [optional] 

## Methods

### NewCollectionAgentOut

`func NewCollectionAgentOut(authenticationType AuthenticationType, deploymentId string, enabled bool, endpoint string, id string, isRemoteUpgradeable bool, ) *CollectionAgentOut`

NewCollectionAgentOut instantiates a new CollectionAgentOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionAgentOutWithDefaults

`func NewCollectionAgentOutWithDefaults() *CollectionAgentOut`

NewCollectionAgentOutWithDefaults instantiates a new CollectionAgentOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *CollectionAgentOut) GetAuthenticationType() AuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *CollectionAgentOut) GetAuthenticationTypeOk() (*AuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *CollectionAgentOut) SetAuthenticationType(v AuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetCreatedTime

`func (o *CollectionAgentOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CollectionAgentOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CollectionAgentOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CollectionAgentOut) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *CollectionAgentOut) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *CollectionAgentOut) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetDeploymentId

`func (o *CollectionAgentOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *CollectionAgentOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *CollectionAgentOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetEnabled

`func (o *CollectionAgentOut) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CollectionAgentOut) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CollectionAgentOut) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEndpoint

`func (o *CollectionAgentOut) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CollectionAgentOut) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CollectionAgentOut) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetId

`func (o *CollectionAgentOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionAgentOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionAgentOut) SetId(v string)`

SetId sets Id field to given value.


### GetImageBuild

`func (o *CollectionAgentOut) GetImageBuild() string`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *CollectionAgentOut) GetImageBuildOk() (*string, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *CollectionAgentOut) SetImageBuild(v string)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *CollectionAgentOut) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### SetImageBuildNil

`func (o *CollectionAgentOut) SetImageBuildNil(b bool)`

 SetImageBuildNil sets the value for ImageBuild to be an explicit nil

### UnsetImageBuild
`func (o *CollectionAgentOut) UnsetImageBuild()`

UnsetImageBuild ensures that no value is present for ImageBuild, not even an explicit nil
### GetImageVersion

`func (o *CollectionAgentOut) GetImageVersion() string`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *CollectionAgentOut) GetImageVersionOk() (*string, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *CollectionAgentOut) SetImageVersion(v string)`

SetImageVersion sets ImageVersion field to given value.

### HasImageVersion

`func (o *CollectionAgentOut) HasImageVersion() bool`

HasImageVersion returns a boolean if a field has been set.

### SetImageVersionNil

`func (o *CollectionAgentOut) SetImageVersionNil(b bool)`

 SetImageVersionNil sets the value for ImageVersion to be an explicit nil

### UnsetImageVersion
`func (o *CollectionAgentOut) UnsetImageVersion()`

UnsetImageVersion ensures that no value is present for ImageVersion, not even an explicit nil
### GetIsRemoteUpgradeable

`func (o *CollectionAgentOut) GetIsRemoteUpgradeable() bool`

GetIsRemoteUpgradeable returns the IsRemoteUpgradeable field if non-nil, zero value otherwise.

### GetIsRemoteUpgradeableOk

`func (o *CollectionAgentOut) GetIsRemoteUpgradeableOk() (*bool, bool)`

GetIsRemoteUpgradeableOk returns a tuple with the IsRemoteUpgradeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteUpgradeable

`func (o *CollectionAgentOut) SetIsRemoteUpgradeable(v bool)`

SetIsRemoteUpgradeable sets IsRemoteUpgradeable field to given value.


### GetLastUpdatedTime

`func (o *CollectionAgentOut) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CollectionAgentOut) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CollectionAgentOut) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *CollectionAgentOut) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### SetLastUpdatedTimeNil

`func (o *CollectionAgentOut) SetLastUpdatedTimeNil(b bool)`

 SetLastUpdatedTimeNil sets the value for LastUpdatedTime to be an explicit nil

### UnsetLastUpdatedTime
`func (o *CollectionAgentOut) UnsetLastUpdatedTime()`

UnsetLastUpdatedTime ensures that no value is present for LastUpdatedTime, not even an explicit nil
### GetName

`func (o *CollectionAgentOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionAgentOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionAgentOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CollectionAgentOut) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CollectionAgentOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CollectionAgentOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPlatform

`func (o *CollectionAgentOut) GetPlatform() RuntimePlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CollectionAgentOut) GetPlatformOk() (*RuntimePlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CollectionAgentOut) SetPlatform(v RuntimePlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CollectionAgentOut) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *CollectionAgentOut) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *CollectionAgentOut) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


