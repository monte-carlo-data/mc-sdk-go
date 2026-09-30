# DatabricksMetastoreSqlWarehouseCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**WorkspaceUrl** | **string** | URL of the Databricks workspace, or its host name. | 
**SqlWarehouseId** | **string** | ID of the Databricks SQL warehouse the connection runs on. | 
**WorkspaceId** | **string** | ID of the Databricks workspace. | 
**OauthClientId** | **NullableString** | Client ID of the OAuth service principal. Null for token credentials. | 
**AzureTenantId** | **NullableString** | Microsoft Entra ID tenant of the service principal. Null unless set. | 
**AzureWorkspaceResourceId** | **NullableString** | Azure resource ID of the workspace. Null unless set. | 

## Methods

### NewDatabricksMetastoreSqlWarehouseCredentialsOut

`func NewDatabricksMetastoreSqlWarehouseCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, workspaceUrl string, sqlWarehouseId string, workspaceId string, oauthClientId NullableString, azureTenantId NullableString, azureWorkspaceResourceId NullableString, ) *DatabricksMetastoreSqlWarehouseCredentialsOut`

NewDatabricksMetastoreSqlWarehouseCredentialsOut instantiates a new DatabricksMetastoreSqlWarehouseCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabricksMetastoreSqlWarehouseCredentialsOutWithDefaults

`func NewDatabricksMetastoreSqlWarehouseCredentialsOutWithDefaults() *DatabricksMetastoreSqlWarehouseCredentialsOut`

NewDatabricksMetastoreSqlWarehouseCredentialsOutWithDefaults instantiates a new DatabricksMetastoreSqlWarehouseCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetWorkspaceUrl

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetWorkspaceUrl() string`

GetWorkspaceUrl returns the WorkspaceUrl field if non-nil, zero value otherwise.

### GetWorkspaceUrlOk

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetWorkspaceUrlOk() (*string, bool)`

GetWorkspaceUrlOk returns a tuple with the WorkspaceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceUrl

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetWorkspaceUrl(v string)`

SetWorkspaceUrl sets WorkspaceUrl field to given value.


### GetSqlWarehouseId

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.


### GetWorkspaceId

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetOauthClientId

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.


### SetOauthClientIdNil

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetOauthClientIdNil(b bool)`

 SetOauthClientIdNil sets the value for OauthClientId to be an explicit nil

### UnsetOauthClientId
`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) UnsetOauthClientId()`

UnsetOauthClientId ensures that no value is present for OauthClientId, not even an explicit nil
### GetAzureTenantId

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.


### SetAzureTenantIdNil

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetAzureTenantIdNil(b bool)`

 SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil

### UnsetAzureTenantId
`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) UnsetAzureTenantId()`

UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
### GetAzureWorkspaceResourceId

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetAzureWorkspaceResourceId() string`

GetAzureWorkspaceResourceId returns the AzureWorkspaceResourceId field if non-nil, zero value otherwise.

### GetAzureWorkspaceResourceIdOk

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) GetAzureWorkspaceResourceIdOk() (*string, bool)`

GetAzureWorkspaceResourceIdOk returns a tuple with the AzureWorkspaceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureWorkspaceResourceId

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetAzureWorkspaceResourceId(v string)`

SetAzureWorkspaceResourceId sets AzureWorkspaceResourceId field to given value.


### SetAzureWorkspaceResourceIdNil

`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) SetAzureWorkspaceResourceIdNil(b bool)`

 SetAzureWorkspaceResourceIdNil sets the value for AzureWorkspaceResourceId to be an explicit nil

### UnsetAzureWorkspaceResourceId
`func (o *DatabricksMetastoreSqlWarehouseCredentialsOut) UnsetAzureWorkspaceResourceId()`

UnsetAzureWorkspaceResourceId ensures that no value is present for AzureWorkspaceResourceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


