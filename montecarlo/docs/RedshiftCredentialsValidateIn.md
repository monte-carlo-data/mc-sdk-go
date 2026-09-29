# RedshiftCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Host** | **string** | Hostname of the database endpoint. | 
**Port** | **int32** | Port the database listens on. | 
**User** | **string** | Database user Monte Carlo logs in as. | 
**Password** | **string** | Password of the database user. Used for this check and not kept. | 
**SslCaData** | Pointer to **NullableString** | PEM text of the CA certificate the server&#39;s certificate is checked against. | [optional] 
**SslDisabled** | Pointer to **NullableBool** | Connect without TLS. | [optional] 
**SslVerifyCert** | Pointer to **NullableBool** | Check the server&#39;s certificate against the CA. | [optional] 
**SslVerifyIdentity** | Pointer to **NullableBool** | Check the certificate and that it names the host. | [optional] 
**SslSkipCertVerification** | Pointer to **NullableBool** | Encrypt the connection without checking the server&#39;s certificate. | [optional] 
**DbName** | **string** | Redshift database to connect to. | 

## Methods

### NewRedshiftCredentialsValidateIn

`func NewRedshiftCredentialsValidateIn(deploymentId string, host string, port int32, user string, password string, dbName string, ) *RedshiftCredentialsValidateIn`

NewRedshiftCredentialsValidateIn instantiates a new RedshiftCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedshiftCredentialsValidateInWithDefaults

`func NewRedshiftCredentialsValidateInWithDefaults() *RedshiftCredentialsValidateIn`

NewRedshiftCredentialsValidateInWithDefaults instantiates a new RedshiftCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *RedshiftCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *RedshiftCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *RedshiftCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetHost

`func (o *RedshiftCredentialsValidateIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RedshiftCredentialsValidateIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RedshiftCredentialsValidateIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *RedshiftCredentialsValidateIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RedshiftCredentialsValidateIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RedshiftCredentialsValidateIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *RedshiftCredentialsValidateIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RedshiftCredentialsValidateIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RedshiftCredentialsValidateIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *RedshiftCredentialsValidateIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RedshiftCredentialsValidateIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RedshiftCredentialsValidateIn) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSslCaData

`func (o *RedshiftCredentialsValidateIn) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *RedshiftCredentialsValidateIn) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *RedshiftCredentialsValidateIn) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.

### HasSslCaData

`func (o *RedshiftCredentialsValidateIn) HasSslCaData() bool`

HasSslCaData returns a boolean if a field has been set.

### SetSslCaDataNil

