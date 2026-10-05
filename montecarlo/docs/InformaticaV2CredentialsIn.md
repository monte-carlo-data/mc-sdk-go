# InformaticaV2CredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **NullableString** | Informatica login URL for your POD. Leave it out for https://dm-us.informaticacloud.com. | [optional] 
**Username** | Pointer to **NullableString** | Informatica user, for &#x60;password&#x60;. | [optional] 
**OrgId** | Pointer to **NullableString** | Informatica organization ID, for &#x60;oauth&#x60;. | [optional] 
**OauthClientId** | Pointer to **NullableString** | Client ID of the app registered with your identity provider, for &#x60;oauth&#x60;. | [optional] 
**OauthGrantType** | Pointer to [**NullableOAuthGrantType**](OAuthGrantType.md) | Grant Monte Carlo requests the token with, for &#x60;oauth&#x60;. &#x60;password&#x60; also takes &#x60;oauth_username&#x60; and &#x60;oauth_password&#x60;. | [optional] 
**OauthAccessTokenEndpoint** | Pointer to **NullableString** | Identity provider URL Monte Carlo requests tokens from. | [optional] 
**OauthScope** | Pointer to **NullableString** | Scope to request the token with. Leave it out to request none. | [optional] 
**OauthUsername** | Pointer to **NullableString** | Identity provider user, for the &#x60;password&#x60; grant. | [optional] 
**Password** | Pointer to **NullableString** | Password of &#x60;username&#x60;, for &#x60;password&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**OauthClientSecret** | Pointer to **NullableString** | Secret of the identity provider app, for &#x60;oauth&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**OauthPassword** | Pointer to **NullableString** | Password of &#x60;oauth_username&#x60;, for the &#x60;password&#x60; grant. Stored by Monte Carlo and never returned. | [optional] 
**AuthMode** | [**InformaticaV2AuthMode**](InformaticaV2AuthMode.md) | How Monte Carlo signs in. &#x60;password&#x60; takes &#x60;username&#x60; and &#x60;password&#x60;. &#x60;oauth&#x60; takes &#x60;org_id&#x60; and the &#x60;oauth_*&#x60; fields. | 

## Methods

### NewInformaticaV2CredentialsIn

`func NewInformaticaV2CredentialsIn(authMode InformaticaV2AuthMode, ) *InformaticaV2CredentialsIn`

NewInformaticaV2CredentialsIn instantiates a new InformaticaV2CredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInformaticaV2CredentialsInWithDefaults

`func NewInformaticaV2CredentialsInWithDefaults() *InformaticaV2CredentialsIn`

NewInformaticaV2CredentialsInWithDefaults instantiates a new InformaticaV2CredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *InformaticaV2CredentialsIn) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *InformaticaV2CredentialsIn) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *InformaticaV2CredentialsIn) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *InformaticaV2CredentialsIn) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *InformaticaV2CredentialsIn) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *InformaticaV2CredentialsIn) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
### GetUsername

`func (o *InformaticaV2CredentialsIn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InformaticaV2CredentialsIn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InformaticaV2CredentialsIn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InformaticaV2CredentialsIn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *InformaticaV2CredentialsIn) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *InformaticaV2CredentialsIn) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetOrgId

`func (o *InformaticaV2CredentialsIn) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *InformaticaV2CredentialsIn) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *InformaticaV2CredentialsIn) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *InformaticaV2CredentialsIn) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *InformaticaV2CredentialsIn) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *InformaticaV2CredentialsIn) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetOauthClientId

