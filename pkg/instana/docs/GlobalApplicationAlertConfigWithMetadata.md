# GlobalApplicationAlertConfigWithMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertChannelIds** | **[]string** |  | 
**ApplicationIds** | Pointer to **[]string** |  | [optional] 
**Applications** | [**map[string]ApplicationNode**](ApplicationNode.md) |  | 
**BoundaryScope** | **string** |  | 
**BuiltIn** | Pointer to **bool** |  | [optional] 
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
**Threshold** | [**Threshold**](Threshold.md) |  | 
**TimeThreshold** | [**ApplicationTimeThreshold**](ApplicationTimeThreshold.md) |  | 
**Triggering** | Pointer to **bool** |  | [optional] 

## Methods

### NewGlobalApplicationAlertConfigWithMetadata

`func NewGlobalApplicationAlertConfigWithMetadata(alertChannelIds []string, applications map[string]ApplicationNode, boundaryScope string, customPayloadFields []CustomPayloadField, description string, evaluationType string, granularity int32, id string, name string, rule ApplicationAlertRule, tagFilterExpression TagFilterExpressionElement, threshold Threshold, timeThreshold ApplicationTimeThreshold, ) *GlobalApplicationAlertConfigWithMetadata`

NewGlobalApplicationAlertConfigWithMetadata instantiates a new GlobalApplicationAlertConfigWithMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalApplicationAlertConfigWithMetadataWithDefaults

`func NewGlobalApplicationAlertConfigWithMetadataWithDefaults() *GlobalApplicationAlertConfigWithMetadata`

NewGlobalApplicationAlertConfigWithMetadataWithDefaults instantiates a new GlobalApplicationAlertConfigWithMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertChannelIds

`func (o *GlobalApplicationAlertConfigWithMetadata) GetAlertChannelIds() []string`

GetAlertChannelIds returns the AlertChannelIds field if non-nil, zero value otherwise.

### GetAlertChannelIdsOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetAlertChannelIdsOk() (*[]string, bool)`

GetAlertChannelIdsOk returns a tuple with the AlertChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertChannelIds

`func (o *GlobalApplicationAlertConfigWithMetadata) SetAlertChannelIds(v []string)`

SetAlertChannelIds sets AlertChannelIds field to given value.


### GetApplicationIds

`func (o *GlobalApplicationAlertConfigWithMetadata) GetApplicationIds() []string`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetApplicationIdsOk() (*[]string, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *GlobalApplicationAlertConfigWithMetadata) SetApplicationIds(v []string)`

SetApplicationIds sets ApplicationIds field to given value.

### HasApplicationIds

`func (o *GlobalApplicationAlertConfigWithMetadata) HasApplicationIds() bool`

HasApplicationIds returns a boolean if a field has been set.

### GetApplications

`func (o *GlobalApplicationAlertConfigWithMetadata) GetApplications() map[string]ApplicationNode`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetApplicationsOk() (*map[string]ApplicationNode, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *GlobalApplicationAlertConfigWithMetadata) SetApplications(v map[string]ApplicationNode)`

SetApplications sets Applications field to given value.


### GetBoundaryScope

`func (o *GlobalApplicationAlertConfigWithMetadata) GetBoundaryScope() string`

GetBoundaryScope returns the BoundaryScope field if non-nil, zero value otherwise.

### GetBoundaryScopeOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetBoundaryScopeOk() (*string, bool)`

GetBoundaryScopeOk returns a tuple with the BoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaryScope

`func (o *GlobalApplicationAlertConfigWithMetadata) SetBoundaryScope(v string)`

SetBoundaryScope sets BoundaryScope field to given value.


### GetBuiltIn

`func (o *GlobalApplicationAlertConfigWithMetadata) GetBuiltIn() bool`

GetBuiltIn returns the BuiltIn field if non-nil, zero value otherwise.

### GetBuiltInOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetBuiltInOk() (*bool, bool)`

GetBuiltInOk returns a tuple with the BuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltIn

`func (o *GlobalApplicationAlertConfigWithMetadata) SetBuiltIn(v bool)`

SetBuiltIn sets BuiltIn field to given value.

### HasBuiltIn

`func (o *GlobalApplicationAlertConfigWithMetadata) HasBuiltIn() bool`

HasBuiltIn returns a boolean if a field has been set.

### GetCreated

`func (o *GlobalApplicationAlertConfigWithMetadata) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GlobalApplicationAlertConfigWithMetadata) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GlobalApplicationAlertConfigWithMetadata) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCustomPayloadFields

`func (o *GlobalApplicationAlertConfigWithMetadata) GetCustomPayloadFields() []CustomPayloadField`

GetCustomPayloadFields returns the CustomPayloadFields field if non-nil, zero value otherwise.

### GetCustomPayloadFieldsOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetCustomPayloadFieldsOk() (*[]CustomPayloadField, bool)`

GetCustomPayloadFieldsOk returns a tuple with the CustomPayloadFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayloadFields

`func (o *GlobalApplicationAlertConfigWithMetadata) SetCustomPayloadFields(v []CustomPayloadField)`

