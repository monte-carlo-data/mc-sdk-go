# AzureCollectionDataStoreOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | [**AuthenticationType**](AuthenticationType.md) | How Monte Carlo authenticates when it reaches the data store. | 
**ContainerName** | **string** | Name of the blob container Monte Carlo uses. Empty until it has been registered. The storage account holding it is part of the credentials, so it is not returned. | 
**CreatedTime** | Pointer to **NullableTime** | When the data store was created, which is when its deployment was provisioned. | [optional] 
**DeploymentId** | **string** | Identifier of the deployment this data store belongs to. | 
**Enabled** | **bool** | Whether Monte Carlo is using this data store. One that is unregistered, or whose validation failed, is not enabled. | 
**Id** | **string** | Unique identifier of the data store. | 
**LastUpdatedTime** | Pointer to **NullableTime** | When the data store was last registered, renamed, or given different storage or credentials. Null until one of those has happened. | [optional] 
**Name** | Pointer to **NullableString** | Display name of the data store. Null when it has no name. | [optional] 
**StorageType** | [**StorageType**](StorageType.md) | Which kind of storage the data store keeps its data in. | 

## Methods

### NewAzureCollectionDataStoreOut

`func NewAzureCollectionDataStoreOut(authenticationType AuthenticationType, containerName string, deploymentId string, enabled bool, id string, storageType StorageType, ) *AzureCollectionDataStoreOut`

NewAzureCollectionDataStoreOut instantiates a new AzureCollectionDataStoreOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCollectionDataStoreOutWithDefaults

`func NewAzureCollectionDataStoreOutWithDefaults() *AzureCollectionDataStoreOut`

NewAzureCollectionDataStoreOutWithDefaults instantiates a new AzureCollectionDataStoreOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *AzureCollectionDataStoreOut) GetAuthenticationType() AuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AzureCollectionDataStoreOut) GetAuthenticationTypeOk() (*AuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AzureCollectionDataStoreOut) SetAuthenticationType(v AuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetContainerName

`func (o *AzureCollectionDataStoreOut) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *AzureCollectionDataStoreOut) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *AzureCollectionDataStoreOut) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetCreatedTime

`func (o *AzureCollectionDataStoreOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AzureCollectionDataStoreOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AzureCollectionDataStoreOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AzureCollectionDataStoreOut) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *AzureCollectionDataStoreOut) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *AzureCollectionDataStoreOut) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetDeploymentId

`func (o *AzureCollectionDataStoreOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AzureCollectionDataStoreOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AzureCollectionDataStoreOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetEnabled

`func (o *AzureCollectionDataStoreOut) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AzureCollectionDataStoreOut) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AzureCollectionDataStoreOut) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *AzureCollectionDataStoreOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureCollectionDataStoreOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureCollectionDataStoreOut) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTime

`func (o *AzureCollectionDataStoreOut) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *AzureCollectionDataStoreOut) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *AzureCollectionDataStoreOut) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *AzureCollectionDataStoreOut) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### SetLastUpdatedTimeNil

`func (o *AzureCollectionDataStoreOut) SetLastUpdatedTimeNil(b bool)`

 SetLastUpdatedTimeNil sets the value for LastUpdatedTime to be an explicit nil

### UnsetLastUpdatedTime
`func (o *AzureCollectionDataStoreOut) UnsetLastUpdatedTime()`

UnsetLastUpdatedTime ensures that no value is present for LastUpdatedTime, not even an explicit nil
### GetName

`func (o *AzureCollectionDataStoreOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureCollectionDataStoreOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureCollectionDataStoreOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureCollectionDataStoreOut) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureCollectionDataStoreOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureCollectionDataStoreOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStorageType

`func (o *AzureCollectionDataStoreOut) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *AzureCollectionDataStoreOut) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *AzureCollectionDataStoreOut) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


