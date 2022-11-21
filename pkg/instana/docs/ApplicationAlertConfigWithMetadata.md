# ApplicationAlertConfigWithMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertChannelIds** | **[]string** |  | 
**ApplicationId** | Pointer to **string** |  | [optional] 
**Applications** | [**map[string]ApplicationNode**](ApplicationNode.md) |  | 
**BoundaryScope** | **string** |  | 
**Created** | Pointer to **int64** |  | [optional] 
**CustomPayloadFields** | [**[]CustomPayloadField**](CustomPayloadField.md) |  | 
**Description** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**EvaluationType** | **string** |  | 
**Granularity** | **int32** |  | [default to 600000]
**Id** | **string** |  | 
**IncludeInternal** | Pointer to **bool** |  | [optional] 
**IncludeSynthetic** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Rule** | [**ApplicationAlertRule**](ApplicationAlertRule.md) |  | 
**Severity** | Pointer to **int32** |  | [optional] 
**TagFilterExpression** | [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | 
**TagFilters** | Pointer to [**[]TagFilter**](TagFilter.md) |  | [optional] 
**Threshold** | [**Threshold**](Threshold.md) |  | 
**TimeThreshold** | [**ApplicationTimeThreshold**](ApplicationTimeThreshold.md) |  | 
**Triggering** | Pointer to **bool** |  | [optional] 

## Methods

### NewApplicationAlertConfigWithMetadata

`func NewApplicationAlertConfigWithMetadata(alertChannelIds []string, applications map[string]ApplicationNode, boundaryScope string, customPayloadFields []CustomPayloadField, description string, evaluationType string, granularity int32, id string, name string, rule ApplicationAlertRule, tagFilterExpression TagFilterExpressionElement, threshold Threshold, timeThreshold ApplicationTimeThreshold, ) *ApplicationAlertConfigWithMetadata`

NewApplicationAlertConfigWithMetadata instantiates a new ApplicationAlertConfigWithMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAlertConfigWithMetadataWithDefaults

`func NewApplicationAlertConfigWithMetadataWithDefaults() *ApplicationAlertConfigWithMetadata`

NewApplicationAlertConfigWithMetadataWithDefaults instantiates a new ApplicationAlertConfigWithMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertChannelIds

`func (o *ApplicationAlertConfigWithMetadata) GetAlertChannelIds() []string`

GetAlertChannelIds returns the AlertChannelIds field if non-nil, zero value otherwise.

### GetAlertChannelIdsOk

`func (o *ApplicationAlertConfigWithMetadata) GetAlertChannelIdsOk() (*[]string, bool)`

GetAlertChannelIdsOk returns a tuple with the AlertChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertChannelIds

`func (o *ApplicationAlertConfigWithMetadata) SetAlertChannelIds(v []string)`

SetAlertChannelIds sets AlertChannelIds field to given value.


### GetApplicationId

`func (o *ApplicationAlertConfigWithMetadata) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationAlertConfigWithMetadata) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationAlertConfigWithMetadata) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApplicationAlertConfigWithMetadata) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplications

`func (o *ApplicationAlertConfigWithMetadata) GetApplications() map[string]ApplicationNode`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ApplicationAlertConfigWithMetadata) GetApplicationsOk() (*map[string]ApplicationNode, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ApplicationAlertConfigWithMetadata) SetApplications(v map[string]ApplicationNode)`

SetApplications sets Applications field to given value.


### GetBoundaryScope

`func (o *ApplicationAlertConfigWithMetadata) GetBoundaryScope() string`

GetBoundaryScope returns the BoundaryScope field if non-nil, zero value otherwise.

### GetBoundaryScopeOk

`func (o *ApplicationAlertConfigWithMetadata) GetBoundaryScopeOk() (*string, bool)`

GetBoundaryScopeOk returns a tuple with the BoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaryScope

`func (o *ApplicationAlertConfigWithMetadata) SetBoundaryScope(v string)`

SetBoundaryScope sets BoundaryScope field to given value.


### GetCreated

`func (o *ApplicationAlertConfigWithMetadata) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationAlertConfigWithMetadata) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApplicationAlertConfigWithMetadata) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApplicationAlertConfigWithMetadata) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCustomPayloadFields

`func (o *ApplicationAlertConfigWithMetadata) GetCustomPayloadFields() []CustomPayloadField`

GetCustomPayloadFields returns the CustomPayloadFields field if non-nil, zero value otherwise.

### GetCustomPayloadFieldsOk

`func (o *ApplicationAlertConfigWithMetadata) GetCustomPayloadFieldsOk() (*[]CustomPayloadField, bool)`

GetCustomPayloadFieldsOk returns a tuple with the CustomPayloadFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadFields

`func (o *ApplicationAlertConfigWithMetadata) SetCustomPayloadFields(v []CustomPayloadField)`

SetCustomPayloadFields sets CustomPayloadFields field to given value.


### GetDescription

`func (o *ApplicationAlertConfigWithMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationAlertConfigWithMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationAlertConfigWithMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *ApplicationAlertConfigWithMetadata) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationAlertConfigWithMetadata) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationAlertConfigWithMetadata) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApplicationAlertConfigWithMetadata) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvaluationType

`func (o *ApplicationAlertConfigWithMetadata) GetEvaluationType() string`

