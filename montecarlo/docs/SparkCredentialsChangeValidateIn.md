# SparkCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Host** | Pointer to **NullableString** | Host name of the Thrift server, for &#x60;binary&#x60;. Changing it requires sending &#x60;password&#x60; again when the credentials hold one. | [optional] 
**Port** | Pointer to **NullableInt32** | Port of the Thrift server, such as &#x60;10000&#x60;, for &#x60;binary&#x60;. | [optional] 
**DbName** | Pointer to **NullableString** | Database Monte Carlo connects to, such as &#x60;default&#x60;, for &#x60;binary&#x60;. | [optional] 
**Url** | Pointer to **NullableString** | URL of the Thrift server&#39;s HTTP endpoint, for &#x60;http&#x60;, such as &#x60;https://thrift.example.com:10001/cliservice&#x60;. Its path is the HTTP path. For a Databricks cluster, use &#x60;https://&lt;workspace-host&gt;:443/sql/protocolv1/o/&lt;org-id&gt;/&lt;cluster-id&gt;&#x60;, not the &#x60;jdbc:&#x60; URL. It cannot hold a user, a password, a query or a fragment. Changing it requires sending &#x60;password&#x60; again when the credentials hold one. | [optional] 
**Mode** | Pointer to [**NullableSparkMode**](SparkMode.md) | How Monte Carlo reaches the Spark Thrift server. &#x60;binary&#x60; takes &#x60;host&#x60;, &#x60;port&#x60; and &#x60;db_name&#x60;. &#x60;http&#x60; takes &#x60;url&#x60; and &#x60;password&#x60;. | [optional] 
**User** | Pointer to **NullableString** | User Monte Carlo signs in as. | [optional] 
**Password** | Pointer to **NullableString** | New password of &#x60;user&#x60;. | [optional] 

## Methods

### NewSparkCredentialsChangeValidateIn

`func NewSparkCredentialsChangeValidateIn(deploymentId string, ) *SparkCredentialsChangeValidateIn`

NewSparkCredentialsChangeValidateIn instantiates a new SparkCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSparkCredentialsChangeValidateInWithDefaults

`func NewSparkCredentialsChangeValidateInWithDefaults() *SparkCredentialsChangeValidateIn`

NewSparkCredentialsChangeValidateInWithDefaults instantiates a new SparkCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *SparkCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *SparkCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *SparkCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetHost

`func (o *SparkCredentialsChangeValidateIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SparkCredentialsChangeValidateIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SparkCredentialsChangeValidateIn) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SparkCredentialsChangeValidateIn) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SparkCredentialsChangeValidateIn) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SparkCredentialsChangeValidateIn) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *SparkCredentialsChangeValidateIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SparkCredentialsChangeValidateIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SparkCredentialsChangeValidateIn) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SparkCredentialsChangeValidateIn) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *SparkCredentialsChangeValidateIn) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *SparkCredentialsChangeValidateIn) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *SparkCredentialsChangeValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *SparkCredentialsChangeValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *SparkCredentialsChangeValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *SparkCredentialsChangeValidateIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *SparkCredentialsChangeValidateIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *SparkCredentialsChangeValidateIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetUrl

`func (o *SparkCredentialsChangeValidateIn) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SparkCredentialsChangeValidateIn) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SparkCredentialsChangeValidateIn) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SparkCredentialsChangeValidateIn) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SparkCredentialsChangeValidateIn) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SparkCredentialsChangeValidateIn) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetMode

`func (o *SparkCredentialsChangeValidateIn) GetMode() SparkMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SparkCredentialsChangeValidateIn) GetModeOk() (*SparkMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SparkCredentialsChangeValidateIn) SetMode(v SparkMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SparkCredentialsChangeValidateIn) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *SparkCredentialsChangeValidateIn) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *SparkCredentialsChangeValidateIn) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetUser

`func (o *SparkCredentialsChangeValidateIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SparkCredentialsChangeValidateIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SparkCredentialsChangeValidateIn) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SparkCredentialsChangeValidateIn) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *SparkCredentialsChangeValidateIn) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SparkCredentialsChangeValidateIn) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *SparkCredentialsChangeValidateIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SparkCredentialsChangeValidateIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SparkCredentialsChangeValidateIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SparkCredentialsChangeValidateIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SparkCredentialsChangeValidateIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SparkCredentialsChangeValidateIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


