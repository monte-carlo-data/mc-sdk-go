# CustomConnectorCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsMetadata** | **bool** | Whether the connector supports metadata collection. | 
**SupportsQueryLogs** | **bool** | Whether the connector supports query log collection. | 
**SupportsLineage** | **bool** | Whether the connector supports lineage extraction. | 
**SupportsFieldLineage** | **bool** | Whether the connector supports field-level lineage. | 
**SupportsSchema** | **bool** | Whether the connector supports schema collection. | 
**SupportsFreshness** | **bool** | Whether the connector supports freshness monitoring. | 
**SupportsVolumeRows** | **bool** | Whether the connector supports row-count volume monitoring. | 
**SupportsVolumeBytes** | **bool** | Whether the connector supports byte-size volume monitoring. | 
**SupportsCustomSqlMonitor** | **bool** | Whether the connector supports custom SQL monitors. | 
**SupportsFullQueryLanguage** | **bool** | Whether the connector supports the full query language. | 
**SupportsStatsMonitor** | **bool** | Whether the connector supports stats monitors. | 
**SupportsComparisonMonitor** | **bool** | Whether the connector supports comparison monitors. | 
**SupportsValidationMonitor** | **bool** | Whether the connector supports validation monitors. | 
**SupportsJsonSchemaMonitor** | **bool** | Whether the connector supports JSON schema monitors. | 
**SupportsTransform** | **bool** | Whether the connector supports transforms. | 
**SupportsAgentMonitor** | **bool** | Whether the connector supports agent monitors. | 
**SupportsAgentTrajectoryMonitor** | **bool** | Whether the connector supports agent trajectory monitors. | 
**SupportsNonMetadataSizeCollection** | **bool** | Whether the connector supports size collection without metadata. | 
**SupportsMetricMonitorsOnUnknownTypes** | **bool** | Whether the connector supports metric monitors on unknown field types. | 

## Methods

### NewCustomConnectorCapabilities

`func NewCustomConnectorCapabilities(supportsMetadata bool, supportsQueryLogs bool, supportsLineage bool, supportsFieldLineage bool, supportsSchema bool, supportsFreshness bool, supportsVolumeRows bool, supportsVolumeBytes bool, supportsCustomSqlMonitor bool, supportsFullQueryLanguage bool, supportsStatsMonitor bool, supportsComparisonMonitor bool, supportsValidationMonitor bool, supportsJsonSchemaMonitor bool, supportsTransform bool, supportsAgentMonitor bool, supportsAgentTrajectoryMonitor bool, supportsNonMetadataSizeCollection bool, supportsMetricMonitorsOnUnknownTypes bool, ) *CustomConnectorCapabilities`

NewCustomConnectorCapabilities instantiates a new CustomConnectorCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomConnectorCapabilitiesWithDefaults

`func NewCustomConnectorCapabilitiesWithDefaults() *CustomConnectorCapabilities`

NewCustomConnectorCapabilitiesWithDefaults instantiates a new CustomConnectorCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportsMetadata

`func (o *CustomConnectorCapabilities) GetSupportsMetadata() bool`

GetSupportsMetadata returns the SupportsMetadata field if non-nil, zero value otherwise.

### GetSupportsMetadataOk

`func (o *CustomConnectorCapabilities) GetSupportsMetadataOk() (*bool, bool)`

GetSupportsMetadataOk returns a tuple with the SupportsMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMetadata

`func (o *CustomConnectorCapabilities) SetSupportsMetadata(v bool)`

SetSupportsMetadata sets SupportsMetadata field to given value.


### GetSupportsQueryLogs

`func (o *CustomConnectorCapabilities) GetSupportsQueryLogs() bool`

GetSupportsQueryLogs returns the SupportsQueryLogs field if non-nil, zero value otherwise.

### GetSupportsQueryLogsOk

`func (o *CustomConnectorCapabilities) GetSupportsQueryLogsOk() (*bool, bool)`

GetSupportsQueryLogsOk returns a tuple with the SupportsQueryLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsQueryLogs

