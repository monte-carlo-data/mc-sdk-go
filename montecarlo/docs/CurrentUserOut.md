# CurrentUserOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountFrozen** | **bool** | Whether the account is paused. While it is, this endpoint still answers but every other one returns 403 with the code &#x60;account_frozen&#x60;. | 
**AccountId** | **string** | Unique identifier of the caller&#39;s account. | 
**AccountName** | Pointer to **NullableString** | Display name of the account. Null when the account has no name. | [optional] 
**AuthGroups** | Pointer to **[]string** | Names of the authorization groups this user belongs to. They determine what the user is permitted to do. | [optional] 
**Email** | **string** | Email address of the user. An identity with no mailbox of its own, such as an AI agent, carries a display label here instead. | 
**FirstName** | Pointer to **NullableString** | Given name of the user. Null when it is not set. | [optional] 
**IdentityType** | [**IdentityType**](IdentityType.md) | What kind of identity this is. It does not decide what the identity may do — that comes from the authorization groups it belongs to. | 
**LastName** | Pointer to **NullableString** | Family name of the user. Null when it is not set. | [optional] 
**UserId** | **string** | Unique identifier of the user. | 

## Methods

### NewCurrentUserOut

`func NewCurrentUserOut(accountFrozen bool, accountId string, email string, identityType IdentityType, userId string, ) *CurrentUserOut`

NewCurrentUserOut instantiates a new CurrentUserOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentUserOutWithDefaults

`func NewCurrentUserOutWithDefaults() *CurrentUserOut`

NewCurrentUserOutWithDefaults instantiates a new CurrentUserOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountFrozen

`func (o *CurrentUserOut) GetAccountFrozen() bool`

GetAccountFrozen returns the AccountFrozen field if non-nil, zero value otherwise.

### GetAccountFrozenOk

`func (o *CurrentUserOut) GetAccountFrozenOk() (*bool, bool)`

GetAccountFrozenOk returns a tuple with the AccountFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFrozen

`func (o *CurrentUserOut) SetAccountFrozen(v bool)`

SetAccountFrozen sets AccountFrozen field to given value.


### GetAccountId

`func (o *CurrentUserOut) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CurrentUserOut) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CurrentUserOut) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountName

`func (o *CurrentUserOut) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *CurrentUserOut) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *CurrentUserOut) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *CurrentUserOut) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *CurrentUserOut) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *CurrentUserOut) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetAuthGroups

`func (o *CurrentUserOut) GetAuthGroups() []string`

GetAuthGroups returns the AuthGroups field if non-nil, zero value otherwise.

### GetAuthGroupsOk

`func (o *CurrentUserOut) GetAuthGroupsOk() (*[]string, bool)`

GetAuthGroupsOk returns a tuple with the AuthGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthGroups

`func (o *CurrentUserOut) SetAuthGroups(v []string)`

SetAuthGroups sets AuthGroups field to given value.

### HasAuthGroups

`func (o *CurrentUserOut) HasAuthGroups() bool`

HasAuthGroups returns a boolean if a field has been set.

### GetEmail

`func (o *CurrentUserOut) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CurrentUserOut) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CurrentUserOut) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *CurrentUserOut) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CurrentUserOut) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CurrentUserOut) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CurrentUserOut) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *CurrentUserOut) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *CurrentUserOut) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetIdentityType

`func (o *CurrentUserOut) GetIdentityType() IdentityType`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *CurrentUserOut) GetIdentityTypeOk() (*IdentityType, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *CurrentUserOut) SetIdentityType(v IdentityType)`

SetIdentityType sets IdentityType field to given value.


### GetLastName

`func (o *CurrentUserOut) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CurrentUserOut) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CurrentUserOut) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CurrentUserOut) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *CurrentUserOut) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *CurrentUserOut) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetUserId

`func (o *CurrentUserOut) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CurrentUserOut) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CurrentUserOut) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


