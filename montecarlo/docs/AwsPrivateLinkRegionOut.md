# AwsPrivateLinkRegionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | The AWS region. | 
**VpcId** | **string** | ID of Monte Carlo&#39;s VPC in the region. | 
**DataStoreVpcEndpointId** | **NullableString** | ID of the S3 VPC endpoint Monte Carlo reaches a data store&#39;s bucket through. Use it in an &#x60;aws:SourceVpce&#x60; condition on the bucket policy. Null when the region has none. | 
**AgentVpcEndpointId** | **NullableString** | ID of the Lambda VPC endpoint Monte Carlo invokes a collection agent through. Use it in an &#x60;aws:SourceVpce&#x60; condition on the function&#39;s policy. Null when the region has none. | 

## Methods

### NewAwsPrivateLinkRegionOut

`func NewAwsPrivateLinkRegionOut(region string, vpcId string, dataStoreVpcEndpointId NullableString, agentVpcEndpointId NullableString, ) *AwsPrivateLinkRegionOut`

NewAwsPrivateLinkRegionOut instantiates a new AwsPrivateLinkRegionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsPrivateLinkRegionOutWithDefaults

`func NewAwsPrivateLinkRegionOutWithDefaults() *AwsPrivateLinkRegionOut`

NewAwsPrivateLinkRegionOutWithDefaults instantiates a new AwsPrivateLinkRegionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AwsPrivateLinkRegionOut) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsPrivateLinkRegionOut) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsPrivateLinkRegionOut) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetVpcId

`func (o *AwsPrivateLinkRegionOut) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsPrivateLinkRegionOut) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsPrivateLinkRegionOut) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetDataStoreVpcEndpointId

`func (o *AwsPrivateLinkRegionOut) GetDataStoreVpcEndpointId() string`

GetDataStoreVpcEndpointId returns the DataStoreVpcEndpointId field if non-nil, zero value otherwise.

### GetDataStoreVpcEndpointIdOk

`func (o *AwsPrivateLinkRegionOut) GetDataStoreVpcEndpointIdOk() (*string, bool)`

GetDataStoreVpcEndpointIdOk returns a tuple with the DataStoreVpcEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreVpcEndpointId

`func (o *AwsPrivateLinkRegionOut) SetDataStoreVpcEndpointId(v string)`

SetDataStoreVpcEndpointId sets DataStoreVpcEndpointId field to given value.


### SetDataStoreVpcEndpointIdNil

`func (o *AwsPrivateLinkRegionOut) SetDataStoreVpcEndpointIdNil(b bool)`

 SetDataStoreVpcEndpointIdNil sets the value for DataStoreVpcEndpointId to be an explicit nil

### UnsetDataStoreVpcEndpointId
`func (o *AwsPrivateLinkRegionOut) UnsetDataStoreVpcEndpointId()`

UnsetDataStoreVpcEndpointId ensures that no value is present for DataStoreVpcEndpointId, not even an explicit nil
### GetAgentVpcEndpointId

`func (o *AwsPrivateLinkRegionOut) GetAgentVpcEndpointId() string`

GetAgentVpcEndpointId returns the AgentVpcEndpointId field if non-nil, zero value otherwise.

### GetAgentVpcEndpointIdOk

`func (o *AwsPrivateLinkRegionOut) GetAgentVpcEndpointIdOk() (*string, bool)`

GetAgentVpcEndpointIdOk returns a tuple with the AgentVpcEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVpcEndpointId

`func (o *AwsPrivateLinkRegionOut) SetAgentVpcEndpointId(v string)`

SetAgentVpcEndpointId sets AgentVpcEndpointId field to given value.


### SetAgentVpcEndpointIdNil

`func (o *AwsPrivateLinkRegionOut) SetAgentVpcEndpointIdNil(b bool)`

 SetAgentVpcEndpointIdNil sets the value for AgentVpcEndpointId to be an explicit nil

### UnsetAgentVpcEndpointId
`func (o *AwsPrivateLinkRegionOut) UnsetAgentVpcEndpointId()`

UnsetAgentVpcEndpointId ensures that no value is present for AgentVpcEndpointId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


