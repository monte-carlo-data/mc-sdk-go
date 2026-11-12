# MicrosoftFabricCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname of the Fabric warehouse or lakehouse SQL analytics endpoint, as Fabric shows it under SQL connection string, such as &#x60;abc123.datawarehouse.fabric.microsoft.com&#x60;. It goes unquoted into the connection string Monte Carlo builds, so it cannot contain whitespace or any of &#x60;;&#x60;, &#x60;{&#x60;, &#x60;}&#x60; and &#x60;&#x3D;&#x60;. | 
**Port** | **int32** | Port of the SQL endpoint. Fabric serves it on 1433. | 
**DbName** | **string** | Warehouse or lakehouse to connect to. It goes unquoted into the connection string Monte Carlo builds, so it cannot contain whitespace or any of &#x60;;&#x60;, &#x60;{&#x60;, &#x60;}&#x60; and &#x60;&#x3D;&#x60;. | 
**TenantId** | **string** | Microsoft Entra ID tenant the Fabric workspace belongs to. | 
**AppClientId** | **string** | Client ID of the Entra ID app registration Monte Carlo signs in with. | 
**AppClientSecret** | **string** | Secret of the app registration. Stored by Monte Carlo and never returned. | 

## Methods

### NewMicrosoftFabricCredentialsIn

`func NewMicrosoftFabricCredentialsIn(host string, port int32, dbName string, tenantId string, appClientId string, appClientSecret string, ) *MicrosoftFabricCredentialsIn`

NewMicrosoftFabricCredentialsIn instantiates a new MicrosoftFabricCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftFabricCredentialsInWithDefaults

`func NewMicrosoftFabricCredentialsInWithDefaults() *MicrosoftFabricCredentialsIn`

NewMicrosoftFabricCredentialsInWithDefaults instantiates a new MicrosoftFabricCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *MicrosoftFabricCredentialsIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MicrosoftFabricCredentialsIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MicrosoftFabricCredentialsIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *MicrosoftFabricCredentialsIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MicrosoftFabricCredentialsIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MicrosoftFabricCredentialsIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDbName

`func (o *MicrosoftFabricCredentialsIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MicrosoftFabricCredentialsIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MicrosoftFabricCredentialsIn) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetTenantId

`func (o *MicrosoftFabricCredentialsIn) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MicrosoftFabricCredentialsIn) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MicrosoftFabricCredentialsIn) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAppClientId

`func (o *MicrosoftFabricCredentialsIn) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *MicrosoftFabricCredentialsIn) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *MicrosoftFabricCredentialsIn) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.


### GetAppClientSecret

`func (o *MicrosoftFabricCredentialsIn) GetAppClientSecret() string`

GetAppClientSecret returns the AppClientSecret field if non-nil, zero value otherwise.

### GetAppClientSecretOk

`func (o *MicrosoftFabricCredentialsIn) GetAppClientSecretOk() (*string, bool)`

GetAppClientSecretOk returns a tuple with the AppClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientSecret

`func (o *MicrosoftFabricCredentialsIn) SetAppClientSecret(v string)`

SetAppClientSecret sets AppClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


