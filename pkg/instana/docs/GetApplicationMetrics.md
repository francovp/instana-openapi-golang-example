# GetApplicationMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeInternal** | Pointer to **bool** |  | [optional] 
**IncludeSynthetic** | Pointer to **bool** |  | [optional] 
**Metrics** | [**[]AppDataMetricConfiguration**](AppDataMetricConfiguration.md) |  | 
**TagFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**TimeFrame** | [**TimeFrame**](TimeFrame.md) |  | 

## Methods

### NewGetApplicationMetrics

`func NewGetApplicationMetrics(metrics []AppDataMetricConfiguration, timeFrame TimeFrame, ) *GetApplicationMetrics`

NewGetApplicationMetrics instantiates a new GetApplicationMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApplicationMetricsWithDefaults

`func NewGetApplicationMetricsWithDefaults() *GetApplicationMetrics`

NewGetApplicationMetricsWithDefaults instantiates a new GetApplicationMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeInternal

`func (o *GetApplicationMetrics) GetIncludeInternal() bool`

GetIncludeInternal returns the IncludeInternal field if non-nil, zero value otherwise.

### GetIncludeInternalOk

`func (o *GetApplicationMetrics) GetIncludeInternalOk() (*bool, bool)`

GetIncludeInternalOk returns a tuple with the IncludeInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInternal

`func (o *GetApplicationMetrics) SetIncludeInternal(v bool)`

SetIncludeInternal sets IncludeInternal field to given value.

### HasIncludeInternal

`func (o *GetApplicationMetrics) HasIncludeInternal() bool`

HasIncludeInternal returns a boolean if a field has been set.

### GetIncludeSynthetic

`func (o *GetApplicationMetrics) GetIncludeSynthetic() bool`

GetIncludeSynthetic returns the IncludeSynthetic field if non-nil, zero value otherwise.

### GetIncludeSyntheticOk

`func (o *GetApplicationMetrics) GetIncludeSyntheticOk() (*bool, bool)`

GetIncludeSyntheticOk returns a tuple with the IncludeSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSynthetic

`func (o *GetApplicationMetrics) SetIncludeSynthetic(v bool)`

SetIncludeSynthetic sets IncludeSynthetic field to given value.

### HasIncludeSynthetic

`func (o *GetApplicationMetrics) HasIncludeSynthetic() bool`

HasIncludeSynthetic returns a boolean if a field has been set.

### GetMetrics

`func (o *GetApplicationMetrics) GetMetrics() []AppDataMetricConfiguration`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetApplicationMetrics) GetMetricsOk() (*[]AppDataMetricConfiguration, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetApplicationMetrics) SetMetrics(v []AppDataMetricConfiguration)`

SetMetrics sets Metrics field to given value.


### GetTagFilterExpression

`func (o *GetApplicationMetrics) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *GetApplicationMetrics) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *GetApplicationMetrics) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.

### HasTagFilterExpression

`func (o *GetApplicationMetrics) HasTagFilterExpression() bool`

HasTagFilterExpression returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetApplicationMetrics) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetApplicationMetrics) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetApplicationMetrics) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


