# MulesoftCredentialsValidateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **string** | Deployment that runs the validations. It has to be one &#x60;GET /deployments&#x60; lists, and it has to be able to reach the system the credentials are for. | 
**AppClientId** | **string** | Client ID of the Anypoint connected app. The app needs the View Environment, Read Deployments and Exchange Viewer scopes. | 
**AppClientSecret** | **string** | Secret of the Anypoint connected app. Used for this check and not kept. | 
**Region** | Pointer to [**NullableMulesoftRegion**](MulesoftRegion.md) | Anypoint Platform instance the organization lives on. US when left out. | [optional] 
**OrgId** | Pointer to **NullableString** | Anypoint organization or business group to collect from. Set it when your Mule applications are deployed in a business group. Leave it out to collect the organization that owns the connected app. | [optional] 

## Methods

### NewMulesoftCredentialsValidateIn

`func NewMulesoftCredentialsValidateIn(deploymentId string, appClientId string, appClientSecret string, ) *MulesoftCredentialsValidateIn`

NewMulesoftCredentialsValidateIn instantiates a new MulesoftCredentialsValidateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulesoftCredentialsValidateInWithDefaults

`func NewMulesoftCredentialsValidateInWithDefaults() *MulesoftCredentialsValidateIn`

NewMulesoftCredentialsValidateInWithDefaults instantiates a new MulesoftCredentialsValidateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *MulesoftCredentialsValidateIn) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *MulesoftCredentialsValidateIn) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *MulesoftCredentialsValidateIn) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetAppClientId

`func (o *MulesoftCredentialsValidateIn) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *MulesoftCredentialsValidateIn) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *MulesoftCredentialsValidateIn) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.


### GetAppClientSecret

`func (o *MulesoftCredentialsValidateIn) GetAppClientSecret() string`

GetAppClientSecret returns the AppClientSecret field if non-nil, zero value otherwise.

### GetAppClientSecretOk

`func (o *MulesoftCredentialsValidateIn) GetAppClientSecretOk() (*string, bool)`

GetAppClientSecretOk returns a tuple with the AppClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientSecret

`func (o *MulesoftCredentialsValidateIn) SetAppClientSecret(v string)`

SetAppClientSecret sets AppClientSecret field to given value.


### GetRegion

`func (o *MulesoftCredentialsValidateIn) GetRegion() MulesoftRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MulesoftCredentialsValidateIn) GetRegionOk() (*MulesoftRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MulesoftCredentialsValidateIn) SetRegion(v MulesoftRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MulesoftCredentialsValidateIn) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *MulesoftCredentialsValidateIn) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *MulesoftCredentialsValidateIn) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetOrgId

`func (o *MulesoftCredentialsValidateIn) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MulesoftCredentialsValidateIn) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MulesoftCredentialsValidateIn) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *MulesoftCredentialsValidateIn) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *MulesoftCredentialsValidateIn) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *MulesoftCredentialsValidateIn) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


