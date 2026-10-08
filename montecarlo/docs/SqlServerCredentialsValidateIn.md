# SqlServerCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Host** | **string** | Hostname of the database endpoint. For &#x60;kerberos&#x60;, the fully qualified name the server&#39;s service principal is registered under. | 
**Port** | **int32** | Port the database listens on. | 
**DbName** | Pointer to **NullableString** | Database to connect to. | [optional] 
**AuthMode** | Pointer to [**SqlServerAuthMode**](SqlServerAuthMode.md) | How Monte Carlo signs in. &#x60;sql&#x60; takes &#x60;user&#x60; and &#x60;password&#x60;. &#x60;kerberos&#x60; takes &#x60;realm&#x60;, &#x60;kdc&#x60;, &#x60;principal&#x60;, and one of &#x60;keytab_base64&#x60; or &#x60;password&#x60;. | [optional] [default to SQLSERVERAUTHMODE_SQL]
**User** | Pointer to **NullableString** | SQL login Monte Carlo signs in as, for &#x60;sql&#x60;. | [optional] 
**Realm** | Pointer to **NullableString** | Kerberos realm, normally the Active Directory domain in capitals, for &#x60;kerberos&#x60;. | [optional] 
**Kdc** | Pointer to **NullableString** | Hostname of the key distribution center, optionally with &#x60;:port&#x60;, for &#x60;kerberos&#x60;. | [optional] 
**Principal** | Pointer to **NullableString** | Active Directory principal Monte Carlo signs in as, for &#x60;kerberos&#x60;. | [optional] 
**Password** | Pointer to **NullableString** | Password of &#x60;user&#x60; for &#x60;sql&#x60;, or of the Active Directory account behind &#x60;principal&#x60; for &#x60;kerberos&#x60;. Used for this check and not kept. | [optional] 
**KeytabBase64** | Pointer to **NullableString** | Base64-encoded keytab for &#x60;principal&#x60;, for &#x60;kerberos&#x60;. Send this or &#x60;password&#x60;. Used for this check and not kept. | [optional] 

## Methods

### NewSqlServerCredentialsValidateIn

`func NewSqlServerCredentialsValidateIn(deploymentId string, host string, port int32, ) *SqlServerCredentialsValidateIn`

NewSqlServerCredentialsValidateIn instantiates a new SqlServerCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlServerCredentialsValidateInWithDefaults

`func NewSqlServerCredentialsValidateInWithDefaults() *SqlServerCredentialsValidateIn`

NewSqlServerCredentialsValidateInWithDefaults instantiates a new SqlServerCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *SqlServerCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *SqlServerCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *SqlServerCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetHost

`func (o *SqlServerCredentialsValidateIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SqlServerCredentialsValidateIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SqlServerCredentialsValidateIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *SqlServerCredentialsValidateIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SqlServerCredentialsValidateIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SqlServerCredentialsValidateIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDbName

`func (o *SqlServerCredentialsValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *SqlServerCredentialsValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *SqlServerCredentialsValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *SqlServerCredentialsValidateIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *SqlServerCredentialsValidateIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *SqlServerCredentialsValidateIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetAuthMode

`func (o *SqlServerCredentialsValidateIn) GetAuthMode() SqlServerAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SqlServerCredentialsValidateIn) GetAuthModeOk() (*SqlServerAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SqlServerCredentialsValidateIn) SetAuthMode(v SqlServerAuthMode)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *SqlServerCredentialsValidateIn) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetUser

`func (o *SqlServerCredentialsValidateIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SqlServerCredentialsValidateIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SqlServerCredentialsValidateIn) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SqlServerCredentialsValidateIn) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *SqlServerCredentialsValidateIn) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SqlServerCredentialsValidateIn) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetRealm

`func (o *SqlServerCredentialsValidateIn) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *SqlServerCredentialsValidateIn) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *SqlServerCredentialsValidateIn) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *SqlServerCredentialsValidateIn) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### SetRealmNil

`func (o *SqlServerCredentialsValidateIn) SetRealmNil(b bool)`

 SetRealmNil sets the value for Realm to be an explicit nil

### UnsetRealm
`func (o *SqlServerCredentialsValidateIn) UnsetRealm()`

UnsetRealm ensures that no value is present for Realm, not even an explicit nil
### GetKdc

`func (o *SqlServerCredentialsValidateIn) GetKdc() string`

GetKdc returns the Kdc field if non-nil, zero value otherwise.

### GetKdcOk

`func (o *SqlServerCredentialsValidateIn) GetKdcOk() (*string, bool)`

GetKdcOk returns a tuple with the Kdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdc

`func (o *SqlServerCredentialsValidateIn) SetKdc(v string)`

SetKdc sets Kdc field to given value.

### HasKdc

`func (o *SqlServerCredentialsValidateIn) HasKdc() bool`

HasKdc returns a boolean if a field has been set.

### SetKdcNil

`func (o *SqlServerCredentialsValidateIn) SetKdcNil(b bool)`

 SetKdcNil sets the value for Kdc to be an explicit nil

### UnsetKdc
`func (o *SqlServerCredentialsValidateIn) UnsetKdc()`

UnsetKdc ensures that no value is present for Kdc, not even an explicit nil
### GetPrincipal

`func (o *SqlServerCredentialsValidateIn) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *SqlServerCredentialsValidateIn) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *SqlServerCredentialsValidateIn) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *SqlServerCredentialsValidateIn) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### SetPrincipalNil

`func (o *SqlServerCredentialsValidateIn) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *SqlServerCredentialsValidateIn) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil
### GetPassword

`func (o *SqlServerCredentialsValidateIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SqlServerCredentialsValidateIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SqlServerCredentialsValidateIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SqlServerCredentialsValidateIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SqlServerCredentialsValidateIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SqlServerCredentialsValidateIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetKeytabBase64

`func (o *SqlServerCredentialsValidateIn) GetKeytabBase64() string`

GetKeytabBase64 returns the KeytabBase64 field if non-nil, zero value otherwise.

### GetKeytabBase64Ok

`func (o *SqlServerCredentialsValidateIn) GetKeytabBase64Ok() (*string, bool)`

GetKeytabBase64Ok returns a tuple with the KeytabBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytabBase64

`func (o *SqlServerCredentialsValidateIn) SetKeytabBase64(v string)`

SetKeytabBase64 sets KeytabBase64 field to given value.

### HasKeytabBase64

`func (o *SqlServerCredentialsValidateIn) HasKeytabBase64() bool`

HasKeytabBase64 returns a boolean if a field has been set.

### SetKeytabBase64Nil

`func (o *SqlServerCredentialsValidateIn) SetKeytabBase64Nil(b bool)`

 SetKeytabBase64Nil sets the value for KeytabBase64 to be an explicit nil

### UnsetKeytabBase64
`func (o *SqlServerCredentialsValidateIn) UnsetKeytabBase64()`

UnsetKeytabBase64 ensures that no value is present for KeytabBase64, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


