# TableauCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteName** | Pointer to **NullableString** | Tableau site to connect to. Leave it out for the default site. | [optional] 
**VerifySsl** | Pointer to **NullableBool** | Whether to verify the server&#39;s TLS certificate. Verified when left out. | [optional] 
**Username** | Pointer to **NullableString** | Tableau user. Send it with &#x60;password&#x60;, or with a connected app. Not used with a personal access token. | [optional] 
**TokenName** | Pointer to **NullableString** | Name of a personal access token. Send it with &#x60;token_value&#x60;. | [optional] 
**ConnectedAppClientId** | Pointer to **NullableString** | Client ID of a Tableau connected app. Send it with &#x60;connected_app_secret_id&#x60;, &#x60;connected_app_secret_value&#x60; and &#x60;username&#x60;. | [optional] 
**ConnectedAppSecretId** | Pointer to **NullableString** | ID of the connected app&#39;s secret. An identifier, not the secret itself. | [optional] 
**Password** | Pointer to **NullableString** | Password of &#x60;username&#x60;. Send this, a personal access token, or a connected app. Stored by Monte Carlo and never returned. | [optional] 
**TokenValue** | Pointer to **NullableString** | Secret of the personal access token in &#x60;token_name&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**ConnectedAppSecretValue** | Pointer to **NullableString** | Value of the connected app&#39;s secret in &#x60;connected_app_secret_id&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**ServerName** | Pointer to **NullableString** | URL of the Tableau server, starting with https:// or http://. | [optional] 

## Methods

### NewTableauCredentialsPatch

`func NewTableauCredentialsPatch() *TableauCredentialsPatch`

NewTableauCredentialsPatch instantiates a new TableauCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableauCredentialsPatchWithDefaults

`func NewTableauCredentialsPatchWithDefaults() *TableauCredentialsPatch`

NewTableauCredentialsPatchWithDefaults instantiates a new TableauCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteName

`func (o *TableauCredentialsPatch) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *TableauCredentialsPatch) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *TableauCredentialsPatch) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *TableauCredentialsPatch) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *TableauCredentialsPatch) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *TableauCredentialsPatch) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetVerifySsl

`func (o *TableauCredentialsPatch) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *TableauCredentialsPatch) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *TableauCredentialsPatch) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *TableauCredentialsPatch) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.

### SetVerifySslNil

`func (o *TableauCredentialsPatch) SetVerifySslNil(b bool)`

 SetVerifySslNil sets the value for VerifySsl to be an explicit nil

### UnsetVerifySsl
`func (o *TableauCredentialsPatch) UnsetVerifySsl()`

UnsetVerifySsl ensures that no value is present for VerifySsl, not even an explicit nil
### GetUsername

`func (o *TableauCredentialsPatch) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TableauCredentialsPatch) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TableauCredentialsPatch) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TableauCredentialsPatch) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TableauCredentialsPatch) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TableauCredentialsPatch) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetTokenName

`func (o *TableauCredentialsPatch) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *TableauCredentialsPatch) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *TableauCredentialsPatch) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *TableauCredentialsPatch) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### SetTokenNameNil

`func (o *TableauCredentialsPatch) SetTokenNameNil(b bool)`

 SetTokenNameNil sets the value for TokenName to be an explicit nil

### UnsetTokenName
`func (o *TableauCredentialsPatch) UnsetTokenName()`

UnsetTokenName ensures that no value is present for TokenName, not even an explicit nil
### GetConnectedAppClientId

`func (o *TableauCredentialsPatch) GetConnectedAppClientId() string`

GetConnectedAppClientId returns the ConnectedAppClientId field if non-nil, zero value otherwise.

### GetConnectedAppClientIdOk

`func (o *TableauCredentialsPatch) GetConnectedAppClientIdOk() (*string, bool)`

GetConnectedAppClientIdOk returns a tuple with the ConnectedAppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppClientId

`func (o *TableauCredentialsPatch) SetConnectedAppClientId(v string)`

SetConnectedAppClientId sets ConnectedAppClientId field to given value.

### HasConnectedAppClientId

`func (o *TableauCredentialsPatch) HasConnectedAppClientId() bool`

HasConnectedAppClientId returns a boolean if a field has been set.

### SetConnectedAppClientIdNil

`func (o *TableauCredentialsPatch) SetConnectedAppClientIdNil(b bool)`

 SetConnectedAppClientIdNil sets the value for ConnectedAppClientId to be an explicit nil

