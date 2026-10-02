# BiContainerPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | New display name for the BI container. Omit it to leave the name unchanged. An explicit null is ignored, the same as omitting the field. | [optional] 

## Methods

### NewBiContainerPatch

`func NewBiContainerPatch() *BiContainerPatch`

NewBiContainerPatch instantiates a new BiContainerPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiContainerPatchWithDefaults

`func NewBiContainerPatchWithDefaults() *BiContainerPatch`

NewBiContainerPatchWithDefaults instantiates a new BiContainerPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BiContainerPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BiContainerPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BiContainerPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BiContainerPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BiContainerPatch) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BiContainerPatch) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


