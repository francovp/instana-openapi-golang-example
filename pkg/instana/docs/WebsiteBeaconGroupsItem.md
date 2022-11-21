# WebsiteBeaconGroupsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | [**IngestionOffsetCursor**](IngestionOffsetCursor.md) |  | 
**EarliestTimestamp** | Pointer to **int64** |  | [optional] 
**Metrics** | [**map[string][][]float32**](array.md) |  | 
**Name** | **string** |  | 

## Methods

### NewWebsiteBeaconGroupsItem

`func NewWebsiteBeaconGroupsItem(cursor IngestionOffsetCursor, metrics map[string][][]float32, name string, ) *WebsiteBeaconGroupsItem`

NewWebsiteBeaconGroupsItem instantiates a new WebsiteBeaconGroupsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteBeaconGroupsItemWithDefaults

`func NewWebsiteBeaconGroupsItemWithDefaults() *WebsiteBeaconGroupsItem`

NewWebsiteBeaconGroupsItemWithDefaults instantiates a new WebsiteBeaconGroupsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *WebsiteBeaconGroupsItem) GetCursor() IngestionOffsetCursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *WebsiteBeaconGroupsItem) GetCursorOk() (*IngestionOffsetCursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *WebsiteBeaconGroupsItem) SetCursor(v IngestionOffsetCursor)`

SetCursor sets Cursor field to given value.


### GetEarliestTimestamp

`func (o *WebsiteBeaconGroupsItem) GetEarliestTimestamp() int64`

GetEarliestTimestamp returns the EarliestTimestamp field if non-nil, zero value otherwise.

### GetEarliestTimestampOk

`func (o *WebsiteBeaconGroupsItem) GetEarliestTimestampOk() (*int64, bool)`

GetEarliestTimestampOk returns a tuple with the EarliestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestTimestamp

`func (o *WebsiteBeaconGroupsItem) SetEarliestTimestamp(v int64)`

SetEarliestTimestamp sets EarliestTimestamp field to given value.

### HasEarliestTimestamp

`func (o *WebsiteBeaconGroupsItem) HasEarliestTimestamp() bool`

HasEarliestTimestamp returns a boolean if a field has been set.

### GetMetrics

`func (o *WebsiteBeaconGroupsItem) GetMetrics() map[string][][]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *WebsiteBeaconGroupsItem) GetMetricsOk() (*map[string][][]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *WebsiteBeaconGroupsItem) SetMetrics(v map[string][][]float32)`

SetMetrics sets Metrics field to given value.


### GetName

`func (o *WebsiteBeaconGroupsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebsiteBeaconGroupsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebsiteBeaconGroupsItem) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


