# AzureDedicatedSqlPoolCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname of the database endpoint. | 
**Port** | **int32** | Port the database listens on. | 
**User** | **string** | Database user Monte Carlo logs in as. | 
**Password** | **string** | Password of the database user. Stored by Monte Carlo and never returned. | 
**DbName** | **string** | Database to connect to. | 

## Methods

### NewAzureDedicatedSqlPoolCredentialsIn

`func NewAzureDedicatedSqlPoolCredentialsIn(host string, port int32, user string, password string, dbName string, ) *AzureDedicatedSqlPoolCredentialsIn`

NewAzureDedicatedSqlPoolCredentialsIn instantiates a new AzureDedicatedSqlPoolCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDedicatedSqlPoolCredentialsInWithDefaults

`func NewAzureDedicatedSqlPoolCredentialsInWithDefaults() *AzureDedicatedSqlPoolCredentialsIn`

NewAzureDedicatedSqlPoolCredentialsInWithDefaults instantiates a new AzureDedicatedSqlPoolCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *AzureDedicatedSqlPoolCredentialsIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AzureDedicatedSqlPoolCredentialsIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AzureDedicatedSqlPoolCredentialsIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *AzureDedicatedSqlPoolCredentialsIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AzureDedicatedSqlPoolCredentialsIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AzureDedicatedSqlPoolCredentialsIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *AzureDedicatedSqlPoolCredentialsIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AzureDedicatedSqlPoolCredentialsIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AzureDedicatedSqlPoolCredentialsIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *AzureDedicatedSqlPoolCredentialsIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AzureDedicatedSqlPoolCredentialsIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AzureDedicatedSqlPoolCredentialsIn) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDbName

`func (o *AzureDedicatedSqlPoolCredentialsIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *AzureDedicatedSqlPoolCredentialsIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *AzureDedicatedSqlPoolCredentialsIn) SetDbName(v string)`

SetDbName sets DbName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


