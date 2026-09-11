# AwsCollectionDataStoreOut

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
**BucketName** | **string** | Name of the S3 bucket Monte Carlo uses. Empty until it has been registered. | 
**ExternalId** | Pointer to **NullableString** | Value to put in the trust policy of the role Monte Carlo assumes to access the bucket. Null before Monte Carlo has generated one, and for a caller who cannot register a data store. Also null if the value could not be read just now, so retry once before treating it as absent. | [optional] 

## Methods

### NewAwsCollectionDataStoreOut

`func NewAwsCollectionDataStoreOut(id string, deploymentId string, storageType StorageType, authenticationType AuthenticationType, enabled bool, bucketName string, ) *AwsCollectionDataStoreOut`

NewAwsCollectionDataStoreOut instantiates a new AwsCollectionDataStoreOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCollectionDataStoreOutWithDefaults

`func NewAwsCollectionDataStoreOutWithDefaults() *AwsCollectionDataStoreOut`

NewAwsCollectionDataStoreOutWithDefaults instantiates a new AwsCollectionDataStoreOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AwsCollectionDataStoreOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsCollectionDataStoreOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsCollectionDataStoreOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AwsCollectionDataStoreOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsCollectionDataStoreOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsCollectionDataStoreOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsCollectionDataStoreOut) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsCollectionDataStoreOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsCollectionDataStoreOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeploymentId

`func (o *AwsCollectionDataStoreOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AwsCollectionDataStoreOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AwsCollectionDataStoreOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetStorageType

`func (o *AwsCollectionDataStoreOut) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *AwsCollectionDataStoreOut) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *AwsCollectionDataStoreOut) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.


### GetAuthenticationType

`func (o *AwsCollectionDataStoreOut) GetAuthenticationType() AuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AwsCollectionDataStoreOut) GetAuthenticationTypeOk() (*AuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AwsCollectionDataStoreOut) SetAuthenticationType(v AuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetEnabled

`func (o *AwsCollectionDataStoreOut) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AwsCollectionDataStoreOut) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AwsCollectionDataStoreOut) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedTime

`func (o *AwsCollectionDataStoreOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AwsCollectionDataStoreOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AwsCollectionDataStoreOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AwsCollectionDataStoreOut) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *AwsCollectionDataStoreOut) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *AwsCollectionDataStoreOut) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetLastUpdatedTime

`func (o *AwsCollectionDataStoreOut) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AwsCollectionDataStoreOut) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AwsCollectionDataStoreOut) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AwsCollectionDataStoreOut) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### SetLastUpdatedTimeNil

`func (o *AwsCollectionDataStoreOut) SetLastUpdatedTimeNil(b bool)`

 SetLastUpdatedTimeNil sets the value for LastUpdatedTime to be an explicit nil

### UnsetLastUpdatedTime
`func (o *AwsCollectionDataStoreOut) UnsetLastUpdatedTime()`

UnsetLastUpdatedTime ensures that no value is present for LastUpdatedTime, not even an explicit nil
### GetBucketName

`func (o *AwsCollectionDataStoreOut) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *AwsCollectionDataStoreOut) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *AwsCollectionDataStoreOut) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetExternalId

`func (o *AwsCollectionDataStoreOut) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsCollectionDataStoreOut) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsCollectionDataStoreOut) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AwsCollectionDataStoreOut) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AwsCollectionDataStoreOut) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AwsCollectionDataStoreOut) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


