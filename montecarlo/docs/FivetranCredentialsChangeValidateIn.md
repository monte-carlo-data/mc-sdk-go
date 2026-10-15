# FivetranCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**ApiKey** | Pointer to **NullableString** | Key of the Fivetran API key. Stored by Monte Carlo and never returned. | [optional] 
**ApiPassword** | Pointer to **NullableString** | Secret of the Fivetran API key. Stored by Monte Carlo and never returned. | [optional] 
**BaseUrl** | Pointer to **NullableString** | URL of the Fivetran REST API. Leave it out for https://api.fivetran.com/v1/. Changing it requires sending the credentials&#39; secrets again. | [optional] 

## Methods

### NewFivetranCredentialsChangeValidateIn

`func NewFivetranCredentialsChangeValidateIn(deploymentId string, ) *FivetranCredentialsChangeValidateIn`

NewFivetranCredentialsChangeValidateIn instantiates a new FivetranCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFivetranCredentialsChangeValidateInWithDefaults

`func NewFivetranCredentialsChangeValidateInWithDefaults() *FivetranCredentialsChangeValidateIn`

NewFivetranCredentialsChangeValidateInWithDefaults instantiates a new FivetranCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *FivetranCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *FivetranCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *FivetranCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetApiKey

`func (o *FivetranCredentialsChangeValidateIn) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *FivetranCredentialsChangeValidateIn) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *FivetranCredentialsChangeValidateIn) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *FivetranCredentialsChangeValidateIn) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *FivetranCredentialsChangeValidateIn) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *FivetranCredentialsChangeValidateIn) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetApiPassword

`func (o *FivetranCredentialsChangeValidateIn) GetApiPassword() string`

GetApiPassword returns the ApiPassword field if non-nil, zero value otherwise.

### GetApiPasswordOk

`func (o *FivetranCredentialsChangeValidateIn) GetApiPasswordOk() (*string, bool)`

GetApiPasswordOk returns a tuple with the ApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPassword

`func (o *FivetranCredentialsChangeValidateIn) SetApiPassword(v string)`

SetApiPassword sets ApiPassword field to given value.

### HasApiPassword

`func (o *FivetranCredentialsChangeValidateIn) HasApiPassword() bool`

HasApiPassword returns a boolean if a field has been set.

### SetApiPasswordNil

`func (o *FivetranCredentialsChangeValidateIn) SetApiPasswordNil(b bool)`

 SetApiPasswordNil sets the value for ApiPassword to be an explicit nil

### UnsetApiPassword
`func (o *FivetranCredentialsChangeValidateIn) UnsetApiPassword()`

UnsetApiPassword ensures that no value is present for ApiPassword, not even an explicit nil
### GetBaseUrl

`func (o *FivetranCredentialsChangeValidateIn) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *FivetranCredentialsChangeValidateIn) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *FivetranCredentialsChangeValidateIn) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *FivetranCredentialsChangeValidateIn) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *FivetranCredentialsChangeValidateIn) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *FivetranCredentialsChangeValidateIn) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


