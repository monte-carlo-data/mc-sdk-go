# PostgresCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname of the database endpoint. | 
**Port** | **int32** | Port the database listens on. | 
**User** | **string** | Database user Monte Carlo logs in as. | 
**Password** | **string** | Password of the database user. Stored by Monte Carlo and never returned. | 
**DbName** | **string** | Database to connect to. | 
**SslCaData** | Pointer to **NullableString** | PEM text of the CA certificate the server&#39;s certificate is checked against. | [optional] 
**SslDisabled** | Pointer to **NullableBool** | Connect without TLS. | [optional] 
**SslVerifyCert** | Pointer to **NullableBool** | Check the server&#39;s certificate against the CA. | [optional] 
**SslVerifyIdentity** | Pointer to **NullableBool** | Check the certificate and that it names the host. | [optional] 
**SslSkipCertVerification** | Pointer to **NullableBool** | Encrypt the connection without checking the server&#39;s certificate. | [optional] 
**RdsProxy** | Pointer to **NullableBool** | Whether &#x60;host&#x60; is an Amazon RDS Proxy endpoint. | [optional] 

## Methods

### NewPostgresCredentialsIn

`func NewPostgresCredentialsIn(host string, port int32, user string, password string, dbName string, ) *PostgresCredentialsIn`

NewPostgresCredentialsIn instantiates a new PostgresCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresCredentialsInWithDefaults

`func NewPostgresCredentialsInWithDefaults() *PostgresCredentialsIn`

NewPostgresCredentialsInWithDefaults instantiates a new PostgresCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *PostgresCredentialsIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PostgresCredentialsIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PostgresCredentialsIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *PostgresCredentialsIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostgresCredentialsIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostgresCredentialsIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *PostgresCredentialsIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PostgresCredentialsIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PostgresCredentialsIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *PostgresCredentialsIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostgresCredentialsIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostgresCredentialsIn) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDbName

`func (o *PostgresCredentialsIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *PostgresCredentialsIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *PostgresCredentialsIn) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetSslCaData

`func (o *PostgresCredentialsIn) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *PostgresCredentialsIn) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *PostgresCredentialsIn) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.

### HasSslCaData

`func (o *PostgresCredentialsIn) HasSslCaData() bool`

HasSslCaData returns a boolean if a field has been set.

### SetSslCaDataNil

