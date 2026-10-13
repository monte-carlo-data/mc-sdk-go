# AirflowCallbacksKeyIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | What this key is for. Monte Carlo generates one naming the warehouse or container if you leave it out. | [optional] 
**EtlContainerId** | **string** | The &#x60;airflow&#x60; ETL container the Airflow environment reports to. | 

## Methods

### NewAirflowCallbacksKeyIn

`func NewAirflowCallbacksKeyIn(etlContainerId string, ) *AirflowCallbacksKeyIn`

NewAirflowCallbacksKeyIn instantiates a new AirflowCallbacksKeyIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirflowCallbacksKeyInWithDefaults

`func NewAirflowCallbacksKeyInWithDefaults() *AirflowCallbacksKeyIn`

NewAirflowCallbacksKeyInWithDefaults instantiates a new AirflowCallbacksKeyIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AirflowCallbacksKeyIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AirflowCallbacksKeyIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AirflowCallbacksKeyIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AirflowCallbacksKeyIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AirflowCallbacksKeyIn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AirflowCallbacksKeyIn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEtlContainerId

`func (o *AirflowCallbacksKeyIn) GetEtlContainerId() string`

GetEtlContainerId returns the EtlContainerId field if non-nil, zero value otherwise.

### GetEtlContainerIdOk

`func (o *AirflowCallbacksKeyIn) GetEtlContainerIdOk() (*string, bool)`

GetEtlContainerIdOk returns a tuple with the EtlContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlContainerId

`func (o *AirflowCallbacksKeyIn) SetEtlContainerId(v string)`

SetEtlContainerId sets EtlContainerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


