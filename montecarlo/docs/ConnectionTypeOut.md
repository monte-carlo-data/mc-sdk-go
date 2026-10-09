# ConnectionTypeOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | The connection type. Credentials for it name this type, and the connection takes its type from them. | 
**Name** | **string** | Display name of the connection type. | 
**Parent** | [**ConnectionParent**](ConnectionParent.md) | What a connection of this type is added to. Create or reuse one first. | 
**WarehouseType** | [**NullableWarehouseType**](WarehouseType.md) | The type of warehouse the connection goes on. Null unless &#x60;parent&#x60; is &#x60;warehouse&#x60;. | 
**BiContainerType** | [**NullableNewBiContainerType**](NewBiContainerType.md) | The type of BI container the connection goes on. Null unless &#x60;parent&#x60; is &#x60;bi_container&#x60;. | 
**EtlContainerType** | [**NullableNewEtlContainerType**](NewEtlContainerType.md) | The type of ETL container the connection goes on. Null unless &#x60;parent&#x60; is &#x60;etl_container&#x60;. | 
**CredentialsRequired** | **bool** | Whether the connection needs credentials. False for a push-only type: its data is sent to Monte Carlo, and Monte Carlo reaches no system for it. | 
**McManagedCredentialsOperation** | **NullableString** | The operation that stores this type&#39;s credentials with Monte Carlo. Null when Monte Carlo cannot store them. | 
**McManagedCredentialsSecret** | **NullableBool** | Whether storing the credentials with Monte Carlo sends a secret, such as a password or a private key. Null when Monte Carlo cannot store them. | 
**SelfHostedCredentialsStorages** | [**[]CredentialsStorageType**](CredentialsStorageType.md) | The stores you run that the credentials can be kept in, with Monte Carlo keeping only where they are. Empty when they cannot be. | 
**RequiresDeployment** | **bool** | Whether the connection runs through its parent&#39;s deployment. False for a type that reports to Monte Carlo itself, such as Airflow, or whose data is pushed. | 
**RequiresOneOfConnectionTypes** | **[]string** | The warehouse needs a connection of one of these types before one of this type is added. Empty when it needs none. | 
**JobTypes** | **[]string** | The collection jobs a connection of this type can run. | 
**IsCustom** | **bool** | Whether a collection agent registered the type for your account. A custom type&#39;s connection runs through the deployment of that agent. | 

## Methods

### NewConnectionTypeOut

`func NewConnectionTypeOut(connectionType string, name string, parent ConnectionParent, warehouseType NullableWarehouseType, biContainerType NullableNewBiContainerType, etlContainerType NullableNewEtlContainerType, credentialsRequired bool, mcManagedCredentialsOperation NullableString, mcManagedCredentialsSecret NullableBool, selfHostedCredentialsStorages []CredentialsStorageType, requiresDeployment bool, requiresOneOfConnectionTypes []string, jobTypes []string, isCustom bool, ) *ConnectionTypeOut`

NewConnectionTypeOut instantiates a new ConnectionTypeOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionTypeOutWithDefaults

`func NewConnectionTypeOutWithDefaults() *ConnectionTypeOut`

NewConnectionTypeOutWithDefaults instantiates a new ConnectionTypeOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *ConnectionTypeOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ConnectionTypeOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ConnectionTypeOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetName

`func (o *ConnectionTypeOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionTypeOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionTypeOut) SetName(v string)`

SetName sets Name field to given value.


### GetParent

`func (o *ConnectionTypeOut) GetParent() ConnectionParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ConnectionTypeOut) GetParentOk() (*ConnectionParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ConnectionTypeOut) SetParent(v ConnectionParent)`

SetParent sets Parent field to given value.


### GetWarehouseType

`func (o *ConnectionTypeOut) GetWarehouseType() WarehouseType`

GetWarehouseType returns the WarehouseType field if non-nil, zero value otherwise.

### GetWarehouseTypeOk

`func (o *ConnectionTypeOut) GetWarehouseTypeOk() (*WarehouseType, bool)`

GetWarehouseTypeOk returns a tuple with the WarehouseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseType

`func (o *ConnectionTypeOut) SetWarehouseType(v WarehouseType)`

SetWarehouseType sets WarehouseType field to given value.


