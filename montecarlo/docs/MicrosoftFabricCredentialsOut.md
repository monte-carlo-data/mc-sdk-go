# MicrosoftFabricCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**Host** | **string** | Hostname of the Fabric warehouse or lakehouse SQL analytics endpoint, as Fabric shows it under SQL connection string, such as &#x60;abc123.datawarehouse.fabric.microsoft.com&#x60;. | 
**Port** | **int32** | Port of the SQL endpoint. Fabric serves it on 1433. | 
**DbName** | **string** | Warehouse or lakehouse to connect to. | 
**TenantId** | **string** | Microsoft Entra ID tenant the Fabric workspace belongs to. | 
**AppClientId** | **string** | Client ID of the Entra ID app registration Monte Carlo signs in with. | 

## Methods

### NewMicrosoftFabricCredentialsOut

`func NewMicrosoftFabricCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, host string, port int32, dbName string, tenantId string, appClientId string, ) *MicrosoftFabricCredentialsOut`

NewMicrosoftFabricCredentialsOut instantiates a new MicrosoftFabricCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftFabricCredentialsOutWithDefaults

`func NewMicrosoftFabricCredentialsOutWithDefaults() *MicrosoftFabricCredentialsOut`

NewMicrosoftFabricCredentialsOutWithDefaults instantiates a new MicrosoftFabricCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftFabricCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftFabricCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftFabricCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *MicrosoftFabricCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *MicrosoftFabricCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *MicrosoftFabricCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *MicrosoftFabricCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *MicrosoftFabricCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *MicrosoftFabricCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *MicrosoftFabricCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *MicrosoftFabricCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *MicrosoftFabricCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetHost

`func (o *MicrosoftFabricCredentialsOut) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MicrosoftFabricCredentialsOut) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MicrosoftFabricCredentialsOut) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *MicrosoftFabricCredentialsOut) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MicrosoftFabricCredentialsOut) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MicrosoftFabricCredentialsOut) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDbName

`func (o *MicrosoftFabricCredentialsOut) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MicrosoftFabricCredentialsOut) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MicrosoftFabricCredentialsOut) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetTenantId

`func (o *MicrosoftFabricCredentialsOut) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MicrosoftFabricCredentialsOut) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MicrosoftFabricCredentialsOut) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAppClientId

`func (o *MicrosoftFabricCredentialsOut) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *MicrosoftFabricCredentialsOut) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *MicrosoftFabricCredentialsOut) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


