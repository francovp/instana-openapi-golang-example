# TestResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **[]string** |  | [optional] 
**Metrics** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**TestId** | **string** |  | 
**TestName** | Pointer to **string** |  | [optional] 

## Methods

### NewTestResultItem

`func NewTestResultItem(testId string, ) *TestResultItem`

NewTestResultItem instantiates a new TestResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultItemWithDefaults

`func NewTestResultItemWithDefaults() *TestResultItem`

NewTestResultItemWithDefaults instantiates a new TestResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *TestResultItem) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *TestResultItem) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *TestResultItem) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *TestResultItem) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetLocationId

`func (o *TestResultItem) GetLocationId() []string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *TestResultItem) GetLocationIdOk() (*[]string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *TestResultItem) SetLocationId(v []string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *TestResultItem) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetMetrics

`func (o *TestResultItem) GetMetrics() []map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *TestResultItem) GetMetricsOk() (*[]map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *TestResultItem) SetMetrics(v []map[string]interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *TestResultItem) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetServiceId

`func (o *TestResultItem) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *TestResultItem) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *TestResultItem) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *TestResultItem) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetTestId

`func (o *TestResultItem) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *TestResultItem) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *TestResultItem) SetTestId(v string)`

SetTestId sets TestId field to given value.


### GetTestName

`func (o *TestResultItem) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *TestResultItem) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *TestResultItem) SetTestName(v string)`

SetTestName sets TestName field to given value.

### HasTestName

`func (o *TestResultItem) HasTestName() bool`

HasTestName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


