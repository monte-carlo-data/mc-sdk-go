# AzureKeyVaultCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**BqProjectId** | **NullableString** | BigQuery project the connection reads from. Null unless set. | 
**SqlWarehouseId** | **NullableString** | Databricks SQL warehouse the connection runs queries on. Null unless set. | 
**AkvSecret** | **string** | Name of the Azure Key Vault secret holding the connection&#39;s credentials. | 
**AkvVaultName** | **NullableString** | Name of the key vault. Null when unset. | 
**AkvVaultUrl** | **NullableString** | URL of the key vault. Null when unset. | 

## Methods

### NewAzureKeyVaultCredentialsOut

`func NewAzureKeyVaultCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, bqProjectId NullableString, sqlWarehouseId NullableString, akvSecret string, akvVaultName NullableString, akvVaultUrl NullableString, ) *AzureKeyVaultCredentialsOut`

NewAzureKeyVaultCredentialsOut instantiates a new AzureKeyVaultCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultCredentialsOutWithDefaults

`func NewAzureKeyVaultCredentialsOutWithDefaults() *AzureKeyVaultCredentialsOut`

NewAzureKeyVaultCredentialsOutWithDefaults instantiates a new AzureKeyVaultCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureKeyVaultCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureKeyVaultCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureKeyVaultCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *AzureKeyVaultCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *AzureKeyVaultCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *AzureKeyVaultCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *AzureKeyVaultCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *AzureKeyVaultCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *AzureKeyVaultCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *AzureKeyVaultCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AzureKeyVaultCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AzureKeyVaultCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetBqProjectId

`func (o *AzureKeyVaultCredentialsOut) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *AzureKeyVaultCredentialsOut) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *AzureKeyVaultCredentialsOut) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.


### SetBqProjectIdNil

`func (o *AzureKeyVaultCredentialsOut) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *AzureKeyVaultCredentialsOut) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetSqlWarehouseId

`func (o *AzureKeyVaultCredentialsOut) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *AzureKeyVaultCredentialsOut) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *AzureKeyVaultCredentialsOut) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.


### SetSqlWarehouseIdNil

`func (o *AzureKeyVaultCredentialsOut) SetSqlWarehouseIdNil(b bool)`

 SetSqlWarehouseIdNil sets the value for SqlWarehouseId to be an explicit nil

### UnsetSqlWarehouseId
`func (o *AzureKeyVaultCredentialsOut) UnsetSqlWarehouseId()`

UnsetSqlWarehouseId ensures that no value is present for SqlWarehouseId, not even an explicit nil
### GetAkvSecret

`func (o *AzureKeyVaultCredentialsOut) GetAkvSecret() string`

GetAkvSecret returns the AkvSecret field if non-nil, zero value otherwise.

### GetAkvSecretOk

`func (o *AzureKeyVaultCredentialsOut) GetAkvSecretOk() (*string, bool)`

GetAkvSecretOk returns a tuple with the AkvSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkvSecret

`func (o *AzureKeyVaultCredentialsOut) SetAkvSecret(v string)`

SetAkvSecret sets AkvSecret field to given value.


### GetAkvVaultName

`func (o *AzureKeyVaultCredentialsOut) GetAkvVaultName() string`

GetAkvVaultName returns the AkvVaultName field if non-nil, zero value otherwise.

### GetAkvVaultNameOk

`func (o *AzureKeyVaultCredentialsOut) GetAkvVaultNameOk() (*string, bool)`

GetAkvVaultNameOk returns a tuple with the AkvVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkvVaultName

`func (o *AzureKeyVaultCredentialsOut) SetAkvVaultName(v string)`

SetAkvVaultName sets AkvVaultName field to given value.


### SetAkvVaultNameNil

`func (o *AzureKeyVaultCredentialsOut) SetAkvVaultNameNil(b bool)`

 SetAkvVaultNameNil sets the value for AkvVaultName to be an explicit nil

### UnsetAkvVaultName
`func (o *AzureKeyVaultCredentialsOut) UnsetAkvVaultName()`

UnsetAkvVaultName ensures that no value is present for AkvVaultName, not even an explicit nil
### GetAkvVaultUrl

`func (o *AzureKeyVaultCredentialsOut) GetAkvVaultUrl() string`

GetAkvVaultUrl returns the AkvVaultUrl field if non-nil, zero value otherwise.

### GetAkvVaultUrlOk

`func (o *AzureKeyVaultCredentialsOut) GetAkvVaultUrlOk() (*string, bool)`

GetAkvVaultUrlOk returns a tuple with the AkvVaultUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkvVaultUrl

`func (o *AzureKeyVaultCredentialsOut) SetAkvVaultUrl(v string)`

SetAkvVaultUrl sets AkvVaultUrl field to given value.


### SetAkvVaultUrlNil

`func (o *AzureKeyVaultCredentialsOut) SetAkvVaultUrlNil(b bool)`

 SetAkvVaultUrlNil sets the value for AkvVaultUrl to be an explicit nil

### UnsetAkvVaultUrl
`func (o *AzureKeyVaultCredentialsOut) UnsetAkvVaultUrl()`

UnsetAkvVaultUrl ensures that no value is present for AkvVaultUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


