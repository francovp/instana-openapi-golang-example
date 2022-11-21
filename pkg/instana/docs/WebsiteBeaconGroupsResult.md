# WebsiteBeaconGroupsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustedTimeframe** | Pointer to [**AdjustedTimeframe**](AdjustedTimeframe.md) |  | [optional] 
**CanLoadMore** | Pointer to **bool** |  | [optional] 
**Items** | [**[]WebsiteBeaconGroupsItem**](WebsiteBeaconGroupsItem.md) |  | 
**TotalHits** | Pointer to **int64** |  | [optional] 
**TotalRepresentedItemCount** | Pointer to **int64** |  | [optional] 
**TotalRetainedItemCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewWebsiteBeaconGroupsResult

`func NewWebsiteBeaconGroupsResult(items []WebsiteBeaconGroupsItem, ) *WebsiteBeaconGroupsResult`

NewWebsiteBeaconGroupsResult instantiates a new WebsiteBeaconGroupsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteBeaconGroupsResultWithDefaults

`func NewWebsiteBeaconGroupsResultWithDefaults() *WebsiteBeaconGroupsResult`

NewWebsiteBeaconGroupsResultWithDefaults instantiates a new WebsiteBeaconGroupsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustedTimeframe

`func (o *WebsiteBeaconGroupsResult) GetAdjustedTimeframe() AdjustedTimeframe`

GetAdjustedTimeframe returns the AdjustedTimeframe field if non-nil, zero value otherwise.

### GetAdjustedTimeframeOk

`func (o *WebsiteBeaconGroupsResult) GetAdjustedTimeframeOk() (*AdjustedTimeframe, bool)`

GetAdjustedTimeframeOk returns a tuple with the AdjustedTimeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedTimeframe

`func (o *WebsiteBeaconGroupsResult) SetAdjustedTimeframe(v AdjustedTimeframe)`

SetAdjustedTimeframe sets AdjustedTimeframe field to given value.

### HasAdjustedTimeframe

`func (o *WebsiteBeaconGroupsResult) HasAdjustedTimeframe() bool`

HasAdjustedTimeframe returns a boolean if a field has been set.

### GetCanLoadMore

`func (o *WebsiteBeaconGroupsResult) GetCanLoadMore() bool`

GetCanLoadMore returns the CanLoadMore field if non-nil, zero value otherwise.

### GetCanLoadMoreOk

`func (o *WebsiteBeaconGroupsResult) GetCanLoadMoreOk() (*bool, bool)`

GetCanLoadMoreOk returns a tuple with the CanLoadMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLoadMore

`func (o *WebsiteBeaconGroupsResult) SetCanLoadMore(v bool)`

SetCanLoadMore sets CanLoadMore field to given value.

### HasCanLoadMore

`func (o *WebsiteBeaconGroupsResult) HasCanLoadMore() bool`

HasCanLoadMore returns a boolean if a field has been set.

### GetItems

`func (o *WebsiteBeaconGroupsResult) GetItems() []WebsiteBeaconGroupsItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WebsiteBeaconGroupsResult) GetItemsOk() (*[]WebsiteBeaconGroupsItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WebsiteBeaconGroupsResult) SetItems(v []WebsiteBeaconGroupsItem)`

SetItems sets Items field to given value.


### GetTotalHits

`func (o *WebsiteBeaconGroupsResult) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *WebsiteBeaconGroupsResult) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *WebsiteBeaconGroupsResult) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *WebsiteBeaconGroupsResult) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetTotalRepresentedItemCount

`func (o *WebsiteBeaconGroupsResult) GetTotalRepresentedItemCount() int64`

GetTotalRepresentedItemCount returns the TotalRepresentedItemCount field if non-nil, zero value otherwise.

### GetTotalRepresentedItemCountOk

`func (o *WebsiteBeaconGroupsResult) GetTotalRepresentedItemCountOk() (*int64, bool)`

GetTotalRepresentedItemCountOk returns a tuple with the TotalRepresentedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRepresentedItemCount

`func (o *WebsiteBeaconGroupsResult) SetTotalRepresentedItemCount(v int64)`

SetTotalRepresentedItemCount sets TotalRepresentedItemCount field to given value.

### HasTotalRepresentedItemCount

`func (o *WebsiteBeaconGroupsResult) HasTotalRepresentedItemCount() bool`

HasTotalRepresentedItemCount returns a boolean if a field has been set.

### GetTotalRetainedItemCount

`func (o *WebsiteBeaconGroupsResult) GetTotalRetainedItemCount() int64`

GetTotalRetainedItemCount returns the TotalRetainedItemCount field if non-nil, zero value otherwise.

### GetTotalRetainedItemCountOk

`func (o *WebsiteBeaconGroupsResult) GetTotalRetainedItemCountOk() (*int64, bool)`

GetTotalRetainedItemCountOk returns a tuple with the TotalRetainedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRetainedItemCount

`func (o *WebsiteBeaconGroupsResult) SetTotalRetainedItemCount(v int64)`

SetTotalRetainedItemCount sets TotalRetainedItemCount field to given value.

### HasTotalRetainedItemCount

`func (o *WebsiteBeaconGroupsResult) HasTotalRetainedItemCount() bool`

HasTotalRetainedItemCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


