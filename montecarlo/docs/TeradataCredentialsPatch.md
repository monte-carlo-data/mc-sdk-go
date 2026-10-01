# TeradataCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | Hostname of the database endpoint. | [optional] 
**Port** | Pointer to **NullableInt32** | Port the database listens on. | [optional] 
**DbName** | Pointer to **NullableString** | Database to connect to. | [optional] 
**User** | Pointer to **NullableString** | Database user Monte Carlo logs in as. | [optional] 
**Password** | Pointer to **NullableString** | New password of the database user. | [optional] 
**SslCaData** | Pointer to **NullableString** | PEM text of the CA certificate the server&#39;s certificate is checked against. | [optional] 
**SslDisabled** | Pointer to **NullableBool** | Do not check the server against &#x60;ssl_ca_data&#x60;. &#x60;td_sslmode&#x60; decides whether the connection uses TLS. | [optional] 
**TdSslmode** | Pointer to [**NullableTeradataSslMode**](TeradataSslMode.md) | How the connection to Teradata uses TLS. | [optional] 
**TdLogmech** | Pointer to [**NullableTeradataLogonMechanism**](TeradataLogonMechanism.md) | How Teradata authenticates the user. | [optional] 

## Methods

### NewTeradataCredentialsPatch

`func NewTeradataCredentialsPatch() *TeradataCredentialsPatch`

NewTeradataCredentialsPatch instantiates a new TeradataCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeradataCredentialsPatchWithDefaults

`func NewTeradataCredentialsPatchWithDefaults() *TeradataCredentialsPatch`

NewTeradataCredentialsPatchWithDefaults instantiates a new TeradataCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *TeradataCredentialsPatch) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TeradataCredentialsPatch) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TeradataCredentialsPatch) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TeradataCredentialsPatch) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *TeradataCredentialsPatch) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *TeradataCredentialsPatch) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *TeradataCredentialsPatch) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TeradataCredentialsPatch) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TeradataCredentialsPatch) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *TeradataCredentialsPatch) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *TeradataCredentialsPatch) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *TeradataCredentialsPatch) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *TeradataCredentialsPatch) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *TeradataCredentialsPatch) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *TeradataCredentialsPatch) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *TeradataCredentialsPatch) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *TeradataCredentialsPatch) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *TeradataCredentialsPatch) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetUser

`func (o *TeradataCredentialsPatch) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TeradataCredentialsPatch) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TeradataCredentialsPatch) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *TeradataCredentialsPatch) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *TeradataCredentialsPatch) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *TeradataCredentialsPatch) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *TeradataCredentialsPatch) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TeradataCredentialsPatch) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TeradataCredentialsPatch) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TeradataCredentialsPatch) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TeradataCredentialsPatch) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TeradataCredentialsPatch) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSslCaData

`func (o *TeradataCredentialsPatch) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *TeradataCredentialsPatch) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *TeradataCredentialsPatch) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.

### HasSslCaData

`func (o *TeradataCredentialsPatch) HasSslCaData() bool`

HasSslCaData returns a boolean if a field has been set.

### SetSslCaDataNil

`func (o *TeradataCredentialsPatch) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *TeradataCredentialsPatch) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslDisabled

`func (o *TeradataCredentialsPatch) GetSslDisabled() bool`

GetSslDisabled returns the SslDisabled field if non-nil, zero value otherwise.

### GetSslDisabledOk

`func (o *TeradataCredentialsPatch) GetSslDisabledOk() (*bool, bool)`

GetSslDisabledOk returns a tuple with the SslDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDisabled

`func (o *TeradataCredentialsPatch) SetSslDisabled(v bool)`

SetSslDisabled sets SslDisabled field to given value.

### HasSslDisabled

`func (o *TeradataCredentialsPatch) HasSslDisabled() bool`

HasSslDisabled returns a boolean if a field has been set.

### SetSslDisabledNil

`func (o *TeradataCredentialsPatch) SetSslDisabledNil(b bool)`

 SetSslDisabledNil sets the value for SslDisabled to be an explicit nil

### UnsetSslDisabled
`func (o *TeradataCredentialsPatch) UnsetSslDisabled()`

UnsetSslDisabled ensures that no value is present for SslDisabled, not even an explicit nil
### GetTdSslmode

`func (o *TeradataCredentialsPatch) GetTdSslmode() TeradataSslMode`

GetTdSslmode returns the TdSslmode field if non-nil, zero value otherwise.

### GetTdSslmodeOk

`func (o *TeradataCredentialsPatch) GetTdSslmodeOk() (*TeradataSslMode, bool)`

GetTdSslmodeOk returns a tuple with the TdSslmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdSslmode

`func (o *TeradataCredentialsPatch) SetTdSslmode(v TeradataSslMode)`

SetTdSslmode sets TdSslmode field to given value.

### HasTdSslmode

`func (o *TeradataCredentialsPatch) HasTdSslmode() bool`

HasTdSslmode returns a boolean if a field has been set.

### SetTdSslmodeNil

`func (o *TeradataCredentialsPatch) SetTdSslmodeNil(b bool)`

 SetTdSslmodeNil sets the value for TdSslmode to be an explicit nil

### UnsetTdSslmode
`func (o *TeradataCredentialsPatch) UnsetTdSslmode()`

UnsetTdSslmode ensures that no value is present for TdSslmode, not even an explicit nil
### GetTdLogmech

`func (o *TeradataCredentialsPatch) GetTdLogmech() TeradataLogonMechanism`

GetTdLogmech returns the TdLogmech field if non-nil, zero value otherwise.

### GetTdLogmechOk

`func (o *TeradataCredentialsPatch) GetTdLogmechOk() (*TeradataLogonMechanism, bool)`

GetTdLogmechOk returns a tuple with the TdLogmech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdLogmech

`func (o *TeradataCredentialsPatch) SetTdLogmech(v TeradataLogonMechanism)`

SetTdLogmech sets TdLogmech field to given value.

### HasTdLogmech

`func (o *TeradataCredentialsPatch) HasTdLogmech() bool`

HasTdLogmech returns a boolean if a field has been set.

### SetTdLogmechNil

`func (o *TeradataCredentialsPatch) SetTdLogmechNil(b bool)`

 SetTdLogmechNil sets the value for TdLogmech to be an explicit nil

### UnsetTdLogmech
`func (o *TeradataCredentialsPatch) UnsetTdLogmech()`

UnsetTdLogmech ensures that no value is present for TdLogmech, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


