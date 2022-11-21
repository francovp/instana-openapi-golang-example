# EndpointResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Endpoint**](Endpoint.md) |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**TotalHits** | Pointer to **int32** |  | [optional] 

## Methods

### NewEndpointResult

`func NewEndpointResult() *EndpointResult`

NewEndpointResult instantiates a new EndpointResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointResultWithDefaults

`func NewEndpointResultWithDefaults() *EndpointResult`

NewEndpointResultWithDefaults instantiates a new EndpointResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EndpointResult) GetItems() []Endpoint`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EndpointResult) GetItemsOk() (*[]Endpoint, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EndpointResult) SetItems(v []Endpoint)`

SetItems sets Items field to given value.

### HasItems

`func (o *EndpointResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPage

`func (o *EndpointResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EndpointResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EndpointResult) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *EndpointResult) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *EndpointResult) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EndpointResult) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EndpointResult) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *EndpointResult) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalHits

`func (o *EndpointResult) GetTotalHits() int32`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *EndpointResult) GetTotalHitsOk() (*int32, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *EndpointResult) SetTotalHits(v int32)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *EndpointResult) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


