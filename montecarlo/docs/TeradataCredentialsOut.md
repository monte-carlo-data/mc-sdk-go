# TeradataCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**Host** | **string** | Hostname of the database endpoint. | 
**Port** | **int32** | Port the database listens on. | 
**DbName** | **NullableString** | Database to connect to. Null when none is set. | 
**User** | **string** | Database user Monte Carlo logs in as. | 
**SslCaData** | **NullableString** | PEM text of the CA certificate the server&#39;s certificate is checked against. Null when none is set. | 
**SslDisabled** | **NullableBool** | Do not check the server against &#x60;ssl_ca_data&#x60;. &#x60;td_sslmode&#x60; decides whether the connection uses TLS. Null when unset. | 
**TdSslmode** | [**NullableTeradataSslMode**](TeradataSslMode.md) | How the connection to Teradata uses TLS. Null when unset. | 
**TdLogmech** | [**NullableTeradataLogonMechanism**](TeradataLogonMechanism.md) | How Teradata authenticates the user. Null when unset. | 

## Methods

### NewTeradataCredentialsOut

`func NewTeradataCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, host string, port int32, dbName NullableString, user string, sslCaData NullableString, sslDisabled NullableBool, tdSslmode NullableTeradataSslMode, tdLogmech NullableTeradataLogonMechanism, ) *TeradataCredentialsOut`

NewTeradataCredentialsOut instantiates a new TeradataCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeradataCredentialsOutWithDefaults

`func NewTeradataCredentialsOutWithDefaults() *TeradataCredentialsOut`

NewTeradataCredentialsOutWithDefaults instantiates a new TeradataCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeradataCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeradataCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeradataCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *TeradataCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *TeradataCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *TeradataCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *TeradataCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *TeradataCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *TeradataCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *TeradataCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *TeradataCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *TeradataCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetHost

`func (o *TeradataCredentialsOut) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TeradataCredentialsOut) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TeradataCredentialsOut) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *TeradataCredentialsOut) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TeradataCredentialsOut) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TeradataCredentialsOut) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDbName

`func (o *TeradataCredentialsOut) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *TeradataCredentialsOut) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *TeradataCredentialsOut) SetDbName(v string)`

SetDbName sets DbName field to given value.


### SetDbNameNil

`func (o *TeradataCredentialsOut) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *TeradataCredentialsOut) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetUser

`func (o *TeradataCredentialsOut) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TeradataCredentialsOut) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TeradataCredentialsOut) SetUser(v string)`

SetUser sets User field to given value.


### GetSslCaData

`func (o *TeradataCredentialsOut) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *TeradataCredentialsOut) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *TeradataCredentialsOut) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.


### SetSslCaDataNil

`func (o *TeradataCredentialsOut) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *TeradataCredentialsOut) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslDisabled

`func (o *TeradataCredentialsOut) GetSslDisabled() bool`

GetSslDisabled returns the SslDisabled field if non-nil, zero value otherwise.

### GetSslDisabledOk

`func (o *TeradataCredentialsOut) GetSslDisabledOk() (*bool, bool)`

GetSslDisabledOk returns a tuple with the SslDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDisabled

`func (o *TeradataCredentialsOut) SetSslDisabled(v bool)`

SetSslDisabled sets SslDisabled field to given value.


### SetSslDisabledNil

`func (o *TeradataCredentialsOut) SetSslDisabledNil(b bool)`

 SetSslDisabledNil sets the value for SslDisabled to be an explicit nil

### UnsetSslDisabled
`func (o *TeradataCredentialsOut) UnsetSslDisabled()`

UnsetSslDisabled ensures that no value is present for SslDisabled, not even an explicit nil
### GetTdSslmode

`func (o *TeradataCredentialsOut) GetTdSslmode() TeradataSslMode`

GetTdSslmode returns the TdSslmode field if non-nil, zero value otherwise.

### GetTdSslmodeOk

`func (o *TeradataCredentialsOut) GetTdSslmodeOk() (*TeradataSslMode, bool)`

GetTdSslmodeOk returns a tuple with the TdSslmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdSslmode

`func (o *TeradataCredentialsOut) SetTdSslmode(v TeradataSslMode)`

SetTdSslmode sets TdSslmode field to given value.


### SetTdSslmodeNil

`func (o *TeradataCredentialsOut) SetTdSslmodeNil(b bool)`

 SetTdSslmodeNil sets the value for TdSslmode to be an explicit nil

### UnsetTdSslmode
`func (o *TeradataCredentialsOut) UnsetTdSslmode()`

UnsetTdSslmode ensures that no value is present for TdSslmode, not even an explicit nil
### GetTdLogmech

`func (o *TeradataCredentialsOut) GetTdLogmech() TeradataLogonMechanism`

GetTdLogmech returns the TdLogmech field if non-nil, zero value otherwise.

### GetTdLogmechOk

`func (o *TeradataCredentialsOut) GetTdLogmechOk() (*TeradataLogonMechanism, bool)`

GetTdLogmechOk returns a tuple with the TdLogmech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdLogmech

`func (o *TeradataCredentialsOut) SetTdLogmech(v TeradataLogonMechanism)`

SetTdLogmech sets TdLogmech field to given value.


### SetTdLogmechNil

`func (o *TeradataCredentialsOut) SetTdLogmechNil(b bool)`

 SetTdLogmechNil sets the value for TdLogmech to be an explicit nil

### UnsetTdLogmech
`func (o *TeradataCredentialsOut) UnsetTdLogmech()`

UnsetTdLogmech ensures that no value is present for TdLogmech, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


