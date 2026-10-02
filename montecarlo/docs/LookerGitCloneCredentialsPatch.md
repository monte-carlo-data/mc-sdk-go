# LookerGitCloneCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** | Git user for an HTTPS clone. Send it with &#x60;token&#x60;. | [optional] 
**SslCaData** | Pointer to **NullableString** | PEM certificate of the CA that signed the git server&#39;s certificate. | [optional] 
**SslSkipCertVerification** | Pointer to **NullableBool** | Skip verifying the git server&#39;s TLS certificate. | [optional] 
**Token** | Pointer to **NullableString** | Access token of &#x60;username&#x60;, for an HTTPS clone. Send this or &#x60;ssh_key&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**SshKey** | Pointer to **NullableString** | Private key for an SSH clone, as PEM text including its BEGIN and END lines. Send this or &#x60;username&#x60; and &#x60;token&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**RepoUrl** | Pointer to **NullableString** | Clone URL of the LookML repository, over HTTPS or SSH. | [optional] 

## Methods

### NewLookerGitCloneCredentialsPatch

`func NewLookerGitCloneCredentialsPatch() *LookerGitCloneCredentialsPatch`

NewLookerGitCloneCredentialsPatch instantiates a new LookerGitCloneCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookerGitCloneCredentialsPatchWithDefaults

`func NewLookerGitCloneCredentialsPatchWithDefaults() *LookerGitCloneCredentialsPatch`

NewLookerGitCloneCredentialsPatchWithDefaults instantiates a new LookerGitCloneCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *LookerGitCloneCredentialsPatch) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LookerGitCloneCredentialsPatch) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LookerGitCloneCredentialsPatch) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LookerGitCloneCredentialsPatch) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *LookerGitCloneCredentialsPatch) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *LookerGitCloneCredentialsPatch) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetSslCaData

`func (o *LookerGitCloneCredentialsPatch) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *LookerGitCloneCredentialsPatch) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *LookerGitCloneCredentialsPatch) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.

### HasSslCaData

`func (o *LookerGitCloneCredentialsPatch) HasSslCaData() bool`

HasSslCaData returns a boolean if a field has been set.

### SetSslCaDataNil

`func (o *LookerGitCloneCredentialsPatch) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *LookerGitCloneCredentialsPatch) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslSkipCertVerification

`func (o *LookerGitCloneCredentialsPatch) GetSslSkipCertVerification() bool`

GetSslSkipCertVerification returns the SslSkipCertVerification field if non-nil, zero value otherwise.

### GetSslSkipCertVerificationOk

`func (o *LookerGitCloneCredentialsPatch) GetSslSkipCertVerificationOk() (*bool, bool)`

GetSslSkipCertVerificationOk returns a tuple with the SslSkipCertVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSkipCertVerification

`func (o *LookerGitCloneCredentialsPatch) SetSslSkipCertVerification(v bool)`

SetSslSkipCertVerification sets SslSkipCertVerification field to given value.

### HasSslSkipCertVerification

`func (o *LookerGitCloneCredentialsPatch) HasSslSkipCertVerification() bool`

HasSslSkipCertVerification returns a boolean if a field has been set.

### SetSslSkipCertVerificationNil

`func (o *LookerGitCloneCredentialsPatch) SetSslSkipCertVerificationNil(b bool)`

 SetSslSkipCertVerificationNil sets the value for SslSkipCertVerification to be an explicit nil

### UnsetSslSkipCertVerification
`func (o *LookerGitCloneCredentialsPatch) UnsetSslSkipCertVerification()`

UnsetSslSkipCertVerification ensures that no value is present for SslSkipCertVerification, not even an explicit nil
### GetToken

`func (o *LookerGitCloneCredentialsPatch) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LookerGitCloneCredentialsPatch) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LookerGitCloneCredentialsPatch) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LookerGitCloneCredentialsPatch) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *LookerGitCloneCredentialsPatch) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *LookerGitCloneCredentialsPatch) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetSshKey

`func (o *LookerGitCloneCredentialsPatch) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *LookerGitCloneCredentialsPatch) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *LookerGitCloneCredentialsPatch) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *LookerGitCloneCredentialsPatch) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *LookerGitCloneCredentialsPatch) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *LookerGitCloneCredentialsPatch) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetRepoUrl

`func (o *LookerGitCloneCredentialsPatch) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *LookerGitCloneCredentialsPatch) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *LookerGitCloneCredentialsPatch) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *LookerGitCloneCredentialsPatch) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.

### SetRepoUrlNil

`func (o *LookerGitCloneCredentialsPatch) SetRepoUrlNil(b bool)`

 SetRepoUrlNil sets the value for RepoUrl to be an explicit nil

### UnsetRepoUrl
`func (o *LookerGitCloneCredentialsPatch) UnsetRepoUrl()`

UnsetRepoUrl ensures that no value is present for RepoUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


