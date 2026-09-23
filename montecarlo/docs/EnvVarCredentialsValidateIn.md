# EnvVarCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**ConnectionType** | **string** | What the credentials are for, hyphenated, such as &#x60;snowflake&#x60; or &#x60;bigquery&#x60;. Decides which checks run. | 
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**DatabricksWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**EnvVarName** | **string** | Name of the environment variable on the deployment that holds the connection&#39;s credentials. Must start with &#x60;MCD_&#x60;. | 
**KmsKeyId** | Pointer to **NullableString** | AWS KMS key the variable&#39;s value is encrypted with. Omit it for a value stored in the clear. | [optional] 

## Methods

### NewEnvVarCredentialsValidateIn

`func NewEnvVarCredentialsValidateIn(deploymentId string, connectionType string, envVarName string, ) *EnvVarCredentialsValidateIn`

NewEnvVarCredentialsValidateIn instantiates a new EnvVarCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvVarCredentialsValidateInWithDefaults

`func NewEnvVarCredentialsValidateInWithDefaults() *EnvVarCredentialsValidateIn`

NewEnvVarCredentialsValidateInWithDefaults instantiates a new EnvVarCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *EnvVarCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *EnvVarCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *EnvVarCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetConnectionType

`func (o *EnvVarCredentialsValidateIn) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *EnvVarCredentialsValidateIn) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *EnvVarCredentialsValidateIn) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetBqProjectId

`func (o *EnvVarCredentialsValidateIn) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *EnvVarCredentialsValidateIn) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *EnvVarCredentialsValidateIn) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *EnvVarCredentialsValidateIn) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *EnvVarCredentialsValidateIn) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *EnvVarCredentialsValidateIn) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetDatabricksWarehouseId

`func (o *EnvVarCredentialsValidateIn) GetDatabricksWarehouseId() string`

GetDatabricksWarehouseId returns the DatabricksWarehouseId field if non-nil, zero value otherwise.

### GetDatabricksWarehouseIdOk

`func (o *EnvVarCredentialsValidateIn) GetDatabricksWarehouseIdOk() (*string, bool)`

GetDatabricksWarehouseIdOk returns a tuple with the DatabricksWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabricksWarehouseId

`func (o *EnvVarCredentialsValidateIn) SetDatabricksWarehouseId(v string)`

SetDatabricksWarehouseId sets DatabricksWarehouseId field to given value.

### HasDatabricksWarehouseId

`func (o *EnvVarCredentialsValidateIn) HasDatabricksWarehouseId() bool`

HasDatabricksWarehouseId returns a boolean if a field has been set.

### SetDatabricksWarehouseIdNil

`func (o *EnvVarCredentialsValidateIn) SetDatabricksWarehouseIdNil(b bool)`

 SetDatabricksWarehouseIdNil sets the value for DatabricksWarehouseId to be an explicit nil

### UnsetDatabricksWarehouseId
`func (o *EnvVarCredentialsValidateIn) UnsetDatabricksWarehouseId()`

UnsetDatabricksWarehouseId ensures that no value is present for DatabricksWarehouseId, not even an explicit nil
### GetEnvVarName

`func (o *EnvVarCredentialsValidateIn) GetEnvVarName() string`

GetEnvVarName returns the EnvVarName field if non-nil, zero value otherwise.

### GetEnvVarNameOk

`func (o *EnvVarCredentialsValidateIn) GetEnvVarNameOk() (*string, bool)`

GetEnvVarNameOk returns a tuple with the EnvVarName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVarName

`func (o *EnvVarCredentialsValidateIn) SetEnvVarName(v string)`

SetEnvVarName sets EnvVarName field to given value.


### GetKmsKeyId

`func (o *EnvVarCredentialsValidateIn) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *EnvVarCredentialsValidateIn) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *EnvVarCredentialsValidateIn) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *EnvVarCredentialsValidateIn) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.

### SetKmsKeyIdNil

`func (o *EnvVarCredentialsValidateIn) SetKmsKeyIdNil(b bool)`

 SetKmsKeyIdNil sets the value for KmsKeyId to be an explicit nil

### UnsetKmsKeyId
`func (o *EnvVarCredentialsValidateIn) UnsetKmsKeyId()`

UnsetKmsKeyId ensures that no value is present for KmsKeyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