`func (o *PostgresCredentialsIn) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *PostgresCredentialsIn) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslDisabled

`func (o *PostgresCredentialsIn) GetSslDisabled() bool`

GetSslDisabled returns the SslDisabled field if non-nil, zero value otherwise.

### GetSslDisabledOk

`func (o *PostgresCredentialsIn) GetSslDisabledOk() (*bool, bool)`

GetSslDisabledOk returns a tuple with the SslDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDisabled

`func (o *PostgresCredentialsIn) SetSslDisabled(v bool)`

SetSslDisabled sets SslDisabled field to given value.

### HasSslDisabled

`func (o *PostgresCredentialsIn) HasSslDisabled() bool`

HasSslDisabled returns a boolean if a field has been set.

### SetSslDisabledNil

`func (o *PostgresCredentialsIn) SetSslDisabledNil(b bool)`

 SetSslDisabledNil sets the value for SslDisabled to be an explicit nil

### UnsetSslDisabled
`func (o *PostgresCredentialsIn) UnsetSslDisabled()`

UnsetSslDisabled ensures that no value is present for SslDisabled, not even an explicit nil
### GetSslVerifyCert

`func (o *PostgresCredentialsIn) GetSslVerifyCert() bool`

GetSslVerifyCert returns the SslVerifyCert field if non-nil, zero value otherwise.

### GetSslVerifyCertOk

`func (o *PostgresCredentialsIn) GetSslVerifyCertOk() (*bool, bool)`

GetSslVerifyCertOk returns a tuple with the SslVerifyCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerifyCert

`func (o *PostgresCredentialsIn) SetSslVerifyCert(v bool)`

SetSslVerifyCert sets SslVerifyCert field to given value.

### HasSslVerifyCert

`func (o *PostgresCredentialsIn) HasSslVerifyCert() bool`

HasSslVerifyCert returns a boolean if a field has been set.

### SetSslVerifyCertNil

`func (o *PostgresCredentialsIn) SetSslVerifyCertNil(b bool)`

 SetSslVerifyCertNil sets the value for SslVerifyCert to be an explicit nil

### UnsetSslVerifyCert
`func (o *PostgresCredentialsIn) UnsetSslVerifyCert()`

UnsetSslVerifyCert ensures that no value is present for SslVerifyCert, not even an explicit nil
### GetSslVerifyIdentity

`func (o *PostgresCredentialsIn) GetSslVerifyIdentity() bool`

GetSslVerifyIdentity returns the SslVerifyIdentity field if non-nil, zero value otherwise.

### GetSslVerifyIdentityOk

`func (o *PostgresCredentialsIn) GetSslVerifyIdentityOk() (*bool, bool)`

GetSslVerifyIdentityOk returns a tuple with the SslVerifyIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerifyIdentity

`func (o *PostgresCredentialsIn) SetSslVerifyIdentity(v bool)`

SetSslVerifyIdentity sets SslVerifyIdentity field to given value.

### HasSslVerifyIdentity

`func (o *PostgresCredentialsIn) HasSslVerifyIdentity() bool`

HasSslVerifyIdentity returns a boolean if a field has been set.

### SetSslVerifyIdentityNil

`func (o *PostgresCredentialsIn) SetSslVerifyIdentityNil(b bool)`

 SetSslVerifyIdentityNil sets the value for SslVerifyIdentity to be an explicit nil

### UnsetSslVerifyIdentity
`func (o *PostgresCredentialsIn) UnsetSslVerifyIdentity()`

UnsetSslVerifyIdentity ensures that no value is present for SslVerifyIdentity, not even an explicit nil
### GetSslSkipCertVerification

`func (o *PostgresCredentialsIn) GetSslSkipCertVerification() bool`

GetSslSkipCertVerification returns the SslSkipCertVerification field if non-nil, zero value otherwise.

### GetSslSkipCertVerificationOk

`func (o *PostgresCredentialsIn) GetSslSkipCertVerificationOk() (*bool, bool)`

GetSslSkipCertVerificationOk returns a tuple with the SslSkipCertVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSkipCertVerification

`func (o *PostgresCredentialsIn) SetSslSkipCertVerification(v bool)`

SetSslSkipCertVerification sets SslSkipCertVerification field to given value.

### HasSslSkipCertVerification

`func (o *PostgresCredentialsIn) HasSslSkipCertVerification() bool`

HasSslSkipCertVerification returns a boolean if a field has been set.

### SetSslSkipCertVerificationNil

`func (o *PostgresCredentialsIn) SetSslSkipCertVerificationNil(b bool)`

 SetSslSkipCertVerificationNil sets the value for SslSkipCertVerification to be an explicit nil

### UnsetSslSkipCertVerification
`func (o *PostgresCredentialsIn) UnsetSslSkipCertVerification()`

UnsetSslSkipCertVerification ensures that no value is present for SslSkipCertVerification, not even an explicit nil
### GetRdsProxy

`func (o *PostgresCredentialsIn) GetRdsProxy() bool`

GetRdsProxy returns the RdsProxy field if non-nil, zero value otherwise.

### GetRdsProxyOk

`func (o *PostgresCredentialsIn) GetRdsProxyOk() (*bool, bool)`

GetRdsProxyOk returns a tuple with the RdsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsProxy

`func (o *PostgresCredentialsIn) SetRdsProxy(v bool)`

SetRdsProxy sets RdsProxy field to given value.

### HasRdsProxy

`func (o *PostgresCredentialsIn) HasRdsProxy() bool`

HasRdsProxy returns a boolean if a field has been set.

### SetRdsProxyNil

`func (o *PostgresCredentialsIn) SetRdsProxyNil(b bool)`

 SetRdsProxyNil sets the value for RdsProxy to be an explicit nil

### UnsetRdsProxy
`func (o *PostgresCredentialsIn) UnsetRdsProxy()`

UnsetRdsProxy ensures that no value is present for RdsProxy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


