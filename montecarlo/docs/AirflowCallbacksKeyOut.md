# AirflowCallbacksKeyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the key. | 
**Description** | **string** | What this key is for. | 
**CreatedTime** | **time.Time** | When the key was created. | 
**EtlContainerId** | Pointer to **NullableString** | The Airflow ETL container the key reports to. Null when the key names no warehouse or container, several of them, or one you cannot see. A key created outside this API with no scope at all is accepted for any of your account&#39;s resources, so replace it with a scoped one. | [optional] 
**McdId** | **string** | Key id to present, as the &#x60;x-mcd-id&#x60; header or a client&#39;s &#x60;mcd_id&#x60; setting. The same value as &#x60;id&#x60;. | 

## Methods

### NewAirflowCallbacksKeyOut

`func NewAirflowCallbacksKeyOut(id string, description string, createdTime time.Time, mcdId string, ) *AirflowCallbacksKeyOut`

NewAirflowCallbacksKeyOut instantiates a new AirflowCallbacksKeyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirflowCallbacksKeyOutWithDefaults

`func NewAirflowCallbacksKeyOutWithDefaults() *AirflowCallbacksKeyOut`

NewAirflowCallbacksKeyOutWithDefaults instantiates a new AirflowCallbacksKeyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AirflowCallbacksKeyOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AirflowCallbacksKeyOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AirflowCallbacksKeyOut) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *AirflowCallbacksKeyOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AirflowCallbacksKeyOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AirflowCallbacksKeyOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedTime

`func (o *AirflowCallbacksKeyOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AirflowCallbacksKeyOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AirflowCallbacksKeyOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetEtlContainerId

`func (o *AirflowCallbacksKeyOut) GetEtlContainerId() string`

GetEtlContainerId returns the EtlContainerId field if non-nil, zero value otherwise.

### GetEtlContainerIdOk

`func (o *AirflowCallbacksKeyOut) GetEtlContainerIdOk() (*string, bool)`

GetEtlContainerIdOk returns a tuple with the EtlContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlContainerId

`func (o *AirflowCallbacksKeyOut) SetEtlContainerId(v string)`

SetEtlContainerId sets EtlContainerId field to given value.

### HasEtlContainerId

`func (o *AirflowCallbacksKeyOut) HasEtlContainerId() bool`

HasEtlContainerId returns a boolean if a field has been set.

### SetEtlContainerIdNil

`func (o *AirflowCallbacksKeyOut) SetEtlContainerIdNil(b bool)`

 SetEtlContainerIdNil sets the value for EtlContainerId to be an explicit nil

### UnsetEtlContainerId
`func (o *AirflowCallbacksKeyOut) UnsetEtlContainerId()`

UnsetEtlContainerId ensures that no value is present for EtlContainerId, not even an explicit nil
### GetMcdId

`func (o *AirflowCallbacksKeyOut) GetMcdId() string`

GetMcdId returns the McdId field if non-nil, zero value otherwise.

### GetMcdIdOk

`func (o *AirflowCallbacksKeyOut) GetMcdIdOk() (*string, bool)`

GetMcdIdOk returns a tuple with the McdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcdId

`func (o *AirflowCallbacksKeyOut) SetMcdId(v string)`

SetMcdId sets McdId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


