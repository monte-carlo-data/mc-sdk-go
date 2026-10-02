# PowerBiCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** | User Monte Carlo signs in as, for &#x60;primary_user&#x60;. | [optional] 
**AppClientSecret** | Pointer to **NullableString** | Secret of the app registration, for &#x60;service_principal&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**Password** | Pointer to **NullableString** | Password of &#x60;username&#x60;, for &#x60;primary_user&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**TenantId** | **string** | Microsoft Entra ID tenant the Power BI service belongs to. | 
**AppClientId** | **string** | Client ID of the Entra ID app registration Monte Carlo signs in with. | 
**AuthMode** | [**PowerBiAuthMode**](PowerBiAuthMode.md) | How Monte Carlo signs in. &#x60;service_principal&#x60; takes &#x60;app_client_secret&#x60;. &#x60;primary_user&#x60; takes &#x60;username&#x60; and &#x60;password&#x60;. | 

## Methods

### NewPowerBiCredentialsIn

`func NewPowerBiCredentialsIn(tenantId string, appClientId string, authMode PowerBiAuthMode, ) *PowerBiCredentialsIn`

NewPowerBiCredentialsIn instantiates a new PowerBiCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerBiCredentialsInWithDefaults

`func NewPowerBiCredentialsInWithDefaults() *PowerBiCredentialsIn`

NewPowerBiCredentialsInWithDefaults instantiates a new PowerBiCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *PowerBiCredentialsIn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PowerBiCredentialsIn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PowerBiCredentialsIn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PowerBiCredentialsIn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *PowerBiCredentialsIn) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *PowerBiCredentialsIn) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetAppClientSecret

`func (o *PowerBiCredentialsIn) GetAppClientSecret() string`

GetAppClientSecret returns the AppClientSecret field if non-nil, zero value otherwise.

### GetAppClientSecretOk

`func (o *PowerBiCredentialsIn) GetAppClientSecretOk() (*string, bool)`

GetAppClientSecretOk returns a tuple with the AppClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientSecret

`func (o *PowerBiCredentialsIn) SetAppClientSecret(v string)`

SetAppClientSecret sets AppClientSecret field to given value.

### HasAppClientSecret

`func (o *PowerBiCredentialsIn) HasAppClientSecret() bool`

HasAppClientSecret returns a boolean if a field has been set.

### SetAppClientSecretNil

`func (o *PowerBiCredentialsIn) SetAppClientSecretNil(b bool)`

 SetAppClientSecretNil sets the value for AppClientSecret to be an explicit nil

### UnsetAppClientSecret
`func (o *PowerBiCredentialsIn) UnsetAppClientSecret()`

UnsetAppClientSecret ensures that no value is present for AppClientSecret, not even an explicit nil
### GetPassword

`func (o *PowerBiCredentialsIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PowerBiCredentialsIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PowerBiCredentialsIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PowerBiCredentialsIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PowerBiCredentialsIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PowerBiCredentialsIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTenantId

`func (o *PowerBiCredentialsIn) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PowerBiCredentialsIn) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PowerBiCredentialsIn) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAppClientId

`func (o *PowerBiCredentialsIn) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *PowerBiCredentialsIn) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *PowerBiCredentialsIn) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.


### GetAuthMode

`func (o *PowerBiCredentialsIn) GetAuthMode() PowerBiAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *PowerBiCredentialsIn) GetAuthModeOk() (*PowerBiAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *PowerBiCredentialsIn) SetAuthMode(v PowerBiAuthMode)`

SetAuthMode sets AuthMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


