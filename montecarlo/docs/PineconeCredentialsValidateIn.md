# PineconeCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Environment** | **string** | Pinecone environment the project runs in, such as &#x60;us-east-1-aws&#x60;. | 
**ProjectId** | **string** | Id of the Pinecone project. | 
**ApiKey** | **string** | API key of the Pinecone project. Used for this check and not kept. | 

## Methods

### NewPineconeCredentialsValidateIn

`func NewPineconeCredentialsValidateIn(deploymentId string, environment string, projectId string, apiKey string, ) *PineconeCredentialsValidateIn`

NewPineconeCredentialsValidateIn instantiates a new PineconeCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPineconeCredentialsValidateInWithDefaults

`func NewPineconeCredentialsValidateInWithDefaults() *PineconeCredentialsValidateIn`

NewPineconeCredentialsValidateInWithDefaults instantiates a new PineconeCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *PineconeCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *PineconeCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *PineconeCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetEnvironment

`func (o *PineconeCredentialsValidateIn) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PineconeCredentialsValidateIn) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PineconeCredentialsValidateIn) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetProjectId

`func (o *PineconeCredentialsValidateIn) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PineconeCredentialsValidateIn) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PineconeCredentialsValidateIn) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetApiKey

`func (o *PineconeCredentialsValidateIn) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *PineconeCredentialsValidateIn) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *PineconeCredentialsValidateIn) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


