# HostAvailabilityRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloseAfter** | Pointer to **int64** |  | [optional] 
**OfflineDuration** | Pointer to **int64** |  | [optional] 
**TagFilter** | Pointer to [**TagFilter**](TagFilter.md) |  | [optional] 

## Methods

### NewHostAvailabilityRule

`func NewHostAvailabilityRule() *HostAvailabilityRule`

NewHostAvailabilityRule instantiates a new HostAvailabilityRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostAvailabilityRuleWithDefaults

`func NewHostAvailabilityRuleWithDefaults() *HostAvailabilityRule`

NewHostAvailabilityRuleWithDefaults instantiates a new HostAvailabilityRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloseAfter

`func (o *HostAvailabilityRule) GetCloseAfter() int64`

GetCloseAfter returns the CloseAfter field if non-nil, zero value otherwise.

### GetCloseAfterOk

`func (o *HostAvailabilityRule) GetCloseAfterOk() (*int64, bool)`

GetCloseAfterOk returns a tuple with the CloseAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseAfter

`func (o *HostAvailabilityRule) SetCloseAfter(v int64)`

SetCloseAfter sets CloseAfter field to given value.

### HasCloseAfter

`func (o *HostAvailabilityRule) HasCloseAfter() bool`

HasCloseAfter returns a boolean if a field has been set.

### GetOfflineDuration

`func (o *HostAvailabilityRule) GetOfflineDuration() int64`

GetOfflineDuration returns the OfflineDuration field if non-nil, zero value otherwise.

### GetOfflineDurationOk

`func (o *HostAvailabilityRule) GetOfflineDurationOk() (*int64, bool)`

GetOfflineDurationOk returns a tuple with the OfflineDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineDuration

`func (o *HostAvailabilityRule) SetOfflineDuration(v int64)`

SetOfflineDuration sets OfflineDuration field to given value.

### HasOfflineDuration

`func (o *HostAvailabilityRule) HasOfflineDuration() bool`

HasOfflineDuration returns a boolean if a field has been set.

### GetTagFilter

`func (o *HostAvailabilityRule) GetTagFilter() TagFilter`

GetTagFilter returns the TagFilter field if non-nil, zero value otherwise.

### GetTagFilterOk

`func (o *HostAvailabilityRule) GetTagFilterOk() (*TagFilter, bool)`

GetTagFilterOk returns a tuple with the TagFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilter

`func (o *HostAvailabilityRule) SetTagFilter(v TagFilter)`

SetTagFilter sets TagFilter field to given value.

### HasTagFilter

`func (o *HostAvailabilityRule) HasTagFilter() bool`

HasTagFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


