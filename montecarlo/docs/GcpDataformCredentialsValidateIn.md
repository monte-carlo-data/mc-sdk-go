# GcpDataformCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**ProjectId** | **string** | Google Cloud project that holds the Dataform repositories. | 
**Locations** | **[]string** | Google Cloud regions to read Dataform repositories in, such as us-central1. | 
**ServiceAccountKey** | **string** | The service account&#39;s JSON key file, as its text. It needs Dataform API access. Used for this check and not kept. | 

## Methods

### NewGcpDataformCredentialsValidateIn

`func NewGcpDataformCredentialsValidateIn(deploymentId string, projectId string, locations []string, serviceAccountKey string, ) *GcpDataformCredentialsValidateIn`

NewGcpDataformCredentialsValidateIn instantiates a new GcpDataformCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpDataformCredentialsValidateInWithDefaults

`func NewGcpDataformCredentialsValidateInWithDefaults() *GcpDataformCredentialsValidateIn`

NewGcpDataformCredentialsValidateInWithDefaults instantiates a new GcpDataformCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *GcpDataformCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GcpDataformCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GcpDataformCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetProjectId

`func (o *GcpDataformCredentialsValidateIn) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GcpDataformCredentialsValidateIn) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GcpDataformCredentialsValidateIn) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetLocations

`func (o *GcpDataformCredentialsValidateIn) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *GcpDataformCredentialsValidateIn) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *GcpDataformCredentialsValidateIn) SetLocations(v []string)`

SetLocations sets Locations field to given value.


### GetServiceAccountKey

`func (o *GcpDataformCredentialsValidateIn) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *GcpDataformCredentialsValidateIn) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *GcpDataformCredentialsValidateIn) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


