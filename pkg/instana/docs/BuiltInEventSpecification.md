# BuiltInEventSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**HyperParams** | [**[]HyperParam**](HyperParam.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**RuleInputs** | [**[]RuleInput**](RuleInput.md) |  | 
**Severity** | Pointer to **int32** |  | [optional] 
**ShortPluginId** | **string** |  | 
**Triggering** | Pointer to **bool** |  | [optional] 

## Methods

### NewBuiltInEventSpecification

`func NewBuiltInEventSpecification(hyperParams []HyperParam, id string, name string, ruleInputs []RuleInput, shortPluginId string, ) *BuiltInEventSpecification`

NewBuiltInEventSpecification instantiates a new BuiltInEventSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuiltInEventSpecificationWithDefaults

`func NewBuiltInEventSpecificationWithDefaults() *BuiltInEventSpecification`

NewBuiltInEventSpecificationWithDefaults instantiates a new BuiltInEventSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BuiltInEventSpecification) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BuiltInEventSpecification) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BuiltInEventSpecification) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BuiltInEventSpecification) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BuiltInEventSpecification) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BuiltInEventSpecification) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BuiltInEventSpecification) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BuiltInEventSpecification) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHidden

`func (o *BuiltInEventSpecification) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *BuiltInEventSpecification) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *BuiltInEventSpecification) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *BuiltInEventSpecification) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetHyperParams

`func (o *BuiltInEventSpecification) GetHyperParams() []HyperParam`

GetHyperParams returns the HyperParams field if non-nil, zero value otherwise.

### GetHyperParamsOk

`func (o *BuiltInEventSpecification) GetHyperParamsOk() (*[]HyperParam, bool)`

GetHyperParamsOk returns a tuple with the HyperParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperParams

`func (o *BuiltInEventSpecification) SetHyperParams(v []HyperParam)`

SetHyperParams sets HyperParams field to given value.


### GetId

`func (o *BuiltInEventSpecification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BuiltInEventSpecification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BuiltInEventSpecification) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BuiltInEventSpecification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BuiltInEventSpecification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BuiltInEventSpecification) SetName(v string)`

SetName sets Name field to given value.


### GetRuleInputs

`func (o *BuiltInEventSpecification) GetRuleInputs() []RuleInput`

GetRuleInputs returns the RuleInputs field if non-nil, zero value otherwise.

### GetRuleInputsOk

`func (o *BuiltInEventSpecification) GetRuleInputsOk() (*[]RuleInput, bool)`

GetRuleInputsOk returns a tuple with the RuleInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleInputs

`func (o *BuiltInEventSpecification) SetRuleInputs(v []RuleInput)`

SetRuleInputs sets RuleInputs field to given value.


### GetSeverity

`func (o *BuiltInEventSpecification) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *BuiltInEventSpecification) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *BuiltInEventSpecification) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *BuiltInEventSpecification) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetShortPluginId

`func (o *BuiltInEventSpecification) GetShortPluginId() string`

GetShortPluginId returns the ShortPluginId field if non-nil, zero value otherwise.

### GetShortPluginIdOk

`func (o *BuiltInEventSpecification) GetShortPluginIdOk() (*string, bool)`

GetShortPluginIdOk returns a tuple with the ShortPluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortPluginId

`func (o *BuiltInEventSpecification) SetShortPluginId(v string)`

SetShortPluginId sets ShortPluginId field to given value.


### GetTriggering

`func (o *BuiltInEventSpecification) GetTriggering() bool`

GetTriggering returns the Triggering field if non-nil, zero value otherwise.

### GetTriggeringOk

`func (o *BuiltInEventSpecification) GetTriggeringOk() (*bool, bool)`

GetTriggeringOk returns a tuple with the Triggering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggering

`func (o *BuiltInEventSpecification) SetTriggering(v bool)`

SetTriggering sets Triggering field to given value.

### HasTriggering

`func (o *BuiltInEventSpecification) HasTriggering() bool`

HasTriggering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


