# ApplicationAlertConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertChannelIds** | **[]string** |  | 
**ApplicationId** | Pointer to **string** |  | [optional] 
**Applications** | [**map[string]ApplicationNode**](ApplicationNode.md) |  | 
**BoundaryScope** | **string** |  | 
**CustomPayloadFields** | [**[]CustomPayloadField**](CustomPayloadField.md) |  | 
**Description** | **string** |  | 
**EvaluationType** | **string** |  | 
**Granularity** | **int32** |  | [default to 600000]
**IncludeInternal** | Pointer to **bool** |  | [optional] 
**IncludeSynthetic** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Rule** | [**ApplicationAlertRule**](ApplicationAlertRule.md) |  | 
**Severity** | Pointer to **int32** |  | [optional] 
**TagFilterExpression** | [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | 
**TagFilters** | Pointer to [**[]TagFilter**](TagFilter.md) |  | [optional] 
**Threshold** | [**Threshold**](Threshold.md) |  | 
**TimeThreshold** | [**ApplicationTimeThreshold**](ApplicationTimeThreshold.md) |  | 
**Triggering** | Pointer to **bool** |  | [optional] 

## Methods

### NewApplicationAlertConfig

`func NewApplicationAlertConfig(alertChannelIds []string, applications map[string]ApplicationNode, boundaryScope string, customPayloadFields []CustomPayloadField, description string, evaluationType string, granularity int32, name string, rule ApplicationAlertRule, tagFilterExpression TagFilterExpressionElement, threshold Threshold, timeThreshold ApplicationTimeThreshold, ) *ApplicationAlertConfig`

NewApplicationAlertConfig instantiates a new ApplicationAlertConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAlertConfigWithDefaults

`func NewApplicationAlertConfigWithDefaults() *ApplicationAlertConfig`

NewApplicationAlertConfigWithDefaults instantiates a new ApplicationAlertConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertChannelIds

`func (o *ApplicationAlertConfig) GetAlertChannelIds() []string`

GetAlertChannelIds returns the AlertChannelIds field if non-nil, zero value otherwise.

### GetAlertChannelIdsOk

`func (o *ApplicationAlertConfig) GetAlertChannelIdsOk() (*[]string, bool)`

GetAlertChannelIdsOk returns a tuple with the AlertChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertChannelIds

`func (o *ApplicationAlertConfig) SetAlertChannelIds(v []string)`

SetAlertChannelIds sets AlertChannelIds field to given value.


### GetApplicationId

`func (o *ApplicationAlertConfig) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationAlertConfig) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationAlertConfig) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApplicationAlertConfig) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplications

`func (o *ApplicationAlertConfig) GetApplications() map[string]ApplicationNode`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ApplicationAlertConfig) GetApplicationsOk() (*map[string]ApplicationNode, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ApplicationAlertConfig) SetApplications(v map[string]ApplicationNode)`

SetApplications sets Applications field to given value.


### GetBoundaryScope

`func (o *ApplicationAlertConfig) GetBoundaryScope() string`

GetBoundaryScope returns the BoundaryScope field if non-nil, zero value otherwise.

### GetBoundaryScopeOk

`func (o *ApplicationAlertConfig) GetBoundaryScopeOk() (*string, bool)`

GetBoundaryScopeOk returns a tuple with the BoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaryScope

`func (o *ApplicationAlertConfig) SetBoundaryScope(v string)`

SetBoundaryScope sets BoundaryScope field to given value.


### GetCustomPayloadFields

`func (o *ApplicationAlertConfig) GetCustomPayloadFields() []CustomPayloadField`

GetCustomPayloadFields returns the CustomPayloadFields field if non-nil, zero value otherwise.

### GetCustomPayloadFieldsOk

`func (o *ApplicationAlertConfig) GetCustomPayloadFieldsOk() (*[]CustomPayloadField, bool)`

GetCustomPayloadFieldsOk returns a tuple with the CustomPayloadFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadFields

`func (o *ApplicationAlertConfig) SetCustomPayloadFields(v []CustomPayloadField)`

SetCustomPayloadFields sets CustomPayloadFields field to given value.


### GetDescription

`func (o *ApplicationAlertConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationAlertConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationAlertConfig) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEvaluationType

`func (o *ApplicationAlertConfig) GetEvaluationType() string`

GetEvaluationType returns the EvaluationType field if non-nil, zero value otherwise.

### GetEvaluationTypeOk

`func (o *ApplicationAlertConfig) GetEvaluationTypeOk() (*string, bool)`

GetEvaluationTypeOk returns a tuple with the EvaluationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationType

`func (o *ApplicationAlertConfig) SetEvaluationType(v string)`

SetEvaluationType sets EvaluationType field to given value.


### GetGranularity

`func (o *ApplicationAlertConfig) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *ApplicationAlertConfig) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *ApplicationAlertConfig) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.


