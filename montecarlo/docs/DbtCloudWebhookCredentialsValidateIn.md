# DbtCloudWebhookCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**ServiceToken** | **string** | A dbt Cloud service token that can read the account. Used for this check and not kept. | 
**HmacSecret** | **string** | The webhook&#39;s HMAC secret, shown by dbt Cloud when the webhook is created. Monte Carlo checks each event&#39;s signature with it. Used for this check and not kept. | 
**AccountId** | **string** | Your dbt Cloud account id, the number in the account&#39;s URLs. | 
**WebhookId** | **string** | Id of the dbt Cloud webhook that sends job.run.completed events to Monte Carlo. | 
**BaseUrl** | Pointer to **NullableString** | Your dbt Cloud access URL followed by /api/v2, such as https://emea.dbt.com/api/v2. Leave it out for https://cloud.getdbt.com/api/v2. | [optional] 

## Methods

### NewDbtCloudWebhookCredentialsValidateIn

`func NewDbtCloudWebhookCredentialsValidateIn(deploymentId string, serviceToken string, hmacSecret string, accountId string, webhookId string, ) *DbtCloudWebhookCredentialsValidateIn`

NewDbtCloudWebhookCredentialsValidateIn instantiates a new DbtCloudWebhookCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbtCloudWebhookCredentialsValidateInWithDefaults

`func NewDbtCloudWebhookCredentialsValidateInWithDefaults() *DbtCloudWebhookCredentialsValidateIn`

NewDbtCloudWebhookCredentialsValidateInWithDefaults instantiates a new DbtCloudWebhookCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *DbtCloudWebhookCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DbtCloudWebhookCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DbtCloudWebhookCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetServiceToken

`func (o *DbtCloudWebhookCredentialsValidateIn) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *DbtCloudWebhookCredentialsValidateIn) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *DbtCloudWebhookCredentialsValidateIn) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.


### GetHmacSecret

`func (o *DbtCloudWebhookCredentialsValidateIn) GetHmacSecret() string`

GetHmacSecret returns the HmacSecret field if non-nil, zero value otherwise.

### GetHmacSecretOk

`func (o *DbtCloudWebhookCredentialsValidateIn) GetHmacSecretOk() (*string, bool)`

GetHmacSecretOk returns a tuple with the HmacSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacSecret

`func (o *DbtCloudWebhookCredentialsValidateIn) SetHmacSecret(v string)`

SetHmacSecret sets HmacSecret field to given value.


### GetAccountId

`func (o *DbtCloudWebhookCredentialsValidateIn) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DbtCloudWebhookCredentialsValidateIn) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DbtCloudWebhookCredentialsValidateIn) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetWebhookId

`func (o *DbtCloudWebhookCredentialsValidateIn) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *DbtCloudWebhookCredentialsValidateIn) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *DbtCloudWebhookCredentialsValidateIn) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetBaseUrl

`func (o *DbtCloudWebhookCredentialsValidateIn) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *DbtCloudWebhookCredentialsValidateIn) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *DbtCloudWebhookCredentialsValidateIn) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *DbtCloudWebhookCredentialsValidateIn) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *DbtCloudWebhookCredentialsValidateIn) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *DbtCloudWebhookCredentialsValidateIn) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


