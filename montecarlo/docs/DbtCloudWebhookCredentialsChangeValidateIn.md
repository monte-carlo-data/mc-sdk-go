# DbtCloudWebhookCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**ServiceToken** | Pointer to **NullableString** | A dbt Cloud service token that can read the account. Stored by Monte Carlo and never returned. | [optional] 
**HmacSecret** | Pointer to **NullableString** | The webhook&#39;s HMAC secret, shown by dbt Cloud when the webhook is created. Monte Carlo checks each event&#39;s signature with it. Stored by Monte Carlo and never returned. | [optional] 
**AccountId** | Pointer to **NullableString** | Your dbt Cloud account id, the number in the account&#39;s URLs. | [optional] 
**WebhookId** | Pointer to **NullableString** | Id of the dbt Cloud webhook that sends job.run.completed events to Monte Carlo. | [optional] 
**BaseUrl** | Pointer to **NullableString** | Your dbt Cloud access URL followed by /api/v2, such as https://emea.dbt.com/api/v2. Leave it out for https://cloud.getdbt.com/api/v2. Setting it to a new URL requires sending service_token again. | [optional] 

## Methods

### NewDbtCloudWebhookCredentialsChangeValidateIn

`func NewDbtCloudWebhookCredentialsChangeValidateIn(deploymentId string, ) *DbtCloudWebhookCredentialsChangeValidateIn`

NewDbtCloudWebhookCredentialsChangeValidateIn instantiates a new DbtCloudWebhookCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbtCloudWebhookCredentialsChangeValidateInWithDefaults

`func NewDbtCloudWebhookCredentialsChangeValidateInWithDefaults() *DbtCloudWebhookCredentialsChangeValidateIn`

NewDbtCloudWebhookCredentialsChangeValidateInWithDefaults instantiates a new DbtCloudWebhookCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetServiceToken

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### SetServiceTokenNil

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) SetServiceTokenNil(b bool)`

 SetServiceTokenNil sets the value for ServiceToken to be an explicit nil

### UnsetServiceToken
`func (o *DbtCloudWebhookCredentialsChangeValidateIn) UnsetServiceToken()`

UnsetServiceToken ensures that no value is present for ServiceToken, not even an explicit nil
### GetHmacSecret

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) GetHmacSecret() string`

GetHmacSecret returns the HmacSecret field if non-nil, zero value otherwise.

### GetHmacSecretOk

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) GetHmacSecretOk() (*string, bool)`

GetHmacSecretOk returns a tuple with the HmacSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacSecret

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) SetHmacSecret(v string)`

SetHmacSecret sets HmacSecret field to given value.

### HasHmacSecret

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) HasHmacSecret() bool`

HasHmacSecret returns a boolean if a field has been set.

### SetHmacSecretNil

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) SetHmacSecretNil(b bool)`

 SetHmacSecretNil sets the value for HmacSecret to be an explicit nil

### UnsetHmacSecret
`func (o *DbtCloudWebhookCredentialsChangeValidateIn) UnsetHmacSecret()`

UnsetHmacSecret ensures that no value is present for HmacSecret, not even an explicit nil
### GetAccountId

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *DbtCloudWebhookCredentialsChangeValidateIn) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetWebhookId

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### SetWebhookIdNil

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) SetWebhookIdNil(b bool)`

 SetWebhookIdNil sets the value for WebhookId to be an explicit nil

### UnsetWebhookId
`func (o *DbtCloudWebhookCredentialsChangeValidateIn) UnsetWebhookId()`

UnsetWebhookId ensures that no value is present for WebhookId, not even an explicit nil
### GetBaseUrl

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *DbtCloudWebhookCredentialsChangeValidateIn) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *DbtCloudWebhookCredentialsChangeValidateIn) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


