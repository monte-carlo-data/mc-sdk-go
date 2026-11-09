# MotherDuckCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**DbName** | **string** | MotherDuck database to connect to. | 
**Token** | **string** | Access token Monte Carlo signs in to MotherDuck with. Used for this check and not kept. | 

## Methods

### NewMotherDuckCredentialsValidateIn

`func NewMotherDuckCredentialsValidateIn(deploymentId string, dbName string, token string, ) *MotherDuckCredentialsValidateIn`

NewMotherDuckCredentialsValidateIn instantiates a new MotherDuckCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotherDuckCredentialsValidateInWithDefaults

`func NewMotherDuckCredentialsValidateInWithDefaults() *MotherDuckCredentialsValidateIn`

NewMotherDuckCredentialsValidateInWithDefaults instantiates a new MotherDuckCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *MotherDuckCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *MotherDuckCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *MotherDuckCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetDbName

`func (o *MotherDuckCredentialsValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MotherDuckCredentialsValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MotherDuckCredentialsValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetToken

`func (o *MotherDuckCredentialsValidateIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MotherDuckCredentialsValidateIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MotherDuckCredentialsValidateIn) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


