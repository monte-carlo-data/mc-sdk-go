# PrestoCredentialsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the credentials. | 
**ConnectionType** | **string** | The connection type the credentials are for, such as &#x60;snowflake&#x60;. Fixed once created. | 
**StorageType** | [**CredentialsStorageType**](CredentialsStorageType.md) | Where the secret lives. Fixed once created. | 
**CreatedTime** | **time.Time** | When the credentials were created. | 
**Host** | **string** | Hostname of the Presto or Trino coordinator. | 
**Port** | **int32** | Port the coordinator listens on. | 
**User** | **string** | User Monte Carlo signs in as. | 
**Catalog** | **NullableString** | Catalog Monte Carlo queries. Null unless set. | 
**Schema** | **NullableString** | Schema Monte Carlo queries by default. Null unless set. | 
**HttpScheme** | **NullableString** | Whether to reach the coordinator over &#x60;http&#x60; or &#x60;https&#x60;. Null means &#x60;http&#x60;. | 
**MetadataCatalogId** | **NullableString** | Name of the catalog in the metadata store. Null means it is &#x60;catalog&#x60;. | 
**SslSkipVerification** | **bool** | Connect over TLS without checking the server&#39;s certificate. | 

## Methods

### NewPrestoCredentialsOut

`func NewPrestoCredentialsOut(id string, connectionType string, storageType CredentialsStorageType, createdTime time.Time, host string, port int32, user string, catalog NullableString, schema NullableString, httpScheme NullableString, metadataCatalogId NullableString, sslSkipVerification bool, ) *PrestoCredentialsOut`

NewPrestoCredentialsOut instantiates a new PrestoCredentialsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrestoCredentialsOutWithDefaults

`func NewPrestoCredentialsOutWithDefaults() *PrestoCredentialsOut`

NewPrestoCredentialsOutWithDefaults instantiates a new PrestoCredentialsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrestoCredentialsOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrestoCredentialsOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrestoCredentialsOut) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionType

`func (o *PrestoCredentialsOut) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *PrestoCredentialsOut) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *PrestoCredentialsOut) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetStorageType

`func (o *PrestoCredentialsOut) GetStorageType() CredentialsStorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *PrestoCredentialsOut) GetStorageTypeOk() (*CredentialsStorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *PrestoCredentialsOut) SetStorageType(v CredentialsStorageType)`

SetStorageType sets StorageType field to given value.


### GetCreatedTime

`func (o *PrestoCredentialsOut) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PrestoCredentialsOut) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PrestoCredentialsOut) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetHost

`func (o *PrestoCredentialsOut) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PrestoCredentialsOut) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PrestoCredentialsOut) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *PrestoCredentialsOut) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PrestoCredentialsOut) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PrestoCredentialsOut) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *PrestoCredentialsOut) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PrestoCredentialsOut) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PrestoCredentialsOut) SetUser(v string)`

SetUser sets User field to given value.


### GetCatalog

`func (o *PrestoCredentialsOut) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *PrestoCredentialsOut) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *PrestoCredentialsOut) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.


### SetCatalogNil

`func (o *PrestoCredentialsOut) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *PrestoCredentialsOut) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
### GetSchema

`func (o *PrestoCredentialsOut) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *PrestoCredentialsOut) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *PrestoCredentialsOut) SetSchema(v string)`

SetSchema sets Schema field to given value.


### SetSchemaNil

`func (o *PrestoCredentialsOut) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *PrestoCredentialsOut) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetHttpScheme

`func (o *PrestoCredentialsOut) GetHttpScheme() string`

GetHttpScheme returns the HttpScheme field if non-nil, zero value otherwise.

### GetHttpSchemeOk

`func (o *PrestoCredentialsOut) GetHttpSchemeOk() (*string, bool)`

GetHttpSchemeOk returns a tuple with the HttpScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpScheme

`func (o *PrestoCredentialsOut) SetHttpScheme(v string)`

SetHttpScheme sets HttpScheme field to given value.


### SetHttpSchemeNil

`func (o *PrestoCredentialsOut) SetHttpSchemeNil(b bool)`

 SetHttpSchemeNil sets the value for HttpScheme to be an explicit nil

### UnsetHttpScheme
`func (o *PrestoCredentialsOut) UnsetHttpScheme()`

UnsetHttpScheme ensures that no value is present for HttpScheme, not even an explicit nil
### GetMetadataCatalogId

`func (o *PrestoCredentialsOut) GetMetadataCatalogId() string`

GetMetadataCatalogId returns the MetadataCatalogId field if non-nil, zero value otherwise.

### GetMetadataCatalogIdOk

`func (o *PrestoCredentialsOut) GetMetadataCatalogIdOk() (*string, bool)`

GetMetadataCatalogIdOk returns a tuple with the MetadataCatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCatalogId

`func (o *PrestoCredentialsOut) SetMetadataCatalogId(v string)`

SetMetadataCatalogId sets MetadataCatalogId field to given value.


### SetMetadataCatalogIdNil

`func (o *PrestoCredentialsOut) SetMetadataCatalogIdNil(b bool)`

 SetMetadataCatalogIdNil sets the value for MetadataCatalogId to be an explicit nil

### UnsetMetadataCatalogId
`func (o *PrestoCredentialsOut) UnsetMetadataCatalogId()`

UnsetMetadataCatalogId ensures that no value is present for MetadataCatalogId, not even an explicit nil
### GetSslSkipVerification

`func (o *PrestoCredentialsOut) GetSslSkipVerification() bool`

GetSslSkipVerification returns the SslSkipVerification field if non-nil, zero value otherwise.

### GetSslSkipVerificationOk

`func (o *PrestoCredentialsOut) GetSslSkipVerificationOk() (*bool, bool)`

GetSslSkipVerificationOk returns a tuple with the SslSkipVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslSkipVerification

`func (o *PrestoCredentialsOut) SetSslSkipVerification(v bool)`

SetSslSkipVerification sets SslSkipVerification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


