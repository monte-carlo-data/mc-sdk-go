# SparkCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | Host name of the Thrift server, for &#x60;binary&#x60;. Changing it requires sending &#x60;password&#x60; again when the credentials hold one. | [optional] 
**Port** | Pointer to **NullableInt32** | Port of the Thrift server, such as &#x60;10000&#x60;, for &#x60;binary&#x60;. | [optional] 
**DbName** | Pointer to **NullableString** | Database Monte Carlo connects to, such as &#x60;default&#x60;, for &#x60;binary&#x60;. | [optional] 
**Url** | Pointer to **NullableString** | URL of the Thrift server&#39;s HTTP endpoint, for &#x60;http&#x60;, such as &#x60;https://thrift.example.com:10001/cliservice&#x60;. Its path is the HTTP path. For a Databricks cluster, use &#x60;https://&lt;workspace-host&gt;:443/sql/protocolv1/o/&lt;org-id&gt;/&lt;cluster-id&gt;&#x60;, not the &#x60;jdbc:&#x60; URL. It cannot hold a user, a password, a query or a fragment. Changing it requires sending &#x60;password&#x60; again when the credentials hold one. | [optional] 
**Mode** | Pointer to [**NullableSparkMode**](SparkMode.md) | How Monte Carlo reaches the Spark Thrift server. &#x60;binary&#x60; takes &#x60;host&#x60;, &#x60;port&#x60; and &#x60;db_name&#x60;. &#x60;http&#x60; takes &#x60;url&#x60; and &#x60;password&#x60;. | [optional] 
**User** | Pointer to **NullableString** | User Monte Carlo signs in as. | [optional] 
**Password** | Pointer to **NullableString** | New password of &#x60;user&#x60;. | [optional] 

## Methods

### NewSparkCredentialsPatch

`func NewSparkCredentialsPatch() *SparkCredentialsPatch`

NewSparkCredentialsPatch instantiates a new SparkCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSparkCredentialsPatchWithDefaults

`func NewSparkCredentialsPatchWithDefaults() *SparkCredentialsPatch`

NewSparkCredentialsPatchWithDefaults instantiates a new SparkCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SparkCredentialsPatch) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SparkCredentialsPatch) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SparkCredentialsPatch) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SparkCredentialsPatch) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SparkCredentialsPatch) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SparkCredentialsPatch) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *SparkCredentialsPatch) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SparkCredentialsPatch) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SparkCredentialsPatch) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SparkCredentialsPatch) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *SparkCredentialsPatch) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *SparkCredentialsPatch) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *SparkCredentialsPatch) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *SparkCredentialsPatch) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *SparkCredentialsPatch) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *SparkCredentialsPatch) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *SparkCredentialsPatch) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *SparkCredentialsPatch) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetUrl

`func (o *SparkCredentialsPatch) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SparkCredentialsPatch) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SparkCredentialsPatch) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SparkCredentialsPatch) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SparkCredentialsPatch) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SparkCredentialsPatch) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetMode

`func (o *SparkCredentialsPatch) GetMode() SparkMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SparkCredentialsPatch) GetModeOk() (*SparkMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SparkCredentialsPatch) SetMode(v SparkMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SparkCredentialsPatch) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *SparkCredentialsPatch) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *SparkCredentialsPatch) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetUser

`func (o *SparkCredentialsPatch) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SparkCredentialsPatch) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SparkCredentialsPatch) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SparkCredentialsPatch) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *SparkCredentialsPatch) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SparkCredentialsPatch) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *SparkCredentialsPatch) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SparkCredentialsPatch) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SparkCredentialsPatch) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SparkCredentialsPatch) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SparkCredentialsPatch) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SparkCredentialsPatch) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


