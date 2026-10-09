# ConnectionIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for the connection. Unique among the connections of its warehouse or BI container. An ETL container holds one connection. | 
**WarehouseId** | Pointer to **NullableString** | The warehouse to add the connection to. Its type has to match what the credentials are for. Send exactly one of this, &#x60;bi_container_id&#x60; and &#x60;etl_container_id&#x60;. | [optional] 
**BiContainerId** | Pointer to **NullableString** | The BI container to add the connection to, for Tableau, Looker or Power BI credentials. Its type has to match what the credentials are for: a &#x60;looker&#x60; container takes both &#x60;looker&#x60; and &#x60;looker-git-clone&#x60; credentials. A &#x60;custom-bi-connector&#x60; container takes a custom BI connector&#39;s credentials, or none when it has no deployment. Send exactly one of this, &#x60;warehouse_id&#x60; and &#x60;etl_container_id&#x60;. | [optional] 
**EtlContainerId** | Pointer to **NullableString** | The ETL container to add the connection to, for ETL tool credentials such as Fivetran or Airflow. The container&#39;s type has to equal the credentials&#39; type, and the container must not have a connection yet. A &#x60;custom-etl-connector&#x60; container takes a custom ETL connector&#39;s credentials, or none when it has no deployment. Send exactly one of this, &#x60;warehouse_id&#x60; and &#x60;bi_container_id&#x60;. | [optional] 
**CredentialsId** | Pointer to **NullableString** | The credentials the connection reads with. They also decide the connection&#39;s type. Create them first, through one of the credentials endpoints. Required, except on a push-only container: a &#x60;custom-etl-connector&#x60; or &#x60;custom-bi-connector&#x60; container with no deployment. Its connection takes no credentials, and has the container&#39;s type. | [optional] 
**JobTypes** | Pointer to **[]string** | The jobs to run on this connection. Omit it to run what the connection type runs by default, which is what the app does. Which values are accepted depends on the connection type. &#x60;etl&#x60; on a Snowflake, Power BI or Salesforce Data Cloud connection also creates its ETL container. An empty list is not accepted; omit the field to take the defaults. | [optional] 

## Methods

### NewConnectionIn

`func NewConnectionIn(name string, ) *ConnectionIn`

NewConnectionIn instantiates a new ConnectionIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionInWithDefaults

`func NewConnectionInWithDefaults() *ConnectionIn`

NewConnectionInWithDefaults instantiates a new ConnectionIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectionIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionIn) SetName(v string)`

SetName sets Name field to given value.


### GetWarehouseId

`func (o *ConnectionIn) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ConnectionIn) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ConnectionIn) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ConnectionIn) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### SetWarehouseIdNil

`func (o *ConnectionIn) SetWarehouseIdNil(b bool)`

 SetWarehouseIdNil sets the value for WarehouseId to be an explicit nil

### UnsetWarehouseId
`func (o *ConnectionIn) UnsetWarehouseId()`

UnsetWarehouseId ensures that no value is present for WarehouseId, not even an explicit nil
### GetBiContainerId

`func (o *ConnectionIn) GetBiContainerId() string`

GetBiContainerId returns the BiContainerId field if non-nil, zero value otherwise.

### GetBiContainerIdOk

`func (o *ConnectionIn) GetBiContainerIdOk() (*string, bool)`

GetBiContainerIdOk returns a tuple with the BiContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiContainerId

`func (o *ConnectionIn) SetBiContainerId(v string)`

SetBiContainerId sets BiContainerId field to given value.

### HasBiContainerId

`func (o *ConnectionIn) HasBiContainerId() bool`

HasBiContainerId returns a boolean if a field has been set.

### SetBiContainerIdNil

`func (o *ConnectionIn) SetBiContainerIdNil(b bool)`

 SetBiContainerIdNil sets the value for BiContainerId to be an explicit nil

### UnsetBiContainerId
`func (o *ConnectionIn) UnsetBiContainerId()`

UnsetBiContainerId ensures that no value is present for BiContainerId, not even an explicit nil
### GetEtlContainerId

`func (o *ConnectionIn) GetEtlContainerId() string`

GetEtlContainerId returns the EtlContainerId field if non-nil, zero value otherwise.

### GetEtlContainerIdOk

`func (o *ConnectionIn) GetEtlContainerIdOk() (*string, bool)`

GetEtlContainerIdOk returns a tuple with the EtlContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlContainerId

`func (o *ConnectionIn) SetEtlContainerId(v string)`

SetEtlContainerId sets EtlContainerId field to given value.

### HasEtlContainerId

`func (o *ConnectionIn) HasEtlContainerId() bool`

HasEtlContainerId returns a boolean if a field has been set.

### SetEtlContainerIdNil

`func (o *ConnectionIn) SetEtlContainerIdNil(b bool)`

 SetEtlContainerIdNil sets the value for EtlContainerId to be an explicit nil

### UnsetEtlContainerId
`func (o *ConnectionIn) UnsetEtlContainerId()`

UnsetEtlContainerId ensures that no value is present for EtlContainerId, not even an explicit nil
### GetCredentialsId

`func (o *ConnectionIn) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *ConnectionIn) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *ConnectionIn) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.

### HasCredentialsId

`func (o *ConnectionIn) HasCredentialsId() bool`

HasCredentialsId returns a boolean if a field has been set.

### SetCredentialsIdNil

`func (o *ConnectionIn) SetCredentialsIdNil(b bool)`

 SetCredentialsIdNil sets the value for CredentialsId to be an explicit nil

### UnsetCredentialsId
`func (o *ConnectionIn) UnsetCredentialsId()`

UnsetCredentialsId ensures that no value is present for CredentialsId, not even an explicit nil
### GetJobTypes

`func (o *ConnectionIn) GetJobTypes() []string`

GetJobTypes returns the JobTypes field if non-nil, zero value otherwise.

### GetJobTypesOk

`func (o *ConnectionIn) GetJobTypesOk() (*[]string, bool)`

GetJobTypesOk returns a tuple with the JobTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypes

`func (o *ConnectionIn) SetJobTypes(v []string)`

SetJobTypes sets JobTypes field to given value.

### HasJobTypes

`func (o *ConnectionIn) HasJobTypes() bool`

HasJobTypes returns a boolean if a field has been set.

### SetJobTypesNil

`func (o *ConnectionIn) SetJobTypesNil(b bool)`

 SetJobTypesNil sets the value for JobTypes to be an explicit nil

### UnsetJobTypes
`func (o *ConnectionIn) UnsetJobTypes()`

UnsetJobTypes ensures that no value is present for JobTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


