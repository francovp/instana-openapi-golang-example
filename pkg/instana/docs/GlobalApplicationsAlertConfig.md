# GlobalApplicationsAlertConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertChannelIds** | **[]string** |  | 
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
**Threshold** | [**Threshold**](Threshold.md) |  | 
**TimeThreshold** | [**ApplicationTimeThreshold**](ApplicationTimeThreshold.md) |  | 
**Triggering** | Pointer to **bool** |  | [optional] 

## Methods

### NewGlobalApplicationsAlertConfig

`func NewGlobalApplicationsAlertConfig(alertChannelIds []string, applications map[string]ApplicationNode, boundaryScope string, customPayloadFields []CustomPayloadField, description string, evaluationType string, granularity int32, name string, rule ApplicationAlertRule, tagFilterExpression TagFilterExpressionElement, threshold Threshold, timeThreshold ApplicationTimeThreshold, ) *GlobalApplicationsAlertConfig`

NewGlobalApplicationsAlertConfig instantiates a new GlobalApplicationsAlertConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalApplicationsAlertConfigWithDefaults

`func NewGlobalApplicationsAlertConfigWithDefaults() *GlobalApplicationsAlertConfig`

NewGlobalApplicationsAlertConfigWithDefaults instantiates a new GlobalApplicationsAlertConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertChannelIds

`func (o *GlobalApplicationsAlertConfig) GetAlertChannelIds() []string`

GetAlertChannelIds returns the AlertChannelIds field if non-nil, zero value otherwise.

### GetAlertChannelIdsOk

`func (o *GlobalApplicationsAlertConfig) GetAlertChannelIdsOk() (*[]string, bool)`

GetAlertChannelIdsOk returns a tuple with the AlertChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertChannelIds

`func (o *GlobalApplicationsAlertConfig) SetAlertChannelIds(v []string)`

SetAlertChannelIds sets AlertChannelIds field to given value.


### GetApplications

`func (o *GlobalApplicationsAlertConfig) GetApplications() map[string]ApplicationNode`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *GlobalApplicationsAlertConfig) GetApplicationsOk() (*map[string]ApplicationNode, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *GlobalApplicationsAlertConfig) SetApplications(v map[string]ApplicationNode)`

SetApplications sets Applications field to given value.


### GetBoundaryScope

`func (o *GlobalApplicationsAlertConfig) GetBoundaryScope() string`

GetBoundaryScope returns the BoundaryScope field if non-nil, zero value otherwise.

### GetBoundaryScopeOk

`func (o *GlobalApplicationsAlertConfig) GetBoundaryScopeOk() (*string, bool)`

GetBoundaryScopeOk returns a tuple with the BoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaryScope

`func (o *GlobalApplicationsAlertConfig) SetBoundaryScope(v string)`

SetBoundaryScope sets BoundaryScope field to given value.


### GetCustomPayloadFields

`func (o *GlobalApplicationsAlertConfig) GetCustomPayloadFields() []CustomPayloadField`

GetCustomPayloadFields returns the CustomPayloadFields field if non-nil, zero value otherwise.

### GetCustomPayloadFieldsOk

`func (o *GlobalApplicationsAlertConfig) GetCustomPayloadFieldsOk() (*[]CustomPayloadField, bool)`

GetCustomPayloadFieldsOk returns a tuple with the CustomPayloadFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadFields

`func (o *GlobalApplicationsAlertConfig) SetCustomPayloadFields(v []CustomPayloadField)`

SetCustomPayloadFields sets CustomPayloadFields field to given value.


### GetDescription

`func (o *GlobalApplicationsAlertConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalApplicationsAlertConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalApplicationsAlertConfig) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEvaluationType

`func (o *GlobalApplicationsAlertConfig) GetEvaluationType() string`

GetEvaluationType returns the EvaluationType field if non-nil, zero value otherwise.

### GetEvaluationTypeOk

`func (o *GlobalApplicationsAlertConfig) GetEvaluationTypeOk() (*string, bool)`

GetEvaluationTypeOk returns a tuple with the EvaluationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationType

`func (o *GlobalApplicationsAlertConfig) SetEvaluationType(v string)`

SetEvaluationType sets EvaluationType field to given value.


