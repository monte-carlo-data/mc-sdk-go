# PineconeCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Pinecone environment the project runs in, such as &#x60;us-east-1-aws&#x60;. Changing it requires sending the credentials&#39; secrets again. | [optional] 
**ProjectId** | Pointer to **NullableString** | Id of the Pinecone project. Changing it requires sending the credentials&#39; secrets again. | [optional] 
**ApiKey** | Pointer to **NullableString** | New API key of the Pinecone project. | [optional] 

## Methods

### NewPineconeCredentialsPatch

`func NewPineconeCredentialsPatch() *PineconeCredentialsPatch`

NewPineconeCredentialsPatch instantiates a new PineconeCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPineconeCredentialsPatchWithDefaults

`func NewPineconeCredentialsPatchWithDefaults() *PineconeCredentialsPatch`

NewPineconeCredentialsPatchWithDefaults instantiates a new PineconeCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *PineconeCredentialsPatch) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PineconeCredentialsPatch) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PineconeCredentialsPatch) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PineconeCredentialsPatch) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *PineconeCredentialsPatch) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *PineconeCredentialsPatch) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetProjectId

`func (o *PineconeCredentialsPatch) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PineconeCredentialsPatch) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PineconeCredentialsPatch) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PineconeCredentialsPatch) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *PineconeCredentialsPatch) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *PineconeCredentialsPatch) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetApiKey

`func (o *PineconeCredentialsPatch) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *PineconeCredentialsPatch) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *PineconeCredentialsPatch) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *PineconeCredentialsPatch) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *PineconeCredentialsPatch) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *PineconeCredentialsPatch) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


