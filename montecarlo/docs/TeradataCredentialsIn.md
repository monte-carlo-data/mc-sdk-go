# TeradataCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname of the database endpoint. | 
**Port** | **int32** | Port the database listens on. | 
**User** | **string** | Database user Monte Carlo logs in as. | 
**Password** | **string** | Password of the database user. Stored by Monte Carlo and never returned. | 
**DbName** | Pointer to **NullableString** | Database to connect to. | [optional] 
**SslCaData** | Pointer to **NullableString** | PEM text of the CA certificate the server&#39;s certificate is checked against. | [optional] 
**SslDisabled** | Pointer to **NullableBool** | Do not check the server against &#x60;ssl_ca_data&#x60;. &#x60;td_sslmode&#x60; decides whether the connection uses TLS. | [optional] 
**TdSslmode** | Pointer to [**NullableTeradataSslMode**](TeradataSslMode.md) | How the connection to Teradata uses TLS. | [optional] 
**TdLogmech** | Pointer to [**NullableTeradataLogonMechanism**](TeradataLogonMechanism.md) | How Teradata authenticates the user. | [optional] 

## Methods

### NewTeradataCredentialsIn

`func NewTeradataCredentialsIn(host string, port int32, user string, password string, ) *TeradataCredentialsIn`

NewTeradataCredentialsIn instantiates a new TeradataCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeradataCredentialsInWithDefaults

`func NewTeradataCredentialsInWithDefaults() *TeradataCredentialsIn`

NewTeradataCredentialsInWithDefaults instantiates a new TeradataCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *TeradataCredentialsIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TeradataCredentialsIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TeradataCredentialsIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *TeradataCredentialsIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TeradataCredentialsIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TeradataCredentialsIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *TeradataCredentialsIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TeradataCredentialsIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TeradataCredentialsIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *TeradataCredentialsIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TeradataCredentialsIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TeradataCredentialsIn) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDbName

`func (o *TeradataCredentialsIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *TeradataCredentialsIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *TeradataCredentialsIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *TeradataCredentialsIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *TeradataCredentialsIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *TeradataCredentialsIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetSslCaData

`func (o *TeradataCredentialsIn) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *TeradataCredentialsIn) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *TeradataCredentialsIn) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.

### HasSslCaData

`func (o *TeradataCredentialsIn) HasSslCaData() bool`

HasSslCaData returns a boolean if a field has been set.

### SetSslCaDataNil

`func (o *TeradataCredentialsIn) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *TeradataCredentialsIn) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslDisabled

`func (o *TeradataCredentialsIn) GetSslDisabled() bool`

GetSslDisabled returns the SslDisabled field if non-nil, zero value otherwise.

### GetSslDisabledOk

`func (o *TeradataCredentialsIn) GetSslDisabledOk() (*bool, bool)`

GetSslDisabledOk returns a tuple with the SslDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDisabled

`func (o *TeradataCredentialsIn) SetSslDisabled(v bool)`

SetSslDisabled sets SslDisabled field to given value.

### HasSslDisabled

`func (o *TeradataCredentialsIn) HasSslDisabled() bool`

HasSslDisabled returns a boolean if a field has been set.

### SetSslDisabledNil

`func (o *TeradataCredentialsIn) SetSslDisabledNil(b bool)`

 SetSslDisabledNil sets the value for SslDisabled to be an explicit nil

### UnsetSslDisabled
`func (o *TeradataCredentialsIn) UnsetSslDisabled()`

UnsetSslDisabled ensures that no value is present for SslDisabled, not even an explicit nil
### GetTdSslmode

`func (o *TeradataCredentialsIn) GetTdSslmode() TeradataSslMode`

GetTdSslmode returns the TdSslmode field if non-nil, zero value otherwise.

### GetTdSslmodeOk

`func (o *TeradataCredentialsIn) GetTdSslmodeOk() (*TeradataSslMode, bool)`

GetTdSslmodeOk returns a tuple with the TdSslmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdSslmode

`func (o *TeradataCredentialsIn) SetTdSslmode(v TeradataSslMode)`

SetTdSslmode sets TdSslmode field to given value.

### HasTdSslmode

`func (o *TeradataCredentialsIn) HasTdSslmode() bool`

HasTdSslmode returns a boolean if a field has been set.

### SetTdSslmodeNil

`func (o *TeradataCredentialsIn) SetTdSslmodeNil(b bool)`

 SetTdSslmodeNil sets the value for TdSslmode to be an explicit nil

### UnsetTdSslmode
`func (o *TeradataCredentialsIn) UnsetTdSslmode()`

UnsetTdSslmode ensures that no value is present for TdSslmode, not even an explicit nil
### GetTdLogmech

`func (o *TeradataCredentialsIn) GetTdLogmech() TeradataLogonMechanism`

GetTdLogmech returns the TdLogmech field if non-nil, zero value otherwise.

### GetTdLogmechOk

`func (o *TeradataCredentialsIn) GetTdLogmechOk() (*TeradataLogonMechanism, bool)`

GetTdLogmechOk returns a tuple with the TdLogmech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdLogmech

`func (o *TeradataCredentialsIn) SetTdLogmech(v TeradataLogonMechanism)`

SetTdLogmech sets TdLogmech field to given value.

### HasTdLogmech

`func (o *TeradataCredentialsIn) HasTdLogmech() bool`

HasTdLogmech returns a boolean if a field has been set.

### SetTdLogmechNil

`func (o *TeradataCredentialsIn) SetTdLogmechNil(b bool)`

 SetTdLogmechNil sets the value for TdLogmech to be an explicit nil

### UnsetTdLogmech
`func (o *TeradataCredentialsIn) UnsetTdLogmech()`

UnsetTdLogmech ensures that no value is present for TdLogmech, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


