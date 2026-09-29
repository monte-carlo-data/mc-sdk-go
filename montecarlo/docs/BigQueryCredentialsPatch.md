# BigQueryCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccountKey** | Pointer to **NullableString** | New JSON key file, as its text. Replaces the stored key whole. | [optional] 

## Methods

### NewBigQueryCredentialsPatch

`func NewBigQueryCredentialsPatch() *BigQueryCredentialsPatch`

NewBigQueryCredentialsPatch instantiates a new BigQueryCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBigQueryCredentialsPatchWithDefaults

`func NewBigQueryCredentialsPatchWithDefaults() *BigQueryCredentialsPatch`

NewBigQueryCredentialsPatchWithDefaults instantiates a new BigQueryCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccountKey

`func (o *BigQueryCredentialsPatch) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *BigQueryCredentialsPatch) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *BigQueryCredentialsPatch) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.

### HasServiceAccountKey

`func (o *BigQueryCredentialsPatch) HasServiceAccountKey() bool`

HasServiceAccountKey returns a boolean if a field has been set.

### SetServiceAccountKeyNil

`func (o *BigQueryCredentialsPatch) SetServiceAccountKeyNil(b bool)`

 SetServiceAccountKeyNil sets the value for ServiceAccountKey to be an explicit nil

### UnsetServiceAccountKey
`func (o *BigQueryCredentialsPatch) UnsetServiceAccountKey()`

UnsetServiceAccountKey ensures that no value is present for ServiceAccountKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


