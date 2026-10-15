# PowerBiCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Username** | Pointer to **NullableString** | User Monte Carlo signs in as, for &#x60;primary_user&#x60;. | [optional] 
**AppClientSecret** | Pointer to **NullableString** | Secret of the app registration, for &#x60;service_principal&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**Password** | Pointer to **NullableString** | Password of &#x60;username&#x60;, for &#x60;primary_user&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**TenantId** | Pointer to **NullableString** | Microsoft Entra ID tenant the Power BI service belongs to. | [optional] 
**AppClientId** | Pointer to **NullableString** | Client ID of the Entra ID app registration Monte Carlo signs in with. | [optional] 
**AuthMode** | Pointer to [**NullablePowerBiAuthMode**](PowerBiAuthMode.md) | How Monte Carlo signs in. &#x60;service_principal&#x60; takes &#x60;app_client_secret&#x60;. &#x60;primary_user&#x60; takes &#x60;username&#x60; and &#x60;password&#x60;. | [optional] 

## Methods

### NewPowerBiCredentialsChangeValidateIn

`func NewPowerBiCredentialsChangeValidateIn(deploymentId string, ) *PowerBiCredentialsChangeValidateIn`

NewPowerBiCredentialsChangeValidateIn instantiates a new PowerBiCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerBiCredentialsChangeValidateInWithDefaults

`func NewPowerBiCredentialsChangeValidateInWithDefaults() *PowerBiCredentialsChangeValidateIn`

NewPowerBiCredentialsChangeValidateInWithDefaults instantiates a new PowerBiCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *PowerBiCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *PowerBiCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *PowerBiCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetUsername

`func (o *PowerBiCredentialsChangeValidateIn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PowerBiCredentialsChangeValidateIn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PowerBiCredentialsChangeValidateIn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PowerBiCredentialsChangeValidateIn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *PowerBiCredentialsChangeValidateIn) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *PowerBiCredentialsChangeValidateIn) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetAppClientSecret

`func (o *PowerBiCredentialsChangeValidateIn) GetAppClientSecret() string`

GetAppClientSecret returns the AppClientSecret field if non-nil, zero value otherwise.

### GetAppClientSecretOk

`func (o *PowerBiCredentialsChangeValidateIn) GetAppClientSecretOk() (*string, bool)`

GetAppClientSecretOk returns a tuple with the AppClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientSecret

`func (o *PowerBiCredentialsChangeValidateIn) SetAppClientSecret(v string)`

SetAppClientSecret sets AppClientSecret field to given value.

### HasAppClientSecret

`func (o *PowerBiCredentialsChangeValidateIn) HasAppClientSecret() bool`

HasAppClientSecret returns a boolean if a field has been set.

### SetAppClientSecretNil

`func (o *PowerBiCredentialsChangeValidateIn) SetAppClientSecretNil(b bool)`

 SetAppClientSecretNil sets the value for AppClientSecret to be an explicit nil

### UnsetAppClientSecret
`func (o *PowerBiCredentialsChangeValidateIn) UnsetAppClientSecret()`

UnsetAppClientSecret ensures that no value is present for AppClientSecret, not even an explicit nil
### GetPassword

`func (o *PowerBiCredentialsChangeValidateIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PowerBiCredentialsChangeValidateIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PowerBiCredentialsChangeValidateIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PowerBiCredentialsChangeValidateIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PowerBiCredentialsChangeValidateIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PowerBiCredentialsChangeValidateIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTenantId

`func (o *PowerBiCredentialsChangeValidateIn) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PowerBiCredentialsChangeValidateIn) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PowerBiCredentialsChangeValidateIn) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PowerBiCredentialsChangeValidateIn) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *PowerBiCredentialsChangeValidateIn) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *PowerBiCredentialsChangeValidateIn) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetAppClientId

`func (o *PowerBiCredentialsChangeValidateIn) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *PowerBiCredentialsChangeValidateIn) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *PowerBiCredentialsChangeValidateIn) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.

### HasAppClientId

`func (o *PowerBiCredentialsChangeValidateIn) HasAppClientId() bool`

HasAppClientId returns a boolean if a field has been set.

### SetAppClientIdNil

`func (o *PowerBiCredentialsChangeValidateIn) SetAppClientIdNil(b bool)`

 SetAppClientIdNil sets the value for AppClientId to be an explicit nil

### UnsetAppClientId
`func (o *PowerBiCredentialsChangeValidateIn) UnsetAppClientId()`

UnsetAppClientId ensures that no value is present for AppClientId, not even an explicit nil
### GetAuthMode

`func (o *PowerBiCredentialsChangeValidateIn) GetAuthMode() PowerBiAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *PowerBiCredentialsChangeValidateIn) GetAuthModeOk() (*PowerBiAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *PowerBiCredentialsChangeValidateIn) SetAuthMode(v PowerBiAuthMode)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *PowerBiCredentialsChangeValidateIn) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### SetAuthModeNil

`func (o *PowerBiCredentialsChangeValidateIn) SetAuthModeNil(b bool)`

 SetAuthModeNil sets the value for AuthMode to be an explicit nil

### UnsetAuthMode
`func (o *PowerBiCredentialsChangeValidateIn) UnsetAuthMode()`

UnsetAuthMode ensures that no value is present for AuthMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


