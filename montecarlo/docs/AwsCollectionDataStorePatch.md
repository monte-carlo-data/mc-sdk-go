# AwsCollectionDataStorePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **NullableString** | Name of the S3 bucket Monte Carlo should use. | [optional] 
**RoleArn** | Pointer to **NullableString** | ARN of the role Monte Carlo assumes to access the bucket. Its trust policy must already carry the deployment&#39;s external id. | [optional] 
**Name** | Pointer to **NullableString** | Display name for the data store. Replaces the name its deployment gave it. | [optional] 

## Methods

### NewAwsCollectionDataStorePatch

`func NewAwsCollectionDataStorePatch() *AwsCollectionDataStorePatch`

NewAwsCollectionDataStorePatch instantiates a new AwsCollectionDataStorePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCollectionDataStorePatchWithDefaults

`func NewAwsCollectionDataStorePatchWithDefaults() *AwsCollectionDataStorePatch`

NewAwsCollectionDataStorePatchWithDefaults instantiates a new AwsCollectionDataStorePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *AwsCollectionDataStorePatch) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *AwsCollectionDataStorePatch) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *AwsCollectionDataStorePatch) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *AwsCollectionDataStorePatch) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### SetBucketNameNil

`func (o *AwsCollectionDataStorePatch) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *AwsCollectionDataStorePatch) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetRoleArn

`func (o *AwsCollectionDataStorePatch) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsCollectionDataStorePatch) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsCollectionDataStorePatch) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *AwsCollectionDataStorePatch) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *AwsCollectionDataStorePatch) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *AwsCollectionDataStorePatch) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetName

`func (o *AwsCollectionDataStorePatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsCollectionDataStorePatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsCollectionDataStorePatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsCollectionDataStorePatch) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsCollectionDataStorePatch) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsCollectionDataStorePatch) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


