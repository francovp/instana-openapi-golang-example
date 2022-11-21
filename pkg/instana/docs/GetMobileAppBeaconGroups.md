# GetMobileAppBeaconGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**MobileAppBeaconTagGroup**](MobileAppBeaconTagGroup.md) |  | 
**Metrics** | [**[]MobileAppMonitoringMetricsConfiguration**](MobileAppMonitoringMetricsConfiguration.md) |  | 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Pagination** | Pointer to [**CursorPagination**](CursorPagination.md) |  | [optional] 
**TagFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**TagFilters** | Pointer to [**[]DeprecatedTagFilter**](DeprecatedTagFilter.md) |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewGetMobileAppBeaconGroups

`func NewGetMobileAppBeaconGroups(group MobileAppBeaconTagGroup, metrics []MobileAppMonitoringMetricsConfiguration, type_ string, ) *GetMobileAppBeaconGroups`

NewGetMobileAppBeaconGroups instantiates a new GetMobileAppBeaconGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMobileAppBeaconGroupsWithDefaults

`func NewGetMobileAppBeaconGroupsWithDefaults() *GetMobileAppBeaconGroups`

NewGetMobileAppBeaconGroupsWithDefaults instantiates a new GetMobileAppBeaconGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *GetMobileAppBeaconGroups) GetGroup() MobileAppBeaconTagGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetMobileAppBeaconGroups) GetGroupOk() (*MobileAppBeaconTagGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetMobileAppBeaconGroups) SetGroup(v MobileAppBeaconTagGroup)`

SetGroup sets Group field to given value.


### GetMetrics

`func (o *GetMobileAppBeaconGroups) GetMetrics() []MobileAppMonitoringMetricsConfiguration`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetMobileAppBeaconGroups) GetMetricsOk() (*[]MobileAppMonitoringMetricsConfiguration, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetMobileAppBeaconGroups) SetMetrics(v []MobileAppMonitoringMetricsConfiguration)`

SetMetrics sets Metrics field to given value.


### GetOrder

`func (o *GetMobileAppBeaconGroups) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetMobileAppBeaconGroups) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetMobileAppBeaconGroups) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetMobileAppBeaconGroups) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPagination

`func (o *GetMobileAppBeaconGroups) GetPagination() CursorPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetMobileAppBeaconGroups) GetPaginationOk() (*CursorPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetMobileAppBeaconGroups) SetPagination(v CursorPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetMobileAppBeaconGroups) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetTagFilterExpression

`func (o *GetMobileAppBeaconGroups) GetTagFilterExpression() TagFilterExpressionElement`

GetTagFilterExpression returns the TagFilterExpression field if non-nil, zero value otherwise.

### GetTagFilterExpressionOk

`func (o *GetMobileAppBeaconGroups) GetTagFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetTagFilterExpressionOk returns a tuple with the TagFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterExpression

`func (o *GetMobileAppBeaconGroups) SetTagFilterExpression(v TagFilterExpressionElement)`

SetTagFilterExpression sets TagFilterExpression field to given value.

### HasTagFilterExpression

`func (o *GetMobileAppBeaconGroups) HasTagFilterExpression() bool`

HasTagFilterExpression returns a boolean if a field has been set.

### GetTagFilters

`func (o *GetMobileAppBeaconGroups) GetTagFilters() []DeprecatedTagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *GetMobileAppBeaconGroups) GetTagFiltersOk() (*[]DeprecatedTagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *GetMobileAppBeaconGroups) SetTagFilters(v []DeprecatedTagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *GetMobileAppBeaconGroups) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetMobileAppBeaconGroups) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetMobileAppBeaconGroups) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetMobileAppBeaconGroups) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetMobileAppBeaconGroups) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetType

`func (o *GetMobileAppBeaconGroups) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMobileAppBeaconGroups) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMobileAppBeaconGroups) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


