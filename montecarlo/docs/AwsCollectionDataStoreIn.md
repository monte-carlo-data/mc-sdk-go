# AwsCollectionDataStoreIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment to register the data store on. It must already hold an unregistered S3 data store. | 
**BucketName** | **string** | Name of the S3 bucket Monte Carlo should use. | 
**RoleArn** | **string** | ARN of the role Monte Carlo assumes to access the bucket. Its trust policy must already carry the deployment&#39;s external id. | 
**Name** | Pointer to **NullableString** | Display name for the data store. Replaces the name its deployment gave it. | [optional] 

## Methods

### NewAwsCollectionDataStoreIn

`func NewAwsCollectionDataStoreIn(deploymentId string, bucketName string, roleArn string, ) *AwsCollectionDataStoreIn`

NewAwsCollectionDataStoreIn instantiates a new AwsCollectionDataStoreIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCollectionDataStoreInWithDefaults

`func NewAwsCollectionDataStoreInWithDefaults() *AwsCollectionDataStoreIn`

NewAwsCollectionDataStoreInWithDefaults instantiates a new AwsCollectionDataStoreIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *AwsCollectionDataStoreIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AwsCollectionDataStoreIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AwsCollectionDataStoreIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetBucketName

`func (o *AwsCollectionDataStoreIn) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *AwsCollectionDataStoreIn) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *AwsCollectionDataStoreIn) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetRoleArn

`func (o *AwsCollectionDataStoreIn) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsCollectionDataStoreIn) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsCollectionDataStoreIn) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### GetName

`func (o *AwsCollectionDataStoreIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsCollectionDataStoreIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsCollectionDataStoreIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsCollectionDataStoreIn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsCollectionDataStoreIn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsCollectionDataStoreIn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


