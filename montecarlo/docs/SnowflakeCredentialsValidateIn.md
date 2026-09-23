# SnowflakeCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**Account** | **string** | Snowflake account identifier, such as &#x60;xy12345.us-east-1&#x60;. Without the &#x60;.snowflakecomputing.com&#x60; suffix. | 
**User** | **string** | Snowflake user the key pair belongs to. | 
**PrivateKey** | **string** | The private key of the pair, as PEM text including its BEGIN and END lines. Used for this check and not kept. | 
**PrivateKeyPassphrase** | Pointer to **NullableString** | Passphrase the private key is encrypted with. Omit it for an unencrypted key. Never returned. | [optional] 
**Warehouse** | Pointer to **NullableString** | Snowflake virtual warehouse to run queries in. Omit it to use the user&#39;s default. | [optional] 

## Methods

### NewSnowflakeCredentialsValidateIn

`func NewSnowflakeCredentialsValidateIn(deploymentId string, account string, user string, privateKey string, ) *SnowflakeCredentialsValidateIn`

NewSnowflakeCredentialsValidateIn instantiates a new SnowflakeCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeCredentialsValidateInWithDefaults

`func NewSnowflakeCredentialsValidateInWithDefaults() *SnowflakeCredentialsValidateIn`

NewSnowflakeCredentialsValidateInWithDefaults instantiates a new SnowflakeCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *SnowflakeCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *SnowflakeCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *SnowflakeCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetAccount

`func (o *SnowflakeCredentialsValidateIn) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SnowflakeCredentialsValidateIn) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SnowflakeCredentialsValidateIn) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetUser

`func (o *SnowflakeCredentialsValidateIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SnowflakeCredentialsValidateIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SnowflakeCredentialsValidateIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPrivateKey

`func (o *SnowflakeCredentialsValidateIn) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SnowflakeCredentialsValidateIn) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SnowflakeCredentialsValidateIn) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetPrivateKeyPassphrase

`func (o *SnowflakeCredentialsValidateIn) GetPrivateKeyPassphrase() string`

GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field if non-nil, zero value otherwise.

### GetPrivateKeyPassphraseOk

`func (o *SnowflakeCredentialsValidateIn) GetPrivateKeyPassphraseOk() (*string, bool)`

GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassphrase

`func (o *SnowflakeCredentialsValidateIn) SetPrivateKeyPassphrase(v string)`

SetPrivateKeyPassphrase sets PrivateKeyPassphrase field to given value.

### HasPrivateKeyPassphrase

`func (o *SnowflakeCredentialsValidateIn) HasPrivateKeyPassphrase() bool`

HasPrivateKeyPassphrase returns a boolean if a field has been set.

### SetPrivateKeyPassphraseNil

`func (o *SnowflakeCredentialsValidateIn) SetPrivateKeyPassphraseNil(b bool)`

 SetPrivateKeyPassphraseNil sets the value for PrivateKeyPassphrase to be an explicit nil

### UnsetPrivateKeyPassphrase
`func (o *SnowflakeCredentialsValidateIn) UnsetPrivateKeyPassphrase()`

UnsetPrivateKeyPassphrase ensures that no value is present for PrivateKeyPassphrase, not even an explicit nil
### GetWarehouse

`func (o *SnowflakeCredentialsValidateIn) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *SnowflakeCredentialsValidateIn) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *SnowflakeCredentialsValidateIn) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *SnowflakeCredentialsValidateIn) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### SetWarehouseNil

`func (o *SnowflakeCredentialsValidateIn) SetWarehouseNil(b bool)`

 SetWarehouseNil sets the value for Warehouse to be an explicit nil

### UnsetWarehouse
`func (o *SnowflakeCredentialsValidateIn) UnsetWarehouse()`

UnsetWarehouse ensures that no value is present for Warehouse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


