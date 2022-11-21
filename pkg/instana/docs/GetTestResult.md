# GetTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **[]string** |  | [optional] 
**Metrics** | [**[]SyntheticMetricConfiguration**](SyntheticMetricConfiguration.md) |  | 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**TagFilterExpression** | Pointer to [**TagFilterExpression**](TagFilterExpression.md) |  | [optional] 
**TagFilters** | Pointer to [**[]TagFilter**](TagFilter.md) |  | [optional] 
**TestId** | Pointer to **string** |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 

## Methods

### NewGetTestResult

`func NewGetTestResult(metrics []SyntheticMetricConfiguration, ) *GetTestResult`

NewGetTestResult instantiates a new GetTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTestResultWithDefaults

`func NewGetTestResultWithDefaults() *GetTestResult`

NewGetTestResultWithDefaults instantiates a new GetTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *GetTestResult) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GetTestResult) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GetTestResult) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *GetTestResult) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetLocationId

`func (o *GetTestResult) GetLocationId() []string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *GetTestResult) GetLocationIdOk() (*[]string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *GetTestResult) SetLocationId(v []string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *GetTestResult) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetMetrics

`func (o *GetTestResult) GetMetrics() []SyntheticMetricConfiguration`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetTestResult) GetMetricsOk() (*[]SyntheticMetricConfiguration, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetTestResult) SetMetrics(v []SyntheticMetricConfiguration)`

SetMetrics sets Metrics field to given value.


### GetOrder

`func (o *GetTestResult) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetTestResult) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetTestResult) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetTestResult) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPagination

`func (o *GetTestResult) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetTestResult) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetTestResult) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetTestResult) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetServiceId

`func (o *GetTestResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GetTestResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GetTestResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *GetTestResult) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetTagFilterExpression

`func (o *GetTestResult) GetTagFilterExpression() TagFilterExpression`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *GetTestResult) GetTagFilterExpressionOk() (*TagFilterExpression, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *GetTestResult) SetTagFilterExpression(v TagFilterExpression)`

SetTagFilterExpression sets TagFilterExpression field to given value.

### HasTagFilterExpression

`func (o *GetTestResult) HasTagFilterExpression() bool`

HasTagFilterExpression returns a boolean if a field has been set.

### GetTagFilters

`func (o *GetTestResult) GetTagFilters() []TagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *GetTestResult) GetTagFiltersOk() (*[]TagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *GetTestResult) SetTagFilters(v []TagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *GetTestResult) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetTestId

`func (o *GetTestResult) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *GetTestResult) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *GetTestResult) SetTestId(v string)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *GetTestResult) HasTestId() bool`

HasTestId returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetTestResult) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetTestResult) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetTestResult) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetTestResult) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


