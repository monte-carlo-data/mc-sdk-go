# TableauCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**SiteName** | Pointer to **NullableString** | Tableau site to connect to. Leave it out for the default site. | [optional] 
**VerifySsl** | Pointer to **NullableBool** | Whether to verify the server&#39;s TLS certificate. Verified when left out. | [optional] 
**Username** | Pointer to **NullableString** | Tableau user. Send it with &#x60;password&#x60;, or with a connected app. Not used with a personal access token. | [optional] 
**TokenName** | Pointer to **NullableString** | Name of a personal access token. Send it with &#x60;token_value&#x60;. | [optional] 
**ConnectedAppClientId** | Pointer to **NullableString** | Client ID of a Tableau connected app. Send it with &#x60;connected_app_secret_id&#x60;, &#x60;connected_app_secret_value&#x60; and &#x60;username&#x60;. | [optional] 
**ConnectedAppSecretId** | Pointer to **NullableString** | ID of the connected app&#39;s secret. An identifier, not the secret itself. | [optional] 
**Password** | Pointer to **NullableString** | Password of &#x60;username&#x60;. Send this, a personal access token, or a connected app. Stored by Monte Carlo and never returned. | [optional] 
**TokenValue** | Pointer to **NullableString** | Secret of the personal access token in &#x60;token_name&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**ConnectedAppSecretValue** | Pointer to **NullableString** | Value of the connected app&#39;s secret in &#x60;connected_app_secret_id&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**ServerName** | Pointer to **NullableString** | URL of the Tableau server, starting with https:// or http://. Changing it requires sending the credentials&#39; secrets again. | [optional] 

## Methods

### NewTableauCredentialsChangeValidateIn

`func NewTableauCredentialsChangeValidateIn(deploymentId string, ) *TableauCredentialsChangeValidateIn`

NewTableauCredentialsChangeValidateIn instantiates a new TableauCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableauCredentialsChangeValidateInWithDefaults

`func NewTableauCredentialsChangeValidateInWithDefaults() *TableauCredentialsChangeValidateIn`

NewTableauCredentialsChangeValidateInWithDefaults instantiates a new TableauCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *TableauCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *TableauCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *TableauCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetSiteName

`func (o *TableauCredentialsChangeValidateIn) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *TableauCredentialsChangeValidateIn) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *TableauCredentialsChangeValidateIn) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *TableauCredentialsChangeValidateIn) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *TableauCredentialsChangeValidateIn) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *TableauCredentialsChangeValidateIn) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetVerifySsl

`func (o *TableauCredentialsChangeValidateIn) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *TableauCredentialsChangeValidateIn) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *TableauCredentialsChangeValidateIn) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *TableauCredentialsChangeValidateIn) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.

### SetVerifySslNil

`func (o *TableauCredentialsChangeValidateIn) SetVerifySslNil(b bool)`

 SetVerifySslNil sets the value for VerifySsl to be an explicit nil

### UnsetVerifySsl
`func (o *TableauCredentialsChangeValidateIn) UnsetVerifySsl()`

UnsetVerifySsl ensures that no value is present for VerifySsl, not even an explicit nil
### GetUsername

`func (o *TableauCredentialsChangeValidateIn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TableauCredentialsChangeValidateIn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TableauCredentialsChangeValidateIn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TableauCredentialsChangeValidateIn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TableauCredentialsChangeValidateIn) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TableauCredentialsChangeValidateIn) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetTokenName

`func (o *TableauCredentialsChangeValidateIn) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *TableauCredentialsChangeValidateIn) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *TableauCredentialsChangeValidateIn) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *TableauCredentialsChangeValidateIn) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### SetTokenNameNil

`func (o *TableauCredentialsChangeValidateIn) SetTokenNameNil(b bool)`

 SetTokenNameNil sets the value for TokenName to be an explicit nil

### UnsetTokenName
`func (o *TableauCredentialsChangeValidateIn) UnsetTokenName()`

UnsetTokenName ensures that no value is present for TokenName, not even an explicit nil
### GetConnectedAppClientId

`func (o *TableauCredentialsChangeValidateIn) GetConnectedAppClientId() string`

GetConnectedAppClientId returns the ConnectedAppClientId field if non-nil, zero value otherwise.

### GetConnectedAppClientIdOk

`func (o *TableauCredentialsChangeValidateIn) GetConnectedAppClientIdOk() (*string, bool)`

GetConnectedAppClientIdOk returns a tuple with the ConnectedAppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppClientId

`func (o *TableauCredentialsChangeValidateIn) SetConnectedAppClientId(v string)`

SetConnectedAppClientId sets ConnectedAppClientId field to given value.

### HasConnectedAppClientId

`func (o *TableauCredentialsChangeValidateIn) HasConnectedAppClientId() bool`

HasConnectedAppClientId returns a boolean if a field has been set.

### SetConnectedAppClientIdNil

