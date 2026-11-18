# OpenTelemetryKeyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the key. | 
**Description** | **string** | What this key is for. | 
**CreatedTime** | **time.Time** | When the key was created. | 
**DeploymentId** | **string** | The &#x60;AGENT_OBSERVABILITY&#x60; deployment the key sends traces to. | 
**McdId** | **string** | Key id to present, as the &#x60;x-mcd-id&#x60; header or a client&#39;s &#x60;mcd_id&#x60; setting. The same value as &#x60;id&#x60;. | 

## Methods

### NewOpenTelemetryKeyOut

`func NewOpenTelemetryKeyOut(id string, description string, createdTime time.Time, deploymentId string, mcdId string, ) *OpenTelemetryKeyOut`

NewOpenTelemetryKeyOut instantiates a new OpenTelemetryKeyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenTelemetryKeyOutWithDefaults

`func NewOpenTelemetryKeyOutWithDefaults() *OpenTelemetryKeyOut`

NewOpenTelemetryKeyOutWithDefaults instantiates a new OpenTelemetryKeyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpenTelemetryKeyOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenTelemetryKeyOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenTelemetryKeyOut) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *OpenTelemetryKeyOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenTelemetryKeyOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenTelemetryKeyOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedTime

`func (o *OpenTelemetryKeyOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *OpenTelemetryKeyOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *OpenTelemetryKeyOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetDeploymentId

`func (o *OpenTelemetryKeyOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *OpenTelemetryKeyOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *OpenTelemetryKeyOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetMcdId

`func (o *OpenTelemetryKeyOut) GetMcdId() string`

GetMcdId returns the McdId field if non-nil, zero value otherwise.

### GetMcdIdOk

`func (o *OpenTelemetryKeyOut) GetMcdIdOk() (*string, bool)`

GetMcdIdOk returns a tuple with the McdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcdId

`func (o *OpenTelemetryKeyOut) SetMcdId(v string)`

SetMcdId sets McdId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


