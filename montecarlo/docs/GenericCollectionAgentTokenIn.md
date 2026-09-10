# GenericCollectionAgentTokenIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment whose generic collection agent will present this credential. It must have been provisioned for a generic collection agent. | 
**Description** | Pointer to **NullableString** | What this credential is for. Monte Carlo generates one naming the agent if you leave it out. | [optional] 

## Methods

### NewGenericCollectionAgentTokenIn

`func NewGenericCollectionAgentTokenIn(deploymentId string, ) *GenericCollectionAgentTokenIn`

NewGenericCollectionAgentTokenIn instantiates a new GenericCollectionAgentTokenIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericCollectionAgentTokenInWithDefaults

`func NewGenericCollectionAgentTokenInWithDefaults() *GenericCollectionAgentTokenIn`

NewGenericCollectionAgentTokenInWithDefaults instantiates a new GenericCollectionAgentTokenIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *GenericCollectionAgentTokenIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GenericCollectionAgentTokenIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GenericCollectionAgentTokenIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetDescription

`func (o *GenericCollectionAgentTokenIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenericCollectionAgentTokenIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenericCollectionAgentTokenIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenericCollectionAgentTokenIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GenericCollectionAgentTokenIn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GenericCollectionAgentTokenIn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


