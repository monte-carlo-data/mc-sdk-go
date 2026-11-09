# DocPageOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | The page&#39;s id. Ask for it with &#x60;slug&#x60; to get its body. | 
**Title** | **string** | The page title. | 
**Summary** | Pointer to **NullableString** | One line on what the page covers. | [optional] 
**Parent** | Pointer to **NullableString** | The slug of the page it sits under. | [optional] 
**Body** | Pointer to **NullableString** | The page as Markdown. Set only when one page was asked for. &#x60;doc:&lt;slug&gt;&#x60; links name other pages. | [optional] 
**Prompt** | Pointer to **NullableString** | A dashboard example&#39;s prompt: send it to an agent to build that dashboard. | [optional] 

## Methods

### NewDocPageOut

`func NewDocPageOut(slug string, title string, ) *DocPageOut`

NewDocPageOut instantiates a new DocPageOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocPageOutWithDefaults

`func NewDocPageOutWithDefaults() *DocPageOut`

NewDocPageOutWithDefaults instantiates a new DocPageOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *DocPageOut) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DocPageOut) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DocPageOut) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *DocPageOut) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocPageOut) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocPageOut) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSummary

`func (o *DocPageOut) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DocPageOut) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DocPageOut) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *DocPageOut) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *DocPageOut) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *DocPageOut) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetParent

`func (o *DocPageOut) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *DocPageOut) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *DocPageOut) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *DocPageOut) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *DocPageOut) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *DocPageOut) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBody

`func (o *DocPageOut) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *DocPageOut) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *DocPageOut) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *DocPageOut) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *DocPageOut) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *DocPageOut) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetPrompt

`func (o *DocPageOut) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *DocPageOut) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *DocPageOut) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *DocPageOut) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### SetPromptNil

`func (o *DocPageOut) SetPromptNil(b bool)`

 SetPromptNil sets the value for Prompt to be an explicit nil

### UnsetPrompt
`func (o *DocPageOut) UnsetPrompt()`

UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


