# PineconeCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Environment** | Pointer to **NullableString** | Pinecone environment the project runs in, such as &#x60;us-east-1-aws&#x60;. Changing it requires sending the credentials&#39; secrets again. | [optional] 
**ProjectId** | Pointer to **NullableString** | Id of the Pinecone project. Changing it requires sending the credentials&#39; secrets again. | [optional] 
**ApiKey** | Pointer to **NullableString** | New API key of the Pinecone project. | [optional] 

## Methods

### NewPineconeCredentialsChangeValidateIn

`func NewPineconeCredentialsChangeValidateIn(deploymentId string, ) *PineconeCredentialsChangeValidateIn`

NewPineconeCredentialsChangeValidateIn instantiates a new PineconeCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPineconeCredentialsChangeValidateInWithDefaults

`func NewPineconeCredentialsChangeValidateInWithDefaults() *PineconeCredentialsChangeValidateIn`

NewPineconeCredentialsChangeValidateInWithDefaults instantiates a new PineconeCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *PineconeCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *PineconeCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *PineconeCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetEnvironment

`func (o *PineconeCredentialsChangeValidateIn) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PineconeCredentialsChangeValidateIn) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PineconeCredentialsChangeValidateIn) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PineconeCredentialsChangeValidateIn) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *PineconeCredentialsChangeValidateIn) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *PineconeCredentialsChangeValidateIn) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetProjectId

`func (o *PineconeCredentialsChangeValidateIn) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PineconeCredentialsChangeValidateIn) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PineconeCredentialsChangeValidateIn) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PineconeCredentialsChangeValidateIn) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *PineconeCredentialsChangeValidateIn) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *PineconeCredentialsChangeValidateIn) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetApiKey

`func (o *PineconeCredentialsChangeValidateIn) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *PineconeCredentialsChangeValidateIn) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *PineconeCredentialsChangeValidateIn) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *PineconeCredentialsChangeValidateIn) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *PineconeCredentialsChangeValidateIn) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *PineconeCredentialsChangeValidateIn) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


