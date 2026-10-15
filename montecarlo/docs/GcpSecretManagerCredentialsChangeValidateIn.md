# GcpSecretManagerCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**SqlWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**GcpSecret** | Pointer to **NullableString** | Name of the GCP Secret Manager secret holding the connection&#39;s credentials. | [optional] 

## Methods

### NewGcpSecretManagerCredentialsChangeValidateIn

`func NewGcpSecretManagerCredentialsChangeValidateIn(deploymentId string, ) *GcpSecretManagerCredentialsChangeValidateIn`

NewGcpSecretManagerCredentialsChangeValidateIn instantiates a new GcpSecretManagerCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpSecretManagerCredentialsChangeValidateInWithDefaults

`func NewGcpSecretManagerCredentialsChangeValidateInWithDefaults() *GcpSecretManagerCredentialsChangeValidateIn`

NewGcpSecretManagerCredentialsChangeValidateInWithDefaults instantiates a new GcpSecretManagerCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *GcpSecretManagerCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GcpSecretManagerCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GcpSecretManagerCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetBqProjectId

`func (o *GcpSecretManagerCredentialsChangeValidateIn) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *GcpSecretManagerCredentialsChangeValidateIn) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *GcpSecretManagerCredentialsChangeValidateIn) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *GcpSecretManagerCredentialsChangeValidateIn) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *GcpSecretManagerCredentialsChangeValidateIn) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *GcpSecretManagerCredentialsChangeValidateIn) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetSqlWarehouseId

`func (o *GcpSecretManagerCredentialsChangeValidateIn) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *GcpSecretManagerCredentialsChangeValidateIn) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *GcpSecretManagerCredentialsChangeValidateIn) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.

### HasSqlWarehouseId

`func (o *GcpSecretManagerCredentialsChangeValidateIn) HasSqlWarehouseId() bool`

HasSqlWarehouseId returns a boolean if a field has been set.

### SetSqlWarehouseIdNil

`func (o *GcpSecretManagerCredentialsChangeValidateIn) SetSqlWarehouseIdNil(b bool)`

 SetSqlWarehouseIdNil sets the value for SqlWarehouseId to be an explicit nil

### UnsetSqlWarehouseId
`func (o *GcpSecretManagerCredentialsChangeValidateIn) UnsetSqlWarehouseId()`

UnsetSqlWarehouseId ensures that no value is present for SqlWarehouseId, not even an explicit nil
### GetGcpSecret

`func (o *GcpSecretManagerCredentialsChangeValidateIn) GetGcpSecret() string`

GetGcpSecret returns the GcpSecret field if non-nil, zero value otherwise.

### GetGcpSecretOk

`func (o *GcpSecretManagerCredentialsChangeValidateIn) GetGcpSecretOk() (*string, bool)`

GetGcpSecretOk returns a tuple with the GcpSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpSecret

`func (o *GcpSecretManagerCredentialsChangeValidateIn) SetGcpSecret(v string)`

SetGcpSecret sets GcpSecret field to given value.

### HasGcpSecret

`func (o *GcpSecretManagerCredentialsChangeValidateIn) HasGcpSecret() bool`

HasGcpSecret returns a boolean if a field has been set.

### SetGcpSecretNil

`func (o *GcpSecretManagerCredentialsChangeValidateIn) SetGcpSecretNil(b bool)`

 SetGcpSecretNil sets the value for GcpSecret to be an explicit nil

### UnsetGcpSecret
`func (o *GcpSecretManagerCredentialsChangeValidateIn) UnsetGcpSecret()`

UnsetGcpSecret ensures that no value is present for GcpSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


