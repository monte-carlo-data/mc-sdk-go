# BiContainerOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the BI container. | 
**Type** | [**BiContainerType**](BiContainerType.md) | The BI tool the container represents. Fixed once created. | 
**Name** | **NullableString** | Display name of the BI container. Null for a container that was never named. | 
**DeploymentId** | **NullableString** | The deployment the container&#39;s connections run through. Null for a container with no deployment, such as a push-only custom BI connector&#39;s. The id may name a deployment on Monte Carlo&#39;s older collection platform. The deployments endpoints do not list those. | 
**DeploymentName** | **NullableString** | Display name of the deployment. Null exactly when &#x60;deployment_id&#x60; is. | 
**CreatedTime** | **time.Time** | When the BI container was created. | 

## Methods

### NewBiContainerOut

`func NewBiContainerOut(id string, type_ BiContainerType, name NullableString, deploymentId NullableString, deploymentName NullableString, createdTime time.Time, ) *BiContainerOut`

NewBiContainerOut instantiates a new BiContainerOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiContainerOutWithDefaults

`func NewBiContainerOutWithDefaults() *BiContainerOut`

NewBiContainerOutWithDefaults instantiates a new BiContainerOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BiContainerOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BiContainerOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BiContainerOut) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BiContainerOut) GetType() BiContainerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BiContainerOut) GetTypeOk() (*BiContainerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BiContainerOut) SetType(v BiContainerType)`

SetType sets Type field to given value.


### GetName

`func (o *BiContainerOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BiContainerOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BiContainerOut) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *BiContainerOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BiContainerOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeploymentId

`func (o *BiContainerOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *BiContainerOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *BiContainerOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### SetDeploymentIdNil

`func (o *BiContainerOut) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *BiContainerOut) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetDeploymentName

`func (o *BiContainerOut) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *BiContainerOut) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *BiContainerOut) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.


### SetDeploymentNameNil

`func (o *BiContainerOut) SetDeploymentNameNil(b bool)`

 SetDeploymentNameNil sets the value for DeploymentName to be an explicit nil

### UnsetDeploymentName
`func (o *BiContainerOut) UnsetDeploymentName()`

UnsetDeploymentName ensures that no value is present for DeploymentName, not even an explicit nil
### GetCreatedTime

`func (o *BiContainerOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *BiContainerOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *BiContainerOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


