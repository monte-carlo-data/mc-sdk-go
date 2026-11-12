# HiveCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | Hostname of the HiveServer2. | [optional] 
**Port** | Pointer to **NullableInt32** | Port the HiveServer2 listens on in binary mode. | [optional] 
**User** | Pointer to **NullableString** | User Monte Carlo signs in as. Hive takes no password. | [optional] 
**DbName** | Pointer to **NullableString** | Database Monte Carlo connects to. | [optional] 
**AuthMode** | Pointer to [**NullableHiveAuthMode**](HiveAuthMode.md) | How Monte Carlo signs in. Omit it for &#x60;SASL&#x60;. | [optional] 

## Methods

### NewHiveCredentialsPatch

`func NewHiveCredentialsPatch() *HiveCredentialsPatch`

NewHiveCredentialsPatch instantiates a new HiveCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveCredentialsPatchWithDefaults

`func NewHiveCredentialsPatchWithDefaults() *HiveCredentialsPatch`

NewHiveCredentialsPatchWithDefaults instantiates a new HiveCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HiveCredentialsPatch) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HiveCredentialsPatch) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HiveCredentialsPatch) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HiveCredentialsPatch) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *HiveCredentialsPatch) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *HiveCredentialsPatch) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *HiveCredentialsPatch) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HiveCredentialsPatch) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HiveCredentialsPatch) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HiveCredentialsPatch) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *HiveCredentialsPatch) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *HiveCredentialsPatch) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUser

`func (o *HiveCredentialsPatch) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *HiveCredentialsPatch) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *HiveCredentialsPatch) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *HiveCredentialsPatch) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *HiveCredentialsPatch) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *HiveCredentialsPatch) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetDbName

`func (o *HiveCredentialsPatch) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *HiveCredentialsPatch) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *HiveCredentialsPatch) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *HiveCredentialsPatch) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *HiveCredentialsPatch) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *HiveCredentialsPatch) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetAuthMode

`func (o *HiveCredentialsPatch) GetAuthMode() HiveAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *HiveCredentialsPatch) GetAuthModeOk() (*HiveAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *HiveCredentialsPatch) SetAuthMode(v HiveAuthMode)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *HiveCredentialsPatch) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### SetAuthModeNil

`func (o *HiveCredentialsPatch) SetAuthModeNil(b bool)`

 SetAuthModeNil sets the value for AuthMode to be an explicit nil

### UnsetAuthMode
`func (o *HiveCredentialsPatch) UnsetAuthMode()`

UnsetAuthMode ensures that no value is present for AuthMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


