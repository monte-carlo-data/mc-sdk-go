# AwsCollectionAgentOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | [**NullableAuthenticationType**](AuthenticationType.md) | How Monte Carlo authenticates when it calls the collection agent. Null for an agent that connects out instead, such as a generic one. | 
**CreatedTime** | Pointer to **NullableTime** | When the collection agent was created. That is when its deployment was provisioned, which is before you register the agent. | [optional] 
**DeploymentId** | **string** | Identifier of the deployment this collection agent runs on. | 
**Enabled** | **bool** | Whether Monte Carlo is using this collection agent. An agent Monte Carlo has not validated is not enabled, either because it has not been registered yet or because validation failed. | 
**ExternalId** | Pointer to **NullableString** | Value to supply in the trust policy of the role Monte Carlo assumes to invoke the function. Null until Monte Carlo has generated one, for a caller who is not permitted to register an agent, and if the value could not be read just now. Retry the request in that last case. | [optional] 
**Id** | **string** | Unique identifier of the collection agent. | 
**ImageBuild** | Pointer to **NullableString** | Build of the image the collection agent is running. Null until Monte Carlo has contacted the agent. | [optional] 
**ImageVersion** | Pointer to **NullableString** | Version of the image the collection agent is running. Null until Monte Carlo has contacted the agent. | [optional] 
**IsRemoteUpgradeable** | **bool** | Whether Monte Carlo can update the collection agent&#39;s image for you. | 
**LambdaFunctionArn** | **string** | ARN of the Lambda function Monte Carlo invokes. Empty until the agent has been registered. | 
**LastUpdatedTime** | Pointer to **NullableTime** | When the collection agent was last changed. Registering it, renaming it, changing how Monte Carlo reaches it, and Monte Carlo picking up a new image version all update this. Null until any of those has happened. | [optional] 
**Name** | Pointer to **NullableString** | Display name of the collection agent. Null when it has no name. | [optional] 

## Methods

### NewAwsCollectionAgentOut

`func NewAwsCollectionAgentOut(authenticationType NullableAuthenticationType, deploymentId string, enabled bool, id string, isRemoteUpgradeable bool, lambdaFunctionArn string, ) *AwsCollectionAgentOut`

NewAwsCollectionAgentOut instantiates a new AwsCollectionAgentOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCollectionAgentOutWithDefaults

`func NewAwsCollectionAgentOutWithDefaults() *AwsCollectionAgentOut`

NewAwsCollectionAgentOutWithDefaults instantiates a new AwsCollectionAgentOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *AwsCollectionAgentOut) GetAuthenticationType() AuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AwsCollectionAgentOut) GetAuthenticationTypeOk() (*AuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AwsCollectionAgentOut) SetAuthenticationType(v AuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


### SetAuthenticationTypeNil

`func (o *AwsCollectionAgentOut) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *AwsCollectionAgentOut) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
### GetCreatedTime

`func (o *AwsCollectionAgentOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AwsCollectionAgentOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AwsCollectionAgentOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AwsCollectionAgentOut) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *AwsCollectionAgentOut) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *AwsCollectionAgentOut) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetDeploymentId

`func (o *AwsCollectionAgentOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AwsCollectionAgentOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AwsCollectionAgentOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetEnabled

`func (o *AwsCollectionAgentOut) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AwsCollectionAgentOut) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AwsCollectionAgentOut) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExternalId

`func (o *AwsCollectionAgentOut) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsCollectionAgentOut) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsCollectionAgentOut) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AwsCollectionAgentOut) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AwsCollectionAgentOut) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AwsCollectionAgentOut) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetId

`func (o *AwsCollectionAgentOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsCollectionAgentOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsCollectionAgentOut) SetId(v string)`

SetId sets Id field to given value.


### GetImageBuild

`func (o *AwsCollectionAgentOut) GetImageBuild() string`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *AwsCollectionAgentOut) GetImageBuildOk() (*string, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *AwsCollectionAgentOut) SetImageBuild(v string)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *AwsCollectionAgentOut) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### SetImageBuildNil

`func (o *AwsCollectionAgentOut) SetImageBuildNil(b bool)`

 SetImageBuildNil sets the value for ImageBuild to be an explicit nil

### UnsetImageBuild
`func (o *AwsCollectionAgentOut) UnsetImageBuild()`

UnsetImageBuild ensures that no value is present for ImageBuild, not even an explicit nil
### GetImageVersion

`func (o *AwsCollectionAgentOut) GetImageVersion() string`

GetImageVersion returns the ImageVersion field if non-nil, zero value otherwise.

### GetImageVersionOk

`func (o *AwsCollectionAgentOut) GetImageVersionOk() (*string, bool)`

GetImageVersionOk returns a tuple with the ImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageVersion

`func (o *AwsCollectionAgentOut) SetImageVersion(v string)`

SetImageVersion sets ImageVersion field to given value.

### HasImageVersion

`func (o *AwsCollectionAgentOut) HasImageVersion() bool`

HasImageVersion returns a boolean if a field has been set.

### SetImageVersionNil

`func (o *AwsCollectionAgentOut) SetImageVersionNil(b bool)`

 SetImageVersionNil sets the value for ImageVersion to be an explicit nil

### UnsetImageVersion
`func (o *AwsCollectionAgentOut) UnsetImageVersion()`

UnsetImageVersion ensures that no value is present for ImageVersion, not even an explicit nil
### GetIsRemoteUpgradeable

`func (o *AwsCollectionAgentOut) GetIsRemoteUpgradeable() bool`

GetIsRemoteUpgradeable returns the IsRemoteUpgradeable field if non-nil, zero value otherwise.

### GetIsRemoteUpgradeableOk

`func (o *AwsCollectionAgentOut) GetIsRemoteUpgradeableOk() (*bool, bool)`

GetIsRemoteUpgradeableOk returns a tuple with the IsRemoteUpgradeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteUpgradeable

`func (o *AwsCollectionAgentOut) SetIsRemoteUpgradeable(v bool)`

SetIsRemoteUpgradeable sets IsRemoteUpgradeable field to given value.


### GetLambdaFunctionArn

`func (o *AwsCollectionAgentOut) GetLambdaFunctionArn() string`

GetLambdaFunctionArn returns the LambdaFunctionArn field if non-nil, zero value otherwise.

### GetLambdaFunctionArnOk

`func (o *AwsCollectionAgentOut) GetLambdaFunctionArnOk() (*string, bool)`

GetLambdaFunctionArnOk returns a tuple with the LambdaFunctionArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaFunctionArn

`func (o *AwsCollectionAgentOut) SetLambdaFunctionArn(v string)`

SetLambdaFunctionArn sets LambdaFunctionArn field to given value.


### GetLastUpdatedTime

`func (o *AwsCollectionAgentOut) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AwsCollectionAgentOut) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AwsCollectionAgentOut) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AwsCollectionAgentOut) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### SetLastUpdatedTimeNil

`func (o *AwsCollectionAgentOut) SetLastUpdatedTimeNil(b bool)`

 SetLastUpdatedTimeNil sets the value for LastUpdatedTime to be an explicit nil

### UnsetLastUpdatedTime
`func (o *AwsCollectionAgentOut) UnsetLastUpdatedTime()`

UnsetLastUpdatedTime ensures that no value is present for LastUpdatedTime, not even an explicit nil
### GetName

`func (o *AwsCollectionAgentOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsCollectionAgentOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsCollectionAgentOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsCollectionAgentOut) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsCollectionAgentOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsCollectionAgentOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


