# WarehousePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | New display name for the warehouse. Omit it to leave the name unchanged. An explicit null is ignored, the same as omitting the field. | [optional] 

## Methods

### NewWarehousePatch

`func NewWarehousePatch() *WarehousePatch`

NewWarehousePatch instantiates a new WarehousePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehousePatchWithDefaults

`func NewWarehousePatchWithDefaults() *WarehousePatch`

NewWarehousePatchWithDefaults instantiates a new WarehousePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WarehousePatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WarehousePatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WarehousePatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WarehousePatch) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WarehousePatch) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WarehousePatch) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


