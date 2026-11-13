# HiveMySqlCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | Hostname of the database endpoint. Changing it requires sending the credentials&#39; secrets again. | [optional] 
**Port** | Pointer to **NullableInt32** | Port the database listens on. | [optional] 
**DbName** | Pointer to **NullableString** | Database to connect to. | [optional] 
**User** | Pointer to **NullableString** | Database user Monte Carlo logs in as. | [optional] 
**Password** | Pointer to **NullableString** | New password of the database user. | [optional] 
**SslCaUrl** | Pointer to **NullableString** | https URL of the PEM bundle of certificate authorities the MySQL server&#39;s certificate is checked against. Monte Carlo downloads it at each connection and connects over TLS. For Amazon RDS, use https://truststore.pki.rds.amazonaws.com/global/global-bundle.pem. A value replaces any setup that read the CA from Monte Carlo storage or skipped the check. Null removes it. Changing it requires sending the credentials&#39; secrets again. | [optional] 

## Methods

### NewHiveMySqlCredentialsPatch

`func NewHiveMySqlCredentialsPatch() *HiveMySqlCredentialsPatch`

NewHiveMySqlCredentialsPatch instantiates a new HiveMySqlCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveMySqlCredentialsPatchWithDefaults

`func NewHiveMySqlCredentialsPatchWithDefaults() *HiveMySqlCredentialsPatch`

NewHiveMySqlCredentialsPatchWithDefaults instantiates a new HiveMySqlCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HiveMySqlCredentialsPatch) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HiveMySqlCredentialsPatch) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HiveMySqlCredentialsPatch) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HiveMySqlCredentialsPatch) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *HiveMySqlCredentialsPatch) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *HiveMySqlCredentialsPatch) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *HiveMySqlCredentialsPatch) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HiveMySqlCredentialsPatch) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HiveMySqlCredentialsPatch) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HiveMySqlCredentialsPatch) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *HiveMySqlCredentialsPatch) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *HiveMySqlCredentialsPatch) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *HiveMySqlCredentialsPatch) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *HiveMySqlCredentialsPatch) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *HiveMySqlCredentialsPatch) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *HiveMySqlCredentialsPatch) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *HiveMySqlCredentialsPatch) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *HiveMySqlCredentialsPatch) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetUser

`func (o *HiveMySqlCredentialsPatch) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *HiveMySqlCredentialsPatch) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *HiveMySqlCredentialsPatch) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *HiveMySqlCredentialsPatch) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *HiveMySqlCredentialsPatch) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *HiveMySqlCredentialsPatch) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *HiveMySqlCredentialsPatch) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HiveMySqlCredentialsPatch) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HiveMySqlCredentialsPatch) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HiveMySqlCredentialsPatch) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *HiveMySqlCredentialsPatch) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *HiveMySqlCredentialsPatch) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSslCaUrl

`func (o *HiveMySqlCredentialsPatch) GetSslCaUrl() string`

GetSslCaUrl returns the SslCaUrl field if non-nil, zero value otherwise.

### GetSslCaUrlOk

`func (o *HiveMySqlCredentialsPatch) GetSslCaUrlOk() (*string, bool)`

GetSslCaUrlOk returns a tuple with the SslCaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaUrl

`func (o *HiveMySqlCredentialsPatch) SetSslCaUrl(v string)`

SetSslCaUrl sets SslCaUrl field to given value.

### HasSslCaUrl

`func (o *HiveMySqlCredentialsPatch) HasSslCaUrl() bool`

HasSslCaUrl returns a boolean if a field has been set.

### SetSslCaUrlNil

`func (o *HiveMySqlCredentialsPatch) SetSslCaUrlNil(b bool)`

 SetSslCaUrlNil sets the value for SslCaUrl to be an explicit nil

### UnsetSslCaUrl
`func (o *HiveMySqlCredentialsPatch) UnsetSslCaUrl()`

UnsetSslCaUrl ensures that no value is present for SslCaUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


