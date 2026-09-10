# GcpCollectionDataStorePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **NullableString** | Name of the Cloud Storage bucket Monte Carlo should use. | [optional] 
**Name** | Pointer to **NullableString** | Display name for the data store. Replaces the name its deployment gave it. | [optional] 
**ServiceAccountKey** | Pointer to **NullableString** | Service account key Monte Carlo reaches the bucket with, as the contents of the JSON key file Google issued for it. It replaces the stored key rather than merging into it. | [optional] 

## Methods

### NewGcpCollectionDataStorePatch

`func NewGcpCollectionDataStorePatch() *GcpCollectionDataStorePatch`

NewGcpCollectionDataStorePatch instantiates a new GcpCollectionDataStorePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpCollectionDataStorePatchWithDefaults

`func NewGcpCollectionDataStorePatchWithDefaults() *GcpCollectionDataStorePatch`

NewGcpCollectionDataStorePatchWithDefaults instantiates a new GcpCollectionDataStorePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *GcpCollectionDataStorePatch) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *GcpCollectionDataStorePatch) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *GcpCollectionDataStorePatch) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *GcpCollectionDataStorePatch) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### SetBucketNameNil

`func (o *GcpCollectionDataStorePatch) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *GcpCollectionDataStorePatch) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetName

`func (o *GcpCollectionDataStorePatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpCollectionDataStorePatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpCollectionDataStorePatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpCollectionDataStorePatch) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GcpCollectionDataStorePatch) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GcpCollectionDataStorePatch) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetServiceAccountKey

`func (o *GcpCollectionDataStorePatch) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *GcpCollectionDataStorePatch) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *GcpCollectionDataStorePatch) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.

### HasServiceAccountKey

`func (o *GcpCollectionDataStorePatch) HasServiceAccountKey() bool`

HasServiceAccountKey returns a boolean if a field has been set.

### SetServiceAccountKeyNil

`func (o *GcpCollectionDataStorePatch) SetServiceAccountKeyNil(b bool)`

 SetServiceAccountKeyNil sets the value for ServiceAccountKey to be an explicit nil

### UnsetServiceAccountKey
`func (o *GcpCollectionDataStorePatch) UnsetServiceAccountKey()`

UnsetServiceAccountKey ensures that no value is present for ServiceAccountKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


