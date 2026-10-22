# SalesforceDataCloudCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**AuthMode** | [**SalesforceDataCloudAuthMode**](SalesforceDataCloudAuthMode.md) | How Monte Carlo signs in. &#x60;oauth_client_credentials&#x60; takes &#x60;consumer_secret&#x60;. &#x60;jwt_bearer&#x60; takes &#x60;username&#x60; and &#x60;private_key&#x60;. Both take &#x60;domain&#x60; and &#x60;consumer_key&#x60;. | 
**Domain** | **string** | My Domain host of your org, such as &#x60;acme.my.salesforce.com&#x60;. | 
**ConsumerKey** | **string** | Consumer key of the connected app Monte Carlo signs in as. | 
**Username** | **NullableString** | Salesforce user the JWT signs in as, for &#x60;jwt_bearer&#x60;. | 
**Dataspaces** | **[]string** | Data spaces Monte Carlo collects metadata from. | 
**Dataspace** | **NullableString** | The one data space a query connection runs in. Once set, it can be changed but not cleared. Sending it in an update, even unchanged, moves row count collection for that data space&#39;s tables to the query connections using these credentials. | 

## Methods

### NewSalesforceDataCloudCredentialsOut

`func NewSalesforceDataCloudCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, authMode SalesforceDataCloudAuthMode, domain string, consumerKey string, username NullableString, dataspaces []string, dataspace NullableString, ) *SalesforceDataCloudCredentialsOut`

NewSalesforceDataCloudCredentialsOut instantiates a new SalesforceDataCloudCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceDataCloudCredentialsOutWithDefaults

`func NewSalesforceDataCloudCredentialsOutWithDefaults() *SalesforceDataCloudCredentialsOut`

NewSalesforceDataCloudCredentialsOutWithDefaults instantiates a new SalesforceDataCloudCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SalesforceDataCloudCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalesforceDataCloudCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalesforceDataCloudCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *SalesforceDataCloudCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SalesforceDataCloudCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SalesforceDataCloudCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *SalesforceDataCloudCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *SalesforceDataCloudCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *SalesforceDataCloudCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *SalesforceDataCloudCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SalesforceDataCloudCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SalesforceDataCloudCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetAuthMode

`func (o *SalesforceDataCloudCredentialsOut) GetAuthMode() SalesforceDataCloudAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SalesforceDataCloudCredentialsOut) GetAuthModeOk() (*SalesforceDataCloudAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SalesforceDataCloudCredentialsOut) SetAuthMode(v SalesforceDataCloudAuthMode)`

SetAuthMode sets AuthMode field to given value.


### GetDomain

`func (o *SalesforceDataCloudCredentialsOut) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SalesforceDataCloudCredentialsOut) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SalesforceDataCloudCredentialsOut) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetConsumerKey

`func (o *SalesforceDataCloudCredentialsOut) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *SalesforceDataCloudCredentialsOut) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *SalesforceDataCloudCredentialsOut) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.


### GetUsername

`func (o *SalesforceDataCloudCredentialsOut) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SalesforceDataCloudCredentialsOut) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SalesforceDataCloudCredentialsOut) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *SalesforceDataCloudCredentialsOut) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *SalesforceDataCloudCredentialsOut) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetDataspaces

`func (o *SalesforceDataCloudCredentialsOut) GetDataspaces() []string`

GetDataspaces returns the Dataspaces field if non-nil, zero value otherwise.

### GetDataspacesOk

`func (o *SalesforceDataCloudCredentialsOut) GetDataspacesOk() (*[]string, bool)`

GetDataspacesOk returns a tuple with the Dataspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataspaces

`func (o *SalesforceDataCloudCredentialsOut) SetDataspaces(v []string)`

SetDataspaces sets Dataspaces field to given value.


### GetDataspace

`func (o *SalesforceDataCloudCredentialsOut) GetDataspace() string`

GetDataspace returns the Dataspace field if non-nil, zero value otherwise.

### GetDataspaceOk

`func (o *SalesforceDataCloudCredentialsOut) GetDataspaceOk() (*string, bool)`

GetDataspaceOk returns a tuple with the Dataspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataspace

`func (o *SalesforceDataCloudCredentialsOut) SetDataspace(v string)`

SetDataspace sets Dataspace field to given value.


### SetDataspaceNil

`func (o *SalesforceDataCloudCredentialsOut) SetDataspaceNil(b bool)`

 SetDataspaceNil sets the value for Dataspace to be an explicit nil

### UnsetDataspace
`func (o *SalesforceDataCloudCredentialsOut) UnsetDataspace()`

UnsetDataspace ensures that no value is present for Dataspace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


