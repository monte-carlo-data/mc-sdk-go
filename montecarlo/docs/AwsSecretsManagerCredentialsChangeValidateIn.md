# AwsSecretsManagerCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**SqlWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**AwsSecret** | Pointer to **NullableString** | Name or ARN of the AWS Secrets Manager secret holding the connection&#39;s credentials. | [optional] 
**AwsRegion** | Pointer to **NullableString** | AWS region of the secret. Omit it to use the deployment&#39;s own region. | [optional] 
**AssumableRole** | Pointer to **NullableString** | ARN of a role the deployment assumes to read the secret. Omit it to read as itself. | [optional] 
**ExternalId** | Pointer to **NullableString** | External id the assumed role&#39;s trust policy requires, if it requires one. | [optional] 

## Methods

### NewAwsSecretsManagerCredentialsChangeValidateIn

`func NewAwsSecretsManagerCredentialsChangeValidateIn(deploymentId string, ) *AwsSecretsManagerCredentialsChangeValidateIn`

NewAwsSecretsManagerCredentialsChangeValidateIn instantiates a new AwsSecretsManagerCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSecretsManagerCredentialsChangeValidateInWithDefaults

`func NewAwsSecretsManagerCredentialsChangeValidateInWithDefaults() *AwsSecretsManagerCredentialsChangeValidateIn`

NewAwsSecretsManagerCredentialsChangeValidateInWithDefaults instantiates a new AwsSecretsManagerCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetBqProjectId

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *AwsSecretsManagerCredentialsChangeValidateIn) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetSqlWarehouseId

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.

### HasSqlWarehouseId

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) HasSqlWarehouseId() bool`

HasSqlWarehouseId returns a boolean if a field has been set.

### SetSqlWarehouseIdNil

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetSqlWarehouseIdNil(b bool)`

 SetSqlWarehouseIdNil sets the value for SqlWarehouseId to be an explicit nil

### UnsetSqlWarehouseId
`func (o *AwsSecretsManagerCredentialsChangeValidateIn) UnsetSqlWarehouseId()`

UnsetSqlWarehouseId ensures that no value is present for SqlWarehouseId, not even an explicit nil
### GetAwsSecret

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetAwsSecret() string`

GetAwsSecret returns the AwsSecret field if non-nil, zero value otherwise.

### GetAwsSecretOk

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetAwsSecretOk() (*string, bool)`

GetAwsSecretOk returns a tuple with the AwsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecret

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetAwsSecret(v string)`

SetAwsSecret sets AwsSecret field to given value.

### HasAwsSecret

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) HasAwsSecret() bool`

HasAwsSecret returns a boolean if a field has been set.

### SetAwsSecretNil

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetAwsSecretNil(b bool)`

 SetAwsSecretNil sets the value for AwsSecret to be an explicit nil

### UnsetAwsSecret
`func (o *AwsSecretsManagerCredentialsChangeValidateIn) UnsetAwsSecret()`

UnsetAwsSecret ensures that no value is present for AwsSecret, not even an explicit nil
### GetAwsRegion

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### SetAwsRegionNil

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *AwsSecretsManagerCredentialsChangeValidateIn) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetAssumableRole

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetAssumableRole() string`

GetAssumableRole returns the AssumableRole field if non-nil, zero value otherwise.

### GetAssumableRoleOk

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetAssumableRoleOk() (*string, bool)`

GetAssumableRoleOk returns a tuple with the AssumableRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumableRole

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetAssumableRole(v string)`

SetAssumableRole sets AssumableRole field to given value.

### HasAssumableRole

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) HasAssumableRole() bool`

HasAssumableRole returns a boolean if a field has been set.

### SetAssumableRoleNil

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetAssumableRoleNil(b bool)`

 SetAssumableRoleNil sets the value for AssumableRole to be an explicit nil

### UnsetAssumableRole
`func (o *AwsSecretsManagerCredentialsChangeValidateIn) UnsetAssumableRole()`

UnsetAssumableRole ensures that no value is present for AssumableRole, not even an explicit nil
### GetExternalId

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AwsSecretsManagerCredentialsChangeValidateIn) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AwsSecretsManagerCredentialsChangeValidateIn) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


