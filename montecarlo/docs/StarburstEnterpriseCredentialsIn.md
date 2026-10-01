# StarburstEnterpriseCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname of the database endpoint. | 
**Port** | **int32** | Port the database listens on. | 
**User** | **string** | Database user Monte Carlo logs in as. | 
**Password** | **string** | Password of the database user. Stored by Monte Carlo and never returned. | 
**DbName** | Pointer to **NullableString** | Database to connect to. | [optional] 
**SslCaData** | Pointer to **NullableString** | PEM text of the CA certificate the server&#39;s certificate is checked against. | [optional] 
**SslDisabled** | Pointer to **NullableBool** | Skip the check of the server&#39;s certificate. The connection always uses TLS. | [optional] 

## Methods

### NewStarburstEnterpriseCredentialsIn

`func NewStarburstEnterpriseCredentialsIn(host string, port int32, user string, password string, ) *StarburstEnterpriseCredentialsIn`

NewStarburstEnterpriseCredentialsIn instantiates a new StarburstEnterpriseCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStarburstEnterpriseCredentialsInWithDefaults

`func NewStarburstEnterpriseCredentialsInWithDefaults() *StarburstEnterpriseCredentialsIn`

NewStarburstEnterpriseCredentialsInWithDefaults instantiates a new StarburstEnterpriseCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *StarburstEnterpriseCredentialsIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StarburstEnterpriseCredentialsIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StarburstEnterpriseCredentialsIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *StarburstEnterpriseCredentialsIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *StarburstEnterpriseCredentialsIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *StarburstEnterpriseCredentialsIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *StarburstEnterpriseCredentialsIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *StarburstEnterpriseCredentialsIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *StarburstEnterpriseCredentialsIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *StarburstEnterpriseCredentialsIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *StarburstEnterpriseCredentialsIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *StarburstEnterpriseCredentialsIn) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDbName

`func (o *StarburstEnterpriseCredentialsIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *StarburstEnterpriseCredentialsIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *StarburstEnterpriseCredentialsIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *StarburstEnterpriseCredentialsIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *StarburstEnterpriseCredentialsIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *StarburstEnterpriseCredentialsIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetSslCaData

`func (o *StarburstEnterpriseCredentialsIn) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *StarburstEnterpriseCredentialsIn) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *StarburstEnterpriseCredentialsIn) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.

### HasSslCaData

`func (o *StarburstEnterpriseCredentialsIn) HasSslCaData() bool`

HasSslCaData returns a boolean if a field has been set.

### SetSslCaDataNil

`func (o *StarburstEnterpriseCredentialsIn) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *StarburstEnterpriseCredentialsIn) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslDisabled

`func (o *StarburstEnterpriseCredentialsIn) GetSslDisabled() bool`

GetSslDisabled returns the SslDisabled field if non-nil, zero value otherwise.

### GetSslDisabledOk

`func (o *StarburstEnterpriseCredentialsIn) GetSslDisabledOk() (*bool, bool)`

GetSslDisabledOk returns a tuple with the SslDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDisabled

`func (o *StarburstEnterpriseCredentialsIn) SetSslDisabled(v bool)`

SetSslDisabled sets SslDisabled field to given value.

### HasSslDisabled

`func (o *StarburstEnterpriseCredentialsIn) HasSslDisabled() bool`

HasSslDisabled returns a boolean if a field has been set.

### SetSslDisabledNil

`func (o *StarburstEnterpriseCredentialsIn) SetSslDisabledNil(b bool)`

 SetSslDisabledNil sets the value for SslDisabled to be an explicit nil

### UnsetSslDisabled
`func (o *StarburstEnterpriseCredentialsIn) UnsetSslDisabled()`

UnsetSslDisabled ensures that no value is present for SslDisabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


