# SnowflakeCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**Account** | **string** | Snowflake account identifier, without the host suffix. | 
**User** | **string** | Snowflake user the key pair belongs to. | 
**Warehouse** | **NullableString** | Snowflake virtual warehouse queries run in. Null when none is set. | 

## Methods

### NewSnowflakeCredentialsOut

`func NewSnowflakeCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, account string, user string, warehouse NullableString, ) *SnowflakeCredentialsOut`

NewSnowflakeCredentialsOut instantiates a new SnowflakeCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeCredentialsOutWithDefaults

`func NewSnowflakeCredentialsOutWithDefaults() *SnowflakeCredentialsOut`

NewSnowflakeCredentialsOutWithDefaults instantiates a new SnowflakeCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnowflakeCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnowflakeCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnowflakeCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *SnowflakeCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SnowflakeCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SnowflakeCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *SnowflakeCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *SnowflakeCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *SnowflakeCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *SnowflakeCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SnowflakeCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SnowflakeCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetAccount

`func (o *SnowflakeCredentialsOut) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SnowflakeCredentialsOut) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SnowflakeCredentialsOut) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetUser

`func (o *SnowflakeCredentialsOut) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SnowflakeCredentialsOut) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SnowflakeCredentialsOut) SetUser(v string)`

SetUser sets User field to given value.


### GetWarehouse

`func (o *SnowflakeCredentialsOut) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *SnowflakeCredentialsOut) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *SnowflakeCredentialsOut) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.


### SetWarehouseNil

`func (o *SnowflakeCredentialsOut) SetWarehouseNil(b bool)`

 SetWarehouseNil sets the value for Warehouse to be an explicit nil

### UnsetWarehouse
`func (o *SnowflakeCredentialsOut) UnsetWarehouse()`

UnsetWarehouse ensures that no value is present for Warehouse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


