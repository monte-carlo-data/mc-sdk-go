# RedshiftCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**Host** | **string** | Hostname of the database endpoint. | 
**Port** | **int32** | Port the database listens on. | 
**DbName** | **string** | Redshift database to connect to. | 
**User** | **string** | Database user Monte Carlo logs in as. | 
**SslCaData** | **NullableString** | PEM text of the CA certificate the server&#39;s certificate is checked against. Null when none is set. | 
**SslDisabled** | **NullableBool** | Connect without TLS. Null when unset. | 
**SslVerifyCert** | **NullableBool** | Check the server&#39;s certificate against the CA. Null when unset. | 
**SslVerifyIdentity** | **NullableBool** | Check the certificate and that it names the host. Null when unset. | 
**SslSkipCertVerification** | **NullableBool** | Encrypt the connection without checking the server&#39;s certificate. Null when unset. | 

## Methods

### NewRedshiftCredentialsOut

`func NewRedshiftCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, host string, port int32, dbName string, user string, sslCaData NullableString, sslDisabled NullableBool, sslVerifyCert NullableBool, sslVerifyIdentity NullableBool, sslSkipCertVerification NullableBool, ) *RedshiftCredentialsOut`

NewRedshiftCredentialsOut instantiates a new RedshiftCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedshiftCredentialsOutWithDefaults

`func NewRedshiftCredentialsOutWithDefaults() *RedshiftCredentialsOut`

NewRedshiftCredentialsOutWithDefaults instantiates a new RedshiftCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RedshiftCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RedshiftCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RedshiftCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *RedshiftCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *RedshiftCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *RedshiftCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *RedshiftCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *RedshiftCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *RedshiftCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *RedshiftCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *RedshiftCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *RedshiftCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetHost

`func (o *RedshiftCredentialsOut) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RedshiftCredentialsOut) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RedshiftCredentialsOut) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *RedshiftCredentialsOut) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RedshiftCredentialsOut) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RedshiftCredentialsOut) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDbName

`func (o *RedshiftCredentialsOut) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *RedshiftCredentialsOut) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *RedshiftCredentialsOut) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetUser

`func (o *RedshiftCredentialsOut) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RedshiftCredentialsOut) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RedshiftCredentialsOut) SetUser(v string)`

SetUser sets User field to given value.


### GetSslCaData

`func (o *RedshiftCredentialsOut) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *RedshiftCredentialsOut) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *RedshiftCredentialsOut) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.


### SetSslCaDataNil

`func (o *RedshiftCredentialsOut) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *RedshiftCredentialsOut) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslDisabled

`func (o *RedshiftCredentialsOut) GetSslDisabled() bool`

GetSslDisabled returns the SslDisabled field if non-nil, zero value otherwise.

### GetSslDisabledOk

`func (o *RedshiftCredentialsOut) GetSslDisabledOk() (*bool, bool)`

GetSslDisabledOk returns a tuple with the SslDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDisabled

`func (o *RedshiftCredentialsOut) SetSslDisabled(v bool)`

SetSslDisabled sets SslDisabled field to given value.


### SetSslDisabledNil

`func (o *RedshiftCredentialsOut) SetSslDisabledNil(b bool)`

 SetSslDisabledNil sets the value for SslDisabled to be an explicit nil

### UnsetSslDisabled
`func (o *RedshiftCredentialsOut) UnsetSslDisabled()`

UnsetSslDisabled ensures that no value is present for SslDisabled, not even an explicit nil
### GetSslVerifyCert

`func (o *RedshiftCredentialsOut) GetSslVerifyCert() bool`

GetSslVerifyCert returns the SslVerifyCert field if non-nil, zero value otherwise.

### GetSslVerifyCertOk

`func (o *RedshiftCredentialsOut) GetSslVerifyCertOk() (*bool, bool)`

GetSslVerifyCertOk returns a tuple with the SslVerifyCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerifyCert

`func (o *RedshiftCredentialsOut) SetSslVerifyCert(v bool)`

SetSslVerifyCert sets SslVerifyCert field to given value.


### SetSslVerifyCertNil

`func (o *RedshiftCredentialsOut) SetSslVerifyCertNil(b bool)`

 SetSslVerifyCertNil sets the value for SslVerifyCert to be an explicit nil

### UnsetSslVerifyCert
`func (o *RedshiftCredentialsOut) UnsetSslVerifyCert()`

UnsetSslVerifyCert ensures that no value is present for SslVerifyCert, not even an explicit nil
### GetSslVerifyIdentity

`func (o *RedshiftCredentialsOut) GetSslVerifyIdentity() bool`

GetSslVerifyIdentity returns the SslVerifyIdentity field if non-nil, zero value otherwise.

### GetSslVerifyIdentityOk

`func (o *RedshiftCredentialsOut) GetSslVerifyIdentityOk() (*bool, bool)`

GetSslVerifyIdentityOk returns a tuple with the SslVerifyIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerifyIdentity

`func (o *RedshiftCredentialsOut) SetSslVerifyIdentity(v bool)`

SetSslVerifyIdentity sets SslVerifyIdentity field to given value.


### SetSslVerifyIdentityNil

`func (o *RedshiftCredentialsOut) SetSslVerifyIdentityNil(b bool)`

 SetSslVerifyIdentityNil sets the value for SslVerifyIdentity to be an explicit nil

### UnsetSslVerifyIdentity
`func (o *RedshiftCredentialsOut) UnsetSslVerifyIdentity()`

UnsetSslVerifyIdentity ensures that no value is present for SslVerifyIdentity, not even an explicit nil
### GetSslSkipCertVerification

`func (o *RedshiftCredentialsOut) GetSslSkipCertVerification() bool`

GetSslSkipCertVerification returns the SslSkipCertVerification field if non-nil, zero value otherwise.

### GetSslSkipCertVerificationOk

`func (o *RedshiftCredentialsOut) GetSslSkipCertVerificationOk() (*bool, bool)`

GetSslSkipCertVerificationOk returns a tuple with the SslSkipCertVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSkipCertVerification

`func (o *RedshiftCredentialsOut) SetSslSkipCertVerification(v bool)`

SetSslSkipCertVerification sets SslSkipCertVerification field to given value.


### SetSslSkipCertVerificationNil

`func (o *RedshiftCredentialsOut) SetSslSkipCertVerificationNil(b bool)`

 SetSslSkipCertVerificationNil sets the value for SslSkipCertVerification to be an explicit nil

### UnsetSslSkipCertVerification
`func (o *RedshiftCredentialsOut) UnsetSslSkipCertVerification()`

UnsetSslSkipCertVerification ensures that no value is present for SslSkipCertVerification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


