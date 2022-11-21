# TestResultListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**map[string][][]float32**](array.md) |  | 
**TestResultCommonProperties** | [**TestResultCommonProperties**](TestResultCommonProperties.md) |  | 

## Methods

### NewTestResultListItem

`func NewTestResultListItem(metrics map[string][][]float32, testResultCommonProperties TestResultCommonProperties, ) *TestResultListItem`

NewTestResultListItem instantiates a new TestResultListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultListItemWithDefaults

`func NewTestResultListItemWithDefaults() *TestResultListItem`

NewTestResultListItemWithDefaults instantiates a new TestResultListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *TestResultListItem) GetMetrics() map[string][][]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *TestResultListItem) GetMetricsOk() (*map[string][][]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *TestResultListItem) SetMetrics(v map[string][][]float32)`

SetMetrics sets Metrics field to given value.


### GetTestResultCommonProperties

`func (o *TestResultListItem) GetTestResultCommonProperties() TestResultCommonProperties`

GetTestResultCommonProperties returns the TestResultCommonProperties field if non-nil, zero value otherwise.

### GetTestResultCommonPropertiesOk

`func (o *TestResultListItem) GetTestResultCommonPropertiesOk() (*TestResultCommonProperties, bool)`

GetTestResultCommonPropertiesOk returns a tuple with the TestResultCommonProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultCommonProperties

`func (o *TestResultListItem) SetTestResultCommonProperties(v TestResultCommonProperties)`

SetTestResultCommonProperties sets TestResultCommonProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


