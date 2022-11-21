# NewApplicationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRules** | [**[]AccessRule**](AccessRule.md) |  | 
**BoundaryScope** | **string** |  | 
**Label** | **string** |  | 
**MatchSpecification** | Pointer to [**MatchExpressionDTO**](MatchExpressionDTO.md) |  | [optional] 
**Scope** | **string** |  | 
**TagFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 

## Methods

### NewNewApplicationConfig

`func NewNewApplicationConfig(accessRules []AccessRule, boundaryScope string, label string, scope string, ) *NewApplicationConfig`

NewNewApplicationConfig instantiates a new NewApplicationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewApplicationConfigWithDefaults

`func NewNewApplicationConfigWithDefaults() *NewApplicationConfig`

NewNewApplicationConfigWithDefaults instantiates a new NewApplicationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRules

`func (o *NewApplicationConfig) GetAccessRules() []AccessRule`

GetAccessRules returns the AccessRules field if non-nil, zero value otherwise.

### GetAccessRulesOk

`func (o *NewApplicationConfig) GetAccessRulesOk() (*[]AccessRule, bool)`

GetAccessRulesOk returns a tuple with the AccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRules

`func (o *NewApplicationConfig) SetAccessRules(v []AccessRule)`

SetAccessRules sets AccessRules field to given value.


### GetBoundaryScope

`func (o *NewApplicationConfig) GetBoundaryScope() string`

GetBoundaryScope returns the BoundaryScope field if non-nil, zero value otherwise.

### GetBoundaryScopeOk

`func (o *NewApplicationConfig) GetBoundaryScopeOk() (*string, bool)`

GetBoundaryScopeOk returns a tuple with the BoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaryScope

`func (o *NewApplicationConfig) SetBoundaryScope(v string)`

SetBoundaryScope sets BoundaryScope field to given value.


### GetLabel

`func (o *NewApplicationConfig) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *NewApplicationConfig) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *NewApplicationConfig) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMatchSpecification

`func (o *NewApplicationConfig) GetMatchSpecification() MatchExpressionDTO`

GetMatchSpecification returns the MatchSpecification field if non-nil, zero value otherwise.

### GetMatchSpecificationOk

`func (o *NewApplicationConfig) GetMatchSpecificationOk() (*MatchExpressionDTO, bool)`

GetMatchSpecificationOk returns a tuple with the MatchSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSpecification

`func (o *NewApplicationConfig) SetMatchSpecification(v MatchExpressionDTO)`

SetMatchSpecification sets MatchSpecification field to given value.

### HasMatchSpecification

`func (o *NewApplicationConfig) HasMatchSpecification() bool`

HasMatchSpecification returns a boolean if a field has been set.

### GetScope

`func (o *NewApplicationConfig) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *NewApplicationConfig) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *NewApplicationConfig) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTagFilterExpression

`func (o *NewApplicationConfig) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *NewApplicationConfig) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *NewApplicationConfig) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.

### HasTagFilterExpression

`func (o *NewApplicationConfig) HasTagFilterExpression() bool`

HasTagFilterExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


