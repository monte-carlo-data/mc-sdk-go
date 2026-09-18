# ConnectionIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for the connection. Unique among the warehouse&#39;s connections. | 
**WarehouseId** | **string** | The warehouse to add the connection to. Its type has to match what the credentials are for. | 
**CredentialsId** | **string** | The credentials the connection reads with. They also decide the connection&#39;s type. Create them first, through one of the credentials endpoints. | 
**JobTypes** | Pointer to **[]string** | The jobs to run on this connection. Omit it to run what the connection type runs by default, which is what the app does. Which values are accepted depends on the connection type. An empty list is not accepted; omit the field to take the defaults. | [optional] 

## Methods

### NewConnectionIn

`func NewConnectionIn(name string, warehouseId string, credentialsId string, ) *ConnectionIn`

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


