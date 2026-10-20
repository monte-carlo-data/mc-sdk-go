# HostingNetworkOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionIpAddresses** | **[]string** | Addresses Monte Carlo collects from, such as when it reaches your warehouse, collection agent or data store over the internet. | 
**SaasIpAddresses** | **[]string** | Addresses the rest of the Monte Carlo platform connects from. | 
**WebhookIpAddresses** | **[]string** | Addresses Monte Carlo sends webhook notifications from. | 
**DataPlatformIpAddresses** | **[]string** | Addresses Monte Carlo&#39;s data processing services connect from. | 

## Methods

### NewHostingNetworkOut

`func NewHostingNetworkOut(collectionIpAddresses []string, saasIpAddresses []string, webhookIpAddresses []string, dataPlatformIpAddresses []string, ) *HostingNetworkOut`

NewHostingNetworkOut instantiates a new HostingNetworkOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostingNetworkOutWithDefaults

`func NewHostingNetworkOutWithDefaults() *HostingNetworkOut`

NewHostingNetworkOutWithDefaults instantiates a new HostingNetworkOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionIpAddresses

`func (o *HostingNetworkOut) GetCollectionIpAddresses() []string`

GetCollectionIpAddresses returns the CollectionIpAddresses field if non-nil, zero value otherwise.

### GetCollectionIpAddressesOk

`func (o *HostingNetworkOut) GetCollectionIpAddressesOk() (*[]string, bool)`

GetCollectionIpAddressesOk returns a tuple with the CollectionIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionIpAddresses

`func (o *HostingNetworkOut) SetCollectionIpAddresses(v []string)`

SetCollectionIpAddresses sets CollectionIpAddresses field to given value.


### GetSaasIpAddresses

`func (o *HostingNetworkOut) GetSaasIpAddresses() []string`

GetSaasIpAddresses returns the SaasIpAddresses field if non-nil, zero value otherwise.

### GetSaasIpAddressesOk

`func (o *HostingNetworkOut) GetSaasIpAddressesOk() (*[]string, bool)`

GetSaasIpAddressesOk returns a tuple with the SaasIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaasIpAddresses

`func (o *HostingNetworkOut) SetSaasIpAddresses(v []string)`

SetSaasIpAddresses sets SaasIpAddresses field to given value.


### GetWebhookIpAddresses

`func (o *HostingNetworkOut) GetWebhookIpAddresses() []string`

GetWebhookIpAddresses returns the WebhookIpAddresses field if non-nil, zero value otherwise.

### GetWebhookIpAddressesOk

`func (o *HostingNetworkOut) GetWebhookIpAddressesOk() (*[]string, bool)`

GetWebhookIpAddressesOk returns a tuple with the WebhookIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookIpAddresses

`func (o *HostingNetworkOut) SetWebhookIpAddresses(v []string)`

SetWebhookIpAddresses sets WebhookIpAddresses field to given value.


### GetDataPlatformIpAddresses

`func (o *HostingNetworkOut) GetDataPlatformIpAddresses() []string`

GetDataPlatformIpAddresses returns the DataPlatformIpAddresses field if non-nil, zero value otherwise.

### GetDataPlatformIpAddressesOk

`func (o *HostingNetworkOut) GetDataPlatformIpAddressesOk() (*[]string, bool)`

GetDataPlatformIpAddressesOk returns a tuple with the DataPlatformIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPlatformIpAddresses

`func (o *HostingNetworkOut) SetDataPlatformIpAddresses(v []string)`

SetDataPlatformIpAddresses sets DataPlatformIpAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


