# MobileAppBeaconGroupsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustedTimeframe** | Pointer to [**AdjustedTimeframe**](AdjustedTimeframe.md) |  | [optional] 
**CanLoadMore** | Pointer to **bool** |  | [optional] 
**Items** | [**[]MobileAppBeaconGroupsItem**](MobileAppBeaconGroupsItem.md) |  | 
**TotalHits** | Pointer to **int64** |  | [optional] 
**TotalRepresentedItemCount** | Pointer to **int64** |  | [optional] 
**TotalRetainedItemCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewMobileAppBeaconGroupsResult

`func NewMobileAppBeaconGroupsResult(items []MobileAppBeaconGroupsItem, ) *MobileAppBeaconGroupsResult`

NewMobileAppBeaconGroupsResult instantiates a new MobileAppBeaconGroupsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileAppBeaconGroupsResultWithDefaults

`func NewMobileAppBeaconGroupsResultWithDefaults() *MobileAppBeaconGroupsResult`

NewMobileAppBeaconGroupsResultWithDefaults instantiates a new MobileAppBeaconGroupsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustedTimeframe

`func (o *MobileAppBeaconGroupsResult) GetAdjustedTimeframe() AdjustedTimeframe`

GetAdjustedTimeframe returns the AdjustedTimeframe field if non-nil, zero value otherwise.

### GetAdjustedTimeframeOk

`func (o *MobileAppBeaconGroupsResult) GetAdjustedTimeframeOk() (*AdjustedTimeframe, bool)`

GetAdjustedTimeframeOk returns a tuple with the AdjustedTimeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedTimeframe

`func (o *MobileAppBeaconGroupsResult) SetAdjustedTimeframe(v AdjustedTimeframe)`

SetAdjustedTimeframe sets AdjustedTimeframe field to given value.

### HasAdjustedTimeframe

`func (o *MobileAppBeaconGroupsResult) HasAdjustedTimeframe() bool`

HasAdjustedTimeframe returns a boolean if a field has been set.

### GetCanLoadMore

`func (o *MobileAppBeaconGroupsResult) GetCanLoadMore() bool`

GetCanLoadMore returns the CanLoadMore field if non-nil, zero value otherwise.

### GetCanLoadMoreOk

`func (o *MobileAppBeaconGroupsResult) GetCanLoadMoreOk() (*bool, bool)`

GetCanLoadMoreOk returns a tuple with the CanLoadMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLoadMore

`func (o *MobileAppBeaconGroupsResult) SetCanLoadMore(v bool)`

SetCanLoadMore sets CanLoadMore field to given value.

### HasCanLoadMore

`func (o *MobileAppBeaconGroupsResult) HasCanLoadMore() bool`

HasCanLoadMore returns a boolean if a field has been set.

### GetItems

`func (o *MobileAppBeaconGroupsResult) GetItems() []MobileAppBeaconGroupsItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MobileAppBeaconGroupsResult) GetItemsOk() (*[]MobileAppBeaconGroupsItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MobileAppBeaconGroupsResult) SetItems(v []MobileAppBeaconGroupsItem)`

SetItems sets Items field to given value.


### GetTotalHits

`func (o *MobileAppBeaconGroupsResult) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *MobileAppBeaconGroupsResult) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *MobileAppBeaconGroupsResult) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *MobileAppBeaconGroupsResult) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetTotalRepresentedItemCount

`func (o *MobileAppBeaconGroupsResult) GetTotalRepresentedItemCount() int64`

GetTotalRepresentedItemCount returns the TotalRepresentedItemCount field if non-nil, zero value otherwise.

### GetTotalRepresentedItemCountOk

`func (o *MobileAppBeaconGroupsResult) GetTotalRepresentedItemCountOk() (*int64, bool)`

GetTotalRepresentedItemCountOk returns a tuple with the TotalRepresentedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRepresentedItemCount

`func (o *MobileAppBeaconGroupsResult) SetTotalRepresentedItemCount(v int64)`

SetTotalRepresentedItemCount sets TotalRepresentedItemCount field to given value.

### HasTotalRepresentedItemCount

`func (o *MobileAppBeaconGroupsResult) HasTotalRepresentedItemCount() bool`

HasTotalRepresentedItemCount returns a boolean if a field has been set.

### GetTotalRetainedItemCount

`func (o *MobileAppBeaconGroupsResult) GetTotalRetainedItemCount() int64`

GetTotalRetainedItemCount returns the TotalRetainedItemCount field if non-nil, zero value otherwise.

### GetTotalRetainedItemCountOk

`func (o *MobileAppBeaconGroupsResult) GetTotalRetainedItemCountOk() (*int64, bool)`

GetTotalRetainedItemCountOk returns a tuple with the TotalRetainedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRetainedItemCount

`func (o *MobileAppBeaconGroupsResult) SetTotalRetainedItemCount(v int64)`

SetTotalRetainedItemCount sets TotalRetainedItemCount field to given value.

### HasTotalRetainedItemCount

`func (o *MobileAppBeaconGroupsResult) HasTotalRetainedItemCount() bool`

HasTotalRetainedItemCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


