# ServiceMapConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | Pointer to **int64** |  | [optional] 
**ErrorRate** | Pointer to **float64** |  | [optional] 
**From** | **string** |  | 
**Latency** | Pointer to **float64** |  | [optional] 
**To** | **string** |  | 

## Methods

### NewServiceMapConnection

`func NewServiceMapConnection(from string, to string, ) *ServiceMapConnection`

NewServiceMapConnection instantiates a new ServiceMapConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMapConnectionWithDefaults

`func NewServiceMapConnectionWithDefaults() *ServiceMapConnection`

NewServiceMapConnectionWithDefaults instantiates a new ServiceMapConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *ServiceMapConnection) GetCalls() int64`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *ServiceMapConnection) GetCallsOk() (*int64, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *ServiceMapConnection) SetCalls(v int64)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *ServiceMapConnection) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetErrorRate

`func (o *ServiceMapConnection) GetErrorRate() float64`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *ServiceMapConnection) GetErrorRateOk() (*float64, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *ServiceMapConnection) SetErrorRate(v float64)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *ServiceMapConnection) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetFrom

`func (o *ServiceMapConnection) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ServiceMapConnection) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ServiceMapConnection) SetFrom(v string)`

SetFrom sets From field to given value.


### GetLatency

`func (o *ServiceMapConnection) GetLatency() float64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *ServiceMapConnection) GetLatencyOk() (*float64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *ServiceMapConnection) SetLatency(v float64)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *ServiceMapConnection) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetTo

`func (o *ServiceMapConnection) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ServiceMapConnection) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ServiceMapConnection) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


