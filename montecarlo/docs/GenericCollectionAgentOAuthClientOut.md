# GenericCollectionAgentOAuthClientOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client id the agent presents, as &#x60;client_id&#x60; in its configuration. The same value as &#x60;id&#x60;. | 
**CreatedTime** | **time.Time** | When the credential was created. | 
**DeploymentId** | **string** | Identifier of the deployment whose agent presents this credential. | 
**Description** | **string** | What this credential is for. | 
**ExpirationTime** | Pointer to **NullableTime** | When the client stops being accepted. Null for a client that does not expire. | [optional] 
**Id** | **string** | Unique identifier of the credential. For a token this is also its key id; for an OAuth client, its client id. | 
**Scopes** | **[]string** | OAuth scopes the client is granted. | 
**Type** | [**CredentialType**](CredentialType.md) | Which kind of credential this is. | 

## Methods

### NewGenericCollectionAgentOAuthClientOut

`func NewGenericCollectionAgentOAuthClientOut(clientId string, createdTime time.Time, deploymentId string, description string, id string, scopes []string, type_ CredentialType, ) *GenericCollectionAgentOAuthClientOut`

NewGenericCollectionAgentOAuthClientOut instantiates a new GenericCollectionAgentOAuthClientOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericCollectionAgentOAuthClientOutWithDefaults

`func NewGenericCollectionAgentOAuthClientOutWithDefaults() *GenericCollectionAgentOAuthClientOut`

NewGenericCollectionAgentOAuthClientOutWithDefaults instantiates a new GenericCollectionAgentOAuthClientOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *GenericCollectionAgentOAuthClientOut) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GenericCollectionAgentOAuthClientOut) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GenericCollectionAgentOAuthClientOut) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCreatedTime

`func (o *GenericCollectionAgentOAuthClientOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GenericCollectionAgentOAuthClientOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GenericCollectionAgentOAuthClientOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetDeploymentId

`func (o *GenericCollectionAgentOAuthClientOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GenericCollectionAgentOAuthClientOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GenericCollectionAgentOAuthClientOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetDescription

`func (o *GenericCollectionAgentOAuthClientOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenericCollectionAgentOAuthClientOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenericCollectionAgentOAuthClientOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExpirationTime

`func (o *GenericCollectionAgentOAuthClientOut) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *GenericCollectionAgentOAuthClientOut) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *GenericCollectionAgentOAuthClientOut) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *GenericCollectionAgentOAuthClientOut) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### SetExpirationTimeNil

`func (o *GenericCollectionAgentOAuthClientOut) SetExpirationTimeNil(b bool)`

 SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil

### UnsetExpirationTime
`func (o *GenericCollectionAgentOAuthClientOut) UnsetExpirationTime()`

UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
### GetId

`func (o *GenericCollectionAgentOAuthClientOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenericCollectionAgentOAuthClientOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenericCollectionAgentOAuthClientOut) SetId(v string)`

SetId sets Id field to given value.


### GetScopes

`func (o *GenericCollectionAgentOAuthClientOut) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GenericCollectionAgentOAuthClientOut) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GenericCollectionAgentOAuthClientOut) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetType

`func (o *GenericCollectionAgentOAuthClientOut) GetType() CredentialType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericCollectionAgentOAuthClientOut) GetTypeOk() (*CredentialType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericCollectionAgentOAuthClientOut) SetType(v CredentialType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


