# SqlServerCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname of the database endpoint. For &#x60;kerberos&#x60;, the fully qualified name the server&#39;s service principal is registered under. | 
**Port** | **int32** | Port the database listens on. | 
**DbName** | Pointer to **NullableString** | Database to connect to. | [optional] 
**AuthMode** | Pointer to [**SqlServerAuthMode**](SqlServerAuthMode.md) | How Monte Carlo signs in. &#x60;sql&#x60; takes &#x60;user&#x60; and &#x60;password&#x60;. &#x60;kerberos&#x60; takes &#x60;realm&#x60;, &#x60;kdc&#x60;, &#x60;principal&#x60;, and one of &#x60;keytab_base64&#x60; or &#x60;password&#x60;. | [optional] [default to SQLSERVERAUTHMODE_SQL]
**User** | Pointer to **NullableString** | SQL login Monte Carlo signs in as, for &#x60;sql&#x60;. | [optional] 
**Realm** | Pointer to **NullableString** | Kerberos realm, normally the Active Directory domain in capitals, for &#x60;kerberos&#x60;. | [optional] 
**Kdc** | Pointer to **NullableString** | Hostname of the key distribution center, optionally with &#x60;:port&#x60;, for &#x60;kerberos&#x60;. | [optional] 
**Principal** | Pointer to **NullableString** | Active Directory principal Monte Carlo signs in as, for &#x60;kerberos&#x60;. | [optional] 
**Password** | Pointer to **NullableString** | Password of &#x60;user&#x60; for &#x60;sql&#x60;, or of the Active Directory account behind &#x60;principal&#x60; for &#x60;kerberos&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**KeytabBase64** | Pointer to **NullableString** | Base64-encoded keytab for &#x60;principal&#x60;, for &#x60;kerberos&#x60;. Send this or &#x60;password&#x60;. Stored by Monte Carlo and never returned. | [optional] 

## Methods

### NewSqlServerCredentialsIn

`func NewSqlServerCredentialsIn(host string, port int32, ) *SqlServerCredentialsIn`

NewSqlServerCredentialsIn instantiates a new SqlServerCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlServerCredentialsInWithDefaults

`func NewSqlServerCredentialsInWithDefaults() *SqlServerCredentialsIn`

NewSqlServerCredentialsInWithDefaults instantiates a new SqlServerCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SqlServerCredentialsIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SqlServerCredentialsIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SqlServerCredentialsIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *SqlServerCredentialsIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SqlServerCredentialsIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SqlServerCredentialsIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDbName

`func (o *SqlServerCredentialsIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *SqlServerCredentialsIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *SqlServerCredentialsIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *SqlServerCredentialsIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *SqlServerCredentialsIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *SqlServerCredentialsIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetAuthMode

`func (o *SqlServerCredentialsIn) GetAuthMode() SqlServerAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SqlServerCredentialsIn) GetAuthModeOk() (*SqlServerAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SqlServerCredentialsIn) SetAuthMode(v SqlServerAuthMode)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *SqlServerCredentialsIn) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetUser

`func (o *SqlServerCredentialsIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SqlServerCredentialsIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SqlServerCredentialsIn) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SqlServerCredentialsIn) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *SqlServerCredentialsIn) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SqlServerCredentialsIn) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetRealm

`func (o *SqlServerCredentialsIn) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *SqlServerCredentialsIn) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *SqlServerCredentialsIn) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *SqlServerCredentialsIn) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### SetRealmNil

`func (o *SqlServerCredentialsIn) SetRealmNil(b bool)`

 SetRealmNil sets the value for Realm to be an explicit nil

### UnsetRealm
`func (o *SqlServerCredentialsIn) UnsetRealm()`

UnsetRealm ensures that no value is present for Realm, not even an explicit nil
### GetKdc

`func (o *SqlServerCredentialsIn) GetKdc() string`

GetKdc returns the Kdc field if non-nil, zero value otherwise.

### GetKdcOk

`func (o *SqlServerCredentialsIn) GetKdcOk() (*string, bool)`

GetKdcOk returns a tuple with the Kdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdc

`func (o *SqlServerCredentialsIn) SetKdc(v string)`

SetKdc sets Kdc field to given value.

### HasKdc

`func (o *SqlServerCredentialsIn) HasKdc() bool`

HasKdc returns a boolean if a field has been set.

### SetKdcNil

`func (o *SqlServerCredentialsIn) SetKdcNil(b bool)`

 SetKdcNil sets the value for Kdc to be an explicit nil

### UnsetKdc
`func (o *SqlServerCredentialsIn) UnsetKdc()`

UnsetKdc ensures that no value is present for Kdc, not even an explicit nil
### GetPrincipal

`func (o *SqlServerCredentialsIn) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *SqlServerCredentialsIn) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *SqlServerCredentialsIn) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *SqlServerCredentialsIn) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### SetPrincipalNil

`func (o *SqlServerCredentialsIn) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *SqlServerCredentialsIn) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil
### GetPassword

`func (o *SqlServerCredentialsIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SqlServerCredentialsIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SqlServerCredentialsIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SqlServerCredentialsIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SqlServerCredentialsIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SqlServerCredentialsIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetKeytabBase64

`func (o *SqlServerCredentialsIn) GetKeytabBase64() string`

GetKeytabBase64 returns the KeytabBase64 field if non-nil, zero value otherwise.

### GetKeytabBase64Ok

`func (o *SqlServerCredentialsIn) GetKeytabBase64Ok() (*string, bool)`

GetKeytabBase64Ok returns a tuple with the KeytabBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytabBase64

`func (o *SqlServerCredentialsIn) SetKeytabBase64(v string)`

SetKeytabBase64 sets KeytabBase64 field to given value.

### HasKeytabBase64

`func (o *SqlServerCredentialsIn) HasKeytabBase64() bool`

HasKeytabBase64 returns a boolean if a field has been set.

### SetKeytabBase64Nil

`func (o *SqlServerCredentialsIn) SetKeytabBase64Nil(b bool)`

 SetKeytabBase64Nil sets the value for KeytabBase64 to be an explicit nil

### UnsetKeytabBase64
`func (o *SqlServerCredentialsIn) UnsetKeytabBase64()`

UnsetKeytabBase64 ensures that no value is present for KeytabBase64, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


