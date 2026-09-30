# GcpSecretManagerCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60; or &#x60;bigquery&#x60;, or one of your custom connector types. Fixed once created. | 
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**SqlWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**GcpSecret** | **string** | Name of the GCP Secret Manager secret holding the connection&#39;s credentials. | 

## Methods

### NewGcpSecretManagerCredentialsIn

`func NewGcpSecretManagerCredentialsIn(connectionType string, gcpSecret string, ) *GcpSecretManagerCredentialsIn`

NewGcpSecretManagerCredentialsIn instantiates a new GcpSecretManagerCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpSecretManagerCredentialsInWithDefaults

`func NewGcpSecretManagerCredentialsInWithDefaults() *GcpSecretManagerCredentialsIn`

NewGcpSecretManagerCredentialsInWithDefaults instantiates a new GcpSecretManagerCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *GcpSecretManagerCredentialsIn) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GcpSecretManagerCredentialsIn) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GcpSecretManagerCredentialsIn) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetBqProjectId

`func (o *GcpSecretManagerCredentialsIn) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *GcpSecretManagerCredentialsIn) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *GcpSecretManagerCredentialsIn) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *GcpSecretManagerCredentialsIn) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *GcpSecretManagerCredentialsIn) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *GcpSecretManagerCredentialsIn) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetSqlWarehouseId

`func (o *GcpSecretManagerCredentialsIn) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *GcpSecretManagerCredentialsIn) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *GcpSecretManagerCredentialsIn) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.

### HasSqlWarehouseId

`func (o *GcpSecretManagerCredentialsIn) HasSqlWarehouseId() bool`

HasSqlWarehouseId returns a boolean if a field has been set.

### SetSqlWarehouseIdNil

`func (o *GcpSecretManagerCredentialsIn) SetSqlWarehouseIdNil(b bool)`

 SetSqlWarehouseIdNil sets the value for SqlWarehouseId to be an explicit nil

### UnsetSqlWarehouseId
`func (o *GcpSecretManagerCredentialsIn) UnsetSqlWarehouseId()`

UnsetSqlWarehouseId ensures that no value is present for SqlWarehouseId, not even an explicit nil
### GetGcpSecret

`func (o *GcpSecretManagerCredentialsIn) GetGcpSecret() string`

GetGcpSecret returns the GcpSecret field if non-nil, zero value otherwise.

### GetGcpSecretOk

`func (o *GcpSecretManagerCredentialsIn) GetGcpSecretOk() (*string, bool)`

GetGcpSecretOk returns a tuple with the GcpSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpSecret

`func (o *GcpSecretManagerCredentialsIn) SetGcpSecret(v string)`

SetGcpSecret sets GcpSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


