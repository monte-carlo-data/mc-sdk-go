# CollectionAgentCredentialOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **NullableString** | Client id the agent presents. Set for an &#x60;OAUTH_CLIENT&#x60;, null otherwise. | [optional] 
**CreatedTime** | **time.Time** | When the credential was created. | 
**DeploymentId** | **string** | Identifier of the deployment whose agent presents this credential. | 
**Description** | **string** | What this credential is for. | 
**ExpirationTime** | Pointer to **NullableTime** | When an &#x60;OAUTH_CLIENT&#x60; stops being accepted. Null for one that does not expire, and for a &#x60;TOKEN&#x60;. | [optional] 
**Id** | **string** | Unique identifier of the credential. For a token this is also its key id; for an OAuth client, its client id. | 
**McdId** | Pointer to **NullableString** | Key id the agent presents. Set for a &#x60;TOKEN&#x60;, null otherwise. | [optional] 
**Scopes** | Pointer to **[]string** | OAuth scopes the client is granted. Set for an &#x60;OAUTH_CLIENT&#x60;, null otherwise. | [optional] 
**Type** | [**CredentialType**](CredentialType.md) | Which kind of credential this is. | 

## Methods

### NewCollectionAgentCredentialOut

`func NewCollectionAgentCredentialOut(createdTime time.Time, deploymentId string, description string, id string, type_ CredentialType, ) *CollectionAgentCredentialOut`

NewCollectionAgentCredentialOut instantiates a new CollectionAgentCredentialOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionAgentCredentialOutWithDefaults

`func NewCollectionAgentCredentialOutWithDefaults() *CollectionAgentCredentialOut`

NewCollectionAgentCredentialOutWithDefaults instantiates a new CollectionAgentCredentialOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CollectionAgentCredentialOut) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CollectionAgentCredentialOut) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CollectionAgentCredentialOut) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CollectionAgentCredentialOut) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *CollectionAgentCredentialOut) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *CollectionAgentCredentialOut) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetCreatedTime

`func (o *CollectionAgentCredentialOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CollectionAgentCredentialOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CollectionAgentCredentialOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetDeploymentId

`func (o *CollectionAgentCredentialOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *CollectionAgentCredentialOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *CollectionAgentCredentialOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetDescription

`func (o *CollectionAgentCredentialOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionAgentCredentialOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionAgentCredentialOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExpirationTime

`func (o *CollectionAgentCredentialOut) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *CollectionAgentCredentialOut) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *CollectionAgentCredentialOut) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *CollectionAgentCredentialOut) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### SetExpirationTimeNil

`func (o *CollectionAgentCredentialOut) SetExpirationTimeNil(b bool)`

 SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil

### UnsetExpirationTime
`func (o *CollectionAgentCredentialOut) UnsetExpirationTime()`

UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
### GetId

`func (o *CollectionAgentCredentialOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionAgentCredentialOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionAgentCredentialOut) SetId(v string)`

SetId sets Id field to given value.


### GetMcdId

`func (o *CollectionAgentCredentialOut) GetMcdId() string`

GetMcdId returns the McdId field if non-nil, zero value otherwise.

### GetMcdIdOk

`func (o *CollectionAgentCredentialOut) GetMcdIdOk() (*string, bool)`

GetMcdIdOk returns a tuple with the McdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcdId

`func (o *CollectionAgentCredentialOut) SetMcdId(v string)`

SetMcdId sets McdId field to given value.

### HasMcdId

`func (o *CollectionAgentCredentialOut) HasMcdId() bool`

HasMcdId returns a boolean if a field has been set.

### SetMcdIdNil

`func (o *CollectionAgentCredentialOut) SetMcdIdNil(b bool)`

 SetMcdIdNil sets the value for McdId to be an explicit nil

### UnsetMcdId
`func (o *CollectionAgentCredentialOut) UnsetMcdId()`

UnsetMcdId ensures that no value is present for McdId, not even an explicit nil
### GetScopes

`func (o *CollectionAgentCredentialOut) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CollectionAgentCredentialOut) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CollectionAgentCredentialOut) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CollectionAgentCredentialOut) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *CollectionAgentCredentialOut) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *CollectionAgentCredentialOut) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetType

`func (o *CollectionAgentCredentialOut) GetType() CredentialType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CollectionAgentCredentialOut) GetTypeOk() (*CredentialType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CollectionAgentCredentialOut) SetType(v CredentialType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


