# PineconeCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **string** | Pinecone environment the project runs in, such as &#x60;us-east-1-aws&#x60;. | 
**ProjectId** | **string** | Id of the Pinecone project. | 
**ApiKey** | **string** | API key of the Pinecone project. Stored by Monte Carlo and never returned. | 

## Methods

### NewPineconeCredentialsIn

`func NewPineconeCredentialsIn(environment string, projectId string, apiKey string, ) *PineconeCredentialsIn`

NewPineconeCredentialsIn instantiates a new PineconeCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPineconeCredentialsInWithDefaults

`func NewPineconeCredentialsInWithDefaults() *PineconeCredentialsIn`

NewPineconeCredentialsInWithDefaults instantiates a new PineconeCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *PineconeCredentialsIn) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PineconeCredentialsIn) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PineconeCredentialsIn) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetProjectId

`func (o *PineconeCredentialsIn) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PineconeCredentialsIn) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PineconeCredentialsIn) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetApiKey

`func (o *PineconeCredentialsIn) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *PineconeCredentialsIn) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *PineconeCredentialsIn) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


