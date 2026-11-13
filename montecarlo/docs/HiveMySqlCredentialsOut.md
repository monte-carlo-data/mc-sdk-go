# HiveMySqlCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**Host** | **string** | Hostname of the database endpoint. | 
**Port** | **int32** | Port the database listens on. | 
**DbName** | **string** | MySQL database that holds the Hive metastore. | 
**User** | **string** | Database user Monte Carlo logs in as. | 
**SslCaUrl** | **NullableString** | https URL of the PEM bundle of certificate authorities the MySQL server&#39;s certificate is checked against. Monte Carlo downloads it at each connection and connects over TLS. For Amazon RDS, use https://truststore.pki.rds.amazonaws.com/global/global-bundle.pem. Null when none is set, or when the stored CA is not fetched from a URL and checked, as when it was set up to read from Monte Carlo storage or to skip the check. Setting &#x60;ssl_ca_url&#x60; replaces such a setup. | 

## Methods

### NewHiveMySqlCredentialsOut

`func NewHiveMySqlCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, host string, port int32, dbName string, user string, sslCaUrl NullableString, ) *HiveMySqlCredentialsOut`

NewHiveMySqlCredentialsOut instantiates a new HiveMySqlCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveMySqlCredentialsOutWithDefaults

`func NewHiveMySqlCredentialsOutWithDefaults() *HiveMySqlCredentialsOut`

NewHiveMySqlCredentialsOutWithDefaults instantiates a new HiveMySqlCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HiveMySqlCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HiveMySqlCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HiveMySqlCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *HiveMySqlCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *HiveMySqlCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *HiveMySqlCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *HiveMySqlCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *HiveMySqlCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *HiveMySqlCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *HiveMySqlCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *HiveMySqlCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *HiveMySqlCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetHost

`func (o *HiveMySqlCredentialsOut) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HiveMySqlCredentialsOut) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HiveMySqlCredentialsOut) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *HiveMySqlCredentialsOut) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HiveMySqlCredentialsOut) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HiveMySqlCredentialsOut) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDbName

`func (o *HiveMySqlCredentialsOut) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *HiveMySqlCredentialsOut) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *HiveMySqlCredentialsOut) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetUser

`func (o *HiveMySqlCredentialsOut) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *HiveMySqlCredentialsOut) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *HiveMySqlCredentialsOut) SetUser(v string)`

SetUser sets User field to given value.


### GetSslCaUrl

`func (o *HiveMySqlCredentialsOut) GetSslCaUrl() string`

GetSslCaUrl returns the SslCaUrl field if non-nil, zero value otherwise.

### GetSslCaUrlOk

`func (o *HiveMySqlCredentialsOut) GetSslCaUrlOk() (*string, bool)`

GetSslCaUrlOk returns a tuple with the SslCaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaUrl

`func (o *HiveMySqlCredentialsOut) SetSslCaUrl(v string)`

SetSslCaUrl sets SslCaUrl field to given value.


### SetSslCaUrlNil

`func (o *HiveMySqlCredentialsOut) SetSslCaUrlNil(b bool)`

 SetSslCaUrlNil sets the value for SslCaUrl to be an explicit nil

### UnsetSslCaUrl
`func (o *HiveMySqlCredentialsOut) UnsetSslCaUrl()`

UnsetSslCaUrl ensures that no value is present for SslCaUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


