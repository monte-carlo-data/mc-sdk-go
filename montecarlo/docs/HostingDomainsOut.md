# HostingDomainsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UiUrl** | **string** | URL of the Monte Carlo web app. | 
**IntegrationsUrl** | **string** | Base URL of the integrations gateway, which receives webhooks and other inbound requests from the tools you connect. | 
**GraphqlApiUrl** | **string** | URL of the GraphQL API, for API keys. | 
**GraphqlOauthApiUrl** | **string** | URL of the GraphQL API, for OAuth clients. | 
**M2mTokenUrl** | **NullableString** | URL an OAuth client requests an access token from, with the client-credentials grant. Null when the instance does not support OAuth clients. | 
**M2mInstanceScope** | **NullableString** | Scope an OAuth client requests for its token to reach this instance. Null when the instance does not support OAuth clients. | 
**DbtCloudWebhookUrl** | **string** | URL to give dbt Cloud when you create the webhook that reports job runs to Monte Carlo. | 

## Methods

### NewHostingDomainsOut

`func NewHostingDomainsOut(uiUrl string, integrationsUrl string, graphqlApiUrl string, graphqlOauthApiUrl string, m2mTokenUrl NullableString, m2mInstanceScope NullableString, dbtCloudWebhookUrl string, ) *HostingDomainsOut`

NewHostingDomainsOut instantiates a new HostingDomainsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostingDomainsOutWithDefaults

`func NewHostingDomainsOutWithDefaults() *HostingDomainsOut`

NewHostingDomainsOutWithDefaults instantiates a new HostingDomainsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUiUrl

`func (o *HostingDomainsOut) GetUiUrl() string`

GetUiUrl returns the UiUrl field if non-nil, zero value otherwise.

### GetUiUrlOk

`func (o *HostingDomainsOut) GetUiUrlOk() (*string, bool)`

GetUiUrlOk returns a tuple with the UiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiUrl

`func (o *HostingDomainsOut) SetUiUrl(v string)`

SetUiUrl sets UiUrl field to given value.


### GetIntegrationsUrl

`func (o *HostingDomainsOut) GetIntegrationsUrl() string`

GetIntegrationsUrl returns the IntegrationsUrl field if non-nil, zero value otherwise.

### GetIntegrationsUrlOk

`func (o *HostingDomainsOut) GetIntegrationsUrlOk() (*string, bool)`

GetIntegrationsUrlOk returns a tuple with the IntegrationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationsUrl

`func (o *HostingDomainsOut) SetIntegrationsUrl(v string)`

SetIntegrationsUrl sets IntegrationsUrl field to given value.


### GetGraphqlApiUrl

`func (o *HostingDomainsOut) GetGraphqlApiUrl() string`

GetGraphqlApiUrl returns the GraphqlApiUrl field if non-nil, zero value otherwise.

### GetGraphqlApiUrlOk

`func (o *HostingDomainsOut) GetGraphqlApiUrlOk() (*string, bool)`

GetGraphqlApiUrlOk returns a tuple with the GraphqlApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphqlApiUrl

`func (o *HostingDomainsOut) SetGraphqlApiUrl(v string)`

SetGraphqlApiUrl sets GraphqlApiUrl field to given value.


### GetGraphqlOauthApiUrl

`func (o *HostingDomainsOut) GetGraphqlOauthApiUrl() string`

GetGraphqlOauthApiUrl returns the GraphqlOauthApiUrl field if non-nil, zero value otherwise.

### GetGraphqlOauthApiUrlOk

`func (o *HostingDomainsOut) GetGraphqlOauthApiUrlOk() (*string, bool)`

GetGraphqlOauthApiUrlOk returns a tuple with the GraphqlOauthApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphqlOauthApiUrl

`func (o *HostingDomainsOut) SetGraphqlOauthApiUrl(v string)`

SetGraphqlOauthApiUrl sets GraphqlOauthApiUrl field to given value.


### GetM2mTokenUrl

`func (o *HostingDomainsOut) GetM2mTokenUrl() string`

GetM2mTokenUrl returns the M2mTokenUrl field if non-nil, zero value otherwise.

### GetM2mTokenUrlOk

`func (o *HostingDomainsOut) GetM2mTokenUrlOk() (*string, bool)`

GetM2mTokenUrlOk returns a tuple with the M2mTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM2mTokenUrl

`func (o *HostingDomainsOut) SetM2mTokenUrl(v string)`

SetM2mTokenUrl sets M2mTokenUrl field to given value.


### SetM2mTokenUrlNil

`func (o *HostingDomainsOut) SetM2mTokenUrlNil(b bool)`

 SetM2mTokenUrlNil sets the value for M2mTokenUrl to be an explicit nil

### UnsetM2mTokenUrl
`func (o *HostingDomainsOut) UnsetM2mTokenUrl()`

UnsetM2mTokenUrl ensures that no value is present for M2mTokenUrl, not even an explicit nil
### GetM2mInstanceScope

`func (o *HostingDomainsOut) GetM2mInstanceScope() string`

GetM2mInstanceScope returns the M2mInstanceScope field if non-nil, zero value otherwise.

### GetM2mInstanceScopeOk

`func (o *HostingDomainsOut) GetM2mInstanceScopeOk() (*string, bool)`

GetM2mInstanceScopeOk returns a tuple with the M2mInstanceScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM2mInstanceScope

`func (o *HostingDomainsOut) SetM2mInstanceScope(v string)`

SetM2mInstanceScope sets M2mInstanceScope field to given value.


### SetM2mInstanceScopeNil

`func (o *HostingDomainsOut) SetM2mInstanceScopeNil(b bool)`

 SetM2mInstanceScopeNil sets the value for M2mInstanceScope to be an explicit nil

### UnsetM2mInstanceScope
`func (o *HostingDomainsOut) UnsetM2mInstanceScope()`

UnsetM2mInstanceScope ensures that no value is present for M2mInstanceScope, not even an explicit nil
### GetDbtCloudWebhookUrl

`func (o *HostingDomainsOut) GetDbtCloudWebhookUrl() string`

GetDbtCloudWebhookUrl returns the DbtCloudWebhookUrl field if non-nil, zero value otherwise.

### GetDbtCloudWebhookUrlOk

`func (o *HostingDomainsOut) GetDbtCloudWebhookUrlOk() (*string, bool)`

GetDbtCloudWebhookUrlOk returns a tuple with the DbtCloudWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbtCloudWebhookUrl

`func (o *HostingDomainsOut) SetDbtCloudWebhookUrl(v string)`

SetDbtCloudWebhookUrl sets DbtCloudWebhookUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


