# SyntheticTypeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarkSyntheticCall** | **bool** |  | 
**Retries** | Pointer to **int32** |  | [optional] 
**RetryInterval** | Pointer to **int32** |  | [optional] 
**SyntheticType** | **string** |  | 
**Timeout** | Pointer to **string** |  | [optional] 

## Methods

### NewSyntheticTypeConfiguration

`func NewSyntheticTypeConfiguration(markSyntheticCall bool, syntheticType string, ) *SyntheticTypeConfiguration`

NewSyntheticTypeConfiguration instantiates a new SyntheticTypeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticTypeConfigurationWithDefaults

`func NewSyntheticTypeConfigurationWithDefaults() *SyntheticTypeConfiguration`

NewSyntheticTypeConfigurationWithDefaults instantiates a new SyntheticTypeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarkSyntheticCall

`func (o *SyntheticTypeConfiguration) GetMarkSyntheticCall() bool`

GetMarkSyntheticCall returns the MarkSyntheticCall field if non-nil, zero value otherwise.

### GetMarkSyntheticCallOk

`func (o *SyntheticTypeConfiguration) GetMarkSyntheticCallOk() (*bool, bool)`

GetMarkSyntheticCallOk returns a tuple with the MarkSyntheticCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkSyntheticCall

`func (o *SyntheticTypeConfiguration) SetMarkSyntheticCall(v bool)`

SetMarkSyntheticCall sets MarkSyntheticCall field to given value.


### GetRetries

`func (o *SyntheticTypeConfiguration) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *SyntheticTypeConfiguration) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *SyntheticTypeConfiguration) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *SyntheticTypeConfiguration) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetRetryInterval

`func (o *SyntheticTypeConfiguration) GetRetryInterval() int32`

GetRetryInterval returns the RetryInterval field if non-nil, zero value otherwise.

### GetRetryIntervalOk

`func (o *SyntheticTypeConfiguration) GetRetryIntervalOk() (*int32, bool)`

GetRetryIntervalOk returns a tuple with the RetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryInterval

`func (o *SyntheticTypeConfiguration) SetRetryInterval(v int32)`

SetRetryInterval sets RetryInterval field to given value.

### HasRetryInterval

`func (o *SyntheticTypeConfiguration) HasRetryInterval() bool`

HasRetryInterval returns a boolean if a field has been set.

### GetSyntheticType

`func (o *SyntheticTypeConfiguration) GetSyntheticType() string`

GetSyntheticType returns the SyntheticType field if non-nil, zero value otherwise.

### GetSyntheticTypeOk

`func (o *SyntheticTypeConfiguration) GetSyntheticTypeOk() (*string, bool)`

GetSyntheticTypeOk returns a tuple with the SyntheticType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticType

`func (o *SyntheticTypeConfiguration) SetSyntheticType(v string)`

SetSyntheticType sets SyntheticType field to given value.


### GetTimeout

`func (o *SyntheticTypeConfiguration) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SyntheticTypeConfiguration) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SyntheticTypeConfiguration) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SyntheticTypeConfiguration) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


