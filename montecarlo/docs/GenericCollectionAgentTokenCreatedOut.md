# GenericCollectionAgentTokenCreatedOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credential. For a token this is also its key id; for an OAuth client, its client id. | 
**DeploymentId** | **string** | Identifier of the deployment whose agent presents this credential. | 
**Type** | [**CredentialType**](CredentialType.md) | Which kind of credential this is. | 
**Description** | **string** | What this credential is for. | 
**CreatedTime** | **time.Time** | When the credential was created. | 
**McdId** | **string** | Key id the agent presents, as &#x60;mcd_id&#x60; in its configuration. The same value as &#x60;id&#x60;. | 
**McdToken** | **string** | Secret the agent presents, as &#x60;mcd_token&#x60; in its configuration. Returned once, by this call. It is not retrievable afterwards, and Terraform holds it in state like any generated credential. To rotate, create a new credential and delete this one. | [readonly] 

## Methods

### NewGenericCollectionAgentTokenCreatedOut

`func NewGenericCollectionAgentTokenCreatedOut(id string, deploymentId string, type_ CredentialType, description string, createdTime time.Time, mcdId string, mcdToken string, ) *GenericCollectionAgentTokenCreatedOut`

NewGenericCollectionAgentTokenCreatedOut instantiates a new GenericCollectionAgentTokenCreatedOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericCollectionAgentTokenCreatedOutWithDefaults

`func NewGenericCollectionAgentTokenCreatedOutWithDefaults() *GenericCollectionAgentTokenCreatedOut`

NewGenericCollectionAgentTokenCreatedOutWithDefaults instantiates a new GenericCollectionAgentTokenCreatedOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GenericCollectionAgentTokenCreatedOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenericCollectionAgentTokenCreatedOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenericCollectionAgentTokenCreatedOut) SetId(v string)`

SetId sets Id field to given value.


### GetDeploymentId

`func (o *GenericCollectionAgentTokenCreatedOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GenericCollectionAgentTokenCreatedOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GenericCollectionAgentTokenCreatedOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetType

`func (o *GenericCollectionAgentTokenCreatedOut) GetType() CredentialType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericCollectionAgentTokenCreatedOut) GetTypeOk() (*CredentialType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericCollectionAgentTokenCreatedOut) SetType(v CredentialType)`

SetType sets Type field to given value.


### GetDescription

`func (o *GenericCollectionAgentTokenCreatedOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenericCollectionAgentTokenCreatedOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenericCollectionAgentTokenCreatedOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedTime

`func (o *GenericCollectionAgentTokenCreatedOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GenericCollectionAgentTokenCreatedOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GenericCollectionAgentTokenCreatedOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetMcdId

`func (o *GenericCollectionAgentTokenCreatedOut) GetMcdId() string`

GetMcdId returns the McdId field if non-nil, zero value otherwise.

### GetMcdIdOk

`func (o *GenericCollectionAgentTokenCreatedOut) GetMcdIdOk() (*string, bool)`

GetMcdIdOk returns a tuple with the McdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcdId

`func (o *GenericCollectionAgentTokenCreatedOut) SetMcdId(v string)`

SetMcdId sets McdId field to given value.


### GetMcdToken

`func (o *GenericCollectionAgentTokenCreatedOut) GetMcdToken() string`

GetMcdToken returns the McdToken field if non-nil, zero value otherwise.

### GetMcdTokenOk

`func (o *GenericCollectionAgentTokenCreatedOut) GetMcdTokenOk() (*string, bool)`

GetMcdTokenOk returns a tuple with the McdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcdToken

`func (o *GenericCollectionAgentTokenCreatedOut) SetMcdToken(v string)`

SetMcdToken sets McdToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