### SetWarehouseTypeNil

`func (o *ConnectionTypeOut) SetWarehouseTypeNil(b bool)`

 SetWarehouseTypeNil sets the value for WarehouseType to be an explicit nil

### UnsetWarehouseType
`func (o *ConnectionTypeOut) UnsetWarehouseType()`

UnsetWarehouseType ensures that no value is present for WarehouseType, not even an explicit nil
### GetBiContainerType

`func (o *ConnectionTypeOut) GetBiContainerType() NewBiContainerType`

GetBiContainerType returns the BiContainerType field if non-nil, zero value otherwise.

### GetBiContainerTypeOk

`func (o *ConnectionTypeOut) GetBiContainerTypeOk() (*NewBiContainerType, bool)`

GetBiContainerTypeOk returns a tuple with the BiContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiContainerType

`func (o *ConnectionTypeOut) SetBiContainerType(v NewBiContainerType)`

SetBiContainerType sets BiContainerType field to given value.


### SetBiContainerTypeNil

`func (o *ConnectionTypeOut) SetBiContainerTypeNil(b bool)`

 SetBiContainerTypeNil sets the value for BiContainerType to be an explicit nil

### UnsetBiContainerType
`func (o *ConnectionTypeOut) UnsetBiContainerType()`

UnsetBiContainerType ensures that no value is present for BiContainerType, not even an explicit nil
### GetEtlContainerType

`func (o *ConnectionTypeOut) GetEtlContainerType() NewEtlContainerType`

GetEtlContainerType returns the EtlContainerType field if non-nil, zero value otherwise.

### GetEtlContainerTypeOk

`func (o *ConnectionTypeOut) GetEtlContainerTypeOk() (*NewEtlContainerType, bool)`

GetEtlContainerTypeOk returns a tuple with the EtlContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlContainerType

`func (o *ConnectionTypeOut) SetEtlContainerType(v NewEtlContainerType)`

SetEtlContainerType sets EtlContainerType field to given value.


### SetEtlContainerTypeNil

`func (o *ConnectionTypeOut) SetEtlContainerTypeNil(b bool)`

 SetEtlContainerTypeNil sets the value for EtlContainerType to be an explicit nil

### UnsetEtlContainerType
`func (o *ConnectionTypeOut) UnsetEtlContainerType()`

UnsetEtlContainerType ensures that no value is present for EtlContainerType, not even an explicit nil
### GetCredentialsRequired

`func (o *ConnectionTypeOut) GetCredentialsRequired() bool`

GetCredentialsRequired returns the CredentialsRequired field if non-nil, zero value otherwise.

### GetCredentialsRequiredOk

`func (o *ConnectionTypeOut) GetCredentialsRequiredOk() (*bool, bool)`

GetCredentialsRequiredOk returns a tuple with the CredentialsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsRequired

`func (o *ConnectionTypeOut) SetCredentialsRequired(v bool)`

SetCredentialsRequired sets CredentialsRequired field to given value.


### GetMcManagedCredentialsOperation

`func (o *ConnectionTypeOut) GetMcManagedCredentialsOperation() string`

GetMcManagedCredentialsOperation returns the McManagedCredentialsOperation field if non-nil, zero value otherwise.

### GetMcManagedCredentialsOperationOk

`func (o *ConnectionTypeOut) GetMcManagedCredentialsOperationOk() (*string, bool)`

GetMcManagedCredentialsOperationOk returns a tuple with the McManagedCredentialsOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcManagedCredentialsOperation

`func (o *ConnectionTypeOut) SetMcManagedCredentialsOperation(v string)`

SetMcManagedCredentialsOperation sets McManagedCredentialsOperation field to given value.


### SetMcManagedCredentialsOperationNil

`func (o *ConnectionTypeOut) SetMcManagedCredentialsOperationNil(b bool)`

 SetMcManagedCredentialsOperationNil sets the value for McManagedCredentialsOperation to be an explicit nil

### UnsetMcManagedCredentialsOperation
`func (o *ConnectionTypeOut) UnsetMcManagedCredentialsOperation()`

UnsetMcManagedCredentialsOperation ensures that no value is present for McManagedCredentialsOperation, not even an explicit nil
### GetMcManagedCredentialsSecret

