# WebsiteBeaconResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustedTimeframe** | Pointer to [**AdjustedTimeframe**](AdjustedTimeframe.md) |  | [optional] 
**CanLoadMore** | Pointer to **bool** |  | [optional] 
**Items** | [**[]WebsiteBeaconsItem**](WebsiteBeaconsItem.md) |  | 
**TotalHits** | Pointer to **int64** |  | [optional] 
**TotalRepresentedItemCount** | Pointer to **int64** |  | [optional] 
**TotalRetainedItemCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewWebsiteBeaconResult

`func NewWebsiteBeaconResult(items []WebsiteBeaconsItem, ) *WebsiteBeaconResult`

NewWebsiteBeaconResult instantiates a new WebsiteBeaconResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteBeaconResultWithDefaults

`func NewWebsiteBeaconResultWithDefaults() *WebsiteBeaconResult`

NewWebsiteBeaconResultWithDefaults instantiates a new WebsiteBeaconResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustedTimeframe

`func (o *WebsiteBeaconResult) GetAdjustedTimeframe() AdjustedTimeframe`

GetAdjustedTimeframe returns the AdjustedTimeframe field if non-nil, zero value otherwise.

### GetAdjustedTimeframeOk

`func (o *WebsiteBeaconResult) GetAdjustedTimeframeOk() (*AdjustedTimeframe, bool)`

GetAdjustedTimeframeOk returns a tuple with the AdjustedTimeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedTimeframe

`func (o *WebsiteBeaconResult) SetAdjustedTimeframe(v AdjustedTimeframe)`

SetAdjustedTimeframe sets AdjustedTimeframe field to given value.

### HasAdjustedTimeframe

`func (o *WebsiteBeaconResult) HasAdjustedTimeframe() bool`

HasAdjustedTimeframe returns a boolean if a field has been set.

### GetCanLoadMore

`func (o *WebsiteBeaconResult) GetCanLoadMore() bool`

GetCanLoadMore returns the CanLoadMore field if non-nil, zero value otherwise.

### GetCanLoadMoreOk

`func (o *WebsiteBeaconResult) GetCanLoadMoreOk() (*bool, bool)`

GetCanLoadMoreOk returns a tuple with the CanLoadMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLoadMore

`func (o *WebsiteBeaconResult) SetCanLoadMore(v bool)`

SetCanLoadMore sets CanLoadMore field to given value.

### HasCanLoadMore

`func (o *WebsiteBeaconResult) HasCanLoadMore() bool`

HasCanLoadMore returns a boolean if a field has been set.

### GetItems

`func (o *WebsiteBeaconResult) GetItems() []WebsiteBeaconsItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WebsiteBeaconResult) GetItemsOk() (*[]WebsiteBeaconsItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WebsiteBeaconResult) SetItems(v []WebsiteBeaconsItem)`

SetItems sets Items field to given value.


### GetTotalHits

`func (o *WebsiteBeaconResult) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *WebsiteBeaconResult) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *WebsiteBeaconResult) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *WebsiteBeaconResult) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetTotalRepresentedItemCount

`func (o *WebsiteBeaconResult) GetTotalRepresentedItemCount() int64`

GetTotalRepresentedItemCount returns the TotalRepresentedItemCount field if non-nil, zero value otherwise.

### GetTotalRepresentedItemCountOk

`func (o *WebsiteBeaconResult) GetTotalRepresentedItemCountOk() (*int64, bool)`

GetTotalRepresentedItemCountOk returns a tuple with the TotalRepresentedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRepresentedItemCount

`func (o *WebsiteBeaconResult) SetTotalRepresentedItemCount(v int64)`

SetTotalRepresentedItemCount sets TotalRepresentedItemCount field to given value.

### HasTotalRepresentedItemCount

`func (o *WebsiteBeaconResult) HasTotalRepresentedItemCount() bool`

HasTotalRepresentedItemCount returns a boolean if a field has been set.

### GetTotalRetainedItemCount

`func (o *WebsiteBeaconResult) GetTotalRetainedItemCount() int64`

GetTotalRetainedItemCount returns the TotalRetainedItemCount field if non-nil, zero value otherwise.

### GetTotalRetainedItemCountOk

`func (o *WebsiteBeaconResult) GetTotalRetainedItemCountOk() (*int64, bool)`

GetTotalRetainedItemCountOk returns a tuple with the TotalRetainedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRetainedItemCount

`func (o *WebsiteBeaconResult) SetTotalRetainedItemCount(v int64)`

SetTotalRetainedItemCount sets TotalRetainedItemCount field to given value.

### HasTotalRetainedItemCount

`func (o *WebsiteBeaconResult) HasTotalRetainedItemCount() bool`

HasTotalRetainedItemCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


