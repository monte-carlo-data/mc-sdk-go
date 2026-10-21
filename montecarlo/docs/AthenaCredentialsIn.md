# AthenaCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssumableRole** | **string** | ARN of the IAM role in your AWS account that the deployment assumes. | 
**ExternalId** | Pointer to **NullableString** | External id your role&#39;s trust policy requires, if it requires one. It is not the deployment&#39;s &#x60;aws_external_id&#x60;. | [optional] 
**AwsRegion** | Pointer to **NullableString** | AWS region to call. Omit it to use the deployment&#39;s own region. | [optional] 
**Catalog** | Pointer to **NullableString** | Athena data catalog to query. Omit it to use &#x60;AwsDataCatalog&#x60;. | [optional] 
**Workgroup** | Pointer to **NullableString** | Athena workgroup to run queries in. Omit it to use the &#x60;primary&#x60; workgroup. The role needs permission to start queries in it. | [optional] 

## Methods

### NewAthenaCredentialsIn

`func NewAthenaCredentialsIn(assumableRole string, ) *AthenaCredentialsIn`

NewAthenaCredentialsIn instantiates a new AthenaCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAthenaCredentialsInWithDefaults

`func NewAthenaCredentialsInWithDefaults() *AthenaCredentialsIn`

NewAthenaCredentialsInWithDefaults instantiates a new AthenaCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssumableRole

`func (o *AthenaCredentialsIn) GetAssumableRole() string`

GetAssumableRole returns the AssumableRole field if non-nil, zero value otherwise.

### GetAssumableRoleOk

`func (o *AthenaCredentialsIn) GetAssumableRoleOk() (*string, bool)`

GetAssumableRoleOk returns a tuple with the AssumableRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumableRole

`func (o *AthenaCredentialsIn) SetAssumableRole(v string)`

SetAssumableRole sets AssumableRole field to given value.


### GetExternalId

`func (o *AthenaCredentialsIn) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AthenaCredentialsIn) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AthenaCredentialsIn) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AthenaCredentialsIn) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AthenaCredentialsIn) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AthenaCredentialsIn) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetAwsRegion

`func (o *AthenaCredentialsIn) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AthenaCredentialsIn) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AthenaCredentialsIn) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *AthenaCredentialsIn) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### SetAwsRegionNil

`func (o *AthenaCredentialsIn) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *AthenaCredentialsIn) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetCatalog

`func (o *AthenaCredentialsIn) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *AthenaCredentialsIn) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *AthenaCredentialsIn) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *AthenaCredentialsIn) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *AthenaCredentialsIn) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *AthenaCredentialsIn) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
### GetWorkgroup

`func (o *AthenaCredentialsIn) GetWorkgroup() string`

GetWorkgroup returns the Workgroup field if non-nil, zero value otherwise.

### GetWorkgroupOk

`func (o *AthenaCredentialsIn) GetWorkgroupOk() (*string, bool)`

GetWorkgroupOk returns a tuple with the Workgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkgroup

`func (o *AthenaCredentialsIn) SetWorkgroup(v string)`

SetWorkgroup sets Workgroup field to given value.

### HasWorkgroup

`func (o *AthenaCredentialsIn) HasWorkgroup() bool`

HasWorkgroup returns a boolean if a field has been set.

### SetWorkgroupNil

`func (o *AthenaCredentialsIn) SetWorkgroupNil(b bool)`

 SetWorkgroupNil sets the value for Workgroup to be an explicit nil

### UnsetWorkgroup
`func (o *AthenaCredentialsIn) UnsetWorkgroup()`

UnsetWorkgroup ensures that no value is present for Workgroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


