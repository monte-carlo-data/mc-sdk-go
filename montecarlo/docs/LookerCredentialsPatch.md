# LookerCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **NullableString** | URL of the Looker API, such as https://acme.cloud.looker.com. | [optional] 
**ApiClientId** | Pointer to **NullableString** | Client ID of the Looker API key. | [optional] 
**ApiClientSecret** | Pointer to **NullableString** | Client secret of the Looker API key. Stored by Monte Carlo and never returned. | [optional] 
**VerifySsl** | Pointer to **NullableBool** | Whether to verify Looker&#39;s TLS certificate. Verified when left out. | [optional] 

## Methods

### NewLookerCredentialsPatch

`func NewLookerCredentialsPatch() *LookerCredentialsPatch`

NewLookerCredentialsPatch instantiates a new LookerCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookerCredentialsPatchWithDefaults

`func NewLookerCredentialsPatchWithDefaults() *LookerCredentialsPatch`

NewLookerCredentialsPatchWithDefaults instantiates a new LookerCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *LookerCredentialsPatch) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *LookerCredentialsPatch) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *LookerCredentialsPatch) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *LookerCredentialsPatch) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *LookerCredentialsPatch) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *LookerCredentialsPatch) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
### GetApiClientId

`func (o *LookerCredentialsPatch) GetApiClientId() string`

GetApiClientId returns the ApiClientId field if non-nil, zero value otherwise.

### GetApiClientIdOk

`func (o *LookerCredentialsPatch) GetApiClientIdOk() (*string, bool)`

GetApiClientIdOk returns a tuple with the ApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientId

`func (o *LookerCredentialsPatch) SetApiClientId(v string)`

SetApiClientId sets ApiClientId field to given value.

### HasApiClientId

`func (o *LookerCredentialsPatch) HasApiClientId() bool`

HasApiClientId returns a boolean if a field has been set.

### SetApiClientIdNil

`func (o *LookerCredentialsPatch) SetApiClientIdNil(b bool)`

 SetApiClientIdNil sets the value for ApiClientId to be an explicit nil

### UnsetApiClientId
`func (o *LookerCredentialsPatch) UnsetApiClientId()`

UnsetApiClientId ensures that no value is present for ApiClientId, not even an explicit nil
### GetApiClientSecret

`func (o *LookerCredentialsPatch) GetApiClientSecret() string`

GetApiClientSecret returns the ApiClientSecret field if non-nil, zero value otherwise.

### GetApiClientSecretOk

`func (o *LookerCredentialsPatch) GetApiClientSecretOk() (*string, bool)`

GetApiClientSecretOk returns a tuple with the ApiClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientSecret

`func (o *LookerCredentialsPatch) SetApiClientSecret(v string)`

SetApiClientSecret sets ApiClientSecret field to given value.

### HasApiClientSecret

`func (o *LookerCredentialsPatch) HasApiClientSecret() bool`

HasApiClientSecret returns a boolean if a field has been set.

### SetApiClientSecretNil

`func (o *LookerCredentialsPatch) SetApiClientSecretNil(b bool)`

 SetApiClientSecretNil sets the value for ApiClientSecret to be an explicit nil

### UnsetApiClientSecret
`func (o *LookerCredentialsPatch) UnsetApiClientSecret()`

UnsetApiClientSecret ensures that no value is present for ApiClientSecret, not even an explicit nil
### GetVerifySsl

`func (o *LookerCredentialsPatch) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *LookerCredentialsPatch) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *LookerCredentialsPatch) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *LookerCredentialsPatch) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.

### SetVerifySslNil

`func (o *LookerCredentialsPatch) SetVerifySslNil(b bool)`

 SetVerifySslNil sets the value for VerifySsl to be an explicit nil

### UnsetVerifySsl
`func (o *LookerCredentialsPatch) UnsetVerifySsl()`

UnsetVerifySsl ensures that no value is present for VerifySsl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


