# AzureKeyVaultCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BqProjectId** | Pointer to **NullableString** | BigQuery project the connection reads from. Only for a BigQuery connection. | [optional] 
**SqlWarehouseId** | Pointer to **NullableString** | Databricks SQL warehouse the connection runs queries on. Required for a &#x60;databricks-sql-warehouse&#x60; or &#x60;databricks-metastore-sql-warehouse&#x60; connection. | [optional] 
**AkvSecret** | Pointer to **NullableString** | Name of the Azure Key Vault secret holding the connection&#39;s credentials. | [optional] 
**AkvVaultName** | Pointer to **NullableString** | Name of the key vault. Send this, &#x60;akv_vault_url&#x60;, or both. | [optional] 
**AkvVaultUrl** | Pointer to **NullableString** | URL of the key vault. Send this, &#x60;akv_vault_name&#x60;, or both. | [optional] 

## Methods

### NewAzureKeyVaultCredentialsPatch

`func NewAzureKeyVaultCredentialsPatch() *AzureKeyVaultCredentialsPatch`

NewAzureKeyVaultCredentialsPatch instantiates a new AzureKeyVaultCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultCredentialsPatchWithDefaults

`func NewAzureKeyVaultCredentialsPatchWithDefaults() *AzureKeyVaultCredentialsPatch`

NewAzureKeyVaultCredentialsPatchWithDefaults instantiates a new AzureKeyVaultCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBqProjectId

`func (o *AzureKeyVaultCredentialsPatch) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *AzureKeyVaultCredentialsPatch) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *AzureKeyVaultCredentialsPatch) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.

### HasBqProjectId

`func (o *AzureKeyVaultCredentialsPatch) HasBqProjectId() bool`

HasBqProjectId returns a boolean if a field has been set.

### SetBqProjectIdNil

`func (o *AzureKeyVaultCredentialsPatch) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *AzureKeyVaultCredentialsPatch) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetSqlWarehouseId

`func (o *AzureKeyVaultCredentialsPatch) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *AzureKeyVaultCredentialsPatch) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *AzureKeyVaultCredentialsPatch) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.

### HasSqlWarehouseId

`func (o *AzureKeyVaultCredentialsPatch) HasSqlWarehouseId() bool`

HasSqlWarehouseId returns a boolean if a field has been set.

### SetSqlWarehouseIdNil

`func (o *AzureKeyVaultCredentialsPatch) SetSqlWarehouseIdNil(b bool)`

 SetSqlWarehouseIdNil sets the value for SqlWarehouseId to be an explicit nil

### UnsetSqlWarehouseId
`func (o *AzureKeyVaultCredentialsPatch) UnsetSqlWarehouseId()`

UnsetSqlWarehouseId ensures that no value is present for SqlWarehouseId, not even an explicit nil
### GetAkvSecret

`func (o *AzureKeyVaultCredentialsPatch) GetAkvSecret() string`

GetAkvSecret returns the AkvSecret field if non-nil, zero value otherwise.

### GetAkvSecretOk

`func (o *AzureKeyVaultCredentialsPatch) GetAkvSecretOk() (*string, bool)`

GetAkvSecretOk returns a tuple with the AkvSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkvSecret

`func (o *AzureKeyVaultCredentialsPatch) SetAkvSecret(v string)`

SetAkvSecret sets AkvSecret field to given value.

### HasAkvSecret

`func (o *AzureKeyVaultCredentialsPatch) HasAkvSecret() bool`

HasAkvSecret returns a boolean if a field has been set.

### SetAkvSecretNil

`func (o *AzureKeyVaultCredentialsPatch) SetAkvSecretNil(b bool)`

 SetAkvSecretNil sets the value for AkvSecret to be an explicit nil

### UnsetAkvSecret
`func (o *AzureKeyVaultCredentialsPatch) UnsetAkvSecret()`

UnsetAkvSecret ensures that no value is present for AkvSecret, not even an explicit nil
### GetAkvVaultName

`func (o *AzureKeyVaultCredentialsPatch) GetAkvVaultName() string`

GetAkvVaultName returns the AkvVaultName field if non-nil, zero value otherwise.

### GetAkvVaultNameOk

`func (o *AzureKeyVaultCredentialsPatch) GetAkvVaultNameOk() (*string, bool)`

GetAkvVaultNameOk returns a tuple with the AkvVaultName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkvVaultName

`func (o *AzureKeyVaultCredentialsPatch) SetAkvVaultName(v string)`

SetAkvVaultName sets AkvVaultName field to given value.

### HasAkvVaultName

`func (o *AzureKeyVaultCredentialsPatch) HasAkvVaultName() bool`

HasAkvVaultName returns a boolean if a field has been set.

### SetAkvVaultNameNil

`func (o *AzureKeyVaultCredentialsPatch) SetAkvVaultNameNil(b bool)`

 SetAkvVaultNameNil sets the value for AkvVaultName to be an explicit nil

### UnsetAkvVaultName
`func (o *AzureKeyVaultCredentialsPatch) UnsetAkvVaultName()`

UnsetAkvVaultName ensures that no value is present for AkvVaultName, not even an explicit nil
### GetAkvVaultUrl

`func (o *AzureKeyVaultCredentialsPatch) GetAkvVaultUrl() string`

GetAkvVaultUrl returns the AkvVaultUrl field if non-nil, zero value otherwise.

### GetAkvVaultUrlOk

`func (o *AzureKeyVaultCredentialsPatch) GetAkvVaultUrlOk() (*string, bool)`

GetAkvVaultUrlOk returns a tuple with the AkvVaultUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkvVaultUrl

`func (o *AzureKeyVaultCredentialsPatch) SetAkvVaultUrl(v string)`

SetAkvVaultUrl sets AkvVaultUrl field to given value.

### HasAkvVaultUrl

`func (o *AzureKeyVaultCredentialsPatch) HasAkvVaultUrl() bool`

HasAkvVaultUrl returns a boolean if a field has been set.

### SetAkvVaultUrlNil

`func (o *AzureKeyVaultCredentialsPatch) SetAkvVaultUrlNil(b bool)`

 SetAkvVaultUrlNil sets the value for AkvVaultUrl to be an explicit nil

### UnsetAkvVaultUrl
`func (o *AzureKeyVaultCredentialsPatch) UnsetAkvVaultUrl()`

UnsetAkvVaultUrl ensures that no value is present for AkvVaultUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


