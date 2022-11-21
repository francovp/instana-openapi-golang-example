# EntityVerificationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchingEntityLabel** | **string** |  | 
**MatchingEntityType** | **string** |  | 
**MatchingOperator** | **string** |  | 
**OfflineDuration** | Pointer to **int64** |  | [optional] 

## Methods

### NewEntityVerificationRule

`func NewEntityVerificationRule(matchingEntityLabel string, matchingEntityType string, matchingOperator string, ) *EntityVerificationRule`

NewEntityVerificationRule instantiates a new EntityVerificationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityVerificationRuleWithDefaults

`func NewEntityVerificationRuleWithDefaults() *EntityVerificationRule`

NewEntityVerificationRuleWithDefaults instantiates a new EntityVerificationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchingEntityLabel

`func (o *EntityVerificationRule) GetMatchingEntityLabel() string`

GetMatchingEntityLabel returns the MatchingEntityLabel field if non-nil, zero value otherwise.

### GetMatchingEntityLabelOk

`func (o *EntityVerificationRule) GetMatchingEntityLabelOk() (*string, bool)`

GetMatchingEntityLabelOk returns a tuple with the MatchingEntityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingEntityLabel

`func (o *EntityVerificationRule) SetMatchingEntityLabel(v string)`

SetMatchingEntityLabel sets MatchingEntityLabel field to given value.


### GetMatchingEntityType

`func (o *EntityVerificationRule) GetMatchingEntityType() string`

GetMatchingEntityType returns the MatchingEntityType field if non-nil, zero value otherwise.

### GetMatchingEntityTypeOk

`func (o *EntityVerificationRule) GetMatchingEntityTypeOk() (*string, bool)`

GetMatchingEntityTypeOk returns a tuple with the MatchingEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingEntityType

`func (o *EntityVerificationRule) SetMatchingEntityType(v string)`

SetMatchingEntityType sets MatchingEntityType field to given value.


### GetMatchingOperator

`func (o *EntityVerificationRule) GetMatchingOperator() string`

GetMatchingOperator returns the MatchingOperator field if non-nil, zero value otherwise.

### GetMatchingOperatorOk

`func (o *EntityVerificationRule) GetMatchingOperatorOk() (*string, bool)`

GetMatchingOperatorOk returns a tuple with the MatchingOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingOperator

`func (o *EntityVerificationRule) SetMatchingOperator(v string)`

SetMatchingOperator sets MatchingOperator field to given value.


### GetOfflineDuration

`func (o *EntityVerificationRule) GetOfflineDuration() int64`

GetOfflineDuration returns the OfflineDuration field if non-nil, zero value otherwise.

### GetOfflineDurationOk

`func (o *EntityVerificationRule) GetOfflineDurationOk() (*int64, bool)`

GetOfflineDurationOk returns a tuple with the OfflineDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineDuration

`func (o *EntityVerificationRule) SetOfflineDuration(v int64)`

SetOfflineDuration sets OfflineDuration field to given value.

### HasOfflineDuration

`func (o *EntityVerificationRule) HasOfflineDuration() bool`

HasOfflineDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