`func (o *CustomConnectorCapabilities) SetSupportsQueryLogs(v bool)`

SetSupportsQueryLogs sets SupportsQueryLogs field to given value.


### GetSupportsLineage

`func (o *CustomConnectorCapabilities) GetSupportsLineage() bool`

GetSupportsLineage returns the SupportsLineage field if non-nil, zero value otherwise.

### GetSupportsLineageOk

`func (o *CustomConnectorCapabilities) GetSupportsLineageOk() (*bool, bool)`

GetSupportsLineageOk returns a tuple with the SupportsLineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsLineage

`func (o *CustomConnectorCapabilities) SetSupportsLineage(v bool)`

SetSupportsLineage sets SupportsLineage field to given value.


### GetSupportsFieldLineage

`func (o *CustomConnectorCapabilities) GetSupportsFieldLineage() bool`

GetSupportsFieldLineage returns the SupportsFieldLineage field if non-nil, zero value otherwise.

### GetSupportsFieldLineageOk

`func (o *CustomConnectorCapabilities) GetSupportsFieldLineageOk() (*bool, bool)`

GetSupportsFieldLineageOk returns a tuple with the SupportsFieldLineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsFieldLineage

`func (o *CustomConnectorCapabilities) SetSupportsFieldLineage(v bool)`

SetSupportsFieldLineage sets SupportsFieldLineage field to given value.


### GetSupportsSchema

`func (o *CustomConnectorCapabilities) GetSupportsSchema() bool`

GetSupportsSchema returns the SupportsSchema field if non-nil, zero value otherwise.

### GetSupportsSchemaOk

`func (o *CustomConnectorCapabilities) GetSupportsSchemaOk() (*bool, bool)`

GetSupportsSchemaOk returns a tuple with the SupportsSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSchema

`func (o *CustomConnectorCapabilities) SetSupportsSchema(v bool)`

SetSupportsSchema sets SupportsSchema field to given value.


### GetSupportsFreshness

`func (o *CustomConnectorCapabilities) GetSupportsFreshness() bool`

GetSupportsFreshness returns the SupportsFreshness field if non-nil, zero value otherwise.

### GetSupportsFreshnessOk

`func (o *CustomConnectorCapabilities) GetSupportsFreshnessOk() (*bool, bool)`

GetSupportsFreshnessOk returns a tuple with the SupportsFreshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsFreshness

`func (o *CustomConnectorCapabilities) SetSupportsFreshness(v bool)`

SetSupportsFreshness sets SupportsFreshness field to given value.


### GetSupportsVolumeRows

`func (o *CustomConnectorCapabilities) GetSupportsVolumeRows() bool`

GetSupportsVolumeRows returns the SupportsVolumeRows field if non-nil, zero value otherwise.

### GetSupportsVolumeRowsOk

`func (o *CustomConnectorCapabilities) GetSupportsVolumeRowsOk() (*bool, bool)`

GetSupportsVolumeRowsOk returns a tuple with the SupportsVolumeRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsVolumeRows

`func (o *CustomConnectorCapabilities) SetSupportsVolumeRows(v bool)`

SetSupportsVolumeRows sets SupportsVolumeRows field to given value.


### GetSupportsVolumeBytes

`func (o *CustomConnectorCapabilities) GetSupportsVolumeBytes() bool`

GetSupportsVolumeBytes returns the SupportsVolumeBytes field if non-nil, zero value otherwise.

### GetSupportsVolumeBytesOk

`func (o *CustomConnectorCapabilities) GetSupportsVolumeBytesOk() (*bool, bool)`

GetSupportsVolumeBytesOk returns a tuple with the SupportsVolumeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsVolumeBytes

`func (o *CustomConnectorCapabilities) SetSupportsVolumeBytes(v bool)`

SetSupportsVolumeBytes sets SupportsVolumeBytes field to given value.


### GetSupportsCustomSqlMonitor

`func (o *CustomConnectorCapabilities) GetSupportsCustomSqlMonitor() bool`

