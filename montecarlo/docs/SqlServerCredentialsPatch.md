# SqlServerCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | Hostname of the database endpoint. For &#x60;kerberos&#x60;, the fully qualified name the server&#39;s service principal is registered under. | [optional] 
**Port** | Pointer to **NullableInt32** | Port the database listens on. | [optional] 
**DbName** | Pointer to **NullableString** | Database to connect to. | [optional] 
**AuthMode** | Pointer to [**NullableSqlServerAuthMode**](SqlServerAuthMode.md) | How Monte Carlo signs in. &#x60;sql&#x60; takes &#x60;user&#x60; and &#x60;password&#x60;. &#x60;kerberos&#x60; takes &#x60;realm&#x60;, &#x60;kdc&#x60;, &#x60;principal&#x60;, and one of &#x60;keytab_base64&#x60; or &#x60;password&#x60;. | [optional] 
**User** | Pointer to **NullableString** | SQL login Monte Carlo signs in as, for &#x60;sql&#x60;. | [optional] 
**Realm** | Pointer to **NullableString** | Kerberos realm, normally the Active Directory domain in capitals, for &#x60;kerberos&#x60;. | [optional] 
**Kdc** | Pointer to **NullableString** | Hostname of the key distribution center, optionally with &#x60;:port&#x60;, for &#x60;kerberos&#x60;. | [optional] 
**Principal** | Pointer to **NullableString** | Active Directory principal Monte Carlo signs in as, for &#x60;kerberos&#x60;. | [optional] 
**Password** | Pointer to **NullableString** | Password of &#x60;user&#x60; for &#x60;sql&#x60;, or of the Active Directory account behind &#x60;principal&#x60; for &#x60;kerberos&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**KeytabBase64** | Pointer to **NullableString** | Base64-encoded keytab for &#x60;principal&#x60;, for &#x60;kerberos&#x60;. Send this or &#x60;password&#x60;. Stored by Monte Carlo and never returned. | [optional] 

## Methods

### NewSqlServerCredentialsPatch

`func NewSqlServerCredentialsPatch() *SqlServerCredentialsPatch`

NewSqlServerCredentialsPatch instantiates a new SqlServerCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlServerCredentialsPatchWithDefaults

`func NewSqlServerCredentialsPatchWithDefaults() *SqlServerCredentialsPatch`

NewSqlServerCredentialsPatchWithDefaults instantiates a new SqlServerCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SqlServerCredentialsPatch) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SqlServerCredentialsPatch) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SqlServerCredentialsPatch) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SqlServerCredentialsPatch) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SqlServerCredentialsPatch) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SqlServerCredentialsPatch) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *SqlServerCredentialsPatch) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SqlServerCredentialsPatch) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SqlServerCredentialsPatch) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SqlServerCredentialsPatch) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *SqlServerCredentialsPatch) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *SqlServerCredentialsPatch) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *SqlServerCredentialsPatch) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *SqlServerCredentialsPatch) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *SqlServerCredentialsPatch) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *SqlServerCredentialsPatch) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *SqlServerCredentialsPatch) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *SqlServerCredentialsPatch) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetAuthMode

`func (o *SqlServerCredentialsPatch) GetAuthMode() SqlServerAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SqlServerCredentialsPatch) GetAuthModeOk() (*SqlServerAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SqlServerCredentialsPatch) SetAuthMode(v SqlServerAuthMode)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *SqlServerCredentialsPatch) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### SetAuthModeNil

`func (o *SqlServerCredentialsPatch) SetAuthModeNil(b bool)`

 SetAuthModeNil sets the value for AuthMode to be an explicit nil

### UnsetAuthMode
`func (o *SqlServerCredentialsPatch) UnsetAuthMode()`

UnsetAuthMode ensures that no value is present for AuthMode, not even an explicit nil
### GetUser

`func (o *SqlServerCredentialsPatch) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SqlServerCredentialsPatch) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SqlServerCredentialsPatch) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SqlServerCredentialsPatch) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *SqlServerCredentialsPatch) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SqlServerCredentialsPatch) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetRealm

`func (o *SqlServerCredentialsPatch) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *SqlServerCredentialsPatch) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *SqlServerCredentialsPatch) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *SqlServerCredentialsPatch) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### SetRealmNil

`func (o *SqlServerCredentialsPatch) SetRealmNil(b bool)`

 SetRealmNil sets the value for Realm to be an explicit nil

### UnsetRealm
`func (o *SqlServerCredentialsPatch) UnsetRealm()`

UnsetRealm ensures that no value is present for Realm, not even an explicit nil
### GetKdc

`func (o *SqlServerCredentialsPatch) GetKdc() string`

GetKdc returns the Kdc field if non-nil, zero value otherwise.

### GetKdcOk

`func (o *SqlServerCredentialsPatch) GetKdcOk() (*string, bool)`

GetKdcOk returns a tuple with the Kdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdc

`func (o *SqlServerCredentialsPatch) SetKdc(v string)`

SetKdc sets Kdc field to given value.

### HasKdc

`func (o *SqlServerCredentialsPatch) HasKdc() bool`

HasKdc returns a boolean if a field has been set.

### SetKdcNil

`func (o *SqlServerCredentialsPatch) SetKdcNil(b bool)`

 SetKdcNil sets the value for Kdc to be an explicit nil

### UnsetKdc
`func (o *SqlServerCredentialsPatch) UnsetKdc()`

UnsetKdc ensures that no value is present for Kdc, not even an explicit nil
### GetPrincipal

`func (o *SqlServerCredentialsPatch) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *SqlServerCredentialsPatch) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *SqlServerCredentialsPatch) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *SqlServerCredentialsPatch) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### SetPrincipalNil

`func (o *SqlServerCredentialsPatch) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *SqlServerCredentialsPatch) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil
### GetPassword

`func (o *SqlServerCredentialsPatch) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SqlServerCredentialsPatch) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SqlServerCredentialsPatch) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SqlServerCredentialsPatch) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SqlServerCredentialsPatch) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SqlServerCredentialsPatch) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetKeytabBase64

`func (o *SqlServerCredentialsPatch) GetKeytabBase64() string`

GetKeytabBase64 returns the KeytabBase64 field if non-nil, zero value otherwise.

### GetKeytabBase64Ok

`func (o *SqlServerCredentialsPatch) GetKeytabBase64Ok() (*string, bool)`

GetKeytabBase64Ok returns a tuple with the KeytabBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytabBase64

`func (o *SqlServerCredentialsPatch) SetKeytabBase64(v string)`

SetKeytabBase64 sets KeytabBase64 field to given value.

### HasKeytabBase64

`func (o *SqlServerCredentialsPatch) HasKeytabBase64() bool`

HasKeytabBase64 returns a boolean if a field has been set.

### SetKeytabBase64Nil

`func (o *SqlServerCredentialsPatch) SetKeytabBase64Nil(b bool)`

 SetKeytabBase64Nil sets the value for KeytabBase64 to be an explicit nil

### UnsetKeytabBase64
`func (o *SqlServerCredentialsPatch) UnsetKeytabBase64()`

UnsetKeytabBase64 ensures that no value is present for KeytabBase64, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


