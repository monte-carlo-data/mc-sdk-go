# PowerBiCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**TenantId** | **string** | Microsoft Entra ID tenant the Power BI service belongs to. | 
**AppClientId** | **string** | Client ID of the Entra ID app registration Monte Carlo signs in with. | 
**AuthMode** | [**PowerBiAuthMode**](PowerBiAuthMode.md) | How Monte Carlo signs in. &#x60;service_principal&#x60; takes &#x60;app_client_secret&#x60;. &#x60;primary_user&#x60; takes &#x60;username&#x60; and &#x60;password&#x60;. | 
**Username** | **NullableString** | User Monte Carlo signs in as. Null unless set. | 

## Methods

### NewPowerBiCredentialsOut

`func NewPowerBiCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, tenantId string, appClientId string, authMode PowerBiAuthMode, username NullableString, ) *PowerBiCredentialsOut`

NewPowerBiCredentialsOut instantiates a new PowerBiCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerBiCredentialsOutWithDefaults

`func NewPowerBiCredentialsOutWithDefaults() *PowerBiCredentialsOut`

NewPowerBiCredentialsOutWithDefaults instantiates a new PowerBiCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerBiCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerBiCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerBiCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *PowerBiCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *PowerBiCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *PowerBiCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *PowerBiCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *PowerBiCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *PowerBiCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *PowerBiCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PowerBiCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PowerBiCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetTenantId

`func (o *PowerBiCredentialsOut) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PowerBiCredentialsOut) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PowerBiCredentialsOut) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAppClientId

`func (o *PowerBiCredentialsOut) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *PowerBiCredentialsOut) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *PowerBiCredentialsOut) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.


### GetAuthMode

`func (o *PowerBiCredentialsOut) GetAuthMode() PowerBiAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *PowerBiCredentialsOut) GetAuthModeOk() (*PowerBiAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *PowerBiCredentialsOut) SetAuthMode(v PowerBiAuthMode)`

SetAuthMode sets AuthMode field to given value.


### GetUsername

`func (o *PowerBiCredentialsOut) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PowerBiCredentialsOut) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PowerBiCredentialsOut) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *PowerBiCredentialsOut) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *PowerBiCredentialsOut) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


