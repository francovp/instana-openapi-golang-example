# TestResultListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]TestResultListItem**](TestResultListItem.md) |  | 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalHits** | Pointer to **int64** |  | [optional] 

## Methods

### NewTestResultListResult

`func NewTestResultListResult(items []TestResultListItem, ) *TestResultListResult`

NewTestResultListResult instantiates a new TestResultListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultListResultWithDefaults

`func NewTestResultListResultWithDefaults() *TestResultListResult`

NewTestResultListResultWithDefaults instantiates a new TestResultListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TestResultListResult) GetItems() []TestResultListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TestResultListResult) GetItemsOk() (*[]TestResultListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TestResultListResult) SetItems(v []TestResultListItem)`

SetItems sets Items field to given value.


### GetPage

`func (o *TestResultListResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TestResultListResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TestResultListResult) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *TestResultListResult) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *TestResultListResult) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *TestResultListResult) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *TestResultListResult) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *TestResultListResult) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalHits

`func (o *TestResultListResult) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *TestResultListResult) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *TestResultListResult) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *TestResultListResult) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

