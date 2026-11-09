# DremioCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | Pointer to **NullableString** | Dremio project or space to connect to. | [optional] 
**Tls** | Pointer to **NullableBool** | Connect over TLS. Dremio Cloud requires it. | [optional] 
**Host** | Pointer to **NullableString** | Hostname of the Dremio Arrow Flight endpoint, such as &#x60;data.dremio.cloud&#x60; for Dremio Cloud (&#x60;data.eu.dremio.cloud&#x60; in the EU). Changing it requires sending the credentials&#39; secrets again. | [optional] 
**Port** | Pointer to **NullableInt32** | Port Dremio serves Arrow Flight on. | [optional] 
**Token** | Pointer to **NullableString** | New personal access token. | [optional] 

## Methods

### NewDremioCredentialsPatch

`func NewDremioCredentialsPatch() *DremioCredentialsPatch`

NewDremioCredentialsPatch instantiates a new DremioCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDremioCredentialsPatchWithDefaults

`func NewDremioCredentialsPatchWithDefaults() *DremioCredentialsPatch`

NewDremioCredentialsPatchWithDefaults instantiates a new DremioCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *DremioCredentialsPatch) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DremioCredentialsPatch) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DremioCredentialsPatch) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *DremioCredentialsPatch) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *DremioCredentialsPatch) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *DremioCredentialsPatch) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetTls

`func (o *DremioCredentialsPatch) GetTls() bool`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *DremioCredentialsPatch) GetTlsOk() (*bool, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *DremioCredentialsPatch) SetTls(v bool)`

SetTls sets Tls field to given value.

### HasTls

`func (o *DremioCredentialsPatch) HasTls() bool`

HasTls returns a boolean if a field has been set.

### SetTlsNil

`func (o *DremioCredentialsPatch) SetTlsNil(b bool)`

 SetTlsNil sets the value for Tls to be an explicit nil

### UnsetTls
`func (o *DremioCredentialsPatch) UnsetTls()`

UnsetTls ensures that no value is present for Tls, not even an explicit nil
### GetHost

`func (o *DremioCredentialsPatch) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DremioCredentialsPatch) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DremioCredentialsPatch) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DremioCredentialsPatch) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *DremioCredentialsPatch) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *DremioCredentialsPatch) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *DremioCredentialsPatch) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DremioCredentialsPatch) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DremioCredentialsPatch) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DremioCredentialsPatch) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *DremioCredentialsPatch) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *DremioCredentialsPatch) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetToken

`func (o *DremioCredentialsPatch) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DremioCredentialsPatch) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DremioCredentialsPatch) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DremioCredentialsPatch) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *DremioCredentialsPatch) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *DremioCredentialsPatch) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


