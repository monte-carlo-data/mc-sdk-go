# HiveS3CredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssumableRole** | **string** | ARN of the IAM role in your AWS account that the deployment assumes. | 
**ExternalId** | Pointer to **NullableString** | External id your role&#39;s trust policy requires, if it requires one. It is not the deployment&#39;s &#x60;aws_external_id&#x60;. | [optional] 
**Bucket** | **string** | S3 bucket Monte Carlo reads from. | 
**Prefix** | **string** | Key prefix of the EMR log root, such as &#x60;emr/prod&#x60;. Leave out the bucket, a leading &#x60;/&#x60; and any &#x60;s3://&#x60; scheme. Send an empty string when the cluster&#39;s EMR log URI is the bucket root. Monte Carlo reads each cluster&#39;s Hive logs under &#x60;&lt;prefix&gt;/&lt;cluster-id&gt;/node/&lt;instance-id&gt;/applications/hive/&#x60;. | 

## Methods

### NewHiveS3CredentialsIn

`func NewHiveS3CredentialsIn(assumableRole string, bucket string, prefix string, ) *HiveS3CredentialsIn`

NewHiveS3CredentialsIn instantiates a new HiveS3CredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveS3CredentialsInWithDefaults

`func NewHiveS3CredentialsInWithDefaults() *HiveS3CredentialsIn`

NewHiveS3CredentialsInWithDefaults instantiates a new HiveS3CredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssumableRole

`func (o *HiveS3CredentialsIn) GetAssumableRole() string`

GetAssumableRole returns the AssumableRole field if non-nil, zero value otherwise.

### GetAssumableRoleOk

`func (o *HiveS3CredentialsIn) GetAssumableRoleOk() (*string, bool)`

GetAssumableRoleOk returns a tuple with the AssumableRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumableRole

`func (o *HiveS3CredentialsIn) SetAssumableRole(v string)`

SetAssumableRole sets AssumableRole field to given value.


### GetExternalId

`func (o *HiveS3CredentialsIn) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *HiveS3CredentialsIn) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *HiveS3CredentialsIn) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *HiveS3CredentialsIn) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *HiveS3CredentialsIn) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *HiveS3CredentialsIn) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetBucket

`func (o *HiveS3CredentialsIn) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *HiveS3CredentialsIn) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *HiveS3CredentialsIn) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetPrefix

`func (o *HiveS3CredentialsIn) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *HiveS3CredentialsIn) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *HiveS3CredentialsIn) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


