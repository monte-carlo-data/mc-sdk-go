# EnvVarCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**DatabricksWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**EnvVarName** | Pointer to **NullableString** | Name of the environment variable on the deployment that holds the connection&#39;s credentials. Must start with &#x60;MCD_&#x60;. | [optional] 
**KmsKeyId** | Pointer to **NullableString** | AWS KMS key the variable&#39;s value is encrypted with. Omit it for a value stored in the clear. | [optional] 

## Methods

### NewEnvVarCredentialsPatch

`func NewEnvVarCredentialsPatch() *EnvVarCredentialsPatch`

NewEnvVarCredentialsPatch instantiates a new EnvVarCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvVarCredentialsPatchWithDefaults

`func NewEnvVarCredentialsPatchWithDefaults() *EnvVarCredentialsPatch`

NewEnvVarCredentialsPatchWithDefaults instantiates a new EnvVarCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBqProjectId

`func (o *EnvVarCredentialsPatch) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *EnvVarCredentialsPatch) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *EnvVarCredentialsPatch) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *EnvVarCredentialsPatch) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *EnvVarCredentialsPatch) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *EnvVarCredentialsPatch) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetDatabricksWarehouseId

`func (o *EnvVarCredentialsPatch) GetDatabricksWarehouseId() string`

GetDatabricksWarehouseId returns the DatabricksWarehouseId field if non-nil, zero value otherwise.

### GetDatabricksWarehouseIdOk

`func (o *EnvVarCredentialsPatch) GetDatabricksWarehouseIdOk() (*string, bool)`

GetDatabricksWarehouseIdOk returns a tuple with the DatabricksWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabricksWarehouseId

`func (o *EnvVarCredentialsPatch) SetDatabricksWarehouseId(v string)`

SetDatabricksWarehouseId sets DatabricksWarehouseId field to given value.

### HasDatabricksWarehouseId

`func (o *EnvVarCredentialsPatch) HasDatabricksWarehouseId() bool`

HasDatabricksWarehouseId returns a boolean if a field has been set.

### SetDatabricksWarehouseIdNil

`func (o *EnvVarCredentialsPatch) SetDatabricksWarehouseIdNil(b bool)`

 SetDatabricksWarehouseIdNil sets the value for DatabricksWarehouseId to be an explicit nil

### UnsetDatabricksWarehouseId
`func (o *EnvVarCredentialsPatch) UnsetDatabricksWarehouseId()`

UnsetDatabricksWarehouseId ensures that no value is present for DatabricksWarehouseId, not even an explicit nil
### GetEnvVarName

`func (o *EnvVarCredentialsPatch) GetEnvVarName() string`

GetEnvVarName returns the EnvVarName field if non-nil, zero value otherwise.

### GetEnvVarNameOk

`func (o *EnvVarCredentialsPatch) GetEnvVarNameOk() (*string, bool)`

GetEnvVarNameOk returns a tuple with the EnvVarName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVarName

`func (o *EnvVarCredentialsPatch) SetEnvVarName(v string)`

SetEnvVarName sets EnvVarName field to given value.

### HasEnvVarName

`func (o *EnvVarCredentialsPatch) HasEnvVarName() bool`

HasEnvVarName returns a boolean if a field has been set.

### SetEnvVarNameNil

`func (o *EnvVarCredentialsPatch) SetEnvVarNameNil(b bool)`

 SetEnvVarNameNil sets the value for EnvVarName to be an explicit nil

### UnsetEnvVarName
`func (o *EnvVarCredentialsPatch) UnsetEnvVarName()`

UnsetEnvVarName ensures that no value is present for EnvVarName, not even an explicit nil
### GetKmsKeyId

`func (o *EnvVarCredentialsPatch) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *EnvVarCredentialsPatch) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *EnvVarCredentialsPatch) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *EnvVarCredentialsPatch) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.

### SetKmsKeyIdNil

`func (o *EnvVarCredentialsPatch) SetKmsKeyIdNil(b bool)`

 SetKmsKeyIdNil sets the value for KmsKeyId to be an explicit nil

### UnsetKmsKeyId
`func (o *EnvVarCredentialsPatch) UnsetKmsKeyId()`

UnsetKmsKeyId ensures that no value is present for KmsKeyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


