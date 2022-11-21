# ApplicationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | [**Application**](Application.md) |  | 
**Metrics** | [**map[string][][]float32**](array.md) |  | 

## Methods

### NewApplicationItem

`func NewApplicationItem(application Application, metrics map[string][][]float32, ) *ApplicationItem`

NewApplicationItem instantiates a new ApplicationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationItemWithDefaults

`func NewApplicationItemWithDefaults() *ApplicationItem`

NewApplicationItemWithDefaults instantiates a new ApplicationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ApplicationItem) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApplicationItem) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApplicationItem) SetApplication(v Application)`

SetApplication sets Application field to given value.


### GetMetrics

`func (o *ApplicationItem) GetMetrics() map[string][][]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ApplicationItem) GetMetricsOk() (*map[string][][]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ApplicationItem) SetMetrics(v map[string][][]float32)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


