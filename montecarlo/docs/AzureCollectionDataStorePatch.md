# AzureCollectionDataStorePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | Pointer to [**NullableAzureDataStoreAuthenticationType**](AzureDataStoreAuthenticationType.md) | How Monte Carlo authenticates to the storage account. Send it together with the matching credentials object. | [optional] 
**ContainerName** | Pointer to **NullableString** | Name of the blob container Monte Carlo should use. | [optional] 
**Name** | Pointer to **NullableString** | Display name for the data store. Replaces the name its deployment gave it. | [optional] 
**ServicePrincipal** | Pointer to [**NullableStorageServicePrincipalCredentialsIn**](StorageServicePrincipalCredentialsIn.md) | Credentials for &#x60;AZURE_STORAGE_SERVICE_PRINCIPAL&#x60;. Send this or &#x60;storage_account_keys&#x60;, never both. | [optional] 
**StorageAccountKeys** | Pointer to [**NullableStorageAccountKeysCredentialsIn**](StorageAccountKeysCredentialsIn.md) | Credentials for &#x60;AZURE_STORAGE_ACCOUNT_KEYS&#x60;. Send this or &#x60;service_principal&#x60;, never both. | [optional] 

## Methods

### NewAzureCollectionDataStorePatch

`func NewAzureCollectionDataStorePatch() *AzureCollectionDataStorePatch`

NewAzureCollectionDataStorePatch instantiates a new AzureCollectionDataStorePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCollectionDataStorePatchWithDefaults

`func NewAzureCollectionDataStorePatchWithDefaults() *AzureCollectionDataStorePatch`

NewAzureCollectionDataStorePatchWithDefaults instantiates a new AzureCollectionDataStorePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *AzureCollectionDataStorePatch) GetAuthenticationType() AzureDataStoreAuthenticationType`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AzureCollectionDataStorePatch) GetAuthenticationTypeOk() (*AzureDataStoreAuthenticationType, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AzureCollectionDataStorePatch) SetAuthenticationType(v AzureDataStoreAuthenticationType)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *AzureCollectionDataStorePatch) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationTypeNil

`func (o *AzureCollectionDataStorePatch) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *AzureCollectionDataStorePatch) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
### GetContainerName

`func (o *AzureCollectionDataStorePatch) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *AzureCollectionDataStorePatch) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *AzureCollectionDataStorePatch) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *AzureCollectionDataStorePatch) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### SetContainerNameNil

`func (o *AzureCollectionDataStorePatch) SetContainerNameNil(b bool)`

 SetContainerNameNil sets the value for ContainerName to be an explicit nil

### UnsetContainerName
`func (o *AzureCollectionDataStorePatch) UnsetContainerName()`

UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
### GetName

`func (o *AzureCollectionDataStorePatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureCollectionDataStorePatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureCollectionDataStorePatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureCollectionDataStorePatch) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureCollectionDataStorePatch) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureCollectionDataStorePatch) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetServicePrincipal

`func (o *AzureCollectionDataStorePatch) GetServicePrincipal() StorageServicePrincipalCredentialsIn`

GetServicePrincipal returns the ServicePrincipal field if non-nil, zero value otherwise.

### GetServicePrincipalOk

`func (o *AzureCollectionDataStorePatch) GetServicePrincipalOk() (*StorageServicePrincipalCredentialsIn, bool)`

GetServicePrincipalOk returns a tuple with the ServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipal

`func (o *AzureCollectionDataStorePatch) SetServicePrincipal(v StorageServicePrincipalCredentialsIn)`

SetServicePrincipal sets ServicePrincipal field to given value.

### HasServicePrincipal

`func (o *AzureCollectionDataStorePatch) HasServicePrincipal() bool`

HasServicePrincipal returns a boolean if a field has been set.

### SetServicePrincipalNil

`func (o *AzureCollectionDataStorePatch) SetServicePrincipalNil(b bool)`

 SetServicePrincipalNil sets the value for ServicePrincipal to be an explicit nil

### UnsetServicePrincipal
`func (o *AzureCollectionDataStorePatch) UnsetServicePrincipal()`

UnsetServicePrincipal ensures that no value is present for ServicePrincipal, not even an explicit nil
### GetStorageAccountKeys

`func (o *AzureCollectionDataStorePatch) GetStorageAccountKeys() StorageAccountKeysCredentialsIn`

GetStorageAccountKeys returns the StorageAccountKeys field if non-nil, zero value otherwise.

### GetStorageAccountKeysOk

`func (o *AzureCollectionDataStorePatch) GetStorageAccountKeysOk() (*StorageAccountKeysCredentialsIn, bool)`

GetStorageAccountKeysOk returns a tuple with the StorageAccountKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountKeys

`func (o *AzureCollectionDataStorePatch) SetStorageAccountKeys(v StorageAccountKeysCredentialsIn)`

SetStorageAccountKeys sets StorageAccountKeys field to given value.

### HasStorageAccountKeys

`func (o *AzureCollectionDataStorePatch) HasStorageAccountKeys() bool`

HasStorageAccountKeys returns a boolean if a field has been set.

### SetStorageAccountKeysNil

`func (o *AzureCollectionDataStorePatch) SetStorageAccountKeysNil(b bool)`

 SetStorageAccountKeysNil sets the value for StorageAccountKeys to be an explicit nil

### UnsetStorageAccountKeys
`func (o *AzureCollectionDataStorePatch) UnsetStorageAccountKeys()`

UnsetStorageAccountKeys ensures that no value is present for StorageAccountKeys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


