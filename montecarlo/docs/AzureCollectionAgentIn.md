# AzureCollectionAgentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionAppKey** | Pointer to [**NullableFunctionAppKeyCredentialsIn**](FunctionAppKeyCredentialsIn.md) | Credentials for &#x60;AZURE_FUNCTION_APP_KEY&#x60;. Send this or &#x60;service_principal&#x60;, never both. | [optional] 
**ServicePrincipal** | Pointer to [**NullableServicePrincipalCredentialsIn**](ServicePrincipalCredentialsIn.md) | Credentials for &#x60;AZURE_FUNCTION_SERVICE_PRINCIPAL&#x60;. Send this or &#x60;function_app_key&#x60;, never both. | [optional] 
**AuthenticationType** | [**AzureAgentAuthenticationType**](AzureAgentAuthenticationType.md) | How Monte Carlo authenticates when it calls the agent. Send it together with the matching credentials object. | 
**DeploymentId** | **string** | Deployment to register the collection agent on. It must already hold an unregistered Azure collection agent. | 
**FunctionAppUrl** | **string** | URL of the function app Monte Carlo should call. | 
**Name** | Pointer to **NullableString** | Display name for the collection agent. Replaces the name it currently has. | [optional] 

## Methods

### NewAzureCollectionAgentIn

`func NewAzureCollectionAgentIn(authenticationType AzureAgentAuthenticationType, deploymentId string, functionAppUrl string, ) *AzureCollectionAgentIn`

NewAzureCollectionAgentIn instantiates a new AzureCollectionAgentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCollectionAgentInWithDefaults

`func NewAzureCollectionAgentInWithDefaults() *AzureCollectionAgentIn`

NewAzureCollectionAgentInWithDefaults instantiates a new AzureCollectionAgentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionAppKey

`func (o *AzureCollectionAgentIn) GetFunctionAppKey() FunctionAppKeyCredentialsIn`

GetFunctionAppKey returns the FunctionAppKey field if non-nil, zero value otherwise.

### GetFunctionAppKeyOk

`func (o *AzureCollectionAgentIn) GetFunctionAppKeyOk() (*FunctionAppKeyCredentialsIn, bool)`

GetFunctionAppKeyOk returns a tuple with the FunctionAppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionAppKey

`func (o *AzureCollectionAgentIn) SetFunctionAppKey(v FunctionAppKeyCredentialsIn)`

SetFunctionAppKey sets FunctionAppKey field to given value.

### HasFunctionAppKey

`func (o *AzureCollectionAgentIn) HasFunctionAppKey() bool`

HasFunctionAppKey returns a boolean if a field has been set.

### SetFunctionAppKeyNil

`func (o *AzureCollectionAgentIn) SetFunctionAppKeyNil(b bool)`

 SetFunctionAppKeyNil sets the value for FunctionAppKey to be an explicit nil

### UnsetFunctionAppKey
`func (o *AzureCollectionAgentIn) UnsetFunctionAppKey()`

UnsetFunctionAppKey ensures that no value is present for FunctionAppKey, not even an explicit nil
### GetServicePrincipal

`func (o *AzureCollectionAgentIn) GetServicePrincipal() ServicePrincipalCredentialsIn`

GetServicePrincipal returns the ServicePrincipal field if non-nil, zero value otherwise.

### GetServicePrincipalOk

`func (o *AzureCollectionAgentIn) GetServicePrincipalOk() (*ServicePrincipalCredentialsIn, bool)`

GetServicePrincipalOk returns a tuple with the ServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipal

`func (o *AzureCollectionAgentIn) SetServicePrincipal(v ServicePrincipalCredentialsIn)`

SetServicePrincipal sets ServicePrincipal field to given value.

### HasServicePrincipal

`func (o *AzureCollectionAgentIn) HasServicePrincipal() bool`

HasServicePrincipal returns a boolean if a field has been set.

### SetServicePrincipalNil

`func (o *AzureCollectionAgentIn) SetServicePrincipalNil(b bool)`

 SetServicePrincipalNil sets the value for ServicePrincipal to be an explicit nil

### UnsetServicePrincipal
`func (o *AzureCollectionAgentIn) UnsetServicePrincipal()`

UnsetServicePrincipal ensures that no value is present for ServicePrincipal, not even an explicit nil
### GetAuthenticationType

`func (o *AzureCollectionAgentIn) GetAuthenticationType() AzureAgentAuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AzureCollectionAgentIn) GetAuthenticationTypeOk() (*AzureAgentAuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AzureCollectionAgentIn) SetAuthenticationType(v AzureAgentAuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetDeploymentId

`func (o *AzureCollectionAgentIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AzureCollectionAgentIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AzureCollectionAgentIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetFunctionAppUrl

`func (o *AzureCollectionAgentIn) GetFunctionAppUrl() string`

GetFunctionAppUrl returns the FunctionAppUrl field if non-nil, zero value otherwise.

### GetFunctionAppUrlOk

`func (o *AzureCollectionAgentIn) GetFunctionAppUrlOk() (*string, bool)`

GetFunctionAppUrlOk returns a tuple with the FunctionAppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionAppUrl

`func (o *AzureCollectionAgentIn) SetFunctionAppUrl(v string)`

SetFunctionAppUrl sets FunctionAppUrl field to given value.


### GetName

`func (o *AzureCollectionAgentIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureCollectionAgentIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureCollectionAgentIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureCollectionAgentIn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureCollectionAgentIn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureCollectionAgentIn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


