# LookerGitCloneCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**RepoUrl** | **string** | Clone URL of the LookML repository, over HTTPS or SSH. | 
**Username** | **NullableString** | Git user. Null for an SSH clone. | 
**SslCaData** | **NullableString** | PEM certificate of the CA that signed the git server&#39;s certificate. Null unless set. | 
**SslSkipCertVerification** | **NullableBool** | Skip verifying the git server&#39;s TLS certificate. Null unless set. | 

## Methods

### NewLookerGitCloneCredentialsOut

`func NewLookerGitCloneCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, repoUrl string, username NullableString, sslCaData NullableString, sslSkipCertVerification NullableBool, ) *LookerGitCloneCredentialsOut`

NewLookerGitCloneCredentialsOut instantiates a new LookerGitCloneCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookerGitCloneCredentialsOutWithDefaults

`func NewLookerGitCloneCredentialsOutWithDefaults() *LookerGitCloneCredentialsOut`

NewLookerGitCloneCredentialsOutWithDefaults instantiates a new LookerGitCloneCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LookerGitCloneCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LookerGitCloneCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LookerGitCloneCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *LookerGitCloneCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *LookerGitCloneCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *LookerGitCloneCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *LookerGitCloneCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *LookerGitCloneCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *LookerGitCloneCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *LookerGitCloneCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *LookerGitCloneCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *LookerGitCloneCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetRepoUrl

`func (o *LookerGitCloneCredentialsOut) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *LookerGitCloneCredentialsOut) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *LookerGitCloneCredentialsOut) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.


### GetUsername

`func (o *LookerGitCloneCredentialsOut) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LookerGitCloneCredentialsOut) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LookerGitCloneCredentialsOut) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *LookerGitCloneCredentialsOut) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *LookerGitCloneCredentialsOut) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetSslCaData

`func (o *LookerGitCloneCredentialsOut) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *LookerGitCloneCredentialsOut) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *LookerGitCloneCredentialsOut) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.


### SetSslCaDataNil

`func (o *LookerGitCloneCredentialsOut) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *LookerGitCloneCredentialsOut) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslSkipCertVerification

`func (o *LookerGitCloneCredentialsOut) GetSslSkipCertVerification() bool`

GetSslSkipCertVerification returns the SslSkipCertVerification field if non-nil, zero value otherwise.

### GetSslSkipCertVerificationOk

`func (o *LookerGitCloneCredentialsOut) GetSslSkipCertVerificationOk() (*bool, bool)`

GetSslSkipCertVerificationOk returns a tuple with the SslSkipCertVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSkipCertVerification

`func (o *LookerGitCloneCredentialsOut) SetSslSkipCertVerification(v bool)`

SetSslSkipCertVerification sets SslSkipCertVerification field to given value.


### SetSslSkipCertVerificationNil

`func (o *LookerGitCloneCredentialsOut) SetSslSkipCertVerificationNil(b bool)`

 SetSslSkipCertVerificationNil sets the value for SslSkipCertVerification to be an explicit nil

### UnsetSslSkipCertVerification
`func (o *LookerGitCloneCredentialsOut) UnsetSslSkipCertVerification()`

UnsetSslSkipCertVerification ensures that no value is present for SslSkipCertVerification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


