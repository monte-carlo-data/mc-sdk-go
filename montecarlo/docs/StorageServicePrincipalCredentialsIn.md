# StorageServicePrincipalCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **NullableString** | Name of the storage account, needed only when &#x60;account_url&#x60; does not start with it. Monte Carlo takes the first label of the host otherwise, which is right for a standard or private-link URL but not for a custom ingress host. | [optional] 
**AccountUrl** | **string** | URL of the storage account Monte Carlo sends requests to. A private endpoint or a custom host works, so this is not required to be under &#x60;blob.core.windows.net&#x60;. | 
**ClientId** | **string** | Application (client) id of the service principal. | 
**ClientSecret** | **string** | Client secret of the service principal. | 
**TenantId** | **string** | Directory (tenant) id the service principal lives in. | 

## Methods

### NewStorageServicePrincipalCredentialsIn

`func NewStorageServicePrincipalCredentialsIn(accountUrl string, clientId string, clientSecret string, tenantId string, ) *StorageServicePrincipalCredentialsIn`

NewStorageServicePrincipalCredentialsIn instantiates a new StorageServicePrincipalCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageServicePrincipalCredentialsInWithDefaults

`func NewStorageServicePrincipalCredentialsInWithDefaults() *StorageServicePrincipalCredentialsIn`

NewStorageServicePrincipalCredentialsInWithDefaults instantiates a new StorageServicePrincipalCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *StorageServicePrincipalCredentialsIn) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *StorageServicePrincipalCredentialsIn) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *StorageServicePrincipalCredentialsIn) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *StorageServicePrincipalCredentialsIn) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *StorageServicePrincipalCredentialsIn) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *StorageServicePrincipalCredentialsIn) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetAccountUrl

`func (o *StorageServicePrincipalCredentialsIn) GetAccountUrl() string`

GetAccountUrl returns the AccountUrl field if non-nil, zero value otherwise.

### GetAccountUrlOk

`func (o *StorageServicePrincipalCredentialsIn) GetAccountUrlOk() (*string, bool)`

GetAccountUrlOk returns a tuple with the AccountUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUrl

`func (o *StorageServicePrincipalCredentialsIn) SetAccountUrl(v string)`

SetAccountUrl sets AccountUrl field to given value.


### GetClientId

`func (o *StorageServicePrincipalCredentialsIn) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *StorageServicePrincipalCredentialsIn) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *StorageServicePrincipalCredentialsIn) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *StorageServicePrincipalCredentialsIn) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *StorageServicePrincipalCredentialsIn) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *StorageServicePrincipalCredentialsIn) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetTenantId

`func (o *StorageServicePrincipalCredentialsIn) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *StorageServicePrincipalCredentialsIn) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *StorageServicePrincipalCredentialsIn) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


