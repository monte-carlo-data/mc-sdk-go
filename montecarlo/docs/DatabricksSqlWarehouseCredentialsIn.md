# DatabricksSqlWarehouseCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceUrl** | **string** | URL of the Databricks workspace, or its host name. | 
**SqlWarehouseId** | **string** | ID of the Databricks SQL warehouse the connection runs on. | 
**WorkspaceId** | Pointer to **NullableString** | ID of the Databricks workspace. | [optional] 
**Token** | Pointer to **NullableString** | Personal access token or service principal token. Send this, or &#x60;oauth_client_id&#x60; and &#x60;oauth_client_secret&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**OauthClientId** | Pointer to **NullableString** | Client ID of the service principal Monte Carlo signs in as with OAuth. Send it with &#x60;oauth_client_secret&#x60;, instead of &#x60;token&#x60;. | [optional] 
**OauthClientSecret** | Pointer to **NullableString** | OAuth secret of the service principal in &#x60;oauth_client_id&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**AzureTenantId** | Pointer to **NullableString** | Microsoft Entra ID tenant, for a service principal Azure manages. Send it with &#x60;azure_workspace_resource_id&#x60; and the OAuth client. | [optional] 
**AzureWorkspaceResourceId** | Pointer to **NullableString** | Azure resource ID of the workspace, for a service principal Azure manages. Send it with &#x60;azure_tenant_id&#x60; and the OAuth client. | [optional] 

## Methods

### NewDatabricksSqlWarehouseCredentialsIn

`func NewDatabricksSqlWarehouseCredentialsIn(workspaceUrl string, sqlWarehouseId string, ) *DatabricksSqlWarehouseCredentialsIn`

NewDatabricksSqlWarehouseCredentialsIn instantiates a new DatabricksSqlWarehouseCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabricksSqlWarehouseCredentialsInWithDefaults

`func NewDatabricksSqlWarehouseCredentialsInWithDefaults() *DatabricksSqlWarehouseCredentialsIn`

NewDatabricksSqlWarehouseCredentialsInWithDefaults instantiates a new DatabricksSqlWarehouseCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceUrl

`func (o *DatabricksSqlWarehouseCredentialsIn) GetWorkspaceUrl() string`

GetWorkspaceUrl returns the WorkspaceUrl field if non-nil, zero value otherwise.

### GetWorkspaceUrlOk

`func (o *DatabricksSqlWarehouseCredentialsIn) GetWorkspaceUrlOk() (*string, bool)`

GetWorkspaceUrlOk returns a tuple with the WorkspaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceUrl

`func (o *DatabricksSqlWarehouseCredentialsIn) SetWorkspaceUrl(v string)`

SetWorkspaceUrl sets WorkspaceUrl field to given value.


### GetSqlWarehouseId

`func (o *DatabricksSqlWarehouseCredentialsIn) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *DatabricksSqlWarehouseCredentialsIn) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *DatabricksSqlWarehouseCredentialsIn) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.


### GetWorkspaceId

`func (o *DatabricksSqlWarehouseCredentialsIn) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DatabricksSqlWarehouseCredentialsIn) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DatabricksSqlWarehouseCredentialsIn) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *DatabricksSqlWarehouseCredentialsIn) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *DatabricksSqlWarehouseCredentialsIn) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *DatabricksSqlWarehouseCredentialsIn) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil
### GetToken

`func (o *DatabricksSqlWarehouseCredentialsIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DatabricksSqlWarehouseCredentialsIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DatabricksSqlWarehouseCredentialsIn) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DatabricksSqlWarehouseCredentialsIn) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *DatabricksSqlWarehouseCredentialsIn) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *DatabricksSqlWarehouseCredentialsIn) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetOauthClientId

`func (o *DatabricksSqlWarehouseCredentialsIn) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *DatabricksSqlWarehouseCredentialsIn) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *DatabricksSqlWarehouseCredentialsIn) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.

### HasOauthClientId

`func (o *DatabricksSqlWarehouseCredentialsIn) HasOauthClientId() bool`

HasOauthClientId returns a boolean if a field has been set.

### SetOauthClientIdNil

`func (o *DatabricksSqlWarehouseCredentialsIn) SetOauthClientIdNil(b bool)`

 SetOauthClientIdNil sets the value for OauthClientId to be an explicit nil

### UnsetOauthClientId
`func (o *DatabricksSqlWarehouseCredentialsIn) UnsetOauthClientId()`

UnsetOauthClientId ensures that no value is present for OauthClientId, not even an explicit nil
### GetOauthClientSecret

`func (o *DatabricksSqlWarehouseCredentialsIn) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *DatabricksSqlWarehouseCredentialsIn) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *DatabricksSqlWarehouseCredentialsIn) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.

### HasOauthClientSecret

`func (o *DatabricksSqlWarehouseCredentialsIn) HasOauthClientSecret() bool`

HasOauthClientSecret returns a boolean if a field has been set.

### SetOauthClientSecretNil

`func (o *DatabricksSqlWarehouseCredentialsIn) SetOauthClientSecretNil(b bool)`

 SetOauthClientSecretNil sets the value for OauthClientSecret to be an explicit nil

### UnsetOauthClientSecret
`func (o *DatabricksSqlWarehouseCredentialsIn) UnsetOauthClientSecret()`

UnsetOauthClientSecret ensures that no value is present for OauthClientSecret, not even an explicit nil
### GetAzureTenantId

`func (o *DatabricksSqlWarehouseCredentialsIn) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *DatabricksSqlWarehouseCredentialsIn) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *DatabricksSqlWarehouseCredentialsIn) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *DatabricksSqlWarehouseCredentialsIn) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### SetAzureTenantIdNil

`func (o *DatabricksSqlWarehouseCredentialsIn) SetAzureTenantIdNil(b bool)`

 SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil

### UnsetAzureTenantId
`func (o *DatabricksSqlWarehouseCredentialsIn) UnsetAzureTenantId()`

UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
### GetAzureWorkspaceResourceId

`func (o *DatabricksSqlWarehouseCredentialsIn) GetAzureWorkspaceResourceId() string`

GetAzureWorkspaceResourceId returns the AzureWorkspaceResourceId field if non-nil, zero value otherwise.

### GetAzureWorkspaceResourceIdOk

`func (o *DatabricksSqlWarehouseCredentialsIn) GetAzureWorkspaceResourceIdOk() (*string, bool)`

GetAzureWorkspaceResourceIdOk returns a tuple with the AzureWorkspaceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureWorkspaceResourceId

`func (o *DatabricksSqlWarehouseCredentialsIn) SetAzureWorkspaceResourceId(v string)`

SetAzureWorkspaceResourceId sets AzureWorkspaceResourceId field to given value.

### HasAzureWorkspaceResourceId

`func (o *DatabricksSqlWarehouseCredentialsIn) HasAzureWorkspaceResourceId() bool`

HasAzureWorkspaceResourceId returns a boolean if a field has been set.

### SetAzureWorkspaceResourceIdNil

`func (o *DatabricksSqlWarehouseCredentialsIn) SetAzureWorkspaceResourceIdNil(b bool)`

 SetAzureWorkspaceResourceIdNil sets the value for AzureWorkspaceResourceId to be an explicit nil

### UnsetAzureWorkspaceResourceId
`func (o *DatabricksSqlWarehouseCredentialsIn) UnsetAzureWorkspaceResourceId()`

UnsetAzureWorkspaceResourceId ensures that no value is present for AzureWorkspaceResourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


