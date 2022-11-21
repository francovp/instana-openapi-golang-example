# AbstractRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | **string** |  | 
**Severity** | Pointer to **int32** |  | [optional] 

## Methods

### NewAbstractRule

`func NewAbstractRule(ruleType string, ) *AbstractRule`

NewAbstractRule instantiates a new AbstractRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractRuleWithDefaults

`func NewAbstractRuleWithDefaults() *AbstractRule`

NewAbstractRuleWithDefaults instantiates a new AbstractRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleType

`func (o *AbstractRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *AbstractRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *AbstractRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetSeverity

`func (o *AbstractRule) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AbstractRule) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AbstractRule) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AbstractRule) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