`func (o *TableauCredentialsChangeValidateIn) SetConnectedAppClientIdNil(b bool)`

 SetConnectedAppClientIdNil sets the value for ConnectedAppClientId to be an explicit nil

### UnsetConnectedAppClientId
`func (o *TableauCredentialsChangeValidateIn) UnsetConnectedAppClientId()`

UnsetConnectedAppClientId ensures that no value is present for ConnectedAppClientId, not even an explicit nil
### GetConnectedAppSecretId

`func (o *TableauCredentialsChangeValidateIn) GetConnectedAppSecretId() string`

GetConnectedAppSecretId returns the ConnectedAppSecretId field if non-nil, zero value otherwise.

### GetConnectedAppSecretIdOk

`func (o *TableauCredentialsChangeValidateIn) GetConnectedAppSecretIdOk() (*string, bool)`

GetConnectedAppSecretIdOk returns a tuple with the ConnectedAppSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppSecretId

`func (o *TableauCredentialsChangeValidateIn) SetConnectedAppSecretId(v string)`

SetConnectedAppSecretId sets ConnectedAppSecretId field to given value.

### HasConnectedAppSecretId

`func (o *TableauCredentialsChangeValidateIn) HasConnectedAppSecretId() bool`

HasConnectedAppSecretId returns a boolean if a field has been set.

### SetConnectedAppSecretIdNil

`func (o *TableauCredentialsChangeValidateIn) SetConnectedAppSecretIdNil(b bool)`

 SetConnectedAppSecretIdNil sets the value for ConnectedAppSecretId to be an explicit nil

### UnsetConnectedAppSecretId
`func (o *TableauCredentialsChangeValidateIn) UnsetConnectedAppSecretId()`

UnsetConnectedAppSecretId ensures that no value is present for ConnectedAppSecretId, not even an explicit nil
### GetPassword

`func (o *TableauCredentialsChangeValidateIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TableauCredentialsChangeValidateIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TableauCredentialsChangeValidateIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TableauCredentialsChangeValidateIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TableauCredentialsChangeValidateIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TableauCredentialsChangeValidateIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTokenValue

`func (o *TableauCredentialsChangeValidateIn) GetTokenValue() string`

GetTokenValue returns the TokenValue field if non-nil, zero value otherwise.

### GetTokenValueOk

`func (o *TableauCredentialsChangeValidateIn) GetTokenValueOk() (*string, bool)`

GetTokenValueOk returns a tuple with the TokenValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenValue

`func (o *TableauCredentialsChangeValidateIn) SetTokenValue(v string)`

SetTokenValue sets TokenValue field to given value.

### HasTokenValue

`func (o *TableauCredentialsChangeValidateIn) HasTokenValue() bool`

HasTokenValue returns a boolean if a field has been set.

### SetTokenValueNil

`func (o *TableauCredentialsChangeValidateIn) SetTokenValueNil(b bool)`

 SetTokenValueNil sets the value for TokenValue to be an explicit nil

### UnsetTokenValue
`func (o *TableauCredentialsChangeValidateIn) UnsetTokenValue()`

UnsetTokenValue ensures that no value is present for TokenValue, not even an explicit nil
### GetConnectedAppSecretValue

`func (o *TableauCredentialsChangeValidateIn) GetConnectedAppSecretValue() string`

GetConnectedAppSecretValue returns the ConnectedAppSecretValue field if non-nil, zero value otherwise.

### GetConnectedAppSecretValueOk

`func (o *TableauCredentialsChangeValidateIn) GetConnectedAppSecretValueOk() (*string, bool)`

GetConnectedAppSecretValueOk returns a tuple with the ConnectedAppSecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppSecretValue

`func (o *TableauCredentialsChangeValidateIn) SetConnectedAppSecretValue(v string)`

SetConnectedAppSecretValue sets ConnectedAppSecretValue field to given value.

### HasConnectedAppSecretValue

`func (o *TableauCredentialsChangeValidateIn) HasConnectedAppSecretValue() bool`

HasConnectedAppSecretValue returns a boolean if a field has been set.

### SetConnectedAppSecretValueNil

`func (o *TableauCredentialsChangeValidateIn) SetConnectedAppSecretValueNil(b bool)`

 SetConnectedAppSecretValueNil sets the value for ConnectedAppSecretValue to be an explicit nil

### UnsetConnectedAppSecretValue
`func (o *TableauCredentialsChangeValidateIn) UnsetConnectedAppSecretValue()`

UnsetConnectedAppSecretValue ensures that no value is present for ConnectedAppSecretValue, not even an explicit nil
### GetServerName

`func (o *TableauCredentialsChangeValidateIn) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *TableauCredentialsChangeValidateIn) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *TableauCredentialsChangeValidateIn) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *TableauCredentialsChangeValidateIn) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *TableauCredentialsChangeValidateIn) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *TableauCredentialsChangeValidateIn) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


