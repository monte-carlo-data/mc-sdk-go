# SnowflakeCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **NullableString** | New Snowflake account identifier, without the &#x60;.snowflakecomputing.com&#x60; suffix. | [optional] 
**User** | Pointer to **NullableString** | New Snowflake user. | [optional] 
**PrivateKey** | Pointer to **NullableString** | New private key, as PEM text. Replaces the stored key and its passphrase; send &#x60;private_key_passphrase&#x60; in the same request if the new key has one. | [optional] 
**PrivateKeyPassphrase** | Pointer to **NullableString** | Passphrase of the new private key. Only accepted together with &#x60;private_key&#x60;. | [optional] 
**Warehouse** | Pointer to **NullableString** | New Snowflake virtual warehouse. An explicit null clears it; queries then run in the user&#39;s default. | [optional] 

## Methods

### NewSnowflakeCredentialsPatch

`func NewSnowflakeCredentialsPatch() *SnowflakeCredentialsPatch`

NewSnowflakeCredentialsPatch instantiates a new SnowflakeCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeCredentialsPatchWithDefaults

`func NewSnowflakeCredentialsPatchWithDefaults() *SnowflakeCredentialsPatch`

NewSnowflakeCredentialsPatchWithDefaults instantiates a new SnowflakeCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SnowflakeCredentialsPatch) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SnowflakeCredentialsPatch) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SnowflakeCredentialsPatch) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SnowflakeCredentialsPatch) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *SnowflakeCredentialsPatch) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *SnowflakeCredentialsPatch) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetUser

`func (o *SnowflakeCredentialsPatch) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SnowflakeCredentialsPatch) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SnowflakeCredentialsPatch) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SnowflakeCredentialsPatch) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *SnowflakeCredentialsPatch) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SnowflakeCredentialsPatch) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPrivateKey

`func (o *SnowflakeCredentialsPatch) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SnowflakeCredentialsPatch) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SnowflakeCredentialsPatch) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *SnowflakeCredentialsPatch) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *SnowflakeCredentialsPatch) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *SnowflakeCredentialsPatch) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetPrivateKeyPassphrase

`func (o *SnowflakeCredentialsPatch) GetPrivateKeyPassphrase() string`

GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field if non-nil, zero value otherwise.

### GetPrivateKeyPassphraseOk

`func (o *SnowflakeCredentialsPatch) GetPrivateKeyPassphraseOk() (*string, bool)`

GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassphrase

`func (o *SnowflakeCredentialsPatch) SetPrivateKeyPassphrase(v string)`

SetPrivateKeyPassphrase sets PrivateKeyPassphrase field to given value.

### HasPrivateKeyPassphrase

`func (o *SnowflakeCredentialsPatch) HasPrivateKeyPassphrase() bool`

HasPrivateKeyPassphrase returns a boolean if a field has been set.

### SetPrivateKeyPassphraseNil

`func (o *SnowflakeCredentialsPatch) SetPrivateKeyPassphraseNil(b bool)`

 SetPrivateKeyPassphraseNil sets the value for PrivateKeyPassphrase to be an explicit nil

### UnsetPrivateKeyPassphrase
`func (o *SnowflakeCredentialsPatch) UnsetPrivateKeyPassphrase()`

UnsetPrivateKeyPassphrase ensures that no value is present for PrivateKeyPassphrase, not even an explicit nil
### GetWarehouse

`func (o *SnowflakeCredentialsPatch) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *SnowflakeCredentialsPatch) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *SnowflakeCredentialsPatch) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *SnowflakeCredentialsPatch) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### SetWarehouseNil

`func (o *SnowflakeCredentialsPatch) SetWarehouseNil(b bool)`

 SetWarehouseNil sets the value for Warehouse to be an explicit nil

### UnsetWarehouse
`func (o *SnowflakeCredentialsPatch) UnsetWarehouse()`

UnsetWarehouse ensures that no value is present for Warehouse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


