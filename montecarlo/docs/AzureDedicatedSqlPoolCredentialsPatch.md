# AzureDedicatedSqlPoolCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | Hostname of the database endpoint. | [optional] 
**Port** | Pointer to **NullableInt32** | Port the database listens on. | [optional] 
**DbName** | Pointer to **NullableString** | Database to connect to. | [optional] 
**User** | Pointer to **NullableString** | Database user Monte Carlo logs in as. | [optional] 
**Password** | Pointer to **NullableString** | New password of the database user. | [optional] 

## Methods

### NewAzureDedicatedSqlPoolCredentialsPatch

`func NewAzureDedicatedSqlPoolCredentialsPatch() *AzureDedicatedSqlPoolCredentialsPatch`

NewAzureDedicatedSqlPoolCredentialsPatch instantiates a new AzureDedicatedSqlPoolCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDedicatedSqlPoolCredentialsPatchWithDefaults

`func NewAzureDedicatedSqlPoolCredentialsPatchWithDefaults() *AzureDedicatedSqlPoolCredentialsPatch`

NewAzureDedicatedSqlPoolCredentialsPatchWithDefaults instantiates a new AzureDedicatedSqlPoolCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *AzureDedicatedSqlPoolCredentialsPatch) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AzureDedicatedSqlPoolCredentialsPatch) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AzureDedicatedSqlPoolCredentialsPatch) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AzureDedicatedSqlPoolCredentialsPatch) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *AzureDedicatedSqlPoolCredentialsPatch) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *AzureDedicatedSqlPoolCredentialsPatch) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *AzureDedicatedSqlPoolCredentialsPatch) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AzureDedicatedSqlPoolCredentialsPatch) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AzureDedicatedSqlPoolCredentialsPatch) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *AzureDedicatedSqlPoolCredentialsPatch) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *AzureDedicatedSqlPoolCredentialsPatch) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *AzureDedicatedSqlPoolCredentialsPatch) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *AzureDedicatedSqlPoolCredentialsPatch) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *AzureDedicatedSqlPoolCredentialsPatch) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *AzureDedicatedSqlPoolCredentialsPatch) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *AzureDedicatedSqlPoolCredentialsPatch) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *AzureDedicatedSqlPoolCredentialsPatch) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *AzureDedicatedSqlPoolCredentialsPatch) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetUser

`func (o *AzureDedicatedSqlPoolCredentialsPatch) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AzureDedicatedSqlPoolCredentialsPatch) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AzureDedicatedSqlPoolCredentialsPatch) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AzureDedicatedSqlPoolCredentialsPatch) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *AzureDedicatedSqlPoolCredentialsPatch) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *AzureDedicatedSqlPoolCredentialsPatch) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *AzureDedicatedSqlPoolCredentialsPatch) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AzureDedicatedSqlPoolCredentialsPatch) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AzureDedicatedSqlPoolCredentialsPatch) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AzureDedicatedSqlPoolCredentialsPatch) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *AzureDedicatedSqlPoolCredentialsPatch) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *AzureDedicatedSqlPoolCredentialsPatch) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


