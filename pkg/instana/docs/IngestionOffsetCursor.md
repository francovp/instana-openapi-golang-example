# IngestionOffsetCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngestionTime** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 

## Methods

### NewIngestionOffsetCursor

`func NewIngestionOffsetCursor() *IngestionOffsetCursor`

NewIngestionOffsetCursor instantiates a new IngestionOffsetCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionOffsetCursorWithDefaults

`func NewIngestionOffsetCursorWithDefaults() *IngestionOffsetCursor`

NewIngestionOffsetCursorWithDefaults instantiates a new IngestionOffsetCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngestionTime

`func (o *IngestionOffsetCursor) GetIngestionTime() int64`

GetIngestionTime returns the IngestionTime field if non-nil, zero value otherwise.

### GetIngestionTimeOk

`func (o *IngestionOffsetCursor) GetIngestionTimeOk() (*int64, bool)`

GetIngestionTimeOk returns a tuple with the IngestionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionTime

`func (o *IngestionOffsetCursor) SetIngestionTime(v int64)`

SetIngestionTime sets IngestionTime field to given value.

### HasIngestionTime

`func (o *IngestionOffsetCursor) HasIngestionTime() bool`

HasIngestionTime returns a boolean if a field has been set.

### GetOffset

`func (o *IngestionOffsetCursor) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *IngestionOffsetCursor) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *IngestionOffsetCursor) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *IngestionOffsetCursor) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


