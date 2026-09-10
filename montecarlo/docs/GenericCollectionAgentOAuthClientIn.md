# GenericCollectionAgentOAuthClientIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment whose generic collection agent will present this credential. It must have been provisioned for a generic collection agent. | 
**Description** | Pointer to **NullableString** | What this credential is for. Monte Carlo generates one naming the agent if you leave it out. | [optional] 
**ExpirationDays** | Pointer to **NullableInt32** | Days until the client stops being accepted. Leave it out for a client that does not expire. | [optional] 

## Methods

### NewGenericCollectionAgentOAuthClientIn

`func NewGenericCollectionAgentOAuthClientIn(deploymentId string, ) *GenericCollectionAgentOAuthClientIn`

NewGenericCollectionAgentOAuthClientIn instantiates a new GenericCollectionAgentOAuthClientIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericCollectionAgentOAuthClientInWithDefaults

`func NewGenericCollectionAgentOAuthClientInWithDefaults() *GenericCollectionAgentOAuthClientIn`

NewGenericCollectionAgentOAuthClientInWithDefaults instantiates a new GenericCollectionAgentOAuthClientIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *GenericCollectionAgentOAuthClientIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GenericCollectionAgentOAuthClientIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GenericCollectionAgentOAuthClientIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetDescription

`func (o *GenericCollectionAgentOAuthClientIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenericCollectionAgentOAuthClientIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenericCollectionAgentOAuthClientIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenericCollectionAgentOAuthClientIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GenericCollectionAgentOAuthClientIn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GenericCollectionAgentOAuthClientIn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpirationDays

`func (o *GenericCollectionAgentOAuthClientIn) GetExpirationDays() int32`

GetExpirationDays returns the ExpirationDays field if non-nil, zero value otherwise.

### GetExpirationDaysOk

`func (o *GenericCollectionAgentOAuthClientIn) GetExpirationDaysOk() (*int32, bool)`

GetExpirationDaysOk returns a tuple with the ExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDays

`func (o *GenericCollectionAgentOAuthClientIn) SetExpirationDays(v int32)`

SetExpirationDays sets ExpirationDays field to given value.

### HasExpirationDays

`func (o *GenericCollectionAgentOAuthClientIn) HasExpirationDays() bool`

HasExpirationDays returns a boolean if a field has been set.

### SetExpirationDaysNil

`func (o *GenericCollectionAgentOAuthClientIn) SetExpirationDaysNil(b bool)`

 SetExpirationDaysNil sets the value for ExpirationDays to be an explicit nil

### UnsetExpirationDays
`func (o *GenericCollectionAgentOAuthClientIn) UnsetExpirationDays()`

UnsetExpirationDays ensures that no value is present for ExpirationDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


