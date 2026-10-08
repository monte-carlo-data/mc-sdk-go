# ConnectionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | New display name for the connection. Omit it to leave the name unchanged. An explicit null is ignored, the same as omitting the field. | [optional] 
**JobTypes** | Pointer to **[]string** | The connection&#39;s job types with &#x60;etl&#x60; added or removed. Adding &#x60;etl&#x60; turns on ETL collection for a Snowflake (Snowflake Tasks), Power BI (dataflows) or Salesforce Data Cloud connection, and creates the ETL container that &#x60;etl_container_id&#x60; then names. Removing it deletes that container. No other job can be added or removed. Omit it to leave the job types unchanged. | [optional] 

## Methods

### NewConnectionPatch

`func NewConnectionPatch() *ConnectionPatch`

NewConnectionPatch instantiates a new ConnectionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionPatchWithDefaults

`func NewConnectionPatchWithDefaults() *ConnectionPatch`

NewConnectionPatchWithDefaults instantiates a new ConnectionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectionPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionPatch) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConnectionPatch) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConnectionPatch) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetJobTypes

`func (o *ConnectionPatch) GetJobTypes() []string`

GetJobTypes returns the JobTypes field if non-nil, zero value otherwise.

### GetJobTypesOk

`func (o *ConnectionPatch) GetJobTypesOk() (*[]string, bool)`

GetJobTypesOk returns a tuple with the JobTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypes

`func (o *ConnectionPatch) SetJobTypes(v []string)`

SetJobTypes sets JobTypes field to given value.

### HasJobTypes

`func (o *ConnectionPatch) HasJobTypes() bool`

HasJobTypes returns a boolean if a field has been set.

### SetJobTypesNil

`func (o *ConnectionPatch) SetJobTypesNil(b bool)`

 SetJobTypesNil sets the value for JobTypes to be an explicit nil

### UnsetJobTypes
`func (o *ConnectionPatch) UnsetJobTypes()`

UnsetJobTypes ensures that no value is present for JobTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


