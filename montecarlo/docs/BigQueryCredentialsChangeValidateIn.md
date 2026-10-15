# BigQueryCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**ServiceAccountKey** | Pointer to **NullableString** | New JSON key file, as its text. Replaces the stored key whole. | [optional] 

## Methods

### NewBigQueryCredentialsChangeValidateIn

`func NewBigQueryCredentialsChangeValidateIn(deploymentId string, ) *BigQueryCredentialsChangeValidateIn`

NewBigQueryCredentialsChangeValidateIn instantiates a new BigQueryCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBigQueryCredentialsChangeValidateInWithDefaults

`func NewBigQueryCredentialsChangeValidateInWithDefaults() *BigQueryCredentialsChangeValidateIn`

NewBigQueryCredentialsChangeValidateInWithDefaults instantiates a new BigQueryCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *BigQueryCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *BigQueryCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *BigQueryCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetServiceAccountKey

`func (o *BigQueryCredentialsChangeValidateIn) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *BigQueryCredentialsChangeValidateIn) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *BigQueryCredentialsChangeValidateIn) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.

### HasServiceAccountKey

`func (o *BigQueryCredentialsChangeValidateIn) HasServiceAccountKey() bool`

HasServiceAccountKey returns a boolean if a field has been set.

### SetServiceAccountKeyNil

`func (o *BigQueryCredentialsChangeValidateIn) SetServiceAccountKeyNil(b bool)`

 SetServiceAccountKeyNil sets the value for ServiceAccountKey to be an explicit nil

### UnsetServiceAccountKey
`func (o *BigQueryCredentialsChangeValidateIn) UnsetServiceAccountKey()`

UnsetServiceAccountKey ensures that no value is present for ServiceAccountKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


