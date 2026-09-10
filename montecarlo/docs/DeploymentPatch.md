# DeploymentPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | New display name for the deployment. Left out, the name is unchanged. | [optional] 

## Methods

### NewDeploymentPatch

`func NewDeploymentPatch() *DeploymentPatch`

NewDeploymentPatch instantiates a new DeploymentPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentPatchWithDefaults

`func NewDeploymentPatchWithDefaults() *DeploymentPatch`

NewDeploymentPatchWithDefaults instantiates a new DeploymentPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeploymentPatch) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeploymentPatch) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


