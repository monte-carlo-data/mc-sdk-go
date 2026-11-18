# OpenTelemetryKeyIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | What this key is for. Monte Carlo generates one naming the deployment if you leave it out. | [optional] 
**DeploymentId** | **string** | The &#x60;AGENT_OBSERVABILITY&#x60; deployment to send traces to. It may still be &#x60;PROVISIONING&#x60;. | 

## Methods

### NewOpenTelemetryKeyIn

`func NewOpenTelemetryKeyIn(deploymentId string, ) *OpenTelemetryKeyIn`

NewOpenTelemetryKeyIn instantiates a new OpenTelemetryKeyIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenTelemetryKeyInWithDefaults

`func NewOpenTelemetryKeyInWithDefaults() *OpenTelemetryKeyIn`

NewOpenTelemetryKeyInWithDefaults instantiates a new OpenTelemetryKeyIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OpenTelemetryKeyIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenTelemetryKeyIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenTelemetryKeyIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenTelemetryKeyIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OpenTelemetryKeyIn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OpenTelemetryKeyIn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDeploymentId

`func (o *OpenTelemetryKeyIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *OpenTelemetryKeyIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *OpenTelemetryKeyIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


