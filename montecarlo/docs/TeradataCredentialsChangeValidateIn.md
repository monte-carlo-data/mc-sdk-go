# TeradataCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Host** | Pointer to **NullableString** | Hostname of the database endpoint. | [optional] 
**Port** | Pointer to **NullableInt32** | Port the database listens on. | [optional] 
**DbName** | Pointer to **NullableString** | Database to connect to. | [optional] 
**User** | Pointer to **NullableString** | Database user Monte Carlo logs in as. | [optional] 
**Password** | Pointer to **NullableString** | New password of the database user. | [optional] 
**SslCaData** | Pointer to **NullableString** | PEM text of the CA certificate the server&#39;s certificate is checked against. | [optional] 
**SslDisabled** | Pointer to **NullableBool** | Do not check the server against &#x60;ssl_ca_data&#x60;. &#x60;td_sslmode&#x60; decides whether the connection uses TLS. | [optional] 
**TdSslmode** | Pointer to [**NullableTeradataSslMode**](TeradataSslMode.md) | How the connection to Teradata uses TLS. | [optional] 
**TdLogmech** | Pointer to [**NullableTeradataLogonMechanism**](TeradataLogonMechanism.md) | How Teradata authenticates the user. | [optional] 

## Methods

### NewTeradataCredentialsChangeValidateIn

`func NewTeradataCredentialsChangeValidateIn(deploymentId string, ) *TeradataCredentialsChangeValidateIn`

NewTeradataCredentialsChangeValidateIn instantiates a new TeradataCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeradataCredentialsChangeValidateInWithDefaults

`func NewTeradataCredentialsChangeValidateInWithDefaults() *TeradataCredentialsChangeValidateIn`

NewTeradataCredentialsChangeValidateInWithDefaults instantiates a new TeradataCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *TeradataCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *TeradataCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *TeradataCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetHost

`func (o *TeradataCredentialsChangeValidateIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TeradataCredentialsChangeValidateIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TeradataCredentialsChangeValidateIn) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TeradataCredentialsChangeValidateIn) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *TeradataCredentialsChangeValidateIn) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *TeradataCredentialsChangeValidateIn) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *TeradataCredentialsChangeValidateIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TeradataCredentialsChangeValidateIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TeradataCredentialsChangeValidateIn) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *TeradataCredentialsChangeValidateIn) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *TeradataCredentialsChangeValidateIn) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *TeradataCredentialsChangeValidateIn) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *TeradataCredentialsChangeValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *TeradataCredentialsChangeValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *TeradataCredentialsChangeValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *TeradataCredentialsChangeValidateIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *TeradataCredentialsChangeValidateIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *TeradataCredentialsChangeValidateIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetUser

`func (o *TeradataCredentialsChangeValidateIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TeradataCredentialsChangeValidateIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TeradataCredentialsChangeValidateIn) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *TeradataCredentialsChangeValidateIn) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *TeradataCredentialsChangeValidateIn) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *TeradataCredentialsChangeValidateIn) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *TeradataCredentialsChangeValidateIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TeradataCredentialsChangeValidateIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TeradataCredentialsChangeValidateIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TeradataCredentialsChangeValidateIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TeradataCredentialsChangeValidateIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TeradataCredentialsChangeValidateIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSslCaData

`func (o *TeradataCredentialsChangeValidateIn) GetSslCaData() string`

GetSslCaData returns the SslCaData field if non-nil, zero value otherwise.

### GetSslCaDataOk

`func (o *TeradataCredentialsChangeValidateIn) GetSslCaDataOk() (*string, bool)`

GetSslCaDataOk returns a tuple with the SslCaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaData

`func (o *TeradataCredentialsChangeValidateIn) SetSslCaData(v string)`

SetSslCaData sets SslCaData field to given value.

### HasSslCaData

`func (o *TeradataCredentialsChangeValidateIn) HasSslCaData() bool`

HasSslCaData returns a boolean if a field has been set.

### SetSslCaDataNil

