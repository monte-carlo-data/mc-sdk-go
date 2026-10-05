# FivetranCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**ApiKey** | **string** | Key of the Fivetran API key. Used for this check and not kept. | 
**ApiPassword** | **string** | Secret of the Fivetran API key. Used for this check and not kept. | 
**BaseUrl** | Pointer to **NullableString** | URL of the Fivetran REST API. Leave it out for https://api.fivetran.com/v1/. | [optional] 

## Methods

### NewFivetranCredentialsValidateIn

`func NewFivetranCredentialsValidateIn(deploymentId string, apiKey string, apiPassword string, ) *FivetranCredentialsValidateIn`

NewFivetranCredentialsValidateIn instantiates a new FivetranCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFivetranCredentialsValidateInWithDefaults

`func NewFivetranCredentialsValidateInWithDefaults() *FivetranCredentialsValidateIn`

NewFivetranCredentialsValidateInWithDefaults instantiates a new FivetranCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *FivetranCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *FivetranCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *FivetranCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetApiKey

`func (o *FivetranCredentialsValidateIn) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *FivetranCredentialsValidateIn) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *FivetranCredentialsValidateIn) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetApiPassword

`func (o *FivetranCredentialsValidateIn) GetApiPassword() string`

GetApiPassword returns the ApiPassword field if non-nil, zero value otherwise.

### GetApiPasswordOk

`func (o *FivetranCredentialsValidateIn) GetApiPasswordOk() (*string, bool)`

GetApiPasswordOk returns a tuple with the ApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPassword

`func (o *FivetranCredentialsValidateIn) SetApiPassword(v string)`

SetApiPassword sets ApiPassword field to given value.


### GetBaseUrl

`func (o *FivetranCredentialsValidateIn) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *FivetranCredentialsValidateIn) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *FivetranCredentialsValidateIn) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *FivetranCredentialsValidateIn) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *FivetranCredentialsValidateIn) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *FivetranCredentialsValidateIn) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


