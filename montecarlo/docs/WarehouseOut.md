# WarehouseOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the warehouse. | 
**Name** | **NullableString** | Display name of the warehouse. Null for a warehouse that was never named. | 
**Type** | [**WarehouseType**](WarehouseType.md) | The kind of data platform the warehouse represents. Fixed once created. | 
**DeploymentId** | **NullableString** | The deployment the warehouse&#39;s connections run through. Null for a warehouse that has no deployment. The id may name a deployment on Monte Carlo&#39;s older collection platform. The deployments endpoints do not list those. | 
**CreatedTime** | **time.Time** | When the warehouse was created. | 

## Methods

### NewWarehouseOut

`func NewWarehouseOut(id string, name NullableString, type_ WarehouseType, deploymentId NullableString, createdTime time.Time, ) *WarehouseOut`

NewWarehouseOut instantiates a new WarehouseOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseOutWithDefaults

`func NewWarehouseOutWithDefaults() *WarehouseOut`

NewWarehouseOutWithDefaults instantiates a new WarehouseOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WarehouseOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarehouseOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarehouseOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WarehouseOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WarehouseOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WarehouseOut) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *WarehouseOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WarehouseOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *WarehouseOut) GetType() WarehouseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WarehouseOut) GetTypeOk() (*WarehouseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WarehouseOut) SetType(v WarehouseType)`

SetType sets Type field to given value.


### GetDeploymentId

`func (o *WarehouseOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *WarehouseOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *WarehouseOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### SetDeploymentIdNil

`func (o *WarehouseOut) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *WarehouseOut) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetCreatedTime

`func (o *WarehouseOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *WarehouseOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *WarehouseOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


