# MotherDuckCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | Pointer to **NullableString** | MotherDuck database to connect to. | [optional] 
**Token** | Pointer to **NullableString** | New access token. | [optional] 

## Methods

### NewMotherDuckCredentialsPatch

`func NewMotherDuckCredentialsPatch() *MotherDuckCredentialsPatch`

NewMotherDuckCredentialsPatch instantiates a new MotherDuckCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotherDuckCredentialsPatchWithDefaults

`func NewMotherDuckCredentialsPatchWithDefaults() *MotherDuckCredentialsPatch`

NewMotherDuckCredentialsPatchWithDefaults instantiates a new MotherDuckCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *MotherDuckCredentialsPatch) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MotherDuckCredentialsPatch) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MotherDuckCredentialsPatch) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *MotherDuckCredentialsPatch) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### SetDbNameNil

`func (o *MotherDuckCredentialsPatch) SetDbNameNil(b bool)`

 SetDbNameNil sets the value for DbName to be an explicit nil

### UnsetDbName
`func (o *MotherDuckCredentialsPatch) UnsetDbName()`

UnsetDbName ensures that no value is present for DbName, not even an explicit nil
### GetToken

`func (o *MotherDuckCredentialsPatch) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MotherDuckCredentialsPatch) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MotherDuckCredentialsPatch) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MotherDuckCredentialsPatch) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *MotherDuckCredentialsPatch) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *MotherDuckCredentialsPatch) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


