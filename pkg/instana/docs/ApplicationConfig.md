# ApplicationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRules** | [**[]AccessRule**](AccessRule.md) |  | 
**BoundaryScope** | **string** |  | 
**Id** | **string** |  | 
**Label** | **string** |  | 
**MatchSpecification** | Pointer to [**MatchExpressionDTO**](MatchExpressionDTO.md) |  | [optional] 
**Scope** | **string** |  | 
**TagFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 

## Methods

### NewApplicationConfig

`func NewApplicationConfig(accessRules []AccessRule, boundaryScope string, id string, label string, scope string, ) *ApplicationConfig`

NewApplicationConfig instantiates a new ApplicationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationConfigWithDefaults

`func NewApplicationConfigWithDefaults() *ApplicationConfig`

NewApplicationConfigWithDefaults instantiates a new ApplicationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRules

`func (o *ApplicationConfig) GetAccessRules() []AccessRule`

GetAccessRules returns the AccessRules field if non-nil, zero value otherwise.

### GetAccessRulesOk

`func (o *ApplicationConfig) GetAccessRulesOk() (*[]AccessRule, bool)`

GetAccessRulesOk returns a tuple with the AccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRules

`func (o *ApplicationConfig) SetAccessRules(v []AccessRule)`

SetAccessRules sets AccessRules field to given value.


### GetBoundaryScope

`func (o *ApplicationConfig) GetBoundaryScope() string`

GetBoundaryScope returns the BoundaryScope field if non-nil, zero value otherwise.

### GetBoundaryScopeOk

`func (o *ApplicationConfig) GetBoundaryScopeOk() (*string, bool)`

GetBoundaryScopeOk returns a tuple with the BoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaryScope

`func (o *ApplicationConfig) SetBoundaryScope(v string)`

SetBoundaryScope sets BoundaryScope field to given value.


### GetId

`func (o *ApplicationConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationConfig) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *ApplicationConfig) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApplicationConfig) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApplicationConfig) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMatchSpecification

`func (o *ApplicationConfig) GetMatchSpecification() MatchExpressionDTO`

GetMatchSpecification returns the MatchSpecification field if non-nil, zero value otherwise.

### GetMatchSpecificationOk

`func (o *ApplicationConfig) GetMatchSpecificationOk() (*MatchExpressionDTO, bool)`

GetMatchSpecificationOk returns a tuple with the MatchSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSpecification

`func (o *ApplicationConfig) SetMatchSpecification(v MatchExpressionDTO)`

SetMatchSpecification sets MatchSpecification field to given value.

### HasMatchSpecification

`func (o *ApplicationConfig) HasMatchSpecification() bool`

HasMatchSpecification returns a boolean if a field has been set.

### GetScope

`func (o *ApplicationConfig) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ApplicationConfig) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ApplicationConfig) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTagFilterExpression

`func (o *ApplicationConfig) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *ApplicationConfig) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *ApplicationConfig) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.

### HasTagFilterExpression

`func (o *ApplicationConfig) HasTagFilterExpression() bool`

HasTagFilterExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


