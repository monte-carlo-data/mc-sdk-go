# DbtCloudWebhookCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceToken** | **string** | A dbt Cloud service token that can read the account. Stored by Monte Carlo and never returned. | 
**HmacSecret** | **string** | The webhook&#39;s HMAC secret, shown by dbt Cloud when the webhook is created. Monte Carlo checks each event&#39;s signature with it. Stored by Monte Carlo and never returned. | 
**AccountId** | **string** | Your dbt Cloud account id, the number in the account&#39;s URLs. | 
**WebhookId** | **string** | Id of the dbt Cloud webhook that sends job.run.completed events to Monte Carlo. | 
**BaseUrl** | Pointer to **NullableString** | Your dbt Cloud access URL followed by /api/v2, such as https://emea.dbt.com/api/v2. Leave it out for https://cloud.getdbt.com/api/v2. | [optional] 

## Methods

### NewDbtCloudWebhookCredentialsIn

`func NewDbtCloudWebhookCredentialsIn(serviceToken string, hmacSecret string, accountId string, webhookId string, ) *DbtCloudWebhookCredentialsIn`

NewDbtCloudWebhookCredentialsIn instantiates a new DbtCloudWebhookCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbtCloudWebhookCredentialsInWithDefaults

`func NewDbtCloudWebhookCredentialsInWithDefaults() *DbtCloudWebhookCredentialsIn`

NewDbtCloudWebhookCredentialsInWithDefaults instantiates a new DbtCloudWebhookCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceToken

`func (o *DbtCloudWebhookCredentialsIn) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *DbtCloudWebhookCredentialsIn) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *DbtCloudWebhookCredentialsIn) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.


### GetHmacSecret

`func (o *DbtCloudWebhookCredentialsIn) GetHmacSecret() string`

GetHmacSecret returns the HmacSecret field if non-nil, zero value otherwise.

### GetHmacSecretOk

`func (o *DbtCloudWebhookCredentialsIn) GetHmacSecretOk() (*string, bool)`

GetHmacSecretOk returns a tuple with the HmacSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacSecret

`func (o *DbtCloudWebhookCredentialsIn) SetHmacSecret(v string)`

SetHmacSecret sets HmacSecret field to given value.


### GetAccountId

`func (o *DbtCloudWebhookCredentialsIn) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DbtCloudWebhookCredentialsIn) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DbtCloudWebhookCredentialsIn) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetWebhookId

`func (o *DbtCloudWebhookCredentialsIn) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *DbtCloudWebhookCredentialsIn) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *DbtCloudWebhookCredentialsIn) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetBaseUrl

`func (o *DbtCloudWebhookCredentialsIn) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *DbtCloudWebhookCredentialsIn) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *DbtCloudWebhookCredentialsIn) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *DbtCloudWebhookCredentialsIn) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *DbtCloudWebhookCredentialsIn) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *DbtCloudWebhookCredentialsIn) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


