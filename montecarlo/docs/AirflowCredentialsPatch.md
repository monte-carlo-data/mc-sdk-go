# AirflowCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | Pointer to **NullableString** | Host name of the Airflow web server, as Airflow reports it to Monte Carlo. | [optional] 

## Methods

### NewAirflowCredentialsPatch

`func NewAirflowCredentialsPatch() *AirflowCredentialsPatch`

NewAirflowCredentialsPatch instantiates a new AirflowCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirflowCredentialsPatchWithDefaults

`func NewAirflowCredentialsPatchWithDefaults() *AirflowCredentialsPatch`

NewAirflowCredentialsPatchWithDefaults instantiates a new AirflowCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *AirflowCredentialsPatch) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *AirflowCredentialsPatch) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *AirflowCredentialsPatch) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *AirflowCredentialsPatch) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *AirflowCredentialsPatch) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *AirflowCredentialsPatch) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