`func (o *TeradataCredentialsChangeValidateIn) SetSslCaDataNil(b bool)`

 SetSslCaDataNil sets the value for SslCaData to be an explicit nil

### UnsetSslCaData
`func (o *TeradataCredentialsChangeValidateIn) UnsetSslCaData()`

UnsetSslCaData ensures that no value is present for SslCaData, not even an explicit nil
### GetSslDisabled

`func (o *TeradataCredentialsChangeValidateIn) GetSslDisabled() bool`

GetSslDisabled returns the SslDisabled field if non-nil, zero value otherwise.

### GetSslDisabledOk

`func (o *TeradataCredentialsChangeValidateIn) GetSslDisabledOk() (*bool, bool)`

GetSslDisabledOk returns a tuple with the SslDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDisabled

`func (o *TeradataCredentialsChangeValidateIn) SetSslDisabled(v bool)`

SetSslDisabled sets SslDisabled field to given value.

### HasSslDisabled

`func (o *TeradataCredentialsChangeValidateIn) HasSslDisabled() bool`

HasSslDisabled returns a boolean if a field has been set.

### SetSslDisabledNil

`func (o *TeradataCredentialsChangeValidateIn) SetSslDisabledNil(b bool)`

 SetSslDisabledNil sets the value for SslDisabled to be an explicit nil

### UnsetSslDisabled
`func (o *TeradataCredentialsChangeValidateIn) UnsetSslDisabled()`

UnsetSslDisabled ensures that no value is present for SslDisabled, not even an explicit nil
### GetTdSslmode

`func (o *TeradataCredentialsChangeValidateIn) GetTdSslmode() TeradataSslMode`

GetTdSslmode returns the TdSslmode field if non-nil, zero value otherwise.

### GetTdSslmodeOk

`func (o *TeradataCredentialsChangeValidateIn) GetTdSslmodeOk() (*TeradataSslMode, bool)`

GetTdSslmodeOk returns a tuple with the TdSslmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdSslmode

`func (o *TeradataCredentialsChangeValidateIn) SetTdSslmode(v TeradataSslMode)`

SetTdSslmode sets TdSslmode field to given value.

### HasTdSslmode

`func (o *TeradataCredentialsChangeValidateIn) HasTdSslmode() bool`

HasTdSslmode returns a boolean if a field has been set.

### SetTdSslmodeNil

`func (o *TeradataCredentialsChangeValidateIn) SetTdSslmodeNil(b bool)`

 SetTdSslmodeNil sets the value for TdSslmode to be an explicit nil

### UnsetTdSslmode
`func (o *TeradataCredentialsChangeValidateIn) UnsetTdSslmode()`

UnsetTdSslmode ensures that no value is present for TdSslmode, not even an explicit nil
### GetTdLogmech

`func (o *TeradataCredentialsChangeValidateIn) GetTdLogmech() TeradataLogonMechanism`

GetTdLogmech returns the TdLogmech field if non-nil, zero value otherwise.

### GetTdLogmechOk

`func (o *TeradataCredentialsChangeValidateIn) GetTdLogmechOk() (*TeradataLogonMechanism, bool)`

GetTdLogmechOk returns a tuple with the TdLogmech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdLogmech

`func (o *TeradataCredentialsChangeValidateIn) SetTdLogmech(v TeradataLogonMechanism)`

SetTdLogmech sets TdLogmech field to given value.

### HasTdLogmech

`func (o *TeradataCredentialsChangeValidateIn) HasTdLogmech() bool`

HasTdLogmech returns a boolean if a field has been set.

### SetTdLogmechNil

`func (o *TeradataCredentialsChangeValidateIn) SetTdLogmechNil(b bool)`

 SetTdLogmechNil sets the value for TdLogmech to be an explicit nil

### UnsetTdLogmech
`func (o *TeradataCredentialsChangeValidateIn) UnsetTdLogmech()`

UnsetTdLogmech ensures that no value is present for TdLogmech, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


