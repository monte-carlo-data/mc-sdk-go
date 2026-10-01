# MySqlCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Host** | **string** | Hostname of the database endpoint. | 
**Port** | **int32** | Port the database listens on. | 
**User** | **string** | Database user Monte Carlo logs in as. | 
**Password** | **string** | Password of the database user. Used for this check and not kept. | 
**DbName** | Pointer to **NullableString** | Database to connect to. | [optional] 
**SslCaData** | Pointer to **NullableString** | PEM text of the CA certificate the server&#39;s certificate is checked against. | [optional] 
**SslDisabled** | Pointer to **NullableBool** | Connect without TLS. | [optional] 
**SslVerifyCert** | Pointer to **NullableBool** | Check the server&#39;s certificate against the CA. | [optional] 
**SslVerifyIdentity** | Pointer to **NullableBool** | Check the certificate and that it names the host. | [optional] 
**SslSkipCertVerification** | Pointer to **NullableBool** | Encrypt the connection without checking the server&#39;s certificate. | [optional] 

## Methods

### NewMySqlCredentialsValidateIn

`func NewMySqlCredentialsValidateIn(deploymentId string, host string, port int32, user string, password string, ) *MySqlCredentialsValidateIn`

NewMySqlCredentialsValidateIn instantiates a new MySqlCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMySqlCredentialsValidateInWithDefaults

`func NewMySqlCredentialsValidateInWithDefaults() *MySqlCredentialsValidateIn`

NewMySqlCredentialsValidateInWithDefaults instantiates a new MySqlCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *MySqlCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *MySqlCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *MySqlCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetHost

`func (o *MySqlCredentialsValidateIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MySqlCredentialsValidateIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MySqlCredentialsValidateIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *MySqlCredentialsValidateIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MySqlCredentialsValidateIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MySqlCredentialsValidateIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *MySqlCredentialsValidateIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MySqlCredentialsValidateIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MySqlCredentialsValidateIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *MySqlCredentialsValidateIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MySqlCredentialsValidateIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MySqlCredentialsValidateIn) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDbName

`func (o *MySqlCredentialsValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MySqlCredentialsValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MySqlCredentialsValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *MySqlCredentialsValidateIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *MySqlCredentialsValidateIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *MySqlCredentialsValidateIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetSslCaData

`func (o *MySqlCredentialsValidateIn) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *MySqlCredentialsValidateIn) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *MySqlCredentialsValidateIn) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.

### HasSslCaData

`func (o *MySqlCredentialsValidateIn) HasSslCaData() bool`

HasSslCaData returns a boolean if a field has been set.

### SetSslCaDataNil

`func (o *MySqlCredentialsValidateIn) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *MySqlCredentialsValidateIn) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslDisabled

`func (o *MySqlCredentialsValidateIn) GetSslDisabled() bool`

GetSslDisabled returns the SslDisabled field if non-nil, zero value otherwise.

### GetSslDisabledOk

`func (o *MySqlCredentialsValidateIn) GetSslDisabledOk() (*bool, bool)`

GetSslDisabledOk returns a tuple with the SslDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDisabled

`func (o *MySqlCredentialsValidateIn) SetSslDisabled(v bool)`

SetSslDisabled sets SslDisabled field to given value.

### HasSslDisabled

`func (o *MySqlCredentialsValidateIn) HasSslDisabled() bool`

HasSslDisabled returns a boolean if a field has been set.

### SetSslDisabledNil

`func (o *MySqlCredentialsValidateIn) SetSslDisabledNil(b bool)`

 SetSslDisabledNil sets the value for SslDisabled to be an explicit nil

### UnsetSslDisabled
`func (o *MySqlCredentialsValidateIn) UnsetSslDisabled()`

UnsetSslDisabled ensures that no value is present for SslDisabled, not even an explicit nil
### GetSslVerifyCert

`func (o *MySqlCredentialsValidateIn) GetSslVerifyCert() bool`

GetSslVerifyCert returns the SslVerifyCert field if non-nil, zero value otherwise.

### GetSslVerifyCertOk

`func (o *MySqlCredentialsValidateIn) GetSslVerifyCertOk() (*bool, bool)`

GetSslVerifyCertOk returns a tuple with the SslVerifyCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerifyCert

`func (o *MySqlCredentialsValidateIn) SetSslVerifyCert(v bool)`

SetSslVerifyCert sets SslVerifyCert field to given value.

### HasSslVerifyCert

`func (o *MySqlCredentialsValidateIn) HasSslVerifyCert() bool`

HasSslVerifyCert returns a boolean if a field has been set.

### SetSslVerifyCertNil

`func (o *MySqlCredentialsValidateIn) SetSslVerifyCertNil(b bool)`

 SetSslVerifyCertNil sets the value for SslVerifyCert to be an explicit nil

### UnsetSslVerifyCert
`func (o *MySqlCredentialsValidateIn) UnsetSslVerifyCert()`

UnsetSslVerifyCert ensures that no value is present for SslVerifyCert, not even an explicit nil
### GetSslVerifyIdentity

`func (o *MySqlCredentialsValidateIn) GetSslVerifyIdentity() bool`

GetSslVerifyIdentity returns the SslVerifyIdentity field if non-nil, zero value otherwise.

### GetSslVerifyIdentityOk

`func (o *MySqlCredentialsValidateIn) GetSslVerifyIdentityOk() (*bool, bool)`

GetSslVerifyIdentityOk returns a tuple with the SslVerifyIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerifyIdentity

`func (o *MySqlCredentialsValidateIn) SetSslVerifyIdentity(v bool)`

SetSslVerifyIdentity sets SslVerifyIdentity field to given value.

### HasSslVerifyIdentity

`func (o *MySqlCredentialsValidateIn) HasSslVerifyIdentity() bool`

HasSslVerifyIdentity returns a boolean if a field has been set.

### SetSslVerifyIdentityNil

`func (o *MySqlCredentialsValidateIn) SetSslVerifyIdentityNil(b bool)`

 SetSslVerifyIdentityNil sets the value for SslVerifyIdentity to be an explicit nil

### UnsetSslVerifyIdentity
`func (o *MySqlCredentialsValidateIn) UnsetSslVerifyIdentity()`

UnsetSslVerifyIdentity ensures that no value is present for SslVerifyIdentity, not even an explicit nil
### GetSslSkipCertVerification

`func (o *MySqlCredentialsValidateIn) GetSslSkipCertVerification() bool`

GetSslSkipCertVerification returns the SslSkipCertVerification field if non-nil, zero value otherwise.

### GetSslSkipCertVerificationOk

`func (o *MySqlCredentialsValidateIn) GetSslSkipCertVerificationOk() (*bool, bool)`

GetSslSkipCertVerificationOk returns a tuple with the SslSkipCertVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSkipCertVerification

`func (o *MySqlCredentialsValidateIn) SetSslSkipCertVerification(v bool)`

SetSslSkipCertVerification sets SslSkipCertVerification field to given value.

### HasSslSkipCertVerification

`func (o *MySqlCredentialsValidateIn) HasSslSkipCertVerification() bool`

HasSslSkipCertVerification returns a boolean if a field has been set.

### SetSslSkipCertVerificationNil

`func (o *MySqlCredentialsValidateIn) SetSslSkipCertVerificationNil(b bool)`

 SetSslSkipCertVerificationNil sets the value for SslSkipCertVerification to be an explicit nil

### UnsetSslSkipCertVerification
`func (o *MySqlCredentialsValidateIn) UnsetSslSkipCertVerification()`

UnsetSslSkipCertVerification ensures that no value is present for SslSkipCertVerification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


