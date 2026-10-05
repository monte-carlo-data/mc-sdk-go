# AzureDataFactoryCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** | Microsoft Entra ID tenant the data factory belongs to. | [optional] 
**AppClientId** | Pointer to **NullableString** | Client ID of the Entra ID app registration Monte Carlo signs in with. | [optional] 
**AppClientSecret** | Pointer to **NullableString** | Secret of the app registration. Stored by Monte Carlo and never returned. | [optional] 
**SubscriptionId** | Pointer to **NullableString** | Azure subscription that holds the data factory. | [optional] 
**ResourceGroupName** | Pointer to **NullableString** | Resource group that holds the data factory. | [optional] 
**FactoryName** | Pointer to **NullableString** | Name of the data factory. | [optional] 

## Methods

### NewAzureDataFactoryCredentialsPatch

`func NewAzureDataFactoryCredentialsPatch() *AzureDataFactoryCredentialsPatch`

NewAzureDataFactoryCredentialsPatch instantiates a new AzureDataFactoryCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataFactoryCredentialsPatchWithDefaults

`func NewAzureDataFactoryCredentialsPatchWithDefaults() *AzureDataFactoryCredentialsPatch`

NewAzureDataFactoryCredentialsPatchWithDefaults instantiates a new AzureDataFactoryCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AzureDataFactoryCredentialsPatch) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureDataFactoryCredentialsPatch) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureDataFactoryCredentialsPatch) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AzureDataFactoryCredentialsPatch) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AzureDataFactoryCredentialsPatch) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AzureDataFactoryCredentialsPatch) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetAppClientId

`func (o *AzureDataFactoryCredentialsPatch) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *AzureDataFactoryCredentialsPatch) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *AzureDataFactoryCredentialsPatch) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.

### HasAppClientId

`func (o *AzureDataFactoryCredentialsPatch) HasAppClientId() bool`

HasAppClientId returns a boolean if a field has been set.

### SetAppClientIdNil

`func (o *AzureDataFactoryCredentialsPatch) SetAppClientIdNil(b bool)`

 SetAppClientIdNil sets the value for AppClientId to be an explicit nil

### UnsetAppClientId
`func (o *AzureDataFactoryCredentialsPatch) UnsetAppClientId()`

UnsetAppClientId ensures that no value is present for AppClientId, not even an explicit nil
### GetAppClientSecret

`func (o *AzureDataFactoryCredentialsPatch) GetAppClientSecret() string`

GetAppClientSecret returns the AppClientSecret field if non-nil, zero value otherwise.

### GetAppClientSecretOk

`func (o *AzureDataFactoryCredentialsPatch) GetAppClientSecretOk() (*string, bool)`

GetAppClientSecretOk returns a tuple with the AppClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientSecret

`func (o *AzureDataFactoryCredentialsPatch) SetAppClientSecret(v string)`

SetAppClientSecret sets AppClientSecret field to given value.

### HasAppClientSecret

`func (o *AzureDataFactoryCredentialsPatch) HasAppClientSecret() bool`

HasAppClientSecret returns a boolean if a field has been set.

### SetAppClientSecretNil

`func (o *AzureDataFactoryCredentialsPatch) SetAppClientSecretNil(b bool)`

 SetAppClientSecretNil sets the value for AppClientSecret to be an explicit nil

### UnsetAppClientSecret
`func (o *AzureDataFactoryCredentialsPatch) UnsetAppClientSecret()`

UnsetAppClientSecret ensures that no value is present for AppClientSecret, not even an explicit nil
### GetSubscriptionId

`func (o *AzureDataFactoryCredentialsPatch) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureDataFactoryCredentialsPatch) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureDataFactoryCredentialsPatch) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureDataFactoryCredentialsPatch) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *AzureDataFactoryCredentialsPatch) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *AzureDataFactoryCredentialsPatch) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetResourceGroupName

`func (o *AzureDataFactoryCredentialsPatch) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureDataFactoryCredentialsPatch) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureDataFactoryCredentialsPatch) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.

### HasResourceGroupName

`func (o *AzureDataFactoryCredentialsPatch) HasResourceGroupName() bool`

HasResourceGroupName returns a boolean if a field has been set.

### SetResourceGroupNameNil

`func (o *AzureDataFactoryCredentialsPatch) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *AzureDataFactoryCredentialsPatch) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetFactoryName

`func (o *AzureDataFactoryCredentialsPatch) GetFactoryName() string`

GetFactoryName returns the FactoryName field if non-nil, zero value otherwise.

### GetFactoryNameOk

`func (o *AzureDataFactoryCredentialsPatch) GetFactoryNameOk() (*string, bool)`

GetFactoryNameOk returns a tuple with the FactoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryName

`func (o *AzureDataFactoryCredentialsPatch) SetFactoryName(v string)`

SetFactoryName sets FactoryName field to given value.

### HasFactoryName

`func (o *AzureDataFactoryCredentialsPatch) HasFactoryName() bool`

HasFactoryName returns a boolean if a field has been set.

### SetFactoryNameNil

`func (o *AzureDataFactoryCredentialsPatch) SetFactoryNameNil(b bool)`

 SetFactoryNameNil sets the value for FactoryName to be an explicit nil

### UnsetFactoryName
`func (o *AzureDataFactoryCredentialsPatch) UnsetFactoryName()`

UnsetFactoryName ensures that no value is present for FactoryName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


