# GenericCollectionAgentOAuthClientCreatedOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credential. For a token this is also its key id; for an OAuth client, its client id. | 
**DeploymentId** | **string** | Identifier of the deployment whose agent presents this credential. | 
**Type** | [**CredentialType**](CredentialType.md) | Which kind of credential this is. | 
**Description** | **string** | What this credential is for. | 
**CreatedTime** | **time.Time** | When the credential was created. | 
**ClientId** | **string** | Client id the agent presents, as &#x60;client_id&#x60; in its configuration. The same value as &#x60;id&#x60;. | 
**Scopes** | **[]string** | OAuth scopes the client is granted. | 
**ExpirationTime** | Pointer to **NullableTime** | When the client stops being accepted. Null for a client that does not expire. | [optional] 
**ClientSecret** | **string** | Secret the agent presents, as &#x60;client_secret&#x60; in its configuration. Returned once, by this call. It is not retrievable afterwards, and Terraform holds it in state like any generated credential. To rotate, create a new credential and delete this one. | [readonly] 
**SecretId** | **string** | Identifier of the secret this call created, as the client&#39;s secrets are listed. Keep it with the secret. | 

## Methods

### NewGenericCollectionAgentOAuthClientCreatedOut

`func NewGenericCollectionAgentOAuthClientCreatedOut(id string, deploymentId string, type_ CredentialType, description string, createdTime time.Time, clientId string, scopes []string, clientSecret string, secretId string, ) *GenericCollectionAgentOAuthClientCreatedOut`

NewGenericCollectionAgentOAuthClientCreatedOut instantiates a new GenericCollectionAgentOAuthClientCreatedOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericCollectionAgentOAuthClientCreatedOutWithDefaults

`func NewGenericCollectionAgentOAuthClientCreatedOutWithDefaults() *GenericCollectionAgentOAuthClientCreatedOut`

NewGenericCollectionAgentOAuthClientCreatedOutWithDefaults instantiates a new GenericCollectionAgentOAuthClientCreatedOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenericCollectionAgentOAuthClientCreatedOut) SetId(v string)`

SetId sets Id field to given value.


### GetDeploymentId

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GenericCollectionAgentOAuthClientCreatedOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetType

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetType() CredentialType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetTypeOk() (*CredentialType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericCollectionAgentOAuthClientCreatedOut) SetType(v CredentialType)`

SetType sets Type field to given value.


### GetDescription

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenericCollectionAgentOAuthClientCreatedOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedTime

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GenericCollectionAgentOAuthClientCreatedOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetClientId

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GenericCollectionAgentOAuthClientCreatedOut) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetScopes

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GenericCollectionAgentOAuthClientCreatedOut) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetExpirationTime

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *GenericCollectionAgentOAuthClientCreatedOut) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *GenericCollectionAgentOAuthClientCreatedOut) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### SetExpirationTimeNil

`func (o *GenericCollectionAgentOAuthClientCreatedOut) SetExpirationTimeNil(b bool)`

 SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil

### UnsetExpirationTime
`func (o *GenericCollectionAgentOAuthClientCreatedOut) UnsetExpirationTime()`

UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
### GetClientSecret

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *GenericCollectionAgentOAuthClientCreatedOut) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetSecretId

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *GenericCollectionAgentOAuthClientCreatedOut) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *GenericCollectionAgentOAuthClientCreatedOut) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


