# WebsiteAlertConfigWithMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertChannelIds** | **[]string** |  | 
**Created** | Pointer to **int64** |  | [optional] 
**CustomPayloadFields** | [**[]CustomPayloadField**](CustomPayloadField.md) |  | 
**Description** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Granularity** | **int32** |  | [default to 600000]
**Id** | **string** |  | 
**Name** | **string** |  | 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Rule** | [**WebsiteAlertRule**](WebsiteAlertRule.md) |  | 
**Severity** | Pointer to **int32** |  | [optional] 
**TagFilterExpression** | [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | 
**TagFilters** | Pointer to [**[]TagFilter**](TagFilter.md) |  | [optional] 
**Threshold** | [**Threshold**](Threshold.md) |  | 
**TimeThreshold** | [**WebsiteTimeThreshold**](WebsiteTimeThreshold.md) |  | 
**Triggering** | Pointer to **bool** |  | [optional] 
**WebsiteId** | **string** |  | 

## Methods

### NewWebsiteAlertConfigWithMetadata

`func NewWebsiteAlertConfigWithMetadata(alertChannelIds []string, customPayloadFields []CustomPayloadField, description string, granularity int32, id string, name string, rule WebsiteAlertRule, tagFilterExpression TagFilterExpressionElement, threshold Threshold, timeThreshold WebsiteTimeThreshold, websiteId string, ) *WebsiteAlertConfigWithMetadata`

NewWebsiteAlertConfigWithMetadata instantiates a new WebsiteAlertConfigWithMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteAlertConfigWithMetadataWithDefaults

`func NewWebsiteAlertConfigWithMetadataWithDefaults() *WebsiteAlertConfigWithMetadata`

NewWebsiteAlertConfigWithMetadataWithDefaults instantiates a new WebsiteAlertConfigWithMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertChannelIds

`func (o *WebsiteAlertConfigWithMetadata) GetAlertChannelIds() []string`

GetAlertChannelIds returns the AlertChannelIds field if non-nil, zero value otherwise.

### GetAlertChannelIdsOk

`func (o *WebsiteAlertConfigWithMetadata) GetAlertChannelIdsOk() (*[]string, bool)`

GetAlertChannelIdsOk returns a tuple with the AlertChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertChannelIds

`func (o *WebsiteAlertConfigWithMetadata) SetAlertChannelIds(v []string)`

SetAlertChannelIds sets AlertChannelIds field to given value.


### GetCreated

`func (o *WebsiteAlertConfigWithMetadata) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebsiteAlertConfigWithMetadata) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebsiteAlertConfigWithMetadata) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WebsiteAlertConfigWithMetadata) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCustomPayloadFields

`func (o *WebsiteAlertConfigWithMetadata) GetCustomPayloadFields() []CustomPayloadField`

GetCustomPayloadFields returns the CustomPayloadFields field if non-nil, zero value otherwise.

### GetCustomPayloadFieldsOk

`func (o *WebsiteAlertConfigWithMetadata) GetCustomPayloadFieldsOk() (*[]CustomPayloadField, bool)`

GetCustomPayloadFieldsOk returns a tuple with the CustomPayloadFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadFields

`func (o *WebsiteAlertConfigWithMetadata) SetCustomPayloadFields(v []CustomPayloadField)`

SetCustomPayloadFields sets CustomPayloadFields field to given value.


### GetDescription

`func (o *WebsiteAlertConfigWithMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebsiteAlertConfigWithMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebsiteAlertConfigWithMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *WebsiteAlertConfigWithMetadata) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebsiteAlertConfigWithMetadata) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebsiteAlertConfigWithMetadata) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebsiteAlertConfigWithMetadata) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGranularity

`func (o *WebsiteAlertConfigWithMetadata) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *WebsiteAlertConfigWithMetadata) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *WebsiteAlertConfigWithMetadata) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.


### GetId

`func (o *WebsiteAlertConfigWithMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebsiteAlertConfigWithMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebsiteAlertConfigWithMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WebsiteAlertConfigWithMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebsiteAlertConfigWithMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebsiteAlertConfigWithMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetReadOnly

`func (o *WebsiteAlertConfigWithMetadata) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *WebsiteAlertConfigWithMetadata) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *WebsiteAlertConfigWithMetadata) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *WebsiteAlertConfigWithMetadata) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRule

`func (o *WebsiteAlertConfigWithMetadata) GetRule() WebsiteAlertRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *WebsiteAlertConfigWithMetadata) GetRuleOk() (*WebsiteAlertRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *WebsiteAlertConfigWithMetadata) SetRule(v WebsiteAlertRule)`

SetRule sets Rule field to given value.


### GetSeverity

`func (o *WebsiteAlertConfigWithMetadata) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *WebsiteAlertConfigWithMetadata) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *WebsiteAlertConfigWithMetadata) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *WebsiteAlertConfigWithMetadata) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTagFilterExpression

`func (o *WebsiteAlertConfigWithMetadata) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *WebsiteAlertConfigWithMetadata) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *WebsiteAlertConfigWithMetadata) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.


### GetTagFilters

`func (o *WebsiteAlertConfigWithMetadata) GetTagFilters() []TagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *WebsiteAlertConfigWithMetadata) GetTagFiltersOk() (*[]TagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *WebsiteAlertConfigWithMetadata) SetTagFilters(v []TagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *WebsiteAlertConfigWithMetadata) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetThreshold

`func (o *WebsiteAlertConfigWithMetadata) GetThreshold() Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *WebsiteAlertConfigWithMetadata) GetThresholdOk() (*Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *WebsiteAlertConfigWithMetadata) SetThreshold(v Threshold)`

SetThreshold sets Threshold field to given value.


### GetTimeThreshold

`func (o *WebsiteAlertConfigWithMetadata) GetTimeThreshold() WebsiteTimeThreshold`

GetTimeThreshold returns the TimeThreshold field if non-nil, zero value otherwise.

### GetTimeThresholdOk

`func (o *WebsiteAlertConfigWithMetadata) GetTimeThresholdOk() (*WebsiteTimeThreshold, bool)`

GetTimeThresholdOk returns a tuple with the TimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeThreshold

`func (o *WebsiteAlertConfigWithMetadata) SetTimeThreshold(v WebsiteTimeThreshold)`

SetTimeThreshold sets TimeThreshold field to given value.


### GetTriggering

`func (o *WebsiteAlertConfigWithMetadata) GetTriggering() bool`

GetTriggering returns the Triggering field if non-nil, zero value otherwise.

### GetTriggeringOk

`func (o *WebsiteAlertConfigWithMetadata) GetTriggeringOk() (*bool, bool)`

GetTriggeringOk returns a tuple with the Triggering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggering

`func (o *WebsiteAlertConfigWithMetadata) SetTriggering(v bool)`

SetTriggering sets Triggering field to given value.

### HasTriggering

`func (o *WebsiteAlertConfigWithMetadata) HasTriggering() bool`

HasTriggering returns a boolean if a field has been set.

### GetWebsiteId

`func (o *WebsiteAlertConfigWithMetadata) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *WebsiteAlertConfigWithMetadata) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *WebsiteAlertConfigWithMetadata) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


