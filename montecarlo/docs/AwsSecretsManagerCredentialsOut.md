# AwsSecretsManagerCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**BqProjectId** | **NullableString** | BigQuery project the connection reads from. Null unless set. | 
**SqlWarehouseId** | **NullableString** | Databricks SQL warehouse the connection runs queries on. Null unless set. | 
**AwsSecret** | **string** | Name or ARN of the AWS Secrets Manager secret holding the connection&#39;s credentials. | 
**AwsRegion** | **NullableString** | AWS region of the secret. Null when unset. | 
**AssumableRole** | **NullableString** | ARN of the role the deployment assumes to read the secret. Null when unset. | 
**ExternalId** | **NullableString** | External id presented when assuming the role. Null when unset. | 

## Methods

### NewAwsSecretsManagerCredentialsOut

`func NewAwsSecretsManagerCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, bqProjectId NullableString, sqlWarehouseId NullableString, awsSecret string, awsRegion NullableString, assumableRole NullableString, externalId NullableString, ) *AwsSecretsManagerCredentialsOut`

NewAwsSecretsManagerCredentialsOut instantiates a new AwsSecretsManagerCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSecretsManagerCredentialsOutWithDefaults

`func NewAwsSecretsManagerCredentialsOutWithDefaults() *AwsSecretsManagerCredentialsOut`

NewAwsSecretsManagerCredentialsOutWithDefaults instantiates a new AwsSecretsManagerCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AwsSecretsManagerCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsSecretsManagerCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsSecretsManagerCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *AwsSecretsManagerCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *AwsSecretsManagerCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *AwsSecretsManagerCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *AwsSecretsManagerCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *AwsSecretsManagerCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *AwsSecretsManagerCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *AwsSecretsManagerCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AwsSecretsManagerCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AwsSecretsManagerCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetBqProjectId

`func (o *AwsSecretsManagerCredentialsOut) GetBqProjectId() string`

GetBqProjectId returns the BqProjectId field if non-nil, zero value otherwise.

### GetBqProjectIdOk

`func (o *AwsSecretsManagerCredentialsOut) GetBqProjectIdOk() (*string, bool)`

GetBqProjectIdOk returns a tuple with the BqProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBqProjectId

`func (o *AwsSecretsManagerCredentialsOut) SetBqProjectId(v string)`

SetBqProjectId sets BqProjectId field to given value.


### SetBqProjectIdNil

`func (o *AwsSecretsManagerCredentialsOut) SetBqProjectIdNil(b bool)`

 SetBqProjectIdNil sets the value for BqProjectId to be an explicit nil

### UnsetBqProjectId
`func (o *AwsSecretsManagerCredentialsOut) UnsetBqProjectId()`

UnsetBqProjectId ensures that no value is present for BqProjectId, not even an explicit nil
### GetSqlWarehouseId

`func (o *AwsSecretsManagerCredentialsOut) GetSqlWarehouseId() string`

GetSqlWarehouseId returns the SqlWarehouseId field if non-nil, zero value otherwise.

### GetSqlWarehouseIdOk

`func (o *AwsSecretsManagerCredentialsOut) GetSqlWarehouseIdOk() (*string, bool)`

GetSqlWarehouseIdOk returns a tuple with the SqlWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlWarehouseId

`func (o *AwsSecretsManagerCredentialsOut) SetSqlWarehouseId(v string)`

SetSqlWarehouseId sets SqlWarehouseId field to given value.


### SetSqlWarehouseIdNil

`func (o *AwsSecretsManagerCredentialsOut) SetSqlWarehouseIdNil(b bool)`

 SetSqlWarehouseIdNil sets the value for SqlWarehouseId to be an explicit nil

### UnsetSqlWarehouseId
`func (o *AwsSecretsManagerCredentialsOut) UnsetSqlWarehouseId()`

UnsetSqlWarehouseId ensures that no value is present for SqlWarehouseId, not even an explicit nil
### GetAwsSecret

`func (o *AwsSecretsManagerCredentialsOut) GetAwsSecret() string`

GetAwsSecret returns the AwsSecret field if non-nil, zero value otherwise.

### GetAwsSecretOk

`func (o *AwsSecretsManagerCredentialsOut) GetAwsSecretOk() (*string, bool)`

GetAwsSecretOk returns a tuple with the AwsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecret

`func (o *AwsSecretsManagerCredentialsOut) SetAwsSecret(v string)`

SetAwsSecret sets AwsSecret field to given value.


### GetAwsRegion

`func (o *AwsSecretsManagerCredentialsOut) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AwsSecretsManagerCredentialsOut) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AwsSecretsManagerCredentialsOut) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### SetAwsRegionNil

`func (o *AwsSecretsManagerCredentialsOut) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *AwsSecretsManagerCredentialsOut) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetAssumableRole

`func (o *AwsSecretsManagerCredentialsOut) GetAssumableRole() string`

GetAssumableRole returns the AssumableRole field if non-nil, zero value otherwise.

### GetAssumableRoleOk

`func (o *AwsSecretsManagerCredentialsOut) GetAssumableRoleOk() (*string, bool)`

GetAssumableRoleOk returns a tuple with the AssumableRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumableRole

`func (o *AwsSecretsManagerCredentialsOut) SetAssumableRole(v string)`

SetAssumableRole sets AssumableRole field to given value.


### SetAssumableRoleNil

`func (o *AwsSecretsManagerCredentialsOut) SetAssumableRoleNil(b bool)`

 SetAssumableRoleNil sets the value for AssumableRole to be an explicit nil

### UnsetAssumableRole
`func (o *AwsSecretsManagerCredentialsOut) UnsetAssumableRole()`

UnsetAssumableRole ensures that no value is present for AssumableRole, not even an explicit nil
### GetExternalId

`func (o *AwsSecretsManagerCredentialsOut) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsSecretsManagerCredentialsOut) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsSecretsManagerCredentialsOut) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *AwsSecretsManagerCredentialsOut) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AwsSecretsManagerCredentialsOut) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


