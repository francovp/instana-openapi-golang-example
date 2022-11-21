# GetWebsiteBeaconGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**WebsiteBeaconTagGroup**](WebsiteBeaconTagGroup.md) |  | 
**Metrics** | [**[]WebsiteMonitoringMetricsConfiguration**](WebsiteMonitoringMetricsConfiguration.md) |  | 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Pagination** | Pointer to [**CursorPagination**](CursorPagination.md) |  | [optional] 
**TagFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**TagFilters** | Pointer to [**[]DeprecatedTagFilter**](DeprecatedTagFilter.md) |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewGetWebsiteBeaconGroups

`func NewGetWebsiteBeaconGroups(group WebsiteBeaconTagGroup, metrics []WebsiteMonitoringMetricsConfiguration, type_ string, ) *GetWebsiteBeaconGroups`

NewGetWebsiteBeaconGroups instantiates a new GetWebsiteBeaconGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebsiteBeaconGroupsWithDefaults

`func NewGetWebsiteBeaconGroupsWithDefaults() *GetWebsiteBeaconGroups`

NewGetWebsiteBeaconGroupsWithDefaults instantiates a new GetWebsiteBeaconGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *GetWebsiteBeaconGroups) GetGroup() WebsiteBeaconTagGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetWebsiteBeaconGroups) GetGroupOk() (*WebsiteBeaconTagGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetWebsiteBeaconGroups) SetGroup(v WebsiteBeaconTagGroup)`

SetGroup sets Group field to given value.


### GetMetrics

`func (o *GetWebsiteBeaconGroups) GetMetrics() []WebsiteMonitoringMetricsConfiguration`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetWebsiteBeaconGroups) GetMetricsOk() (*[]WebsiteMonitoringMetricsConfiguration, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetWebsiteBeaconGroups) SetMetrics(v []WebsiteMonitoringMetricsConfiguration)`

SetMetrics sets Metrics field to given value.


### GetOrder

`func (o *GetWebsiteBeaconGroups) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetWebsiteBeaconGroups) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetWebsiteBeaconGroups) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetWebsiteBeaconGroups) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPagination

`func (o *GetWebsiteBeaconGroups) GetPagination() CursorPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetWebsiteBeaconGroups) GetPaginationOk() (*CursorPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetWebsiteBeaconGroups) SetPagination(v CursorPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetWebsiteBeaconGroups) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetTagFilterExpression

`func (o *GetWebsiteBeaconGroups) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *GetWebsiteBeaconGroups) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *GetWebsiteBeaconGroups) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.

### HasTagFilterExpression

`func (o *GetWebsiteBeaconGroups) HasTagFilterExpression() bool`

HasTagFilterExpression returns a boolean if a field has been set.

### GetTagFilters

`func (o *GetWebsiteBeaconGroups) GetTagFilters() []DeprecatedTagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *GetWebsiteBeaconGroups) GetTagFiltersOk() (*[]DeprecatedTagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *GetWebsiteBeaconGroups) SetTagFilters(v []DeprecatedTagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *GetWebsiteBeaconGroups) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetWebsiteBeaconGroups) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetWebsiteBeaconGroups) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetWebsiteBeaconGroups) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetWebsiteBeaconGroups) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetType

`func (o *GetWebsiteBeaconGroups) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetWebsiteBeaconGroups) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetWebsiteBeaconGroups) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


