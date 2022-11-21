# Trace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int64** |  | [optional] 
**Endpoint** | Pointer to [**Endpoint**](Endpoint.md) |  | [optional] 
**Erroneous** | Pointer to **bool** |  | [optional] 
**Id** | **string** |  | 
**Label** | **string** |  | 
**Service** | Pointer to [**Service**](Service.md) |  | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewTrace

`func NewTrace(id string, label string, ) *Trace`

NewTrace instantiates a new Trace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceWithDefaults

`func NewTraceWithDefaults() *Trace`

NewTraceWithDefaults instantiates a new Trace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *Trace) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Trace) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Trace) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Trace) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndpoint

`func (o *Trace) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Trace) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Trace) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *Trace) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetErroneous

`func (o *Trace) GetErroneous() bool`

GetErroneous returns the Erroneous field if non-nil, zero value otherwise.

### GetErroneousOk

`func (o *Trace) GetErroneousOk() (*bool, bool)`

GetErroneousOk returns a tuple with the Erroneous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErroneous

`func (o *Trace) SetErroneous(v bool)`

SetErroneous sets Erroneous field to given value.

### HasErroneous

`func (o *Trace) HasErroneous() bool`

HasErroneous returns a boolean if a field has been set.

### GetId

`func (o *Trace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trace) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *Trace) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Trace) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Trace) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetService

`func (o *Trace) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Trace) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Trace) SetService(v Service)`

SetService sets Service field to given value.

### HasService

`func (o *Trace) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStartTime

`func (o *Trace) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Trace) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Trace) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Trace) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


