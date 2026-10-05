# GcpDataformCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**ProjectId** | **string** | Google Cloud project that holds the Dataform repositories. | 
**Locations** | **[]string** | Google Cloud regions to read Dataform repositories in, such as us-central1. | 
**ClientEmail** | **string** | Email address of the service account. | 

## Methods

### NewGcpDataformCredentialsOut

`func NewGcpDataformCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, projectId string, locations []string, clientEmail string, ) *GcpDataformCredentialsOut`

NewGcpDataformCredentialsOut instantiates a new GcpDataformCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpDataformCredentialsOutWithDefaults

`func NewGcpDataformCredentialsOutWithDefaults() *GcpDataformCredentialsOut`

NewGcpDataformCredentialsOutWithDefaults instantiates a new GcpDataformCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GcpDataformCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GcpDataformCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GcpDataformCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *GcpDataformCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GcpDataformCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GcpDataformCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *GcpDataformCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *GcpDataformCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *GcpDataformCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *GcpDataformCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GcpDataformCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GcpDataformCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetProjectId

`func (o *GcpDataformCredentialsOut) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GcpDataformCredentialsOut) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GcpDataformCredentialsOut) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetLocations

`func (o *GcpDataformCredentialsOut) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *GcpDataformCredentialsOut) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *GcpDataformCredentialsOut) SetLocations(v []string)`

SetLocations sets Locations field to given value.


### GetClientEmail

`func (o *GcpDataformCredentialsOut) GetClientEmail() string`

GetClientEmail returns the ClientEmail field if non-nil, zero value otherwise.

### GetClientEmailOk

`func (o *GcpDataformCredentialsOut) GetClientEmailOk() (*string, bool)`

GetClientEmailOk returns a tuple with the ClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmail

`func (o *GcpDataformCredentialsOut) SetClientEmail(v string)`

SetClientEmail sets ClientEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


