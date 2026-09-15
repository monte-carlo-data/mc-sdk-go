# WarehouseIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for the warehouse. Unique among your warehouses of the same type. | 
**Type** | [**WarehouseType**](WarehouseType.md) | The kind of data platform the warehouse represents. Every connection added to it has to fit. Cannot be changed after the warehouse is created. | 
**DeploymentId** | **string** | The deployment the warehouse&#39;s connections will run through. Pick one from the deployments list. Only a deployment on Monte Carlo&#39;s current collection platform is accepted. | 

## Methods

### NewWarehouseIn

`func NewWarehouseIn(name string, type_ WarehouseType, deploymentId string, ) *WarehouseIn`

NewWarehouseIn instantiates a new WarehouseIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseInWithDefaults

`func NewWarehouseInWithDefaults() *WarehouseIn`

NewWarehouseInWithDefaults instantiates a new WarehouseIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WarehouseIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WarehouseIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WarehouseIn) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WarehouseIn) GetType() WarehouseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WarehouseIn) GetTypeOk() (*WarehouseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WarehouseIn) SetType(v WarehouseType)`

SetType sets Type field to given value.


### GetDeploymentId

`func (o *WarehouseIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *WarehouseIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *WarehouseIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


