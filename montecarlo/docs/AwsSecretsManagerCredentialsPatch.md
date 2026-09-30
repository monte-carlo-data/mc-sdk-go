# AwsSecretsManagerCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**SqlWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**AwsSecret** | Pointer to **NullableString** | Name or ARN of the AWS Secrets Manager secret holding the connection&#39;s credentials. | [optional] 
**AwsRegion** | Pointer to **NullableString** | AWS region of the secret. Omit it to use the deployment&#39;s own region. | [optional] 
**AssumableRole** | Pointer to **NullableString** | ARN of a role the deployment assumes to read the secret. Omit it to read as itself. | [optional] 
**ExternalId** | Pointer to **NullableString** | External id the assumed role&#39;s trust policy requires, if it requires one. | [optional] 

## Methods

### NewAwsSecretsManagerCredentialsPatch

`func NewAwsSecretsManagerCredentialsPatch() *AwsSecretsManagerCredentialsPatch`

NewAwsSecretsManagerCredentialsPatch instantiates a new AwsSecretsManagerCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSecretsManagerCredentialsPatchWithDefaults

`func NewAwsSecretsManagerCredentialsPatchWithDefaults() *AwsSecretsManagerCredentialsPatch`

NewAwsSecretsManagerCredentialsPatchWithDefaults instantiates a new AwsSecretsManagerCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBqProjectId

`func (o *AwsSecretsManagerCredentialsPatch) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *AwsSecretsManagerCredentialsPatch) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *AwsSecretsManagerCredentialsPatch) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *AwsSecretsManagerCredentialsPatch) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *AwsSecretsManagerCredentialsPatch) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *AwsSecretsManagerCredentialsPatch) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetSqlWarehouseId

`func (o *AwsSecretsManagerCredentialsPatch) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *AwsSecretsManagerCredentialsPatch) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *AwsSecretsManagerCredentialsPatch) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.

### HasSqlWarehouseId

`func (o *AwsSecretsManagerCredentialsPatch) HasSqlWarehouseId() bool`

HasSqlWarehouseId returns a boolean if a field has been set.

### SetSqlWarehouseIdNil

`func (o *AwsSecretsManagerCredentialsPatch) SetSqlWarehouseIdNil(b bool)`

 SetSqlWarehouseIdNil sets the value for SqlWarehouseId to be an explicit nil

### UnsetSqlWarehouseId
`func (o *AwsSecretsManagerCredentialsPatch) UnsetSqlWarehouseId()`

UnsetSqlWarehouseId ensures that no value is present for SqlWarehouseId, not even an explicit nil
### GetAwsSecret

`func (o *AwsSecretsManagerCredentialsPatch) GetAwsSecret() string`

GetAwsSecret returns the AwsSecret field if non-nil, zero value otherwise.

### GetAwsSecretOk

`func (o *AwsSecretsManagerCredentialsPatch) GetAwsSecretOk() (*string, bool)`

GetAwsSecretOk returns a tuple with the AwsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecret

`func (o *AwsSecretsManagerCredentialsPatch) SetAwsSecret(v string)`

SetAwsSecret sets AwsSecret field to given value.

### HasAwsSecret

`func (o *AwsSecretsManagerCredentialsPatch) HasAwsSecret() bool`

HasAwsSecret returns a boolean if a field has been set.

### SetAwsSecretNil

`func (o *AwsSecretsManagerCredentialsPatch) SetAwsSecretNil(b bool)`

 SetAwsSecretNil sets the value for AwsSecret to be an explicit nil

### UnsetAwsSecret
`func (o *AwsSecretsManagerCredentialsPatch) UnsetAwsSecret()`

UnsetAwsSecret ensures that no value is present for AwsSecret, not even an explicit nil
### GetAwsRegion

`func (o *AwsSecretsManagerCredentialsPatch) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AwsSecretsManagerCredentialsPatch) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AwsSecretsManagerCredentialsPatch) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *AwsSecretsManagerCredentialsPatch) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### SetAwsRegionNil

`func (o *AwsSecretsManagerCredentialsPatch) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *AwsSecretsManagerCredentialsPatch) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetAssumableRole

`func (o *AwsSecretsManagerCredentialsPatch) GetAssumableRole() string`

GetAssumableRole returns the AssumableRole field if non-nil, zero value otherwise.

### GetAssumableRoleOk

`func (o *AwsSecretsManagerCredentialsPatch) GetAssumableRoleOk() (*string, bool)`

GetAssumableRoleOk returns a tuple with the AssumableRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumableRole

`func (o *AwsSecretsManagerCredentialsPatch) SetAssumableRole(v string)`

SetAssumableRole sets AssumableRole field to given value.

### HasAssumableRole

`func (o *AwsSecretsManagerCredentialsPatch) HasAssumableRole() bool`

HasAssumableRole returns a boolean if a field has been set.

### SetAssumableRoleNil

`func (o *AwsSecretsManagerCredentialsPatch) SetAssumableRoleNil(b bool)`

 SetAssumableRoleNil sets the value for AssumableRole to be an explicit nil

### UnsetAssumableRole
`func (o *AwsSecretsManagerCredentialsPatch) UnsetAssumableRole()`

UnsetAssumableRole ensures that no value is present for AssumableRole, not even an explicit nil
### GetExternalId

`func (o *AwsSecretsManagerCredentialsPatch) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsSecretsManagerCredentialsPatch) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsSecretsManagerCredentialsPatch) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AwsSecretsManagerCredentialsPatch) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AwsSecretsManagerCredentialsPatch) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AwsSecretsManagerCredentialsPatch) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


