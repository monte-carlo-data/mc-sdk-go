# EnvVarCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**SqlWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**EnvVarName** | Pointer to **NullableString** | Name of the environment variable on the deployment that holds the connection&#39;s credentials. Must start with &#x60;MCD_&#x60;. | [optional] 
**KmsKeyId** | Pointer to **NullableString** | AWS KMS key the variable&#39;s value is encrypted with. Omit it for a value stored in the clear. | [optional] 

## Methods

### NewEnvVarCredentialsChangeValidateIn

`func NewEnvVarCredentialsChangeValidateIn(deploymentId string, ) *EnvVarCredentialsChangeValidateIn`

NewEnvVarCredentialsChangeValidateIn instantiates a new EnvVarCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvVarCredentialsChangeValidateInWithDefaults

`func NewEnvVarCredentialsChangeValidateInWithDefaults() *EnvVarCredentialsChangeValidateIn`

NewEnvVarCredentialsChangeValidateInWithDefaults instantiates a new EnvVarCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *EnvVarCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *EnvVarCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *EnvVarCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetBqProjectId

`func (o *EnvVarCredentialsChangeValidateIn) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *EnvVarCredentialsChangeValidateIn) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *EnvVarCredentialsChangeValidateIn) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *EnvVarCredentialsChangeValidateIn) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *EnvVarCredentialsChangeValidateIn) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *EnvVarCredentialsChangeValidateIn) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetSqlWarehouseId

`func (o *EnvVarCredentialsChangeValidateIn) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *EnvVarCredentialsChangeValidateIn) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *EnvVarCredentialsChangeValidateIn) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.

### HasSqlWarehouseId

`func (o *EnvVarCredentialsChangeValidateIn) HasSqlWarehouseId() bool`

HasSqlWarehouseId returns a boolean if a field has been set.

### SetSqlWarehouseIdNil

`func (o *EnvVarCredentialsChangeValidateIn) SetSqlWarehouseIdNil(b bool)`

 SetSqlWarehouseIdNil sets the value for SqlWarehouseId to be an explicit nil

### UnsetSqlWarehouseId
`func (o *EnvVarCredentialsChangeValidateIn) UnsetSqlWarehouseId()`

UnsetSqlWarehouseId ensures that no value is present for SqlWarehouseId, not even an explicit nil
### GetEnvVarName

`func (o *EnvVarCredentialsChangeValidateIn) GetEnvVarName() string`

GetEnvVarName returns the EnvVarName field if non-nil, zero value otherwise.

### GetEnvVarNameOk

`func (o *EnvVarCredentialsChangeValidateIn) GetEnvVarNameOk() (*string, bool)`

GetEnvVarNameOk returns a tuple with the EnvVarName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVarName

`func (o *EnvVarCredentialsChangeValidateIn) SetEnvVarName(v string)`

SetEnvVarName sets EnvVarName field to given value.

### HasEnvVarName

`func (o *EnvVarCredentialsChangeValidateIn) HasEnvVarName() bool`

HasEnvVarName returns a boolean if a field has been set.

### SetEnvVarNameNil

`func (o *EnvVarCredentialsChangeValidateIn) SetEnvVarNameNil(b bool)`

 SetEnvVarNameNil sets the value for EnvVarName to be an explicit nil

### UnsetEnvVarName
`func (o *EnvVarCredentialsChangeValidateIn) UnsetEnvVarName()`

UnsetEnvVarName ensures that no value is present for EnvVarName, not even an explicit nil
### GetKmsKeyId

`func (o *EnvVarCredentialsChangeValidateIn) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *EnvVarCredentialsChangeValidateIn) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *EnvVarCredentialsChangeValidateIn) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *EnvVarCredentialsChangeValidateIn) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.

### SetKmsKeyIdNil

`func (o *EnvVarCredentialsChangeValidateIn) SetKmsKeyIdNil(b bool)`

 SetKmsKeyIdNil sets the value for KmsKeyId to be an explicit nil

### UnsetKmsKeyId
`func (o *EnvVarCredentialsChangeValidateIn) UnsetKmsKeyId()`

UnsetKmsKeyId ensures that no value is present for KmsKeyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