`func (o *InformaticaV2CredentialsIn) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *InformaticaV2CredentialsIn) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *InformaticaV2CredentialsIn) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *InformaticaV2CredentialsIn) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### SetOauthClientIdNil

`func (o *InformaticaV2CredentialsIn) SetOauthClientIdNil(b bool)`

 SetOauthClientIdNil sets the value for OauthClientId to be an explicit nil

### UnsetOauthClientId
`func (o *InformaticaV2CredentialsIn) UnsetOauthClientId()`

UnsetOauthClientId ensures that no value is present for OauthClientId, not even an explicit nil
### GetOauthGrantType

`func (o *InformaticaV2CredentialsIn) GetOauthGrantType() OAuthGrantType`

GetOauthGrantType returns the OauthGrantType field if non-nil, zero value otherwise.

### GetOauthGrantTypeOk

`func (o *InformaticaV2CredentialsIn) GetOauthGrantTypeOk() (*OAuthGrantType, bool)`

GetOauthGrantTypeOk returns a tuple with the OauthGrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthGrantType

`func (o *InformaticaV2CredentialsIn) SetOauthGrantType(v OAuthGrantType)`

SetOauthGrantType sets OauthGrantType field to given value.

### HasOauthGrantType

`func (o *InformaticaV2CredentialsIn) HasOauthGrantType() bool`

HasOauthGrantType returns a boolean if a field has been set.

### SetOauthGrantTypeNil

`func (o *InformaticaV2CredentialsIn) SetOauthGrantTypeNil(b bool)`

 SetOauthGrantTypeNil sets the value for OauthGrantType to be an explicit nil

### UnsetOauthGrantType
`func (o *InformaticaV2CredentialsIn) UnsetOauthGrantType()`

UnsetOauthGrantType ensures that no value is present for OauthGrantType, not even an explicit nil
### GetOauthAccessTokenEndpoint

`func (o *InformaticaV2CredentialsIn) GetOauthAccessTokenEndpoint() string`

GetOauthAccessTokenEndpoint returns the OauthAccessTokenEndpoint field if non-nil, zero value otherwise.

### GetOauthAccessTokenEndpointOk

`func (o *InformaticaV2CredentialsIn) GetOauthAccessTokenEndpointOk() (*string, bool)`

GetOauthAccessTokenEndpointOk returns a tuple with the OauthAccessTokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthAccessTokenEndpoint

`func (o *InformaticaV2CredentialsIn) SetOauthAccessTokenEndpoint(v string)`

SetOauthAccessTokenEndpoint sets OauthAccessTokenEndpoint field to given value.

### HasOauthAccessTokenEndpoint

`func (o *InformaticaV2CredentialsIn) HasOauthAccessTokenEndpoint() bool`

HasOauthAccessTokenEndpoint returns a boolean if a field has been set.

### SetOauthAccessTokenEndpointNil

`func (o *InformaticaV2CredentialsIn) SetOauthAccessTokenEndpointNil(b bool)`

 SetOauthAccessTokenEndpointNil sets the value for OauthAccessTokenEndpoint to be an explicit nil

### UnsetOauthAccessTokenEndpoint
`func (o *InformaticaV2CredentialsIn) UnsetOauthAccessTokenEndpoint()`

UnsetOauthAccessTokenEndpoint ensures that no value is present for OauthAccessTokenEndpoint, not even an explicit nil
### GetOauthScope

`func (o *InformaticaV2CredentialsIn) GetOauthScope() string`

GetOauthScope returns the OauthScope field if non-nil, zero value otherwise.

### GetOauthScopeOk

`func (o *InformaticaV2CredentialsIn) GetOauthScopeOk() (*string, bool)`

GetOauthScopeOk returns a tuple with the OauthScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthScope

`func (o *InformaticaV2CredentialsIn) SetOauthScope(v string)`

SetOauthScope sets OauthScope field to given value.

### HasOauthScope

`func (o *InformaticaV2CredentialsIn) HasOauthScope() bool`

HasOauthScope returns a boolean if a field has been set.

### SetOauthScopeNil

`func (o *InformaticaV2CredentialsIn) SetOauthScopeNil(b bool)`

 SetOauthScopeNil sets the value for OauthScope to be an explicit nil

### UnsetOauthScope
`func (o *InformaticaV2CredentialsIn) UnsetOauthScope()`

UnsetOauthScope ensures that no value is present for OauthScope, not even an explicit nil
### GetOauthUsername

`func (o *InformaticaV2CredentialsIn) GetOauthUsername() string`

GetOauthUsername returns the OauthUsername field if non-nil, zero value otherwise.

### GetOauthUsernameOk

`func (o *InformaticaV2CredentialsIn) GetOauthUsernameOk() (*string, bool)`

GetOauthUsernameOk returns a tuple with the OauthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthUsername

`func (o *InformaticaV2CredentialsIn) SetOauthUsername(v string)`

SetOauthUsername sets OauthUsername field to given value.

### HasOauthUsername

`func (o *InformaticaV2CredentialsIn) HasOauthUsername() bool`

HasOauthUsername returns a boolean if a field has been set.

### SetOauthUsernameNil

`func (o *InformaticaV2CredentialsIn) SetOauthUsernameNil(b bool)`

 SetOauthUsernameNil sets the value for OauthUsername to be an explicit nil

### UnsetOauthUsername
`func (o *InformaticaV2CredentialsIn) UnsetOauthUsername()`

UnsetOauthUsername ensures that no value is present for OauthUsername, not even an explicit nil
### GetPassword

`func (o *InformaticaV2CredentialsIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InformaticaV2CredentialsIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InformaticaV2CredentialsIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InformaticaV2CredentialsIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *InformaticaV2CredentialsIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *InformaticaV2CredentialsIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetOauthClientSecret

`func (o *InformaticaV2CredentialsIn) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *InformaticaV2CredentialsIn) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *InformaticaV2CredentialsIn) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.

### HasOauthClientSecret

`func (o *InformaticaV2CredentialsIn) HasOauthClientSecret() bool`

HasOauthClientSecret returns a boolean if a field has been set.

### SetOauthClientSecretNil

`func (o *InformaticaV2CredentialsIn) SetOauthClientSecretNil(b bool)`

 SetOauthClientSecretNil sets the value for OauthClientSecret to be an explicit nil

### UnsetOauthClientSecret
`func (o *InformaticaV2CredentialsIn) UnsetOauthClientSecret()`

UnsetOauthClientSecret ensures that no value is present for OauthClientSecret, not even an explicit nil
### GetOauthPassword

`func (o *InformaticaV2CredentialsIn) GetOauthPassword() string`

GetOauthPassword returns the OauthPassword field if non-nil, zero value otherwise.

### GetOauthPasswordOk

`func (o *InformaticaV2CredentialsIn) GetOauthPasswordOk() (*string, bool)`

GetOauthPasswordOk returns a tuple with the OauthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthPassword

`func (o *InformaticaV2CredentialsIn) SetOauthPassword(v string)`

SetOauthPassword sets OauthPassword field to given value.

### HasOauthPassword

`func (o *InformaticaV2CredentialsIn) HasOauthPassword() bool`

HasOauthPassword returns a boolean if a field has been set.

### SetOauthPasswordNil

`func (o *InformaticaV2CredentialsIn) SetOauthPasswordNil(b bool)`

 SetOauthPasswordNil sets the value for OauthPassword to be an explicit nil

### UnsetOauthPassword
`func (o *InformaticaV2CredentialsIn) UnsetOauthPassword()`

UnsetOauthPassword ensures that no value is present for OauthPassword, not even an explicit nil
### GetAuthMode

`func (o *InformaticaV2CredentialsIn) GetAuthMode() InformaticaV2AuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *InformaticaV2CredentialsIn) GetAuthModeOk() (*InformaticaV2AuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *InformaticaV2CredentialsIn) SetAuthMode(v InformaticaV2AuthMode)`

SetAuthMode sets AuthMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


