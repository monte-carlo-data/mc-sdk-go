# InformaticaV2CredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**AuthMode** | [**InformaticaV2AuthMode**](InformaticaV2AuthMode.md) | How Monte Carlo signs in. &#x60;password&#x60; takes &#x60;username&#x60; and &#x60;password&#x60;. &#x60;oauth&#x60; takes &#x60;org_id&#x60; and the &#x60;oauth_*&#x60; fields. | 
**BaseUrl** | **NullableString** | Informatica login URL for your POD. Leave it out for https://dm-us.informaticacloud.com. Null unless set. | 
**Username** | **NullableString** | Informatica user. Null unless &#x60;auth_mode&#x60; is password. | 
**OrgId** | **NullableString** | Informatica organization ID. Null unless &#x60;auth_mode&#x60; is oauth. | 
**OauthClientId** | **NullableString** | Client ID of the identity provider app. Null unless &#x60;auth_mode&#x60; is oauth. | 
**OauthGrantType** | [**NullableOAuthGrantType**](OAuthGrantType.md) | Grant Monte Carlo requests the token with. Null unless &#x60;auth_mode&#x60; is oauth. | 
**OauthAccessTokenEndpoint** | **NullableString** | Identity provider URL Monte Carlo requests tokens from. Null unless &#x60;auth_mode&#x60; is oauth. | 
**OauthScope** | **NullableString** | Scope the token is requested with. Null unless set. | 
**OauthUsername** | **NullableString** | Identity provider user. Null unless the grant is password. | 

## Methods

### NewInformaticaV2CredentialsOut

`func NewInformaticaV2CredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, authMode InformaticaV2AuthMode, baseUrl NullableString, username NullableString, orgId NullableString, oauthClientId NullableString, oauthGrantType NullableOAuthGrantType, oauthAccessTokenEndpoint NullableString, oauthScope NullableString, oauthUsername NullableString, ) *InformaticaV2CredentialsOut`

NewInformaticaV2CredentialsOut instantiates a new InformaticaV2CredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInformaticaV2CredentialsOutWithDefaults

`func NewInformaticaV2CredentialsOutWithDefaults() *InformaticaV2CredentialsOut`

NewInformaticaV2CredentialsOutWithDefaults instantiates a new InformaticaV2CredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InformaticaV2CredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InformaticaV2CredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InformaticaV2CredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *InformaticaV2CredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *InformaticaV2CredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *InformaticaV2CredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *InformaticaV2CredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *InformaticaV2CredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *InformaticaV2CredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *InformaticaV2CredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *InformaticaV2CredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *InformaticaV2CredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetAuthMode

`func (o *InformaticaV2CredentialsOut) GetAuthMode() InformaticaV2AuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *InformaticaV2CredentialsOut) GetAuthModeOk() (*InformaticaV2AuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *InformaticaV2CredentialsOut) SetAuthMode(v InformaticaV2AuthMode)`

SetAuthMode sets AuthMode field to given value.


### GetBaseUrl

`func (o *InformaticaV2CredentialsOut) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *InformaticaV2CredentialsOut) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *InformaticaV2CredentialsOut) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### SetBaseUrlNil

`func (o *InformaticaV2CredentialsOut) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *InformaticaV2CredentialsOut) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
### GetUsername

`func (o *InformaticaV2CredentialsOut) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InformaticaV2CredentialsOut) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InformaticaV2CredentialsOut) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *InformaticaV2CredentialsOut) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *InformaticaV2CredentialsOut) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetOrgId

`func (o *InformaticaV2CredentialsOut) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *InformaticaV2CredentialsOut) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *InformaticaV2CredentialsOut) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### SetOrgIdNil

`func (o *InformaticaV2CredentialsOut) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *InformaticaV2CredentialsOut) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetOauthClientId

`func (o *InformaticaV2CredentialsOut) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *InformaticaV2CredentialsOut) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *InformaticaV2CredentialsOut) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.


### SetOauthClientIdNil

`func (o *InformaticaV2CredentialsOut) SetOauthClientIdNil(b bool)`

 SetOauthClientIdNil sets the value for OauthClientId to be an explicit nil

### UnsetOauthClientId
`func (o *InformaticaV2CredentialsOut) UnsetOauthClientId()`

