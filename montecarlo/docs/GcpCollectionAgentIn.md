# GcpCollectionAgentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthHeaders** | Pointer to [**NullableAuthHeadersCredentialsIn**](AuthHeadersCredentialsIn.md) | Credentials for &#x60;CUSTOM_AUTH_HEADERS&#x60;. Send this or &#x60;service_account_key&#x60;, never both. It replaces the stored credentials rather than merging into them. | [optional] 
**AuthenticationType** | [**GcpAgentAuthenticationType**](GcpAgentAuthenticationType.md) | How Monte Carlo authenticates when it calls the agent. Send it together with the matching credentials. | 
**CloudRunUrl** | **string** | URL of the Cloud Run service Monte Carlo should call. | 
**DeploymentId** | **string** | Deployment to register the collection agent on. It must already hold an unregistered GCP collection agent. | 
**Name** | Pointer to **NullableString** | Display name for the collection agent. Replaces the name it currently has. | [optional] 
**ServiceAccountKey** | Pointer to **NullableString** | Credentials for &#x60;GCP_JSON_SERVICE_ACCOUNT_KEY&#x60;, as the contents of the JSON key file Google issued for the service account. Send this or &#x60;auth_headers&#x60;, never both. It replaces the stored credentials rather than merging into them. | [optional] 

## Methods

### NewGcpCollectionAgentIn

`func NewGcpCollectionAgentIn(authenticationType GcpAgentAuthenticationType, cloudRunUrl string, deploymentId string, ) *GcpCollectionAgentIn`

NewGcpCollectionAgentIn instantiates a new GcpCollectionAgentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpCollectionAgentInWithDefaults

`func NewGcpCollectionAgentInWithDefaults() *GcpCollectionAgentIn`

NewGcpCollectionAgentInWithDefaults instantiates a new GcpCollectionAgentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthHeaders

`func (o *GcpCollectionAgentIn) GetAuthHeaders() AuthHeadersCredentialsIn`

GetAuthHeaders returns the AuthHeaders field if non-nil, zero value otherwise.

### GetAuthHeadersOk

`func (o *GcpCollectionAgentIn) GetAuthHeadersOk() (*AuthHeadersCredentialsIn, bool)`

GetAuthHeadersOk returns a tuple with the AuthHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthHeaders

`func (o *GcpCollectionAgentIn) SetAuthHeaders(v AuthHeadersCredentialsIn)`

SetAuthHeaders sets AuthHeaders field to given value.

### HasAuthHeaders

`func (o *GcpCollectionAgentIn) HasAuthHeaders() bool`

HasAuthHeaders returns a boolean if a field has been set.

### SetAuthHeadersNil

`func (o *GcpCollectionAgentIn) SetAuthHeadersNil(b bool)`

 SetAuthHeadersNil sets the value for AuthHeaders to be an explicit nil

### UnsetAuthHeaders
`func (o *GcpCollectionAgentIn) UnsetAuthHeaders()`

UnsetAuthHeaders ensures that no value is present for AuthHeaders, not even an explicit nil
### GetAuthenticationType

`func (o *GcpCollectionAgentIn) GetAuthenticationType() GcpAgentAuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *GcpCollectionAgentIn) GetAuthenticationTypeOk() (*GcpAgentAuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *GcpCollectionAgentIn) SetAuthenticationType(v GcpAgentAuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetCloudRunUrl

`func (o *GcpCollectionAgentIn) GetCloudRunUrl() string`

GetCloudRunUrl returns the CloudRunUrl field if non-nil, zero value otherwise.

### GetCloudRunUrlOk

`func (o *GcpCollectionAgentIn) GetCloudRunUrlOk() (*string, bool)`

GetCloudRunUrlOk returns a tuple with the CloudRunUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudRunUrl

`func (o *GcpCollectionAgentIn) SetCloudRunUrl(v string)`

SetCloudRunUrl sets CloudRunUrl field to given value.


### GetDeploymentId

`func (o *GcpCollectionAgentIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GcpCollectionAgentIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GcpCollectionAgentIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetName

`func (o *GcpCollectionAgentIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpCollectionAgentIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpCollectionAgentIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpCollectionAgentIn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GcpCollectionAgentIn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GcpCollectionAgentIn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetServiceAccountKey

`func (o *GcpCollectionAgentIn) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *GcpCollectionAgentIn) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *GcpCollectionAgentIn) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.

### HasServiceAccountKey

`func (o *GcpCollectionAgentIn) HasServiceAccountKey() bool`

HasServiceAccountKey returns a boolean if a field has been set.

### SetServiceAccountKeyNil

`func (o *GcpCollectionAgentIn) SetServiceAccountKeyNil(b bool)`

 SetServiceAccountKeyNil sets the value for ServiceAccountKey to be an explicit nil

### UnsetServiceAccountKey
`func (o *GcpCollectionAgentIn) UnsetServiceAccountKey()`

UnsetServiceAccountKey ensures that no value is present for ServiceAccountKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


