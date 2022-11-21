# AuditLogUiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]AuditLogEntry**](AuditLogEntry.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewAuditLogUiResponse

`func NewAuditLogUiResponse() *AuditLogUiResponse`

NewAuditLogUiResponse instantiates a new AuditLogUiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogUiResponseWithDefaults

`func NewAuditLogUiResponseWithDefaults() *AuditLogUiResponse`

NewAuditLogUiResponseWithDefaults instantiates a new AuditLogUiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *AuditLogUiResponse) GetEntries() []AuditLogEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *AuditLogUiResponse) GetEntriesOk() (*[]AuditLogEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *AuditLogUiResponse) SetEntries(v []AuditLogEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *AuditLogUiResponse) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetTotal

`func (o *AuditLogUiResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AuditLogUiResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AuditLogUiResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AuditLogUiResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


