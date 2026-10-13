# IngestionKeyCreatedOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the key. | 
**Description** | **string** | What this key is for. | 
**CreatedTime** | **time.Time** | When the key was created. | 
**WarehouseId** | Pointer to **NullableString** | The warehouse the key pushes metadata, lineage and query logs to. Null when the key is scoped to a container. | [optional] 
**BiContainerId** | Pointer to **NullableString** | The BI container the key pushes BI assets to. Null when the key is scoped to something else. | [optional] 
**EtlContainerId** | Pointer to **NullableString** | The ETL container the key pushes ETL jobs and runs to. Null when the key is scoped to something else. | [optional] 
**McdId** | **string** | Key id to present, as the &#x60;x-mcd-id&#x60; header or a client&#39;s &#x60;mcd_id&#x60; setting. The same value as &#x60;id&#x60;. | 
**McdToken** | **string** | Secret to present, as the &#x60;x-mcd-token&#x60; header or a client&#39;s &#x60;mcd_token&#x60; setting. Returned once, by this call, and not retrievable afterwards. Terraform keeps it in state, so anyone who can read the state can read the secret: &#x60;sensitive&#x60; only hides it from CLI output. Keep state in a backend that encrypts it and limits who can read it. Hand the secret on through a write-only argument, such as &#x60;aws_secretsmanager_secret_version.secret_string_wo&#x60;, not through an output, because a &#x60;sensitive&#x60; output is stored in state too. To rotate, create a new key and delete this one. | [readonly] 

## Methods

### NewIngestionKeyCreatedOut

`func NewIngestionKeyCreatedOut(id string, description string, createdTime time.Time, mcdId string, mcdToken string, ) *IngestionKeyCreatedOut`

NewIngestionKeyCreatedOut instantiates a new IngestionKeyCreatedOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionKeyCreatedOutWithDefaults

`func NewIngestionKeyCreatedOutWithDefaults() *IngestionKeyCreatedOut`

NewIngestionKeyCreatedOutWithDefaults instantiates a new IngestionKeyCreatedOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IngestionKeyCreatedOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestionKeyCreatedOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestionKeyCreatedOut) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *IngestionKeyCreatedOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IngestionKeyCreatedOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IngestionKeyCreatedOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedTime

`func (o *IngestionKeyCreatedOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IngestionKeyCreatedOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IngestionKeyCreatedOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetWarehouseId

`func (o *IngestionKeyCreatedOut) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *IngestionKeyCreatedOut) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *IngestionKeyCreatedOut) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *IngestionKeyCreatedOut) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### SetWarehouseIdNil

`func (o *IngestionKeyCreatedOut) SetWarehouseIdNil(b bool)`

 SetWarehouseIdNil sets the value for WarehouseId to be an explicit nil

### UnsetWarehouseId
`func (o *IngestionKeyCreatedOut) UnsetWarehouseId()`

UnsetWarehouseId ensures that no value is present for WarehouseId, not even an explicit nil
### GetBiContainerId

`func (o *IngestionKeyCreatedOut) GetBiContainerId() string`

GetBiContainerId returns the BiContainerId field if non-nil, zero value otherwise.

### GetBiContainerIdOk

`func (o *IngestionKeyCreatedOut) GetBiContainerIdOk() (*string, bool)`

GetBiContainerIdOk returns a tuple with the BiContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiContainerId

`func (o *IngestionKeyCreatedOut) SetBiContainerId(v string)`

SetBiContainerId sets BiContainerId field to given value.

### HasBiContainerId

`func (o *IngestionKeyCreatedOut) HasBiContainerId() bool`

HasBiContainerId returns a boolean if a field has been set.

### SetBiContainerIdNil

`func (o *IngestionKeyCreatedOut) SetBiContainerIdNil(b bool)`

 SetBiContainerIdNil sets the value for BiContainerId to be an explicit nil

### UnsetBiContainerId
`func (o *IngestionKeyCreatedOut) UnsetBiContainerId()`

UnsetBiContainerId ensures that no value is present for BiContainerId, not even an explicit nil
### GetEtlContainerId

`func (o *IngestionKeyCreatedOut) GetEtlContainerId() string`

GetEtlContainerId returns the EtlContainerId field if non-nil, zero value otherwise.

### GetEtlContainerIdOk

`func (o *IngestionKeyCreatedOut) GetEtlContainerIdOk() (*string, bool)`

GetEtlContainerIdOk returns a tuple with the EtlContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlContainerId

`func (o *IngestionKeyCreatedOut) SetEtlContainerId(v string)`

SetEtlContainerId sets EtlContainerId field to given value.

### HasEtlContainerId

`func (o *IngestionKeyCreatedOut) HasEtlContainerId() bool`

HasEtlContainerId returns a boolean if a field has been set.

### SetEtlContainerIdNil

`func (o *IngestionKeyCreatedOut) SetEtlContainerIdNil(b bool)`

 SetEtlContainerIdNil sets the value for EtlContainerId to be an explicit nil

### UnsetEtlContainerId
`func (o *IngestionKeyCreatedOut) UnsetEtlContainerId()`

UnsetEtlContainerId ensures that no value is present for EtlContainerId, not even an explicit nil
### GetMcdId

`func (o *IngestionKeyCreatedOut) GetMcdId() string`

GetMcdId returns the McdId field if non-nil, zero value otherwise.

### GetMcdIdOk

`func (o *IngestionKeyCreatedOut) GetMcdIdOk() (*string, bool)`

GetMcdIdOk returns a tuple with the McdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcdId

`func (o *IngestionKeyCreatedOut) SetMcdId(v string)`

SetMcdId sets McdId field to given value.


### GetMcdToken

`func (o *IngestionKeyCreatedOut) GetMcdToken() string`

GetMcdToken returns the McdToken field if non-nil, zero value otherwise.

### GetMcdTokenOk

`func (o *IngestionKeyCreatedOut) GetMcdTokenOk() (*string, bool)`

GetMcdTokenOk returns a tuple with the McdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcdToken

`func (o *IngestionKeyCreatedOut) SetMcdToken(v string)`

SetMcdToken sets McdToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


