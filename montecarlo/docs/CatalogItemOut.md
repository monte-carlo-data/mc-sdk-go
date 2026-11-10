# CatalogItemOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name to use in a query. | 
**Label** | Pointer to **NullableString** | A short display name. | [optional] 
**Description** | Pointer to **NullableString** | What it counts or describes. | [optional] 
**RequiresMetricTime** | Pointer to **NullableBool** | Metrics only. True when the metric answers only grouped by &#x60;metric_time&#x60;, for example a month-to-date total. Left out for a dimension. | [optional] 
**HistoryStart** | Pointer to **NullableString** | Metrics only. The first date the metric keeps. A query that filters &#x60;metric_time&#x60; before it is refused, so start a window on or after it. Left out for a metric that keeps all its history. | [optional] 
**LongerHistory** | Pointer to **NullableString** | Metrics only. A metric that keeps dates before &#x60;history_start&#x60;. Left out when there is none. | [optional] 
**Format** | Pointer to [**NullableCatalogMetricFormat**](CatalogMetricFormat.md) | Metrics only. How to show the value: &#x60;percent&#x60; for a share from 0 to 1 (0.16 is 16%), &#x60;credits&#x60; or &#x60;hours&#x60;. Left out for a plain number. | [optional] 
**Type** | Pointer to [**NullableCatalogDimensionType**](CatalogDimensionType.md) | Dimensions only. Left out for a metric. | [optional] 
**TimeGrains** | Pointer to **[]string** | Time dimensions only. The grains a query can append to the name, such as &#x60;metric_time__week&#x60; for &#x60;week&#x60;. Left out otherwise. | [optional] 

## Methods

### NewCatalogItemOut

`func NewCatalogItemOut(name string, ) *CatalogItemOut`

NewCatalogItemOut instantiates a new CatalogItemOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemOutWithDefaults

`func NewCatalogItemOutWithDefaults() *CatalogItemOut`

NewCatalogItemOutWithDefaults instantiates a new CatalogItemOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogItemOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogItemOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogItemOut) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *CatalogItemOut) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CatalogItemOut) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CatalogItemOut) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CatalogItemOut) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *CatalogItemOut) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *CatalogItemOut) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDescription

`func (o *CatalogItemOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogItemOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogItemOut) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogItemOut) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CatalogItemOut) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CatalogItemOut) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRequiresMetricTime

`func (o *CatalogItemOut) GetRequiresMetricTime() bool`

GetRequiresMetricTime returns the RequiresMetricTime field if non-nil, zero value otherwise.

### GetRequiresMetricTimeOk

`func (o *CatalogItemOut) GetRequiresMetricTimeOk() (*bool, bool)`

GetRequiresMetricTimeOk returns a tuple with the RequiresMetricTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresMetricTime

`func (o *CatalogItemOut) SetRequiresMetricTime(v bool)`

SetRequiresMetricTime sets RequiresMetricTime field to given value.

### HasRequiresMetricTime

`func (o *CatalogItemOut) HasRequiresMetricTime() bool`

HasRequiresMetricTime returns a boolean if a field has been set.

### SetRequiresMetricTimeNil

`func (o *CatalogItemOut) SetRequiresMetricTimeNil(b bool)`

 SetRequiresMetricTimeNil sets the value for RequiresMetricTime to be an explicit nil

### UnsetRequiresMetricTime
`func (o *CatalogItemOut) UnsetRequiresMetricTime()`

UnsetRequiresMetricTime ensures that no value is present for RequiresMetricTime, not even an explicit nil
### GetHistoryStart

`func (o *CatalogItemOut) GetHistoryStart() string`

GetHistoryStart returns the HistoryStart field if non-nil, zero value otherwise.

### GetHistoryStartOk

`func (o *CatalogItemOut) GetHistoryStartOk() (*string, bool)`

GetHistoryStartOk returns a tuple with the HistoryStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryStart

`func (o *CatalogItemOut) SetHistoryStart(v string)`

SetHistoryStart sets HistoryStart field to given value.

### HasHistoryStart

`func (o *CatalogItemOut) HasHistoryStart() bool`

HasHistoryStart returns a boolean if a field has been set.

### SetHistoryStartNil

`func (o *CatalogItemOut) SetHistoryStartNil(b bool)`

 SetHistoryStartNil sets the value for HistoryStart to be an explicit nil

### UnsetHistoryStart
`func (o *CatalogItemOut) UnsetHistoryStart()`

UnsetHistoryStart ensures that no value is present for HistoryStart, not even an explicit nil
### GetLongerHistory

`func (o *CatalogItemOut) GetLongerHistory() string`

GetLongerHistory returns the LongerHistory field if non-nil, zero value otherwise.

### GetLongerHistoryOk

`func (o *CatalogItemOut) GetLongerHistoryOk() (*string, bool)`

GetLongerHistoryOk returns a tuple with the LongerHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongerHistory

`func (o *CatalogItemOut) SetLongerHistory(v string)`

SetLongerHistory sets LongerHistory field to given value.

### HasLongerHistory

`func (o *CatalogItemOut) HasLongerHistory() bool`

HasLongerHistory returns a boolean if a field has been set.

### SetLongerHistoryNil

`func (o *CatalogItemOut) SetLongerHistoryNil(b bool)`

 SetLongerHistoryNil sets the value for LongerHistory to be an explicit nil

### UnsetLongerHistory
`func (o *CatalogItemOut) UnsetLongerHistory()`

UnsetLongerHistory ensures that no value is present for LongerHistory, not even an explicit nil
### GetFormat

`func (o *CatalogItemOut) GetFormat() CatalogMetricFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CatalogItemOut) GetFormatOk() (*CatalogMetricFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CatalogItemOut) SetFormat(v CatalogMetricFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CatalogItemOut) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *CatalogItemOut) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *CatalogItemOut) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetType

`func (o *CatalogItemOut) GetType() CatalogDimensionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogItemOut) GetTypeOk() (*CatalogDimensionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogItemOut) SetType(v CatalogDimensionType)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogItemOut) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CatalogItemOut) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CatalogItemOut) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTimeGrains

`func (o *CatalogItemOut) GetTimeGrains() []string`

GetTimeGrains returns the TimeGrains field if non-nil, zero value otherwise.

### GetTimeGrainsOk

`func (o *CatalogItemOut) GetTimeGrainsOk() (*[]string, bool)`

GetTimeGrainsOk returns a tuple with the TimeGrains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeGrains

`func (o *CatalogItemOut) SetTimeGrains(v []string)`

SetTimeGrains sets TimeGrains field to given value.

### HasTimeGrains

`func (o *CatalogItemOut) HasTimeGrains() bool`

HasTimeGrains returns a boolean if a field has been set.

### SetTimeGrainsNil

`func (o *CatalogItemOut) SetTimeGrainsNil(b bool)`

 SetTimeGrainsNil sets the value for TimeGrains to be an explicit nil

### UnsetTimeGrains
`func (o *CatalogItemOut) UnsetTimeGrains()`

UnsetTimeGrains ensures that no value is present for TimeGrains, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


