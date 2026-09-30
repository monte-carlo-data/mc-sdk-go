# AwsSecretsManagerCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**ConnectionType** | **string** | What the credentials are for, hyphenated, such as &#x60;snowflake&#x60; or &#x60;bigquery&#x60;. Decides which checks run. | 
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**SqlWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**AwsSecret** | **string** | Name or ARN of the AWS Secrets Manager secret holding the connection&#39;s credentials. | 
**AwsRegion** | Pointer to **NullableString** | AWS region of the secret. Omit it to use the deployment&#39;s own region. | [optional] 
**AssumableRole** | Pointer to **NullableString** | ARN of a role the deployment assumes to read the secret. Omit it to read as itself. | [optional] 
**ExternalId** | Pointer to **NullableString** | External id the assumed role&#39;s trust policy requires, if it requires one. | [optional] 

## Methods

### NewAwsSecretsManagerCredentialsValidateIn

`func NewAwsSecretsManagerCredentialsValidateIn(deploymentId string, connectionType string, awsSecret string, ) *AwsSecretsManagerCredentialsValidateIn`

NewAwsSecretsManagerCredentialsValidateIn instantiates a new AwsSecretsManagerCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSecretsManagerCredentialsValidateInWithDefaults

`func NewAwsSecretsManagerCredentialsValidateInWithDefaults() *AwsSecretsManagerCredentialsValidateIn`

NewAwsSecretsManagerCredentialsValidateInWithDefaults instantiates a new AwsSecretsManagerCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *AwsSecretsManagerCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AwsSecretsManagerCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AwsSecretsManagerCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetConnectionType

`func (o *AwsSecretsManagerCredentialsValidateIn) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *AwsSecretsManagerCredentialsValidateIn) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *AwsSecretsManagerCredentialsValidateIn) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetBqProjectId

`func (o *AwsSecretsManagerCredentialsValidateIn) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *AwsSecretsManagerCredentialsValidateIn) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *AwsSecretsManagerCredentialsValidateIn) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *AwsSecretsManagerCredentialsValidateIn) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *AwsSecretsManagerCredentialsValidateIn) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *AwsSecretsManagerCredentialsValidateIn) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetSqlWarehouseId

`func (o *AwsSecretsManagerCredentialsValidateIn) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *AwsSecretsManagerCredentialsValidateIn) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *AwsSecretsManagerCredentialsValidateIn) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.

### HasSqlWarehouseId

`func (o *AwsSecretsManagerCredentialsValidateIn) HasSqlWarehouseId() bool`

HasSqlWarehouseId returns a boolean if a field has been set.

### SetSqlWarehouseIdNil

`func (o *AwsSecretsManagerCredentialsValidateIn) SetSqlWarehouseIdNil(b bool)`

 SetSqlWarehouseIdNil sets the value for SqlWarehouseId to be an explicit nil

### UnsetSqlWarehouseId
`func (o *AwsSecretsManagerCredentialsValidateIn) UnsetSqlWarehouseId()`

UnsetSqlWarehouseId ensures that no value is present for SqlWarehouseId, not even an explicit nil
### GetAwsSecret

`func (o *AwsSecretsManagerCredentialsValidateIn) GetAwsSecret() string`

GetAwsSecret returns the AwsSecret field if non-nil, zero value otherwise.

### GetAwsSecretOk

`func (o *AwsSecretsManagerCredentialsValidateIn) GetAwsSecretOk() (*string, bool)`

GetAwsSecretOk returns a tuple with the AwsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecret

`func (o *AwsSecretsManagerCredentialsValidateIn) SetAwsSecret(v string)`

SetAwsSecret sets AwsSecret field to given value.


### GetAwsRegion

`func (o *AwsSecretsManagerCredentialsValidateIn) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AwsSecretsManagerCredentialsValidateIn) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AwsSecretsManagerCredentialsValidateIn) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *AwsSecretsManagerCredentialsValidateIn) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### SetAwsRegionNil

`func (o *AwsSecretsManagerCredentialsValidateIn) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *AwsSecretsManagerCredentialsValidateIn) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetAssumableRole

`func (o *AwsSecretsManagerCredentialsValidateIn) GetAssumableRole() string`

GetAssumableRole returns the AssumableRole field if non-nil, zero value otherwise.

### GetAssumableRoleOk

`func (o *AwsSecretsManagerCredentialsValidateIn) GetAssumableRoleOk() (*string, bool)`

GetAssumableRoleOk returns a tuple with the AssumableRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumableRole

`func (o *AwsSecretsManagerCredentialsValidateIn) SetAssumableRole(v string)`

SetAssumableRole sets AssumableRole field to given value.

### HasAssumableRole

`func (o *AwsSecretsManagerCredentialsValidateIn) HasAssumableRole() bool`

HasAssumableRole returns a boolean if a field has been set.

### SetAssumableRoleNil

`func (o *AwsSecretsManagerCredentialsValidateIn) SetAssumableRoleNil(b bool)`

 SetAssumableRoleNil sets the value for AssumableRole to be an explicit nil

### UnsetAssumableRole
`func (o *AwsSecretsManagerCredentialsValidateIn) UnsetAssumableRole()`

UnsetAssumableRole ensures that no value is present for AssumableRole, not even an explicit nil
### GetExternalId

`func (o *AwsSecretsManagerCredentialsValidateIn) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsSecretsManagerCredentialsValidateIn) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsSecretsManagerCredentialsValidateIn) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AwsSecretsManagerCredentialsValidateIn) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AwsSecretsManagerCredentialsValidateIn) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AwsSecretsManagerCredentialsValidateIn) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


