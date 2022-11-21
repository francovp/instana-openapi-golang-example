# ThresholdRuleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** |  | [optional] 
**ConditionOperator** | Pointer to **string** |  | [optional] 
**ConditionValue** | Pointer to **float64** |  | [optional] 
**MetricName** | Pointer to **string** |  | [optional] 
**MetricPattern** | Pointer to [**MetricPattern**](MetricPattern.md) |  | [optional] 
**Rollup** | Pointer to **int64** |  | [optional] 
**Window** | Pointer to **int64** |  | [optional] 

## Methods

### NewThresholdRuleAllOf

`func NewThresholdRuleAllOf() *ThresholdRuleAllOf`

NewThresholdRuleAllOf instantiates a new ThresholdRuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdRuleAllOfWithDefaults

`func NewThresholdRuleAllOfWithDefaults() *ThresholdRuleAllOf`

NewThresholdRuleAllOfWithDefaults instantiates a new ThresholdRuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *ThresholdRuleAllOf) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *ThresholdRuleAllOf) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *ThresholdRuleAllOf) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *ThresholdRuleAllOf) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetConditionOperator

`func (o *ThresholdRuleAllOf) GetConditionOperator() string`

GetConditionOperator returns the ConditionOperator field if non-nil, zero value otherwise.

### GetConditionOperatorOk

`func (o *ThresholdRuleAllOf) GetConditionOperatorOk() (*string, bool)`

GetConditionOperatorOk returns a tuple with the ConditionOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionOperator

`func (o *ThresholdRuleAllOf) SetConditionOperator(v string)`

SetConditionOperator sets ConditionOperator field to given value.

### HasConditionOperator

`func (o *ThresholdRuleAllOf) HasConditionOperator() bool`

HasConditionOperator returns a boolean if a field has been set.

### GetConditionValue

`func (o *ThresholdRuleAllOf) GetConditionValue() float64`

GetConditionValue returns the ConditionValue field if non-nil, zero value otherwise.

### GetConditionValueOk

`func (o *ThresholdRuleAllOf) GetConditionValueOk() (*float64, bool)`

GetConditionValueOk returns a tuple with the ConditionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionValue

`func (o *ThresholdRuleAllOf) SetConditionValue(v float64)`

SetConditionValue sets ConditionValue field to given value.

### HasConditionValue

`func (o *ThresholdRuleAllOf) HasConditionValue() bool`

HasConditionValue returns a boolean if a field has been set.

### GetMetricName

`func (o *ThresholdRuleAllOf) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ThresholdRuleAllOf) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ThresholdRuleAllOf) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *ThresholdRuleAllOf) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMetricPattern

`func (o *ThresholdRuleAllOf) GetMetricPattern() MetricPattern`

GetMetricPattern returns the MetricPattern field if non-nil, zero value otherwise.

### GetMetricPatternOk

`func (o *ThresholdRuleAllOf) GetMetricPatternOk() (*MetricPattern, bool)`

GetMetricPatternOk returns a tuple with the MetricPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricPattern

`func (o *ThresholdRuleAllOf) SetMetricPattern(v MetricPattern)`

SetMetricPattern sets MetricPattern field to given value.

### HasMetricPattern

`func (o *ThresholdRuleAllOf) HasMetricPattern() bool`

HasMetricPattern returns a boolean if a field has been set.

### GetRollup

`func (o *ThresholdRuleAllOf) GetRollup() int64`

GetRollup returns the Rollup field if non-nil, zero value otherwise.

### GetRollupOk

`func (o *ThresholdRuleAllOf) GetRollupOk() (*int64, bool)`

GetRollupOk returns a tuple with the Rollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollup

`func (o *ThresholdRuleAllOf) SetRollup(v int64)`

SetRollup sets Rollup field to given value.

### HasRollup

`func (o *ThresholdRuleAllOf) HasRollup() bool`

HasRollup returns a boolean if a field has been set.

### GetWindow

`func (o *ThresholdRuleAllOf) GetWindow() int64`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ThresholdRuleAllOf) GetWindowOk() (*int64, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ThresholdRuleAllOf) SetWindow(v int64)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *ThresholdRuleAllOf) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