`func (o *RedshiftCredentialsValidateIn) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *RedshiftCredentialsValidateIn) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslDisabled

`func (o *RedshiftCredentialsValidateIn) GetSslDisabled() bool`

GetSslDisabled returns the SslDisabled field if non-nil, zero value otherwise.

### GetSslDisabledOk

`func (o *RedshiftCredentialsValidateIn) GetSslDisabledOk() (*bool, bool)`

GetSslDisabledOk returns a tuple with the SslDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDisabled

`func (o *RedshiftCredentialsValidateIn) SetSslDisabled(v bool)`

SetSslDisabled sets SslDisabled field to given value.

### HasSslDisabled

`func (o *RedshiftCredentialsValidateIn) HasSslDisabled() bool`

HasSslDisabled returns a boolean if a field has been set.

### SetSslDisabledNil

`func (o *RedshiftCredentialsValidateIn) SetSslDisabledNil(b bool)`

 SetSslDisabledNil sets the value for SslDisabled to be an explicit nil

### UnsetSslDisabled
`func (o *RedshiftCredentialsValidateIn) UnsetSslDisabled()`

UnsetSslDisabled ensures that no value is present for SslDisabled, not even an explicit nil
### GetSslVerifyCert

`func (o *RedshiftCredentialsValidateIn) GetSslVerifyCert() bool`

GetSslVerifyCert returns the SslVerifyCert field if non-nil, zero value otherwise.

### GetSslVerifyCertOk

`func (o *RedshiftCredentialsValidateIn) GetSslVerifyCertOk() (*bool, bool)`

GetSslVerifyCertOk returns a tuple with the SslVerifyCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerifyCert

`func (o *RedshiftCredentialsValidateIn) SetSslVerifyCert(v bool)`

SetSslVerifyCert sets SslVerifyCert field to given value.

### HasSslVerifyCert

`func (o *RedshiftCredentialsValidateIn) HasSslVerifyCert() bool`

HasSslVerifyCert returns a boolean if a field has been set.

### SetSslVerifyCertNil

`func (o *RedshiftCredentialsValidateIn) SetSslVerifyCertNil(b bool)`

 SetSslVerifyCertNil sets the value for SslVerifyCert to be an explicit nil

### UnsetSslVerifyCert
`func (o *RedshiftCredentialsValidateIn) UnsetSslVerifyCert()`

UnsetSslVerifyCert ensures that no value is present for SslVerifyCert, not even an explicit nil
### GetSslVerifyIdentity

`func (o *RedshiftCredentialsValidateIn) GetSslVerifyIdentity() bool`

GetSslVerifyIdentity returns the SslVerifyIdentity field if non-nil, zero value otherwise.

### GetSslVerifyIdentityOk

`func (o *RedshiftCredentialsValidateIn) GetSslVerifyIdentityOk() (*bool, bool)`

GetSslVerifyIdentityOk returns a tuple with the SslVerifyIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerifyIdentity

`func (o *RedshiftCredentialsValidateIn) SetSslVerifyIdentity(v bool)`

SetSslVerifyIdentity sets SslVerifyIdentity field to given value.

### HasSslVerifyIdentity

`func (o *RedshiftCredentialsValidateIn) HasSslVerifyIdentity() bool`

HasSslVerifyIdentity returns a boolean if a field has been set.

### SetSslVerifyIdentityNil

`func (o *RedshiftCredentialsValidateIn) SetSslVerifyIdentityNil(b bool)`

 SetSslVerifyIdentityNil sets the value for SslVerifyIdentity to be an explicit nil

### UnsetSslVerifyIdentity
`func (o *RedshiftCredentialsValidateIn) UnsetSslVerifyIdentity()`

UnsetSslVerifyIdentity ensures that no value is present for SslVerifyIdentity, not even an explicit nil
### GetSslSkipCertVerification

`func (o *RedshiftCredentialsValidateIn) GetSslSkipCertVerification() bool`

GetSslSkipCertVerification returns the SslSkipCertVerification field if non-nil, zero value otherwise.

### GetSslSkipCertVerificationOk

`func (o *RedshiftCredentialsValidateIn) GetSslSkipCertVerificationOk() (*bool, bool)`

GetSslSkipCertVerificationOk returns a tuple with the SslSkipCertVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSkipCertVerification

`func (o *RedshiftCredentialsValidateIn) SetSslSkipCertVerification(v bool)`

SetSslSkipCertVerification sets SslSkipCertVerification field to given value.

### HasSslSkipCertVerification

`func (o *RedshiftCredentialsValidateIn) HasSslSkipCertVerification() bool`

HasSslSkipCertVerification returns a boolean if a field has been set.

### SetSslSkipCertVerificationNil

`func (o *RedshiftCredentialsValidateIn) SetSslSkipCertVerificationNil(b bool)`

 SetSslSkipCertVerificationNil sets the value for SslSkipCertVerification to be an explicit nil

### UnsetSslSkipCertVerification
`func (o *RedshiftCredentialsValidateIn) UnsetSslSkipCertVerification()`

UnsetSslSkipCertVerification ensures that no value is present for SslSkipCertVerification, not even an explicit nil
### GetDbName

`func (o *RedshiftCredentialsValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *RedshiftCredentialsValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *RedshiftCredentialsValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


