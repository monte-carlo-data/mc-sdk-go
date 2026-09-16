# SnowflakeCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Snowflake account identifier, such as &#x60;xy12345.us-east-1&#x60;. Without the &#x60;.snowflakecomputing.com&#x60; suffix. | 
**User** | **string** | Snowflake user the key pair belongs to. | 
**PrivateKey** | **string** | The private key of the pair, as PEM text including its BEGIN and END lines. Stored by Monte Carlo and never returned. | 
**PrivateKeyPassphrase** | Pointer to **NullableString** | Passphrase the private key is encrypted with. Omit it for an unencrypted key. Never returned. | [optional] 
**Warehouse** | Pointer to **NullableString** | Snowflake virtual warehouse to run queries in. Omit it to use the user&#39;s default. | [optional] 

## Methods

### NewSnowflakeCredentialsIn

`func NewSnowflakeCredentialsIn(account string, user string, privateKey string, ) *SnowflakeCredentialsIn`

NewSnowflakeCredentialsIn instantiates a new SnowflakeCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeCredentialsInWithDefaults

`func NewSnowflakeCredentialsInWithDefaults() *SnowflakeCredentialsIn`

NewSnowflakeCredentialsInWithDefaults instantiates a new SnowflakeCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SnowflakeCredentialsIn) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SnowflakeCredentialsIn) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SnowflakeCredentialsIn) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetUser

`func (o *SnowflakeCredentialsIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SnowflakeCredentialsIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SnowflakeCredentialsIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPrivateKey

`func (o *SnowflakeCredentialsIn) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SnowflakeCredentialsIn) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SnowflakeCredentialsIn) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetPrivateKeyPassphrase

`func (o *SnowflakeCredentialsIn) GetPrivateKeyPassphrase() string`

GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field if non-nil, zero value otherwise.

### GetPrivateKeyPassphraseOk

`func (o *SnowflakeCredentialsIn) GetPrivateKeyPassphraseOk() (*string, bool)`

GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassphrase

`func (o *SnowflakeCredentialsIn) SetPrivateKeyPassphrase(v string)`

SetPrivateKeyPassphrase sets PrivateKeyPassphrase field to given value.

### HasPrivateKeyPassphrase

`func (o *SnowflakeCredentialsIn) HasPrivateKeyPassphrase() bool`

HasPrivateKeyPassphrase returns a boolean if a field has been set.

### SetPrivateKeyPassphraseNil

`func (o *SnowflakeCredentialsIn) SetPrivateKeyPassphraseNil(b bool)`

 SetPrivateKeyPassphraseNil sets the value for PrivateKeyPassphrase to be an explicit nil

### UnsetPrivateKeyPassphrase
`func (o *SnowflakeCredentialsIn) UnsetPrivateKeyPassphrase()`

UnsetPrivateKeyPassphrase ensures that no value is present for PrivateKeyPassphrase, not even an explicit nil
### GetWarehouse

`func (o *SnowflakeCredentialsIn) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *SnowflakeCredentialsIn) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *SnowflakeCredentialsIn) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *SnowflakeCredentialsIn) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### SetWarehouseNil

`func (o *SnowflakeCredentialsIn) SetWarehouseNil(b bool)`

 SetWarehouseNil sets the value for Warehouse to be an explicit nil

### UnsetWarehouse
`func (o *SnowflakeCredentialsIn) UnsetWarehouse()`

UnsetWarehouse ensures that no value is present for Warehouse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


