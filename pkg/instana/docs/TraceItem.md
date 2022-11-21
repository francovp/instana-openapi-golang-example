# TraceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | [**IngestionOffsetCursor**](IngestionOffsetCursor.md) |  | 
**Trace** | [**Trace**](Trace.md) |  | 

## Methods

### NewTraceItem

`func NewTraceItem(cursor IngestionOffsetCursor, trace Trace, ) *TraceItem`

NewTraceItem instantiates a new TraceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceItemWithDefaults

`func NewTraceItemWithDefaults() *TraceItem`

NewTraceItemWithDefaults instantiates a new TraceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *TraceItem) GetCursor() IngestionOffsetCursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *TraceItem) GetCursorOk() (*IngestionOffsetCursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *TraceItem) SetCursor(v IngestionOffsetCursor)`

SetCursor sets Cursor field to given value.


### GetTrace

`func (o *TraceItem) GetTrace() Trace`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *TraceItem) GetTraceOk() (*Trace, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *TraceItem) SetTrace(v Trace)`

SetTrace sets Trace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


