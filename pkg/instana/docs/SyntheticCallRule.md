# SyntheticCallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MatchSpecification** | [**MatchExpressionDTO**](MatchExpressionDTO.md) |  | 
**Name** | **string** |  | 

## Methods

### NewSyntheticCallRule

`func NewSyntheticCallRule(matchSpecification MatchExpressionDTO, name string, ) *SyntheticCallRule`

NewSyntheticCallRule instantiates a new SyntheticCallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticCallRuleWithDefaults

`func NewSyntheticCallRuleWithDefaults() *SyntheticCallRule`

NewSyntheticCallRuleWithDefaults instantiates a new SyntheticCallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SyntheticCallRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyntheticCallRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyntheticCallRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyntheticCallRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SyntheticCallRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SyntheticCallRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SyntheticCallRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SyntheticCallRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMatchSpecification

`func (o *SyntheticCallRule) GetMatchSpecification() MatchExpressionDTO`

GetMatchSpecification returns the MatchSpecification field if non-nil, zero value otherwise.

### GetMatchSpecificationOk

`func (o *SyntheticCallRule) GetMatchSpecificationOk() (*MatchExpressionDTO, bool)`

GetMatchSpecificationOk returns a tuple with the MatchSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSpecification

`func (o *SyntheticCallRule) SetMatchSpecification(v MatchExpressionDTO)`

SetMatchSpecification sets MatchSpecification field to given value.


### GetName

`func (o *SyntheticCallRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticCallRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticCallRule) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


