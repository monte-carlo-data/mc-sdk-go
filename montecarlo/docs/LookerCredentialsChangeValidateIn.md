# LookerCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**BaseUrl** | Pointer to **NullableString** | URL of the Looker API, such as https://acme.cloud.looker.com. Changing it requires sending the credentials&#39; secrets again. | [optional] 
**ApiClientId** | Pointer to **NullableString** | Client ID of the Looker API key. | [optional] 
**ApiClientSecret** | Pointer to **NullableString** | Client secret of the Looker API key. Stored by Monte Carlo and never returned. | [optional] 
**VerifySsl** | Pointer to **NullableBool** | Whether to verify Looker&#39;s TLS certificate. Verified when left out. | [optional] 

## Methods

### NewLookerCredentialsChangeValidateIn

`func NewLookerCredentialsChangeValidateIn(deploymentId string, ) *LookerCredentialsChangeValidateIn`

NewLookerCredentialsChangeValidateIn instantiates a new LookerCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookerCredentialsChangeValidateInWithDefaults

`func NewLookerCredentialsChangeValidateInWithDefaults() *LookerCredentialsChangeValidateIn`

NewLookerCredentialsChangeValidateInWithDefaults instantiates a new LookerCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *LookerCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *LookerCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *LookerCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetBaseUrl

`func (o *LookerCredentialsChangeValidateIn) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *LookerCredentialsChangeValidateIn) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *LookerCredentialsChangeValidateIn) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *LookerCredentialsChangeValidateIn) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *LookerCredentialsChangeValidateIn) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *LookerCredentialsChangeValidateIn) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
### GetApiClientId

`func (o *LookerCredentialsChangeValidateIn) GetApiClientId() string`

GetApiClientId returns the ApiClientId field if non-nil, zero value otherwise.

### GetApiClientIdOk

`func (o *LookerCredentialsChangeValidateIn) GetApiClientIdOk() (*string, bool)`

GetApiClientIdOk returns a tuple with the ApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientId

`func (o *LookerCredentialsChangeValidateIn) SetApiClientId(v string)`

SetApiClientId sets ApiClientId field to given value.

### HasApiClientId

`func (o *LookerCredentialsChangeValidateIn) HasApiClientId() bool`

HasApiClientId returns a boolean if a field has been set.

### SetApiClientIdNil

`func (o *LookerCredentialsChangeValidateIn) SetApiClientIdNil(b bool)`

 SetApiClientIdNil sets the value for ApiClientId to be an explicit nil

### UnsetApiClientId
`func (o *LookerCredentialsChangeValidateIn) UnsetApiClientId()`

UnsetApiClientId ensures that no value is present for ApiClientId, not even an explicit nil
### GetApiClientSecret

`func (o *LookerCredentialsChangeValidateIn) GetApiClientSecret() string`

GetApiClientSecret returns the ApiClientSecret field if non-nil, zero value otherwise.

### GetApiClientSecretOk

`func (o *LookerCredentialsChangeValidateIn) GetApiClientSecretOk() (*string, bool)`

GetApiClientSecretOk returns a tuple with the ApiClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientSecret

`func (o *LookerCredentialsChangeValidateIn) SetApiClientSecret(v string)`

SetApiClientSecret sets ApiClientSecret field to given value.

### HasApiClientSecret

`func (o *LookerCredentialsChangeValidateIn) HasApiClientSecret() bool`

HasApiClientSecret returns a boolean if a field has been set.

### SetApiClientSecretNil

`func (o *LookerCredentialsChangeValidateIn) SetApiClientSecretNil(b bool)`

 SetApiClientSecretNil sets the value for ApiClientSecret to be an explicit nil

### UnsetApiClientSecret
`func (o *LookerCredentialsChangeValidateIn) UnsetApiClientSecret()`

UnsetApiClientSecret ensures that no value is present for ApiClientSecret, not even an explicit nil
### GetVerifySsl

`func (o *LookerCredentialsChangeValidateIn) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *LookerCredentialsChangeValidateIn) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *LookerCredentialsChangeValidateIn) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *LookerCredentialsChangeValidateIn) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.

### SetVerifySslNil

`func (o *LookerCredentialsChangeValidateIn) SetVerifySslNil(b bool)`

 SetVerifySslNil sets the value for VerifySsl to be an explicit nil

### UnsetVerifySsl
`func (o *LookerCredentialsChangeValidateIn) UnsetVerifySsl()`

UnsetVerifySsl ensures that no value is present for VerifySsl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


