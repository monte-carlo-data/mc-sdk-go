# AthenaCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**AssumableRole** | Pointer to **NullableString** | ARN of the IAM role in your AWS account that the deployment assumes. | [optional] 
**ExternalId** | Pointer to **NullableString** | External id your role&#39;s trust policy requires, if it requires one. It is not the deployment&#39;s &#x60;aws_external_id&#x60;. | [optional] 
**AwsRegion** | Pointer to **NullableString** | AWS region to call. Omit it to use the deployment&#39;s own region. | [optional] 
**Catalog** | Pointer to **NullableString** | Athena data catalog to query. Omit it to use &#x60;AwsDataCatalog&#x60;. | [optional] 
**Workgroup** | Pointer to **NullableString** | Athena workgroup to run queries in. Omit it to use the &#x60;primary&#x60; workgroup. The role needs permission to start queries in it. | [optional] 

## Methods

### NewAthenaCredentialsChangeValidateIn

`func NewAthenaCredentialsChangeValidateIn(deploymentId string, ) *AthenaCredentialsChangeValidateIn`

NewAthenaCredentialsChangeValidateIn instantiates a new AthenaCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAthenaCredentialsChangeValidateInWithDefaults

`func NewAthenaCredentialsChangeValidateInWithDefaults() *AthenaCredentialsChangeValidateIn`

NewAthenaCredentialsChangeValidateInWithDefaults instantiates a new AthenaCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *AthenaCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AthenaCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AthenaCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetAssumableRole

`func (o *AthenaCredentialsChangeValidateIn) GetAssumableRole() string`

GetAssumableRole returns the AssumableRole field if non-nil, zero value otherwise.

### GetAssumableRoleOk

`func (o *AthenaCredentialsChangeValidateIn) GetAssumableRoleOk() (*string, bool)`

GetAssumableRoleOk returns a tuple with the AssumableRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumableRole

`func (o *AthenaCredentialsChangeValidateIn) SetAssumableRole(v string)`

SetAssumableRole sets AssumableRole field to given value.

### HasAssumableRole

`func (o *AthenaCredentialsChangeValidateIn) HasAssumableRole() bool`

HasAssumableRole returns a boolean if a field has been set.

### SetAssumableRoleNil

`func (o *AthenaCredentialsChangeValidateIn) SetAssumableRoleNil(b bool)`

 SetAssumableRoleNil sets the value for AssumableRole to be an explicit nil

### UnsetAssumableRole
`func (o *AthenaCredentialsChangeValidateIn) UnsetAssumableRole()`

UnsetAssumableRole ensures that no value is present for AssumableRole, not even an explicit nil
### GetExternalId

`func (o *AthenaCredentialsChangeValidateIn) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AthenaCredentialsChangeValidateIn) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AthenaCredentialsChangeValidateIn) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AthenaCredentialsChangeValidateIn) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AthenaCredentialsChangeValidateIn) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AthenaCredentialsChangeValidateIn) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetAwsRegion

`func (o *AthenaCredentialsChangeValidateIn) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AthenaCredentialsChangeValidateIn) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AthenaCredentialsChangeValidateIn) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *AthenaCredentialsChangeValidateIn) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### SetAwsRegionNil

`func (o *AthenaCredentialsChangeValidateIn) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *AthenaCredentialsChangeValidateIn) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetCatalog

`func (o *AthenaCredentialsChangeValidateIn) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *AthenaCredentialsChangeValidateIn) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *AthenaCredentialsChangeValidateIn) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *AthenaCredentialsChangeValidateIn) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *AthenaCredentialsChangeValidateIn) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *AthenaCredentialsChangeValidateIn) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
### GetWorkgroup

`func (o *AthenaCredentialsChangeValidateIn) GetWorkgroup() string`

GetWorkgroup returns the Workgroup field if non-nil, zero value otherwise.

### GetWorkgroupOk

`func (o *AthenaCredentialsChangeValidateIn) GetWorkgroupOk() (*string, bool)`

GetWorkgroupOk returns a tuple with the Workgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkgroup

`func (o *AthenaCredentialsChangeValidateIn) SetWorkgroup(v string)`

SetWorkgroup sets Workgroup field to given value.

### HasWorkgroup

`func (o *AthenaCredentialsChangeValidateIn) HasWorkgroup() bool`

HasWorkgroup returns a boolean if a field has been set.

### SetWorkgroupNil

`func (o *AthenaCredentialsChangeValidateIn) SetWorkgroupNil(b bool)`

 SetWorkgroupNil sets the value for Workgroup to be an explicit nil

### UnsetWorkgroup
`func (o *AthenaCredentialsChangeValidateIn) UnsetWorkgroup()`

UnsetWorkgroup ensures that no value is present for Workgroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


