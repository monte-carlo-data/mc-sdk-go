# PrestoCredentialsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname of the Presto or Trino coordinator. | 
**Port** | **int32** | Port the coordinator listens on. | 
**User** | **string** | User Monte Carlo signs in as. | 
**Password** | Pointer to **NullableString** | Password of &#x60;user&#x60;. Leave it out for a cluster that takes none. Stored by Monte Carlo and never returned. | [optional] 
**Catalog** | Pointer to **NullableString** | Catalog Monte Carlo queries. | [optional] 
**Schema** | Pointer to **NullableString** | Schema Monte Carlo queries by default. | [optional] 
**HttpScheme** | Pointer to **NullableString** | Whether to reach the coordinator over &#x60;http&#x60; or &#x60;https&#x60;. | [optional] 
**MetadataCatalogId** | Pointer to **NullableString** | Name of the catalog in the metadata store, when it differs from &#x60;catalog&#x60;. Requires &#x60;catalog&#x60;, which may then hold only lowercase letters, digits, &#x60;_&#x60; and &#x60;-&#x60;. | [optional] 
**SslSkipVerification** | Pointer to **NullableBool** | Connect over TLS without checking the server&#39;s certificate. | [optional] 

## Methods

### NewPrestoCredentialsIn

`func NewPrestoCredentialsIn(host string, port int32, user string, ) *PrestoCredentialsIn`

NewPrestoCredentialsIn instantiates a new PrestoCredentialsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrestoCredentialsInWithDefaults

`func NewPrestoCredentialsInWithDefaults() *PrestoCredentialsIn`

NewPrestoCredentialsInWithDefaults instantiates a new PrestoCredentialsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *PrestoCredentialsIn) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PrestoCredentialsIn) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PrestoCredentialsIn) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *PrestoCredentialsIn) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PrestoCredentialsIn) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PrestoCredentialsIn) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *PrestoCredentialsIn) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PrestoCredentialsIn) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PrestoCredentialsIn) SetUser(v string)`

SetUser sets User field to given value.


### GetPassword

`func (o *PrestoCredentialsIn) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PrestoCredentialsIn) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PrestoCredentialsIn) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PrestoCredentialsIn) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PrestoCredentialsIn) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PrestoCredentialsIn) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetCatalog

`func (o *PrestoCredentialsIn) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *PrestoCredentialsIn) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *PrestoCredentialsIn) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *PrestoCredentialsIn) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *PrestoCredentialsIn) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *PrestoCredentialsIn) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
### GetSchema

`func (o *PrestoCredentialsIn) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *PrestoCredentialsIn) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *PrestoCredentialsIn) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *PrestoCredentialsIn) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *PrestoCredentialsIn) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *PrestoCredentialsIn) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetHttpScheme

`func (o *PrestoCredentialsIn) GetHttpScheme() string`

GetHttpScheme returns the HttpScheme field if non-nil, zero value otherwise.

### GetHttpSchemeOk

`func (o *PrestoCredentialsIn) GetHttpSchemeOk() (*string, bool)`

GetHttpSchemeOk returns a tuple with the HttpScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpScheme

`func (o *PrestoCredentialsIn) SetHttpScheme(v string)`

SetHttpScheme sets HttpScheme field to given value.

### HasHttpScheme

`func (o *PrestoCredentialsIn) HasHttpScheme() bool`

HasHttpScheme returns a boolean if a field has been set.

### SetHttpSchemeNil

`func (o *PrestoCredentialsIn) SetHttpSchemeNil(b bool)`

 SetHttpSchemeNil sets the value for HttpScheme to be an explicit nil

### UnsetHttpScheme
`func (o *PrestoCredentialsIn) UnsetHttpScheme()`

UnsetHttpScheme ensures that no value is present for HttpScheme, not even an explicit nil
### GetMetadataCatalogId

`func (o *PrestoCredentialsIn) GetMetadataCatalogId() string`

GetMetadataCatalogId returns the MetadataCatalogId field if non-nil, zero value otherwise.

### GetMetadataCatalogIdOk

`func (o *PrestoCredentialsIn) GetMetadataCatalogIdOk() (*string, bool)`

GetMetadataCatalogIdOk returns a tuple with the MetadataCatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCatalogId

`func (o *PrestoCredentialsIn) SetMetadataCatalogId(v string)`

SetMetadataCatalogId sets MetadataCatalogId field to given value.

### HasMetadataCatalogId

`func (o *PrestoCredentialsIn) HasMetadataCatalogId() bool`

HasMetadataCatalogId returns a boolean if a field has been set.

### SetMetadataCatalogIdNil

`func (o *PrestoCredentialsIn) SetMetadataCatalogIdNil(b bool)`

 SetMetadataCatalogIdNil sets the value for MetadataCatalogId to be an explicit nil

### UnsetMetadataCatalogId
`func (o *PrestoCredentialsIn) UnsetMetadataCatalogId()`

UnsetMetadataCatalogId ensures that no value is present for MetadataCatalogId, not even an explicit nil
### GetSslSkipVerification

`func (o *PrestoCredentialsIn) GetSslSkipVerification() bool`

GetSslSkipVerification returns the SslSkipVerification field if non-nil, zero value otherwise.

### GetSslSkipVerificationOk

`func (o *PrestoCredentialsIn) GetSslSkipVerificationOk() (*bool, bool)`

GetSslSkipVerificationOk returns a tuple with the SslSkipVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSkipVerification

`func (o *PrestoCredentialsIn) SetSslSkipVerification(v bool)`

SetSslSkipVerification sets SslSkipVerification field to given value.

### HasSslSkipVerification

`func (o *PrestoCredentialsIn) HasSslSkipVerification() bool`

HasSslSkipVerification returns a boolean if a field has been set.

### SetSslSkipVerificationNil

`func (o *PrestoCredentialsIn) SetSslSkipVerificationNil(b bool)`

 SetSslSkipVerificationNil sets the value for SslSkipVerification to be an explicit nil

### UnsetSslSkipVerification
`func (o *PrestoCredentialsIn) UnsetSslSkipVerification()`

UnsetSslSkipVerification ensures that no value is present for SslSkipVerification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


