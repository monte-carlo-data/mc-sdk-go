# TableauCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**ServerName** | **string** | URL of the Tableau server, starting with https:// or http://. | 
**SiteName** | **NullableString** | Tableau site. Null for the default site. | 
**VerifySsl** | **NullableBool** | Whether to verify the server&#39;s TLS certificate. Verified when left out. Null unless set. | 
**Username** | **NullableString** | Tableau user. Null for a personal access token. | 
**TokenName** | **NullableString** | Name of the personal access token. Null unless the credentials use one. | 
**ConnectedAppClientId** | **NullableString** | Client ID of the connected app. Null unless the credentials use one. | 
**ConnectedAppSecretId** | **NullableString** | ID of the connected app&#39;s secret. Null unless the credentials use one. | 

## Methods

### NewTableauCredentialsOut

`func NewTableauCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, serverName string, siteName NullableString, verifySsl NullableBool, username NullableString, tokenName NullableString, connectedAppClientId NullableString, connectedAppSecretId NullableString, ) *TableauCredentialsOut`

NewTableauCredentialsOut instantiates a new TableauCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableauCredentialsOutWithDefaults

`func NewTableauCredentialsOutWithDefaults() *TableauCredentialsOut`

NewTableauCredentialsOutWithDefaults instantiates a new TableauCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TableauCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TableauCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TableauCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *TableauCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *TableauCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *TableauCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *TableauCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *TableauCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *TableauCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *TableauCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *TableauCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *TableauCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetServerName

`func (o *TableauCredentialsOut) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *TableauCredentialsOut) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *TableauCredentialsOut) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetSiteName

`func (o *TableauCredentialsOut) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *TableauCredentialsOut) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *TableauCredentialsOut) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.


### SetSiteNameNil

`func (o *TableauCredentialsOut) SetSiteNameNil(b bool)`

 SetSiteNameNil sets the value for SiteName to be an explicit nil

### UnsetSiteName
`func (o *TableauCredentialsOut) UnsetSiteName()`

UnsetSiteName ensures that no value is present for SiteName, not even an explicit nil
### GetVerifySsl

`func (o *TableauCredentialsOut) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *TableauCredentialsOut) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *TableauCredentialsOut) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.


### SetVerifySslNil

`func (o *TableauCredentialsOut) SetVerifySslNil(b bool)`

 SetVerifySslNil sets the value for VerifySsl to be an explicit nil

### UnsetVerifySsl
`func (o *TableauCredentialsOut) UnsetVerifySsl()`

UnsetVerifySsl ensures that no value is present for VerifySsl, not even an explicit nil
### GetUsername

`func (o *TableauCredentialsOut) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TableauCredentialsOut) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TableauCredentialsOut) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *TableauCredentialsOut) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TableauCredentialsOut) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetTokenName

`func (o *TableauCredentialsOut) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *TableauCredentialsOut) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *TableauCredentialsOut) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### SetTokenNameNil

`func (o *TableauCredentialsOut) SetTokenNameNil(b bool)`

 SetTokenNameNil sets the value for TokenName to be an explicit nil

### UnsetTokenName
`func (o *TableauCredentialsOut) UnsetTokenName()`

UnsetTokenName ensures that no value is present for TokenName, not even an explicit nil
### GetConnectedAppClientId

`func (o *TableauCredentialsOut) GetConnectedAppClientId() string`

GetConnectedAppClientId returns the ConnectedAppClientId field if non-nil, zero value otherwise.

### GetConnectedAppClientIdOk

`func (o *TableauCredentialsOut) GetConnectedAppClientIdOk() (*string, bool)`

GetConnectedAppClientIdOk returns a tuple with the ConnectedAppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppClientId

`func (o *TableauCredentialsOut) SetConnectedAppClientId(v string)`

SetConnectedAppClientId sets ConnectedAppClientId field to given value.


### SetConnectedAppClientIdNil

`func (o *TableauCredentialsOut) SetConnectedAppClientIdNil(b bool)`

 SetConnectedAppClientIdNil sets the value for ConnectedAppClientId to be an explicit nil

### UnsetConnectedAppClientId
`func (o *TableauCredentialsOut) UnsetConnectedAppClientId()`

UnsetConnectedAppClientId ensures that no value is present for ConnectedAppClientId, not even an explicit nil
### GetConnectedAppSecretId

`func (o *TableauCredentialsOut) GetConnectedAppSecretId() string`

GetConnectedAppSecretId returns the ConnectedAppSecretId field if non-nil, zero value otherwise.

### GetConnectedAppSecretIdOk

`func (o *TableauCredentialsOut) GetConnectedAppSecretIdOk() (*string, bool)`

GetConnectedAppSecretIdOk returns a tuple with the ConnectedAppSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAppSecretId

`func (o *TableauCredentialsOut) SetConnectedAppSecretId(v string)`

SetConnectedAppSecretId sets ConnectedAppSecretId field to given value.


### SetConnectedAppSecretIdNil

`func (o *TableauCredentialsOut) SetConnectedAppSecretIdNil(b bool)`

 SetConnectedAppSecretIdNil sets the value for ConnectedAppSecretId to be an explicit nil

### UnsetConnectedAppSecretId
`func (o *TableauCredentialsOut) UnsetConnectedAppSecretId()`

UnsetConnectedAppSecretId ensures that no value is present for ConnectedAppSecretId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


