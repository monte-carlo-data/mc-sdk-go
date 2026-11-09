# DremioCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | Pointer to **NullableString** | Dremio project or space to connect to. | [optional] 
**Tls** | Pointer to **NullableBool** | Connect over TLS. Dremio Cloud requires it. | [optional] 
**Host** | **string** | Hostname of the Dremio Arrow Flight endpoint, such as &#x60;data.dremio.cloud&#x60; for Dremio Cloud (&#x60;data.eu.dremio.cloud&#x60; in the EU). | 
**Port** | **int32** | Port Dremio serves Arrow Flight on. | 
**Token** | **string** | Personal access token Monte Carlo signs in to Dremio with. Stored by Monte Carlo and never returned. | 

## Methods

### NewDremioCredentialsIn

`func NewDremioCredentialsIn(host string, port int32, token string, ) *DremioCredentialsIn`

NewDremioCredentialsIn instantiates a new DremioCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDremioCredentialsInWithDefaults

`func NewDremioCredentialsInWithDefaults() *DremioCredentialsIn`

NewDremioCredentialsInWithDefaults instantiates a new DremioCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *DremioCredentialsIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DremioCredentialsIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DremioCredentialsIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *DremioCredentialsIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *DremioCredentialsIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *DremioCredentialsIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetTls

`func (o *DremioCredentialsIn) GetTls() bool`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *DremioCredentialsIn) GetTlsOk() (*bool, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *DremioCredentialsIn) SetTls(v bool)`

SetTls sets Tls field to given value.

### HasTls

`func (o *DremioCredentialsIn) HasTls() bool`

HasTls returns a boolean if a field has been set.

### SetTlsNil

`func (o *DremioCredentialsIn) SetTlsNil(b bool)`

 SetTlsNil sets the value for Tls to be an explicit nil

### UnsetTls
`func (o *DremioCredentialsIn) UnsetTls()`

UnsetTls ensures that no value is present for Tls, not even an explicit nil
### GetHost

`func (o *DremioCredentialsIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DremioCredentialsIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DremioCredentialsIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *DremioCredentialsIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DremioCredentialsIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DremioCredentialsIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetToken

`func (o *DremioCredentialsIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DremioCredentialsIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DremioCredentialsIn) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


