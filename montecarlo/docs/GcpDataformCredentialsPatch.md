# GcpDataformCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **NullableString** | Google Cloud project that holds the Dataform repositories. | [optional] 
**Locations** | Pointer to **[]string** | Google Cloud regions to read Dataform repositories in, such as us-central1. | [optional] 
**ServiceAccountKey** | Pointer to **NullableString** | The service account&#39;s JSON key file, as its text. It needs Dataform API access. Stored by Monte Carlo and never returned. | [optional] 

## Methods

### NewGcpDataformCredentialsPatch

`func NewGcpDataformCredentialsPatch() *GcpDataformCredentialsPatch`

NewGcpDataformCredentialsPatch instantiates a new GcpDataformCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpDataformCredentialsPatchWithDefaults

`func NewGcpDataformCredentialsPatchWithDefaults() *GcpDataformCredentialsPatch`

NewGcpDataformCredentialsPatchWithDefaults instantiates a new GcpDataformCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GcpDataformCredentialsPatch) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GcpDataformCredentialsPatch) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GcpDataformCredentialsPatch) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GcpDataformCredentialsPatch) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *GcpDataformCredentialsPatch) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *GcpDataformCredentialsPatch) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetLocations

`func (o *GcpDataformCredentialsPatch) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *GcpDataformCredentialsPatch) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *GcpDataformCredentialsPatch) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *GcpDataformCredentialsPatch) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocationsNil

`func (o *GcpDataformCredentialsPatch) SetLocationsNil(b bool)`

 SetLocationsNil sets the value for Locations to be an explicit nil

### UnsetLocations
`func (o *GcpDataformCredentialsPatch) UnsetLocations()`

UnsetLocations ensures that no value is present for Locations, not even an explicit nil
### GetServiceAccountKey

`func (o *GcpDataformCredentialsPatch) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *GcpDataformCredentialsPatch) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *GcpDataformCredentialsPatch) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.

### HasServiceAccountKey

`func (o *GcpDataformCredentialsPatch) HasServiceAccountKey() bool`

HasServiceAccountKey returns a boolean if a field has been set.

### SetServiceAccountKeyNil

`func (o *GcpDataformCredentialsPatch) SetServiceAccountKeyNil(b bool)`

 SetServiceAccountKeyNil sets the value for ServiceAccountKey to be an explicit nil

### UnsetServiceAccountKey
`func (o *GcpDataformCredentialsPatch) UnsetServiceAccountKey()`

UnsetServiceAccountKey ensures that no value is present for ServiceAccountKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


