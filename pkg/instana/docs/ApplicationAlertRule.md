# ApplicationAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** |  | [optional] 
**AlertType** | **string** |  | 
**MetricName** | **string** |  | 
**StableHash** | Pointer to **int32** |  | [optional] 

## Methods

### NewApplicationAlertRule

`func NewApplicationAlertRule(alertType string, metricName string, ) *ApplicationAlertRule`

NewApplicationAlertRule instantiates a new ApplicationAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAlertRuleWithDefaults

`func NewApplicationAlertRuleWithDefaults() *ApplicationAlertRule`

NewApplicationAlertRuleWithDefaults instantiates a new ApplicationAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *ApplicationAlertRule) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *ApplicationAlertRule) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *ApplicationAlertRule) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *ApplicationAlertRule) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetAlertType

`func (o *ApplicationAlertRule) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *ApplicationAlertRule) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *ApplicationAlertRule) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.


### GetMetricName

`func (o *ApplicationAlertRule) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ApplicationAlertRule) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ApplicationAlertRule) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetStableHash

`func (o *ApplicationAlertRule) GetStableHash() int32`

GetStableHash returns the StableHash field if non-nil, zero value otherwise.

### GetStableHashOk

`func (o *ApplicationAlertRule) GetStableHashOk() (*int32, bool)`

GetStableHashOk returns a tuple with the StableHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableHash

`func (o *ApplicationAlertRule) SetStableHash(v int32)`

SetStableHash sets StableHash field to given value.

### HasStableHash

`func (o *ApplicationAlertRule) HasStableHash() bool`

HasStableHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


