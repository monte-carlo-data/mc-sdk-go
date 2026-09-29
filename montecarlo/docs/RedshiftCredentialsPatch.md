# RedshiftCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | Hostname of the database endpoint. | [optional] 
**Port** | Pointer to **NullableInt32** | Port the database listens on. | [optional] 
**DbName** | Pointer to **NullableString** | Database to connect to. | [optional] 
**User** | Pointer to **NullableString** | Database user Monte Carlo logs in as. | [optional] 
**Password** | Pointer to **NullableString** | New password of the database user. | [optional] 
**SslCaData** | Pointer to **NullableString** | PEM text of the CA certificate the server&#39;s certificate is checked against. | [optional] 
**SslDisabled** | Pointer to **NullableBool** | Connect without TLS. | [optional] 
**SslVerifyCert** | Pointer to **NullableBool** | Check the server&#39;s certificate against the CA. | [optional] 
**SslVerifyIdentity** | Pointer to **NullableBool** | Check the certificate and that it names the host. | [optional] 
**SslSkipCertVerification** | Pointer to **NullableBool** | Encrypt the connection without checking the server&#39;s certificate. | [optional] 

## Methods

### NewRedshiftCredentialsPatch

`func NewRedshiftCredentialsPatch() *RedshiftCredentialsPatch`

NewRedshiftCredentialsPatch instantiates a new RedshiftCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedshiftCredentialsPatchWithDefaults

`func NewRedshiftCredentialsPatchWithDefaults() *RedshiftCredentialsPatch`

NewRedshiftCredentialsPatchWithDefaults instantiates a new RedshiftCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RedshiftCredentialsPatch) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RedshiftCredentialsPatch) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RedshiftCredentialsPatch) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RedshiftCredentialsPatch) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *RedshiftCredentialsPatch) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *RedshiftCredentialsPatch) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *RedshiftCredentialsPatch) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RedshiftCredentialsPatch) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RedshiftCredentialsPatch) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RedshiftCredentialsPatch) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *RedshiftCredentialsPatch) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *RedshiftCredentialsPatch) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *RedshiftCredentialsPatch) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *RedshiftCredentialsPatch) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *RedshiftCredentialsPatch) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *RedshiftCredentialsPatch) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *RedshiftCredentialsPatch) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *RedshiftCredentialsPatch) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetUser

`func (o *RedshiftCredentialsPatch) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RedshiftCredentialsPatch) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RedshiftCredentialsPatch) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RedshiftCredentialsPatch) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *RedshiftCredentialsPatch) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *RedshiftCredentialsPatch) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *RedshiftCredentialsPatch) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RedshiftCredentialsPatch) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RedshiftCredentialsPatch) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RedshiftCredentialsPatch) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RedshiftCredentialsPatch) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RedshiftCredentialsPatch) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSslCaData

`func (o *RedshiftCredentialsPatch) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *RedshiftCredentialsPatch) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *RedshiftCredentialsPatch) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.

### HasSslCaData

`func (o *RedshiftCredentialsPatch) HasSslCaData() bool`

HasSslCaData returns a boolean if a field has been set.

### SetSslCaDataNil

