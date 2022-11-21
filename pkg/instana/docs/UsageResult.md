# UsageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]UsageResultItems**](UsageResultItems.md) |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewUsageResult

`func NewUsageResult() *UsageResult`

NewUsageResult instantiates a new UsageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageResultWithDefaults

`func NewUsageResultWithDefaults() *UsageResult`

NewUsageResultWithDefaults instantiates a new UsageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *UsageResult) GetItems() []UsageResultItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UsageResult) GetItemsOk() (*[]UsageResultItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UsageResult) SetItems(v []UsageResultItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *UsageResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTime

`func (o *UsageResult) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UsageResult) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UsageResult) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *UsageResult) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


