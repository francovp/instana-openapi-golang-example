# TraceGroupsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | [**IngestionOffsetCursor**](IngestionOffsetCursor.md) |  | 
**Metrics** | [**map[string][][]float32**](array.md) |  | 
**Name** | **string** |  | 
**Timestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewTraceGroupsItem

`func NewTraceGroupsItem(cursor IngestionOffsetCursor, metrics map[string][][]float32, name string, ) *TraceGroupsItem`

NewTraceGroupsItem instantiates a new TraceGroupsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceGroupsItemWithDefaults

`func NewTraceGroupsItemWithDefaults() *TraceGroupsItem`

NewTraceGroupsItemWithDefaults instantiates a new TraceGroupsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *TraceGroupsItem) GetCursor() IngestionOffsetCursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *TraceGroupsItem) GetCursorOk() (*IngestionOffsetCursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *TraceGroupsItem) SetCursor(v IngestionOffsetCursor)`

SetCursor sets Cursor field to given value.


### GetMetrics

`func (o *TraceGroupsItem) GetMetrics() map[string][][]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *TraceGroupsItem) GetMetricsOk() (*map[string][][]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *TraceGroupsItem) SetMetrics(v map[string][][]float32)`

SetMetrics sets Metrics field to given value.


### GetName

`func (o *TraceGroupsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TraceGroupsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TraceGroupsItem) SetName(v string)`

SetName sets Name field to given value.


### GetTimestamp

`func (o *TraceGroupsItem) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TraceGroupsItem) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TraceGroupsItem) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TraceGroupsItem) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


