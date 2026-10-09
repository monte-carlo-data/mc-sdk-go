# CustomConnectorTypeOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the custom connector type, used as a connection type. | 
**Name** | **string** | Display name of the connector. Supplied by the collection agent that registered it, and not validated. | 
**AssetClass** | [**CustomConnectorAssetClass**](CustomConnectorAssetClass.md) | What the connector connects to: a warehouse, an ETL tool or a BI tool. | 
**CollectionAgentId** | **string** | The collection agent that registered the connector. It may name an agent the collection agents endpoints no longer list. | 
**DeploymentId** | **string** | The deployment of the collection agent that registered the connector. Its connections run there. The id may name a deployment the deployments endpoints do not list. | 
**JobTypes** | **[]string** | The collection jobs the connector runs. | 
**IconUrl** | **NullableString** | Icon for the connector. Supplied by the collection agent, not validated, and never fetched by Monte Carlo. Null for a warehouse connector, or when none was supplied. | 
**Terminology** | [**NullableCustomConnectorTerminology**](CustomConnectorTerminology.md) | The connector&#39;s nouns for its tiers. Set only for an ETL connector. | 
**Capabilities** | [**NullableCustomConnectorCapabilities**](CustomConnectorCapabilities.md) | What the connector supports. Set only for a warehouse connector. | 
**CreatedTime** | **time.Time** | When the connector type was first registered. | 
**UpdatedTime** | **time.Time** | When the connector type was last registered. | 

## Methods

### NewCustomConnectorTypeOut

`func NewCustomConnectorTypeOut(id string, name string, assetClass CustomConnectorAssetClass, collectionAgentId string, deploymentId string, jobTypes []string, iconUrl NullableString, terminology NullableCustomConnectorTerminology, capabilities NullableCustomConnectorCapabilities, createdTime time.Time, updatedTime time.Time, ) *CustomConnectorTypeOut`

NewCustomConnectorTypeOut instantiates a new CustomConnectorTypeOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomConnectorTypeOutWithDefaults

`func NewCustomConnectorTypeOutWithDefaults() *CustomConnectorTypeOut`

NewCustomConnectorTypeOutWithDefaults instantiates a new CustomConnectorTypeOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomConnectorTypeOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomConnectorTypeOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomConnectorTypeOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomConnectorTypeOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomConnectorTypeOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomConnectorTypeOut) SetName(v string)`

SetName sets Name field to given value.


### GetAssetClass

`func (o *CustomConnectorTypeOut) GetAssetClass() CustomConnectorAssetClass`

GetAssetClass returns the AssetClass field if non-nil, zero value otherwise.

### GetAssetClassOk

`func (o *CustomConnectorTypeOut) GetAssetClassOk() (*CustomConnectorAssetClass, bool)`

GetAssetClassOk returns a tuple with the AssetClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetClass

`func (o *CustomConnectorTypeOut) SetAssetClass(v CustomConnectorAssetClass)`

SetAssetClass sets AssetClass field to given value.


### GetCollectionAgentId

`func (o *CustomConnectorTypeOut) GetCollectionAgentId() string`

GetCollectionAgentId returns the CollectionAgentId field if non-nil, zero value otherwise.

### GetCollectionAgentIdOk

`func (o *CustomConnectorTypeOut) GetCollectionAgentIdOk() (*string, bool)`

GetCollectionAgentIdOk returns a tuple with the CollectionAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAgentId

`func (o *CustomConnectorTypeOut) SetCollectionAgentId(v string)`

SetCollectionAgentId sets CollectionAgentId field to given value.


### GetDeploymentId

`func (o *CustomConnectorTypeOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *CustomConnectorTypeOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *CustomConnectorTypeOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### GetJobTypes

`func (o *CustomConnectorTypeOut) GetJobTypes() []string`

GetJobTypes returns the JobTypes field if non-nil, zero value otherwise.

### GetJobTypesOk

`func (o *CustomConnectorTypeOut) GetJobTypesOk() (*[]string, bool)`

GetJobTypesOk returns a tuple with the JobTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypes

`func (o *CustomConnectorTypeOut) SetJobTypes(v []string)`

SetJobTypes sets JobTypes field to given value.


### GetIconUrl

`func (o *CustomConnectorTypeOut) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *CustomConnectorTypeOut) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *CustomConnectorTypeOut) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### SetIconUrlNil

`func (o *CustomConnectorTypeOut) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *CustomConnectorTypeOut) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetTerminology

`func (o *CustomConnectorTypeOut) GetTerminology() CustomConnectorTerminology`

GetTerminology returns the Terminology field if non-nil, zero value otherwise.

### GetTerminologyOk

`func (o *CustomConnectorTypeOut) GetTerminologyOk() (*CustomConnectorTerminology, bool)`

GetTerminologyOk returns a tuple with the Terminology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminology

`func (o *CustomConnectorTypeOut) SetTerminology(v CustomConnectorTerminology)`

SetTerminology sets Terminology field to given value.


### SetTerminologyNil

`func (o *CustomConnectorTypeOut) SetTerminologyNil(b bool)`

 SetTerminologyNil sets the value for Terminology to be an explicit nil

### UnsetTerminology
`func (o *CustomConnectorTypeOut) UnsetTerminology()`

UnsetTerminology ensures that no value is present for Terminology, not even an explicit nil
### GetCapabilities

`func (o *CustomConnectorTypeOut) GetCapabilities() CustomConnectorCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CustomConnectorTypeOut) GetCapabilitiesOk() (*CustomConnectorCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CustomConnectorTypeOut) SetCapabilities(v CustomConnectorCapabilities)`

SetCapabilities sets Capabilities field to given value.


### SetCapabilitiesNil

`func (o *CustomConnectorTypeOut) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *CustomConnectorTypeOut) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetCreatedTime

`func (o *CustomConnectorTypeOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CustomConnectorTypeOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CustomConnectorTypeOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetUpdatedTime

`func (o *CustomConnectorTypeOut) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *CustomConnectorTypeOut) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *CustomConnectorTypeOut) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


