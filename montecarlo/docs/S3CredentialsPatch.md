# S3CredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssumableRole** | Pointer to **NullableString** | ARN of the IAM role in your AWS account that the deployment assumes. | [optional] 
**ExternalId** | Pointer to **NullableString** | External id your role&#39;s trust policy requires, if it requires one. It is not the deployment&#39;s &#x60;aws_external_id&#x60;. | [optional] 
**Bucket** | Pointer to **NullableString** | S3 bucket Monte Carlo reads from. | [optional] 
**Prefix** | Pointer to **NullableString** | Key prefix the objects sit under, such as &#x60;emr/prod&#x60;. Leave out the bucket, a leading &#x60;/&#x60; and any &#x60;s3://&#x60; scheme. | [optional] 

## Methods

### NewS3CredentialsPatch

`func NewS3CredentialsPatch() *S3CredentialsPatch`

NewS3CredentialsPatch instantiates a new S3CredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CredentialsPatchWithDefaults

`func NewS3CredentialsPatchWithDefaults() *S3CredentialsPatch`

NewS3CredentialsPatchWithDefaults instantiates a new S3CredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssumableRole

`func (o *S3CredentialsPatch) GetAssumableRole() string`

GetAssumableRole returns the AssumableRole field if non-nil, zero value otherwise.

### GetAssumableRoleOk

`func (o *S3CredentialsPatch) GetAssumableRoleOk() (*string, bool)`

GetAssumableRoleOk returns a tuple with the AssumableRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumableRole

`func (o *S3CredentialsPatch) SetAssumableRole(v string)`

SetAssumableRole sets AssumableRole field to given value.

### HasAssumableRole

`func (o *S3CredentialsPatch) HasAssumableRole() bool`

HasAssumableRole returns a boolean if a field has been set.

### SetAssumableRoleNil

`func (o *S3CredentialsPatch) SetAssumableRoleNil(b bool)`

 SetAssumableRoleNil sets the value for AssumableRole to be an explicit nil

### UnsetAssumableRole
`func (o *S3CredentialsPatch) UnsetAssumableRole()`

UnsetAssumableRole ensures that no value is present for AssumableRole, not even an explicit nil
### GetExternalId

`func (o *S3CredentialsPatch) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *S3CredentialsPatch) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *S3CredentialsPatch) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *S3CredentialsPatch) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *S3CredentialsPatch) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *S3CredentialsPatch) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetBucket

`func (o *S3CredentialsPatch) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *S3CredentialsPatch) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *S3CredentialsPatch) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *S3CredentialsPatch) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### SetBucketNil

`func (o *S3CredentialsPatch) SetBucketNil(b bool)`

 SetBucketNil sets the value for Bucket to be an explicit nil

### UnsetBucket
`func (o *S3CredentialsPatch) UnsetBucket()`

UnsetBucket ensures that no value is present for Bucket, not even an explicit nil
### GetPrefix

`func (o *S3CredentialsPatch) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *S3CredentialsPatch) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *S3CredentialsPatch) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *S3CredentialsPatch) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *S3CredentialsPatch) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *S3CredentialsPatch) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


