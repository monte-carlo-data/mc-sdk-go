# FileCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**SqlWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**FilePath** | Pointer to **NullableString** | Path of the file on the deployment that holds the connection&#39;s credentials. | [optional] 

## Methods

### NewFileCredentialsChangeValidateIn

`func NewFileCredentialsChangeValidateIn(deploymentId string, ) *FileCredentialsChangeValidateIn`

NewFileCredentialsChangeValidateIn instantiates a new FileCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCredentialsChangeValidateInWithDefaults

`func NewFileCredentialsChangeValidateInWithDefaults() *FileCredentialsChangeValidateIn`

NewFileCredentialsChangeValidateInWithDefaults instantiates a new FileCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *FileCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *FileCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *FileCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetBqProjectId

`func (o *FileCredentialsChangeValidateIn) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *FileCredentialsChangeValidateIn) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *FileCredentialsChangeValidateIn) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *FileCredentialsChangeValidateIn) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *FileCredentialsChangeValidateIn) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *FileCredentialsChangeValidateIn) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetSqlWarehouseId

`func (o *FileCredentialsChangeValidateIn) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *FileCredentialsChangeValidateIn) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *FileCredentialsChangeValidateIn) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.

### HasSqlWarehouseId

`func (o *FileCredentialsChangeValidateIn) HasSqlWarehouseId() bool`

HasSqlWarehouseId returns a boolean if a field has been set.

### SetSqlWarehouseIdNil

`func (o *FileCredentialsChangeValidateIn) SetSqlWarehouseIdNil(b bool)`

 SetSqlWarehouseIdNil sets the value for SqlWarehouseId to be an explicit nil

### UnsetSqlWarehouseId
`func (o *FileCredentialsChangeValidateIn) UnsetSqlWarehouseId()`

UnsetSqlWarehouseId ensures that no value is present for SqlWarehouseId, not even an explicit nil
### GetFilePath

`func (o *FileCredentialsChangeValidateIn) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FileCredentialsChangeValidateIn) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FileCredentialsChangeValidateIn) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *FileCredentialsChangeValidateIn) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *FileCredentialsChangeValidateIn) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *FileCredentialsChangeValidateIn) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


