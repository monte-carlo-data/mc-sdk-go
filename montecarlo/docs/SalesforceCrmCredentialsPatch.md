# SalesforceCrmCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | Pointer to [**NullableSalesforceCrmAuthMode**](SalesforceCrmAuthMode.md) | How Monte Carlo signs in. &#x60;password&#x60; takes &#x60;username&#x60;, &#x60;password&#x60; and &#x60;security_token&#x60;. &#x60;oauth_client_credentials&#x60; takes &#x60;consumer_key&#x60;, &#x60;consumer_secret&#x60; and &#x60;domain&#x60;. | [optional] 
**Username** | Pointer to **NullableString** | Salesforce user Monte Carlo signs in as, for &#x60;password&#x60;. | [optional] 
**Password** | Pointer to **NullableString** | Password of &#x60;username&#x60;, for &#x60;password&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**SecurityToken** | Pointer to **NullableString** | Security token of &#x60;username&#x60;, for &#x60;password&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**ConsumerKey** | Pointer to **NullableString** | Consumer key of the connected app Monte Carlo signs in as, for &#x60;oauth_client_credentials&#x60;. | [optional] 
**ConsumerSecret** | Pointer to **NullableString** | Consumer secret of the connected app, for &#x60;oauth_client_credentials&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**Domain** | Pointer to **NullableString** | My Domain of your org, such as &#x60;acme.my.salesforce.com&#x60;, for &#x60;oauth_client_credentials&#x60;. The &#x60;.salesforce.com&#x60; suffix is optional. Changing it requires sending the credentials&#39; secrets again. | [optional] 

## Methods

### NewSalesforceCrmCredentialsPatch

`func NewSalesforceCrmCredentialsPatch() *SalesforceCrmCredentialsPatch`

NewSalesforceCrmCredentialsPatch instantiates a new SalesforceCrmCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceCrmCredentialsPatchWithDefaults

`func NewSalesforceCrmCredentialsPatchWithDefaults() *SalesforceCrmCredentialsPatch`

NewSalesforceCrmCredentialsPatchWithDefaults instantiates a new SalesforceCrmCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *SalesforceCrmCredentialsPatch) GetAuthMode() SalesforceCrmAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SalesforceCrmCredentialsPatch) GetAuthModeOk() (*SalesforceCrmAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SalesforceCrmCredentialsPatch) SetAuthMode(v SalesforceCrmAuthMode)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *SalesforceCrmCredentialsPatch) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### SetAuthModeNil

`func (o *SalesforceCrmCredentialsPatch) SetAuthModeNil(b bool)`

 SetAuthModeNil sets the value for AuthMode to be an explicit nil

### UnsetAuthMode
`func (o *SalesforceCrmCredentialsPatch) UnsetAuthMode()`

UnsetAuthMode ensures that no value is present for AuthMode, not even an explicit nil
### GetUsername

`func (o *SalesforceCrmCredentialsPatch) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SalesforceCrmCredentialsPatch) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SalesforceCrmCredentialsPatch) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SalesforceCrmCredentialsPatch) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *SalesforceCrmCredentialsPatch) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *SalesforceCrmCredentialsPatch) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *SalesforceCrmCredentialsPatch) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SalesforceCrmCredentialsPatch) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SalesforceCrmCredentialsPatch) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SalesforceCrmCredentialsPatch) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SalesforceCrmCredentialsPatch) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SalesforceCrmCredentialsPatch) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSecurityToken

`func (o *SalesforceCrmCredentialsPatch) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *SalesforceCrmCredentialsPatch) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *SalesforceCrmCredentialsPatch) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *SalesforceCrmCredentialsPatch) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### SetSecurityTokenNil

`func (o *SalesforceCrmCredentialsPatch) SetSecurityTokenNil(b bool)`

 SetSecurityTokenNil sets the value for SecurityToken to be an explicit nil

### UnsetSecurityToken
`func (o *SalesforceCrmCredentialsPatch) UnsetSecurityToken()`

UnsetSecurityToken ensures that no value is present for SecurityToken, not even an explicit nil
### GetConsumerKey

`func (o *SalesforceCrmCredentialsPatch) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *SalesforceCrmCredentialsPatch) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *SalesforceCrmCredentialsPatch) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.

### HasConsumerKey

`func (o *SalesforceCrmCredentialsPatch) HasConsumerKey() bool`

HasConsumerKey returns a boolean if a field has been set.

### SetConsumerKeyNil

`func (o *SalesforceCrmCredentialsPatch) SetConsumerKeyNil(b bool)`

 SetConsumerKeyNil sets the value for ConsumerKey to be an explicit nil

### UnsetConsumerKey
`func (o *SalesforceCrmCredentialsPatch) UnsetConsumerKey()`

UnsetConsumerKey ensures that no value is present for ConsumerKey, not even an explicit nil
### GetConsumerSecret

`func (o *SalesforceCrmCredentialsPatch) GetConsumerSecret() string`

GetConsumerSecret returns the ConsumerSecret field if non-nil, zero value otherwise.

### GetConsumerSecretOk

`func (o *SalesforceCrmCredentialsPatch) GetConsumerSecretOk() (*string, bool)`

GetConsumerSecretOk returns a tuple with the ConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerSecret

`func (o *SalesforceCrmCredentialsPatch) SetConsumerSecret(v string)`

SetConsumerSecret sets ConsumerSecret field to given value.

### HasConsumerSecret

`func (o *SalesforceCrmCredentialsPatch) HasConsumerSecret() bool`

HasConsumerSecret returns a boolean if a field has been set.

### SetConsumerSecretNil

`func (o *SalesforceCrmCredentialsPatch) SetConsumerSecretNil(b bool)`

 SetConsumerSecretNil sets the value for ConsumerSecret to be an explicit nil

### UnsetConsumerSecret
`func (o *SalesforceCrmCredentialsPatch) UnsetConsumerSecret()`

UnsetConsumerSecret ensures that no value is present for ConsumerSecret, not even an explicit nil
### GetDomain

`func (o *SalesforceCrmCredentialsPatch) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SalesforceCrmCredentialsPatch) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SalesforceCrmCredentialsPatch) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SalesforceCrmCredentialsPatch) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *SalesforceCrmCredentialsPatch) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *SalesforceCrmCredentialsPatch) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


