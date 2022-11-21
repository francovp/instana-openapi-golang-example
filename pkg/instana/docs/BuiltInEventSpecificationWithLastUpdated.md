# BuiltInEventSpecificationWithLastUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**HyperParams** | [**[]HyperParam**](HyperParam.md) |  | 
**Id** | **string** |  | 
**LastUpdated** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**RuleInputs** | [**[]RuleInput**](RuleInput.md) |  | 
**Severity** | Pointer to **int32** |  | [optional] 
**ShortPluginId** | **string** |  | 
**Triggering** | Pointer to **bool** |  | [optional] 

## Methods

### NewBuiltInEventSpecificationWithLastUpdated

`func NewBuiltInEventSpecificationWithLastUpdated(hyperParams []HyperParam, id string, name string, ruleInputs []RuleInput, shortPluginId string, ) *BuiltInEventSpecificationWithLastUpdated`

NewBuiltInEventSpecificationWithLastUpdated instantiates a new BuiltInEventSpecificationWithLastUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuiltInEventSpecificationWithLastUpdatedWithDefaults

`func NewBuiltInEventSpecificationWithLastUpdatedWithDefaults() *BuiltInEventSpecificationWithLastUpdated`

NewBuiltInEventSpecificationWithLastUpdatedWithDefaults instantiates a new BuiltInEventSpecificationWithLastUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BuiltInEventSpecificationWithLastUpdated) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BuiltInEventSpecificationWithLastUpdated) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BuiltInEventSpecificationWithLastUpdated) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BuiltInEventSpecificationWithLastUpdated) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BuiltInEventSpecificationWithLastUpdated) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BuiltInEventSpecificationWithLastUpdated) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BuiltInEventSpecificationWithLastUpdated) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BuiltInEventSpecificationWithLastUpdated) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHidden

`func (o *BuiltInEventSpecificationWithLastUpdated) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *BuiltInEventSpecificationWithLastUpdated) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *BuiltInEventSpecificationWithLastUpdated) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *BuiltInEventSpecificationWithLastUpdated) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHyperParams

`func (o *BuiltInEventSpecificationWithLastUpdated) GetHyperParams() []HyperParam`

GetHyperParams returns the HyperParams field if non-nil, zero value otherwise.

### GetHyperParamsOk

`func (o *BuiltInEventSpecificationWithLastUpdated) GetHyperParamsOk() (*[]HyperParam, bool)`

GetHyperParamsOk returns a tuple with the HyperParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperParams

`func (o *BuiltInEventSpecificationWithLastUpdated) SetHyperParams(v []HyperParam)`

SetHyperParams sets HyperParams field to given value.


### GetId

`func (o *BuiltInEventSpecificationWithLastUpdated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BuiltInEventSpecificationWithLastUpdated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BuiltInEventSpecificationWithLastUpdated) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdated

`func (o *BuiltInEventSpecificationWithLastUpdated) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *BuiltInEventSpecificationWithLastUpdated) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *BuiltInEventSpecificationWithLastUpdated) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *BuiltInEventSpecificationWithLastUpdated) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *BuiltInEventSpecificationWithLastUpdated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuiltInEventSpecificationWithLastUpdated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuiltInEventSpecificationWithLastUpdated) SetName(v string)`

SetName sets Name field to given value.


### GetRuleInputs

`func (o *BuiltInEventSpecificationWithLastUpdated) GetRuleInputs() []RuleInput`

GetRuleInputs returns the RuleInputs field if non-nil, zero value otherwise.

### GetRuleInputsOk

`func (o *BuiltInEventSpecificationWithLastUpdated) GetRuleInputsOk() (*[]RuleInput, bool)`

GetRuleInputsOk returns a tuple with the RuleInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleInputs

`func (o *BuiltInEventSpecificationWithLastUpdated) SetRuleInputs(v []RuleInput)`

SetRuleInputs sets RuleInputs field to given value.


### GetSeverity

`func (o *BuiltInEventSpecificationWithLastUpdated) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *BuiltInEventSpecificationWithLastUpdated) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *BuiltInEventSpecificationWithLastUpdated) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *BuiltInEventSpecificationWithLastUpdated) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetShortPluginId

`func (o *BuiltInEventSpecificationWithLastUpdated) GetShortPluginId() string`

GetShortPluginId returns the ShortPluginId field if non-nil, zero value otherwise.

### GetShortPluginIdOk

`func (o *BuiltInEventSpecificationWithLastUpdated) GetShortPluginIdOk() (*string, bool)`

GetShortPluginIdOk returns a tuple with the ShortPluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortPluginId

`func (o *BuiltInEventSpecificationWithLastUpdated) SetShortPluginId(v string)`

SetShortPluginId sets ShortPluginId field to given value.


### GetTriggering

`func (o *BuiltInEventSpecificationWithLastUpdated) GetTriggering() bool`

GetTriggering returns the Triggering field if non-nil, zero value otherwise.

### GetTriggeringOk

`func (o *BuiltInEventSpecificationWithLastUpdated) GetTriggeringOk() (*bool, bool)`

GetTriggeringOk returns a tuple with the Triggering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggering

`func (o *BuiltInEventSpecificationWithLastUpdated) SetTriggering(v bool)`

SetTriggering sets Triggering field to given value.

### HasTriggering

`func (o *BuiltInEventSpecificationWithLastUpdated) HasTriggering() bool`

HasTriggering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