`func (o *RedshiftCredentialsPatch) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *RedshiftCredentialsPatch) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslDisabled

`func (o *RedshiftCredentialsPatch) GetSslDisabled() bool`

GetSslDisabled returns the SslDisabled field if non-nil, zero value otherwise.

### GetSslDisabledOk

`func (o *RedshiftCredentialsPatch) GetSslDisabledOk() (*bool, bool)`

GetSslDisabledOk returns a tuple with the SslDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDisabled

`func (o *RedshiftCredentialsPatch) SetSslDisabled(v bool)`

SetSslDisabled sets SslDisabled field to given value.

### HasSslDisabled

`func (o *RedshiftCredentialsPatch) HasSslDisabled() bool`

HasSslDisabled returns a boolean if a field has been set.

### SetSslDisabledNil

`func (o *RedshiftCredentialsPatch) SetSslDisabledNil(b bool)`

 SetSslDisabledNil sets the value for SslDisabled to be an explicit nil

### UnsetSslDisabled
`func (o *RedshiftCredentialsPatch) UnsetSslDisabled()`

UnsetSslDisabled ensures that no value is present for SslDisabled, not even an explicit nil
### GetSslVerifyCert

`func (o *RedshiftCredentialsPatch) GetSslVerifyCert() bool`

GetSslVerifyCert returns the SslVerifyCert field if non-nil, zero value otherwise.

### GetSslVerifyCertOk

`func (o *RedshiftCredentialsPatch) GetSslVerifyCertOk() (*bool, bool)`

GetSslVerifyCertOk returns a tuple with the SslVerifyCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerifyCert

`func (o *RedshiftCredentialsPatch) SetSslVerifyCert(v bool)`

SetSslVerifyCert sets SslVerifyCert field to given value.

### HasSslVerifyCert

`func (o *RedshiftCredentialsPatch) HasSslVerifyCert() bool`

HasSslVerifyCert returns a boolean if a field has been set.

### SetSslVerifyCertNil

`func (o *RedshiftCredentialsPatch) SetSslVerifyCertNil(b bool)`

 SetSslVerifyCertNil sets the value for SslVerifyCert to be an explicit nil

### UnsetSslVerifyCert
`func (o *RedshiftCredentialsPatch) UnsetSslVerifyCert()`

UnsetSslVerifyCert ensures that no value is present for SslVerifyCert, not even an explicit nil
### GetSslVerifyIdentity

`func (o *RedshiftCredentialsPatch) GetSslVerifyIdentity() bool`

GetSslVerifyIdentity returns the SslVerifyIdentity field if non-nil, zero value otherwise.

### GetSslVerifyIdentityOk

`func (o *RedshiftCredentialsPatch) GetSslVerifyIdentityOk() (*bool, bool)`

GetSslVerifyIdentityOk returns a tuple with the SslVerifyIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerifyIdentity

`func (o *RedshiftCredentialsPatch) SetSslVerifyIdentity(v bool)`

SetSslVerifyIdentity sets SslVerifyIdentity field to given value.

### HasSslVerifyIdentity

`func (o *RedshiftCredentialsPatch) HasSslVerifyIdentity() bool`

HasSslVerifyIdentity returns a boolean if a field has been set.

### SetSslVerifyIdentityNil

`func (o *RedshiftCredentialsPatch) SetSslVerifyIdentityNil(b bool)`

 SetSslVerifyIdentityNil sets the value for SslVerifyIdentity to be an explicit nil

### UnsetSslVerifyIdentity
`func (o *RedshiftCredentialsPatch) UnsetSslVerifyIdentity()`

UnsetSslVerifyIdentity ensures that no value is present for SslVerifyIdentity, not even an explicit nil
### GetSslSkipCertVerification

`func (o *RedshiftCredentialsPatch) GetSslSkipCertVerification() bool`

GetSslSkipCertVerification returns the SslSkipCertVerification field if non-nil, zero value otherwise.

### GetSslSkipCertVerificationOk

`func (o *RedshiftCredentialsPatch) GetSslSkipCertVerificationOk() (*bool, bool)`

GetSslSkipCertVerificationOk returns a tuple with the SslSkipCertVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSkipCertVerification

`func (o *RedshiftCredentialsPatch) SetSslSkipCertVerification(v bool)`

SetSslSkipCertVerification sets SslSkipCertVerification field to given value.

### HasSslSkipCertVerification

`func (o *RedshiftCredentialsPatch) HasSslSkipCertVerification() bool`

HasSslSkipCertVerification returns a boolean if a field has been set.

### SetSslSkipCertVerificationNil

`func (o *RedshiftCredentialsPatch) SetSslSkipCertVerificationNil(b bool)`

 SetSslSkipCertVerificationNil sets the value for SslSkipCertVerification to be an explicit nil

### UnsetSslSkipCertVerification
`func (o *RedshiftCredentialsPatch) UnsetSslSkipCertVerification()`

UnsetSslSkipCertVerification ensures that no value is present for SslSkipCertVerification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


