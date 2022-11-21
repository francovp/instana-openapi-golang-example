# TagFilterExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | [**[]TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | 
**LogicalOperator** | **string** |  | 

## Methods

### NewTagFilterExpression

`func NewTagFilterExpression(elements []TagFilterExpressionElement, logicalOperator string, ) *TagFilterExpression`

NewTagFilterExpression instantiates a new TagFilterExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagFilterExpressionWithDefaults

`func NewTagFilterExpressionWithDefaults() *TagFilterExpression`

NewTagFilterExpressionWithDefaults instantiates a new TagFilterExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *TagFilterExpression) GetElements() []TagFilterExpressionElement`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *TagFilterExpression) GetElementsOk() (*[]TagFilterExpressionElement, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *TagFilterExpression) SetElements(v []TagFilterExpressionElement)`

SetElements sets Elements field to given value.


### GetLogicalOperator

`func (o *TagFilterExpression) GetLogicalOperator() string`

GetLogicalOperator returns the LogicalOperator field if non-nil, zero value otherwise.

### GetLogicalOperatorOk

`func (o *TagFilterExpression) GetLogicalOperatorOk() (*string, bool)`

GetLogicalOperatorOk returns a tuple with the LogicalOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalOperator

`func (o *TagFilterExpression) SetLogicalOperator(v string)`

SetLogicalOperator sets LogicalOperator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


