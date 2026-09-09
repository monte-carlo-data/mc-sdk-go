# DeploymentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Display name for the deployment. Monte Carlo generates one if you leave it out. | [optional] 
**RuntimePlatform** | [**RuntimePlatform**](RuntimePlatform.md) | Where the deployment&#39;s collection agent or data store will run. Both can be provisioned on &#x60;AWS&#x60; or &#x60;AZURE&#x60;. Any other value is rejected. | 
**Type** | [**DeploymentType**](DeploymentType.md) | What the deployment will host. Only &#x60;COLLECTION_AGENT&#x60; and &#x60;COLLECTION_DATA_STORE&#x60; can be provisioned today. Any other value is rejected. | 

## Methods

### NewDeploymentIn

`func NewDeploymentIn(runtimePlatform RuntimePlatform, type_ DeploymentType, ) *DeploymentIn`

NewDeploymentIn instantiates a new DeploymentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentInWithDefaults

`func NewDeploymentInWithDefaults() *DeploymentIn`

NewDeploymentInWithDefaults instantiates a new DeploymentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentIn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeploymentIn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeploymentIn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRuntimePlatform

`func (o *DeploymentIn) GetRuntimePlatform() RuntimePlatform`

GetRuntimePlatform returns the RuntimePlatform field if non-nil, zero value otherwise.

### GetRuntimePlatformOk

`func (o *DeploymentIn) GetRuntimePlatformOk() (*RuntimePlatform, bool)`

GetRuntimePlatformOk returns a tuple with the RuntimePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimePlatform

`func (o *DeploymentIn) SetRuntimePlatform(v RuntimePlatform)`

SetRuntimePlatform sets RuntimePlatform field to given value.


### GetType

`func (o *DeploymentIn) GetType() DeploymentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentIn) GetTypeOk() (*DeploymentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentIn) SetType(v DeploymentType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


