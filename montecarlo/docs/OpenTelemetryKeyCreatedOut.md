# OpenTelemetryKeyCreatedOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the key. | 
**Description** | **string** | What this key is for. | 
**CreatedTime** | **time.Time** | When the key was created. | 
**DeploymentId** | **string** | The &#x60;AGENT_OBSERVABILITY&#x60; deployment the key sends traces to. | 
**McdId** | **string** | Key id to present, as the &#x60;x-mcd-id&#x60; header or a client&#39;s &#x60;mcd_id&#x60; setting. The same value as &#x60;id&#x60;. | 
**McdToken** | **string** | Secret to present, as the &#x60;x-mcd-token&#x60; header or a client&#39;s &#x60;mcd_token&#x60; setting. Returned once, by this call, and not retrievable afterwards. Terraform keeps it in state, so anyone who can read the state can read the secret: &#x60;sensitive&#x60; only hides it from CLI output. Keep state in a backend that encrypts it and limits who can read it. Hand the secret on through a write-only argument, such as &#x60;aws_secretsmanager_secret_version.secret_string_wo&#x60;, not through an output, because a &#x60;sensitive&#x60; output is stored in state too. To rotate, create a new key and delete this one. | [readonly] 

## Methods

### NewOpenTelemetryKeyCreatedOut

`func NewOpenTelemetryKeyCreatedOut(id string, description string, createdTime time.Time, deploymentId string, mcdId string, mcdToken string, ) *OpenTelemetryKeyCreatedOut`

NewOpenTelemetryKeyCreatedOut instantiates a new OpenTelemetryKeyCreatedOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenTelemetryKeyCreatedOutWithDefaults

`func NewOpenTelemetryKeyCreatedOutWithDefaults() *OpenTelemetryKeyCreatedOut`

NewOpenTelemetryKeyCreatedOutWithDefaults instantiates a new OpenTelemetryKeyCreatedOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpenTelemetryKeyCreatedOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenTelemetryKeyCreatedOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenTelemetryKeyCreatedOut) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *OpenTelemetryKeyCreatedOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenTelemetryKeyCreatedOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenTelemetryKeyCreatedOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedTime

`func (o *OpenTelemetryKeyCreatedOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *OpenTelemetryKeyCreatedOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *OpenTelemetryKeyCreatedOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetDeploymentId

`func (o *OpenTelemetryKeyCreatedOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *OpenTelemetryKeyCreatedOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *OpenTelemetryKeyCreatedOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetMcdId

`func (o *OpenTelemetryKeyCreatedOut) GetMcdId() string`

GetMcdId returns the McdId field if non-nil, zero value otherwise.

### GetMcdIdOk

`func (o *OpenTelemetryKeyCreatedOut) GetMcdIdOk() (*string, bool)`

GetMcdIdOk returns a tuple with the McdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcdId

`func (o *OpenTelemetryKeyCreatedOut) SetMcdId(v string)`

SetMcdId sets McdId field to given value.


### GetMcdToken

`func (o *OpenTelemetryKeyCreatedOut) GetMcdToken() string`

GetMcdToken returns the McdToken field if non-nil, zero value otherwise.

### GetMcdTokenOk

`func (o *OpenTelemetryKeyCreatedOut) GetMcdTokenOk() (*string, bool)`

GetMcdTokenOk returns a tuple with the McdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcdToken

`func (o *OpenTelemetryKeyCreatedOut) SetMcdToken(v string)`

SetMcdToken sets McdToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


