# GenericCollectionAgentTokenOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **time.Time** | When the credential was created. | 
**DeploymentId** | **string** | Identifier of the deployment whose agent presents this credential. | 
**Description** | **string** | What this credential is for. | 
**Id** | **string** | Unique identifier of the credential. For a token this is also its key id; for an OAuth client, its client id. | 
**McdId** | **string** | Key id the agent presents, as &#x60;mcd_id&#x60; in its configuration. The same value as &#x60;id&#x60;. | 
**Type** | [**CredentialType**](CredentialType.md) | Which kind of credential this is. | 

## Methods

### NewGenericCollectionAgentTokenOut

`func NewGenericCollectionAgentTokenOut(createdTime time.Time, deploymentId string, description string, id string, mcdId string, type_ CredentialType, ) *GenericCollectionAgentTokenOut`

NewGenericCollectionAgentTokenOut instantiates a new GenericCollectionAgentTokenOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericCollectionAgentTokenOutWithDefaults

`func NewGenericCollectionAgentTokenOutWithDefaults() *GenericCollectionAgentTokenOut`

NewGenericCollectionAgentTokenOutWithDefaults instantiates a new GenericCollectionAgentTokenOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *GenericCollectionAgentTokenOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GenericCollectionAgentTokenOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GenericCollectionAgentTokenOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetDeploymentId

`func (o *GenericCollectionAgentTokenOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GenericCollectionAgentTokenOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GenericCollectionAgentTokenOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetDescription

`func (o *GenericCollectionAgentTokenOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenericCollectionAgentTokenOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenericCollectionAgentTokenOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *GenericCollectionAgentTokenOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenericCollectionAgentTokenOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenericCollectionAgentTokenOut) SetId(v string)`

SetId sets Id field to given value.


### GetMcdId

`func (o *GenericCollectionAgentTokenOut) GetMcdId() string`

GetMcdId returns the McdId field if non-nil, zero value otherwise.

### GetMcdIdOk

`func (o *GenericCollectionAgentTokenOut) GetMcdIdOk() (*string, bool)`

GetMcdIdOk returns a tuple with the McdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcdId

`func (o *GenericCollectionAgentTokenOut) SetMcdId(v string)`

SetMcdId sets McdId field to given value.


### GetType

`func (o *GenericCollectionAgentTokenOut) GetType() CredentialType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericCollectionAgentTokenOut) GetTypeOk() (*CredentialType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericCollectionAgentTokenOut) SetType(v CredentialType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


