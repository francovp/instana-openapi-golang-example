# AccessLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]AccessLogEntry**](AccessLogEntry.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccessLogResponse

`func NewAccessLogResponse() *AccessLogResponse`

NewAccessLogResponse instantiates a new AccessLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessLogResponseWithDefaults

`func NewAccessLogResponseWithDefaults() *AccessLogResponse`

NewAccessLogResponseWithDefaults instantiates a new AccessLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *AccessLogResponse) GetEntries() []AccessLogEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *AccessLogResponse) GetEntriesOk() (*[]AccessLogEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *AccessLogResponse) SetEntries(v []AccessLogEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *AccessLogResponse) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetTotal

`func (o *AccessLogResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccessLogResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccessLogResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AccessLogResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


