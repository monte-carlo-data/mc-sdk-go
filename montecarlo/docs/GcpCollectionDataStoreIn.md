# GcpCollectionDataStoreIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment to register the data store on. It must already hold an unregistered Cloud Storage data store. | 
**BucketName** | **string** | Name of the Cloud Storage bucket Monte Carlo should use. | 
**ServiceAccountKey** | **string** | Service account key Monte Carlo reaches the bucket with, as the contents of the JSON key file Google issued for it. It replaces the stored key rather than merging into it. | 
**Name** | Pointer to **NullableString** | Display name for the data store. Replaces the name its deployment gave it. | [optional] 

## Methods

### NewGcpCollectionDataStoreIn

`func NewGcpCollectionDataStoreIn(deploymentId string, bucketName string, serviceAccountKey string, ) *GcpCollectionDataStoreIn`

NewGcpCollectionDataStoreIn instantiates a new GcpCollectionDataStoreIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpCollectionDataStoreInWithDefaults

`func NewGcpCollectionDataStoreInWithDefaults() *GcpCollectionDataStoreIn`

NewGcpCollectionDataStoreInWithDefaults instantiates a new GcpCollectionDataStoreIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *GcpCollectionDataStoreIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GcpCollectionDataStoreIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GcpCollectionDataStoreIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetBucketName

`func (o *GcpCollectionDataStoreIn) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *GcpCollectionDataStoreIn) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *GcpCollectionDataStoreIn) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetServiceAccountKey

`func (o *GcpCollectionDataStoreIn) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *GcpCollectionDataStoreIn) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *GcpCollectionDataStoreIn) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.


### GetName

`func (o *GcpCollectionDataStoreIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpCollectionDataStoreIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpCollectionDataStoreIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpCollectionDataStoreIn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GcpCollectionDataStoreIn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GcpCollectionDataStoreIn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


