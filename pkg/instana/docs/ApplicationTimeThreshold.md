# ApplicationTimeThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeWindow** | Pointer to **int64** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewApplicationTimeThreshold

`func NewApplicationTimeThreshold(type_ string, ) *ApplicationTimeThreshold`

NewApplicationTimeThreshold instantiates a new ApplicationTimeThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationTimeThresholdWithDefaults

`func NewApplicationTimeThresholdWithDefaults() *ApplicationTimeThreshold`

NewApplicationTimeThresholdWithDefaults instantiates a new ApplicationTimeThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeWindow

`func (o *ApplicationTimeThreshold) GetTimeWindow() int64`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *ApplicationTimeThreshold) GetTimeWindowOk() (*int64, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *ApplicationTimeThreshold) SetTimeWindow(v int64)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *ApplicationTimeThreshold) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetType

`func (o *ApplicationTimeThreshold) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationTimeThreshold) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationTimeThreshold) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


