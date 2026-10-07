# ConnectionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the connection. | 
**ConnectionType** | **string** | What the connection reaches, such as &#x60;snowflake&#x60;. Taken from the credentials the connection was created with, and fixed once created. | 
**Name** | **NullableString** | Display name of the connection. Null for a connection that was never named. | 
**WarehouseId** | **NullableString** | The warehouse the connection belongs to. Null for a connection on a BI or ETL container. Fixed once created. | 
**WarehouseName** | **NullableString** | Display name of that warehouse. Null for a warehouse that was never named, and for a connection on a BI or ETL container. | 
**BiContainerId** | **NullableString** | The BI container the connection belongs to. Null for a connection on a warehouse or an ETL container. Fixed once created. | 
**BiContainerName** | **NullableString** | Display name of that BI container. Null for a container that was never named, and for a connection on a warehouse or an ETL container. | 
**EtlContainerId** | **NullableString** | The ETL container the connection belongs to. Fixed once created for a connection on an ETL container. A warehouse or BI connection has one while it collects ETL jobs through a container of its own, such as Snowflake Tasks, a Databricks metastore or Power BI dataflows, and it changes when that is turned on or off. Null otherwise. | 
**EtlContainerName** | **NullableString** | Display name of that ETL container. Null when there is no ETL container. | 
**DeploymentId** | **NullableString** | The deployment the connection runs through, taken from its warehouse, BI container or ETL container. Null when that has no deployment, as an Airflow ETL container has none. The id may name a deployment on Monte Carlo&#39;s older collection platform, which the deployments endpoints do not list. | 
**DeploymentName** | **NullableString** | Display name of that deployment. Null when there is no deployment to name. | 
**CredentialsId** | **NullableString** | The credentials the connection reads with. Null for a connection created before credentials became their own resource, and for one created outside this API. | 
**CredentialsStorageType** | [**NullableCredentialsStorageType**](CredentialsStorageType.md) | Where that secret lives. Null when there are no credentials to describe. | 
**JobTypes** | **[]string** | The jobs Monte Carlo runs on this connection, such as &#x60;metadata&#x60;. | 
**CreatedTime** | **time.Time** | When the connection was created. | 

## Methods

### NewConnectionOut

`func NewConnectionOut(id string, connectionType string, name NullableString, warehouseId NullableString, warehouseName NullableString, biContainerId NullableString, biContainerName NullableString, etlContainerId NullableString, etlContainerName NullableString, deploymentId NullableString, deploymentName NullableString, credentialsId NullableString, credentialsStorageType NullableCredentialsStorageType, jobTypes []string, createdTime time.Time, ) *ConnectionOut`

NewConnectionOut instantiates a new ConnectionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOutWithDefaults

`func NewConnectionOutWithDefaults() *ConnectionOut`

NewConnectionOutWithDefaults instantiates a new ConnectionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *ConnectionOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ConnectionOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ConnectionOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetName

`func (o *ConnectionOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionOut) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ConnectionOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConnectionOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetWarehouseId

`func (o *ConnectionOut) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ConnectionOut) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ConnectionOut) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.


### SetWarehouseIdNil

`func (o *ConnectionOut) SetWarehouseIdNil(b bool)`

 SetWarehouseIdNil sets the value for WarehouseId to be an explicit nil

### UnsetWarehouseId
`func (o *ConnectionOut) UnsetWarehouseId()`

UnsetWarehouseId ensures that no value is present for WarehouseId, not even an explicit nil
### GetWarehouseName

`func (o *ConnectionOut) GetWarehouseName() string`

GetWarehouseName returns the WarehouseName field if non-nil, zero value otherwise.

### GetWarehouseNameOk

`func (o *ConnectionOut) GetWarehouseNameOk() (*string, bool)`

GetWarehouseNameOk returns a tuple with the WarehouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseName

`func (o *ConnectionOut) SetWarehouseName(v string)`

SetWarehouseName sets WarehouseName field to given value.


### SetWarehouseNameNil

`func (o *ConnectionOut) SetWarehouseNameNil(b bool)`

 SetWarehouseNameNil sets the value for WarehouseName to be an explicit nil

### UnsetWarehouseName
`func (o *ConnectionOut) UnsetWarehouseName()`

UnsetWarehouseName ensures that no value is present for WarehouseName, not even an explicit nil
### GetBiContainerId

`func (o *ConnectionOut) GetBiContainerId() string`

GetBiContainerId returns the BiContainerId field if non-nil, zero value otherwise.

### GetBiContainerIdOk

`func (o *ConnectionOut) GetBiContainerIdOk() (*string, bool)`

GetBiContainerIdOk returns a tuple with the BiContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiContainerId

`func (o *ConnectionOut) SetBiContainerId(v string)`

SetBiContainerId sets BiContainerId field to given value.


### SetBiContainerIdNil

`func (o *ConnectionOut) SetBiContainerIdNil(b bool)`

 SetBiContainerIdNil sets the value for BiContainerId to be an explicit nil

### UnsetBiContainerId
`func (o *ConnectionOut) UnsetBiContainerId()`

UnsetBiContainerId ensures that no value is present for BiContainerId, not even an explicit nil
### GetBiContainerName

`func (o *ConnectionOut) GetBiContainerName() string`

GetBiContainerName returns the BiContainerName field if non-nil, zero value otherwise.