GetSupportsCustomSqlMonitor returns the SupportsCustomSqlMonitor field if non-nil, zero value otherwise.

### GetSupportsCustomSqlMonitorOk

`func (o *CustomConnectorCapabilities) GetSupportsCustomSqlMonitorOk() (*bool, bool)`

GetSupportsCustomSqlMonitorOk returns a tuple with the SupportsCustomSqlMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCustomSqlMonitor

`func (o *CustomConnectorCapabilities) SetSupportsCustomSqlMonitor(v bool)`

SetSupportsCustomSqlMonitor sets SupportsCustomSqlMonitor field to given value.


### GetSupportsFullQueryLanguage

`func (o *CustomConnectorCapabilities) GetSupportsFullQueryLanguage() bool`

GetSupportsFullQueryLanguage returns the SupportsFullQueryLanguage field if non-nil, zero value otherwise.

### GetSupportsFullQueryLanguageOk

`func (o *CustomConnectorCapabilities) GetSupportsFullQueryLanguageOk() (*bool, bool)`

GetSupportsFullQueryLanguageOk returns a tuple with the SupportsFullQueryLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsFullQueryLanguage

`func (o *CustomConnectorCapabilities) SetSupportsFullQueryLanguage(v bool)`

SetSupportsFullQueryLanguage sets SupportsFullQueryLanguage field to given value.


### GetSupportsStatsMonitor

`func (o *CustomConnectorCapabilities) GetSupportsStatsMonitor() bool`

GetSupportsStatsMonitor returns the SupportsStatsMonitor field if non-nil, zero value otherwise.

### GetSupportsStatsMonitorOk

`func (o *CustomConnectorCapabilities) GetSupportsStatsMonitorOk() (*bool, bool)`

GetSupportsStatsMonitorOk returns a tuple with the SupportsStatsMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsStatsMonitor

`func (o *CustomConnectorCapabilities) SetSupportsStatsMonitor(v bool)`

SetSupportsStatsMonitor sets SupportsStatsMonitor field to given value.


### GetSupportsComparisonMonitor

`func (o *CustomConnectorCapabilities) GetSupportsComparisonMonitor() bool`

GetSupportsComparisonMonitor returns the SupportsComparisonMonitor field if non-nil, zero value otherwise.

### GetSupportsComparisonMonitorOk

`func (o *CustomConnectorCapabilities) GetSupportsComparisonMonitorOk() (*bool, bool)`

GetSupportsComparisonMonitorOk returns a tuple with the SupportsComparisonMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsComparisonMonitor

`func (o *CustomConnectorCapabilities) SetSupportsComparisonMonitor(v bool)`

SetSupportsComparisonMonitor sets SupportsComparisonMonitor field to given value.


### GetSupportsValidationMonitor

`func (o *CustomConnectorCapabilities) GetSupportsValidationMonitor() bool`

GetSupportsValidationMonitor returns the SupportsValidationMonitor field if non-nil, zero value otherwise.

### GetSupportsValidationMonitorOk

`func (o *CustomConnectorCapabilities) GetSupportsValidationMonitorOk() (*bool, bool)`

GetSupportsValidationMonitorOk returns a tuple with the SupportsValidationMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsValidationMonitor

`func (o *CustomConnectorCapabilities) SetSupportsValidationMonitor(v bool)`

SetSupportsValidationMonitor sets SupportsValidationMonitor field to given value.


### GetSupportsJsonSchemaMonitor

`func (o *CustomConnectorCapabilities) GetSupportsJsonSchemaMonitor() bool`

GetSupportsJsonSchemaMonitor returns the SupportsJsonSchemaMonitor field if non-nil, zero value otherwise.

### GetSupportsJsonSchemaMonitorOk

`func (o *CustomConnectorCapabilities) GetSupportsJsonSchemaMonitorOk() (*bool, bool)`

GetSupportsJsonSchemaMonitorOk returns a tuple with the SupportsJsonSchemaMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsJsonSchemaMonitor

`func (o *CustomConnectorCapabilities) SetSupportsJsonSchemaMonitor(v bool)`

