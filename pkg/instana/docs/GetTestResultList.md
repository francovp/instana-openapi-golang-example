# GetTestResultList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**SyntheticMetrics** | **[]string** |  | 
**TagFilters** | Pointer to [**[]TagFilter**](TagFilter.md) |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 

## Methods

### NewGetTestResultList

`func NewGetTestResultList(syntheticMetrics []string, ) *GetTestResultList`

NewGetTestResultList instantiates a new GetTestResultList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTestResultListWithDefaults

`func NewGetTestResultListWithDefaults() *GetTestResultList`

NewGetTestResultListWithDefaults instantiates a new GetTestResultList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *GetTestResultList) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GetTestResultList) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GetTestResultList) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *GetTestResultList) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetOrder

`func (o *GetTestResultList) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetTestResultList) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetTestResultList) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetTestResultList) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPagination

`func (o *GetTestResultList) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetTestResultList) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetTestResultList) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetTestResultList) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSyntheticMetrics

`func (o *GetTestResultList) GetSyntheticMetrics() []string`

GetSyntheticMetrics returns the SyntheticMetrics field if non-nil, zero value otherwise.

### GetSyntheticMetricsOk

`func (o *GetTestResultList) GetSyntheticMetricsOk() (*[]string, bool)`

GetSyntheticMetricsOk returns a tuple with the SyntheticMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticMetrics

`func (o *GetTestResultList) SetSyntheticMetrics(v []string)`

SetSyntheticMetrics sets SyntheticMetrics field to given value.


### GetTagFilters

`func (o *GetTestResultList) GetTagFilters() []TagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *GetTestResultList) GetTagFiltersOk() (*[]TagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *GetTestResultList) SetTagFilters(v []TagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *GetTestResultList) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetTestResultList) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetTestResultList) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetTestResultList) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetTestResultList) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


