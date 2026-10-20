# HostingPrivateLinkOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | [**NullableAwsPrivateLinkOut**](AwsPrivateLinkOut.md) | AWS private connectivity. Null when the region has none. | 

## Methods

### NewHostingPrivateLinkOut

`func NewHostingPrivateLinkOut(aws NullableAwsPrivateLinkOut, ) *HostingPrivateLinkOut`

NewHostingPrivateLinkOut instantiates a new HostingPrivateLinkOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostingPrivateLinkOutWithDefaults

`func NewHostingPrivateLinkOutWithDefaults() *HostingPrivateLinkOut`

NewHostingPrivateLinkOutWithDefaults instantiates a new HostingPrivateLinkOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *HostingPrivateLinkOut) GetAws() AwsPrivateLinkOut`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *HostingPrivateLinkOut) GetAwsOk() (*AwsPrivateLinkOut, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *HostingPrivateLinkOut) SetAws(v AwsPrivateLinkOut)`

SetAws sets Aws field to given value.


### SetAwsNil

`func (o *HostingPrivateLinkOut) SetAwsNil(b bool)`

 SetAwsNil sets the value for Aws to be an explicit nil

### UnsetAws
`func (o *HostingPrivateLinkOut) UnsetAws()`

UnsetAws ensures that no value is present for Aws, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


