# ServiceMetricResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustedTimeframe** | Pointer to [**AdjustedTimeframe**](AdjustedTimeframe.md) |  | [optional] 
**Items** | [**[]ServiceItem**](ServiceItem.md) |  | 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalHits** | Pointer to **int64** |  | [optional] 

## Methods

### NewServiceMetricResult

`func NewServiceMetricResult(items []ServiceItem, ) *ServiceMetricResult`

NewServiceMetricResult instantiates a new ServiceMetricResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMetricResultWithDefaults

`func NewServiceMetricResultWithDefaults() *ServiceMetricResult`

NewServiceMetricResultWithDefaults instantiates a new ServiceMetricResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustedTimeframe

`func (o *ServiceMetricResult) GetAdjustedTimeframe() AdjustedTimeframe`

GetAdjustedTimeframe returns the AdjustedTimeframe field if non-nil, zero value otherwise.

### GetAdjustedTimeframeOk

`func (o *ServiceMetricResult) GetAdjustedTimeframeOk() (*AdjustedTimeframe, bool)`

GetAdjustedTimeframeOk returns a tuple with the AdjustedTimeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedTimeframe

`func (o *ServiceMetricResult) SetAdjustedTimeframe(v AdjustedTimeframe)`

SetAdjustedTimeframe sets AdjustedTimeframe field to given value.

### HasAdjustedTimeframe

`func (o *ServiceMetricResult) HasAdjustedTimeframe() bool`

HasAdjustedTimeframe returns a boolean if a field has been set.

### GetItems

`func (o *ServiceMetricResult) GetItems() []ServiceItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ServiceMetricResult) GetItemsOk() (*[]ServiceItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ServiceMetricResult) SetItems(v []ServiceItem)`

SetItems sets Items field to given value.


### GetPage

`func (o *ServiceMetricResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ServiceMetricResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ServiceMetricResult) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ServiceMetricResult) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ServiceMetricResult) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ServiceMetricResult) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ServiceMetricResult) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ServiceMetricResult) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalHits

`func (o *ServiceMetricResult) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *ServiceMetricResult) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *ServiceMetricResult) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *ServiceMetricResult) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


