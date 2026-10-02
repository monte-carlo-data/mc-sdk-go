# LookerCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**BaseUrl** | **string** | URL of the Looker API, such as https://acme.cloud.looker.com. | 
**ApiClientId** | **string** | Client ID of the Looker API key. | 
**ApiClientSecret** | **string** | Client secret of the Looker API key. Used for this check and not kept. | 
**VerifySsl** | Pointer to **NullableBool** | Whether to verify Looker&#39;s TLS certificate. Verified when left out. | [optional] 

## Methods

### NewLookerCredentialsValidateIn

`func NewLookerCredentialsValidateIn(deploymentId string, baseUrl string, apiClientId string, apiClientSecret string, ) *LookerCredentialsValidateIn`

NewLookerCredentialsValidateIn instantiates a new LookerCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookerCredentialsValidateInWithDefaults

`func NewLookerCredentialsValidateInWithDefaults() *LookerCredentialsValidateIn`

NewLookerCredentialsValidateInWithDefaults instantiates a new LookerCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *LookerCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *LookerCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *LookerCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetBaseUrl

`func (o *LookerCredentialsValidateIn) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *LookerCredentialsValidateIn) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *LookerCredentialsValidateIn) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetApiClientId

`func (o *LookerCredentialsValidateIn) GetApiClientId() string`

GetApiClientId returns the ApiClientId field if non-nil, zero value otherwise.

### GetApiClientIdOk

`func (o *LookerCredentialsValidateIn) GetApiClientIdOk() (*string, bool)`

GetApiClientIdOk returns a tuple with the ApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientId

`func (o *LookerCredentialsValidateIn) SetApiClientId(v string)`

SetApiClientId sets ApiClientId field to given value.


### GetApiClientSecret

`func (o *LookerCredentialsValidateIn) GetApiClientSecret() string`

GetApiClientSecret returns the ApiClientSecret field if non-nil, zero value otherwise.

### GetApiClientSecretOk

`func (o *LookerCredentialsValidateIn) GetApiClientSecretOk() (*string, bool)`

GetApiClientSecretOk returns a tuple with the ApiClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientSecret

`func (o *LookerCredentialsValidateIn) SetApiClientSecret(v string)`

SetApiClientSecret sets ApiClientSecret field to given value.


### GetVerifySsl

`func (o *LookerCredentialsValidateIn) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *LookerCredentialsValidateIn) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *LookerCredentialsValidateIn) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *LookerCredentialsValidateIn) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.

### SetVerifySslNil

`func (o *LookerCredentialsValidateIn) SetVerifySslNil(b bool)`

 SetVerifySslNil sets the value for VerifySsl to be an explicit nil

### UnsetVerifySsl
`func (o *LookerCredentialsValidateIn) UnsetVerifySsl()`

UnsetVerifySsl ensures that no value is present for VerifySsl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


