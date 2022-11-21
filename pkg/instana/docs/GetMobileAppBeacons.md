# GetMobileAppBeacons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**CursorPagination**](CursorPagination.md) |  | [optional] 
**TagFilters** | Pointer to [**[]DeprecatedTagFilter**](DeprecatedTagFilter.md) |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewGetMobileAppBeacons

`func NewGetMobileAppBeacons(type_ string, ) *GetMobileAppBeacons`

NewGetMobileAppBeacons instantiates a new GetMobileAppBeacons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMobileAppBeaconsWithDefaults

`func NewGetMobileAppBeaconsWithDefaults() *GetMobileAppBeacons`

NewGetMobileAppBeaconsWithDefaults instantiates a new GetMobileAppBeacons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *GetMobileAppBeacons) GetPagination() CursorPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetMobileAppBeacons) GetPaginationOk() (*CursorPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetMobileAppBeacons) SetPagination(v CursorPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetMobileAppBeacons) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetTagFilters

`func (o *GetMobileAppBeacons) GetTagFilters() []DeprecatedTagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *GetMobileAppBeacons) GetTagFiltersOk() (*[]DeprecatedTagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *GetMobileAppBeacons) SetTagFilters(v []DeprecatedTagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *GetMobileAppBeacons) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetMobileAppBeacons) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetMobileAppBeacons) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetMobileAppBeacons) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetMobileAppBeacons) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetType

`func (o *GetMobileAppBeacons) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMobileAppBeacons) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMobileAppBeacons) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