`func (o *ConnectionTypeOut) GetMcManagedCredentialsSecret() bool`

GetMcManagedCredentialsSecret returns the McManagedCredentialsSecret field if non-nil, zero value otherwise.

### GetMcManagedCredentialsSecretOk

`func (o *ConnectionTypeOut) GetMcManagedCredentialsSecretOk() (*bool, bool)`

GetMcManagedCredentialsSecretOk returns a tuple with the McManagedCredentialsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcManagedCredentialsSecret

`func (o *ConnectionTypeOut) SetMcManagedCredentialsSecret(v bool)`

SetMcManagedCredentialsSecret sets McManagedCredentialsSecret field to given value.


### SetMcManagedCredentialsSecretNil

`func (o *ConnectionTypeOut) SetMcManagedCredentialsSecretNil(b bool)`

 SetMcManagedCredentialsSecretNil sets the value for McManagedCredentialsSecret to be an explicit nil

### UnsetMcManagedCredentialsSecret
`func (o *ConnectionTypeOut) UnsetMcManagedCredentialsSecret()`

UnsetMcManagedCredentialsSecret ensures that no value is present for McManagedCredentialsSecret, not even an explicit nil
### GetSelfHostedCredentialsStorages

`func (o *ConnectionTypeOut) GetSelfHostedCredentialsStorages() []CredentialsStorageType`

GetSelfHostedCredentialsStorages returns the SelfHostedCredentialsStorages field if non-nil, zero value otherwise.

### GetSelfHostedCredentialsStoragesOk

`func (o *ConnectionTypeOut) GetSelfHostedCredentialsStoragesOk() (*[]CredentialsStorageType, bool)`

GetSelfHostedCredentialsStoragesOk returns a tuple with the SelfHostedCredentialsStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfHostedCredentialsStorages

`func (o *ConnectionTypeOut) SetSelfHostedCredentialsStorages(v []CredentialsStorageType)`

SetSelfHostedCredentialsStorages sets SelfHostedCredentialsStorages field to given value.


### GetRequiresDeployment

`func (o *ConnectionTypeOut) GetRequiresDeployment() bool`

GetRequiresDeployment returns the RequiresDeployment field if non-nil, zero value otherwise.

### GetRequiresDeploymentOk

`func (o *ConnectionTypeOut) GetRequiresDeploymentOk() (*bool, bool)`

GetRequiresDeploymentOk returns a tuple with the RequiresDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresDeployment

`func (o *ConnectionTypeOut) SetRequiresDeployment(v bool)`

SetRequiresDeployment sets RequiresDeployment field to given value.


### GetRequiresOneOfConnectionTypes

`func (o *ConnectionTypeOut) GetRequiresOneOfConnectionTypes() []string`

GetRequiresOneOfConnectionTypes returns the RequiresOneOfConnectionTypes field if non-nil, zero value otherwise.

### GetRequiresOneOfConnectionTypesOk

`func (o *ConnectionTypeOut) GetRequiresOneOfConnectionTypesOk() (*[]string, bool)`

GetRequiresOneOfConnectionTypesOk returns a tuple with the RequiresOneOfConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresOneOfConnectionTypes

`func (o *ConnectionTypeOut) SetRequiresOneOfConnectionTypes(v []string)`

SetRequiresOneOfConnectionTypes sets RequiresOneOfConnectionTypes field to given value.


### GetJobTypes

`func (o *ConnectionTypeOut) GetJobTypes() []string`

GetJobTypes returns the JobTypes field if non-nil, zero value otherwise.

### GetJobTypesOk

`func (o *ConnectionTypeOut) GetJobTypesOk() (*[]string, bool)`

GetJobTypesOk returns a tuple with the JobTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypes

`func (o *ConnectionTypeOut) SetJobTypes(v []string)`

SetJobTypes sets JobTypes field to given value.


### GetIsCustom

`func (o *ConnectionTypeOut) GetIsCustom() bool`

GetIsCustom returns the IsCustom field if non-nil, zero value otherwise.

### GetIsCustomOk

`func (o *ConnectionTypeOut) GetIsCustomOk() (*bool, bool)`

GetIsCustomOk returns a tuple with the IsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustom

`func (o *ConnectionTypeOut) SetIsCustom(v bool)`

SetIsCustom sets IsCustom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


