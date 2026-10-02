# LookerGitCloneCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** | Git user for an HTTPS clone. Send it with &#x60;token&#x60;. | [optional] 
**SslCaData** | Pointer to **NullableString** | PEM certificate of the CA that signed the git server&#39;s certificate. | [optional] 
**SslSkipCertVerification** | Pointer to **NullableBool** | Skip verifying the git server&#39;s TLS certificate. | [optional] 
**Token** | Pointer to **NullableString** | Access token of &#x60;username&#x60;, for an HTTPS clone. Send this or &#x60;ssh_key&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**SshKey** | Pointer to **NullableString** | Private key for an SSH clone, as PEM text including its BEGIN and END lines. Send this or &#x60;username&#x60; and &#x60;token&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**RepoUrl** | **string** | Clone URL of the LookML repository, over HTTPS or SSH. | 

## Methods

### NewLookerGitCloneCredentialsIn

`func NewLookerGitCloneCredentialsIn(repoUrl string, ) *LookerGitCloneCredentialsIn`

NewLookerGitCloneCredentialsIn instantiates a new LookerGitCloneCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookerGitCloneCredentialsInWithDefaults

`func NewLookerGitCloneCredentialsInWithDefaults() *LookerGitCloneCredentialsIn`

NewLookerGitCloneCredentialsInWithDefaults instantiates a new LookerGitCloneCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *LookerGitCloneCredentialsIn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LookerGitCloneCredentialsIn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LookerGitCloneCredentialsIn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LookerGitCloneCredentialsIn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *LookerGitCloneCredentialsIn) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *LookerGitCloneCredentialsIn) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetSslCaData

`func (o *LookerGitCloneCredentialsIn) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *LookerGitCloneCredentialsIn) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *LookerGitCloneCredentialsIn) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.

### HasSslCaData

`func (o *LookerGitCloneCredentialsIn) HasSslCaData() bool`

HasSslCaData returns a boolean if a field has been set.

### SetSslCaDataNil

`func (o *LookerGitCloneCredentialsIn) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *LookerGitCloneCredentialsIn) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslSkipCertVerification

`func (o *LookerGitCloneCredentialsIn) GetSslSkipCertVerification() bool`

GetSslSkipCertVerification returns the SslSkipCertVerification field if non-nil, zero value otherwise.

### GetSslSkipCertVerificationOk

`func (o *LookerGitCloneCredentialsIn) GetSslSkipCertVerificationOk() (*bool, bool)`

GetSslSkipCertVerificationOk returns a tuple with the SslSkipCertVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSkipCertVerification

`func (o *LookerGitCloneCredentialsIn) SetSslSkipCertVerification(v bool)`

SetSslSkipCertVerification sets SslSkipCertVerification field to given value.

### HasSslSkipCertVerification

`func (o *LookerGitCloneCredentialsIn) HasSslSkipCertVerification() bool`

HasSslSkipCertVerification returns a boolean if a field has been set.

### SetSslSkipCertVerificationNil

`func (o *LookerGitCloneCredentialsIn) SetSslSkipCertVerificationNil(b bool)`

 SetSslSkipCertVerificationNil sets the value for SslSkipCertVerification to be an explicit nil

### UnsetSslSkipCertVerification
`func (o *LookerGitCloneCredentialsIn) UnsetSslSkipCertVerification()`

UnsetSslSkipCertVerification ensures that no value is present for SslSkipCertVerification, not even an explicit nil
### GetToken

`func (o *LookerGitCloneCredentialsIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LookerGitCloneCredentialsIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LookerGitCloneCredentialsIn) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LookerGitCloneCredentialsIn) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *LookerGitCloneCredentialsIn) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *LookerGitCloneCredentialsIn) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetSshKey

`func (o *LookerGitCloneCredentialsIn) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *LookerGitCloneCredentialsIn) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *LookerGitCloneCredentialsIn) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *LookerGitCloneCredentialsIn) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *LookerGitCloneCredentialsIn) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *LookerGitCloneCredentialsIn) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetRepoUrl

`func (o *LookerGitCloneCredentialsIn) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *LookerGitCloneCredentialsIn) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *LookerGitCloneCredentialsIn) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


