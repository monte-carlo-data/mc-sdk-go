# SalesforceDataCloudCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMode** | [**SalesforceDataCloudAuthMode**](SalesforceDataCloudAuthMode.md) | How Monte Carlo signs in. &#x60;oauth_client_credentials&#x60; takes &#x60;consumer_secret&#x60;. &#x60;jwt_bearer&#x60; takes &#x60;username&#x60; and &#x60;private_key&#x60;. Both take &#x60;domain&#x60; and &#x60;consumer_key&#x60;. | 
**Domain** | **string** | My Domain host of your org, such as &#x60;acme.my.salesforce.com&#x60;. | 
**ConsumerKey** | **string** | Consumer key of the connected app Monte Carlo signs in as. | 
**Username** | Pointer to **NullableString** | Salesforce user the JWT signs in as, for &#x60;jwt_bearer&#x60;. | [optional] 
**ConsumerSecret** | Pointer to **NullableString** | Consumer secret of the connected app, for &#x60;oauth_client_credentials&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**PrivateKey** | Pointer to **NullableString** | PEM private key whose certificate is on the connected app, for &#x60;jwt_bearer&#x60;. It must be an unencrypted RSA key. Stored by Monte Carlo and never returned. | [optional] 
**Dataspaces** | Pointer to **[]string** | Data spaces Monte Carlo collects metadata from. | [optional] 
**Dataspace** | Pointer to **NullableString** | The one data space a query connection runs in. Once set, it can be changed but not cleared. Sending it in an update, even unchanged, moves row count collection for that data space&#39;s tables to the query connections using these credentials. | [optional] 

## Methods

### NewSalesforceDataCloudCredentialsIn

`func NewSalesforceDataCloudCredentialsIn(authMode SalesforceDataCloudAuthMode, domain string, consumerKey string, ) *SalesforceDataCloudCredentialsIn`

NewSalesforceDataCloudCredentialsIn instantiates a new SalesforceDataCloudCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceDataCloudCredentialsInWithDefaults

`func NewSalesforceDataCloudCredentialsInWithDefaults() *SalesforceDataCloudCredentialsIn`

NewSalesforceDataCloudCredentialsInWithDefaults instantiates a new SalesforceDataCloudCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMode

`func (o *SalesforceDataCloudCredentialsIn) GetAuthMode() SalesforceDataCloudAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SalesforceDataCloudCredentialsIn) GetAuthModeOk() (*SalesforceDataCloudAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SalesforceDataCloudCredentialsIn) SetAuthMode(v SalesforceDataCloudAuthMode)`

SetAuthMode sets AuthMode field to given value.


### GetDomain

`func (o *SalesforceDataCloudCredentialsIn) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SalesforceDataCloudCredentialsIn) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SalesforceDataCloudCredentialsIn) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetConsumerKey

`func (o *SalesforceDataCloudCredentialsIn) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *SalesforceDataCloudCredentialsIn) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *SalesforceDataCloudCredentialsIn) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.


### GetUsername

`func (o *SalesforceDataCloudCredentialsIn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SalesforceDataCloudCredentialsIn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SalesforceDataCloudCredentialsIn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SalesforceDataCloudCredentialsIn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *SalesforceDataCloudCredentialsIn) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *SalesforceDataCloudCredentialsIn) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetConsumerSecret

`func (o *SalesforceDataCloudCredentialsIn) GetConsumerSecret() string`

GetConsumerSecret returns the ConsumerSecret field if non-nil, zero value otherwise.

### GetConsumerSecretOk

`func (o *SalesforceDataCloudCredentialsIn) GetConsumerSecretOk() (*string, bool)`

GetConsumerSecretOk returns a tuple with the ConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerSecret

`func (o *SalesforceDataCloudCredentialsIn) SetConsumerSecret(v string)`

SetConsumerSecret sets ConsumerSecret field to given value.

### HasConsumerSecret

`func (o *SalesforceDataCloudCredentialsIn) HasConsumerSecret() bool`

HasConsumerSecret returns a boolean if a field has been set.

### SetConsumerSecretNil

`func (o *SalesforceDataCloudCredentialsIn) SetConsumerSecretNil(b bool)`

 SetConsumerSecretNil sets the value for ConsumerSecret to be an explicit nil

### UnsetConsumerSecret
`func (o *SalesforceDataCloudCredentialsIn) UnsetConsumerSecret()`

UnsetConsumerSecret ensures that no value is present for ConsumerSecret, not even an explicit nil
### GetPrivateKey

`func (o *SalesforceDataCloudCredentialsIn) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SalesforceDataCloudCredentialsIn) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SalesforceDataCloudCredentialsIn) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *SalesforceDataCloudCredentialsIn) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *SalesforceDataCloudCredentialsIn) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *SalesforceDataCloudCredentialsIn) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetDataspaces

`func (o *SalesforceDataCloudCredentialsIn) GetDataspaces() []string`

GetDataspaces returns the Dataspaces field if non-nil, zero value otherwise.

### GetDataspacesOk

`func (o *SalesforceDataCloudCredentialsIn) GetDataspacesOk() (*[]string, bool)`

GetDataspacesOk returns a tuple with the Dataspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataspaces

`func (o *SalesforceDataCloudCredentialsIn) SetDataspaces(v []string)`

SetDataspaces sets Dataspaces field to given value.

### HasDataspaces

`func (o *SalesforceDataCloudCredentialsIn) HasDataspaces() bool`

HasDataspaces returns a boolean if a field has been set.

### SetDataspacesNil

`func (o *SalesforceDataCloudCredentialsIn) SetDataspacesNil(b bool)`

 SetDataspacesNil sets the value for Dataspaces to be an explicit nil

### UnsetDataspaces
`func (o *SalesforceDataCloudCredentialsIn) UnsetDataspaces()`

UnsetDataspaces ensures that no value is present for Dataspaces, not even an explicit nil
### GetDataspace

`func (o *SalesforceDataCloudCredentialsIn) GetDataspace() string`

GetDataspace returns the Dataspace field if non-nil, zero value otherwise.

### GetDataspaceOk

`func (o *SalesforceDataCloudCredentialsIn) GetDataspaceOk() (*string, bool)`

GetDataspaceOk returns a tuple with the Dataspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataspace

`func (o *SalesforceDataCloudCredentialsIn) SetDataspace(v string)`

SetDataspace sets Dataspace field to given value.

### HasDataspace

`func (o *SalesforceDataCloudCredentialsIn) HasDataspace() bool`

HasDataspace returns a boolean if a field has been set.

### SetDataspaceNil

`func (o *SalesforceDataCloudCredentialsIn) SetDataspaceNil(b bool)`

 SetDataspaceNil sets the value for Dataspace to be an explicit nil

### UnsetDataspace
`func (o *SalesforceDataCloudCredentialsIn) UnsetDataspace()`

UnsetDataspace ensures that no value is present for Dataspace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


