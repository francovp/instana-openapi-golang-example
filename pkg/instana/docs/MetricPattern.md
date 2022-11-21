# MetricPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** |  | 
**Placeholder** | Pointer to **string** |  | [optional] 
**Postfix** | Pointer to **string** |  | [optional] 
**Prefix** | **string** |  | 

## Methods

### NewMetricPattern

`func NewMetricPattern(operator string, prefix string, ) *MetricPattern`

NewMetricPattern instantiates a new MetricPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricPatternWithDefaults

`func NewMetricPatternWithDefaults() *MetricPattern`

NewMetricPatternWithDefaults instantiates a new MetricPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *MetricPattern) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MetricPattern) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MetricPattern) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetPlaceholder

`func (o *MetricPattern) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *MetricPattern) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *MetricPattern) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *MetricPattern) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetPostfix

`func (o *MetricPattern) GetPostfix() string`

GetPostfix returns the Postfix field if non-nil, zero value otherwise.

### GetPostfixOk

`func (o *MetricPattern) GetPostfixOk() (*string, bool)`

GetPostfixOk returns a tuple with the Postfix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostfix

`func (o *MetricPattern) SetPostfix(v string)`

SetPostfix sets Postfix field to given value.

### HasPostfix

`func (o *MetricPattern) HasPostfix() bool`

HasPostfix returns a boolean if a field has been set.

### GetPrefix

`func (o *MetricPattern) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *MetricPattern) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *MetricPattern) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


