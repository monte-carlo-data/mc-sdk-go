# BiContainerIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**NewBiContainerType**](NewBiContainerType.md) | The BI tool the container represents. Every connection added to it has to be for this tool. Cannot be changed after the container is created. | 
**Name** | **string** | Display name for the BI container. | 
**DeploymentId** | Pointer to **NullableString** | The deployment the container&#39;s connections will run through. Pick one from the deployments list. Only a deployment on Monte Carlo&#39;s current collection platform is accepted. &#x60;custom-bi-connector&#x60; takes the deployment of the agent that registered the connector, or none for a push-only connector. Every other type requires one. | [optional] 

## Methods

### NewBiContainerIn

`func NewBiContainerIn(type_ NewBiContainerType, name string, ) *BiContainerIn`

NewBiContainerIn instantiates a new BiContainerIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiContainerInWithDefaults

`func NewBiContainerInWithDefaults() *BiContainerIn`

NewBiContainerInWithDefaults instantiates a new BiContainerIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BiContainerIn) GetType() NewBiContainerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BiContainerIn) GetTypeOk() (*NewBiContainerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BiContainerIn) SetType(v NewBiContainerType)`

SetType sets Type field to given value.


### GetName

`func (o *BiContainerIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BiContainerIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BiContainerIn) SetName(v string)`

SetName sets Name field to given value.


### GetDeploymentId

`func (o *BiContainerIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *BiContainerIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *BiContainerIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *BiContainerIn) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *BiContainerIn) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *BiContainerIn) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


