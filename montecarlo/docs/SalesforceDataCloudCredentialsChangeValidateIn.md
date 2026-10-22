# SalesforceDataCloudCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**AuthMode** | Pointer to [**NullableSalesforceDataCloudAuthMode**](SalesforceDataCloudAuthMode.md) | How Monte Carlo signs in. &#x60;oauth_client_credentials&#x60; takes &#x60;consumer_secret&#x60;. &#x60;jwt_bearer&#x60; takes &#x60;username&#x60; and &#x60;private_key&#x60;. Both take &#x60;domain&#x60; and &#x60;consumer_key&#x60;. | [optional] 
**Domain** | Pointer to **NullableString** | My Domain host of your org, such as &#x60;acme.my.salesforce.com&#x60;. Changing it requires sending the credentials&#39; secrets again. | [optional] 
**ConsumerKey** | Pointer to **NullableString** | Consumer key of the connected app Monte Carlo signs in as. | [optional] 
**Username** | Pointer to **NullableString** | Salesforce user the JWT signs in as, for &#x60;jwt_bearer&#x60;. | [optional] 
**ConsumerSecret** | Pointer to **NullableString** | Consumer secret of the connected app, for &#x60;oauth_client_credentials&#x60;. Stored by Monte Carlo and never returned. | [optional] 
**PrivateKey** | Pointer to **NullableString** | PEM private key whose certificate is on the connected app, for &#x60;jwt_bearer&#x60;. It must be an unencrypted RSA key. Stored by Monte Carlo and never returned. | [optional] 
**Dataspaces** | Pointer to **[]string** | Data spaces Monte Carlo collects metadata from. | [optional] 
**Dataspace** | Pointer to **NullableString** | The one data space a query connection runs in. Once set, it can be changed but not cleared. Sending it in an update, even unchanged, moves row count collection for that data space&#39;s tables to the query connections using these credentials. | [optional] 

## Methods

### NewSalesforceDataCloudCredentialsChangeValidateIn

`func NewSalesforceDataCloudCredentialsChangeValidateIn(deploymentId string, ) *SalesforceDataCloudCredentialsChangeValidateIn`

NewSalesforceDataCloudCredentialsChangeValidateIn instantiates a new SalesforceDataCloudCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceDataCloudCredentialsChangeValidateInWithDefaults

`func NewSalesforceDataCloudCredentialsChangeValidateInWithDefaults() *SalesforceDataCloudCredentialsChangeValidateIn`

NewSalesforceDataCloudCredentialsChangeValidateInWithDefaults instantiates a new SalesforceDataCloudCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetAuthMode

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetAuthMode() SalesforceDataCloudAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetAuthModeOk() (*SalesforceDataCloudAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetAuthMode(v SalesforceDataCloudAuthMode)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### SetAuthModeNil

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetAuthModeNil(b bool)`

 SetAuthModeNil sets the value for AuthMode to be an explicit nil

### UnsetAuthMode
`func (o *SalesforceDataCloudCredentialsChangeValidateIn) UnsetAuthMode()`

UnsetAuthMode ensures that no value is present for AuthMode, not even an explicit nil
### GetDomain

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *SalesforceDataCloudCredentialsChangeValidateIn) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetConsumerKey

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.

### HasConsumerKey

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) HasConsumerKey() bool`

HasConsumerKey returns a boolean if a field has been set.

### SetConsumerKeyNil

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetConsumerKeyNil(b bool)`

 SetConsumerKeyNil sets the value for ConsumerKey to be an explicit nil

### UnsetConsumerKey
`func (o *SalesforceDataCloudCredentialsChangeValidateIn) UnsetConsumerKey()`

UnsetConsumerKey ensures that no value is present for ConsumerKey, not even an explicit nil
### GetUsername

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *SalesforceDataCloudCredentialsChangeValidateIn) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetConsumerSecret

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetConsumerSecret() string`

GetConsumerSecret returns the ConsumerSecret field if non-nil, zero value otherwise.

### GetConsumerSecretOk

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetConsumerSecretOk() (*string, bool)`

GetConsumerSecretOk returns a tuple with the ConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerSecret

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetConsumerSecret(v string)`

SetConsumerSecret sets ConsumerSecret field to given value.

### HasConsumerSecret

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) HasConsumerSecret() bool`

HasConsumerSecret returns a boolean if a field has been set.

### SetConsumerSecretNil

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetConsumerSecretNil(b bool)`

 SetConsumerSecretNil sets the value for ConsumerSecret to be an explicit nil

### UnsetConsumerSecret
`func (o *SalesforceDataCloudCredentialsChangeValidateIn) UnsetConsumerSecret()`

UnsetConsumerSecret ensures that no value is present for ConsumerSecret, not even an explicit nil
### GetPrivateKey

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *SalesforceDataCloudCredentialsChangeValidateIn) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetDataspaces

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetDataspaces() []string`

GetDataspaces returns the Dataspaces field if non-nil, zero value otherwise.

### GetDataspacesOk

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetDataspacesOk() (*[]string, bool)`

GetDataspacesOk returns a tuple with the Dataspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataspaces

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetDataspaces(v []string)`

SetDataspaces sets Dataspaces field to given value.

### HasDataspaces

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) HasDataspaces() bool`

HasDataspaces returns a boolean if a field has been set.

### SetDataspacesNil

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetDataspacesNil(b bool)`

 SetDataspacesNil sets the value for Dataspaces to be an explicit nil

### UnsetDataspaces
`func (o *SalesforceDataCloudCredentialsChangeValidateIn) UnsetDataspaces()`

UnsetDataspaces ensures that no value is present for Dataspaces, not even an explicit nil
### GetDataspace

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetDataspace() string`

GetDataspace returns the Dataspace field if non-nil, zero value otherwise.

### GetDataspaceOk

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) GetDataspaceOk() (*string, bool)`

GetDataspaceOk returns a tuple with the Dataspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataspace

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetDataspace(v string)`

SetDataspace sets Dataspace field to given value.

### HasDataspace

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) HasDataspace() bool`

HasDataspace returns a boolean if a field has been set.

### SetDataspaceNil

`func (o *SalesforceDataCloudCredentialsChangeValidateIn) SetDataspaceNil(b bool)`

 SetDataspaceNil sets the value for Dataspace to be an explicit nil

### UnsetDataspace
`func (o *SalesforceDataCloudCredentialsChangeValidateIn) UnsetDataspace()`

UnsetDataspace ensures that no value is present for Dataspace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


