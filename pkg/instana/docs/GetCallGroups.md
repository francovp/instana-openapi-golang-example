# GetCallGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**Group**](Group.md) |  | 
**IncludeInternal** | Pointer to **bool** |  | [optional] 
**IncludeSynthetic** | Pointer to **bool** |  | [optional] 
**Metrics** | [**[]MetricConfig**](MetricConfig.md) |  | 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Pagination** | Pointer to [**CursorPagination**](CursorPagination.md) |  | [optional] 
**TagFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**TagFilters** | Pointer to [**[]DeprecatedTagFilter**](DeprecatedTagFilter.md) |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 

## Methods

### NewGetCallGroups

`func NewGetCallGroups(group Group, metrics []MetricConfig, ) *GetCallGroups`

NewGetCallGroups instantiates a new GetCallGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCallGroupsWithDefaults

`func NewGetCallGroupsWithDefaults() *GetCallGroups`

NewGetCallGroupsWithDefaults instantiates a new GetCallGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *GetCallGroups) GetGroup() Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetCallGroups) GetGroupOk() (*Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetCallGroups) SetGroup(v Group)`

SetGroup sets Group field to given value.


### GetIncludeInternal

`func (o *GetCallGroups) GetIncludeInternal() bool`

GetIncludeInternal returns the IncludeInternal field if non-nil, zero value otherwise.

### GetIncludeInternalOk

`func (o *GetCallGroups) GetIncludeInternalOk() (*bool, bool)`

GetIncludeInternalOk returns a tuple with the IncludeInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInternal

`func (o *GetCallGroups) SetIncludeInternal(v bool)`

SetIncludeInternal sets IncludeInternal field to given value.

### HasIncludeInternal

`func (o *GetCallGroups) HasIncludeInternal() bool`

HasIncludeInternal returns a boolean if a field has been set.

### GetIncludeSynthetic

`func (o *GetCallGroups) GetIncludeSynthetic() bool`

GetIncludeSynthetic returns the IncludeSynthetic field if non-nil, zero value otherwise.

### GetIncludeSyntheticOk

`func (o *GetCallGroups) GetIncludeSyntheticOk() (*bool, bool)`

GetIncludeSyntheticOk returns a tuple with the IncludeSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSynthetic

`func (o *GetCallGroups) SetIncludeSynthetic(v bool)`

SetIncludeSynthetic sets IncludeSynthetic field to given value.

### HasIncludeSynthetic

`func (o *GetCallGroups) HasIncludeSynthetic() bool`

HasIncludeSynthetic returns a boolean if a field has been set.

### GetMetrics

`func (o *GetCallGroups) GetMetrics() []MetricConfig`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetCallGroups) GetMetricsOk() (*[]MetricConfig, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetCallGroups) SetMetrics(v []MetricConfig)`

SetMetrics sets Metrics field to given value.


### GetOrder

`func (o *GetCallGroups) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetCallGroups) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetCallGroups) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetCallGroups) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPagination

`func (o *GetCallGroups) GetPagination() CursorPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetCallGroups) GetPaginationOk() (*CursorPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetCallGroups) SetPagination(v CursorPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetCallGroups) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetTagFilterExpression

`func (o *GetCallGroups) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *GetCallGroups) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *GetCallGroups) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.

### HasTagFilterExpression

`func (o *GetCallGroups) HasTagFilterExpression() bool`

HasTagFilterExpression returns a boolean if a field has been set.

### GetTagFilters

`func (o *GetCallGroups) GetTagFilters() []DeprecatedTagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *GetCallGroups) GetTagFiltersOk() (*[]DeprecatedTagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *GetCallGroups) SetTagFilters(v []DeprecatedTagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *GetCallGroups) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetCallGroups) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetCallGroups) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetCallGroups) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetCallGroups) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


