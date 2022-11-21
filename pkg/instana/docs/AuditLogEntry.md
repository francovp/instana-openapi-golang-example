# AuditLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Actor** | [**LogEntryActor**](LogEntryActor.md) |  | 
**Id** | **string** |  | 
**Message** | **string** |  | 
**Meta** | **map[string]interface{}** |  | 
**Timestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewAuditLogEntry

`func NewAuditLogEntry(action string, actor LogEntryActor, id string, message string, meta map[string]interface{}, ) *AuditLogEntry`

NewAuditLogEntry instantiates a new AuditLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogEntryWithDefaults

`func NewAuditLogEntryWithDefaults() *AuditLogEntry`

NewAuditLogEntryWithDefaults instantiates a new AuditLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AuditLogEntry) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditLogEntry) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditLogEntry) SetAction(v string)`

SetAction sets Action field to given value.


### GetActor

`func (o *AuditLogEntry) GetActor() LogEntryActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AuditLogEntry) GetActorOk() (*LogEntryActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AuditLogEntry) SetActor(v LogEntryActor)`

SetActor sets Actor field to given value.


### GetId

`func (o *AuditLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogEntry) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *AuditLogEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuditLogEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuditLogEntry) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMeta

`func (o *AuditLogEntry) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuditLogEntry) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuditLogEntry) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### GetTimestamp

`func (o *AuditLogEntry) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuditLogEntry) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuditLogEntry) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AuditLogEntry) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


