# SyntheticCallConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomRules** | [**[]SyntheticCallRule**](SyntheticCallRule.md) |  | 
**DefaultRulesEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewSyntheticCallConfig

`func NewSyntheticCallConfig(customRules []SyntheticCallRule, ) *SyntheticCallConfig`

NewSyntheticCallConfig instantiates a new SyntheticCallConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticCallConfigWithDefaults

`func NewSyntheticCallConfigWithDefaults() *SyntheticCallConfig`

NewSyntheticCallConfigWithDefaults instantiates a new SyntheticCallConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomRules

`func (o *SyntheticCallConfig) GetCustomRules() []SyntheticCallRule`

GetCustomRules returns the CustomRules field if non-nil, zero value otherwise.

### GetCustomRulesOk

`func (o *SyntheticCallConfig) GetCustomRulesOk() (*[]SyntheticCallRule, bool)`

GetCustomRulesOk returns a tuple with the CustomRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRules

`func (o *SyntheticCallConfig) SetCustomRules(v []SyntheticCallRule)`

SetCustomRules sets CustomRules field to given value.


### GetDefaultRulesEnabled

`func (o *SyntheticCallConfig) GetDefaultRulesEnabled() bool`

GetDefaultRulesEnabled returns the DefaultRulesEnabled field if non-nil, zero value otherwise.

### GetDefaultRulesEnabledOk

`func (o *SyntheticCallConfig) GetDefaultRulesEnabledOk() (*bool, bool)`

GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRulesEnabled

`func (o *SyntheticCallConfig) SetDefaultRulesEnabled(v bool)`

SetDefaultRulesEnabled sets DefaultRulesEnabled field to given value.

### HasDefaultRulesEnabled

`func (o *SyntheticCallConfig) HasDefaultRulesEnabled() bool`

HasDefaultRulesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


