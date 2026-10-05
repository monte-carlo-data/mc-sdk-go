# FivetranCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | Key of the Fivetran API key. Stored by Monte Carlo and never returned. | 
**ApiPassword** | **string** | Secret of the Fivetran API key. Stored by Monte Carlo and never returned. | 
**BaseUrl** | Pointer to **NullableString** | URL of the Fivetran REST API. Leave it out for https://api.fivetran.com/v1/. | [optional] 

## Methods

### NewFivetranCredentialsIn

`func NewFivetranCredentialsIn(apiKey string, apiPassword string, ) *FivetranCredentialsIn`

NewFivetranCredentialsIn instantiates a new FivetranCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFivetranCredentialsInWithDefaults

`func NewFivetranCredentialsInWithDefaults() *FivetranCredentialsIn`

NewFivetranCredentialsInWithDefaults instantiates a new FivetranCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *FivetranCredentialsIn) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *FivetranCredentialsIn) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *FivetranCredentialsIn) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetApiPassword

`func (o *FivetranCredentialsIn) GetApiPassword() string`

GetApiPassword returns the ApiPassword field if non-nil, zero value otherwise.

### GetApiPasswordOk

`func (o *FivetranCredentialsIn) GetApiPasswordOk() (*string, bool)`

GetApiPasswordOk returns a tuple with the ApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPassword

`func (o *FivetranCredentialsIn) SetApiPassword(v string)`

SetApiPassword sets ApiPassword field to given value.


### GetBaseUrl

`func (o *FivetranCredentialsIn) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *FivetranCredentialsIn) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *FivetranCredentialsIn) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *FivetranCredentialsIn) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *FivetranCredentialsIn) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *FivetranCredentialsIn) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


