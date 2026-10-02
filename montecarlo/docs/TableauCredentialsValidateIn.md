# TableauCredentialsValidateIn

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
**Password** | Pointer to **NullableString** | Password of &#x60;username&#x60;. Send this, a personal access token, or a connected app. Used for this check and not kept. | [optional] 
**TokenValue** | Pointer to **NullableString** | Secret of the personal access token in &#x60;token_name&#x60;. Used for this check and not kept. | [optional] 
**ConnectedAppSecretValue** | Pointer to **NullableString** | Value of the connected app&#39;s secret in &#x60;connected_app_secret_id&#x60;. Used for this check and not kept. | [optional] 
**ServerName** | **string** | URL of the Tableau server, starting with https:// or http://. | 

## Methods

### NewTableauCredentialsValidateIn

`func NewTableauCredentialsValidateIn(deploymentId string, serverName string, ) *TableauCredentialsValidateIn`

NewTableauCredentialsValidateIn instantiates a new TableauCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableauCredentialsValidateInWithDefaults

`func NewTableauCredentialsValidateInWithDefaults() *TableauCredentialsValidateIn`

NewTableauCredentialsValidateInWithDefaults instantiates a new TableauCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *TableauCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *TableauCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *TableauCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetSiteName

`func (o *TableauCredentialsValidateIn) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *TableauCredentialsValidateIn) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *TableauCredentialsValidateIn) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *TableauCredentialsValidateIn) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *TableauCredentialsValidateIn) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *TableauCredentialsValidateIn) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetVerifySsl

`func (o *TableauCredentialsValidateIn) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *TableauCredentialsValidateIn) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *TableauCredentialsValidateIn) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *TableauCredentialsValidateIn) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.

### SetVerifySslNil

`func (o *TableauCredentialsValidateIn) SetVerifySslNil(b bool)`

 SetVerifySslNil sets the value for VerifySsl to be an explicit nil

### UnsetVerifySsl
`func (o *TableauCredentialsValidateIn) UnsetVerifySsl()`

UnsetVerifySsl ensures that no value is present for VerifySsl, not even an explicit nil
### GetUsername

`func (o *TableauCredentialsValidateIn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TableauCredentialsValidateIn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TableauCredentialsValidateIn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TableauCredentialsValidateIn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TableauCredentialsValidateIn) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TableauCredentialsValidateIn) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetTokenName

`func (o *TableauCredentialsValidateIn) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *TableauCredentialsValidateIn) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *TableauCredentialsValidateIn) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *TableauCredentialsValidateIn) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### SetTokenNameNil

`func (o *TableauCredentialsValidateIn) SetTokenNameNil(b bool)`

 SetTokenNameNil sets the value for TokenName to be an explicit nil

### UnsetTokenName
`func (o *TableauCredentialsValidateIn) UnsetTokenName()`

UnsetTokenName ensures that no value is present for TokenName, not even an explicit nil
### GetConnectedAppClientId

`func (o *TableauCredentialsValidateIn) GetConnectedAppClientId() string`

GetConnectedAppClientId returns the ConnectedAppClientId field if non-nil, zero value otherwise.

### GetConnectedAppClientIdOk

`func (o *TableauCredentialsValidateIn) GetConnectedAppClientIdOk() (*string, bool)`

GetConnectedAppClientIdOk returns a tuple with the ConnectedAppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppClientId

`func (o *TableauCredentialsValidateIn) SetConnectedAppClientId(v string)`

SetConnectedAppClientId sets ConnectedAppClientId field to given value.

### HasConnectedAppClientId

`func (o *TableauCredentialsValidateIn) HasConnectedAppClientId() bool`

HasConnectedAppClientId returns a boolean if a field has been set.

### SetConnectedAppClientIdNil

`func (o *TableauCredentialsValidateIn) SetConnectedAppClientIdNil(b bool)`

 SetConnectedAppClientIdNil sets the value for ConnectedAppClientId to be an explicit nil

### UnsetConnectedAppClientId
`func (o *TableauCredentialsValidateIn) UnsetConnectedAppClientId()`

UnsetConnectedAppClientId ensures that no value is present for ConnectedAppClientId, not even an explicit nil
### GetConnectedAppSecretId