### UnsetConnectedAppClientId
`func (o *TableauCredentialsPatch) UnsetConnectedAppClientId()`

UnsetConnectedAppClientId ensures that no value is present for ConnectedAppClientId, not even an explicit nil
### GetConnectedAppSecretId

`func (o *TableauCredentialsPatch) GetConnectedAppSecretId() string`

GetConnectedAppSecretId returns the ConnectedAppSecretId field if non-nil, zero value otherwise.

### GetConnectedAppSecretIdOk

`func (o *TableauCredentialsPatch) GetConnectedAppSecretIdOk() (*string, bool)`

GetConnectedAppSecretIdOk returns a tuple with the ConnectedAppSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppSecretId

`func (o *TableauCredentialsPatch) SetConnectedAppSecretId(v string)`

SetConnectedAppSecretId sets ConnectedAppSecretId field to given value.

### HasConnectedAppSecretId

`func (o *TableauCredentialsPatch) HasConnectedAppSecretId() bool`

HasConnectedAppSecretId returns a boolean if a field has been set.

### SetConnectedAppSecretIdNil

`func (o *TableauCredentialsPatch) SetConnectedAppSecretIdNil(b bool)`

 SetConnectedAppSecretIdNil sets the value for ConnectedAppSecretId to be an explicit nil

### UnsetConnectedAppSecretId
`func (o *TableauCredentialsPatch) UnsetConnectedAppSecretId()`

UnsetConnectedAppSecretId ensures that no value is present for ConnectedAppSecretId, not even an explicit nil
### GetPassword

`func (o *TableauCredentialsPatch) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TableauCredentialsPatch) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TableauCredentialsPatch) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TableauCredentialsPatch) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TableauCredentialsPatch) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TableauCredentialsPatch) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTokenValue

`func (o *TableauCredentialsPatch) GetTokenValue() string`

GetTokenValue returns the TokenValue field if non-nil, zero value otherwise.

### GetTokenValueOk

`func (o *TableauCredentialsPatch) GetTokenValueOk() (*string, bool)`

GetTokenValueOk returns a tuple with the TokenValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenValue

`func (o *TableauCredentialsPatch) SetTokenValue(v string)`

SetTokenValue sets TokenValue field to given value.

### HasTokenValue

`func (o *TableauCredentialsPatch) HasTokenValue() bool`

HasTokenValue returns a boolean if a field has been set.

### SetTokenValueNil

`func (o *TableauCredentialsPatch) SetTokenValueNil(b bool)`

 SetTokenValueNil sets the value for TokenValue to be an explicit nil

### UnsetTokenValue
`func (o *TableauCredentialsPatch) UnsetTokenValue()`

UnsetTokenValue ensures that no value is present for TokenValue, not even an explicit nil
### GetConnectedAppSecretValue

`func (o *TableauCredentialsPatch) GetConnectedAppSecretValue() string`

GetConnectedAppSecretValue returns the ConnectedAppSecretValue field if non-nil, zero value otherwise.

### GetConnectedAppSecretValueOk

`func (o *TableauCredentialsPatch) GetConnectedAppSecretValueOk() (*string, bool)`

GetConnectedAppSecretValueOk returns a tuple with the ConnectedAppSecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppSecretValue

`func (o *TableauCredentialsPatch) SetConnectedAppSecretValue(v string)`

SetConnectedAppSecretValue sets ConnectedAppSecretValue field to given value.

### HasConnectedAppSecretValue

`func (o *TableauCredentialsPatch) HasConnectedAppSecretValue() bool`

HasConnectedAppSecretValue returns a boolean if a field has been set.

### SetConnectedAppSecretValueNil

`func (o *TableauCredentialsPatch) SetConnectedAppSecretValueNil(b bool)`

 SetConnectedAppSecretValueNil sets the value for ConnectedAppSecretValue to be an explicit nil

### UnsetConnectedAppSecretValue
`func (o *TableauCredentialsPatch) UnsetConnectedAppSecretValue()`

UnsetConnectedAppSecretValue ensures that no value is present for ConnectedAppSecretValue, not even an explicit nil
### GetServerName

`func (o *TableauCredentialsPatch) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *TableauCredentialsPatch) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *TableauCredentialsPatch) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *TableauCredentialsPatch) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *TableauCredentialsPatch) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *TableauCredentialsPatch) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


