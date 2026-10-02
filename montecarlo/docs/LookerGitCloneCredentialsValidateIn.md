# LookerGitCloneCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Username** | Pointer to **NullableString** | Git user for an HTTPS clone. Send it with &#x60;token&#x60;. | [optional] 
**SslCaData** | Pointer to **NullableString** | PEM certificate of the CA that signed the git server&#39;s certificate. | [optional] 
**SslSkipCertVerification** | Pointer to **NullableBool** | Skip verifying the git server&#39;s TLS certificate. | [optional] 
**Token** | Pointer to **NullableString** | Access token of &#x60;username&#x60;, for an HTTPS clone. Send this or &#x60;ssh_key&#x60;. Used for this check and not kept. | [optional] 
**SshKey** | Pointer to **NullableString** | Private key for an SSH clone, as PEM text including its BEGIN and END lines. Send this or &#x60;username&#x60; and &#x60;token&#x60;. Used for this check and not kept. | [optional] 
**RepoUrl** | **string** | Clone URL of the LookML repository, over HTTPS or SSH. | 

## Methods

### NewLookerGitCloneCredentialsValidateIn

`func NewLookerGitCloneCredentialsValidateIn(deploymentId string, repoUrl string, ) *LookerGitCloneCredentialsValidateIn`

NewLookerGitCloneCredentialsValidateIn instantiates a new LookerGitCloneCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookerGitCloneCredentialsValidateInWithDefaults

`func NewLookerGitCloneCredentialsValidateInWithDefaults() *LookerGitCloneCredentialsValidateIn`

NewLookerGitCloneCredentialsValidateInWithDefaults instantiates a new LookerGitCloneCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *LookerGitCloneCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *LookerGitCloneCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *LookerGitCloneCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetUsername

`func (o *LookerGitCloneCredentialsValidateIn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LookerGitCloneCredentialsValidateIn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LookerGitCloneCredentialsValidateIn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LookerGitCloneCredentialsValidateIn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *LookerGitCloneCredentialsValidateIn) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *LookerGitCloneCredentialsValidateIn) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetSslCaData

`func (o *LookerGitCloneCredentialsValidateIn) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *LookerGitCloneCredentialsValidateIn) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *LookerGitCloneCredentialsValidateIn) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.

### HasSslCaData

`func (o *LookerGitCloneCredentialsValidateIn) HasSslCaData() bool`

HasSslCaData returns a boolean if a field has been set.

### SetSslCaDataNil

`func (o *LookerGitCloneCredentialsValidateIn) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *LookerGitCloneCredentialsValidateIn) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslSkipCertVerification

`func (o *LookerGitCloneCredentialsValidateIn) GetSslSkipCertVerification() bool`

GetSslSkipCertVerification returns the SslSkipCertVerification field if non-nil, zero value otherwise.

### GetSslSkipCertVerificationOk

`func (o *LookerGitCloneCredentialsValidateIn) GetSslSkipCertVerificationOk() (*bool, bool)`

GetSslSkipCertVerificationOk returns a tuple with the SslSkipCertVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSkipCertVerification

`func (o *LookerGitCloneCredentialsValidateIn) SetSslSkipCertVerification(v bool)`

SetSslSkipCertVerification sets SslSkipCertVerification field to given value.

### HasSslSkipCertVerification

`func (o *LookerGitCloneCredentialsValidateIn) HasSslSkipCertVerification() bool`

HasSslSkipCertVerification returns a boolean if a field has been set.

### SetSslSkipCertVerificationNil

`func (o *LookerGitCloneCredentialsValidateIn) SetSslSkipCertVerificationNil(b bool)`

 SetSslSkipCertVerificationNil sets the value for SslSkipCertVerification to be an explicit nil

### UnsetSslSkipCertVerification
`func (o *LookerGitCloneCredentialsValidateIn) UnsetSslSkipCertVerification()`

UnsetSslSkipCertVerification ensures that no value is present for SslSkipCertVerification, not even an explicit nil
### GetToken

`func (o *LookerGitCloneCredentialsValidateIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LookerGitCloneCredentialsValidateIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LookerGitCloneCredentialsValidateIn) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LookerGitCloneCredentialsValidateIn) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *LookerGitCloneCredentialsValidateIn) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *LookerGitCloneCredentialsValidateIn) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetSshKey

`func (o *LookerGitCloneCredentialsValidateIn) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *LookerGitCloneCredentialsValidateIn) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *LookerGitCloneCredentialsValidateIn) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *LookerGitCloneCredentialsValidateIn) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *LookerGitCloneCredentialsValidateIn) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *LookerGitCloneCredentialsValidateIn) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetRepoUrl

`func (o *LookerGitCloneCredentialsValidateIn) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *LookerGitCloneCredentialsValidateIn) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *LookerGitCloneCredentialsValidateIn) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


