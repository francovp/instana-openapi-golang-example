# WebsiteAlertConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertChannelIds** | **[]string** |  | 
**CustomPayloadFields** | [**[]CustomPayloadField**](CustomPayloadField.md) |  | 
**Description** | **string** |  | 
**Granularity** | **int32** |  | [default to 600000]
**Name** | **string** |  | 
**Rule** | [**WebsiteAlertRule**](WebsiteAlertRule.md) |  | 
**Severity** | Pointer to **int32** |  | [optional] 
**TagFilterExpression** | [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | 
**TagFilters** | Pointer to [**[]TagFilter**](TagFilter.md) |  | [optional] 
**Threshold** | [**Threshold**](Threshold.md) |  | 
**TimeThreshold** | [**WebsiteTimeThreshold**](WebsiteTimeThreshold.md) |  | 
**Triggering** | Pointer to **bool** |  | [optional] 
**WebsiteId** | **string** |  | 

## Methods

### NewWebsiteAlertConfig

`func NewWebsiteAlertConfig(alertChannelIds []string, customPayloadFields []CustomPayloadField, description string, granularity int32, name string, rule WebsiteAlertRule, tagFilterExpression TagFilterExpressionElement, threshold Threshold, timeThreshold WebsiteTimeThreshold, websiteId string, ) *WebsiteAlertConfig`

NewWebsiteAlertConfig instantiates a new WebsiteAlertConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteAlertConfigWithDefaults

`func NewWebsiteAlertConfigWithDefaults() *WebsiteAlertConfig`

NewWebsiteAlertConfigWithDefaults instantiates a new WebsiteAlertConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertChannelIds

`func (o *WebsiteAlertConfig) GetAlertChannelIds() []string`

GetAlertChannelIds returns the AlertChannelIds field if non-nil, zero value otherwise.

### GetAlertChannelIdsOk

`func (o *WebsiteAlertConfig) GetAlertChannelIdsOk() (*[]string, bool)`

GetAlertChannelIdsOk returns a tuple with the AlertChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertChannelIds

`func (o *WebsiteAlertConfig) SetAlertChannelIds(v []string)`

SetAlertChannelIds sets AlertChannelIds field to given value.


### GetCustomPayloadFields

`func (o *WebsiteAlertConfig) GetCustomPayloadFields() []CustomPayloadField`

GetCustomPayloadFields returns the CustomPayloadFields field if non-nil, zero value otherwise.

### GetCustomPayloadFieldsOk

`func (o *WebsiteAlertConfig) GetCustomPayloadFieldsOk() (*[]CustomPayloadField, bool)`

GetCustomPayloadFieldsOk returns a tuple with the CustomPayloadFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadFields

`func (o *WebsiteAlertConfig) SetCustomPayloadFields(v []CustomPayloadField)`

SetCustomPayloadFields sets CustomPayloadFields field to given value.


### GetDescription

`func (o *WebsiteAlertConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebsiteAlertConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebsiteAlertConfig) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGranularity

`func (o *WebsiteAlertConfig) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *WebsiteAlertConfig) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *WebsiteAlertConfig) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.


### GetName

`func (o *WebsiteAlertConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebsiteAlertConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebsiteAlertConfig) SetName(v string)`

SetName sets Name field to given value.


### GetRule

`func (o *WebsiteAlertConfig) GetRule() WebsiteAlertRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *WebsiteAlertConfig) GetRuleOk() (*WebsiteAlertRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *WebsiteAlertConfig) SetRule(v WebsiteAlertRule)`

SetRule sets Rule field to given value.


### GetSeverity

`func (o *WebsiteAlertConfig) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *WebsiteAlertConfig) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *WebsiteAlertConfig) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *WebsiteAlertConfig) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTagFilterExpression

`func (o *WebsiteAlertConfig) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *WebsiteAlertConfig) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *WebsiteAlertConfig) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.


### GetTagFilters

`func (o *WebsiteAlertConfig) GetTagFilters() []TagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *WebsiteAlertConfig) GetTagFiltersOk() (*[]TagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *WebsiteAlertConfig) SetTagFilters(v []TagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *WebsiteAlertConfig) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetThreshold

`func (o *WebsiteAlertConfig) GetThreshold() Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *WebsiteAlertConfig) GetThresholdOk() (*Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *WebsiteAlertConfig) SetThreshold(v Threshold)`

SetThreshold sets Threshold field to given value.


### GetTimeThreshold

`func (o *WebsiteAlertConfig) GetTimeThreshold() WebsiteTimeThreshold`

GetTimeThreshold returns the TimeThreshold field if non-nil, zero value otherwise.

### GetTimeThresholdOk

`func (o *WebsiteAlertConfig) GetTimeThresholdOk() (*WebsiteTimeThreshold, bool)`

GetTimeThresholdOk returns a tuple with the TimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeThreshold

`func (o *WebsiteAlertConfig) SetTimeThreshold(v WebsiteTimeThreshold)`

SetTimeThreshold sets TimeThreshold field to given value.


### GetTriggering

`func (o *WebsiteAlertConfig) GetTriggering() bool`

GetTriggering returns the Triggering field if non-nil, zero value otherwise.

### GetTriggeringOk

`func (o *WebsiteAlertConfig) GetTriggeringOk() (*bool, bool)`

GetTriggeringOk returns a tuple with the Triggering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggering

`func (o *WebsiteAlertConfig) SetTriggering(v bool)`

SetTriggering sets Triggering field to given value.

### HasTriggering

`func (o *WebsiteAlertConfig) HasTriggering() bool`

HasTriggering returns a boolean if a field has been set.

### GetWebsiteId

`func (o *WebsiteAlertConfig) GetWebsiteId() string`

GetWebsiteId returns the WebsiteId field if non-nil, zero value otherwise.

### GetWebsiteIdOk

`func (o *WebsiteAlertConfig) GetWebsiteIdOk() (*string, bool)`

GetWebsiteIdOk returns a tuple with the WebsiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteId

`func (o *WebsiteAlertConfig) SetWebsiteId(v string)`

SetWebsiteId sets WebsiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


