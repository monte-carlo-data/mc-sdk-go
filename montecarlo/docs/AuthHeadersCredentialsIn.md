# AuthHeadersCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | **map[string]string** | Header names, each mapped to the value to send for it. Names have to be valid HTTP header names. At most 20 headers, each value up to 8192 characters. | 

## Methods

### NewAuthHeadersCredentialsIn

`func NewAuthHeadersCredentialsIn(headers map[string]string, ) *AuthHeadersCredentialsIn`

NewAuthHeadersCredentialsIn instantiates a new AuthHeadersCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthHeadersCredentialsInWithDefaults

`func NewAuthHeadersCredentialsInWithDefaults() *AuthHeadersCredentialsIn`

NewAuthHeadersCredentialsInWithDefaults instantiates a new AuthHeadersCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *AuthHeadersCredentialsIn) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AuthHeadersCredentialsIn) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AuthHeadersCredentialsIn) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