### GetGranularity

`func (o *GlobalApplicationsAlertConfig) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *GlobalApplicationsAlertConfig) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *GlobalApplicationsAlertConfig) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.


### GetIncludeInternal

`func (o *GlobalApplicationsAlertConfig) GetIncludeInternal() bool`

GetIncludeInternal returns the IncludeInternal field if non-nil, zero value otherwise.

### GetIncludeInternalOk

`func (o *GlobalApplicationsAlertConfig) GetIncludeInternalOk() (*bool, bool)`

GetIncludeInternalOk returns a tuple with the IncludeInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInternal

`func (o *GlobalApplicationsAlertConfig) SetIncludeInternal(v bool)`

SetIncludeInternal sets IncludeInternal field to given value.

### HasIncludeInternal

`func (o *GlobalApplicationsAlertConfig) HasIncludeInternal() bool`

HasIncludeInternal returns a boolean if a field has been set.

### GetIncludeSynthetic

`func (o *GlobalApplicationsAlertConfig) GetIncludeSynthetic() bool`

GetIncludeSynthetic returns the IncludeSynthetic field if non-nil, zero value otherwise.

### GetIncludeSyntheticOk

`func (o *GlobalApplicationsAlertConfig) GetIncludeSyntheticOk() (*bool, bool)`

GetIncludeSyntheticOk returns a tuple with the IncludeSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSynthetic

`func (o *GlobalApplicationsAlertConfig) SetIncludeSynthetic(v bool)`

SetIncludeSynthetic sets IncludeSynthetic field to given value.

### HasIncludeSynthetic

`func (o *GlobalApplicationsAlertConfig) HasIncludeSynthetic() bool`

HasIncludeSynthetic returns a boolean if a field has been set.

### GetName

`func (o *GlobalApplicationsAlertConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalApplicationsAlertConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalApplicationsAlertConfig) SetName(v string)`

SetName sets Name field to given value.


### GetRule

`func (o *GlobalApplicationsAlertConfig) GetRule() ApplicationAlertRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *GlobalApplicationsAlertConfig) GetRuleOk() (*ApplicationAlertRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *GlobalApplicationsAlertConfig) SetRule(v ApplicationAlertRule)`

SetRule sets Rule field to given value.


### GetSeverity

`func (o *GlobalApplicationsAlertConfig) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GlobalApplicationsAlertConfig) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GlobalApplicationsAlertConfig) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GlobalApplicationsAlertConfig) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTagFilterExpression

`func (o *GlobalApplicationsAlertConfig) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *GlobalApplicationsAlertConfig) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *GlobalApplicationsAlertConfig) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.


### GetThreshold

`func (o *GlobalApplicationsAlertConfig) GetThreshold() Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *GlobalApplicationsAlertConfig) GetThresholdOk() (*Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *GlobalApplicationsAlertConfig) SetThreshold(v Threshold)`

SetThreshold sets Threshold field to given value.


### GetTimeThreshold

`func (o *GlobalApplicationsAlertConfig) GetTimeThreshold() ApplicationTimeThreshold`

GetTimeThreshold returns the TimeThreshold field if non-nil, zero value otherwise.

### GetTimeThresholdOk

`func (o *GlobalApplicationsAlertConfig) GetTimeThresholdOk() (*ApplicationTimeThreshold, bool)`

GetTimeThresholdOk returns a tuple with the TimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeThreshold

`func (o *GlobalApplicationsAlertConfig) SetTimeThreshold(v ApplicationTimeThreshold)`

SetTimeThreshold sets TimeThreshold field to given value.


### GetTriggering

`func (o *GlobalApplicationsAlertConfig) GetTriggering() bool`

GetTriggering returns the Triggering field if non-nil, zero value otherwise.

### GetTriggeringOk

`func (o *GlobalApplicationsAlertConfig) GetTriggeringOk() (*bool, bool)`

GetTriggeringOk returns a tuple with the Triggering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggering

`func (o *GlobalApplicationsAlertConfig) SetTriggering(v bool)`

SetTriggering sets Triggering field to given value.

### HasTriggering

`func (o *GlobalApplicationsAlertConfig) HasTriggering() bool`

HasTriggering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


