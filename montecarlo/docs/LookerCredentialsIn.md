# LookerCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | **string** | URL of the Looker API, such as https://acme.cloud.looker.com. | 
**ApiClientId** | **string** | Client ID of the Looker API key. | 
**ApiClientSecret** | **string** | Client secret of the Looker API key. Stored by Monte Carlo and never returned. | 
**VerifySsl** | Pointer to **NullableBool** | Whether to verify Looker&#39;s TLS certificate. Verified when left out. | [optional] 

## Methods

### NewLookerCredentialsIn

`func NewLookerCredentialsIn(baseUrl string, apiClientId string, apiClientSecret string, ) *LookerCredentialsIn`

NewLookerCredentialsIn instantiates a new LookerCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookerCredentialsInWithDefaults

`func NewLookerCredentialsInWithDefaults() *LookerCredentialsIn`

NewLookerCredentialsInWithDefaults instantiates a new LookerCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *LookerCredentialsIn) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *LookerCredentialsIn) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *LookerCredentialsIn) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetApiClientId

`func (o *LookerCredentialsIn) GetApiClientId() string`

GetApiClientId returns the ApiClientId field if non-nil, zero value otherwise.

### GetApiClientIdOk

`func (o *LookerCredentialsIn) GetApiClientIdOk() (*string, bool)`

GetApiClientIdOk returns a tuple with the ApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientId

`func (o *LookerCredentialsIn) SetApiClientId(v string)`

SetApiClientId sets ApiClientId field to given value.


### GetApiClientSecret

`func (o *LookerCredentialsIn) GetApiClientSecret() string`

GetApiClientSecret returns the ApiClientSecret field if non-nil, zero value otherwise.

### GetApiClientSecretOk

`func (o *LookerCredentialsIn) GetApiClientSecretOk() (*string, bool)`

GetApiClientSecretOk returns a tuple with the ApiClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientSecret

`func (o *LookerCredentialsIn) SetApiClientSecret(v string)`

SetApiClientSecret sets ApiClientSecret field to given value.


### GetVerifySsl

`func (o *LookerCredentialsIn) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *LookerCredentialsIn) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *LookerCredentialsIn) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *LookerCredentialsIn) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.

### SetVerifySslNil

`func (o *LookerCredentialsIn) SetVerifySslNil(b bool)`

 SetVerifySslNil sets the value for VerifySsl to be an explicit nil

### UnsetVerifySsl
`func (o *LookerCredentialsIn) UnsetVerifySsl()`

UnsetVerifySsl ensures that no value is present for VerifySsl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


