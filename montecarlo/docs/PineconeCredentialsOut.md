# PineconeCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**Environment** | **string** | Pinecone environment the project runs in, such as &#x60;us-east-1-aws&#x60;. | 
**ProjectId** | **string** | Id of the Pinecone project. | 

## Methods

### NewPineconeCredentialsOut

`func NewPineconeCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, environment string, projectId string, ) *PineconeCredentialsOut`

NewPineconeCredentialsOut instantiates a new PineconeCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPineconeCredentialsOutWithDefaults

`func NewPineconeCredentialsOutWithDefaults() *PineconeCredentialsOut`

NewPineconeCredentialsOutWithDefaults instantiates a new PineconeCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PineconeCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PineconeCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PineconeCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *PineconeCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *PineconeCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *PineconeCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *PineconeCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *PineconeCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *PineconeCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *PineconeCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PineconeCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PineconeCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetEnvironment

`func (o *PineconeCredentialsOut) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PineconeCredentialsOut) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PineconeCredentialsOut) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetProjectId

`func (o *PineconeCredentialsOut) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PineconeCredentialsOut) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PineconeCredentialsOut) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