GetEvaluationType returns the EvaluationType field if non-nil, zero value otherwise.

### GetEvaluationTypeOk

`func (o *ApplicationAlertConfigWithMetadata) GetEvaluationTypeOk() (*string, bool)`

GetEvaluationTypeOk returns a tuple with the EvaluationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationType

`func (o *ApplicationAlertConfigWithMetadata) SetEvaluationType(v string)`

SetEvaluationType sets EvaluationType field to given value.


### GetGranularity

`func (o *ApplicationAlertConfigWithMetadata) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *ApplicationAlertConfigWithMetadata) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *ApplicationAlertConfigWithMetadata) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.


### GetId

`func (o *ApplicationAlertConfigWithMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationAlertConfigWithMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationAlertConfigWithMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetIncludeInternal

`func (o *ApplicationAlertConfigWithMetadata) GetIncludeInternal() bool`

GetIncludeInternal returns the IncludeInternal field if non-nil, zero value otherwise.

### GetIncludeInternalOk

`func (o *ApplicationAlertConfigWithMetadata) GetIncludeInternalOk() (*bool, bool)`

GetIncludeInternalOk returns a tuple with the IncludeInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInternal

`func (o *ApplicationAlertConfigWithMetadata) SetIncludeInternal(v bool)`

SetIncludeInternal sets IncludeInternal field to given value.

### HasIncludeInternal

`func (o *ApplicationAlertConfigWithMetadata) HasIncludeInternal() bool`

HasIncludeInternal returns a boolean if a field has been set.

### GetIncludeSynthetic

`func (o *ApplicationAlertConfigWithMetadata) GetIncludeSynthetic() bool`

GetIncludeSynthetic returns the IncludeSynthetic field if non-nil, zero value otherwise.

### GetIncludeSyntheticOk

`func (o *ApplicationAlertConfigWithMetadata) GetIncludeSyntheticOk() (*bool, bool)`

GetIncludeSyntheticOk returns a tuple with the IncludeSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSynthetic

`func (o *ApplicationAlertConfigWithMetadata) SetIncludeSynthetic(v bool)`

SetIncludeSynthetic sets IncludeSynthetic field to given value.

### HasIncludeSynthetic

`func (o *ApplicationAlertConfigWithMetadata) HasIncludeSynthetic() bool`

HasIncludeSynthetic returns a boolean if a field has been set.

### GetName

`func (o *ApplicationAlertConfigWithMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationAlertConfigWithMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationAlertConfigWithMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetReadOnly

`func (o *ApplicationAlertConfigWithMetadata) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ApplicationAlertConfigWithMetadata) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ApplicationAlertConfigWithMetadata) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ApplicationAlertConfigWithMetadata) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRule

`func (o *ApplicationAlertConfigWithMetadata) GetRule() ApplicationAlertRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *ApplicationAlertConfigWithMetadata) GetRuleOk() (*ApplicationAlertRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *ApplicationAlertConfigWithMetadata) SetRule(v ApplicationAlertRule)`

SetRule sets Rule field to given value.


### GetSeverity

`func (o *ApplicationAlertConfigWithMetadata) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ApplicationAlertConfigWithMetadata) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ApplicationAlertConfigWithMetadata) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ApplicationAlertConfigWithMetadata) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTagFilterExpression

`func (o *ApplicationAlertConfigWithMetadata) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *ApplicationAlertConfigWithMetadata) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *ApplicationAlertConfigWithMetadata) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.


### GetTagFilters

`func (o *ApplicationAlertConfigWithMetadata) GetTagFilters() []TagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *ApplicationAlertConfigWithMetadata) GetTagFiltersOk() (*[]TagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *ApplicationAlertConfigWithMetadata) SetTagFilters(v []TagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *ApplicationAlertConfigWithMetadata) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetThreshold

`func (o *ApplicationAlertConfigWithMetadata) GetThreshold() Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ApplicationAlertConfigWithMetadata) GetThresholdOk() (*Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ApplicationAlertConfigWithMetadata) SetThreshold(v Threshold)`

SetThreshold sets Threshold field to given value.


### GetTimeThreshold

`func (o *ApplicationAlertConfigWithMetadata) GetTimeThreshold() ApplicationTimeThreshold`

GetTimeThreshold returns the TimeThreshold field if non-nil, zero value otherwise.

### GetTimeThresholdOk

`func (o *ApplicationAlertConfigWithMetadata) GetTimeThresholdOk() (*ApplicationTimeThreshold, bool)`

GetTimeThresholdOk returns a tuple with the TimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeThreshold

`func (o *ApplicationAlertConfigWithMetadata) SetTimeThreshold(v ApplicationTimeThreshold)`

SetTimeThreshold sets TimeThreshold field to given value.


### GetTriggering

`func (o *ApplicationAlertConfigWithMetadata) GetTriggering() bool`

GetTriggering returns the Triggering field if non-nil, zero value otherwise.

### GetTriggeringOk

`func (o *ApplicationAlertConfigWithMetadata) GetTriggeringOk() (*bool, bool)`

GetTriggeringOk returns a tuple with the Triggering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggering

`func (o *ApplicationAlertConfigWithMetadata) SetTriggering(v bool)`

SetTriggering sets Triggering field to given value.

### HasTriggering

`func (o *ApplicationAlertConfigWithMetadata) HasTriggering() bool`

HasTriggering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


