# DbtCloudSettingsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionMatch** | [**ConnectionMatch**](ConnectionMatch.md) | How a run is matched to the connection&#39;s warehouse. &#x60;off&#x60;, the default, drops no run for its dbt connection. | 
**AllowedConnections** | **[]string** | The dbt connections, by id or name, whose runs are kept. At least one for &#x60;allow_list&#x60;. Empty otherwise. | 
**AllowedEnvironments** | **[]string** | The dbt environments, by name and ignoring case, whose runs are kept. Empty keeps runs from every environment. | 

## Methods

### NewDbtCloudSettingsOut

`func NewDbtCloudSettingsOut(connectionMatch ConnectionMatch, allowedConnections []string, allowedEnvironments []string, ) *DbtCloudSettingsOut`

NewDbtCloudSettingsOut instantiates a new DbtCloudSettingsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbtCloudSettingsOutWithDefaults

`func NewDbtCloudSettingsOutWithDefaults() *DbtCloudSettingsOut`

NewDbtCloudSettingsOutWithDefaults instantiates a new DbtCloudSettingsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionMatch

`func (o *DbtCloudSettingsOut) GetConnectionMatch() ConnectionMatch`

GetConnectionMatch returns the ConnectionMatch field if non-nil, zero value otherwise.

### GetConnectionMatchOk

`func (o *DbtCloudSettingsOut) GetConnectionMatchOk() (*ConnectionMatch, bool)`

GetConnectionMatchOk returns a tuple with the ConnectionMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionMatch

`func (o *DbtCloudSettingsOut) SetConnectionMatch(v ConnectionMatch)`

SetConnectionMatch sets ConnectionMatch field to given value.


### GetAllowedConnections

`func (o *DbtCloudSettingsOut) GetAllowedConnections() []string`

GetAllowedConnections returns the AllowedConnections field if non-nil, zero value otherwise.

### GetAllowedConnectionsOk

`func (o *DbtCloudSettingsOut) GetAllowedConnectionsOk() (*[]string, bool)`

GetAllowedConnectionsOk returns a tuple with the AllowedConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedConnections

`func (o *DbtCloudSettingsOut) SetAllowedConnections(v []string)`

SetAllowedConnections sets AllowedConnections field to given value.


### GetAllowedEnvironments

`func (o *DbtCloudSettingsOut) GetAllowedEnvironments() []string`

GetAllowedEnvironments returns the AllowedEnvironments field if non-nil, zero value otherwise.

### GetAllowedEnvironmentsOk

`func (o *DbtCloudSettingsOut) GetAllowedEnvironmentsOk() (*[]string, bool)`

GetAllowedEnvironmentsOk returns a tuple with the AllowedEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedEnvironments

`func (o *DbtCloudSettingsOut) SetAllowedEnvironments(v []string)`

SetAllowedEnvironments sets AllowedEnvironments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


