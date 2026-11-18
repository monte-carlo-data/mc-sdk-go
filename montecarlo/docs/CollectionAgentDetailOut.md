# CollectionAgentDetailOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the collection agent. | 
**Name** | Pointer to **NullableString** | Display name of the collection agent. Null when it has no name. | [optional] 
**DeploymentId** | **string** | Identifier of the deployment this collection agent runs on. | 
**AuthenticationType** | [**NullableAuthenticationType**](AuthenticationType.md) | How Monte Carlo authenticates when it calls the collection agent. Null for an agent that connects out instead, such as a generic one. | 
**Enabled** | **bool** | Whether Monte Carlo is using this collection agent. An agent Monte Carlo has not validated is not enabled, either because it has not been registered yet or because validation failed. | 
**CreatedTime** | Pointer to **NullableTime** | When the collection agent was created. That is when its deployment was provisioned, which is before you register the agent. | [optional] 
**LastUpdatedTime** | Pointer to **NullableTime** | When the collection agent was last changed. Registering it, renaming it, changing how Monte Carlo reaches it, and Monte Carlo picking up a new image version all update this. Null until any of those has happened. | [optional] 
**ImageBuild** | Pointer to **NullableString** | Build of the image the collection agent is running. Null until Monte Carlo has contacted the agent. | [optional] 
**ImageVersion** | Pointer to **NullableString** | Version of the image the collection agent is running. Null until Monte Carlo has contacted the agent. | [optional] 
**IsRemoteUpgradeable** | **bool** | Whether Monte Carlo can update the collection agent&#39;s image for you. | 
**Platform** | [**RuntimePlatform**](RuntimePlatform.md) | Where the collection agent runs. | 
**Endpoint** | **NullableString** | Address Monte Carlo reaches the collection agent at, in whatever form its platform uses. On AWS that is the ARN of a Lambda function, on Azure the URL of a function app, and on GCP the URL of a Cloud Run service. Empty until the agent has been registered. Null for a generic or Snowflake agent, which connect to Monte Carlo rather than being reached. | 
**Details** | **map[string]interface{}** | The fields the platform-specific read of this agent returns beyond the ones above, such as &#x60;lambda_function_arn&#x60; and &#x60;external_id&#x60; on AWS. Empty for a generic agent. Never a secret. | 

## Methods

### NewCollectionAgentDetailOut

`func NewCollectionAgentDetailOut(id string, deploymentId string, authenticationType NullableAuthenticationType, enabled bool, isRemoteUpgradeable bool, platform RuntimePlatform, endpoint NullableString, details map[string]interface{}, ) *CollectionAgentDetailOut`

NewCollectionAgentDetailOut instantiates a new CollectionAgentDetailOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionAgentDetailOutWithDefaults

`func NewCollectionAgentDetailOutWithDefaults() *CollectionAgentDetailOut`

NewCollectionAgentDetailOutWithDefaults instantiates a new CollectionAgentDetailOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CollectionAgentDetailOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionAgentDetailOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionAgentDetailOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CollectionAgentDetailOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionAgentDetailOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionAgentDetailOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CollectionAgentDetailOut) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CollectionAgentDetailOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CollectionAgentDetailOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeploymentId

`func (o *CollectionAgentDetailOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *CollectionAgentDetailOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *CollectionAgentDetailOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetAuthenticationType

`func (o *CollectionAgentDetailOut) GetAuthenticationType() AuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *CollectionAgentDetailOut) GetAuthenticationTypeOk() (*AuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *CollectionAgentDetailOut) SetAuthenticationType(v AuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


### SetAuthenticationTypeNil

`func (o *CollectionAgentDetailOut) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *CollectionAgentDetailOut) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
### GetEnabled

`func (o *CollectionAgentDetailOut) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CollectionAgentDetailOut) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CollectionAgentDetailOut) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedTime

`func (o *CollectionAgentDetailOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CollectionAgentDetailOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CollectionAgentDetailOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CollectionAgentDetailOut) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *CollectionAgentDetailOut) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *CollectionAgentDetailOut) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetLastUpdatedTime

`func (o *CollectionAgentDetailOut) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CollectionAgentDetailOut) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CollectionAgentDetailOut) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *CollectionAgentDetailOut) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### SetLastUpdatedTimeNil

`func (o *CollectionAgentDetailOut) SetLastUpdatedTimeNil(b bool)`

 SetLastUpdatedTimeNil sets the value for LastUpdatedTime to be an explicit nil

### UnsetLastUpdatedTime
`func (o *CollectionAgentDetailOut) UnsetLastUpdatedTime()`

UnsetLastUpdatedTime ensures that no value is present for LastUpdatedTime, not even an explicit nil
### GetImageBuild

`func (o *CollectionAgentDetailOut) GetImageBuild() string`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *CollectionAgentDetailOut) GetImageBuildOk() (*string, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *CollectionAgentDetailOut) SetImageBuild(v string)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *CollectionAgentDetailOut) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### SetImageBuildNil

`func (o *CollectionAgentDetailOut) SetImageBuildNil(b bool)`

 SetImageBuildNil sets the value for ImageBuild to be an explicit nil

### UnsetImageBuild
`func (o *CollectionAgentDetailOut) UnsetImageBuild()`

UnsetImageBuild ensures that no value is present for ImageBuild, not even an explicit nil
### GetImageVersion

`func (o *CollectionAgentDetailOut) GetImageVersion() string`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *CollectionAgentDetailOut) GetImageVersionOk() (*string, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *CollectionAgentDetailOut) SetImageVersion(v string)`

SetImageVersion sets ImageVersion field to given value.

### HasImageVersion

`func (o *CollectionAgentDetailOut) HasImageVersion() bool`

HasImageVersion returns a boolean if a field has been set.

### SetImageVersionNil

`func (o *CollectionAgentDetailOut) SetImageVersionNil(b bool)`

 SetImageVersionNil sets the value for ImageVersion to be an explicit nil

### UnsetImageVersion
`func (o *CollectionAgentDetailOut) UnsetImageVersion()`

UnsetImageVersion ensures that no value is present for ImageVersion, not even an explicit nil
### GetIsRemoteUpgradeable

`func (o *CollectionAgentDetailOut) GetIsRemoteUpgradeable() bool`

GetIsRemoteUpgradeable returns the IsRemoteUpgradeable field if non-nil, zero value otherwise.

### GetIsRemoteUpgradeableOk

`func (o *CollectionAgentDetailOut) GetIsRemoteUpgradeableOk() (*bool, bool)`

GetIsRemoteUpgradeableOk returns a tuple with the IsRemoteUpgradeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteUpgradeable

`func (o *CollectionAgentDetailOut) SetIsRemoteUpgradeable(v bool)`

SetIsRemoteUpgradeable sets IsRemoteUpgradeable field to given value.


### GetPlatform

`func (o *CollectionAgentDetailOut) GetPlatform() RuntimePlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CollectionAgentDetailOut) GetPlatformOk() (*RuntimePlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CollectionAgentDetailOut) SetPlatform(v RuntimePlatform)`

SetPlatform sets Platform field to given value.


### GetEndpoint

`func (o *CollectionAgentDetailOut) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CollectionAgentDetailOut) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CollectionAgentDetailOut) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### SetEndpointNil

`func (o *CollectionAgentDetailOut) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *CollectionAgentDetailOut) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetDetails

`func (o *CollectionAgentDetailOut) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CollectionAgentDetailOut) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CollectionAgentDetailOut) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


