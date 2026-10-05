# AzureDataFactoryCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**TenantId** | **string** | Microsoft Entra ID tenant the data factory belongs to. | 
**AppClientId** | **string** | Client ID of the Entra ID app registration Monte Carlo signs in with. | 
**SubscriptionId** | **string** | Azure subscription that holds the data factory. | 
**ResourceGroupName** | **string** | Resource group that holds the data factory. | 
**FactoryName** | **string** | Name of the data factory. | 

## Methods

### NewAzureDataFactoryCredentialsOut

`func NewAzureDataFactoryCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, tenantId string, appClientId string, subscriptionId string, resourceGroupName string, factoryName string, ) *AzureDataFactoryCredentialsOut`

NewAzureDataFactoryCredentialsOut instantiates a new AzureDataFactoryCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataFactoryCredentialsOutWithDefaults

`func NewAzureDataFactoryCredentialsOutWithDefaults() *AzureDataFactoryCredentialsOut`

NewAzureDataFactoryCredentialsOutWithDefaults instantiates a new AzureDataFactoryCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureDataFactoryCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureDataFactoryCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureDataFactoryCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *AzureDataFactoryCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *AzureDataFactoryCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *AzureDataFactoryCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *AzureDataFactoryCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *AzureDataFactoryCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *AzureDataFactoryCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *AzureDataFactoryCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AzureDataFactoryCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AzureDataFactoryCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetTenantId

`func (o *AzureDataFactoryCredentialsOut) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureDataFactoryCredentialsOut) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureDataFactoryCredentialsOut) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetAppClientId

`func (o *AzureDataFactoryCredentialsOut) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *AzureDataFactoryCredentialsOut) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *AzureDataFactoryCredentialsOut) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.


### GetSubscriptionId

`func (o *AzureDataFactoryCredentialsOut) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureDataFactoryCredentialsOut) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureDataFactoryCredentialsOut) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetResourceGroupName

`func (o *AzureDataFactoryCredentialsOut) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureDataFactoryCredentialsOut) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureDataFactoryCredentialsOut) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### GetFactoryName

`func (o *AzureDataFactoryCredentialsOut) GetFactoryName() string`

GetFactoryName returns the FactoryName field if non-nil, zero value otherwise.

### GetFactoryNameOk

`func (o *AzureDataFactoryCredentialsOut) GetFactoryNameOk() (*string, bool)`

GetFactoryNameOk returns a tuple with the FactoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryName

`func (o *AzureDataFactoryCredentialsOut) SetFactoryName(v string)`

SetFactoryName sets FactoryName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


