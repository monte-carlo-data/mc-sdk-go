# AwsSecretsManagerCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60; or &#x60;bigquery&#x60;, or one of your custom connector types. Fixed once created. | 
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**DatabricksWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**AwsSecret** | **string** | Name or ARN of the AWS Secrets Manager secret holding the connection&#39;s credentials. | 
**AwsRegion** | Pointer to **NullableString** | AWS region of the secret. Omit it to use the deployment&#39;s own region. | [optional] 
**AssumableRole** | Pointer to **NullableString** | ARN of a role the deployment assumes to read the secret. Omit it to read as itself. | [optional] 
**ExternalId** | Pointer to **NullableString** | External id the assumed role&#39;s trust policy requires, if it requires one. | [optional] 

## Methods

### NewAwsSecretsManagerCredentialsIn

`func NewAwsSecretsManagerCredentialsIn(connectionType string, awsSecret string, ) *AwsSecretsManagerCredentialsIn`

NewAwsSecretsManagerCredentialsIn instantiates a new AwsSecretsManagerCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSecretsManagerCredentialsInWithDefaults

`func NewAwsSecretsManagerCredentialsInWithDefaults() *AwsSecretsManagerCredentialsIn`

NewAwsSecretsManagerCredentialsInWithDefaults instantiates a new AwsSecretsManagerCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *AwsSecretsManagerCredentialsIn) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *AwsSecretsManagerCredentialsIn) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *AwsSecretsManagerCredentialsIn) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetBqProjectId

`func (o *AwsSecretsManagerCredentialsIn) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *AwsSecretsManagerCredentialsIn) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *AwsSecretsManagerCredentialsIn) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *AwsSecretsManagerCredentialsIn) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *AwsSecretsManagerCredentialsIn) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *AwsSecretsManagerCredentialsIn) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetDatabricksWarehouseId

`func (o *AwsSecretsManagerCredentialsIn) GetDatabricksWarehouseId() string`

GetDatabricksWarehouseId returns the DatabricksWarehouseId field if non-nil, zero value otherwise.

### GetDatabricksWarehouseIdOk

`func (o *AwsSecretsManagerCredentialsIn) GetDatabricksWarehouseIdOk() (*string, bool)`

GetDatabricksWarehouseIdOk returns a tuple with the DatabricksWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabricksWarehouseId

`func (o *AwsSecretsManagerCredentialsIn) SetDatabricksWarehouseId(v string)`

SetDatabricksWarehouseId sets DatabricksWarehouseId field to given value.

### HasDatabricksWarehouseId

`func (o *AwsSecretsManagerCredentialsIn) HasDatabricksWarehouseId() bool`

HasDatabricksWarehouseId returns a boolean if a field has been set.

### SetDatabricksWarehouseIdNil

`func (o *AwsSecretsManagerCredentialsIn) SetDatabricksWarehouseIdNil(b bool)`

 SetDatabricksWarehouseIdNil sets the value for DatabricksWarehouseId to be an explicit nil

### UnsetDatabricksWarehouseId
`func (o *AwsSecretsManagerCredentialsIn) UnsetDatabricksWarehouseId()`

UnsetDatabricksWarehouseId ensures that no value is present for DatabricksWarehouseId, not even an explicit nil
### GetAwsSecret

`func (o *AwsSecretsManagerCredentialsIn) GetAwsSecret() string`

GetAwsSecret returns the AwsSecret field if non-nil, zero value otherwise.

### GetAwsSecretOk

`func (o *AwsSecretsManagerCredentialsIn) GetAwsSecretOk() (*string, bool)`

GetAwsSecretOk returns a tuple with the AwsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecret

`func (o *AwsSecretsManagerCredentialsIn) SetAwsSecret(v string)`

SetAwsSecret sets AwsSecret field to given value.


### GetAwsRegion

`func (o *AwsSecretsManagerCredentialsIn) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AwsSecretsManagerCredentialsIn) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AwsSecretsManagerCredentialsIn) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *AwsSecretsManagerCredentialsIn) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### SetAwsRegionNil

`func (o *AwsSecretsManagerCredentialsIn) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *AwsSecretsManagerCredentialsIn) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetAssumableRole

`func (o *AwsSecretsManagerCredentialsIn) GetAssumableRole() string`

GetAssumableRole returns the AssumableRole field if non-nil, zero value otherwise.

### GetAssumableRoleOk

`func (o *AwsSecretsManagerCredentialsIn) GetAssumableRoleOk() (*string, bool)`

GetAssumableRoleOk returns a tuple with the AssumableRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumableRole

`func (o *AwsSecretsManagerCredentialsIn) SetAssumableRole(v string)`

SetAssumableRole sets AssumableRole field to given value.

### HasAssumableRole

`func (o *AwsSecretsManagerCredentialsIn) HasAssumableRole() bool`

HasAssumableRole returns a boolean if a field has been set.

### SetAssumableRoleNil

`func (o *AwsSecretsManagerCredentialsIn) SetAssumableRoleNil(b bool)`

 SetAssumableRoleNil sets the value for AssumableRole to be an explicit nil

### UnsetAssumableRole
`func (o *AwsSecretsManagerCredentialsIn) UnsetAssumableRole()`

UnsetAssumableRole ensures that no value is present for AssumableRole, not even an explicit nil
### GetExternalId

`func (o *AwsSecretsManagerCredentialsIn) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsSecretsManagerCredentialsIn) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsSecretsManagerCredentialsIn) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AwsSecretsManagerCredentialsIn) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AwsSecretsManagerCredentialsIn) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AwsSecretsManagerCredentialsIn) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


