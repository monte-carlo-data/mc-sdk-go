# DremioCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**DbName** | Pointer to **NullableString** | Dremio project or space to connect to. | [optional] 
**Tls** | Pointer to **NullableBool** | Connect over TLS. Dremio Cloud requires it. | [optional] 
**Host** | **string** | Hostname of the Dremio Arrow Flight endpoint, such as &#x60;data.dremio.cloud&#x60; for Dremio Cloud (&#x60;data.eu.dremio.cloud&#x60; in the EU). | 
**Port** | **int32** | Port Dremio serves Arrow Flight on. | 
**Token** | **string** | Personal access token Monte Carlo signs in to Dremio with. Used for this check and not kept. | 

## Methods

### NewDremioCredentialsValidateIn

`func NewDremioCredentialsValidateIn(deploymentId string, host string, port int32, token string, ) *DremioCredentialsValidateIn`

NewDremioCredentialsValidateIn instantiates a new DremioCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDremioCredentialsValidateInWithDefaults

`func NewDremioCredentialsValidateInWithDefaults() *DremioCredentialsValidateIn`

NewDremioCredentialsValidateInWithDefaults instantiates a new DremioCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *DremioCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DremioCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DremioCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetDbName

`func (o *DremioCredentialsValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DremioCredentialsValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DremioCredentialsValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *DremioCredentialsValidateIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *DremioCredentialsValidateIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *DremioCredentialsValidateIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetTls

`func (o *DremioCredentialsValidateIn) GetTls() bool`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *DremioCredentialsValidateIn) GetTlsOk() (*bool, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *DremioCredentialsValidateIn) SetTls(v bool)`

SetTls sets Tls field to given value.

### HasTls

`func (o *DremioCredentialsValidateIn) HasTls() bool`

HasTls returns a boolean if a field has been set.

### SetTlsNil

`func (o *DremioCredentialsValidateIn) SetTlsNil(b bool)`

 SetTlsNil sets the value for Tls to be an explicit nil

### UnsetTls
`func (o *DremioCredentialsValidateIn) UnsetTls()`

UnsetTls ensures that no value is present for Tls, not even an explicit nil
### GetHost

`func (o *DremioCredentialsValidateIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DremioCredentialsValidateIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DremioCredentialsValidateIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *DremioCredentialsValidateIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DremioCredentialsValidateIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DremioCredentialsValidateIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetToken

`func (o *DremioCredentialsValidateIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DremioCredentialsValidateIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DremioCredentialsValidateIn) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


