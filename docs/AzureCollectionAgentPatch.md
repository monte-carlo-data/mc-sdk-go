# AzureCollectionAgentPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to [**NullableAzureAgentAuthenticationType**](AzureAgentAuthenticationType.md) | How Monte Carlo authenticates when it calls the agent. Send it together with the matching credentials object. | [optional] 
**FunctionAppKey** | Pointer to [**NullableFunctionAppKeyCredentialsIn**](FunctionAppKeyCredentialsIn.md) | Credentials for &#x60;AZURE_FUNCTION_APP_KEY&#x60;. Send this or &#x60;service_principal&#x60;, never both. | [optional] 
**FunctionAppUrl** | Pointer to **NullableString** | URL of the function app Monte Carlo should call. | [optional] 
**Name** | Pointer to **NullableString** | Display name for the collection agent. Replaces the name it currently has. | [optional] 
**ServicePrincipal** | Pointer to [**NullableServicePrincipalCredentialsIn**](ServicePrincipalCredentialsIn.md) | Credentials for &#x60;AZURE_FUNCTION_SERVICE_PRINCIPAL&#x60;. Send this or &#x60;function_app_key&#x60;, never both. | [optional] 

## Methods

### NewAzureCollectionAgentPatch

`func NewAzureCollectionAgentPatch() *AzureCollectionAgentPatch`

NewAzureCollectionAgentPatch instantiates a new AzureCollectionAgentPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCollectionAgentPatchWithDefaults

`func NewAzureCollectionAgentPatchWithDefaults() *AzureCollectionAgentPatch`

NewAzureCollectionAgentPatchWithDefaults instantiates a new AzureCollectionAgentPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *AzureCollectionAgentPatch) GetAuthenticationType() AzureAgentAuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AzureCollectionAgentPatch) GetAuthenticationTypeOk() (*AzureAgentAuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AzureCollectionAgentPatch) SetAuthenticationType(v AzureAgentAuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *AzureCollectionAgentPatch) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationTypeNil

`func (o *AzureCollectionAgentPatch) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *AzureCollectionAgentPatch) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
### GetFunctionAppKey

`func (o *AzureCollectionAgentPatch) GetFunctionAppKey() FunctionAppKeyCredentialsIn`

GetFunctionAppKey returns the FunctionAppKey field if non-nil, zero value otherwise.

### GetFunctionAppKeyOk

`func (o *AzureCollectionAgentPatch) GetFunctionAppKeyOk() (*FunctionAppKeyCredentialsIn, bool)`

GetFunctionAppKeyOk returns a tuple with the FunctionAppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionAppKey

`func (o *AzureCollectionAgentPatch) SetFunctionAppKey(v FunctionAppKeyCredentialsIn)`

SetFunctionAppKey sets FunctionAppKey field to given value.

### HasFunctionAppKey

`func (o *AzureCollectionAgentPatch) HasFunctionAppKey() bool`

HasFunctionAppKey returns a boolean if a field has been set.

### SetFunctionAppKeyNil

`func (o *AzureCollectionAgentPatch) SetFunctionAppKeyNil(b bool)`

 SetFunctionAppKeyNil sets the value for FunctionAppKey to be an explicit nil

### UnsetFunctionAppKey
`func (o *AzureCollectionAgentPatch) UnsetFunctionAppKey()`

UnsetFunctionAppKey ensures that no value is present for FunctionAppKey, not even an explicit nil
### GetFunctionAppUrl

`func (o *AzureCollectionAgentPatch) GetFunctionAppUrl() string`

GetFunctionAppUrl returns the FunctionAppUrl field if non-nil, zero value otherwise.

### GetFunctionAppUrlOk

`func (o *AzureCollectionAgentPatch) GetFunctionAppUrlOk() (*string, bool)`

GetFunctionAppUrlOk returns a tuple with the FunctionAppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionAppUrl

`func (o *AzureCollectionAgentPatch) SetFunctionAppUrl(v string)`

SetFunctionAppUrl sets FunctionAppUrl field to given value.

### HasFunctionAppUrl

`func (o *AzureCollectionAgentPatch) HasFunctionAppUrl() bool`

HasFunctionAppUrl returns a boolean if a field has been set.

### SetFunctionAppUrlNil

`func (o *AzureCollectionAgentPatch) SetFunctionAppUrlNil(b bool)`

 SetFunctionAppUrlNil sets the value for FunctionAppUrl to be an explicit nil

### UnsetFunctionAppUrl
`func (o *AzureCollectionAgentPatch) UnsetFunctionAppUrl()`

UnsetFunctionAppUrl ensures that no value is present for FunctionAppUrl, not even an explicit nil
### GetName

`func (o *AzureCollectionAgentPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureCollectionAgentPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureCollectionAgentPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureCollectionAgentPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureCollectionAgentPatch) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureCollectionAgentPatch) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetServicePrincipal

`func (o *AzureCollectionAgentPatch) GetServicePrincipal() ServicePrincipalCredentialsIn`

GetServicePrincipal returns the ServicePrincipal field if non-nil, zero value otherwise.

### GetServicePrincipalOk

`func (o *AzureCollectionAgentPatch) GetServicePrincipalOk() (*ServicePrincipalCredentialsIn, bool)`

GetServicePrincipalOk returns a tuple with the ServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipal

`func (o *AzureCollectionAgentPatch) SetServicePrincipal(v ServicePrincipalCredentialsIn)`

SetServicePrincipal sets ServicePrincipal field to given value.

### HasServicePrincipal

`func (o *AzureCollectionAgentPatch) HasServicePrincipal() bool`

HasServicePrincipal returns a boolean if a field has been set.

### SetServicePrincipalNil

`func (o *AzureCollectionAgentPatch) SetServicePrincipalNil(b bool)`

 SetServicePrincipalNil sets the value for ServicePrincipal to be an explicit nil

### UnsetServicePrincipal
`func (o *AzureCollectionAgentPatch) UnsetServicePrincipal()`

UnsetServicePrincipal ensures that no value is present for ServicePrincipal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


