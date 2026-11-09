# MotherDuckCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | **string** | MotherDuck database to connect to. | 
**Token** | **string** | Access token Monte Carlo signs in to MotherDuck with. Stored by Monte Carlo and never returned. | 

## Methods

### NewMotherDuckCredentialsIn

`func NewMotherDuckCredentialsIn(dbName string, token string, ) *MotherDuckCredentialsIn`

NewMotherDuckCredentialsIn instantiates a new MotherDuckCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotherDuckCredentialsInWithDefaults

`func NewMotherDuckCredentialsInWithDefaults() *MotherDuckCredentialsIn`

NewMotherDuckCredentialsInWithDefaults instantiates a new MotherDuckCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *MotherDuckCredentialsIn) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MotherDuckCredentialsIn) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MotherDuckCredentialsIn) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetToken

`func (o *MotherDuckCredentialsIn) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MotherDuckCredentialsIn) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MotherDuckCredentialsIn) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


