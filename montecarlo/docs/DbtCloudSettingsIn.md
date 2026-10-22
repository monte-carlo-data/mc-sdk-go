# DbtCloudSettingsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionMatch** | [**ConnectionMatch**](ConnectionMatch.md) | How a run is matched to the connection&#39;s warehouse. &#x60;off&#x60;, the default, drops no run for its dbt connection. | 
**AllowedConnectionFields** | **[]string** | The dbt connection fields to compare, such as &#x60;account&#x60; for Snowflake. &#x60;name&#x60; and &#x60;id&#x60; match a dbt connection by its own name or id. At least one for &#x60;allow_list&#x60;. Empty otherwise. | 
**AllowedConnectionValues** | **[]string** | The values to match, exactly as they appear in dbt Cloud. A run is kept when one of &#x60;allowed_connection_fields&#x60; on its dbt connection equals one of them. At least one for &#x60;allow_list&#x60;. Empty otherwise. | 
**AllowedEnvironments** | **[]string** | The dbt environments, by name and ignoring case, whose runs are kept. Empty keeps runs from every environment. | 

## Methods

### NewDbtCloudSettingsIn

`func NewDbtCloudSettingsIn(connectionMatch ConnectionMatch, allowedConnectionFields []string, allowedConnectionValues []string, allowedEnvironments []string, ) *DbtCloudSettingsIn`

NewDbtCloudSettingsIn instantiates a new DbtCloudSettingsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbtCloudSettingsInWithDefaults

`func NewDbtCloudSettingsInWithDefaults() *DbtCloudSettingsIn`

NewDbtCloudSettingsInWithDefaults instantiates a new DbtCloudSettingsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionMatch

`func (o *DbtCloudSettingsIn) GetConnectionMatch() ConnectionMatch`

GetConnectionMatch returns the ConnectionMatch field if non-nil, zero value otherwise.

### GetConnectionMatchOk

`func (o *DbtCloudSettingsIn) GetConnectionMatchOk() (*ConnectionMatch, bool)`

GetConnectionMatchOk returns a tuple with the ConnectionMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionMatch

`func (o *DbtCloudSettingsIn) SetConnectionMatch(v ConnectionMatch)`

SetConnectionMatch sets ConnectionMatch field to given value.


### GetAllowedConnectionFields

`func (o *DbtCloudSettingsIn) GetAllowedConnectionFields() []string`

GetAllowedConnectionFields returns the AllowedConnectionFields field if non-nil, zero value otherwise.

### GetAllowedConnectionFieldsOk

`func (o *DbtCloudSettingsIn) GetAllowedConnectionFieldsOk() (*[]string, bool)`

GetAllowedConnectionFieldsOk returns a tuple with the AllowedConnectionFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedConnectionFields

`func (o *DbtCloudSettingsIn) SetAllowedConnectionFields(v []string)`

SetAllowedConnectionFields sets AllowedConnectionFields field to given value.


### GetAllowedConnectionValues

`func (o *DbtCloudSettingsIn) GetAllowedConnectionValues() []string`

GetAllowedConnectionValues returns the AllowedConnectionValues field if non-nil, zero value otherwise.

### GetAllowedConnectionValuesOk

`func (o *DbtCloudSettingsIn) GetAllowedConnectionValuesOk() (*[]string, bool)`

GetAllowedConnectionValuesOk returns a tuple with the AllowedConnectionValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedConnectionValues

`func (o *DbtCloudSettingsIn) SetAllowedConnectionValues(v []string)`

SetAllowedConnectionValues sets AllowedConnectionValues field to given value.


### GetAllowedEnvironments

`func (o *DbtCloudSettingsIn) GetAllowedEnvironments() []string`

GetAllowedEnvironments returns the AllowedEnvironments field if non-nil, zero value otherwise.

### GetAllowedEnvironmentsOk

`func (o *DbtCloudSettingsIn) GetAllowedEnvironmentsOk() (*[]string, bool)`

GetAllowedEnvironmentsOk returns a tuple with the AllowedEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnvironments

`func (o *DbtCloudSettingsIn) SetAllowedEnvironments(v []string)`

SetAllowedEnvironments sets AllowedEnvironments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


