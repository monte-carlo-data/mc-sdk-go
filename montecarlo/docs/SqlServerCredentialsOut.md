# SqlServerCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**Host** | **string** | Hostname of the database endpoint. For &#x60;kerberos&#x60;, the fully qualified name the server&#39;s service principal is registered under. | 
**Port** | **int32** | Port the database listens on. | 
**DbName** | **NullableString** | Database to connect to. Null when none is set. | 
**AuthMode** | [**SqlServerAuthMode**](SqlServerAuthMode.md) | How Monte Carlo signs in. &#x60;sql&#x60; takes &#x60;user&#x60; and &#x60;password&#x60;. &#x60;kerberos&#x60; takes &#x60;realm&#x60;, &#x60;kdc&#x60;, &#x60;principal&#x60;, and one of &#x60;keytab_base64&#x60; or &#x60;password&#x60;. | 
**User** | **NullableString** | SQL login Monte Carlo signs in as, for &#x60;sql&#x60;. Null for &#x60;kerberos&#x60;. | 
**Realm** | **NullableString** | Kerberos realm, normally the Active Directory domain in capitals, for &#x60;kerberos&#x60;. Null for &#x60;sql&#x60;. | 
**Kdc** | **NullableString** | Hostname of the key distribution center, optionally with &#x60;:port&#x60;, for &#x60;kerberos&#x60;. Null for &#x60;sql&#x60;. | 
**Principal** | **NullableString** | Active Directory principal Monte Carlo signs in as, for &#x60;kerberos&#x60;. Null for &#x60;sql&#x60;. | 

## Methods

### NewSqlServerCredentialsOut

`func NewSqlServerCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, host string, port int32, dbName NullableString, authMode SqlServerAuthMode, user NullableString, realm NullableString, kdc NullableString, principal NullableString, ) *SqlServerCredentialsOut`

NewSqlServerCredentialsOut instantiates a new SqlServerCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlServerCredentialsOutWithDefaults

`func NewSqlServerCredentialsOutWithDefaults() *SqlServerCredentialsOut`

NewSqlServerCredentialsOutWithDefaults instantiates a new SqlServerCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SqlServerCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SqlServerCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SqlServerCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *SqlServerCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SqlServerCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SqlServerCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *SqlServerCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *SqlServerCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *SqlServerCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *SqlServerCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SqlServerCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SqlServerCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetHost

`func (o *SqlServerCredentialsOut) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SqlServerCredentialsOut) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SqlServerCredentialsOut) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *SqlServerCredentialsOut) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SqlServerCredentialsOut) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SqlServerCredentialsOut) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDbName

`func (o *SqlServerCredentialsOut) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *SqlServerCredentialsOut) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *SqlServerCredentialsOut) SetDbName(v string)`

SetDbName sets DbName field to given value.


### SetDbNameNil

`func (o *SqlServerCredentialsOut) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *SqlServerCredentialsOut) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetAuthMode

`func (o *SqlServerCredentialsOut) GetAuthMode() SqlServerAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SqlServerCredentialsOut) GetAuthModeOk() (*SqlServerAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SqlServerCredentialsOut) SetAuthMode(v SqlServerAuthMode)`

SetAuthMode sets AuthMode field to given value.


### GetUser

`func (o *SqlServerCredentialsOut) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SqlServerCredentialsOut) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SqlServerCredentialsOut) SetUser(v string)`

SetUser sets User field to given value.


### SetUserNil

`func (o *SqlServerCredentialsOut) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SqlServerCredentialsOut) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetRealm

`func (o *SqlServerCredentialsOut) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *SqlServerCredentialsOut) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *SqlServerCredentialsOut) SetRealm(v string)`

SetRealm sets Realm field to given value.


### SetRealmNil

`func (o *SqlServerCredentialsOut) SetRealmNil(b bool)`

 SetRealmNil sets the value for Realm to be an explicit nil

### UnsetRealm
`func (o *SqlServerCredentialsOut) UnsetRealm()`

UnsetRealm ensures that no value is present for Realm, not even an explicit nil
### GetKdc

`func (o *SqlServerCredentialsOut) GetKdc() string`

GetKdc returns the Kdc field if non-nil, zero value otherwise.

### GetKdcOk

`func (o *SqlServerCredentialsOut) GetKdcOk() (*string, bool)`

GetKdcOk returns a tuple with the Kdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdc

`func (o *SqlServerCredentialsOut) SetKdc(v string)`

SetKdc sets Kdc field to given value.


### SetKdcNil

`func (o *SqlServerCredentialsOut) SetKdcNil(b bool)`

 SetKdcNil sets the value for Kdc to be an explicit nil

### UnsetKdc
`func (o *SqlServerCredentialsOut) UnsetKdc()`

UnsetKdc ensures that no value is present for Kdc, not even an explicit nil
### GetPrincipal

`func (o *SqlServerCredentialsOut) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *SqlServerCredentialsOut) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *SqlServerCredentialsOut) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### SetPrincipalNil

`func (o *SqlServerCredentialsOut) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *SqlServerCredentialsOut) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


