# GcpSecretManagerCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**BqProjectId** | **NullableString** | BigQuery project the connection reads from. Null unless set. | 
**DatabricksWarehouseId** | **NullableString** | Databricks SQL warehouse the connection runs queries on. Null unless set. | 
**GcpSecret** | **string** | Name of the GCP Secret Manager secret holding the connection&#39;s credentials. | 

## Methods

### NewGcpSecretManagerCredentialsOut

`func NewGcpSecretManagerCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, bqProjectId NullableString, databricksWarehouseId NullableString, gcpSecret string, ) *GcpSecretManagerCredentialsOut`

NewGcpSecretManagerCredentialsOut instantiates a new GcpSecretManagerCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpSecretManagerCredentialsOutWithDefaults

`func NewGcpSecretManagerCredentialsOutWithDefaults() *GcpSecretManagerCredentialsOut`

NewGcpSecretManagerCredentialsOutWithDefaults instantiates a new GcpSecretManagerCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GcpSecretManagerCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GcpSecretManagerCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GcpSecretManagerCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *GcpSecretManagerCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GcpSecretManagerCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GcpSecretManagerCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *GcpSecretManagerCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *GcpSecretManagerCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *GcpSecretManagerCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *GcpSecretManagerCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GcpSecretManagerCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GcpSecretManagerCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetBqProjectId

`func (o *GcpSecretManagerCredentialsOut) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *GcpSecretManagerCredentialsOut) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *GcpSecretManagerCredentialsOut) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.


### SetBqProjectIdNil

`func (o *GcpSecretManagerCredentialsOut) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *GcpSecretManagerCredentialsOut) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetDatabricksWarehouseId

`func (o *GcpSecretManagerCredentialsOut) GetDatabricksWarehouseId() string`

GetDatabricksWarehouseId returns the DatabricksWarehouseId field if non-nil, zero value otherwise.

### GetDatabricksWarehouseIdOk

`func (o *GcpSecretManagerCredentialsOut) GetDatabricksWarehouseIdOk() (*string, bool)`

GetDatabricksWarehouseIdOk returns a tuple with the DatabricksWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabricksWarehouseId

`func (o *GcpSecretManagerCredentialsOut) SetDatabricksWarehouseId(v string)`

SetDatabricksWarehouseId sets DatabricksWarehouseId field to given value.


### SetDatabricksWarehouseIdNil

`func (o *GcpSecretManagerCredentialsOut) SetDatabricksWarehouseIdNil(b bool)`

 SetDatabricksWarehouseIdNil sets the value for DatabricksWarehouseId to be an explicit nil

### UnsetDatabricksWarehouseId
`func (o *GcpSecretManagerCredentialsOut) UnsetDatabricksWarehouseId()`

UnsetDatabricksWarehouseId ensures that no value is present for DatabricksWarehouseId, not even an explicit nil
### GetGcpSecret

`func (o *GcpSecretManagerCredentialsOut) GetGcpSecret() string`

GetGcpSecret returns the GcpSecret field if non-nil, zero value otherwise.

### GetGcpSecretOk

`func (o *GcpSecretManagerCredentialsOut) GetGcpSecretOk() (*string, bool)`

GetGcpSecretOk returns a tuple with the GcpSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpSecret

`func (o *GcpSecretManagerCredentialsOut) SetGcpSecret(v string)`

SetGcpSecret sets GcpSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