### GetBiContainerNameOk

`func (o *ConnectionOut) GetBiContainerNameOk() (*string, bool)`

GetBiContainerNameOk returns a tuple with the BiContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiContainerName

`func (o *ConnectionOut) SetBiContainerName(v string)`

SetBiContainerName sets BiContainerName field to given value.


### SetBiContainerNameNil

`func (o *ConnectionOut) SetBiContainerNameNil(b bool)`

 SetBiContainerNameNil sets the value for BiContainerName to be an explicit nil

### UnsetBiContainerName
`func (o *ConnectionOut) UnsetBiContainerName()`

UnsetBiContainerName ensures that no value is present for BiContainerName, not even an explicit nil
### GetEtlContainerId

`func (o *ConnectionOut) GetEtlContainerId() string`

GetEtlContainerId returns the EtlContainerId field if non-nil, zero value otherwise.

### GetEtlContainerIdOk

`func (o *ConnectionOut) GetEtlContainerIdOk() (*string, bool)`

GetEtlContainerIdOk returns a tuple with the EtlContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlContainerId

`func (o *ConnectionOut) SetEtlContainerId(v string)`

SetEtlContainerId sets EtlContainerId field to given value.


### SetEtlContainerIdNil

`func (o *ConnectionOut) SetEtlContainerIdNil(b bool)`

 SetEtlContainerIdNil sets the value for EtlContainerId to be an explicit nil

### UnsetEtlContainerId
`func (o *ConnectionOut) UnsetEtlContainerId()`

UnsetEtlContainerId ensures that no value is present for EtlContainerId, not even an explicit nil
### GetEtlContainerName

`func (o *ConnectionOut) GetEtlContainerName() string`

GetEtlContainerName returns the EtlContainerName field if non-nil, zero value otherwise.

### GetEtlContainerNameOk

`func (o *ConnectionOut) GetEtlContainerNameOk() (*string, bool)`

GetEtlContainerNameOk returns a tuple with the EtlContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlContainerName

`func (o *ConnectionOut) SetEtlContainerName(v string)`

SetEtlContainerName sets EtlContainerName field to given value.


### SetEtlContainerNameNil

`func (o *ConnectionOut) SetEtlContainerNameNil(b bool)`

 SetEtlContainerNameNil sets the value for EtlContainerName to be an explicit nil

### UnsetEtlContainerName
`func (o *ConnectionOut) UnsetEtlContainerName()`

UnsetEtlContainerName ensures that no value is present for EtlContainerName, not even an explicit nil
### GetDeploymentId

`func (o *ConnectionOut) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ConnectionOut) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ConnectionOut) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.


### SetDeploymentIdNil

`func (o *ConnectionOut) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *ConnectionOut) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetDeploymentName

`func (o *ConnectionOut) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *ConnectionOut) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *ConnectionOut) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.


### SetDeploymentNameNil

`func (o *ConnectionOut) SetDeploymentNameNil(b bool)`

 SetDeploymentNameNil sets the value for DeploymentName to be an explicit nil

### UnsetDeploymentName
`func (o *ConnectionOut) UnsetDeploymentName()`

UnsetDeploymentName ensures that no value is present for DeploymentName, not even an explicit nil
### GetCredentialsId

`func (o *ConnectionOut) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *ConnectionOut) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *ConnectionOut) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### SetCredentialsIdNil

`func (o *ConnectionOut) SetCredentialsIdNil(b bool)`

 SetCredentialsIdNil sets the value for CredentialsId to be an explicit nil

### UnsetCredentialsId
`func (o *ConnectionOut) UnsetCredentialsId()`

UnsetCredentialsId ensures that no value is present for CredentialsId, not even an explicit nil
### GetCredentialsStorageType

`func (o *ConnectionOut) GetCredentialsStorageType() CredentialsStorageType`

GetCredentialsStorageType returns the CredentialsStorageType field if non-nil, zero value otherwise.

### GetCredentialsStorageTypeOk

`func (o *ConnectionOut) GetCredentialsStorageTypeOk() (*CredentialsStorageType, bool)`

GetCredentialsStorageTypeOk returns a tuple with the CredentialsStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsStorageType

`func (o *ConnectionOut) SetCredentialsStorageType(v CredentialsStorageType)`

SetCredentialsStorageType sets CredentialsStorageType field to given value.


### SetCredentialsStorageTypeNil

`func (o *ConnectionOut) SetCredentialsStorageTypeNil(b bool)`

 SetCredentialsStorageTypeNil sets the value for CredentialsStorageType to be an explicit nil

### UnsetCredentialsStorageType
`func (o *ConnectionOut) UnsetCredentialsStorageType()`

UnsetCredentialsStorageType ensures that no value is present for CredentialsStorageType, not even an explicit nil
### GetJobTypes

`func (o *ConnectionOut) GetJobTypes() []string`

GetJobTypes returns the JobTypes field if non-nil, zero value otherwise.

### GetJobTypesOk

`func (o *ConnectionOut) GetJobTypesOk() (*[]string, bool)`

GetJobTypesOk returns a tuple with the JobTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypes

`func (o *ConnectionOut) SetJobTypes(v []string)`

SetJobTypes sets JobTypes field to given value.


### GetCreatedTime

`func (o *ConnectionOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ConnectionOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ConnectionOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


