# CollectionDataStoreDetailOut

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
**Platform** | [**RuntimePlatform**](RuntimePlatform.md) | Where the data store runs. | 
**Endpoint** | **string** | Address of the data store, in the form its platform uses. On AWS that is an S3 bucket name, on Azure the name of a blob container, and on GCP a Cloud Storage bucket name. Empty until it has been registered. | 
**Details** | **map[string]interface{}** | The fields the platform-specific read of this data store returns beyond the ones above, such as &#x60;bucket_name&#x60; and &#x60;external_id&#x60; on AWS. Never a secret. | 

## Methods

### NewCollectionDataStoreDetailOut

`func NewCollectionDataStoreDetailOut(id string, deploymentId string, storageType StorageType, authenticationType AuthenticationType, enabled bool, platform RuntimePlatform, endpoint string, details map[string]interface{}, ) *CollectionDataStoreDetailOut`

NewCollectionDataStoreDetailOut instantiates a new CollectionDataStoreDetailOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionDataStoreDetailOutWithDefaults

`func NewCollectionDataStoreDetailOutWithDefaults() *CollectionDataStoreDetailOut`

NewCollectionDataStoreDetailOutWithDefaults instantiates a new CollectionDataStoreDetailOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CollectionDataStoreDetailOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionDataStoreDetailOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionDataStoreDetailOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CollectionDataStoreDetailOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionDataStoreDetailOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionDataStoreDetailOut) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CollectionDataStoreDetailOut) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CollectionDataStoreDetailOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CollectionDataStoreDetailOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeploymentId

`func (o *CollectionDataStoreDetailOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *CollectionDataStoreDetailOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *CollectionDataStoreDetailOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetStorageType

`func (o *CollectionDataStoreDetailOut) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *CollectionDataStoreDetailOut) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *CollectionDataStoreDetailOut) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.


### GetAuthenticationType

`func (o *CollectionDataStoreDetailOut) GetAuthenticationType() AuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *CollectionDataStoreDetailOut) GetAuthenticationTypeOk() (*AuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *CollectionDataStoreDetailOut) SetAuthenticationType(v AuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetEnabled

`func (o *CollectionDataStoreDetailOut) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CollectionDataStoreDetailOut) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CollectionDataStoreDetailOut) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedTime

`func (o *CollectionDataStoreDetailOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CollectionDataStoreDetailOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CollectionDataStoreDetailOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CollectionDataStoreDetailOut) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *CollectionDataStoreDetailOut) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *CollectionDataStoreDetailOut) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetLastUpdatedTime

`func (o *CollectionDataStoreDetailOut) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *CollectionDataStoreDetailOut) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *CollectionDataStoreDetailOut) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *CollectionDataStoreDetailOut) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### SetLastUpdatedTimeNil

`func (o *CollectionDataStoreDetailOut) SetLastUpdatedTimeNil(b bool)`

 SetLastUpdatedTimeNil sets the value for LastUpdatedTime to be an explicit nil

### UnsetLastUpdatedTime
`func (o *CollectionDataStoreDetailOut) UnsetLastUpdatedTime()`

UnsetLastUpdatedTime ensures that no value is present for LastUpdatedTime, not even an explicit nil
### GetPlatform

`func (o *CollectionDataStoreDetailOut) GetPlatform() RuntimePlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CollectionDataStoreDetailOut) GetPlatformOk() (*RuntimePlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CollectionDataStoreDetailOut) SetPlatform(v RuntimePlatform)`

SetPlatform sets Platform field to given value.


### GetEndpoint

`func (o *CollectionDataStoreDetailOut) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CollectionDataStoreDetailOut) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CollectionDataStoreDetailOut) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetDetails

`func (o *CollectionDataStoreDetailOut) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CollectionDataStoreDetailOut) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CollectionDataStoreDetailOut) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


