# SalesforceCrmCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**AuthMode** | Pointer to [**NullableSalesforceCrmAuthMode**](SalesforceCrmAuthMode.md) | How Monte Carlo signs in. &#x60;password&#x60; takes &#x60;username&#x60;, &#x60;password&#x60; and &#x60;security_token&#x60;. &#x60;oauth_client_credentials&#x60; takes &#x60;consumer_key&#x60;, &#x60;consumer_secret&#x60; and &#x60;domain&#x60;. | [optional] 
**Username** | Pointer to **NullableString** | Salesforce user Monte Carlo signs in as, for &#x60;password&#x60;. | [optional] 
**Password** | Pointer to **NullableString** | Password of &#x60;username&#x60;, for &#x60;password&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**SecurityToken** | Pointer to **NullableString** | Security token of &#x60;username&#x60;, for &#x60;password&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**ConsumerKey** | Pointer to **NullableString** | Consumer key of the connected app Monte Carlo signs in as, for &#x60;oauth_client_credentials&#x60;. | [optional] 
**ConsumerSecret** | Pointer to **NullableString** | Consumer secret of the connected app, for &#x60;oauth_client_credentials&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**Domain** | Pointer to **NullableString** | My Domain of your org, such as &#x60;acme.my.salesforce.com&#x60;, for &#x60;oauth_client_credentials&#x60;. The &#x60;.salesforce.com&#x60; suffix is optional. Changing it requires sending the credentials&#39; secrets again. | [optional] 

## Methods

### NewSalesforceCrmCredentialsChangeValidateIn

`func NewSalesforceCrmCredentialsChangeValidateIn(deploymentId string, ) *SalesforceCrmCredentialsChangeValidateIn`

NewSalesforceCrmCredentialsChangeValidateIn instantiates a new SalesforceCrmCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceCrmCredentialsChangeValidateInWithDefaults

`func NewSalesforceCrmCredentialsChangeValidateInWithDefaults() *SalesforceCrmCredentialsChangeValidateIn`

NewSalesforceCrmCredentialsChangeValidateInWithDefaults instantiates a new SalesforceCrmCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetAuthMode

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetAuthMode() SalesforceCrmAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetAuthModeOk() (*SalesforceCrmAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetAuthMode(v SalesforceCrmAuthMode)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *SalesforceCrmCredentialsChangeValidateIn) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### SetAuthModeNil

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetAuthModeNil(b bool)`

 SetAuthModeNil sets the value for AuthMode to be an explicit nil

### UnsetAuthMode
`func (o *SalesforceCrmCredentialsChangeValidateIn) UnsetAuthMode()`

UnsetAuthMode ensures that no value is present for AuthMode, not even an explicit nil
### GetUsername

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SalesforceCrmCredentialsChangeValidateIn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *SalesforceCrmCredentialsChangeValidateIn) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SalesforceCrmCredentialsChangeValidateIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SalesforceCrmCredentialsChangeValidateIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSecurityToken

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *SalesforceCrmCredentialsChangeValidateIn) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### SetSecurityTokenNil

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetSecurityTokenNil(b bool)`

 SetSecurityTokenNil sets the value for SecurityToken to be an explicit nil

### UnsetSecurityToken
`func (o *SalesforceCrmCredentialsChangeValidateIn) UnsetSecurityToken()`

UnsetSecurityToken ensures that no value is present for SecurityToken, not even an explicit nil
### GetConsumerKey

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.

### HasConsumerKey

`func (o *SalesforceCrmCredentialsChangeValidateIn) HasConsumerKey() bool`

HasConsumerKey returns a boolean if a field has been set.

### SetConsumerKeyNil

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetConsumerKeyNil(b bool)`

 SetConsumerKeyNil sets the value for ConsumerKey to be an explicit nil

### UnsetConsumerKey
`func (o *SalesforceCrmCredentialsChangeValidateIn) UnsetConsumerKey()`

UnsetConsumerKey ensures that no value is present for ConsumerKey, not even an explicit nil
### GetConsumerSecret

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetConsumerSecret() string`

GetConsumerSecret returns the ConsumerSecret field if non-nil, zero value otherwise.

### GetConsumerSecretOk

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetConsumerSecretOk() (*string, bool)`

GetConsumerSecretOk returns a tuple with the ConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerSecret

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetConsumerSecret(v string)`

SetConsumerSecret sets ConsumerSecret field to given value.

### HasConsumerSecret

`func (o *SalesforceCrmCredentialsChangeValidateIn) HasConsumerSecret() bool`

HasConsumerSecret returns a boolean if a field has been set.

### SetConsumerSecretNil

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetConsumerSecretNil(b bool)`

 SetConsumerSecretNil sets the value for ConsumerSecret to be an explicit nil

### UnsetConsumerSecret
`func (o *SalesforceCrmCredentialsChangeValidateIn) UnsetConsumerSecret()`

UnsetConsumerSecret ensures that no value is present for ConsumerSecret, not even an explicit nil
### GetDomain

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SalesforceCrmCredentialsChangeValidateIn) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SalesforceCrmCredentialsChangeValidateIn) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *SalesforceCrmCredentialsChangeValidateIn) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *SalesforceCrmCredentialsChangeValidateIn) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


