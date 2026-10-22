# SalesforceDataCloudCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**AuthMode** | [**SalesforceDataCloudAuthMode**](SalesforceDataCloudAuthMode.md) | How Monte Carlo signs in. &#x60;oauth_client_credentials&#x60; takes &#x60;consumer_secret&#x60;. &#x60;jwt_bearer&#x60; takes &#x60;username&#x60; and &#x60;private_key&#x60;. Both take &#x60;domain&#x60; and &#x60;consumer_key&#x60;. | 
**Domain** | **string** | My Domain host of your org, such as &#x60;acme.my.salesforce.com&#x60;. | 
**ConsumerKey** | **string** | Consumer key of the connected app Monte Carlo signs in as. | 
**Username** | Pointer to **NullableString** | Salesforce user the JWT signs in as, for &#x60;jwt_bearer&#x60;. | [optional] 
**ConsumerSecret** | Pointer to **NullableString** | Consumer secret of the connected app, for &#x60;oauth_client_credentials&#x60;. Used for this check and not kept. | [optional] 
**PrivateKey** | Pointer to **NullableString** | PEM private key whose certificate is on the connected app, for &#x60;jwt_bearer&#x60;. It must be an unencrypted RSA key. Used for this check and not kept. | [optional] 
**Dataspaces** | Pointer to **[]string** | Data spaces Monte Carlo collects metadata from. | [optional] 
**Dataspace** | Pointer to **NullableString** | The one data space a query connection runs in. Once set, it can be changed but not cleared. Sending it in an update, even unchanged, moves row count collection for that data space&#39;s tables to the query connections using these credentials. | [optional] 

## Methods

### NewSalesforceDataCloudCredentialsValidateIn

`func NewSalesforceDataCloudCredentialsValidateIn(deploymentId string, authMode SalesforceDataCloudAuthMode, domain string, consumerKey string, ) *SalesforceDataCloudCredentialsValidateIn`

NewSalesforceDataCloudCredentialsValidateIn instantiates a new SalesforceDataCloudCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceDataCloudCredentialsValidateInWithDefaults

`func NewSalesforceDataCloudCredentialsValidateInWithDefaults() *SalesforceDataCloudCredentialsValidateIn`

NewSalesforceDataCloudCredentialsValidateInWithDefaults instantiates a new SalesforceDataCloudCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *SalesforceDataCloudCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *SalesforceDataCloudCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *SalesforceDataCloudCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetAuthMode

`func (o *SalesforceDataCloudCredentialsValidateIn) GetAuthMode() SalesforceDataCloudAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SalesforceDataCloudCredentialsValidateIn) GetAuthModeOk() (*SalesforceDataCloudAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SalesforceDataCloudCredentialsValidateIn) SetAuthMode(v SalesforceDataCloudAuthMode)`

SetAuthMode sets AuthMode field to given value.


### GetDomain

`func (o *SalesforceDataCloudCredentialsValidateIn) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SalesforceDataCloudCredentialsValidateIn) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SalesforceDataCloudCredentialsValidateIn) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetConsumerKey

`func (o *SalesforceDataCloudCredentialsValidateIn) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *SalesforceDataCloudCredentialsValidateIn) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *SalesforceDataCloudCredentialsValidateIn) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.


### GetUsername

`func (o *SalesforceDataCloudCredentialsValidateIn) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SalesforceDataCloudCredentialsValidateIn) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SalesforceDataCloudCredentialsValidateIn) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SalesforceDataCloudCredentialsValidateIn) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *SalesforceDataCloudCredentialsValidateIn) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *SalesforceDataCloudCredentialsValidateIn) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetConsumerSecret

`func (o *SalesforceDataCloudCredentialsValidateIn) GetConsumerSecret() string`

GetConsumerSecret returns the ConsumerSecret field if non-nil, zero value otherwise.

### GetConsumerSecretOk

`func (o *SalesforceDataCloudCredentialsValidateIn) GetConsumerSecretOk() (*string, bool)`

GetConsumerSecretOk returns a tuple with the ConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerSecret

`func (o *SalesforceDataCloudCredentialsValidateIn) SetConsumerSecret(v string)`

SetConsumerSecret sets ConsumerSecret field to given value.

### HasConsumerSecret

`func (o *SalesforceDataCloudCredentialsValidateIn) HasConsumerSecret() bool`

HasConsumerSecret returns a boolean if a field has been set.

### SetConsumerSecretNil

`func (o *SalesforceDataCloudCredentialsValidateIn) SetConsumerSecretNil(b bool)`

 SetConsumerSecretNil sets the value for ConsumerSecret to be an explicit nil

### UnsetConsumerSecret
`func (o *SalesforceDataCloudCredentialsValidateIn) UnsetConsumerSecret()`

UnsetConsumerSecret ensures that no value is present for ConsumerSecret, not even an explicit nil
### GetPrivateKey

`func (o *SalesforceDataCloudCredentialsValidateIn) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SalesforceDataCloudCredentialsValidateIn) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SalesforceDataCloudCredentialsValidateIn) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *SalesforceDataCloudCredentialsValidateIn) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *SalesforceDataCloudCredentialsValidateIn) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *SalesforceDataCloudCredentialsValidateIn) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetDataspaces

`func (o *SalesforceDataCloudCredentialsValidateIn) GetDataspaces() []string`

GetDataspaces returns the Dataspaces field if non-nil, zero value otherwise.

### GetDataspacesOk

`func (o *SalesforceDataCloudCredentialsValidateIn) GetDataspacesOk() (*[]string, bool)`

GetDataspacesOk returns a tuple with the Dataspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataspaces

`func (o *SalesforceDataCloudCredentialsValidateIn) SetDataspaces(v []string)`

SetDataspaces sets Dataspaces field to given value.

### HasDataspaces

`func (o *SalesforceDataCloudCredentialsValidateIn) HasDataspaces() bool`

HasDataspaces returns a boolean if a field has been set.

### SetDataspacesNil

`func (o *SalesforceDataCloudCredentialsValidateIn) SetDataspacesNil(b bool)`

 SetDataspacesNil sets the value for Dataspaces to be an explicit nil

### UnsetDataspaces
`func (o *SalesforceDataCloudCredentialsValidateIn) UnsetDataspaces()`

UnsetDataspaces ensures that no value is present for Dataspaces, not even an explicit nil
### GetDataspace

`func (o *SalesforceDataCloudCredentialsValidateIn) GetDataspace() string`

GetDataspace returns the Dataspace field if non-nil, zero value otherwise.

### GetDataspaceOk

`func (o *SalesforceDataCloudCredentialsValidateIn) GetDataspaceOk() (*string, bool)`

GetDataspaceOk returns a tuple with the Dataspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataspace

`func (o *SalesforceDataCloudCredentialsValidateIn) SetDataspace(v string)`

SetDataspace sets Dataspace field to given value.

### HasDataspace

`func (o *SalesforceDataCloudCredentialsValidateIn) HasDataspace() bool`

HasDataspace returns a boolean if a field has been set.

### SetDataspaceNil

`func (o *SalesforceDataCloudCredentialsValidateIn) SetDataspaceNil(b bool)`

 SetDataspaceNil sets the value for Dataspace to be an explicit nil

### UnsetDataspace
`func (o *SalesforceDataCloudCredentialsValidateIn) UnsetDataspace()`

UnsetDataspace ensures that no value is present for Dataspace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


