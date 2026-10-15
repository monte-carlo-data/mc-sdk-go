# MulesoftCredentialsChangeValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**AppClientId** | Pointer to **NullableString** | Client ID of the Anypoint connected app. The app needs the View Environment, Read Deployments and Exchange Viewer scopes. | [optional] 
**AppClientSecret** | Pointer to **NullableString** | Secret of the Anypoint connected app. Stored by Monte Carlo and never returned. | [optional] 
**Region** | Pointer to [**NullableMulesoftRegion**](MulesoftRegion.md) | Anypoint Platform instance the organization lives on. | [optional] 
**OrgId** | Pointer to **NullableString** | Anypoint organization or business group to collect from. Set it when your Mule applications are deployed in a business group. Leave it out to collect the organization that owns the connected app. | [optional] 

## Methods

### NewMulesoftCredentialsChangeValidateIn

`func NewMulesoftCredentialsChangeValidateIn(deploymentId string, ) *MulesoftCredentialsChangeValidateIn`

NewMulesoftCredentialsChangeValidateIn instantiates a new MulesoftCredentialsChangeValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulesoftCredentialsChangeValidateInWithDefaults

`func NewMulesoftCredentialsChangeValidateInWithDefaults() *MulesoftCredentialsChangeValidateIn`

NewMulesoftCredentialsChangeValidateInWithDefaults instantiates a new MulesoftCredentialsChangeValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *MulesoftCredentialsChangeValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *MulesoftCredentialsChangeValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *MulesoftCredentialsChangeValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetAppClientId

`func (o *MulesoftCredentialsChangeValidateIn) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *MulesoftCredentialsChangeValidateIn) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *MulesoftCredentialsChangeValidateIn) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.

### HasAppClientId

`func (o *MulesoftCredentialsChangeValidateIn) HasAppClientId() bool`

HasAppClientId returns a boolean if a field has been set.

### SetAppClientIdNil

`func (o *MulesoftCredentialsChangeValidateIn) SetAppClientIdNil(b bool)`

 SetAppClientIdNil sets the value for AppClientId to be an explicit nil

### UnsetAppClientId
`func (o *MulesoftCredentialsChangeValidateIn) UnsetAppClientId()`

UnsetAppClientId ensures that no value is present for AppClientId, not even an explicit nil
### GetAppClientSecret

`func (o *MulesoftCredentialsChangeValidateIn) GetAppClientSecret() string`

GetAppClientSecret returns the AppClientSecret field if non-nil, zero value otherwise.

### GetAppClientSecretOk

`func (o *MulesoftCredentialsChangeValidateIn) GetAppClientSecretOk() (*string, bool)`

GetAppClientSecretOk returns a tuple with the AppClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientSecret

`func (o *MulesoftCredentialsChangeValidateIn) SetAppClientSecret(v string)`

SetAppClientSecret sets AppClientSecret field to given value.

### HasAppClientSecret

`func (o *MulesoftCredentialsChangeValidateIn) HasAppClientSecret() bool`

HasAppClientSecret returns a boolean if a field has been set.

### SetAppClientSecretNil

`func (o *MulesoftCredentialsChangeValidateIn) SetAppClientSecretNil(b bool)`

 SetAppClientSecretNil sets the value for AppClientSecret to be an explicit nil

### UnsetAppClientSecret
`func (o *MulesoftCredentialsChangeValidateIn) UnsetAppClientSecret()`

UnsetAppClientSecret ensures that no value is present for AppClientSecret, not even an explicit nil
### GetRegion

`func (o *MulesoftCredentialsChangeValidateIn) GetRegion() MulesoftRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MulesoftCredentialsChangeValidateIn) GetRegionOk() (*MulesoftRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MulesoftCredentialsChangeValidateIn) SetRegion(v MulesoftRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MulesoftCredentialsChangeValidateIn) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *MulesoftCredentialsChangeValidateIn) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *MulesoftCredentialsChangeValidateIn) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetOrgId

`func (o *MulesoftCredentialsChangeValidateIn) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MulesoftCredentialsChangeValidateIn) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MulesoftCredentialsChangeValidateIn) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *MulesoftCredentialsChangeValidateIn) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *MulesoftCredentialsChangeValidateIn) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *MulesoftCredentialsChangeValidateIn) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


