# SparkCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | Host name of the Thrift server, for &#x60;binary&#x60;. | [optional] 
**Port** | Pointer to **NullableInt32** | Port of the Thrift server, such as &#x60;10000&#x60;, for &#x60;binary&#x60;. | [optional] 
**DbName** | Pointer to **NullableString** | Database Monte Carlo connects to, such as &#x60;default&#x60;, for &#x60;binary&#x60;. | [optional] 
**Url** | Pointer to **NullableString** | URL of the Thrift server&#39;s HTTP endpoint, for &#x60;http&#x60;, such as &#x60;https://thrift.example.com:10001/cliservice&#x60;. Its path is the HTTP path. For a Databricks cluster, use &#x60;https://&lt;workspace-host&gt;:443/sql/protocolv1/o/&lt;org-id&gt;/&lt;cluster-id&gt;&#x60;, not the &#x60;jdbc:&#x60; URL. It cannot hold a user, a password, a query or a fragment. | [optional] 
**Mode** | [**SparkMode**](SparkMode.md) | How Monte Carlo reaches the Spark Thrift server. &#x60;binary&#x60; takes &#x60;host&#x60;, &#x60;port&#x60; and &#x60;db_name&#x60;. &#x60;http&#x60; takes &#x60;url&#x60; and &#x60;password&#x60;. | 
**User** | **string** | User Monte Carlo signs in as. | 
**Password** | Pointer to **NullableString** | Password of &#x60;user&#x60;. &#x60;http&#x60; needs one. &#x60;binary&#x60; does not send it to the server, so it is optional there. Stored by Monte Carlo and never returned. | [optional] 

## Methods

### NewSparkCredentialsIn

`func NewSparkCredentialsIn(mode SparkMode, user string, ) *SparkCredentialsIn`

NewSparkCredentialsIn instantiates a new SparkCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSparkCredentialsInWithDefaults

`func NewSparkCredentialsInWithDefaults() *SparkCredentialsIn`

NewSparkCredentialsInWithDefaults instantiates a new SparkCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SparkCredentialsIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SparkCredentialsIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SparkCredentialsIn) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SparkCredentialsIn) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SparkCredentialsIn) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SparkCredentialsIn) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *SparkCredentialsIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SparkCredentialsIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SparkCredentialsIn) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SparkCredentialsIn) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *SparkCredentialsIn) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *SparkCredentialsIn) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *SparkCredentialsIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *SparkCredentialsIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *SparkCredentialsIn) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *SparkCredentialsIn) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *SparkCredentialsIn) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *SparkCredentialsIn) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetUrl

`func (o *SparkCredentialsIn) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SparkCredentialsIn) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SparkCredentialsIn) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SparkCredentialsIn) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SparkCredentialsIn) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SparkCredentialsIn) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetMode

`func (o *SparkCredentialsIn) GetMode() SparkMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SparkCredentialsIn) GetModeOk() (*SparkMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SparkCredentialsIn) SetMode(v SparkMode)`

SetMode sets Mode field to given value.


### GetUser

`func (o *SparkCredentialsIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SparkCredentialsIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SparkCredentialsIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *SparkCredentialsIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SparkCredentialsIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SparkCredentialsIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SparkCredentialsIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SparkCredentialsIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SparkCredentialsIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


