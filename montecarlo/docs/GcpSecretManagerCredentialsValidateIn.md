# GcpSecretManagerCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**ConnectionType** | **string** | What the credentials are for, hyphenated, such as &#x60;snowflake&#x60; or &#x60;bigquery&#x60;. Decides which checks run. | 
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**DatabricksWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**GcpSecret** | **string** | Name of the GCP Secret Manager secret holding the connection&#39;s credentials. | 

## Methods

### NewGcpSecretManagerCredentialsValidateIn

`func NewGcpSecretManagerCredentialsValidateIn(deploymentId string, connectionType string, gcpSecret string, ) *GcpSecretManagerCredentialsValidateIn`

NewGcpSecretManagerCredentialsValidateIn instantiates a new GcpSecretManagerCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpSecretManagerCredentialsValidateInWithDefaults

`func NewGcpSecretManagerCredentialsValidateInWithDefaults() *GcpSecretManagerCredentialsValidateIn`

NewGcpSecretManagerCredentialsValidateInWithDefaults instantiates a new GcpSecretManagerCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *GcpSecretManagerCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GcpSecretManagerCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GcpSecretManagerCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetConnectionType

`func (o *GcpSecretManagerCredentialsValidateIn) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GcpSecretManagerCredentialsValidateIn) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GcpSecretManagerCredentialsValidateIn) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetBqProjectId

`func (o *GcpSecretManagerCredentialsValidateIn) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *GcpSecretManagerCredentialsValidateIn) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *GcpSecretManagerCredentialsValidateIn) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *GcpSecretManagerCredentialsValidateIn) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *GcpSecretManagerCredentialsValidateIn) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *GcpSecretManagerCredentialsValidateIn) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetDatabricksWarehouseId

`func (o *GcpSecretManagerCredentialsValidateIn) GetDatabricksWarehouseId() string`

GetDatabricksWarehouseId returns the DatabricksWarehouseId field if non-nil, zero value otherwise.

### GetDatabricksWarehouseIdOk

`func (o *GcpSecretManagerCredentialsValidateIn) GetDatabricksWarehouseIdOk() (*string, bool)`

GetDatabricksWarehouseIdOk returns a tuple with the DatabricksWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabricksWarehouseId

`func (o *GcpSecretManagerCredentialsValidateIn) SetDatabricksWarehouseId(v string)`

SetDatabricksWarehouseId sets DatabricksWarehouseId field to given value.

### HasDatabricksWarehouseId

`func (o *GcpSecretManagerCredentialsValidateIn) HasDatabricksWarehouseId() bool`

HasDatabricksWarehouseId returns a boolean if a field has been set.

### SetDatabricksWarehouseIdNil

`func (o *GcpSecretManagerCredentialsValidateIn) SetDatabricksWarehouseIdNil(b bool)`

 SetDatabricksWarehouseIdNil sets the value for DatabricksWarehouseId to be an explicit nil

### UnsetDatabricksWarehouseId
`func (o *GcpSecretManagerCredentialsValidateIn) UnsetDatabricksWarehouseId()`

UnsetDatabricksWarehouseId ensures that no value is present for DatabricksWarehouseId, not even an explicit nil
### GetGcpSecret

`func (o *GcpSecretManagerCredentialsValidateIn) GetGcpSecret() string`

GetGcpSecret returns the GcpSecret field if non-nil, zero value otherwise.

### GetGcpSecretOk

`func (o *GcpSecretManagerCredentialsValidateIn) GetGcpSecretOk() (*string, bool)`

GetGcpSecretOk returns a tuple with the GcpSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpSecret

`func (o *GcpSecretManagerCredentialsValidateIn) SetGcpSecret(v string)`

SetGcpSecret sets GcpSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