SetSupportsJsonSchemaMonitor sets SupportsJsonSchemaMonitor field to given value.


### GetSupportsTransform

`func (o *CustomConnectorCapabilities) GetSupportsTransform() bool`

GetSupportsTransform returns the SupportsTransform field if non-nil, zero value otherwise.

### GetSupportsTransformOk

`func (o *CustomConnectorCapabilities) GetSupportsTransformOk() (*bool, bool)`

GetSupportsTransformOk returns a tuple with the SupportsTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsTransform

`func (o *CustomConnectorCapabilities) SetSupportsTransform(v bool)`

SetSupportsTransform sets SupportsTransform field to given value.


### GetSupportsAgentMonitor

`func (o *CustomConnectorCapabilities) GetSupportsAgentMonitor() bool`

GetSupportsAgentMonitor returns the SupportsAgentMonitor field if non-nil, zero value otherwise.

### GetSupportsAgentMonitorOk

`func (o *CustomConnectorCapabilities) GetSupportsAgentMonitorOk() (*bool, bool)`

GetSupportsAgentMonitorOk returns a tuple with the SupportsAgentMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAgentMonitor

`func (o *CustomConnectorCapabilities) SetSupportsAgentMonitor(v bool)`

SetSupportsAgentMonitor sets SupportsAgentMonitor field to given value.


### GetSupportsAgentTrajectoryMonitor

`func (o *CustomConnectorCapabilities) GetSupportsAgentTrajectoryMonitor() bool`

GetSupportsAgentTrajectoryMonitor returns the SupportsAgentTrajectoryMonitor field if non-nil, zero value otherwise.

### GetSupportsAgentTrajectoryMonitorOk

`func (o *CustomConnectorCapabilities) GetSupportsAgentTrajectoryMonitorOk() (*bool, bool)`

GetSupportsAgentTrajectoryMonitorOk returns a tuple with the SupportsAgentTrajectoryMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAgentTrajectoryMonitor

`func (o *CustomConnectorCapabilities) SetSupportsAgentTrajectoryMonitor(v bool)`

SetSupportsAgentTrajectoryMonitor sets SupportsAgentTrajectoryMonitor field to given value.


### GetSupportsNonMetadataSizeCollection

`func (o *CustomConnectorCapabilities) GetSupportsNonMetadataSizeCollection() bool`

GetSupportsNonMetadataSizeCollection returns the SupportsNonMetadataSizeCollection field if non-nil, zero value otherwise.

### GetSupportsNonMetadataSizeCollectionOk

`func (o *CustomConnectorCapabilities) GetSupportsNonMetadataSizeCollectionOk() (*bool, bool)`

GetSupportsNonMetadataSizeCollectionOk returns a tuple with the SupportsNonMetadataSizeCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsNonMetadataSizeCollection

`func (o *CustomConnectorCapabilities) SetSupportsNonMetadataSizeCollection(v bool)`

SetSupportsNonMetadataSizeCollection sets SupportsNonMetadataSizeCollection field to given value.


### GetSupportsMetricMonitorsOnUnknownTypes

`func (o *CustomConnectorCapabilities) GetSupportsMetricMonitorsOnUnknownTypes() bool`

GetSupportsMetricMonitorsOnUnknownTypes returns the SupportsMetricMonitorsOnUnknownTypes field if non-nil, zero value otherwise.

### GetSupportsMetricMonitorsOnUnknownTypesOk

`func (o *CustomConnectorCapabilities) GetSupportsMetricMonitorsOnUnknownTypesOk() (*bool, bool)`

GetSupportsMetricMonitorsOnUnknownTypesOk returns a tuple with the SupportsMetricMonitorsOnUnknownTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMetricMonitorsOnUnknownTypes

`func (o *CustomConnectorCapabilities) SetSupportsMetricMonitorsOnUnknownTypes(v bool)`

SetSupportsMetricMonitorsOnUnknownTypes sets SupportsMetricMonitorsOnUnknownTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


