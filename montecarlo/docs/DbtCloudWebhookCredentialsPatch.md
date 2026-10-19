# DbtCloudWebhookCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceToken** | Pointer to **NullableString** | A dbt Cloud service token that can read the account. Stored by Monte Carlo and never returned. | [optional] 
**HmacSecret** | Pointer to **NullableString** | The webhook&#39;s HMAC secret, shown by dbt Cloud when the webhook is created. Monte Carlo checks each event&#39;s signature with it. Stored by Monte Carlo and never returned. | [optional] 
**AccountId** | Pointer to **NullableString** | Your dbt Cloud account id, the number in the account&#39;s URLs. | [optional] 
**WebhookId** | Pointer to **NullableString** | Id of the dbt Cloud webhook that sends job.run.completed events to Monte Carlo. | [optional] 
**BaseUrl** | Pointer to **NullableString** | Your dbt Cloud access URL followed by /api/v2, such as https://emea.dbt.com/api/v2. Leave it out for https://cloud.getdbt.com/api/v2. Setting it to a new URL requires sending service_token again. | [optional] 

## Methods

### NewDbtCloudWebhookCredentialsPatch

`func NewDbtCloudWebhookCredentialsPatch() *DbtCloudWebhookCredentialsPatch`

NewDbtCloudWebhookCredentialsPatch instantiates a new DbtCloudWebhookCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbtCloudWebhookCredentialsPatchWithDefaults

`func NewDbtCloudWebhookCredentialsPatchWithDefaults() *DbtCloudWebhookCredentialsPatch`

NewDbtCloudWebhookCredentialsPatchWithDefaults instantiates a new DbtCloudWebhookCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceToken

`func (o *DbtCloudWebhookCredentialsPatch) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *DbtCloudWebhookCredentialsPatch) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *DbtCloudWebhookCredentialsPatch) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *DbtCloudWebhookCredentialsPatch) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### SetServiceTokenNil

`func (o *DbtCloudWebhookCredentialsPatch) SetServiceTokenNil(b bool)`

 SetServiceTokenNil sets the value for ServiceToken to be an explicit nil

### UnsetServiceToken
`func (o *DbtCloudWebhookCredentialsPatch) UnsetServiceToken()`

UnsetServiceToken ensures that no value is present for ServiceToken, not even an explicit nil
### GetHmacSecret

`func (o *DbtCloudWebhookCredentialsPatch) GetHmacSecret() string`

GetHmacSecret returns the HmacSecret field if non-nil, zero value otherwise.

### GetHmacSecretOk

`func (o *DbtCloudWebhookCredentialsPatch) GetHmacSecretOk() (*string, bool)`

GetHmacSecretOk returns a tuple with the HmacSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacSecret

`func (o *DbtCloudWebhookCredentialsPatch) SetHmacSecret(v string)`

SetHmacSecret sets HmacSecret field to given value.

### HasHmacSecret

`func (o *DbtCloudWebhookCredentialsPatch) HasHmacSecret() bool`

HasHmacSecret returns a boolean if a field has been set.

### SetHmacSecretNil

`func (o *DbtCloudWebhookCredentialsPatch) SetHmacSecretNil(b bool)`

 SetHmacSecretNil sets the value for HmacSecret to be an explicit nil

### UnsetHmacSecret
`func (o *DbtCloudWebhookCredentialsPatch) UnsetHmacSecret()`

UnsetHmacSecret ensures that no value is present for HmacSecret, not even an explicit nil
### GetAccountId

`func (o *DbtCloudWebhookCredentialsPatch) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DbtCloudWebhookCredentialsPatch) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DbtCloudWebhookCredentialsPatch) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *DbtCloudWebhookCredentialsPatch) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *DbtCloudWebhookCredentialsPatch) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *DbtCloudWebhookCredentialsPatch) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetWebhookId

`func (o *DbtCloudWebhookCredentialsPatch) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *DbtCloudWebhookCredentialsPatch) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *DbtCloudWebhookCredentialsPatch) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *DbtCloudWebhookCredentialsPatch) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### SetWebhookIdNil

`func (o *DbtCloudWebhookCredentialsPatch) SetWebhookIdNil(b bool)`

 SetWebhookIdNil sets the value for WebhookId to be an explicit nil

### UnsetWebhookId
`func (o *DbtCloudWebhookCredentialsPatch) UnsetWebhookId()`

UnsetWebhookId ensures that no value is present for WebhookId, not even an explicit nil
### GetBaseUrl

`func (o *DbtCloudWebhookCredentialsPatch) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *DbtCloudWebhookCredentialsPatch) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *DbtCloudWebhookCredentialsPatch) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *DbtCloudWebhookCredentialsPatch) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *DbtCloudWebhookCredentialsPatch) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *DbtCloudWebhookCredentialsPatch) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


