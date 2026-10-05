# EtlContainerOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the ETL container. | 
**Type** | [**EtlContainerType**](EtlContainerType.md) | The ETL tool the container represents. Fixed once created. | 
**Name** | **string** | Display name of the ETL container. | 
**IsSynthetic** | **bool** | True for a container this API does not create or delete. Most belong to another connection, such as a warehouse or BI connection, and go away with it. | 
**DeploymentId** | **NullableString** | The deployment the container&#39;s connection runs through. Null for a type that runs on none, such as &#x60;airflow&#x60;. The id may name a deployment on Monte Carlo&#39;s older collection platform. The deployments endpoints do not list those. | 
**DeploymentName** | **NullableString** | Display name of the deployment. Null exactly when &#x60;deployment_id&#x60; is. | 
**CreatedTime** | **time.Time** | When the ETL container was created. | 

## Methods

### NewEtlContainerOut

`func NewEtlContainerOut(id string, type_ EtlContainerType, name string, isSynthetic bool, deploymentId NullableString, deploymentName NullableString, createdTime time.Time, ) *EtlContainerOut`

NewEtlContainerOut instantiates a new EtlContainerOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtlContainerOutWithDefaults

`func NewEtlContainerOutWithDefaults() *EtlContainerOut`

NewEtlContainerOutWithDefaults instantiates a new EtlContainerOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EtlContainerOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EtlContainerOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EtlContainerOut) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *EtlContainerOut) GetType() EtlContainerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EtlContainerOut) GetTypeOk() (*EtlContainerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EtlContainerOut) SetType(v EtlContainerType)`

SetType sets Type field to given value.


### GetName

`func (o *EtlContainerOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EtlContainerOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EtlContainerOut) SetName(v string)`

SetName sets Name field to given value.


### GetIsSynthetic

`func (o *EtlContainerOut) GetIsSynthetic() bool`

GetIsSynthetic returns the IsSynthetic field if non-nil, zero value otherwise.

### GetIsSyntheticOk

`func (o *EtlContainerOut) GetIsSyntheticOk() (*bool, bool)`

GetIsSyntheticOk returns a tuple with the IsSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynthetic

`func (o *EtlContainerOut) SetIsSynthetic(v bool)`

SetIsSynthetic sets IsSynthetic field to given value.


### GetDeploymentId

`func (o *EtlContainerOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *EtlContainerOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *EtlContainerOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### SetDeploymentIdNil

`func (o *EtlContainerOut) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *EtlContainerOut) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetDeploymentName

`func (o *EtlContainerOut) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *EtlContainerOut) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *EtlContainerOut) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.


### SetDeploymentNameNil

`func (o *EtlContainerOut) SetDeploymentNameNil(b bool)`

 SetDeploymentNameNil sets the value for DeploymentName to be an explicit nil

### UnsetDeploymentName
`func (o *EtlContainerOut) UnsetDeploymentName()`

UnsetDeploymentName ensures that no value is present for DeploymentName, not even an explicit nil
### GetCreatedTime

`func (o *EtlContainerOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *EtlContainerOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *EtlContainerOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


