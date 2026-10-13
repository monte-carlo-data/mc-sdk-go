# IngestionKeyIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | What this key is for. Monte Carlo generates one naming the warehouse or container if you leave it out. | [optional] 
**WarehouseId** | Pointer to **NullableString** | The warehouse to push metadata, lineage and query logs to. Send exactly one of this, &#x60;bi_container_id&#x60; and &#x60;etl_container_id&#x60;. | [optional] 
**BiContainerId** | Pointer to **NullableString** | The BI container to push BI assets to, such as a &#x60;custom-bi-connector&#x60; container. Send exactly one of this, &#x60;warehouse_id&#x60; and &#x60;etl_container_id&#x60;. | [optional] 
**EtlContainerId** | Pointer to **NullableString** | The ETL container to push ETL jobs and runs to, such as a &#x60;custom-etl-connector&#x60; container. Send exactly one of this, &#x60;warehouse_id&#x60; and &#x60;bi_container_id&#x60;. | [optional] 

## Methods

### NewIngestionKeyIn

`func NewIngestionKeyIn() *IngestionKeyIn`

NewIngestionKeyIn instantiates a new IngestionKeyIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionKeyInWithDefaults

`func NewIngestionKeyInWithDefaults() *IngestionKeyIn`

NewIngestionKeyInWithDefaults instantiates a new IngestionKeyIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IngestionKeyIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IngestionKeyIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IngestionKeyIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IngestionKeyIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IngestionKeyIn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IngestionKeyIn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWarehouseId

`func (o *IngestionKeyIn) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *IngestionKeyIn) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *IngestionKeyIn) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *IngestionKeyIn) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### SetWarehouseIdNil

`func (o *IngestionKeyIn) SetWarehouseIdNil(b bool)`

 SetWarehouseIdNil sets the value for WarehouseId to be an explicit nil

### UnsetWarehouseId
`func (o *IngestionKeyIn) UnsetWarehouseId()`

UnsetWarehouseId ensures that no value is present for WarehouseId, not even an explicit nil
### GetBiContainerId

`func (o *IngestionKeyIn) GetBiContainerId() string`

GetBiContainerId returns the BiContainerId field if non-nil, zero value otherwise.

### GetBiContainerIdOk

`func (o *IngestionKeyIn) GetBiContainerIdOk() (*string, bool)`

GetBiContainerIdOk returns a tuple with the BiContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiContainerId

`func (o *IngestionKeyIn) SetBiContainerId(v string)`

SetBiContainerId sets BiContainerId field to given value.

### HasBiContainerId

`func (o *IngestionKeyIn) HasBiContainerId() bool`

HasBiContainerId returns a boolean if a field has been set.

### SetBiContainerIdNil

`func (o *IngestionKeyIn) SetBiContainerIdNil(b bool)`

 SetBiContainerIdNil sets the value for BiContainerId to be an explicit nil

### UnsetBiContainerId
`func (o *IngestionKeyIn) UnsetBiContainerId()`

UnsetBiContainerId ensures that no value is present for BiContainerId, not even an explicit nil
### GetEtlContainerId

`func (o *IngestionKeyIn) GetEtlContainerId() string`

GetEtlContainerId returns the EtlContainerId field if non-nil, zero value otherwise.

### GetEtlContainerIdOk

`func (o *IngestionKeyIn) GetEtlContainerIdOk() (*string, bool)`

GetEtlContainerIdOk returns a tuple with the EtlContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlContainerId

`func (o *IngestionKeyIn) SetEtlContainerId(v string)`

SetEtlContainerId sets EtlContainerId field to given value.

### HasEtlContainerId

`func (o *IngestionKeyIn) HasEtlContainerId() bool`

HasEtlContainerId returns a boolean if a field has been set.

### SetEtlContainerIdNil

`func (o *IngestionKeyIn) SetEtlContainerIdNil(b bool)`

 SetEtlContainerIdNil sets the value for EtlContainerId to be an explicit nil

### UnsetEtlContainerId
`func (o *IngestionKeyIn) UnsetEtlContainerId()`

UnsetEtlContainerId ensures that no value is present for EtlContainerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


