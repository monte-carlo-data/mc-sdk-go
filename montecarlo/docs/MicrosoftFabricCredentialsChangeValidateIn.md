# MicrosoftFabricCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Host** | Pointer to **NullableString** | Hostname of the Fabric warehouse or lakehouse SQL analytics endpoint, as Fabric shows it under SQL connection string, such as &#x60;abc123.datawarehouse.fabric.microsoft.com&#x60;. It goes unquoted into the connection string Monte Carlo builds, so it cannot contain whitespace or any of &#x60;;&#x60;, &#x60;{&#x60;, &#x60;}&#x60; and &#x60;&#x3D;&#x60;. Changing it requires sending the credentials&#39; secrets again. | [optional] 
**Port** | Pointer to **NullableInt32** | Port of the SQL endpoint. Fabric serves it on 1433. | [optional] 
**DbName** | Pointer to **NullableString** | Warehouse or lakehouse to connect to. It goes unquoted into the connection string Monte Carlo builds, so it cannot contain whitespace or any of &#x60;;&#x60;, &#x60;{&#x60;, &#x60;}&#x60; and &#x60;&#x3D;&#x60;. | [optional] 
**TenantId** | Pointer to **NullableString** | Microsoft Entra ID tenant the Fabric workspace belongs to. | [optional] 
**AppClientId** | Pointer to **NullableString** | Client ID of the Entra ID app registration Monte Carlo signs in with. | [optional] 
**AppClientSecret** | Pointer to **NullableString** | Secret of the app registration. Stored by Monte Carlo and never returned. | [optional] 

## Methods

### NewMicrosoftFabricCredentialsChangeValidateIn

`func NewMicrosoftFabricCredentialsChangeValidateIn(deploymentId string, ) *MicrosoftFabricCredentialsChangeValidateIn`

NewMicrosoftFabricCredentialsChangeValidateIn instantiates a new MicrosoftFabricCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftFabricCredentialsChangeValidateInWithDefaults

`func NewMicrosoftFabricCredentialsChangeValidateInWithDefaults() *MicrosoftFabricCredentialsChangeValidateIn`

NewMicrosoftFabricCredentialsChangeValidateInWithDefaults instantiates a new MicrosoftFabricCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetHost

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MicrosoftFabricCredentialsChangeValidateIn) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *MicrosoftFabricCredentialsChangeValidateIn) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MicrosoftFabricCredentialsChangeValidateIn) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *MicrosoftFabricCredentialsChangeValidateIn) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *MicrosoftFabricCredentialsChangeValidateIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *MicrosoftFabricCredentialsChangeValidateIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetTenantId

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MicrosoftFabricCredentialsChangeValidateIn) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *MicrosoftFabricCredentialsChangeValidateIn) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetAppClientId

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.

### HasAppClientId

`func (o *MicrosoftFabricCredentialsChangeValidateIn) HasAppClientId() bool`

HasAppClientId returns a boolean if a field has been set.

### SetAppClientIdNil

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetAppClientIdNil(b bool)`

 SetAppClientIdNil sets the value for AppClientId to be an explicit nil

### UnsetAppClientId
`func (o *MicrosoftFabricCredentialsChangeValidateIn) UnsetAppClientId()`

UnsetAppClientId ensures that no value is present for AppClientId, not even an explicit nil
### GetAppClientSecret

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetAppClientSecret() string`

GetAppClientSecret returns the AppClientSecret field if non-nil, zero value otherwise.

### GetAppClientSecretOk

`func (o *MicrosoftFabricCredentialsChangeValidateIn) GetAppClientSecretOk() (*string, bool)`

GetAppClientSecretOk returns a tuple with the AppClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientSecret

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetAppClientSecret(v string)`

SetAppClientSecret sets AppClientSecret field to given value.

### HasAppClientSecret

`func (o *MicrosoftFabricCredentialsChangeValidateIn) HasAppClientSecret() bool`

HasAppClientSecret returns a boolean if a field has been set.

### SetAppClientSecretNil

`func (o *MicrosoftFabricCredentialsChangeValidateIn) SetAppClientSecretNil(b bool)`

 SetAppClientSecretNil sets the value for AppClientSecret to be an explicit nil

### UnsetAppClientSecret
`func (o *MicrosoftFabricCredentialsChangeValidateIn) UnsetAppClientSecret()`

UnsetAppClientSecret ensures that no value is present for AppClientSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


