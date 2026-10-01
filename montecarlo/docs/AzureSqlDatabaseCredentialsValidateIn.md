# AzureSqlDatabaseCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Host** | **string** | Hostname of the database endpoint. | 
**Port** | **int32** | Port the database listens on. | 
**User** | **string** | Database user Monte Carlo logs in as. | 
**Password** | **string** | Password of the database user. Used for this check and not kept. | 
**DbName** | **string** | Database to connect to. | 

## Methods

### NewAzureSqlDatabaseCredentialsValidateIn

`func NewAzureSqlDatabaseCredentialsValidateIn(deploymentId string, host string, port int32, user string, password string, dbName string, ) *AzureSqlDatabaseCredentialsValidateIn`

NewAzureSqlDatabaseCredentialsValidateIn instantiates a new AzureSqlDatabaseCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSqlDatabaseCredentialsValidateInWithDefaults

`func NewAzureSqlDatabaseCredentialsValidateInWithDefaults() *AzureSqlDatabaseCredentialsValidateIn`

NewAzureSqlDatabaseCredentialsValidateInWithDefaults instantiates a new AzureSqlDatabaseCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *AzureSqlDatabaseCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AzureSqlDatabaseCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AzureSqlDatabaseCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetHost

`func (o *AzureSqlDatabaseCredentialsValidateIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AzureSqlDatabaseCredentialsValidateIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AzureSqlDatabaseCredentialsValidateIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *AzureSqlDatabaseCredentialsValidateIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AzureSqlDatabaseCredentialsValidateIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AzureSqlDatabaseCredentialsValidateIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *AzureSqlDatabaseCredentialsValidateIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AzureSqlDatabaseCredentialsValidateIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AzureSqlDatabaseCredentialsValidateIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *AzureSqlDatabaseCredentialsValidateIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AzureSqlDatabaseCredentialsValidateIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AzureSqlDatabaseCredentialsValidateIn) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDbName

`func (o *AzureSqlDatabaseCredentialsValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *AzureSqlDatabaseCredentialsValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *AzureSqlDatabaseCredentialsValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


