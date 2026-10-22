# SalesforceCrmCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**AuthMode** | [**SalesforceCrmAuthMode**](SalesforceCrmAuthMode.md) | How Monte Carlo signs in. &#x60;password&#x60; takes &#x60;username&#x60;, &#x60;password&#x60; and &#x60;security_token&#x60;. &#x60;oauth_client_credentials&#x60; takes &#x60;consumer_key&#x60;, &#x60;consumer_secret&#x60; and &#x60;domain&#x60;. | 
**Username** | **NullableString** | Salesforce user Monte Carlo signs in as, for &#x60;password&#x60;. | 
**ConsumerKey** | **NullableString** | Consumer key of the connected app Monte Carlo signs in as, for &#x60;oauth_client_credentials&#x60;. | 
**Domain** | **NullableString** | My Domain of your org, such as &#x60;acme.my.salesforce.com&#x60;, for &#x60;oauth_client_credentials&#x60;. The &#x60;.salesforce.com&#x60; suffix is optional. | 

## Methods

### NewSalesforceCrmCredentialsOut

`func NewSalesforceCrmCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, authMode SalesforceCrmAuthMode, username NullableString, consumerKey NullableString, domain NullableString, ) *SalesforceCrmCredentialsOut`

NewSalesforceCrmCredentialsOut instantiates a new SalesforceCrmCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceCrmCredentialsOutWithDefaults

`func NewSalesforceCrmCredentialsOutWithDefaults() *SalesforceCrmCredentialsOut`

NewSalesforceCrmCredentialsOutWithDefaults instantiates a new SalesforceCrmCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SalesforceCrmCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalesforceCrmCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalesforceCrmCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *SalesforceCrmCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SalesforceCrmCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SalesforceCrmCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *SalesforceCrmCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *SalesforceCrmCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *SalesforceCrmCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *SalesforceCrmCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SalesforceCrmCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SalesforceCrmCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetAuthMode

`func (o *SalesforceCrmCredentialsOut) GetAuthMode() SalesforceCrmAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SalesforceCrmCredentialsOut) GetAuthModeOk() (*SalesforceCrmAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SalesforceCrmCredentialsOut) SetAuthMode(v SalesforceCrmAuthMode)`

SetAuthMode sets AuthMode field to given value.


### GetUsername

`func (o *SalesforceCrmCredentialsOut) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SalesforceCrmCredentialsOut) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SalesforceCrmCredentialsOut) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *SalesforceCrmCredentialsOut) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *SalesforceCrmCredentialsOut) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetConsumerKey

`func (o *SalesforceCrmCredentialsOut) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *SalesforceCrmCredentialsOut) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *SalesforceCrmCredentialsOut) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.


### SetConsumerKeyNil

`func (o *SalesforceCrmCredentialsOut) SetConsumerKeyNil(b bool)`

 SetConsumerKeyNil sets the value for ConsumerKey to be an explicit nil

### UnsetConsumerKey
`func (o *SalesforceCrmCredentialsOut) UnsetConsumerKey()`

UnsetConsumerKey ensures that no value is present for ConsumerKey, not even an explicit nil
### GetDomain

`func (o *SalesforceCrmCredentialsOut) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SalesforceCrmCredentialsOut) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SalesforceCrmCredentialsOut) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *SalesforceCrmCredentialsOut) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *SalesforceCrmCredentialsOut) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


