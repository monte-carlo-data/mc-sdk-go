# DeploymentReprovisionIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**DeploymentType**](DeploymentType.md) | What the deployment will host. Only &#x60;COLLECTION_AGENT&#x60; and &#x60;COLLECTION_DATA_STORE&#x60; can be provisioned today. Any other value is rejected. | 
**RuntimePlatform** | [**RuntimePlatform**](RuntimePlatform.md) | Where the deployment&#39;s collection agent or data store will run. Either can be provisioned on &#x60;AWS&#x60;, &#x60;AZURE&#x60; or &#x60;GCP&#x60;, and a collection agent also on &#x60;GENERIC&#x60;. Any other combination is rejected. | 
**Name** | Pointer to **NullableString** | Display name for the deployment. Monte Carlo generates one if you leave it out. | [optional] 

## Methods

### NewDeploymentReprovisionIn

`func NewDeploymentReprovisionIn(type_ DeploymentType, runtimePlatform RuntimePlatform, ) *DeploymentReprovisionIn`

NewDeploymentReprovisionIn instantiates a new DeploymentReprovisionIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentReprovisionInWithDefaults

`func NewDeploymentReprovisionInWithDefaults() *DeploymentReprovisionIn`

NewDeploymentReprovisionInWithDefaults instantiates a new DeploymentReprovisionIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeploymentReprovisionIn) GetType() DeploymentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentReprovisionIn) GetTypeOk() (*DeploymentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentReprovisionIn) SetType(v DeploymentType)`

SetType sets Type field to given value.


### GetRuntimePlatform

`func (o *DeploymentReprovisionIn) GetRuntimePlatform() RuntimePlatform`

GetRuntimePlatform returns the RuntimePlatform field if non-nil, zero value otherwise.

### GetRuntimePlatformOk

`func (o *DeploymentReprovisionIn) GetRuntimePlatformOk() (*RuntimePlatform, bool)`

GetRuntimePlatformOk returns a tuple with the RuntimePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimePlatform

`func (o *DeploymentReprovisionIn) SetRuntimePlatform(v RuntimePlatform)`

SetRuntimePlatform sets RuntimePlatform field to given value.


### GetName

`func (o *DeploymentReprovisionIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentReprovisionIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentReprovisionIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentReprovisionIn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeploymentReprovisionIn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeploymentReprovisionIn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


