# EventResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** |  | [optional] 
**End** | Pointer to **int64** |  | [optional] 
**EntityLabel** | Pointer to **string** |  | [optional] 
**EntityName** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 
**EventId** | Pointer to **string** |  | [optional] 
**FixSuggestion** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Problem** | Pointer to **string** |  | [optional] 
**RecentEvents** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**SnapshotId** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**TriggeringTime** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewEventResult

`func NewEventResult() *EventResult`

NewEventResult instantiates a new EventResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResultWithDefaults

`func NewEventResultWithDefaults() *EventResult`

NewEventResultWithDefaults instantiates a new EventResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *EventResult) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *EventResult) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *EventResult) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *EventResult) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetEnd

`func (o *EventResult) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *EventResult) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *EventResult) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *EventResult) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetEntityLabel

`func (o *EventResult) GetEntityLabel() string`

GetEntityLabel returns the EntityLabel field if non-nil, zero value otherwise.

### GetEntityLabelOk

`func (o *EventResult) GetEntityLabelOk() (*string, bool)`

GetEntityLabelOk returns a tuple with the EntityLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityLabel

`func (o *EventResult) SetEntityLabel(v string)`

SetEntityLabel sets EntityLabel field to given value.

### HasEntityLabel

`func (o *EventResult) HasEntityLabel() bool`

HasEntityLabel returns a boolean if a field has been set.

### GetEntityName

`func (o *EventResult) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *EventResult) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *EventResult) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *EventResult) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetEntityType

`func (o *EventResult) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EventResult) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EventResult) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *EventResult) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEventId

`func (o *EventResult) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EventResult) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EventResult) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *EventResult) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetFixSuggestion

`func (o *EventResult) GetFixSuggestion() string`

GetFixSuggestion returns the FixSuggestion field if non-nil, zero value otherwise.

### GetFixSuggestionOk

`func (o *EventResult) GetFixSuggestionOk() (*string, bool)`

GetFixSuggestionOk returns a tuple with the FixSuggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixSuggestion

`func (o *EventResult) SetFixSuggestion(v string)`

SetFixSuggestion sets FixSuggestion field to given value.

### HasFixSuggestion

`func (o *EventResult) HasFixSuggestion() bool`

HasFixSuggestion returns a boolean if a field has been set.

### GetMetrics

`func (o *EventResult) GetMetrics() []map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *EventResult) GetMetricsOk() (*[]map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *EventResult) SetMetrics(v []map[string]interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *EventResult) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetProblem

`func (o *EventResult) GetProblem() string`

GetProblem returns the Problem field if non-nil, zero value otherwise.

### GetProblemOk

`func (o *EventResult) GetProblemOk() (*string, bool)`

GetProblemOk returns a tuple with the Problem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblem

`func (o *EventResult) SetProblem(v string)`

SetProblem sets Problem field to given value.

### HasProblem

`func (o *EventResult) HasProblem() bool`

HasProblem returns a boolean if a field has been set.

### GetRecentEvents

`func (o *EventResult) GetRecentEvents() []map[string]interface{}`

GetRecentEvents returns the RecentEvents field if non-nil, zero value otherwise.

### GetRecentEventsOk

`func (o *EventResult) GetRecentEventsOk() (*[]map[string]interface{}, bool)`

GetRecentEventsOk returns a tuple with the RecentEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentEvents

`func (o *EventResult) SetRecentEvents(v []map[string]interface{})`

SetRecentEvents sets RecentEvents field to given value.

### HasRecentEvents

`func (o *EventResult) HasRecentEvents() bool`

HasRecentEvents returns a boolean if a field has been set.

### GetSeverity

`func (o *EventResult) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EventResult) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EventResult) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *EventResult) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSnapshotId

`func (o *EventResult) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *EventResult) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *EventResult) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *EventResult) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetStart

`func (o *EventResult) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EventResult) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EventResult) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *EventResult) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetState

`func (o *EventResult) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EventResult) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EventResult) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EventResult) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTriggeringTime

`func (o *EventResult) GetTriggeringTime() int64`

GetTriggeringTime returns the TriggeringTime field if non-nil, zero value otherwise.

### GetTriggeringTimeOk

`func (o *EventResult) GetTriggeringTimeOk() (*int64, bool)`

GetTriggeringTimeOk returns a tuple with the TriggeringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeringTime

`func (o *EventResult) SetTriggeringTime(v int64)`

SetTriggeringTime sets TriggeringTime field to given value.

### HasTriggeringTime

`func (o *EventResult) HasTriggeringTime() bool`

HasTriggeringTime returns a boolean if a field has been set.

### GetType

`func (o *EventResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventResult) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


