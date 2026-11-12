# HiveCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname of the HiveServer2. | 
**Port** | Pointer to **NullableInt32** | Port the HiveServer2 listens on in binary mode. | [optional] 
**User** | **string** | User Monte Carlo signs in as. Hive takes no password. | 
**DbName** | Pointer to **NullableString** | Database Monte Carlo connects to. | [optional] 
**AuthMode** | Pointer to [**HiveAuthMode**](HiveAuthMode.md) | How Monte Carlo signs in. Omit it for &#x60;SASL&#x60;. | [optional] [default to HIVEAUTHMODE_SASL]

## Methods

### NewHiveCredentialsIn

`func NewHiveCredentialsIn(host string, user string, ) *HiveCredentialsIn`

NewHiveCredentialsIn instantiates a new HiveCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveCredentialsInWithDefaults

`func NewHiveCredentialsInWithDefaults() *HiveCredentialsIn`

NewHiveCredentialsInWithDefaults instantiates a new HiveCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HiveCredentialsIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HiveCredentialsIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HiveCredentialsIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *HiveCredentialsIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HiveCredentialsIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HiveCredentialsIn) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HiveCredentialsIn) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *HiveCredentialsIn) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *HiveCredentialsIn) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUser

`func (o *HiveCredentialsIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *HiveCredentialsIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *HiveCredentialsIn) SetUser(v string)`

SetUser sets User field to given value.


### GetDbName

`func (o *HiveCredentialsIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *HiveCredentialsIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *HiveCredentialsIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *HiveCredentialsIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *HiveCredentialsIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *HiveCredentialsIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetAuthMode

`func (o *HiveCredentialsIn) GetAuthMode() HiveAuthMode`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *HiveCredentialsIn) GetAuthModeOk() (*HiveAuthMode, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *HiveCredentialsIn) SetAuthMode(v HiveAuthMode)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *HiveCredentialsIn) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


