# LocationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | **string** |  | 
**SuccessRate** | Pointer to **float64** |  | [optional] 
**SuccessRuns** | Pointer to **int64** |  | [optional] 
**TotalTestRuns** | Pointer to **int64** |  | [optional] 

## Methods

### NewLocationStatus

`func NewLocationStatus(locationId string, ) *LocationStatus`

NewLocationStatus instantiates a new LocationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationStatusWithDefaults

`func NewLocationStatusWithDefaults() *LocationStatus`

NewLocationStatusWithDefaults instantiates a new LocationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *LocationStatus) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *LocationStatus) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *LocationStatus) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetSuccessRate

`func (o *LocationStatus) GetSuccessRate() float64`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *LocationStatus) GetSuccessRateOk() (*float64, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *LocationStatus) SetSuccessRate(v float64)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *LocationStatus) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetSuccessRuns

`func (o *LocationStatus) GetSuccessRuns() int64`

GetSuccessRuns returns the SuccessRuns field if non-nil, zero value otherwise.

### GetSuccessRunsOk

`func (o *LocationStatus) GetSuccessRunsOk() (*int64, bool)`

GetSuccessRunsOk returns a tuple with the SuccessRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRuns

`func (o *LocationStatus) SetSuccessRuns(v int64)`

SetSuccessRuns sets SuccessRuns field to given value.

### HasSuccessRuns

`func (o *LocationStatus) HasSuccessRuns() bool`

HasSuccessRuns returns a boolean if a field has been set.

### GetTotalTestRuns

`func (o *LocationStatus) GetTotalTestRuns() int64`

GetTotalTestRuns returns the TotalTestRuns field if non-nil, zero value otherwise.

### GetTotalTestRunsOk

`func (o *LocationStatus) GetTotalTestRunsOk() (*int64, bool)`

GetTotalTestRunsOk returns a tuple with the TotalTestRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTestRuns

`func (o *LocationStatus) SetTotalTestRuns(v int64)`

SetTotalTestRuns sets TotalTestRuns field to given value.

### HasTotalTestRuns

`func (o *LocationStatus) HasTotalTestRuns() bool`

HasTotalTestRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


