# FullTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**RootSpan** | [**Span**](Span.md) |  | 
**TotalErrorCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewFullTrace

`func NewFullTrace(id string, rootSpan Span, ) *FullTrace`

NewFullTrace instantiates a new FullTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullTraceWithDefaults

`func NewFullTraceWithDefaults() *FullTrace`

NewFullTraceWithDefaults instantiates a new FullTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullTrace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullTrace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullTrace) SetId(v string)`

SetId sets Id field to given value.


### GetRootSpan

`func (o *FullTrace) GetRootSpan() Span`

GetRootSpan returns the RootSpan field if non-nil, zero value otherwise.

### GetRootSpanOk

`func (o *FullTrace) GetRootSpanOk() (*Span, bool)`

GetRootSpanOk returns a tuple with the RootSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootSpan

`func (o *FullTrace) SetRootSpan(v Span)`

SetRootSpan sets RootSpan field to given value.


### GetTotalErrorCount

`func (o *FullTrace) GetTotalErrorCount() int32`

GetTotalErrorCount returns the TotalErrorCount field if non-nil, zero value otherwise.

### GetTotalErrorCountOk

`func (o *FullTrace) GetTotalErrorCountOk() (*int32, bool)`

GetTotalErrorCountOk returns a tuple with the TotalErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalErrorCount

`func (o *FullTrace) SetTotalErrorCount(v int32)`

SetTotalErrorCount sets TotalErrorCount field to given value.

### HasTotalErrorCount

`func (o *FullTrace) HasTotalErrorCount() bool`

HasTotalErrorCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


