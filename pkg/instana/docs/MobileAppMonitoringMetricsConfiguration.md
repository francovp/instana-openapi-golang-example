# MobileAppMonitoringMetricsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | **string** |  | 
**Granularity** | Pointer to **int32** |  | [optional] 
**Metric** | **string** |  | 

## Methods

### NewMobileAppMonitoringMetricsConfiguration

`func NewMobileAppMonitoringMetricsConfiguration(aggregation string, metric string, ) *MobileAppMonitoringMetricsConfiguration`

NewMobileAppMonitoringMetricsConfiguration instantiates a new MobileAppMonitoringMetricsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileAppMonitoringMetricsConfigurationWithDefaults

`func NewMobileAppMonitoringMetricsConfigurationWithDefaults() *MobileAppMonitoringMetricsConfiguration`

NewMobileAppMonitoringMetricsConfigurationWithDefaults instantiates a new MobileAppMonitoringMetricsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *MobileAppMonitoringMetricsConfiguration) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MobileAppMonitoringMetricsConfiguration) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MobileAppMonitoringMetricsConfiguration) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetGranularity

`func (o *MobileAppMonitoringMetricsConfiguration) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *MobileAppMonitoringMetricsConfiguration) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *MobileAppMonitoringMetricsConfiguration) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *MobileAppMonitoringMetricsConfiguration) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetMetric

`func (o *MobileAppMonitoringMetricsConfiguration) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MobileAppMonitoringMetricsConfiguration) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MobileAppMonitoringMetricsConfiguration) SetMetric(v string)`

SetMetric sets Metric field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


