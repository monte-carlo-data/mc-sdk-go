# SparkCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**Mode** | [**SparkMode**](SparkMode.md) | How Monte Carlo reaches the Spark Thrift server. &#x60;binary&#x60; takes &#x60;host&#x60;, &#x60;port&#x60; and &#x60;db_name&#x60;. &#x60;http&#x60; takes &#x60;url&#x60; and &#x60;password&#x60;. | 
**Host** | **NullableString** | Host name of the Thrift server, for &#x60;binary&#x60;. Null for &#x60;http&#x60;. | 
**Port** | **NullableInt32** | Port of the Thrift server, such as &#x60;10000&#x60;, for &#x60;binary&#x60;. Null for &#x60;http&#x60;. | 
**DbName** | **NullableString** | Database Monte Carlo connects to, such as &#x60;default&#x60;, for &#x60;binary&#x60;. Null for &#x60;http&#x60;. | 
**Url** | **NullableString** | URL of the Thrift server&#39;s HTTP endpoint, for &#x60;http&#x60;, such as &#x60;https://thrift.example.com:10001/cliservice&#x60;. Its path is the HTTP path. For a Databricks cluster, use &#x60;https://&lt;workspace-host&gt;:443/sql/protocolv1/o/&lt;org-id&gt;/&lt;cluster-id&gt;&#x60;, not the &#x60;jdbc:&#x60; URL. It cannot hold a user, a password, a query or a fragment. Null for &#x60;binary&#x60;. | 
**User** | **string** | User Monte Carlo signs in as. | 

## Methods

### NewSparkCredentialsOut

`func NewSparkCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, mode SparkMode, host NullableString, port NullableInt32, dbName NullableString, url NullableString, user string, ) *SparkCredentialsOut`

NewSparkCredentialsOut instantiates a new SparkCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSparkCredentialsOutWithDefaults

`func NewSparkCredentialsOutWithDefaults() *SparkCredentialsOut`

NewSparkCredentialsOutWithDefaults instantiates a new SparkCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SparkCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SparkCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SparkCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *SparkCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SparkCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SparkCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *SparkCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *SparkCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *SparkCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *SparkCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *SparkCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *SparkCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetMode

`func (o *SparkCredentialsOut) GetMode() SparkMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SparkCredentialsOut) GetModeOk() (*SparkMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SparkCredentialsOut) SetMode(v SparkMode)`

SetMode sets Mode field to given value.


### GetHost

`func (o *SparkCredentialsOut) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SparkCredentialsOut) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SparkCredentialsOut) SetHost(v string)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *SparkCredentialsOut) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SparkCredentialsOut) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *SparkCredentialsOut) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SparkCredentialsOut) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SparkCredentialsOut) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *SparkCredentialsOut) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *SparkCredentialsOut) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetDbName

`func (o *SparkCredentialsOut) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *SparkCredentialsOut) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *SparkCredentialsOut) SetDbName(v string)`

SetDbName sets DbName field to given value.


### SetDbNameNil

`func (o *SparkCredentialsOut) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *SparkCredentialsOut) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetUrl

`func (o *SparkCredentialsOut) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SparkCredentialsOut) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SparkCredentialsOut) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *SparkCredentialsOut) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SparkCredentialsOut) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetUser

`func (o *SparkCredentialsOut) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SparkCredentialsOut) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SparkCredentialsOut) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


