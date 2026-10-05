# FivetranCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**BaseUrl** | **NullableString** | URL of the Fivetran REST API. Leave it out for https://api.fivetran.com/v1/. Null unless set. | 

## Methods

### NewFivetranCredentialsOut

`func NewFivetranCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, baseUrl NullableString, ) *FivetranCredentialsOut`

NewFivetranCredentialsOut instantiates a new FivetranCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFivetranCredentialsOutWithDefaults

`func NewFivetranCredentialsOutWithDefaults() *FivetranCredentialsOut`

NewFivetranCredentialsOutWithDefaults instantiates a new FivetranCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FivetranCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FivetranCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FivetranCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *FivetranCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *FivetranCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *FivetranCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *FivetranCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *FivetranCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *FivetranCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *FivetranCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *FivetranCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *FivetranCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetBaseUrl

`func (o *FivetranCredentialsOut) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *FivetranCredentialsOut) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *FivetranCredentialsOut) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### SetBaseUrlNil

`func (o *FivetranCredentialsOut) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *FivetranCredentialsOut) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


