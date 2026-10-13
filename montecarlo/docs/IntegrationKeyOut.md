# IntegrationKeyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the key. | 
**Description** | **string** | What this key is for. | 
**CreatedTime** | **time.Time** | When the key was created. | 
**Scope** | [**IntegrationKeyScope**](IntegrationKeyScope.md) | What the key lets its holder do. | 
**WarehouseId** | Pointer to **NullableString** | The warehouse the key is scoped to, if it is scoped to one. | [optional] 
**BiContainerId** | Pointer to **NullableString** | The BI container the key is scoped to, if it is scoped to one. | [optional] 
**EtlContainerId** | Pointer to **NullableString** | The ETL container the key is scoped to, if it is scoped to one. | [optional] 

## Methods

### NewIntegrationKeyOut

`func NewIntegrationKeyOut(id string, description string, createdTime time.Time, scope IntegrationKeyScope, ) *IntegrationKeyOut`

NewIntegrationKeyOut instantiates a new IntegrationKeyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationKeyOutWithDefaults

`func NewIntegrationKeyOutWithDefaults() *IntegrationKeyOut`

NewIntegrationKeyOutWithDefaults instantiates a new IntegrationKeyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationKeyOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationKeyOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationKeyOut) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *IntegrationKeyOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationKeyOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationKeyOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedTime

`func (o *IntegrationKeyOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IntegrationKeyOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IntegrationKeyOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetScope

`func (o *IntegrationKeyOut) GetScope() IntegrationKeyScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IntegrationKeyOut) GetScopeOk() (*IntegrationKeyScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IntegrationKeyOut) SetScope(v IntegrationKeyScope)`

SetScope sets Scope field to given value.


### GetWarehouseId

`func (o *IntegrationKeyOut) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *IntegrationKeyOut) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *IntegrationKeyOut) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *IntegrationKeyOut) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### SetWarehouseIdNil

`func (o *IntegrationKeyOut) SetWarehouseIdNil(b bool)`

 SetWarehouseIdNil sets the value for WarehouseId to be an explicit nil

### UnsetWarehouseId
`func (o *IntegrationKeyOut) UnsetWarehouseId()`

UnsetWarehouseId ensures that no value is present for WarehouseId, not even an explicit nil
### GetBiContainerId

`func (o *IntegrationKeyOut) GetBiContainerId() string`

GetBiContainerId returns the BiContainerId field if non-nil, zero value otherwise.

### GetBiContainerIdOk

`func (o *IntegrationKeyOut) GetBiContainerIdOk() (*string, bool)`

GetBiContainerIdOk returns a tuple with the BiContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiContainerId

`func (o *IntegrationKeyOut) SetBiContainerId(v string)`

SetBiContainerId sets BiContainerId field to given value.

### HasBiContainerId

`func (o *IntegrationKeyOut) HasBiContainerId() bool`

HasBiContainerId returns a boolean if a field has been set.

### SetBiContainerIdNil

`func (o *IntegrationKeyOut) SetBiContainerIdNil(b bool)`

 SetBiContainerIdNil sets the value for BiContainerId to be an explicit nil

### UnsetBiContainerId
`func (o *IntegrationKeyOut) UnsetBiContainerId()`

UnsetBiContainerId ensures that no value is present for BiContainerId, not even an explicit nil
### GetEtlContainerId

`func (o *IntegrationKeyOut) GetEtlContainerId() string`

GetEtlContainerId returns the EtlContainerId field if non-nil, zero value otherwise.

### GetEtlContainerIdOk

`func (o *IntegrationKeyOut) GetEtlContainerIdOk() (*string, bool)`

GetEtlContainerIdOk returns a tuple with the EtlContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlContainerId

`func (o *IntegrationKeyOut) SetEtlContainerId(v string)`

SetEtlContainerId sets EtlContainerId field to given value.

### HasEtlContainerId

`func (o *IntegrationKeyOut) HasEtlContainerId() bool`

HasEtlContainerId returns a boolean if a field has been set.

### SetEtlContainerIdNil

`func (o *IntegrationKeyOut) SetEtlContainerIdNil(b bool)`

 SetEtlContainerIdNil sets the value for EtlContainerId to be an explicit nil

### UnsetEtlContainerId
`func (o *IntegrationKeyOut) UnsetEtlContainerId()`

UnsetEtlContainerId ensures that no value is present for EtlContainerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


