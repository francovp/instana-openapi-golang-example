# GetMobileAppMetricsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**[]MobileAppMonitoringMetricsConfiguration**](MobileAppMonitoringMetricsConfiguration.md) |  | 
**TagFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewGetMobileAppMetricsV2

`func NewGetMobileAppMetricsV2(metrics []MobileAppMonitoringMetricsConfiguration, type_ string, ) *GetMobileAppMetricsV2`

NewGetMobileAppMetricsV2 instantiates a new GetMobileAppMetricsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMobileAppMetricsV2WithDefaults

`func NewGetMobileAppMetricsV2WithDefaults() *GetMobileAppMetricsV2`

NewGetMobileAppMetricsV2WithDefaults instantiates a new GetMobileAppMetricsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *GetMobileAppMetricsV2) GetMetrics() []MobileAppMonitoringMetricsConfiguration`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetMobileAppMetricsV2) GetMetricsOk() (*[]MobileAppMonitoringMetricsConfiguration, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetMobileAppMetricsV2) SetMetrics(v []MobileAppMonitoringMetricsConfiguration)`

SetMetrics sets Metrics field to given value.


### GetTagFilterExpression

`func (o *GetMobileAppMetricsV2) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *GetMobileAppMetricsV2) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *GetMobileAppMetricsV2) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.

### HasTagFilterExpression

`func (o *GetMobileAppMetricsV2) HasTagFilterExpression() bool`

HasTagFilterExpression returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetMobileAppMetricsV2) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetMobileAppMetricsV2) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetMobileAppMetricsV2) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetMobileAppMetricsV2) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetType

`func (o *GetMobileAppMetricsV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMobileAppMetricsV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMobileAppMetricsV2) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


