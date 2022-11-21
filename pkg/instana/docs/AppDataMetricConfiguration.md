# AppDataMetricConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | **string** |  | 
**Granularity** | Pointer to **int32** |  | [optional] 
**Metric** | **string** |  | 
**NumeratorTagFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 

## Methods

### NewAppDataMetricConfiguration

`func NewAppDataMetricConfiguration(aggregation string, metric string, ) *AppDataMetricConfiguration`

NewAppDataMetricConfiguration instantiates a new AppDataMetricConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDataMetricConfigurationWithDefaults

`func NewAppDataMetricConfigurationWithDefaults() *AppDataMetricConfiguration`

NewAppDataMetricConfigurationWithDefaults instantiates a new AppDataMetricConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *AppDataMetricConfiguration) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *AppDataMetricConfiguration) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *AppDataMetricConfiguration) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetGranularity

`func (o *AppDataMetricConfiguration) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *AppDataMetricConfiguration) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *AppDataMetricConfiguration) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *AppDataMetricConfiguration) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetMetric

`func (o *AppDataMetricConfiguration) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AppDataMetricConfiguration) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AppDataMetricConfiguration) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetNumeratorTagFilterExpression

`func (o *AppDataMetricConfiguration) GetNumeratorTagFilterExpression() TagFilterExpressionElement`

GetNumeratorTagFilterExpression returns the NumeratorTagFilterExpression field if non-nil, zero value otherwise.

### GetNumeratorTagFilterExpressionOk

`func (o *AppDataMetricConfiguration) GetNumeratorTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetNumeratorTagFilterExpressionOk returns a tuple with the NumeratorTagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeratorTagFilterExpression

`func (o *AppDataMetricConfiguration) SetNumeratorTagFilterExpression(v TagFilterExpressionElement)`

SetNumeratorTagFilterExpression sets NumeratorTagFilterExpression field to given value.

### HasNumeratorTagFilterExpression

`func (o *AppDataMetricConfiguration) HasNumeratorTagFilterExpression() bool`

HasNumeratorTagFilterExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


