# TableauCredentialsIn

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
**ServerName** | **string** | URL of the Tableau server, starting with https:// or http://. | 

## Methods

### NewTableauCredentialsIn

`func NewTableauCredentialsIn(serverName string, ) *TableauCredentialsIn`

NewTableauCredentialsIn instantiates a new TableauCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableauCredentialsInWithDefaults

`func NewTableauCredentialsInWithDefaults() *TableauCredentialsIn`

NewTableauCredentialsInWithDefaults instantiates a new TableauCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteName

`func (o *TableauCredentialsIn) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *TableauCredentialsIn) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *TableauCredentialsIn) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *TableauCredentialsIn) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### SetSiteNameNil

`func (o *TableauCredentialsIn) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *TableauCredentialsIn) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetVerifySsl

`func (o *TableauCredentialsIn) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *TableauCredentialsIn) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *TableauCredentialsIn) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *TableauCredentialsIn) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.

### SetVerifySslNil

`func (o *TableauCredentialsIn) SetVerifySslNil(b bool)`

 SetVerifySslNil sets the value for VerifySsl to be an explicit nil

### UnsetVerifySsl
`func (o *TableauCredentialsIn) UnsetVerifySsl()`

UnsetVerifySsl ensures that no value is present for VerifySsl, not even an explicit nil
### GetUsername

`func (o *TableauCredentialsIn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TableauCredentialsIn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TableauCredentialsIn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TableauCredentialsIn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TableauCredentialsIn) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TableauCredentialsIn) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetTokenName

`func (o *TableauCredentialsIn) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *TableauCredentialsIn) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *TableauCredentialsIn) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *TableauCredentialsIn) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### SetTokenNameNil

`func (o *TableauCredentialsIn) SetTokenNameNil(b bool)`

 SetTokenNameNil sets the value for TokenName to be an explicit nil

### UnsetTokenName
`func (o *TableauCredentialsIn) UnsetTokenName()`

UnsetTokenName ensures that no value is present for TokenName, not even an explicit nil
### GetConnectedAppClientId

`func (o *TableauCredentialsIn) GetConnectedAppClientId() string`

GetConnectedAppClientId returns the ConnectedAppClientId field if non-nil, zero value otherwise.

### GetConnectedAppClientIdOk

`func (o *TableauCredentialsIn) GetConnectedAppClientIdOk() (*string, bool)`

GetConnectedAppClientIdOk returns a tuple with the ConnectedAppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppClientId

`func (o *TableauCredentialsIn) SetConnectedAppClientId(v string)`

SetConnectedAppClientId sets ConnectedAppClientId field to given value.

### HasConnectedAppClientId

`func (o *TableauCredentialsIn) HasConnectedAppClientId() bool`

HasConnectedAppClientId returns a boolean if a field has been set.

### SetConnectedAppClientIdNil

`func (o *TableauCredentialsIn) SetConnectedAppClientIdNil(b bool)`

 SetConnectedAppClientIdNil sets the value for ConnectedAppClientId to be an explicit nil

### UnsetConnectedAppClientId
`func (o *TableauCredentialsIn) UnsetConnectedAppClientId()`

UnsetConnectedAppClientId ensures that no value is present for ConnectedAppClientId, not even an explicit nil
### GetConnectedAppSecretId

`func (o *TableauCredentialsIn) GetConnectedAppSecretId() string`

GetConnectedAppSecretId returns the ConnectedAppSecretId field if non-nil, zero value otherwise.

### GetConnectedAppSecretIdOk

`func (o *TableauCredentialsIn) GetConnectedAppSecretIdOk() (*string, bool)`

GetConnectedAppSecretIdOk returns a tuple with the ConnectedAppSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppSecretId

`func (o *TableauCredentialsIn) SetConnectedAppSecretId(v string)`

SetConnectedAppSecretId sets ConnectedAppSecretId field to given value.

### HasConnectedAppSecretId

`func (o *TableauCredentialsIn) HasConnectedAppSecretId() bool`

HasConnectedAppSecretId returns a boolean if a field has been set.

### SetConnectedAppSecretIdNil

`func (o *TableauCredentialsIn) SetConnectedAppSecretIdNil(b bool)`

 SetConnectedAppSecretIdNil sets the value for ConnectedAppSecretId to be an explicit nil

### UnsetConnectedAppSecretId
`func (o *TableauCredentialsIn) UnsetConnectedAppSecretId()`

UnsetConnectedAppSecretId ensures that no value is present for ConnectedAppSecretId, not even an explicit nil
### GetPassword

`func (o *TableauCredentialsIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TableauCredentialsIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TableauCredentialsIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TableauCredentialsIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TableauCredentialsIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TableauCredentialsIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTokenValue

`func (o *TableauCredentialsIn) GetTokenValue() string`

GetTokenValue returns the TokenValue field if non-nil, zero value otherwise.

### GetTokenValueOk

`func (o *TableauCredentialsIn) GetTokenValueOk() (*string, bool)`

GetTokenValueOk returns a tuple with the TokenValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenValue

`func (o *TableauCredentialsIn) SetTokenValue(v string)`

SetTokenValue sets TokenValue field to given value.

### HasTokenValue

`func (o *TableauCredentialsIn) HasTokenValue() bool`

HasTokenValue returns a boolean if a field has been set.

### SetTokenValueNil

`func (o *TableauCredentialsIn) SetTokenValueNil(b bool)`

 SetTokenValueNil sets the value for TokenValue to be an explicit nil

### UnsetTokenValue
`func (o *TableauCredentialsIn) UnsetTokenValue()`

UnsetTokenValue ensures that no value is present for TokenValue, not even an explicit nil
### GetConnectedAppSecretValue

`func (o *TableauCredentialsIn) GetConnectedAppSecretValue() string`

GetConnectedAppSecretValue returns the ConnectedAppSecretValue field if non-nil, zero value otherwise.

### GetConnectedAppSecretValueOk

`func (o *TableauCredentialsIn) GetConnectedAppSecretValueOk() (*string, bool)`

GetConnectedAppSecretValueOk returns a tuple with the ConnectedAppSecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppSecretValue

`func (o *TableauCredentialsIn) SetConnectedAppSecretValue(v string)`

SetConnectedAppSecretValue sets ConnectedAppSecretValue field to given value.

### HasConnectedAppSecretValue

`func (o *TableauCredentialsIn) HasConnectedAppSecretValue() bool`

HasConnectedAppSecretValue returns a boolean if a field has been set.

### SetConnectedAppSecretValueNil

`func (o *TableauCredentialsIn) SetConnectedAppSecretValueNil(b bool)`

 SetConnectedAppSecretValueNil sets the value for ConnectedAppSecretValue to be an explicit nil

### UnsetConnectedAppSecretValue
`func (o *TableauCredentialsIn) UnsetConnectedAppSecretValue()`

UnsetConnectedAppSecretValue ensures that no value is present for ConnectedAppSecretValue, not even an explicit nil
### GetServerName

`func (o *TableauCredentialsIn) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *TableauCredentialsIn) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *TableauCredentialsIn) SetServerName(v string)`

SetServerName sets ServerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


