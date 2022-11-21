# MetricConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricAggregation** | Pointer to **string** |  | [optional] 
**MetricName** | **string** |  | 
**Threshold** | Pointer to **float64** |  | [optional] 

## Methods

### NewMetricConfiguration

`func NewMetricConfiguration(metricName string, ) *MetricConfiguration`

NewMetricConfiguration instantiates a new MetricConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricConfigurationWithDefaults

`func NewMetricConfigurationWithDefaults() *MetricConfiguration`

NewMetricConfigurationWithDefaults instantiates a new MetricConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricAggregation

`func (o *MetricConfiguration) GetMetricAggregation() string`

GetMetricAggregation returns the MetricAggregation field if non-nil, zero value otherwise.

### GetMetricAggregationOk

`func (o *MetricConfiguration) GetMetricAggregationOk() (*string, bool)`

GetMetricAggregationOk returns a tuple with the MetricAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricAggregation

`func (o *MetricConfiguration) SetMetricAggregation(v string)`

SetMetricAggregation sets MetricAggregation field to given value.

### HasMetricAggregation

`func (o *MetricConfiguration) HasMetricAggregation() bool`

HasMetricAggregation returns a boolean if a field has been set.

### GetMetricName

`func (o *MetricConfiguration) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricConfiguration) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricConfiguration) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetThreshold

`func (o *MetricConfiguration) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MetricConfiguration) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MetricConfiguration) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *MetricConfiguration) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