SetCustomPayloadFields sets CustomPayloadFields field to given value.


### GetDescription

`func (o *GlobalApplicationAlertConfigWithMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalApplicationAlertConfigWithMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *GlobalApplicationAlertConfigWithMetadata) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalApplicationAlertConfigWithMetadata) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GlobalApplicationAlertConfigWithMetadata) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvaluationType

`func (o *GlobalApplicationAlertConfigWithMetadata) GetEvaluationType() string`

GetEvaluationType returns the EvaluationType field if non-nil, zero value otherwise.

### GetEvaluationTypeOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetEvaluationTypeOk() (*string, bool)`

GetEvaluationTypeOk returns a tuple with the EvaluationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationType

`func (o *GlobalApplicationAlertConfigWithMetadata) SetEvaluationType(v string)`

SetEvaluationType sets EvaluationType field to given value.


### GetGranularity

`func (o *GlobalApplicationAlertConfigWithMetadata) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *GlobalApplicationAlertConfigWithMetadata) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.


### GetId

`func (o *GlobalApplicationAlertConfigWithMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalApplicationAlertConfigWithMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetIncludeInternal

`func (o *GlobalApplicationAlertConfigWithMetadata) GetIncludeInternal() bool`

GetIncludeInternal returns the IncludeInternal field if non-nil, zero value otherwise.

### GetIncludeInternalOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetIncludeInternalOk() (*bool, bool)`

GetIncludeInternalOk returns a tuple with the IncludeInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInternal

`func (o *GlobalApplicationAlertConfigWithMetadata) SetIncludeInternal(v bool)`

SetIncludeInternal sets IncludeInternal field to given value.

### HasIncludeInternal

`func (o *GlobalApplicationAlertConfigWithMetadata) HasIncludeInternal() bool`

HasIncludeInternal returns a boolean if a field has been set.

### GetIncludeSynthetic

`func (o *GlobalApplicationAlertConfigWithMetadata) GetIncludeSynthetic() bool`

GetIncludeSynthetic returns the IncludeSynthetic field if non-nil, zero value otherwise.

### GetIncludeSyntheticOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetIncludeSyntheticOk() (*bool, bool)`

GetIncludeSyntheticOk returns a tuple with the IncludeSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSynthetic

`func (o *GlobalApplicationAlertConfigWithMetadata) SetIncludeSynthetic(v bool)`

SetIncludeSynthetic sets IncludeSynthetic field to given value.

### HasIncludeSynthetic

`func (o *GlobalApplicationAlertConfigWithMetadata) HasIncludeSynthetic() bool`

HasIncludeSynthetic returns a boolean if a field has been set.

### GetName

`func (o *GlobalApplicationAlertConfigWithMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalApplicationAlertConfigWithMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetReadOnly

`func (o *GlobalApplicationAlertConfigWithMetadata) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *GlobalApplicationAlertConfigWithMetadata) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *GlobalApplicationAlertConfigWithMetadata) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRule

`func (o *GlobalApplicationAlertConfigWithMetadata) GetRule() ApplicationAlertRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetRuleOk() (*ApplicationAlertRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *GlobalApplicationAlertConfigWithMetadata) SetRule(v ApplicationAlertRule)`

SetRule sets Rule field to given value.


### GetSeverity

`func (o *GlobalApplicationAlertConfigWithMetadata) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GlobalApplicationAlertConfigWithMetadata) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GlobalApplicationAlertConfigWithMetadata) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTagFilterExpression

`func (o *GlobalApplicationAlertConfigWithMetadata) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *GlobalApplicationAlertConfigWithMetadata) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.


### GetThreshold

`func (o *GlobalApplicationAlertConfigWithMetadata) GetThreshold() Threshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetThresholdOk() (*Threshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *GlobalApplicationAlertConfigWithMetadata) SetThreshold(v Threshold)`

SetThreshold sets Threshold field to given value.


### GetTimeThreshold

`func (o *GlobalApplicationAlertConfigWithMetadata) GetTimeThreshold() ApplicationTimeThreshold`

GetTimeThreshold returns the TimeThreshold field if non-nil, zero value otherwise.

### GetTimeThresholdOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetTimeThresholdOk() (*ApplicationTimeThreshold, bool)`

GetTimeThresholdOk returns a tuple with the TimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeThreshold

`func (o *GlobalApplicationAlertConfigWithMetadata) SetTimeThreshold(v ApplicationTimeThreshold)`

SetTimeThreshold sets TimeThreshold field to given value.


### GetTriggering

`func (o *GlobalApplicationAlertConfigWithMetadata) GetTriggering() bool`

GetTriggering returns the Triggering field if non-nil, zero value otherwise.

### GetTriggeringOk

`func (o *GlobalApplicationAlertConfigWithMetadata) GetTriggeringOk() (*bool, bool)`

GetTriggeringOk returns a tuple with the Triggering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggering

`func (o *GlobalApplicationAlertConfigWithMetadata) SetTriggering(v bool)`

SetTriggering sets Triggering field to given value.

### HasTriggering

`func (o *GlobalApplicationAlertConfigWithMetadata) HasTriggering() bool`

HasTriggering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


