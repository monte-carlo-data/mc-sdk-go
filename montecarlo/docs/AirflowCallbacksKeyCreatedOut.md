# AirflowCallbacksKeyCreatedOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the key. | 
**Description** | **string** | What this key is for. | 
**CreatedTime** | **time.Time** | When the key was created. | 
**EtlContainerId** | Pointer to **NullableString** | The Airflow ETL container the key reports to. Null when the key names no warehouse or container, several of them, or one you cannot see. A key created outside this API with no scope at all is accepted for any of your account&#39;s resources, so replace it with a scoped one. | [optional] 
**McdId** | **string** | Key id to present, as the &#x60;x-mcd-id&#x60; header or a client&#39;s &#x60;mcd_id&#x60; setting. The same value as &#x60;id&#x60;. | 
**McdToken** | **string** | Secret to present, as the &#x60;x-mcd-token&#x60; header or a client&#39;s &#x60;mcd_token&#x60; setting. Returned once, by this call, and not retrievable afterwards. Terraform keeps it in state, so anyone who can read the state can read the secret: &#x60;sensitive&#x60; only hides it from CLI output. Keep state in a backend that encrypts it and limits who can read it. Hand the secret on through a write-only argument, such as &#x60;aws_secretsmanager_secret_version.secret_string_wo&#x60;, not through an output, because a &#x60;sensitive&#x60; output is stored in state too. To rotate, create a new key and delete this one. | [readonly] 

## Methods

### NewAirflowCallbacksKeyCreatedOut

`func NewAirflowCallbacksKeyCreatedOut(id string, description string, createdTime time.Time, mcdId string, mcdToken string, ) *AirflowCallbacksKeyCreatedOut`

NewAirflowCallbacksKeyCreatedOut instantiates a new AirflowCallbacksKeyCreatedOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirflowCallbacksKeyCreatedOutWithDefaults

`func NewAirflowCallbacksKeyCreatedOutWithDefaults() *AirflowCallbacksKeyCreatedOut`

NewAirflowCallbacksKeyCreatedOutWithDefaults instantiates a new AirflowCallbacksKeyCreatedOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AirflowCallbacksKeyCreatedOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AirflowCallbacksKeyCreatedOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AirflowCallbacksKeyCreatedOut) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *AirflowCallbacksKeyCreatedOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AirflowCallbacksKeyCreatedOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AirflowCallbacksKeyCreatedOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedTime

`func (o *AirflowCallbacksKeyCreatedOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AirflowCallbacksKeyCreatedOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AirflowCallbacksKeyCreatedOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetEtlContainerId

`func (o *AirflowCallbacksKeyCreatedOut) GetEtlContainerId() string`

GetEtlContainerId returns the EtlContainerId field if non-nil, zero value otherwise.

### GetEtlContainerIdOk

`func (o *AirflowCallbacksKeyCreatedOut) GetEtlContainerIdOk() (*string, bool)`

GetEtlContainerIdOk returns a tuple with the EtlContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlContainerId

`func (o *AirflowCallbacksKeyCreatedOut) SetEtlContainerId(v string)`

SetEtlContainerId sets EtlContainerId field to given value.

### HasEtlContainerId

`func (o *AirflowCallbacksKeyCreatedOut) HasEtlContainerId() bool`

HasEtlContainerId returns a boolean if a field has been set.

### SetEtlContainerIdNil

`func (o *AirflowCallbacksKeyCreatedOut) SetEtlContainerIdNil(b bool)`

 SetEtlContainerIdNil sets the value for EtlContainerId to be an explicit nil

### UnsetEtlContainerId
`func (o *AirflowCallbacksKeyCreatedOut) UnsetEtlContainerId()`

UnsetEtlContainerId ensures that no value is present for EtlContainerId, not even an explicit nil
### GetMcdId

`func (o *AirflowCallbacksKeyCreatedOut) GetMcdId() string`

GetMcdId returns the McdId field if non-nil, zero value otherwise.

### GetMcdIdOk

`func (o *AirflowCallbacksKeyCreatedOut) GetMcdIdOk() (*string, bool)`

GetMcdIdOk returns a tuple with the McdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcdId

`func (o *AirflowCallbacksKeyCreatedOut) SetMcdId(v string)`

SetMcdId sets McdId field to given value.


### GetMcdToken

`func (o *AirflowCallbacksKeyCreatedOut) GetMcdToken() string`

GetMcdToken returns the McdToken field if non-nil, zero value otherwise.

### GetMcdTokenOk

`func (o *AirflowCallbacksKeyCreatedOut) GetMcdTokenOk() (*string, bool)`

GetMcdTokenOk returns a tuple with the McdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcdToken

`func (o *AirflowCallbacksKeyCreatedOut) SetMcdToken(v string)`

SetMcdToken sets McdToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


