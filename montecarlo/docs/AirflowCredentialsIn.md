# AirflowCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | **string** | Host name of the Airflow web server, as Airflow reports it to Monte Carlo. | 

## Methods

### NewAirflowCredentialsIn

`func NewAirflowCredentialsIn(hostName string, ) *AirflowCredentialsIn`

NewAirflowCredentialsIn instantiates a new AirflowCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirflowCredentialsInWithDefaults

`func NewAirflowCredentialsInWithDefaults() *AirflowCredentialsIn`

NewAirflowCredentialsInWithDefaults instantiates a new AirflowCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *AirflowCredentialsIn) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *AirflowCredentialsIn) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *AirflowCredentialsIn) SetHostName(v string)`

SetHostName sets HostName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


