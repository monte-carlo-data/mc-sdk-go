# PrestoCredentialsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | Hostname of the Presto or Trino coordinator. Changing it requires sending &#x60;password&#x60; again when one is stored. | [optional] 
**Port** | Pointer to **NullableInt32** | Port the coordinator listens on. | [optional] 
**User** | Pointer to **NullableString** | User Monte Carlo signs in as. | [optional] 
**Password** | Pointer to **NullableString** | New password of &#x60;user&#x60;. | [optional] 
**Catalog** | Pointer to **NullableString** | Catalog Monte Carlo queries. | [optional] 
**Schema** | Pointer to **NullableString** | Schema Monte Carlo queries by default. | [optional] 
**HttpScheme** | Pointer to **NullableString** | Whether to reach the coordinator over &#x60;http&#x60; or &#x60;https&#x60;. Changing it requires sending &#x60;password&#x60; again when one is stored. | [optional] 
**MetadataCatalogId** | Pointer to **NullableString** | Name of the catalog in the metadata store, when it differs from &#x60;catalog&#x60;. Requires &#x60;catalog&#x60;, which may then hold only lowercase letters, digits, &#x60;_&#x60; and &#x60;-&#x60;. | [optional] 
**SslSkipVerification** | Pointer to **NullableBool** | Connect over TLS without checking the server&#39;s certificate. Changing it requires sending &#x60;password&#x60; again when one is stored. | [optional] 

## Methods

### NewPrestoCredentialsPatch

`func NewPrestoCredentialsPatch() *PrestoCredentialsPatch`

NewPrestoCredentialsPatch instantiates a new PrestoCredentialsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrestoCredentialsPatchWithDefaults

`func NewPrestoCredentialsPatchWithDefaults() *PrestoCredentialsPatch`

NewPrestoCredentialsPatchWithDefaults instantiates a new PrestoCredentialsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *PrestoCredentialsPatch) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PrestoCredentialsPatch) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PrestoCredentialsPatch) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PrestoCredentialsPatch) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *PrestoCredentialsPatch) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *PrestoCredentialsPatch) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *PrestoCredentialsPatch) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PrestoCredentialsPatch) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PrestoCredentialsPatch) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PrestoCredentialsPatch) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *PrestoCredentialsPatch) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *PrestoCredentialsPatch) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUser

`func (o *PrestoCredentialsPatch) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PrestoCredentialsPatch) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PrestoCredentialsPatch) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *PrestoCredentialsPatch) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PrestoCredentialsPatch) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PrestoCredentialsPatch) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *PrestoCredentialsPatch) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PrestoCredentialsPatch) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PrestoCredentialsPatch) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PrestoCredentialsPatch) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PrestoCredentialsPatch) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PrestoCredentialsPatch) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetCatalog

`func (o *PrestoCredentialsPatch) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *PrestoCredentialsPatch) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *PrestoCredentialsPatch) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *PrestoCredentialsPatch) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *PrestoCredentialsPatch) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *PrestoCredentialsPatch) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
### GetSchema

`func (o *PrestoCredentialsPatch) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *PrestoCredentialsPatch) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *PrestoCredentialsPatch) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *PrestoCredentialsPatch) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *PrestoCredentialsPatch) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *PrestoCredentialsPatch) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetHttpScheme

`func (o *PrestoCredentialsPatch) GetHttpScheme() string`

GetHttpScheme returns the HttpScheme field if non-nil, zero value otherwise.

### GetHttpSchemeOk

`func (o *PrestoCredentialsPatch) GetHttpSchemeOk() (*string, bool)`

GetHttpSchemeOk returns a tuple with the HttpScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpScheme

`func (o *PrestoCredentialsPatch) SetHttpScheme(v string)`

SetHttpScheme sets HttpScheme field to given value.

### HasHttpScheme

`func (o *PrestoCredentialsPatch) HasHttpScheme() bool`

HasHttpScheme returns a boolean if a field has been set.

### SetHttpSchemeNil

`func (o *PrestoCredentialsPatch) SetHttpSchemeNil(b bool)`

 SetHttpSchemeNil sets the value for HttpScheme to be an explicit nil

### UnsetHttpScheme
`func (o *PrestoCredentialsPatch) UnsetHttpScheme()`

UnsetHttpScheme ensures that no value is present for HttpScheme, not even an explicit nil
### GetMetadataCatalogId

`func (o *PrestoCredentialsPatch) GetMetadataCatalogId() string`

GetMetadataCatalogId returns the MetadataCatalogId field if non-nil, zero value otherwise.

### GetMetadataCatalogIdOk

`func (o *PrestoCredentialsPatch) GetMetadataCatalogIdOk() (*string, bool)`

GetMetadataCatalogIdOk returns a tuple with the MetadataCatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCatalogId

`func (o *PrestoCredentialsPatch) SetMetadataCatalogId(v string)`

SetMetadataCatalogId sets MetadataCatalogId field to given value.

### HasMetadataCatalogId

`func (o *PrestoCredentialsPatch) HasMetadataCatalogId() bool`

HasMetadataCatalogId returns a boolean if a field has been set.

### SetMetadataCatalogIdNil

`func (o *PrestoCredentialsPatch) SetMetadataCatalogIdNil(b bool)`

 SetMetadataCatalogIdNil sets the value for MetadataCatalogId to be an explicit nil

### UnsetMetadataCatalogId
`func (o *PrestoCredentialsPatch) UnsetMetadataCatalogId()`

UnsetMetadataCatalogId ensures that no value is present for MetadataCatalogId, not even an explicit nil
### GetSslSkipVerification

`func (o *PrestoCredentialsPatch) GetSslSkipVerification() bool`

GetSslSkipVerification returns the SslSkipVerification field if non-nil, zero value otherwise.

### GetSslSkipVerificationOk

`func (o *PrestoCredentialsPatch) GetSslSkipVerificationOk() (*bool, bool)`

GetSslSkipVerificationOk returns a tuple with the SslSkipVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSkipVerification

`func (o *PrestoCredentialsPatch) SetSslSkipVerification(v bool)`

SetSslSkipVerification sets SslSkipVerification field to given value.

### HasSslSkipVerification

`func (o *PrestoCredentialsPatch) HasSslSkipVerification() bool`

HasSslSkipVerification returns a boolean if a field has been set.

### SetSslSkipVerificationNil

`func (o *PrestoCredentialsPatch) SetSslSkipVerificationNil(b bool)`

 SetSslSkipVerificationNil sets the value for SslSkipVerification to be an explicit nil

### UnsetSslSkipVerification
`func (o *PrestoCredentialsPatch) UnsetSslSkipVerification()`

UnsetSslSkipVerification ensures that no value is present for SslSkipVerification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


