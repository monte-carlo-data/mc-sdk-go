# DbtCloudWebhookCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**AccountId** | **string** | Your dbt Cloud account id, the number in the account&#39;s URLs. | 
**WebhookId** | **string** | Id of the dbt Cloud webhook that sends job.run.completed events to Monte Carlo. | 
**BaseUrl** | **NullableString** | Your dbt Cloud access URL followed by /api/v2. Null means https://cloud.getdbt.com/api/v2. | 

## Methods

### NewDbtCloudWebhookCredentialsOut

`func NewDbtCloudWebhookCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, accountId string, webhookId string, baseUrl NullableString, ) *DbtCloudWebhookCredentialsOut`

NewDbtCloudWebhookCredentialsOut instantiates a new DbtCloudWebhookCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbtCloudWebhookCredentialsOutWithDefaults

`func NewDbtCloudWebhookCredentialsOutWithDefaults() *DbtCloudWebhookCredentialsOut`

NewDbtCloudWebhookCredentialsOutWithDefaults instantiates a new DbtCloudWebhookCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbtCloudWebhookCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbtCloudWebhookCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbtCloudWebhookCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *DbtCloudWebhookCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *DbtCloudWebhookCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *DbtCloudWebhookCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *DbtCloudWebhookCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DbtCloudWebhookCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DbtCloudWebhookCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *DbtCloudWebhookCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DbtCloudWebhookCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DbtCloudWebhookCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetAccountId

`func (o *DbtCloudWebhookCredentialsOut) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DbtCloudWebhookCredentialsOut) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DbtCloudWebhookCredentialsOut) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetWebhookId

`func (o *DbtCloudWebhookCredentialsOut) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *DbtCloudWebhookCredentialsOut) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *DbtCloudWebhookCredentialsOut) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetBaseUrl

`func (o *DbtCloudWebhookCredentialsOut) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *DbtCloudWebhookCredentialsOut) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *DbtCloudWebhookCredentialsOut) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### SetBaseUrlNil

`func (o *DbtCloudWebhookCredentialsOut) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *DbtCloudWebhookCredentialsOut) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


