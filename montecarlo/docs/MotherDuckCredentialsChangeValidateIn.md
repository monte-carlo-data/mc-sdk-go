# MotherDuckCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**DbName** | Pointer to **NullableString** | MotherDuck database to connect to. | [optional] 
**Token** | Pointer to **NullableString** | New access token. | [optional] 

## Methods

### NewMotherDuckCredentialsChangeValidateIn

`func NewMotherDuckCredentialsChangeValidateIn(deploymentId string, ) *MotherDuckCredentialsChangeValidateIn`

NewMotherDuckCredentialsChangeValidateIn instantiates a new MotherDuckCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotherDuckCredentialsChangeValidateInWithDefaults

`func NewMotherDuckCredentialsChangeValidateInWithDefaults() *MotherDuckCredentialsChangeValidateIn`

NewMotherDuckCredentialsChangeValidateInWithDefaults instantiates a new MotherDuckCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *MotherDuckCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *MotherDuckCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *MotherDuckCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetDbName

`func (o *MotherDuckCredentialsChangeValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MotherDuckCredentialsChangeValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MotherDuckCredentialsChangeValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *MotherDuckCredentialsChangeValidateIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *MotherDuckCredentialsChangeValidateIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *MotherDuckCredentialsChangeValidateIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetToken

`func (o *MotherDuckCredentialsChangeValidateIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MotherDuckCredentialsChangeValidateIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MotherDuckCredentialsChangeValidateIn) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MotherDuckCredentialsChangeValidateIn) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *MotherDuckCredentialsChangeValidateIn) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *MotherDuckCredentialsChangeValidateIn) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


