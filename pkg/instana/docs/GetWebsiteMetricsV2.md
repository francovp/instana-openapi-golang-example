# GetWebsiteMetricsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**[]WebsiteMonitoringMetricsConfiguration**](WebsiteMonitoringMetricsConfiguration.md) |  | 
**TagFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewGetWebsiteMetricsV2

`func NewGetWebsiteMetricsV2(metrics []WebsiteMonitoringMetricsConfiguration, type_ string, ) *GetWebsiteMetricsV2`

NewGetWebsiteMetricsV2 instantiates a new GetWebsiteMetricsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebsiteMetricsV2WithDefaults

`func NewGetWebsiteMetricsV2WithDefaults() *GetWebsiteMetricsV2`

NewGetWebsiteMetricsV2WithDefaults instantiates a new GetWebsiteMetricsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *GetWebsiteMetricsV2) GetMetrics() []WebsiteMonitoringMetricsConfiguration`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetWebsiteMetricsV2) GetMetricsOk() (*[]WebsiteMonitoringMetricsConfiguration, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetWebsiteMetricsV2) SetMetrics(v []WebsiteMonitoringMetricsConfiguration)`

SetMetrics sets Metrics field to given value.


### GetTagFilterExpression

`func (o *GetWebsiteMetricsV2) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *GetWebsiteMetricsV2) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *GetWebsiteMetricsV2) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.

### HasTagFilterExpression

`func (o *GetWebsiteMetricsV2) HasTagFilterExpression() bool`

HasTagFilterExpression returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetWebsiteMetricsV2) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetWebsiteMetricsV2) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetWebsiteMetricsV2) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetWebsiteMetricsV2) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetType

`func (o *GetWebsiteMetricsV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetWebsiteMetricsV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetWebsiteMetricsV2) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


