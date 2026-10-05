# GcpDataformCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Google Cloud project that holds the Dataform repositories. | 
**Locations** | **[]string** | Google Cloud regions to read Dataform repositories in, such as us-central1. | 
**ServiceAccountKey** | **string** | The service account&#39;s JSON key file, as its text. It needs Dataform API access. Stored by Monte Carlo and never returned. | 

## Methods

### NewGcpDataformCredentialsIn

`func NewGcpDataformCredentialsIn(projectId string, locations []string, serviceAccountKey string, ) *GcpDataformCredentialsIn`

NewGcpDataformCredentialsIn instantiates a new GcpDataformCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpDataformCredentialsInWithDefaults

`func NewGcpDataformCredentialsInWithDefaults() *GcpDataformCredentialsIn`

NewGcpDataformCredentialsInWithDefaults instantiates a new GcpDataformCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GcpDataformCredentialsIn) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GcpDataformCredentialsIn) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GcpDataformCredentialsIn) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetLocations

`func (o *GcpDataformCredentialsIn) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *GcpDataformCredentialsIn) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *GcpDataformCredentialsIn) SetLocations(v []string)`

SetLocations sets Locations field to given value.


### GetServiceAccountKey

`func (o *GcpDataformCredentialsIn) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *GcpDataformCredentialsIn) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *GcpDataformCredentialsIn) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


