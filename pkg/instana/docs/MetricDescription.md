# MetricDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | **[]string** |  | 
**DefaultAggregation** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Formatter** | **string** |  | 
**Label** | **string** |  | 
**MetricId** | **string** |  | 

## Methods

### NewMetricDescription

`func NewMetricDescription(aggregations []string, formatter string, label string, metricId string, ) *MetricDescription`

NewMetricDescription instantiates a new MetricDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDescriptionWithDefaults

`func NewMetricDescriptionWithDefaults() *MetricDescription`

NewMetricDescriptionWithDefaults instantiates a new MetricDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *MetricDescription) GetAggregations() []string`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *MetricDescription) GetAggregationsOk() (*[]string, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *MetricDescription) SetAggregations(v []string)`

SetAggregations sets Aggregations field to given value.


### GetDefaultAggregation

`func (o *MetricDescription) GetDefaultAggregation() string`

GetDefaultAggregation returns the DefaultAggregation field if non-nil, zero value otherwise.

### GetDefaultAggregationOk

`func (o *MetricDescription) GetDefaultAggregationOk() (*string, bool)`

GetDefaultAggregationOk returns a tuple with the DefaultAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAggregation

`func (o *MetricDescription) SetDefaultAggregation(v string)`

SetDefaultAggregation sets DefaultAggregation field to given value.

### HasDefaultAggregation

`func (o *MetricDescription) HasDefaultAggregation() bool`

HasDefaultAggregation returns a boolean if a field has been set.

### GetDescription

`func (o *MetricDescription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricDescription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricDescription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricDescription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormatter

`func (o *MetricDescription) GetFormatter() string`

GetFormatter returns the Formatter field if non-nil, zero value otherwise.

### GetFormatterOk

`func (o *MetricDescription) GetFormatterOk() (*string, bool)`

GetFormatterOk returns a tuple with the Formatter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatter

`func (o *MetricDescription) SetFormatter(v string)`

SetFormatter sets Formatter field to given value.


### GetLabel

`func (o *MetricDescription) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MetricDescription) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MetricDescription) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMetricId

`func (o *MetricDescription) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricDescription) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricDescription) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


