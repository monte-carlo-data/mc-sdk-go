# GenericCollectionAgentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment whose generic collection agent to enable. It must have been provisioned for one, and the agent must be running with a credential created for this deployment. | 
**Name** | Pointer to **NullableString** | Display name for the collection agent. Replaces the name it currently has. | [optional] 

## Methods

### NewGenericCollectionAgentIn

`func NewGenericCollectionAgentIn(deploymentId string, ) *GenericCollectionAgentIn`

NewGenericCollectionAgentIn instantiates a new GenericCollectionAgentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericCollectionAgentInWithDefaults

`func NewGenericCollectionAgentInWithDefaults() *GenericCollectionAgentIn`

NewGenericCollectionAgentInWithDefaults instantiates a new GenericCollectionAgentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *GenericCollectionAgentIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GenericCollectionAgentIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GenericCollectionAgentIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetName

`func (o *GenericCollectionAgentIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericCollectionAgentIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericCollectionAgentIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenericCollectionAgentIn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GenericCollectionAgentIn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GenericCollectionAgentIn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


