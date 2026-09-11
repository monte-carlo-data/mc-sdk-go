# GcpCollectionDataStoreOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the data store. | 
**Name** | Pointer to **NullableString** | Display name of the data store. Null when it has no name. | [optional] 
**DeploymentId** | **string** | Identifier of the deployment this data store belongs to. | 
**StorageType** | [**StorageType**](StorageType.md) | Which kind of storage the data store keeps its data in. | 
**AuthenticationType** | [**AuthenticationType**](AuthenticationType.md) | How Monte Carlo authenticates when it reaches the data store. | 
**Enabled** | **bool** | Whether Monte Carlo is using this data store. One that is unregistered, or whose validation failed, is not enabled. | 
**CreatedTime** | Pointer to **NullableTime** | When the data store was created, which is when its deployment was provisioned. | [optional] 
**LastUpdatedTime** | Pointer to **NullableTime** | When the data store was last registered, renamed, or given different storage or credentials. Null until one of those has happened. | [optional] 
**BucketName** | **string** | Name of the Cloud Storage bucket Monte Carlo uses. Empty until it has been registered. | 

## Methods

### NewGcpCollectionDataStoreOut

`func NewGcpCollectionDataStoreOut(id string, deploymentId string, storageType StorageType, authenticationType AuthenticationType, enabled bool, bucketName string, ) *GcpCollectionDataStoreOut`

NewGcpCollectionDataStoreOut instantiates a new GcpCollectionDataStoreOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpCollectionDataStoreOutWithDefaults

`func NewGcpCollectionDataStoreOutWithDefaults() *GcpCollectionDataStoreOut`

NewGcpCollectionDataStoreOutWithDefaults instantiates a new GcpCollectionDataStoreOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GcpCollectionDataStoreOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GcpCollectionDataStoreOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GcpCollectionDataStoreOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GcpCollectionDataStoreOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpCollectionDataStoreOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpCollectionDataStoreOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpCollectionDataStoreOut) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GcpCollectionDataStoreOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GcpCollectionDataStoreOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeploymentId

`func (o *GcpCollectionDataStoreOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GcpCollectionDataStoreOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GcpCollectionDataStoreOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetStorageType

`func (o *GcpCollectionDataStoreOut) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *GcpCollectionDataStoreOut) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *GcpCollectionDataStoreOut) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.


### GetAuthenticationType

`func (o *GcpCollectionDataStoreOut) GetAuthenticationType() AuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *GcpCollectionDataStoreOut) GetAuthenticationTypeOk() (*AuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *GcpCollectionDataStoreOut) SetAuthenticationType(v AuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetEnabled

`func (o *GcpCollectionDataStoreOut) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GcpCollectionDataStoreOut) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GcpCollectionDataStoreOut) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedTime

`func (o *GcpCollectionDataStoreOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GcpCollectionDataStoreOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GcpCollectionDataStoreOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *GcpCollectionDataStoreOut) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *GcpCollectionDataStoreOut) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *GcpCollectionDataStoreOut) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetLastUpdatedTime

`func (o *GcpCollectionDataStoreOut) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *GcpCollectionDataStoreOut) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *GcpCollectionDataStoreOut) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *GcpCollectionDataStoreOut) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### SetLastUpdatedTimeNil

`func (o *GcpCollectionDataStoreOut) SetLastUpdatedTimeNil(b bool)`

 SetLastUpdatedTimeNil sets the value for LastUpdatedTime to be an explicit nil

### UnsetLastUpdatedTime
`func (o *GcpCollectionDataStoreOut) UnsetLastUpdatedTime()`

UnsetLastUpdatedTime ensures that no value is present for LastUpdatedTime, not even an explicit nil
### GetBucketName

`func (o *GcpCollectionDataStoreOut) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *GcpCollectionDataStoreOut) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *GcpCollectionDataStoreOut) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


