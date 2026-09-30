# DatabricksSqlWarehouseCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceUrl** | Pointer to **NullableString** | URL of the Databricks workspace, or its host name. | [optional] 
**SqlWarehouseId** | Pointer to **NullableString** | ID of the Databricks SQL warehouse the connection runs on. | [optional] 
**WorkspaceId** | Pointer to **NullableString** | ID of the Databricks workspace. | [optional] 
**Token** | Pointer to **NullableString** | Personal access token or service principal token. Send this, or &#x60;oauth_client_id&#x60; and &#x60;oauth_client_secret&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**OauthClientId** | Pointer to **NullableString** | Client ID of the service principal Monte Carlo signs in as with OAuth. Send it with &#x60;oauth_client_secret&#x60;, instead of &#x60;token&#x60;. | [optional] 
**OauthClientSecret** | Pointer to **NullableString** | OAuth secret of the service principal in &#x60;oauth_client_id&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**AzureTenantId** | Pointer to **NullableString** | Microsoft Entra ID tenant, for a service principal Azure manages. Send it with &#x60;azure_workspace_resource_id&#x60; and the OAuth client. | [optional] 
**AzureWorkspaceResourceId** | Pointer to **NullableString** | Azure resource ID of the workspace, for a service principal Azure manages. Send it with &#x60;azure_tenant_id&#x60; and the OAuth client. | [optional] 

## Methods

### NewDatabricksSqlWarehouseCredentialsPatch

`func NewDatabricksSqlWarehouseCredentialsPatch() *DatabricksSqlWarehouseCredentialsPatch`

NewDatabricksSqlWarehouseCredentialsPatch instantiates a new DatabricksSqlWarehouseCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabricksSqlWarehouseCredentialsPatchWithDefaults

`func NewDatabricksSqlWarehouseCredentialsPatchWithDefaults() *DatabricksSqlWarehouseCredentialsPatch`

NewDatabricksSqlWarehouseCredentialsPatchWithDefaults instantiates a new DatabricksSqlWarehouseCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceUrl

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetWorkspaceUrl() string`

GetWorkspaceUrl returns the WorkspaceUrl field if non-nil, zero value otherwise.

### GetWorkspaceUrlOk

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetWorkspaceUrlOk() (*string, bool)`

GetWorkspaceUrlOk returns a tuple with the WorkspaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceUrl

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetWorkspaceUrl(v string)`

SetWorkspaceUrl sets WorkspaceUrl field to given value.

### HasWorkspaceUrl

`func (o *DatabricksSqlWarehouseCredentialsPatch) HasWorkspaceUrl() bool`

HasWorkspaceUrl returns a boolean if a field has been set.

### SetWorkspaceUrlNil

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetWorkspaceUrlNil(b bool)`

 SetWorkspaceUrlNil sets the value for WorkspaceUrl to be an explicit nil

### UnsetWorkspaceUrl
`func (o *DatabricksSqlWarehouseCredentialsPatch) UnsetWorkspaceUrl()`

UnsetWorkspaceUrl ensures that no value is present for WorkspaceUrl, not even an explicit nil
### GetSqlWarehouseId

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.

### HasSqlWarehouseId

`func (o *DatabricksSqlWarehouseCredentialsPatch) HasSqlWarehouseId() bool`

HasSqlWarehouseId returns a boolean if a field has been set.

### SetSqlWarehouseIdNil

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetSqlWarehouseIdNil(b bool)`

 SetSqlWarehouseIdNil sets the value for SqlWarehouseId to be an explicit nil

### UnsetSqlWarehouseId
`func (o *DatabricksSqlWarehouseCredentialsPatch) UnsetSqlWarehouseId()`

UnsetSqlWarehouseId ensures that no value is present for SqlWarehouseId, not even an explicit nil
### GetWorkspaceId

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *DatabricksSqlWarehouseCredentialsPatch) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *DatabricksSqlWarehouseCredentialsPatch) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
### GetToken

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DatabricksSqlWarehouseCredentialsPatch) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *DatabricksSqlWarehouseCredentialsPatch) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetOauthClientId

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *DatabricksSqlWarehouseCredentialsPatch) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### SetOauthClientIdNil

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetOauthClientIdNil(b bool)`

 SetOauthClientIdNil sets the value for OauthClientId to be an explicit nil

### UnsetOauthClientId
`func (o *DatabricksSqlWarehouseCredentialsPatch) UnsetOauthClientId()`

UnsetOauthClientId ensures that no value is present for OauthClientId, not even an explicit nil
### GetOauthClientSecret

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.

### HasOauthClientSecret

`func (o *DatabricksSqlWarehouseCredentialsPatch) HasOauthClientSecret() bool`

HasOauthClientSecret returns a boolean if a field has been set.

### SetOauthClientSecretNil

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetOauthClientSecretNil(b bool)`

 SetOauthClientSecretNil sets the value for OauthClientSecret to be an explicit nil

### UnsetOauthClientSecret
`func (o *DatabricksSqlWarehouseCredentialsPatch) UnsetOauthClientSecret()`

UnsetOauthClientSecret ensures that no value is present for OauthClientSecret, not even an explicit nil
### GetAzureTenantId

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *DatabricksSqlWarehouseCredentialsPatch) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### SetAzureTenantIdNil

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetAzureTenantIdNil(b bool)`

 SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil

### UnsetAzureTenantId
`func (o *DatabricksSqlWarehouseCredentialsPatch) UnsetAzureTenantId()`

UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
### GetAzureWorkspaceResourceId

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetAzureWorkspaceResourceId() string`

GetAzureWorkspaceResourceId returns the AzureWorkspaceResourceId field if non-nil, zero value otherwise.

### GetAzureWorkspaceResourceIdOk

`func (o *DatabricksSqlWarehouseCredentialsPatch) GetAzureWorkspaceResourceIdOk() (*string, bool)`

GetAzureWorkspaceResourceIdOk returns a tuple with the AzureWorkspaceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureWorkspaceResourceId

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetAzureWorkspaceResourceId(v string)`

SetAzureWorkspaceResourceId sets AzureWorkspaceResourceId field to given value.

### HasAzureWorkspaceResourceId

`func (o *DatabricksSqlWarehouseCredentialsPatch) HasAzureWorkspaceResourceId() bool`

HasAzureWorkspaceResourceId returns a boolean if a field has been set.

### SetAzureWorkspaceResourceIdNil

`func (o *DatabricksSqlWarehouseCredentialsPatch) SetAzureWorkspaceResourceIdNil(b bool)`

 SetAzureWorkspaceResourceIdNil sets the value for AzureWorkspaceResourceId to be an explicit nil

### UnsetAzureWorkspaceResourceId
`func (o *DatabricksSqlWarehouseCredentialsPatch) UnsetAzureWorkspaceResourceId()`

UnsetAzureWorkspaceResourceId ensures that no value is present for AzureWorkspaceResourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


