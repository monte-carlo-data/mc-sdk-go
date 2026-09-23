# AzureKeyVaultCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**ConnectionType** | **string** | What the credentials are for, hyphenated, such as &#x60;snowflake&#x60; or &#x60;bigquery&#x60;. Decides which checks run. | 
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**DatabricksWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**AkvSecret** | **string** | Name of the Azure Key Vault secret holding the connection&#39;s credentials. | 
**AkvVaultName** | Pointer to **NullableString** | Name of the key vault. Send this, &#x60;akv_vault_url&#x60;, or both. | [optional] 
**AkvVaultUrl** | Pointer to **NullableString** | URL of the key vault. Send this, &#x60;akv_vault_name&#x60;, or both. | [optional] 

## Methods

### NewAzureKeyVaultCredentialsValidateIn

`func NewAzureKeyVaultCredentialsValidateIn(deploymentId string, connectionType string, akvSecret string, ) *AzureKeyVaultCredentialsValidateIn`

NewAzureKeyVaultCredentialsValidateIn instantiates a new AzureKeyVaultCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultCredentialsValidateInWithDefaults

`func NewAzureKeyVaultCredentialsValidateInWithDefaults() *AzureKeyVaultCredentialsValidateIn`

NewAzureKeyVaultCredentialsValidateInWithDefaults instantiates a new AzureKeyVaultCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *AzureKeyVaultCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AzureKeyVaultCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AzureKeyVaultCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetConnectionType

`func (o *AzureKeyVaultCredentialsValidateIn) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *AzureKeyVaultCredentialsValidateIn) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *AzureKeyVaultCredentialsValidateIn) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetBqProjectId

`func (o *AzureKeyVaultCredentialsValidateIn) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *AzureKeyVaultCredentialsValidateIn) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *AzureKeyVaultCredentialsValidateIn) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *AzureKeyVaultCredentialsValidateIn) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *AzureKeyVaultCredentialsValidateIn) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *AzureKeyVaultCredentialsValidateIn) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetDatabricksWarehouseId

`func (o *AzureKeyVaultCredentialsValidateIn) GetDatabricksWarehouseId() string`

GetDatabricksWarehouseId returns the DatabricksWarehouseId field if non-nil, zero value otherwise.

### GetDatabricksWarehouseIdOk

`func (o *AzureKeyVaultCredentialsValidateIn) GetDatabricksWarehouseIdOk() (*string, bool)`

GetDatabricksWarehouseIdOk returns a tuple with the DatabricksWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabricksWarehouseId

`func (o *AzureKeyVaultCredentialsValidateIn) SetDatabricksWarehouseId(v string)`

SetDatabricksWarehouseId sets DatabricksWarehouseId field to given value.

### HasDatabricksWarehouseId

`func (o *AzureKeyVaultCredentialsValidateIn) HasDatabricksWarehouseId() bool`

HasDatabricksWarehouseId returns a boolean if a field has been set.

### SetDatabricksWarehouseIdNil

`func (o *AzureKeyVaultCredentialsValidateIn) SetDatabricksWarehouseIdNil(b bool)`

 SetDatabricksWarehouseIdNil sets the value for DatabricksWarehouseId to be an explicit nil

### UnsetDatabricksWarehouseId
`func (o *AzureKeyVaultCredentialsValidateIn) UnsetDatabricksWarehouseId()`

UnsetDatabricksWarehouseId ensures that no value is present for DatabricksWarehouseId, not even an explicit nil
### GetAkvSecret

`func (o *AzureKeyVaultCredentialsValidateIn) GetAkvSecret() string`

GetAkvSecret returns the AkvSecret field if non-nil, zero value otherwise.

### GetAkvSecretOk

`func (o *AzureKeyVaultCredentialsValidateIn) GetAkvSecretOk() (*string, bool)`

GetAkvSecretOk returns a tuple with the AkvSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkvSecret

`func (o *AzureKeyVaultCredentialsValidateIn) SetAkvSecret(v string)`

SetAkvSecret sets AkvSecret field to given value.


### GetAkvVaultName

`func (o *AzureKeyVaultCredentialsValidateIn) GetAkvVaultName() string`

GetAkvVaultName returns the AkvVaultName field if non-nil, zero value otherwise.

### GetAkvVaultNameOk

`func (o *AzureKeyVaultCredentialsValidateIn) GetAkvVaultNameOk() (*string, bool)`

GetAkvVaultNameOk returns a tuple with the AkvVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkvVaultName

`func (o *AzureKeyVaultCredentialsValidateIn) SetAkvVaultName(v string)`

SetAkvVaultName sets AkvVaultName field to given value.

### HasAkvVaultName

`func (o *AzureKeyVaultCredentialsValidateIn) HasAkvVaultName() bool`

HasAkvVaultName returns a boolean if a field has been set.

### SetAkvVaultNameNil

`func (o *AzureKeyVaultCredentialsValidateIn) SetAkvVaultNameNil(b bool)`

 SetAkvVaultNameNil sets the value for AkvVaultName to be an explicit nil

### UnsetAkvVaultName
`func (o *AzureKeyVaultCredentialsValidateIn) UnsetAkvVaultName()`

UnsetAkvVaultName ensures that no value is present for AkvVaultName, not even an explicit nil
### GetAkvVaultUrl

`func (o *AzureKeyVaultCredentialsValidateIn) GetAkvVaultUrl() string`

GetAkvVaultUrl returns the AkvVaultUrl field if non-nil, zero value otherwise.

### GetAkvVaultUrlOk

`func (o *AzureKeyVaultCredentialsValidateIn) GetAkvVaultUrlOk() (*string, bool)`

GetAkvVaultUrlOk returns a tuple with the AkvVaultUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkvVaultUrl

`func (o *AzureKeyVaultCredentialsValidateIn) SetAkvVaultUrl(v string)`

SetAkvVaultUrl sets AkvVaultUrl field to given value.

### HasAkvVaultUrl

`func (o *AzureKeyVaultCredentialsValidateIn) HasAkvVaultUrl() bool`

HasAkvVaultUrl returns a boolean if a field has been set.

### SetAkvVaultUrlNil

`func (o *AzureKeyVaultCredentialsValidateIn) SetAkvVaultUrlNil(b bool)`

 SetAkvVaultUrlNil sets the value for AkvVaultUrl to be an explicit nil

### UnsetAkvVaultUrl
`func (o *AzureKeyVaultCredentialsValidateIn) UnsetAkvVaultUrl()`

UnsetAkvVaultUrl ensures that no value is present for AkvVaultUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


