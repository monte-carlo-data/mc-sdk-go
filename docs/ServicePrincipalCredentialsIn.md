# ServicePrincipalCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | **string** | Audience the issued token is for, usually the function app&#39;s application id. | 
**ClientId** | **string** | Application (client) id of the service principal. | 
**ClientSecret** | **string** | Client secret of the service principal. | 
**TenantId** | **string** | Directory (tenant) id the service principal lives in. | 

## Methods

### NewServicePrincipalCredentialsIn

`func NewServicePrincipalCredentialsIn(audience string, clientId string, clientSecret string, tenantId string, ) *ServicePrincipalCredentialsIn`

NewServicePrincipalCredentialsIn instantiates a new ServicePrincipalCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePrincipalCredentialsInWithDefaults

`func NewServicePrincipalCredentialsInWithDefaults() *ServicePrincipalCredentialsIn`

NewServicePrincipalCredentialsInWithDefaults instantiates a new ServicePrincipalCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *ServicePrincipalCredentialsIn) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *ServicePrincipalCredentialsIn) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *ServicePrincipalCredentialsIn) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetClientId

`func (o *ServicePrincipalCredentialsIn) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ServicePrincipalCredentialsIn) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ServicePrincipalCredentialsIn) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *ServicePrincipalCredentialsIn) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ServicePrincipalCredentialsIn) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ServicePrincipalCredentialsIn) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetTenantId

`func (o *ServicePrincipalCredentialsIn) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ServicePrincipalCredentialsIn) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ServicePrincipalCredentialsIn) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


