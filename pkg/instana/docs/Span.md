# Span

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchSelfTime** | Pointer to **int64** |  | [optional] 
**BatchSize** | Pointer to **int32** |  | [optional] 
**CalculatedSelfTime** | Pointer to **int64** |  | [optional] 
**CallId** | **string** |  | 
**ChildSpans** | [**[]Span**](Span.md) |  | 
**Data** | **map[string]interface{}** |  | 
**Destination** | Pointer to [**SpanRelation**](SpanRelation.md) |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**ErrorCount** | Pointer to **int32** |  | [optional] 
**ForeignParentId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**IsSynthetic** | Pointer to **bool** |  | [optional] 
**Kind** | **string** |  | 
**Label** | **string** |  | 
**Name** | **string** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**SpanRelation**](SpanRelation.md) |  | [optional] 
**StackTrace** | [**[]StackTraceItem**](StackTraceItem.md) |  | 
**Start** | Pointer to **int64** |  | [optional] 

## Methods

### NewSpan

`func NewSpan(callId string, childSpans []Span, data map[string]interface{}, id string, kind string, label string, name string, stackTrace []StackTraceItem, ) *Span`

NewSpan instantiates a new Span object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanWithDefaults

`func NewSpanWithDefaults() *Span`

NewSpanWithDefaults instantiates a new Span object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchSelfTime

`func (o *Span) GetBatchSelfTime() int64`

GetBatchSelfTime returns the BatchSelfTime field if non-nil, zero value otherwise.

### GetBatchSelfTimeOk

`func (o *Span) GetBatchSelfTimeOk() (*int64, bool)`

GetBatchSelfTimeOk returns a tuple with the BatchSelfTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSelfTime

`func (o *Span) SetBatchSelfTime(v int64)`

SetBatchSelfTime sets BatchSelfTime field to given value.

### HasBatchSelfTime

`func (o *Span) HasBatchSelfTime() bool`

HasBatchSelfTime returns a boolean if a field has been set.

### GetBatchSize

`func (o *Span) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *Span) GetBatchSizeOk() (*int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *Span) SetBatchSize(v int32)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *Span) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetCalculatedSelfTime

`func (o *Span) GetCalculatedSelfTime() int64`

GetCalculatedSelfTime returns the CalculatedSelfTime field if non-nil, zero value otherwise.

### GetCalculatedSelfTimeOk

`func (o *Span) GetCalculatedSelfTimeOk() (*int64, bool)`

GetCalculatedSelfTimeOk returns a tuple with the CalculatedSelfTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedSelfTime

`func (o *Span) SetCalculatedSelfTime(v int64)`

SetCalculatedSelfTime sets CalculatedSelfTime field to given value.

### HasCalculatedSelfTime

`func (o *Span) HasCalculatedSelfTime() bool`

HasCalculatedSelfTime returns a boolean if a field has been set.

### GetCallId

`func (o *Span) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *Span) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *Span) SetCallId(v string)`

SetCallId sets CallId field to given value.


### GetChildSpans

`func (o *Span) GetChildSpans() []Span`

GetChildSpans returns the ChildSpans field if non-nil, zero value otherwise.

### GetChildSpansOk

`func (o *Span) GetChildSpansOk() (*[]Span, bool)`

GetChildSpansOk returns a tuple with the ChildSpans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSpans

`func (o *Span) SetChildSpans(v []Span)`

SetChildSpans sets ChildSpans field to given value.


### GetData

`func (o *Span) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Span) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Span) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetDestination

`func (o *Span) GetDestination() SpanRelation`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Span) GetDestinationOk() (*SpanRelation, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Span) SetDestination(v SpanRelation)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Span) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDuration

`func (o *Span) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Span) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Span) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Span) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetErrorCount

`func (o *Span) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *Span) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *Span) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *Span) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetForeignParentId

`func (o *Span) GetForeignParentId() string`

GetForeignParentId returns the ForeignParentId field if non-nil, zero value otherwise.

### GetForeignParentIdOk

`func (o *Span) GetForeignParentIdOk() (*string, bool)`

GetForeignParentIdOk returns a tuple with the ForeignParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignParentId

`func (o *Span) SetForeignParentId(v string)`

SetForeignParentId sets ForeignParentId field to given value.

### HasForeignParentId

`func (o *Span) HasForeignParentId() bool`

HasForeignParentId returns a boolean if a field has been set.

### GetId

`func (o *Span) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Span) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Span) SetId(v string)`

SetId sets Id field to given value.


### GetIsSynthetic

`func (o *Span) GetIsSynthetic() bool`

GetIsSynthetic returns the IsSynthetic field if non-nil, zero value otherwise.

### GetIsSyntheticOk

`func (o *Span) GetIsSyntheticOk() (*bool, bool)`

GetIsSyntheticOk returns a tuple with the IsSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynthetic

`func (o *Span) SetIsSynthetic(v bool)`

SetIsSynthetic sets IsSynthetic field to given value.

### HasIsSynthetic

`func (o *Span) HasIsSynthetic() bool`

HasIsSynthetic returns a boolean if a field has been set.

### GetKind

`func (o *Span) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Span) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Span) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLabel

`func (o *Span) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Span) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Span) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *Span) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Span) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Span) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *Span) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Span) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Span) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Span) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetSource

`func (o *Span) GetSource() SpanRelation`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Span) GetSourceOk() (*SpanRelation, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Span) SetSource(v SpanRelation)`

SetSource sets Source field to given value.

### HasSource

`func (o *Span) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStackTrace

`func (o *Span) GetStackTrace() []StackTraceItem`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *Span) GetStackTraceOk() (*[]StackTraceItem, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *Span) SetStackTrace(v []StackTraceItem)`

SetStackTrace sets StackTrace field to given value.


### GetStart

`func (o *Span) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Span) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Span) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *Span) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


