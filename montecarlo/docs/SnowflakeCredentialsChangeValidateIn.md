# SnowflakeCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Account** | Pointer to **NullableString** | New Snowflake account identifier, without the &#x60;.snowflakecomputing.com&#x60; suffix. | [optional] 
**User** | Pointer to **NullableString** | New Snowflake user. | [optional] 
**PrivateKey** | Pointer to **NullableString** | New private key, as PEM text. Replaces the stored key and its passphrase; send &#x60;private_key_passphrase&#x60; in the same request if the new key has one. | [optional] 
**PrivateKeyPassphrase** | Pointer to **NullableString** | Passphrase of the new private key. Only accepted together with &#x60;private_key&#x60;. | [optional] 
**Warehouse** | Pointer to **NullableString** | New Snowflake virtual warehouse. An explicit null clears it; queries then run in the user&#39;s default. | [optional] 

## Methods

### NewSnowflakeCredentialsChangeValidateIn

`func NewSnowflakeCredentialsChangeValidateIn(deploymentId string, ) *SnowflakeCredentialsChangeValidateIn`

NewSnowflakeCredentialsChangeValidateIn instantiates a new SnowflakeCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeCredentialsChangeValidateInWithDefaults

`func NewSnowflakeCredentialsChangeValidateInWithDefaults() *SnowflakeCredentialsChangeValidateIn`

NewSnowflakeCredentialsChangeValidateInWithDefaults instantiates a new SnowflakeCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *SnowflakeCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *SnowflakeCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *SnowflakeCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetAccount

`func (o *SnowflakeCredentialsChangeValidateIn) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SnowflakeCredentialsChangeValidateIn) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SnowflakeCredentialsChangeValidateIn) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SnowflakeCredentialsChangeValidateIn) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *SnowflakeCredentialsChangeValidateIn) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *SnowflakeCredentialsChangeValidateIn) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetUser

`func (o *SnowflakeCredentialsChangeValidateIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SnowflakeCredentialsChangeValidateIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SnowflakeCredentialsChangeValidateIn) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SnowflakeCredentialsChangeValidateIn) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *SnowflakeCredentialsChangeValidateIn) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SnowflakeCredentialsChangeValidateIn) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPrivateKey

`func (o *SnowflakeCredentialsChangeValidateIn) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SnowflakeCredentialsChangeValidateIn) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SnowflakeCredentialsChangeValidateIn) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *SnowflakeCredentialsChangeValidateIn) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *SnowflakeCredentialsChangeValidateIn) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *SnowflakeCredentialsChangeValidateIn) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetPrivateKeyPassphrase

`func (o *SnowflakeCredentialsChangeValidateIn) GetPrivateKeyPassphrase() string`

GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field if non-nil, zero value otherwise.

### GetPrivateKeyPassphraseOk

`func (o *SnowflakeCredentialsChangeValidateIn) GetPrivateKeyPassphraseOk() (*string, bool)`

GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassphrase

`func (o *SnowflakeCredentialsChangeValidateIn) SetPrivateKeyPassphrase(v string)`

SetPrivateKeyPassphrase sets PrivateKeyPassphrase field to given value.

### HasPrivateKeyPassphrase

`func (o *SnowflakeCredentialsChangeValidateIn) HasPrivateKeyPassphrase() bool`

HasPrivateKeyPassphrase returns a boolean if a field has been set.

### SetPrivateKeyPassphraseNil

`func (o *SnowflakeCredentialsChangeValidateIn) SetPrivateKeyPassphraseNil(b bool)`

 SetPrivateKeyPassphraseNil sets the value for PrivateKeyPassphrase to be an explicit nil

### UnsetPrivateKeyPassphrase
`func (o *SnowflakeCredentialsChangeValidateIn) UnsetPrivateKeyPassphrase()`

UnsetPrivateKeyPassphrase ensures that no value is present for PrivateKeyPassphrase, not even an explicit nil
### GetWarehouse

`func (o *SnowflakeCredentialsChangeValidateIn) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *SnowflakeCredentialsChangeValidateIn) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *SnowflakeCredentialsChangeValidateIn) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *SnowflakeCredentialsChangeValidateIn) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### SetWarehouseNil

`func (o *SnowflakeCredentialsChangeValidateIn) SetWarehouseNil(b bool)`

 SetWarehouseNil sets the value for Warehouse to be an explicit nil

### UnsetWarehouse
`func (o *SnowflakeCredentialsChangeValidateIn) UnsetWarehouse()`

UnsetWarehouse ensures that no value is present for Warehouse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


