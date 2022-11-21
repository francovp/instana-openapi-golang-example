# TagFilterExpressionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | Pointer to [**[]TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**LogicalOperator** | Pointer to **string** |  | [optional] 

## Methods

### NewTagFilterExpressionAllOf

`func NewTagFilterExpressionAllOf() *TagFilterExpressionAllOf`

NewTagFilterExpressionAllOf instantiates a new TagFilterExpressionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagFilterExpressionAllOfWithDefaults

`func NewTagFilterExpressionAllOfWithDefaults() *TagFilterExpressionAllOf`

NewTagFilterExpressionAllOfWithDefaults instantiates a new TagFilterExpressionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *TagFilterExpressionAllOf) GetElements() []TagFilterExpressionElement`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *TagFilterExpressionAllOf) GetElementsOk() (*[]TagFilterExpressionElement, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *TagFilterExpressionAllOf) SetElements(v []TagFilterExpressionElement)`

SetElements sets Elements field to given value.

### HasElements

`func (o *TagFilterExpressionAllOf) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetLogicalOperator

`func (o *TagFilterExpressionAllOf) GetLogicalOperator() string`

GetLogicalOperator returns the LogicalOperator field if non-nil, zero value otherwise.

### GetLogicalOperatorOk

`func (o *TagFilterExpressionAllOf) GetLogicalOperatorOk() (*string, bool)`

GetLogicalOperatorOk returns a tuple with the LogicalOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalOperator

`func (o *TagFilterExpressionAllOf) SetLogicalOperator(v string)`

SetLogicalOperator sets LogicalOperator field to given value.

### HasLogicalOperator

`func (o *TagFilterExpressionAllOf) HasLogicalOperator() bool`

HasLogicalOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


