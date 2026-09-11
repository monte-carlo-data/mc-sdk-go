# GcpCollectionAgentPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccountKey** | Pointer to **NullableString** | Credentials for &#x60;GCP_JSON_SERVICE_ACCOUNT_KEY&#x60;, as the contents of the JSON key file Google issued for the service account. Send this or &#x60;auth_headers&#x60;, never both. It replaces the stored credentials rather than merging into them. | [optional] 
**AuthHeaders** | Pointer to [**NullableAuthHeadersCredentialsIn**](AuthHeadersCredentialsIn.md) | Credentials for &#x60;CUSTOM_AUTH_HEADERS&#x60;. Send this or &#x60;service_account_key&#x60;, never both. It replaces the stored credentials rather than merging into them. | [optional] 
**CloudRunUrl** | Pointer to **NullableString** | URL of the Cloud Run service Monte Carlo should call. | [optional] 
**Name** | Pointer to **NullableString** | Display name for the collection agent. Replaces the name it currently has. | [optional] 
**AuthenticationType** | Pointer to [**NullableGcpAgentAuthenticationType**](GcpAgentAuthenticationType.md) | How Monte Carlo authenticates when it calls the agent. Send it together with the matching credentials. | [optional] 

## Methods

### NewGcpCollectionAgentPatch

`func NewGcpCollectionAgentPatch() *GcpCollectionAgentPatch`

NewGcpCollectionAgentPatch instantiates a new GcpCollectionAgentPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpCollectionAgentPatchWithDefaults

`func NewGcpCollectionAgentPatchWithDefaults() *GcpCollectionAgentPatch`

NewGcpCollectionAgentPatchWithDefaults instantiates a new GcpCollectionAgentPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccountKey

`func (o *GcpCollectionAgentPatch) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *GcpCollectionAgentPatch) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *GcpCollectionAgentPatch) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.

### HasServiceAccountKey

`func (o *GcpCollectionAgentPatch) HasServiceAccountKey() bool`

HasServiceAccountKey returns a boolean if a field has been set.

### SetServiceAccountKeyNil

`func (o *GcpCollectionAgentPatch) SetServiceAccountKeyNil(b bool)`

 SetServiceAccountKeyNil sets the value for ServiceAccountKey to be an explicit nil

### UnsetServiceAccountKey
`func (o *GcpCollectionAgentPatch) UnsetServiceAccountKey()`

UnsetServiceAccountKey ensures that no value is present for ServiceAccountKey, not even an explicit nil
### GetAuthHeaders

`func (o *GcpCollectionAgentPatch) GetAuthHeaders() AuthHeadersCredentialsIn`

GetAuthHeaders returns the AuthHeaders field if non-nil, zero value otherwise.

### GetAuthHeadersOk

`func (o *GcpCollectionAgentPatch) GetAuthHeadersOk() (*AuthHeadersCredentialsIn, bool)`

GetAuthHeadersOk returns a tuple with the AuthHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeaders

`func (o *GcpCollectionAgentPatch) SetAuthHeaders(v AuthHeadersCredentialsIn)`

SetAuthHeaders sets AuthHeaders field to given value.

### HasAuthHeaders

`func (o *GcpCollectionAgentPatch) HasAuthHeaders() bool`

HasAuthHeaders returns a boolean if a field has been set.

### SetAuthHeadersNil

`func (o *GcpCollectionAgentPatch) SetAuthHeadersNil(b bool)`

 SetAuthHeadersNil sets the value for AuthHeaders to be an explicit nil

### UnsetAuthHeaders
`func (o *GcpCollectionAgentPatch) UnsetAuthHeaders()`

UnsetAuthHeaders ensures that no value is present for AuthHeaders, not even an explicit nil
### GetCloudRunUrl

`func (o *GcpCollectionAgentPatch) GetCloudRunUrl() string`

GetCloudRunUrl returns the CloudRunUrl field if non-nil, zero value otherwise.

### GetCloudRunUrlOk

`func (o *GcpCollectionAgentPatch) GetCloudRunUrlOk() (*string, bool)`

GetCloudRunUrlOk returns a tuple with the CloudRunUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudRunUrl

`func (o *GcpCollectionAgentPatch) SetCloudRunUrl(v string)`

SetCloudRunUrl sets CloudRunUrl field to given value.

### HasCloudRunUrl

`func (o *GcpCollectionAgentPatch) HasCloudRunUrl() bool`

HasCloudRunUrl returns a boolean if a field has been set.

### SetCloudRunUrlNil

`func (o *GcpCollectionAgentPatch) SetCloudRunUrlNil(b bool)`

 SetCloudRunUrlNil sets the value for CloudRunUrl to be an explicit nil

### UnsetCloudRunUrl
`func (o *GcpCollectionAgentPatch) UnsetCloudRunUrl()`

UnsetCloudRunUrl ensures that no value is present for CloudRunUrl, not even an explicit nil
### GetName

`func (o *GcpCollectionAgentPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpCollectionAgentPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpCollectionAgentPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpCollectionAgentPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GcpCollectionAgentPatch) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GcpCollectionAgentPatch) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAuthenticationType

`func (o *GcpCollectionAgentPatch) GetAuthenticationType() GcpAgentAuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *GcpCollectionAgentPatch) GetAuthenticationTypeOk() (*GcpAgentAuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *GcpCollectionAgentPatch) SetAuthenticationType(v GcpAgentAuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *GcpCollectionAgentPatch) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationTypeNil

`func (o *GcpCollectionAgentPatch) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *GcpCollectionAgentPatch) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