### GetIncludeInternal

`func (o *ApplicationAlertConfig) GetIncludeInternal() bool`

GetIncludeInternal returns the IncludeInternal field if non-nil, zero value otherwise.

### GetIncludeInternalOk

`func (o *ApplicationAlertConfig) GetIncludeInternalOk() (*bool, bool)`

GetIncludeInternalOk returns a tuple with the IncludeInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInternal

`func (o *ApplicationAlertConfig) SetIncludeInternal(v bool)`

SetIncludeInternal sets IncludeInternal field to given value.

### HasIncludeInternal

`func (o *ApplicationAlertConfig) HasIncludeInternal() bool`

HasIncludeInternal returns a boolean if a field has been set.

### GetIncludeSynthetic

`func (o *ApplicationAlertConfig) GetIncludeSynthetic() bool`

GetIncludeSynthetic returns the IncludeSynthetic field if non-nil, zero value otherwise.

### GetIncludeSyntheticOk

`func (o *ApplicationAlertConfig) GetIncludeSyntheticOk() (*bool, bool)`

GetIncludeSyntheticOk returns a tuple with the IncludeSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSynthetic

`func (o *ApplicationAlertConfig) SetIncludeSynthetic(v bool)`

SetIncludeSynthetic sets IncludeSynthetic field to given value.

### HasIncludeSynthetic

`func (o *ApplicationAlertConfig) HasIncludeSynthetic() bool`

HasIncludeSynthetic returns a boolean if a field has been set.

### GetName

`func (o *ApplicationAlertConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationAlertConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationAlertConfig) SetName(v string)`

SetName sets Name field to given value.


### GetRule

`func (o *ApplicationAlertConfig) GetRule() ApplicationAlertRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *ApplicationAlertConfig) GetRuleOk() (*ApplicationAlertRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *ApplicationAlertConfig) SetRule(v ApplicationAlertRule)`

SetRule sets Rule field to given value.


### GetSeverity

`func (o *ApplicationAlertConfig) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ApplicationAlertConfig) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ApplicationAlertConfig) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ApplicationAlertConfig) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTagFilterExpression

`func (o *ApplicationAlertConfig) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *ApplicationAlertConfig) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *ApplicationAlertConfig) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.


### GetTagFilters

`func (o *ApplicationAlertConfig) GetTagFilters() []TagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *ApplicationAlertConfig) GetTagFiltersOk() (*[]TagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *ApplicationAlertConfig) SetTagFilters(v []TagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *ApplicationAlertConfig) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetThreshold

`func (o *ApplicationAlertConfig) GetThreshold() Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ApplicationAlertConfig) GetThresholdOk() (*Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ApplicationAlertConfig) SetThreshold(v Threshold)`

SetThreshold sets Threshold field to given value.


### GetTimeThreshold

`func (o *ApplicationAlertConfig) GetTimeThreshold() ApplicationTimeThreshold`

GetTimeThreshold returns the TimeThreshold field if non-nil, zero value otherwise.

### GetTimeThresholdOk

`func (o *ApplicationAlertConfig) GetTimeThresholdOk() (*ApplicationTimeThreshold, bool)`

GetTimeThresholdOk returns a tuple with the TimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeThreshold

`func (o *ApplicationAlertConfig) SetTimeThreshold(v ApplicationTimeThreshold)`

SetTimeThreshold sets TimeThreshold field to given value.


### GetTriggering

`func (o *ApplicationAlertConfig) GetTriggering() bool`

GetTriggering returns the Triggering field if non-nil, zero value otherwise.

### GetTriggeringOk

`func (o *ApplicationAlertConfig) GetTriggeringOk() (*bool, bool)`

GetTriggeringOk returns a tuple with the Triggering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggering

`func (o *ApplicationAlertConfig) SetTriggering(v bool)`

SetTriggering sets Triggering field to given value.

### HasTriggering

`func (o *ApplicationAlertConfig) HasTriggering() bool`

HasTriggering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


