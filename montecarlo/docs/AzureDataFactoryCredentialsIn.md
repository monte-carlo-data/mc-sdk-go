# AzureDataFactoryCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | Microsoft Entra ID tenant the data factory belongs to. | 
**AppClientId** | **string** | Client ID of the Entra ID app registration Monte Carlo signs in with. | 
**AppClientSecret** | **string** | Secret of the app registration. Stored by Monte Carlo and never returned. | 
**SubscriptionId** | **string** | Azure subscription that holds the data factory. | 
**ResourceGroupName** | **string** | Resource group that holds the data factory. | 
**FactoryName** | **string** | Name of the data factory. | 

## Methods

### NewAzureDataFactoryCredentialsIn

`func NewAzureDataFactoryCredentialsIn(tenantId string, appClientId string, appClientSecret string, subscriptionId string, resourceGroupName string, factoryName string, ) *AzureDataFactoryCredentialsIn`

NewAzureDataFactoryCredentialsIn instantiates a new AzureDataFactoryCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataFactoryCredentialsInWithDefaults

`func NewAzureDataFactoryCredentialsInWithDefaults() *AzureDataFactoryCredentialsIn`

NewAzureDataFactoryCredentialsInWithDefaults instantiates a new AzureDataFactoryCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AzureDataFactoryCredentialsIn) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureDataFactoryCredentialsIn) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureDataFactoryCredentialsIn) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAppClientId

`func (o *AzureDataFactoryCredentialsIn) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *AzureDataFactoryCredentialsIn) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *AzureDataFactoryCredentialsIn) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.


### GetAppClientSecret

`func (o *AzureDataFactoryCredentialsIn) GetAppClientSecret() string`

GetAppClientSecret returns the AppClientSecret field if non-nil, zero value otherwise.

### GetAppClientSecretOk

`func (o *AzureDataFactoryCredentialsIn) GetAppClientSecretOk() (*string, bool)`

GetAppClientSecretOk returns a tuple with the AppClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientSecret

`func (o *AzureDataFactoryCredentialsIn) SetAppClientSecret(v string)`

SetAppClientSecret sets AppClientSecret field to given value.


### GetSubscriptionId

`func (o *AzureDataFactoryCredentialsIn) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureDataFactoryCredentialsIn) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureDataFactoryCredentialsIn) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetResourceGroupName

`func (o *AzureDataFactoryCredentialsIn) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureDataFactoryCredentialsIn) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureDataFactoryCredentialsIn) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetFactoryName

`func (o *AzureDataFactoryCredentialsIn) GetFactoryName() string`

GetFactoryName returns the FactoryName field if non-nil, zero value otherwise.

### GetFactoryNameOk

`func (o *AzureDataFactoryCredentialsIn) GetFactoryNameOk() (*string, bool)`

GetFactoryNameOk returns a tuple with the FactoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryName

`func (o *AzureDataFactoryCredentialsIn) SetFactoryName(v string)`

SetFactoryName sets FactoryName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