UnsetOauthClientId ensures that no value is present for OauthClientId, not even an explicit nil
### GetOauthGrantType

`func (o *InformaticaV2CredentialsOut) GetOauthGrantType() OAuthGrantType`

GetOauthGrantType returns the OauthGrantType field if non-nil, zero value otherwise.

### GetOauthGrantTypeOk

`func (o *InformaticaV2CredentialsOut) GetOauthGrantTypeOk() (*OAuthGrantType, bool)`

GetOauthGrantTypeOk returns a tuple with the OauthGrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthGrantType

`func (o *InformaticaV2CredentialsOut) SetOauthGrantType(v OAuthGrantType)`

SetOauthGrantType sets OauthGrantType field to given value.


### SetOauthGrantTypeNil

`func (o *InformaticaV2CredentialsOut) SetOauthGrantTypeNil(b bool)`

 SetOauthGrantTypeNil sets the value for OauthGrantType to be an explicit nil

### UnsetOauthGrantType
`func (o *InformaticaV2CredentialsOut) UnsetOauthGrantType()`

UnsetOauthGrantType ensures that no value is present for OauthGrantType, not even an explicit nil
### GetOauthAccessTokenEndpoint

`func (o *InformaticaV2CredentialsOut) GetOauthAccessTokenEndpoint() string`

GetOauthAccessTokenEndpoint returns the OauthAccessTokenEndpoint field if non-nil, zero value otherwise.

### GetOauthAccessTokenEndpointOk

`func (o *InformaticaV2CredentialsOut) GetOauthAccessTokenEndpointOk() (*string, bool)`

GetOauthAccessTokenEndpointOk returns a tuple with the OauthAccessTokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthAccessTokenEndpoint

`func (o *InformaticaV2CredentialsOut) SetOauthAccessTokenEndpoint(v string)`

SetOauthAccessTokenEndpoint sets OauthAccessTokenEndpoint field to given value.


### SetOauthAccessTokenEndpointNil

`func (o *InformaticaV2CredentialsOut) SetOauthAccessTokenEndpointNil(b bool)`

 SetOauthAccessTokenEndpointNil sets the value for OauthAccessTokenEndpoint to be an explicit nil

### UnsetOauthAccessTokenEndpoint
`func (o *InformaticaV2CredentialsOut) UnsetOauthAccessTokenEndpoint()`

UnsetOauthAccessTokenEndpoint ensures that no value is present for OauthAccessTokenEndpoint, not even an explicit nil
### GetOauthScope

`func (o *InformaticaV2CredentialsOut) GetOauthScope() string`

GetOauthScope returns the OauthScope field if non-nil, zero value otherwise.

### GetOauthScopeOk

`func (o *InformaticaV2CredentialsOut) GetOauthScopeOk() (*string, bool)`

GetOauthScopeOk returns a tuple with the OauthScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthScope

`func (o *InformaticaV2CredentialsOut) SetOauthScope(v string)`

SetOauthScope sets OauthScope field to given value.


### SetOauthScopeNil

`func (o *InformaticaV2CredentialsOut) SetOauthScopeNil(b bool)`

 SetOauthScopeNil sets the value for OauthScope to be an explicit nil

### UnsetOauthScope
`func (o *InformaticaV2CredentialsOut) UnsetOauthScope()`

UnsetOauthScope ensures that no value is present for OauthScope, not even an explicit nil
### GetOauthUsername

`func (o *InformaticaV2CredentialsOut) GetOauthUsername() string`

GetOauthUsername returns the OauthUsername field if non-nil, zero value otherwise.

### GetOauthUsernameOk

`func (o *InformaticaV2CredentialsOut) GetOauthUsernameOk() (*string, bool)`

GetOauthUsernameOk returns a tuple with the OauthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthUsername

`func (o *InformaticaV2CredentialsOut) SetOauthUsername(v string)`

SetOauthUsername sets OauthUsername field to given value.


### SetOauthUsernameNil

`func (o *InformaticaV2CredentialsOut) SetOauthUsernameNil(b bool)`

 SetOauthUsernameNil sets the value for OauthUsername to be an explicit nil

### UnsetOauthUsername
`func (o *InformaticaV2CredentialsOut) UnsetOauthUsername()`

UnsetOauthUsername ensures that no value is present for OauthUsername, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


