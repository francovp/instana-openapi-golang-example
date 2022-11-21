# AccessLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Email** | **string** |  | 
**FullName** | **string** |  | 
**TenantId** | **string** |  | 
**TenantUnitId** | Pointer to **string** |  | [optional] 
**Timestamp** | **int64** |  | 

## Methods

### NewAccessLogEntry

`func NewAccessLogEntry(action string, email string, fullName string, tenantId string, timestamp int64, ) *AccessLogEntry`

NewAccessLogEntry instantiates a new AccessLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessLogEntryWithDefaults

`func NewAccessLogEntryWithDefaults() *AccessLogEntry`

NewAccessLogEntryWithDefaults instantiates a new AccessLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AccessLogEntry) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AccessLogEntry) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AccessLogEntry) SetAction(v string)`

SetAction sets Action field to given value.


### GetEmail

`func (o *AccessLogEntry) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccessLogEntry) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccessLogEntry) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFullName

`func (o *AccessLogEntry) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *AccessLogEntry) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *AccessLogEntry) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetTenantId

`func (o *AccessLogEntry) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AccessLogEntry) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AccessLogEntry) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTenantUnitId

`func (o *AccessLogEntry) GetTenantUnitId() string`

GetTenantUnitId returns the TenantUnitId field if non-nil, zero value otherwise.

### GetTenantUnitIdOk

`func (o *AccessLogEntry) GetTenantUnitIdOk() (*string, bool)`

GetTenantUnitIdOk returns a tuple with the TenantUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUnitId

`func (o *AccessLogEntry) SetTenantUnitId(v string)`

SetTenantUnitId sets TenantUnitId field to given value.

### HasTenantUnitId

`func (o *AccessLogEntry) HasTenantUnitId() bool`

HasTenantUnitId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AccessLogEntry) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccessLogEntry) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccessLogEntry) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


