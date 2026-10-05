# AzureDataFactoryCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**TenantId** | **string** | Microsoft Entra ID tenant the data factory belongs to. | 
**AppClientId** | **string** | Client ID of the Entra ID app registration Monte Carlo signs in with. | 
**AppClientSecret** | **string** | Secret of the app registration. Used for this check and not kept. | 
**SubscriptionId** | **string** | Azure subscription that holds the data factory. | 
**ResourceGroupName** | **string** | Resource group that holds the data factory. | 
**FactoryName** | **string** | Name of the data factory. | 

## Methods

### NewAzureDataFactoryCredentialsValidateIn

`func NewAzureDataFactoryCredentialsValidateIn(deploymentId string, tenantId string, appClientId string, appClientSecret string, subscriptionId string, resourceGroupName string, factoryName string, ) *AzureDataFactoryCredentialsValidateIn`

NewAzureDataFactoryCredentialsValidateIn instantiates a new AzureDataFactoryCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataFactoryCredentialsValidateInWithDefaults

`func NewAzureDataFactoryCredentialsValidateInWithDefaults() *AzureDataFactoryCredentialsValidateIn`

NewAzureDataFactoryCredentialsValidateInWithDefaults instantiates a new AzureDataFactoryCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *AzureDataFactoryCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AzureDataFactoryCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AzureDataFactoryCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetTenantId

`func (o *AzureDataFactoryCredentialsValidateIn) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureDataFactoryCredentialsValidateIn) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureDataFactoryCredentialsValidateIn) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAppClientId

`func (o *AzureDataFactoryCredentialsValidateIn) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *AzureDataFactoryCredentialsValidateIn) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *AzureDataFactoryCredentialsValidateIn) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.


### GetAppClientSecret

`func (o *AzureDataFactoryCredentialsValidateIn) GetAppClientSecret() string`

GetAppClientSecret returns the AppClientSecret field if non-nil, zero value otherwise.

### GetAppClientSecretOk

`func (o *AzureDataFactoryCredentialsValidateIn) GetAppClientSecretOk() (*string, bool)`

GetAppClientSecretOk returns a tuple with the AppClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientSecret

`func (o *AzureDataFactoryCredentialsValidateIn) SetAppClientSecret(v string)`

SetAppClientSecret sets AppClientSecret field to given value.


### GetSubscriptionId

`func (o *AzureDataFactoryCredentialsValidateIn) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureDataFactoryCredentialsValidateIn) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureDataFactoryCredentialsValidateIn) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetResourceGroupName

`func (o *AzureDataFactoryCredentialsValidateIn) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureDataFactoryCredentialsValidateIn) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureDataFactoryCredentialsValidateIn) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetFactoryName

`func (o *AzureDataFactoryCredentialsValidateIn) GetFactoryName() string`

GetFactoryName returns the FactoryName field if non-nil, zero value otherwise.

### GetFactoryNameOk

`func (o *AzureDataFactoryCredentialsValidateIn) GetFactoryNameOk() (*string, bool)`

GetFactoryNameOk returns a tuple with the FactoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryName

`func (o *AzureDataFactoryCredentialsValidateIn) SetFactoryName(v string)`

SetFactoryName sets FactoryName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


