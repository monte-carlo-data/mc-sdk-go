# EtlContainerIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**NewEtlContainerType**](NewEtlContainerType.md) | The ETL tool the container represents. Its connection has to be for this tool. Cannot be changed after the container is created. | 
**Name** | **string** | Display name for the ETL container. No two containers of the same type can share a name. | 
**DeploymentId** | Pointer to **NullableString** | The deployment the container&#39;s connection will run through. Pick one from the deployments list. Only a deployment on Monte Carlo&#39;s current collection platform is accepted. &#x60;airflow&#x60; takes none. &#x60;custom-etl-connector&#x60; takes the deployment of the agent that registered the connector, or none for a push-only connector. Every other type requires one. | [optional] 

## Methods

### NewEtlContainerIn

`func NewEtlContainerIn(type_ NewEtlContainerType, name string, ) *EtlContainerIn`

NewEtlContainerIn instantiates a new EtlContainerIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtlContainerInWithDefaults

`func NewEtlContainerInWithDefaults() *EtlContainerIn`

NewEtlContainerInWithDefaults instantiates a new EtlContainerIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EtlContainerIn) GetType() NewEtlContainerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EtlContainerIn) GetTypeOk() (*NewEtlContainerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EtlContainerIn) SetType(v NewEtlContainerType)`

SetType sets Type field to given value.


### GetName

`func (o *EtlContainerIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EtlContainerIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EtlContainerIn) SetName(v string)`

SetName sets Name field to given value.


### GetDeploymentId

`func (o *EtlContainerIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *EtlContainerIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *EtlContainerIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *EtlContainerIn) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *EtlContainerIn) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *EtlContainerIn) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


