# AzureDataFactoryCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**TenantId** | Pointer to **NullableString** | Microsoft Entra ID tenant the data factory belongs to. | [optional] 
**AppClientId** | Pointer to **NullableString** | Client ID of the Entra ID app registration Monte Carlo signs in with. | [optional] 
**AppClientSecret** | Pointer to **NullableString** | Secret of the app registration. Stored by Monte Carlo and never returned. | [optional] 
**SubscriptionId** | Pointer to **NullableString** | Azure subscription that holds the data factory. | [optional] 
**ResourceGroupName** | Pointer to **NullableString** | Resource group that holds the data factory. | [optional] 
**FactoryName** | Pointer to **NullableString** | Name of the data factory. | [optional] 

## Methods

### NewAzureDataFactoryCredentialsChangeValidateIn

`func NewAzureDataFactoryCredentialsChangeValidateIn(deploymentId string, ) *AzureDataFactoryCredentialsChangeValidateIn`

NewAzureDataFactoryCredentialsChangeValidateIn instantiates a new AzureDataFactoryCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataFactoryCredentialsChangeValidateInWithDefaults

`func NewAzureDataFactoryCredentialsChangeValidateInWithDefaults() *AzureDataFactoryCredentialsChangeValidateIn`

NewAzureDataFactoryCredentialsChangeValidateInWithDefaults instantiates a new AzureDataFactoryCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetTenantId

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AzureDataFactoryCredentialsChangeValidateIn) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AzureDataFactoryCredentialsChangeValidateIn) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetAppClientId

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.

### HasAppClientId

`func (o *AzureDataFactoryCredentialsChangeValidateIn) HasAppClientId() bool`

HasAppClientId returns a boolean if a field has been set.

### SetAppClientIdNil

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetAppClientIdNil(b bool)`

 SetAppClientIdNil sets the value for AppClientId to be an explicit nil

### UnsetAppClientId
`func (o *AzureDataFactoryCredentialsChangeValidateIn) UnsetAppClientId()`

UnsetAppClientId ensures that no value is present for AppClientId, not even an explicit nil
### GetAppClientSecret

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetAppClientSecret() string`

GetAppClientSecret returns the AppClientSecret field if non-nil, zero value otherwise.

### GetAppClientSecretOk

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetAppClientSecretOk() (*string, bool)`

GetAppClientSecretOk returns a tuple with the AppClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientSecret

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetAppClientSecret(v string)`

SetAppClientSecret sets AppClientSecret field to given value.

### HasAppClientSecret

`func (o *AzureDataFactoryCredentialsChangeValidateIn) HasAppClientSecret() bool`

HasAppClientSecret returns a boolean if a field has been set.

### SetAppClientSecretNil

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetAppClientSecretNil(b bool)`

 SetAppClientSecretNil sets the value for AppClientSecret to be an explicit nil

### UnsetAppClientSecret
`func (o *AzureDataFactoryCredentialsChangeValidateIn) UnsetAppClientSecret()`

UnsetAppClientSecret ensures that no value is present for AppClientSecret, not even an explicit nil
### GetSubscriptionId

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureDataFactoryCredentialsChangeValidateIn) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *AzureDataFactoryCredentialsChangeValidateIn) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetResourceGroupName

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureDataFactoryCredentialsChangeValidateIn) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *AzureDataFactoryCredentialsChangeValidateIn) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetFactoryName

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetFactoryName() string`

GetFactoryName returns the FactoryName field if non-nil, zero value otherwise.

### GetFactoryNameOk

`func (o *AzureDataFactoryCredentialsChangeValidateIn) GetFactoryNameOk() (*string, bool)`

GetFactoryNameOk returns a tuple with the FactoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryName

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetFactoryName(v string)`

SetFactoryName sets FactoryName field to given value.

### HasFactoryName

`func (o *AzureDataFactoryCredentialsChangeValidateIn) HasFactoryName() bool`

HasFactoryName returns a boolean if a field has been set.

### SetFactoryNameNil

`func (o *AzureDataFactoryCredentialsChangeValidateIn) SetFactoryNameNil(b bool)`

 SetFactoryNameNil sets the value for FactoryName to be an explicit nil

### UnsetFactoryName
`func (o *AzureDataFactoryCredentialsChangeValidateIn) UnsetFactoryName()`

UnsetFactoryName ensures that no value is present for FactoryName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


