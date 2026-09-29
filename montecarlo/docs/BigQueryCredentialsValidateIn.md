# BigQueryCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**ServiceAccountKey** | **string** | The service account&#39;s JSON key file, as its text. Used for this check and not kept. | 

## Methods

### NewBigQueryCredentialsValidateIn

`func NewBigQueryCredentialsValidateIn(deploymentId string, serviceAccountKey string, ) *BigQueryCredentialsValidateIn`

NewBigQueryCredentialsValidateIn instantiates a new BigQueryCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBigQueryCredentialsValidateInWithDefaults

`func NewBigQueryCredentialsValidateInWithDefaults() *BigQueryCredentialsValidateIn`

NewBigQueryCredentialsValidateInWithDefaults instantiates a new BigQueryCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *BigQueryCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *BigQueryCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *BigQueryCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetServiceAccountKey

`func (o *BigQueryCredentialsValidateIn) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *BigQueryCredentialsValidateIn) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *BigQueryCredentialsValidateIn) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


