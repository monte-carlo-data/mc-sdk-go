# AwsPrivateLinkOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The AWS account Monte Carlo collects from. | 
**PrincipalArn** | **string** | ARN of the IAM role Monte Carlo reaches your collection agent and data store with. | 
**PrivateSubnetCidrs** | **[]string** | CIDR blocks of the private subnets Monte Carlo collects from. Allow these rather than individual private IP addresses, which AWS can change. | 
**Regions** | [**[]AwsPrivateLinkRegionOut**](AwsPrivateLinkRegionOut.md) | The regions where Monte Carlo has a VPC for collection. | 

## Methods

### NewAwsPrivateLinkOut

`func NewAwsPrivateLinkOut(accountId string, principalArn string, privateSubnetCidrs []string, regions []AwsPrivateLinkRegionOut, ) *AwsPrivateLinkOut`

NewAwsPrivateLinkOut instantiates a new AwsPrivateLinkOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsPrivateLinkOutWithDefaults

`func NewAwsPrivateLinkOutWithDefaults() *AwsPrivateLinkOut`

NewAwsPrivateLinkOutWithDefaults instantiates a new AwsPrivateLinkOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AwsPrivateLinkOut) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AwsPrivateLinkOut) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AwsPrivateLinkOut) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetPrincipalArn

`func (o *AwsPrivateLinkOut) GetPrincipalArn() string`

GetPrincipalArn returns the PrincipalArn field if non-nil, zero value otherwise.

### GetPrincipalArnOk

`func (o *AwsPrivateLinkOut) GetPrincipalArnOk() (*string, bool)`

GetPrincipalArnOk returns a tuple with the PrincipalArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalArn

`func (o *AwsPrivateLinkOut) SetPrincipalArn(v string)`

SetPrincipalArn sets PrincipalArn field to given value.


### GetPrivateSubnetCidrs

`func (o *AwsPrivateLinkOut) GetPrivateSubnetCidrs() []string`

GetPrivateSubnetCidrs returns the PrivateSubnetCidrs field if non-nil, zero value otherwise.

### GetPrivateSubnetCidrsOk

`func (o *AwsPrivateLinkOut) GetPrivateSubnetCidrsOk() (*[]string, bool)`

GetPrivateSubnetCidrsOk returns a tuple with the PrivateSubnetCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnetCidrs

`func (o *AwsPrivateLinkOut) SetPrivateSubnetCidrs(v []string)`

SetPrivateSubnetCidrs sets PrivateSubnetCidrs field to given value.


### GetRegions

`func (o *AwsPrivateLinkOut) GetRegions() []AwsPrivateLinkRegionOut`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AwsPrivateLinkOut) GetRegionsOk() (*[]AwsPrivateLinkRegionOut, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AwsPrivateLinkOut) SetRegions(v []AwsPrivateLinkRegionOut)`

SetRegions sets Regions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


