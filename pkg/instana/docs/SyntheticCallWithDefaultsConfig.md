# SyntheticCallWithDefaultsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomRules** | [**[]SyntheticCallRule**](SyntheticCallRule.md) |  | 
**DefaultRules** | [**[]SyntheticCallRule**](SyntheticCallRule.md) |  | 
**DefaultRulesEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewSyntheticCallWithDefaultsConfig

`func NewSyntheticCallWithDefaultsConfig(customRules []SyntheticCallRule, defaultRules []SyntheticCallRule, ) *SyntheticCallWithDefaultsConfig`

NewSyntheticCallWithDefaultsConfig instantiates a new SyntheticCallWithDefaultsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticCallWithDefaultsConfigWithDefaults

`func NewSyntheticCallWithDefaultsConfigWithDefaults() *SyntheticCallWithDefaultsConfig`

NewSyntheticCallWithDefaultsConfigWithDefaults instantiates a new SyntheticCallWithDefaultsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomRules

`func (o *SyntheticCallWithDefaultsConfig) GetCustomRules() []SyntheticCallRule`

GetCustomRules returns the CustomRules field if non-nil, zero value otherwise.

### GetCustomRulesOk

`func (o *SyntheticCallWithDefaultsConfig) GetCustomRulesOk() (*[]SyntheticCallRule, bool)`

GetCustomRulesOk returns a tuple with the CustomRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRules

`func (o *SyntheticCallWithDefaultsConfig) SetCustomRules(v []SyntheticCallRule)`

SetCustomRules sets CustomRules field to given value.


### GetDefaultRules

`func (o *SyntheticCallWithDefaultsConfig) GetDefaultRules() []SyntheticCallRule`

GetDefaultRules returns the DefaultRules field if non-nil, zero value otherwise.

### GetDefaultRulesOk

`func (o *SyntheticCallWithDefaultsConfig) GetDefaultRulesOk() (*[]SyntheticCallRule, bool)`

GetDefaultRulesOk returns a tuple with the DefaultRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRules

`func (o *SyntheticCallWithDefaultsConfig) SetDefaultRules(v []SyntheticCallRule)`

SetDefaultRules sets DefaultRules field to given value.


### GetDefaultRulesEnabled

`func (o *SyntheticCallWithDefaultsConfig) GetDefaultRulesEnabled() bool`

GetDefaultRulesEnabled returns the DefaultRulesEnabled field if non-nil, zero value otherwise.

### GetDefaultRulesEnabledOk

`func (o *SyntheticCallWithDefaultsConfig) GetDefaultRulesEnabledOk() (*bool, bool)`

GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRulesEnabled

`func (o *SyntheticCallWithDefaultsConfig) SetDefaultRulesEnabled(v bool)`

SetDefaultRulesEnabled sets DefaultRulesEnabled field to given value.

### HasDefaultRulesEnabled

`func (o *SyntheticCallWithDefaultsConfig) HasDefaultRulesEnabled() bool`

HasDefaultRulesEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


