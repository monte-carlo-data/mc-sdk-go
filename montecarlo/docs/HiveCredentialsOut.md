# HiveCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**Host** | **string** | Hostname of the HiveServer2. | 
**Port** | **NullableInt32** | Port the HiveServer2 listens on in binary mode. Null means the client&#39;s default. | 
**User** | **string** | User Monte Carlo signs in as. Hive takes no password. | 
**DbName** | **NullableString** | Database Monte Carlo connects to. Null means the default database. | 
**AuthMode** | [**HiveAuthMode**](HiveAuthMode.md) | How Monte Carlo signs in. | 

## Methods

### NewHiveCredentialsOut

`func NewHiveCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, host string, port NullableInt32, user string, dbName NullableString, authMode HiveAuthMode, ) *HiveCredentialsOut`

NewHiveCredentialsOut instantiates a new HiveCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveCredentialsOutWithDefaults

`func NewHiveCredentialsOutWithDefaults() *HiveCredentialsOut`

NewHiveCredentialsOutWithDefaults instantiates a new HiveCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HiveCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HiveCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HiveCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *HiveCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HiveCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HiveCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *HiveCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *HiveCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *HiveCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *HiveCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *HiveCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *HiveCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetHost

`func (o *HiveCredentialsOut) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HiveCredentialsOut) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HiveCredentialsOut) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *HiveCredentialsOut) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HiveCredentialsOut) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HiveCredentialsOut) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *HiveCredentialsOut) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *HiveCredentialsOut) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUser

`func (o *HiveCredentialsOut) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *HiveCredentialsOut) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *HiveCredentialsOut) SetUser(v string)`

SetUser sets User field to given value.


### GetDbName

`func (o *HiveCredentialsOut) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *HiveCredentialsOut) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *HiveCredentialsOut) SetDbName(v string)`

SetDbName sets DbName field to given value.


### SetDbNameNil

`func (o *HiveCredentialsOut) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *HiveCredentialsOut) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetAuthMode

`func (o *HiveCredentialsOut) GetAuthMode() HiveAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *HiveCredentialsOut) GetAuthModeOk() (*HiveAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *HiveCredentialsOut) SetAuthMode(v HiveAuthMode)`

SetAuthMode sets AuthMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


