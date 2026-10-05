# MulesoftCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**AppClientId** | **string** | Client ID of the Anypoint connected app. The app needs the View Environment, Read Deployments and Exchange Viewer scopes. | 
**Region** | [**MulesoftRegion**](MulesoftRegion.md) | Anypoint Platform instance the organization lives on. | 
**OrgId** | **NullableString** | Anypoint organization or business group to collect from. Set it when your Mule applications are deployed in a business group. Leave it out to collect the organization that owns the connected app. Null unless set. | 

## Methods

### NewMulesoftCredentialsOut

`func NewMulesoftCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, appClientId string, region MulesoftRegion, orgId NullableString, ) *MulesoftCredentialsOut`

NewMulesoftCredentialsOut instantiates a new MulesoftCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulesoftCredentialsOutWithDefaults

`func NewMulesoftCredentialsOutWithDefaults() *MulesoftCredentialsOut`

NewMulesoftCredentialsOutWithDefaults instantiates a new MulesoftCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MulesoftCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MulesoftCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MulesoftCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *MulesoftCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *MulesoftCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *MulesoftCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *MulesoftCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *MulesoftCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *MulesoftCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *MulesoftCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *MulesoftCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *MulesoftCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetAppClientId

`func (o *MulesoftCredentialsOut) GetAppClientId() string`

GetAppClientId returns the AppClientId field if non-nil, zero value otherwise.

### GetAppClientIdOk

`func (o *MulesoftCredentialsOut) GetAppClientIdOk() (*string, bool)`

GetAppClientIdOk returns a tuple with the AppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppClientId

`func (o *MulesoftCredentialsOut) SetAppClientId(v string)`

SetAppClientId sets AppClientId field to given value.


### GetRegion

`func (o *MulesoftCredentialsOut) GetRegion() MulesoftRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MulesoftCredentialsOut) GetRegionOk() (*MulesoftRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MulesoftCredentialsOut) SetRegion(v MulesoftRegion)`

SetRegion sets Region field to given value.


### GetOrgId

`func (o *MulesoftCredentialsOut) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MulesoftCredentialsOut) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MulesoftCredentialsOut) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### SetOrgIdNil

`func (o *MulesoftCredentialsOut) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *MulesoftCredentialsOut) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


