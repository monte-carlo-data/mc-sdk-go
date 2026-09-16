# EnvVarCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**BqProjectId** | **NullableString** | BigQuery project the connection reads from. Null unless set. | 
**DatabricksWarehouseId** | **NullableString** | Databricks SQL warehouse the connection runs queries on. Null unless set. | 
**EnvVarName** | **string** | Name of the environment variable on the deployment that holds the connection&#39;s credentials. Must start with &#x60;MCD_&#x60;. | 
**KmsKeyId** | **NullableString** | AWS KMS key the value is encrypted with. Null for a value in the clear. | 

## Methods

### NewEnvVarCredentialsOut

`func NewEnvVarCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, bqProjectId NullableString, databricksWarehouseId NullableString, envVarName string, kmsKeyId NullableString, ) *EnvVarCredentialsOut`

NewEnvVarCredentialsOut instantiates a new EnvVarCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvVarCredentialsOutWithDefaults

`func NewEnvVarCredentialsOutWithDefaults() *EnvVarCredentialsOut`

NewEnvVarCredentialsOutWithDefaults instantiates a new EnvVarCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvVarCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvVarCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvVarCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *EnvVarCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *EnvVarCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *EnvVarCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *EnvVarCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *EnvVarCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *EnvVarCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *EnvVarCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *EnvVarCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *EnvVarCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetBqProjectId

`func (o *EnvVarCredentialsOut) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *EnvVarCredentialsOut) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *EnvVarCredentialsOut) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.


### SetBqProjectIdNil

`func (o *EnvVarCredentialsOut) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *EnvVarCredentialsOut) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetDatabricksWarehouseId

`func (o *EnvVarCredentialsOut) GetDatabricksWarehouseId() string`

GetDatabricksWarehouseId returns the DatabricksWarehouseId field if non-nil, zero value otherwise.

### GetDatabricksWarehouseIdOk

`func (o *EnvVarCredentialsOut) GetDatabricksWarehouseIdOk() (*string, bool)`

GetDatabricksWarehouseIdOk returns a tuple with the DatabricksWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabricksWarehouseId

`func (o *EnvVarCredentialsOut) SetDatabricksWarehouseId(v string)`

SetDatabricksWarehouseId sets DatabricksWarehouseId field to given value.


### SetDatabricksWarehouseIdNil

`func (o *EnvVarCredentialsOut) SetDatabricksWarehouseIdNil(b bool)`

 SetDatabricksWarehouseIdNil sets the value for DatabricksWarehouseId to be an explicit nil

### UnsetDatabricksWarehouseId
`func (o *EnvVarCredentialsOut) UnsetDatabricksWarehouseId()`

UnsetDatabricksWarehouseId ensures that no value is present for DatabricksWarehouseId, not even an explicit nil
### GetEnvVarName

`func (o *EnvVarCredentialsOut) GetEnvVarName() string`

GetEnvVarName returns the EnvVarName field if non-nil, zero value otherwise.

### GetEnvVarNameOk

`func (o *EnvVarCredentialsOut) GetEnvVarNameOk() (*string, bool)`

GetEnvVarNameOk returns a tuple with the EnvVarName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVarName

`func (o *EnvVarCredentialsOut) SetEnvVarName(v string)`

SetEnvVarName sets EnvVarName field to given value.


### GetKmsKeyId

`func (o *EnvVarCredentialsOut) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *EnvVarCredentialsOut) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *EnvVarCredentialsOut) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.


### SetKmsKeyIdNil

`func (o *EnvVarCredentialsOut) SetKmsKeyIdNil(b bool)`

 SetKmsKeyIdNil sets the value for KmsKeyId to be an explicit nil

### UnsetKmsKeyId
`func (o *EnvVarCredentialsOut) UnsetKmsKeyId()`

UnsetKmsKeyId ensures that no value is present for KmsKeyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


