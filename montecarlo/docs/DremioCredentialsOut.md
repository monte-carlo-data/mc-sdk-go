# DremioCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**Host** | **string** | Hostname of the Dremio Arrow Flight endpoint, such as &#x60;data.dremio.cloud&#x60; for Dremio Cloud (&#x60;data.eu.dremio.cloud&#x60; in the EU). | 
**Port** | **int32** | Port Dremio serves Arrow Flight on. | 
**DbName** | **NullableString** | Dremio project or space to connect to. Null when none is set. | 
**Tls** | **NullableBool** | Connect over TLS. Dremio Cloud requires it. Null when unset. | 

## Methods

### NewDremioCredentialsOut

`func NewDremioCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, host string, port int32, dbName NullableString, tls NullableBool, ) *DremioCredentialsOut`

NewDremioCredentialsOut instantiates a new DremioCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDremioCredentialsOutWithDefaults

`func NewDremioCredentialsOutWithDefaults() *DremioCredentialsOut`

NewDremioCredentialsOutWithDefaults instantiates a new DremioCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DremioCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DremioCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DremioCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *DremioCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *DremioCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *DremioCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *DremioCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DremioCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DremioCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *DremioCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DremioCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DremioCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetHost

`func (o *DremioCredentialsOut) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DremioCredentialsOut) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DremioCredentialsOut) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *DremioCredentialsOut) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DremioCredentialsOut) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DremioCredentialsOut) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDbName

`func (o *DremioCredentialsOut) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DremioCredentialsOut) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DremioCredentialsOut) SetDbName(v string)`

SetDbName sets DbName field to given value.


### SetDbNameNil

`func (o *DremioCredentialsOut) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *DremioCredentialsOut) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetTls

`func (o *DremioCredentialsOut) GetTls() bool`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *DremioCredentialsOut) GetTlsOk() (*bool, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *DremioCredentialsOut) SetTls(v bool)`

SetTls sets Tls field to given value.


### SetTlsNil

`func (o *DremioCredentialsOut) SetTlsNil(b bool)`

 SetTlsNil sets the value for Tls to be an explicit nil

### UnsetTls
`func (o *DremioCredentialsOut) UnsetTls()`

UnsetTls ensures that no value is present for Tls, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


