# HiveMySqlCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname of the database endpoint. | 
**Port** | **int32** | Port the database listens on. | 
**User** | **string** | Database user Monte Carlo logs in as. | 
**Password** | **string** | Password of the database user. Stored by Monte Carlo and never returned. | 
**DbName** | **string** | MySQL database that holds the Hive metastore. | 
**SslCaUrl** | Pointer to **NullableString** | https URL of the PEM bundle of certificate authorities the MySQL server&#39;s certificate is checked against. Monte Carlo downloads it at each connection and connects over TLS. For Amazon RDS, use https://truststore.pki.rds.amazonaws.com/global/global-bundle.pem. Leave it out and the server&#39;s certificate is not checked. | [optional] 

## Methods

### NewHiveMySqlCredentialsIn

`func NewHiveMySqlCredentialsIn(host string, port int32, user string, password string, dbName string, ) *HiveMySqlCredentialsIn`

NewHiveMySqlCredentialsIn instantiates a new HiveMySqlCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHiveMySqlCredentialsInWithDefaults

`func NewHiveMySqlCredentialsInWithDefaults() *HiveMySqlCredentialsIn`

NewHiveMySqlCredentialsInWithDefaults instantiates a new HiveMySqlCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HiveMySqlCredentialsIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HiveMySqlCredentialsIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HiveMySqlCredentialsIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *HiveMySqlCredentialsIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HiveMySqlCredentialsIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HiveMySqlCredentialsIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *HiveMySqlCredentialsIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *HiveMySqlCredentialsIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *HiveMySqlCredentialsIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *HiveMySqlCredentialsIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HiveMySqlCredentialsIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HiveMySqlCredentialsIn) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDbName

`func (o *HiveMySqlCredentialsIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *HiveMySqlCredentialsIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *HiveMySqlCredentialsIn) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetSslCaUrl

`func (o *HiveMySqlCredentialsIn) GetSslCaUrl() string`

GetSslCaUrl returns the SslCaUrl field if non-nil, zero value otherwise.

### GetSslCaUrlOk

`func (o *HiveMySqlCredentialsIn) GetSslCaUrlOk() (*string, bool)`

GetSslCaUrlOk returns a tuple with the SslCaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCaUrl

`func (o *HiveMySqlCredentialsIn) SetSslCaUrl(v string)`

SetSslCaUrl sets SslCaUrl field to given value.

### HasSslCaUrl

`func (o *HiveMySqlCredentialsIn) HasSslCaUrl() bool`

HasSslCaUrl returns a boolean if a field has been set.

### SetSslCaUrlNil

`func (o *HiveMySqlCredentialsIn) SetSslCaUrlNil(b bool)`

 SetSslCaUrlNil sets the value for SslCaUrl to be an explicit nil

### UnsetSslCaUrl
`func (o *HiveMySqlCredentialsIn) UnsetSslCaUrl()`

UnsetSslCaUrl ensures that no value is present for SslCaUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


