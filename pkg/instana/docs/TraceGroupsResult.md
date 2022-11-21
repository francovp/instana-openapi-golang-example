# TraceGroupsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustedTimeframe** | Pointer to [**AdjustedTimeframe**](AdjustedTimeframe.md) |  | [optional] 
**CanLoadMore** | Pointer to **bool** |  | [optional] 
**Items** | [**[]TraceGroupsItem**](TraceGroupsItem.md) |  | 
**TotalHits** | Pointer to **int64** |  | [optional] 
**TotalRepresentedItemCount** | Pointer to **int64** |  | [optional] 
**TotalRetainedItemCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewTraceGroupsResult

`func NewTraceGroupsResult(items []TraceGroupsItem, ) *TraceGroupsResult`

NewTraceGroupsResult instantiates a new TraceGroupsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceGroupsResultWithDefaults

`func NewTraceGroupsResultWithDefaults() *TraceGroupsResult`

NewTraceGroupsResultWithDefaults instantiates a new TraceGroupsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustedTimeframe

`func (o *TraceGroupsResult) GetAdjustedTimeframe() AdjustedTimeframe`

GetAdjustedTimeframe returns the AdjustedTimeframe field if non-nil, zero value otherwise.

### GetAdjustedTimeframeOk

`func (o *TraceGroupsResult) GetAdjustedTimeframeOk() (*AdjustedTimeframe, bool)`

GetAdjustedTimeframeOk returns a tuple with the AdjustedTimeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedTimeframe

`func (o *TraceGroupsResult) SetAdjustedTimeframe(v AdjustedTimeframe)`

SetAdjustedTimeframe sets AdjustedTimeframe field to given value.

### HasAdjustedTimeframe

`func (o *TraceGroupsResult) HasAdjustedTimeframe() bool`

HasAdjustedTimeframe returns a boolean if a field has been set.

### GetCanLoadMore

`func (o *TraceGroupsResult) GetCanLoadMore() bool`

GetCanLoadMore returns the CanLoadMore field if non-nil, zero value otherwise.

### GetCanLoadMoreOk

`func (o *TraceGroupsResult) GetCanLoadMoreOk() (*bool, bool)`

GetCanLoadMoreOk returns a tuple with the CanLoadMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLoadMore

`func (o *TraceGroupsResult) SetCanLoadMore(v bool)`

SetCanLoadMore sets CanLoadMore field to given value.

### HasCanLoadMore

`func (o *TraceGroupsResult) HasCanLoadMore() bool`

HasCanLoadMore returns a boolean if a field has been set.

### GetItems

`func (o *TraceGroupsResult) GetItems() []TraceGroupsItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TraceGroupsResult) GetItemsOk() (*[]TraceGroupsItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TraceGroupsResult) SetItems(v []TraceGroupsItem)`

SetItems sets Items field to given value.


### GetTotalHits

`func (o *TraceGroupsResult) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *TraceGroupsResult) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *TraceGroupsResult) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *TraceGroupsResult) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetTotalRepresentedItemCount

`func (o *TraceGroupsResult) GetTotalRepresentedItemCount() int64`

GetTotalRepresentedItemCount returns the TotalRepresentedItemCount field if non-nil, zero value otherwise.

### GetTotalRepresentedItemCountOk

`func (o *TraceGroupsResult) GetTotalRepresentedItemCountOk() (*int64, bool)`

GetTotalRepresentedItemCountOk returns a tuple with the TotalRepresentedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRepresentedItemCount

`func (o *TraceGroupsResult) SetTotalRepresentedItemCount(v int64)`

SetTotalRepresentedItemCount sets TotalRepresentedItemCount field to given value.

### HasTotalRepresentedItemCount

`func (o *TraceGroupsResult) HasTotalRepresentedItemCount() bool`

HasTotalRepresentedItemCount returns a boolean if a field has been set.

### GetTotalRetainedItemCount

`func (o *TraceGroupsResult) GetTotalRetainedItemCount() int64`

GetTotalRetainedItemCount returns the TotalRetainedItemCount field if non-nil, zero value otherwise.

### GetTotalRetainedItemCountOk

`func (o *TraceGroupsResult) GetTotalRetainedItemCountOk() (*int64, bool)`

GetTotalRetainedItemCountOk returns a tuple with the TotalRetainedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRetainedItemCount

`func (o *TraceGroupsResult) SetTotalRetainedItemCount(v int64)`

SetTotalRetainedItemCount sets TotalRetainedItemCount field to given value.

### HasTotalRetainedItemCount

`func (o *TraceGroupsResult) HasTotalRetainedItemCount() bool`

HasTotalRetainedItemCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


