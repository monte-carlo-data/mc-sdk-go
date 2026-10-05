# FivetranCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **NullableString** | Key of the Fivetran API key. Stored by Monte Carlo and never returned. | [optional] 
**ApiPassword** | Pointer to **NullableString** | Secret of the Fivetran API key. Stored by Monte Carlo and never returned. | [optional] 
**BaseUrl** | Pointer to **NullableString** | URL of the Fivetran REST API. Leave it out for https://api.fivetran.com/v1/. | [optional] 

## Methods

### NewFivetranCredentialsPatch

`func NewFivetranCredentialsPatch() *FivetranCredentialsPatch`

NewFivetranCredentialsPatch instantiates a new FivetranCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFivetranCredentialsPatchWithDefaults

`func NewFivetranCredentialsPatchWithDefaults() *FivetranCredentialsPatch`

NewFivetranCredentialsPatchWithDefaults instantiates a new FivetranCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *FivetranCredentialsPatch) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *FivetranCredentialsPatch) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *FivetranCredentialsPatch) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *FivetranCredentialsPatch) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *FivetranCredentialsPatch) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *FivetranCredentialsPatch) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetApiPassword

`func (o *FivetranCredentialsPatch) GetApiPassword() string`

GetApiPassword returns the ApiPassword field if non-nil, zero value otherwise.

### GetApiPasswordOk

`func (o *FivetranCredentialsPatch) GetApiPasswordOk() (*string, bool)`

GetApiPasswordOk returns a tuple with the ApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPassword

`func (o *FivetranCredentialsPatch) SetApiPassword(v string)`

SetApiPassword sets ApiPassword field to given value.

### HasApiPassword

`func (o *FivetranCredentialsPatch) HasApiPassword() bool`

HasApiPassword returns a boolean if a field has been set.

### SetApiPasswordNil

`func (o *FivetranCredentialsPatch) SetApiPasswordNil(b bool)`

 SetApiPasswordNil sets the value for ApiPassword to be an explicit nil

### UnsetApiPassword
`func (o *FivetranCredentialsPatch) UnsetApiPassword()`

UnsetApiPassword ensures that no value is present for ApiPassword, not even an explicit nil
### GetBaseUrl

`func (o *FivetranCredentialsPatch) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *FivetranCredentialsPatch) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *FivetranCredentialsPatch) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *FivetranCredentialsPatch) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *FivetranCredentialsPatch) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *FivetranCredentialsPatch) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


