# AwsCollectionAgentPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LambdaFunctionArn** | Pointer to **NullableString** | ARN of the Lambda function Monte Carlo should invoke. | [optional] 
**RoleArn** | Pointer to **NullableString** | ARN of the role Monte Carlo assumes to invoke the function. Its trust policy must already carry the deployment&#39;s external id. | [optional] 
**Name** | Pointer to **NullableString** | Display name for the collection agent. Replaces the name it currently has. | [optional] 

## Methods

### NewAwsCollectionAgentPatch

`func NewAwsCollectionAgentPatch() *AwsCollectionAgentPatch`

NewAwsCollectionAgentPatch instantiates a new AwsCollectionAgentPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCollectionAgentPatchWithDefaults

`func NewAwsCollectionAgentPatchWithDefaults() *AwsCollectionAgentPatch`

NewAwsCollectionAgentPatchWithDefaults instantiates a new AwsCollectionAgentPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLambdaFunctionArn

`func (o *AwsCollectionAgentPatch) GetLambdaFunctionArn() string`

GetLambdaFunctionArn returns the LambdaFunctionArn field if non-nil, zero value otherwise.

### GetLambdaFunctionArnOk

`func (o *AwsCollectionAgentPatch) GetLambdaFunctionArnOk() (*string, bool)`

GetLambdaFunctionArnOk returns a tuple with the LambdaFunctionArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaFunctionArn

`func (o *AwsCollectionAgentPatch) SetLambdaFunctionArn(v string)`

SetLambdaFunctionArn sets LambdaFunctionArn field to given value.

### HasLambdaFunctionArn

`func (o *AwsCollectionAgentPatch) HasLambdaFunctionArn() bool`

HasLambdaFunctionArn returns a boolean if a field has been set.

### SetLambdaFunctionArnNil

`func (o *AwsCollectionAgentPatch) SetLambdaFunctionArnNil(b bool)`

 SetLambdaFunctionArnNil sets the value for LambdaFunctionArn to be an explicit nil

### UnsetLambdaFunctionArn
`func (o *AwsCollectionAgentPatch) UnsetLambdaFunctionArn()`

UnsetLambdaFunctionArn ensures that no value is present for LambdaFunctionArn, not even an explicit nil
### GetRoleArn

`func (o *AwsCollectionAgentPatch) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsCollectionAgentPatch) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsCollectionAgentPatch) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *AwsCollectionAgentPatch) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *AwsCollectionAgentPatch) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *AwsCollectionAgentPatch) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetName

`func (o *AwsCollectionAgentPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsCollectionAgentPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsCollectionAgentPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsCollectionAgentPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsCollectionAgentPatch) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsCollectionAgentPatch) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


