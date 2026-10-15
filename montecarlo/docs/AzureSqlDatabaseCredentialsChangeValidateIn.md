# AzureSqlDatabaseCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Host** | Pointer to **NullableString** | Hostname of the database endpoint. | [optional] 
**Port** | Pointer to **NullableInt32** | Port the database listens on. | [optional] 
**DbName** | Pointer to **NullableString** | Database to connect to. | [optional] 
**User** | Pointer to **NullableString** | Database user Monte Carlo logs in as. | [optional] 
**Password** | Pointer to **NullableString** | New password of the database user. | [optional] 

## Methods

### NewAzureSqlDatabaseCredentialsChangeValidateIn

`func NewAzureSqlDatabaseCredentialsChangeValidateIn(deploymentId string, ) *AzureSqlDatabaseCredentialsChangeValidateIn`

NewAzureSqlDatabaseCredentialsChangeValidateIn instantiates a new AzureSqlDatabaseCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSqlDatabaseCredentialsChangeValidateInWithDefaults

`func NewAzureSqlDatabaseCredentialsChangeValidateInWithDefaults() *AzureSqlDatabaseCredentialsChangeValidateIn`

NewAzureSqlDatabaseCredentialsChangeValidateInWithDefaults instantiates a new AzureSqlDatabaseCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetHost

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetUser

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AzureSqlDatabaseCredentialsChangeValidateIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


