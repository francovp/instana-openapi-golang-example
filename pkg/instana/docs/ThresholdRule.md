# ThresholdRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** |  | [optional] 
**ConditionOperator** | **string** |  | 
**ConditionValue** | Pointer to **float64** |  | [optional] 
**MetricName** | Pointer to **string** |  | [optional] 
**MetricPattern** | Pointer to [**MetricPattern**](MetricPattern.md) |  | [optional] 
**Rollup** | Pointer to **int64** |  | [optional] 
**Window** | Pointer to **int64** |  | [optional] 

## Methods

### NewThresholdRule

`func NewThresholdRule(conditionOperator string, ) *ThresholdRule`

NewThresholdRule instantiates a new ThresholdRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdRuleWithDefaults

`func NewThresholdRuleWithDefaults() *ThresholdRule`

NewThresholdRuleWithDefaults instantiates a new ThresholdRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *ThresholdRule) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *ThresholdRule) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *ThresholdRule) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *ThresholdRule) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetConditionOperator

`func (o *ThresholdRule) GetConditionOperator() string`

GetConditionOperator returns the ConditionOperator field if non-nil, zero value otherwise.

### GetConditionOperatorOk

`func (o *ThresholdRule) GetConditionOperatorOk() (*string, bool)`

GetConditionOperatorOk returns a tuple with the ConditionOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionOperator

`func (o *ThresholdRule) SetConditionOperator(v string)`

SetConditionOperator sets ConditionOperator field to given value.


### GetConditionValue

`func (o *ThresholdRule) GetConditionValue() float64`

GetConditionValue returns the ConditionValue field if non-nil, zero value otherwise.

### GetConditionValueOk

`func (o *ThresholdRule) GetConditionValueOk() (*float64, bool)`

GetConditionValueOk returns a tuple with the ConditionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionValue

`func (o *ThresholdRule) SetConditionValue(v float64)`

SetConditionValue sets ConditionValue field to given value.

### HasConditionValue

`func (o *ThresholdRule) HasConditionValue() bool`

HasConditionValue returns a boolean if a field has been set.

### GetMetricName

`func (o *ThresholdRule) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ThresholdRule) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ThresholdRule) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *ThresholdRule) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMetricPattern

`func (o *ThresholdRule) GetMetricPattern() MetricPattern`

GetMetricPattern returns the MetricPattern field if non-nil, zero value otherwise.

### GetMetricPatternOk

`func (o *ThresholdRule) GetMetricPatternOk() (*MetricPattern, bool)`

GetMetricPatternOk returns a tuple with the MetricPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricPattern

`func (o *ThresholdRule) SetMetricPattern(v MetricPattern)`

SetMetricPattern sets MetricPattern field to given value.

### HasMetricPattern

`func (o *ThresholdRule) HasMetricPattern() bool`

HasMetricPattern returns a boolean if a field has been set.

### GetRollup

`func (o *ThresholdRule) GetRollup() int64`

GetRollup returns the Rollup field if non-nil, zero value otherwise.

### GetRollupOk

`func (o *ThresholdRule) GetRollupOk() (*int64, bool)`

GetRollupOk returns a tuple with the Rollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollup

`func (o *ThresholdRule) SetRollup(v int64)`

SetRollup sets Rollup field to given value.

### HasRollup

`func (o *ThresholdRule) HasRollup() bool`

HasRollup returns a boolean if a field has been set.

### GetWindow

`func (o *ThresholdRule) GetWindow() int64`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ThresholdRule) GetWindowOk() (*int64, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ThresholdRule) SetWindow(v int64)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *ThresholdRule) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