`func (o *TableauCredentialsValidateIn) GetConnectedAppSecretId() string`

GetConnectedAppSecretId returns the ConnectedAppSecretId field if non-nil, zero value otherwise.

### GetConnectedAppSecretIdOk

`func (o *TableauCredentialsValidateIn) GetConnectedAppSecretIdOk() (*string, bool)`

GetConnectedAppSecretIdOk returns a tuple with the ConnectedAppSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppSecretId

`func (o *TableauCredentialsValidateIn) SetConnectedAppSecretId(v string)`

SetConnectedAppSecretId sets ConnectedAppSecretId field to given value.

### HasConnectedAppSecretId

`func (o *TableauCredentialsValidateIn) HasConnectedAppSecretId() bool`

HasConnectedAppSecretId returns a boolean if a field has been set.

### SetConnectedAppSecretIdNil

`func (o *TableauCredentialsValidateIn) SetConnectedAppSecretIdNil(b bool)`

 SetConnectedAppSecretIdNil sets the value for ConnectedAppSecretId to be an explicit nil

### UnsetConnectedAppSecretId
`func (o *TableauCredentialsValidateIn) UnsetConnectedAppSecretId()`

UnsetConnectedAppSecretId ensures that no value is present for ConnectedAppSecretId, not even an explicit nil
### GetPassword

`func (o *TableauCredentialsValidateIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TableauCredentialsValidateIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TableauCredentialsValidateIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TableauCredentialsValidateIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TableauCredentialsValidateIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TableauCredentialsValidateIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTokenValue

`func (o *TableauCredentialsValidateIn) GetTokenValue() string`

GetTokenValue returns the TokenValue field if non-nil, zero value otherwise.

### GetTokenValueOk

`func (o *TableauCredentialsValidateIn) GetTokenValueOk() (*string, bool)`

GetTokenValueOk returns a tuple with the TokenValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenValue

`func (o *TableauCredentialsValidateIn) SetTokenValue(v string)`

SetTokenValue sets TokenValue field to given value.

### HasTokenValue

`func (o *TableauCredentialsValidateIn) HasTokenValue() bool`

HasTokenValue returns a boolean if a field has been set.

### SetTokenValueNil

`func (o *TableauCredentialsValidateIn) SetTokenValueNil(b bool)`

 SetTokenValueNil sets the value for TokenValue to be an explicit nil

### UnsetTokenValue
`func (o *TableauCredentialsValidateIn) UnsetTokenValue()`

UnsetTokenValue ensures that no value is present for TokenValue, not even an explicit nil
### GetConnectedAppSecretValue

`func (o *TableauCredentialsValidateIn) GetConnectedAppSecretValue() string`

GetConnectedAppSecretValue returns the ConnectedAppSecretValue field if non-nil, zero value otherwise.

### GetConnectedAppSecretValueOk

`func (o *TableauCredentialsValidateIn) GetConnectedAppSecretValueOk() (*string, bool)`

GetConnectedAppSecretValueOk returns a tuple with the ConnectedAppSecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppSecretValue

`func (o *TableauCredentialsValidateIn) SetConnectedAppSecretValue(v string)`

SetConnectedAppSecretValue sets ConnectedAppSecretValue field to given value.

### HasConnectedAppSecretValue

`func (o *TableauCredentialsValidateIn) HasConnectedAppSecretValue() bool`

HasConnectedAppSecretValue returns a boolean if a field has been set.

### SetConnectedAppSecretValueNil

`func (o *TableauCredentialsValidateIn) SetConnectedAppSecretValueNil(b bool)`

 SetConnectedAppSecretValueNil sets the value for ConnectedAppSecretValue to be an explicit nil

### UnsetConnectedAppSecretValue
`func (o *TableauCredentialsValidateIn) UnsetConnectedAppSecretValue()`

UnsetConnectedAppSecretValue ensures that no value is present for ConnectedAppSecretValue, not even an explicit nil
### GetServerName

`func (o *TableauCredentialsValidateIn) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *TableauCredentialsValidateIn) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *TableauCredentialsValidateIn) SetServerName(v string)`

SetServerName sets ServerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


