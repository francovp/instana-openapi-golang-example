# MobileAppBeaconResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustedTimeframe** | Pointer to [**AdjustedTimeframe**](AdjustedTimeframe.md) |  | [optional] 
**CanLoadMore** | Pointer to **bool** |  | [optional] 
**Items** | [**[]MobileAppBeaconsItem**](MobileAppBeaconsItem.md) |  | 
**TotalHits** | Pointer to **int64** |  | [optional] 
**TotalRepresentedItemCount** | Pointer to **int64** |  | [optional] 
**TotalRetainedItemCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewMobileAppBeaconResult

`func NewMobileAppBeaconResult(items []MobileAppBeaconsItem, ) *MobileAppBeaconResult`

NewMobileAppBeaconResult instantiates a new MobileAppBeaconResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileAppBeaconResultWithDefaults

`func NewMobileAppBeaconResultWithDefaults() *MobileAppBeaconResult`

NewMobileAppBeaconResultWithDefaults instantiates a new MobileAppBeaconResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustedTimeframe

`func (o *MobileAppBeaconResult) GetAdjustedTimeframe() AdjustedTimeframe`

GetAdjustedTimeframe returns the AdjustedTimeframe field if non-nil, zero value otherwise.

### GetAdjustedTimeframeOk

`func (o *MobileAppBeaconResult) GetAdjustedTimeframeOk() (*AdjustedTimeframe, bool)`

GetAdjustedTimeframeOk returns a tuple with the AdjustedTimeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedTimeframe

`func (o *MobileAppBeaconResult) SetAdjustedTimeframe(v AdjustedTimeframe)`

SetAdjustedTimeframe sets AdjustedTimeframe field to given value.

### HasAdjustedTimeframe

`func (o *MobileAppBeaconResult) HasAdjustedTimeframe() bool`

HasAdjustedTimeframe returns a boolean if a field has been set.

### GetCanLoadMore

`func (o *MobileAppBeaconResult) GetCanLoadMore() bool`

GetCanLoadMore returns the CanLoadMore field if non-nil, zero value otherwise.

### GetCanLoadMoreOk

`func (o *MobileAppBeaconResult) GetCanLoadMoreOk() (*bool, bool)`

GetCanLoadMoreOk returns a tuple with the CanLoadMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLoadMore

`func (o *MobileAppBeaconResult) SetCanLoadMore(v bool)`

SetCanLoadMore sets CanLoadMore field to given value.

### HasCanLoadMore

`func (o *MobileAppBeaconResult) HasCanLoadMore() bool`

HasCanLoadMore returns a boolean if a field has been set.

### GetItems

`func (o *MobileAppBeaconResult) GetItems() []MobileAppBeaconsItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MobileAppBeaconResult) GetItemsOk() (*[]MobileAppBeaconsItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MobileAppBeaconResult) SetItems(v []MobileAppBeaconsItem)`

SetItems sets Items field to given value.


### GetTotalHits

`func (o *MobileAppBeaconResult) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *MobileAppBeaconResult) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *MobileAppBeaconResult) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *MobileAppBeaconResult) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetTotalRepresentedItemCount

`func (o *MobileAppBeaconResult) GetTotalRepresentedItemCount() int64`

GetTotalRepresentedItemCount returns the TotalRepresentedItemCount field if non-nil, zero value otherwise.

### GetTotalRepresentedItemCountOk

`func (o *MobileAppBeaconResult) GetTotalRepresentedItemCountOk() (*int64, bool)`

GetTotalRepresentedItemCountOk returns a tuple with the TotalRepresentedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRepresentedItemCount

`func (o *MobileAppBeaconResult) SetTotalRepresentedItemCount(v int64)`

SetTotalRepresentedItemCount sets TotalRepresentedItemCount field to given value.

### HasTotalRepresentedItemCount

`func (o *MobileAppBeaconResult) HasTotalRepresentedItemCount() bool`

HasTotalRepresentedItemCount returns a boolean if a field has been set.

### GetTotalRetainedItemCount

`func (o *MobileAppBeaconResult) GetTotalRetainedItemCount() int64`

GetTotalRetainedItemCount returns the TotalRetainedItemCount field if non-nil, zero value otherwise.

### GetTotalRetainedItemCountOk

`func (o *MobileAppBeaconResult) GetTotalRetainedItemCountOk() (*int64, bool)`

GetTotalRetainedItemCountOk returns a tuple with the TotalRetainedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRetainedItemCount

`func (o *MobileAppBeaconResult) SetTotalRetainedItemCount(v int64)`

SetTotalRetainedItemCount sets TotalRetainedItemCount field to given value.

### HasTotalRetainedItemCount

`func (o *MobileAppBeaconResult) HasTotalRetainedItemCount() bool`

HasTotalRetainedItemCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


