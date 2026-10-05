# MulesoftCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppClientId** | **string** | Client ID of the Anypoint connected app. The app needs the View Environment, Read Deployments and Exchange Viewer scopes. | 
**AppClientSecret** | **string** | Secret of the Anypoint connected app. Stored by Monte Carlo and never returned. | 
**Region** | Pointer to [**NullableMulesoftRegion**](MulesoftRegion.md) | Anypoint Platform instance the organization lives on. US when left out. | [optional] 
**OrgId** | Pointer to **NullableString** | Anypoint organization or business group to collect from. Set it when your Mule applications are deployed in a business group. Leave it out to collect the organization that owns the connected app. | [optional] 

## Methods

### NewMulesoftCredentialsIn

`func NewMulesoftCredentialsIn(appClientId string, appClientSecret string, ) *MulesoftCredentialsIn`

NewMulesoftCredentialsIn instantiates a new MulesoftCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulesoftCredentialsInWithDefaults

`func NewMulesoftCredentialsInWithDefaults() *MulesoftCredentialsIn`

NewMulesoftCredentialsInWithDefaults instantiates a new MulesoftCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppClientId

`func (o *MulesoftCredentialsIn) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *MulesoftCredentialsIn) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *MulesoftCredentialsIn) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.


### GetAppClientSecret

`func (o *MulesoftCredentialsIn) GetAppClientSecret() string`

GetAppClientSecret returns the AppClientSecret field if non-nil, zero value otherwise.

### GetAppClientSecretOk

`func (o *MulesoftCredentialsIn) GetAppClientSecretOk() (*string, bool)`

GetAppClientSecretOk returns a tuple with the AppClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientSecret

`func (o *MulesoftCredentialsIn) SetAppClientSecret(v string)`

SetAppClientSecret sets AppClientSecret field to given value.


### GetRegion

`func (o *MulesoftCredentialsIn) GetRegion() MulesoftRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MulesoftCredentialsIn) GetRegionOk() (*MulesoftRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MulesoftCredentialsIn) SetRegion(v MulesoftRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MulesoftCredentialsIn) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *MulesoftCredentialsIn) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *MulesoftCredentialsIn) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetOrgId

`func (o *MulesoftCredentialsIn) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MulesoftCredentialsIn) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MulesoftCredentialsIn) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *MulesoftCredentialsIn) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *MulesoftCredentialsIn) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *MulesoftCredentialsIn) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


