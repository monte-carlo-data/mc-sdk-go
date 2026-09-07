# AwsCollectionAgentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment to register the collection agent on. It must already hold an unregistered AWS collection agent. | 
**LambdaFunctionArn** | **string** | ARN of the Lambda function Monte Carlo should invoke. | 
**Name** | Pointer to **NullableString** | Display name for the collection agent. Replaces the name it currently has. | [optional] 
**RoleArn** | **string** | ARN of the role Monte Carlo assumes to invoke the function. Its trust policy must already carry the deployment&#39;s external id. | 

## Methods

### NewAwsCollectionAgentIn

`func NewAwsCollectionAgentIn(deploymentId string, lambdaFunctionArn string, roleArn string, ) *AwsCollectionAgentIn`

NewAwsCollectionAgentIn instantiates a new AwsCollectionAgentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsCollectionAgentInWithDefaults

`func NewAwsCollectionAgentInWithDefaults() *AwsCollectionAgentIn`

NewAwsCollectionAgentInWithDefaults instantiates a new AwsCollectionAgentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *AwsCollectionAgentIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AwsCollectionAgentIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AwsCollectionAgentIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetLambdaFunctionArn

`func (o *AwsCollectionAgentIn) GetLambdaFunctionArn() string`

GetLambdaFunctionArn returns the LambdaFunctionArn field if non-nil, zero value otherwise.

### GetLambdaFunctionArnOk

`func (o *AwsCollectionAgentIn) GetLambdaFunctionArnOk() (*string, bool)`

GetLambdaFunctionArnOk returns a tuple with the LambdaFunctionArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaFunctionArn

`func (o *AwsCollectionAgentIn) SetLambdaFunctionArn(v string)`

SetLambdaFunctionArn sets LambdaFunctionArn field to given value.


### GetName

`func (o *AwsCollectionAgentIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsCollectionAgentIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsCollectionAgentIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsCollectionAgentIn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AwsCollectionAgentIn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsCollectionAgentIn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRoleArn

`func (o *AwsCollectionAgentIn) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsCollectionAgentIn) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsCollectionAgentIn) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


