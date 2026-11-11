# SparkCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Host** | Pointer to **NullableString** | Host name of the Thrift server, for &#x60;binary&#x60;. | [optional] 
**Port** | Pointer to **NullableInt32** | Port of the Thrift server, such as &#x60;10000&#x60;, for &#x60;binary&#x60;. | [optional] 
**DbName** | Pointer to **NullableString** | Database Monte Carlo connects to, such as &#x60;default&#x60;, for &#x60;binary&#x60;. | [optional] 
**Url** | Pointer to **NullableString** | URL of the Thrift server&#39;s HTTP endpoint, for &#x60;http&#x60;, such as &#x60;https://thrift.example.com:10001/cliservice&#x60;. Its path is the HTTP path. For a Databricks cluster, use &#x60;https://&lt;workspace-host&gt;:443/sql/protocolv1/o/&lt;org-id&gt;/&lt;cluster-id&gt;&#x60;, not the &#x60;jdbc:&#x60; URL. It cannot hold a user, a password, a query or a fragment. | [optional] 
**Mode** | [**SparkMode**](SparkMode.md) | How Monte Carlo reaches the Spark Thrift server. &#x60;binary&#x60; takes &#x60;host&#x60;, &#x60;port&#x60; and &#x60;db_name&#x60;. &#x60;http&#x60; takes &#x60;url&#x60; and &#x60;password&#x60;. | 
**User** | **string** | User Monte Carlo signs in as. | 
**Password** | Pointer to **NullableString** | Password of &#x60;user&#x60;. &#x60;http&#x60; needs one, and &#x60;binary&#x60; ignores it. Used for this check and not kept. | [optional] 

## Methods

### NewSparkCredentialsValidateIn

`func NewSparkCredentialsValidateIn(deploymentId string, mode SparkMode, user string, ) *SparkCredentialsValidateIn`

NewSparkCredentialsValidateIn instantiates a new SparkCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSparkCredentialsValidateInWithDefaults

`func NewSparkCredentialsValidateInWithDefaults() *SparkCredentialsValidateIn`

NewSparkCredentialsValidateInWithDefaults instantiates a new SparkCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *SparkCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *SparkCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *SparkCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetHost

`func (o *SparkCredentialsValidateIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SparkCredentialsValidateIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SparkCredentialsValidateIn) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SparkCredentialsValidateIn) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SparkCredentialsValidateIn) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SparkCredentialsValidateIn) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *SparkCredentialsValidateIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SparkCredentialsValidateIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SparkCredentialsValidateIn) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SparkCredentialsValidateIn) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *SparkCredentialsValidateIn) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *SparkCredentialsValidateIn) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *SparkCredentialsValidateIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *SparkCredentialsValidateIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *SparkCredentialsValidateIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *SparkCredentialsValidateIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *SparkCredentialsValidateIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *SparkCredentialsValidateIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetUrl

`func (o *SparkCredentialsValidateIn) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SparkCredentialsValidateIn) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SparkCredentialsValidateIn) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SparkCredentialsValidateIn) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SparkCredentialsValidateIn) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SparkCredentialsValidateIn) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetMode

`func (o *SparkCredentialsValidateIn) GetMode() SparkMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SparkCredentialsValidateIn) GetModeOk() (*SparkMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SparkCredentialsValidateIn) SetMode(v SparkMode)`

SetMode sets Mode field to given value.


### GetUser

`func (o *SparkCredentialsValidateIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SparkCredentialsValidateIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SparkCredentialsValidateIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *SparkCredentialsValidateIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SparkCredentialsValidateIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SparkCredentialsValidateIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SparkCredentialsValidateIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SparkCredentialsValidateIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SparkCredentialsValidateIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


