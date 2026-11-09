# DremioCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**DbName** | Pointer to **NullableString** | Dremio project or space to connect to. | [optional] 
**Tls** | Pointer to **NullableBool** | Connect over TLS. Dremio Cloud requires it. | [optional] 
**Host** | Pointer to **NullableString** | Hostname of the Dremio Arrow Flight endpoint, such as &#x60;data.dremio.cloud&#x60; for Dremio Cloud (&#x60;data.eu.dremio.cloud&#x60; in the EU). Changing it requires sending the credentials&#39; secrets again. | [optional] 
**Port** | Pointer to **NullableInt32** | Port Dremio serves Arrow Flight on. | [optional] 
**Token** | Pointer to **NullableString** | New personal access token. | [optional] 

## Methods

### NewDremioCredentialsChangeValidateIn

`func NewDremioCredentialsChangeValidateIn(deploymentId string, ) *DremioCredentialsChangeValidateIn`

NewDremioCredentialsChangeValidateIn instantiates a new DremioCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDremioCredentialsChangeValidateInWithDefaults

`func NewDremioCredentialsChangeValidateInWithDefaults() *DremioCredentialsChangeValidateIn`

NewDremioCredentialsChangeValidateInWithDefaults instantiates a new DremioCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *DremioCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DremioCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DremioCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetDbName

`func (o *DremioCredentialsChangeValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DremioCredentialsChangeValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DremioCredentialsChangeValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *DremioCredentialsChangeValidateIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *DremioCredentialsChangeValidateIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *DremioCredentialsChangeValidateIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetTls

`func (o *DremioCredentialsChangeValidateIn) GetTls() bool`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *DremioCredentialsChangeValidateIn) GetTlsOk() (*bool, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *DremioCredentialsChangeValidateIn) SetTls(v bool)`

SetTls sets Tls field to given value.

### HasTls

`func (o *DremioCredentialsChangeValidateIn) HasTls() bool`

HasTls returns a boolean if a field has been set.

### SetTlsNil

`func (o *DremioCredentialsChangeValidateIn) SetTlsNil(b bool)`

 SetTlsNil sets the value for Tls to be an explicit nil

### UnsetTls
`func (o *DremioCredentialsChangeValidateIn) UnsetTls()`

UnsetTls ensures that no value is present for Tls, not even an explicit nil
### GetHost

`func (o *DremioCredentialsChangeValidateIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DremioCredentialsChangeValidateIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DremioCredentialsChangeValidateIn) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DremioCredentialsChangeValidateIn) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *DremioCredentialsChangeValidateIn) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *DremioCredentialsChangeValidateIn) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *DremioCredentialsChangeValidateIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DremioCredentialsChangeValidateIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DremioCredentialsChangeValidateIn) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DremioCredentialsChangeValidateIn) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *DremioCredentialsChangeValidateIn) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *DremioCredentialsChangeValidateIn) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetToken

`func (o *DremioCredentialsChangeValidateIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DremioCredentialsChangeValidateIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DremioCredentialsChangeValidateIn) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DremioCredentialsChangeValidateIn) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *DremioCredentialsChangeValidateIn) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *DremioCredentialsChangeValidateIn) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


