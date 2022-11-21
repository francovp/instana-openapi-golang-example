# TestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestResult** | Pointer to [**[]TestResultItem**](TestResultItem.md) |  | [optional] 
**TestResultItems** | Pointer to [**[]TestResultItem**](TestResultItem.md) |  | [optional] 

## Methods

### NewTestResult

`func NewTestResult() *TestResult`

NewTestResult instantiates a new TestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultWithDefaults

`func NewTestResultWithDefaults() *TestResult`

NewTestResultWithDefaults instantiates a new TestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestResult

`func (o *TestResult) GetTestResult() []TestResultItem`

GetTestResult returns the TestResult field if non-nil, zero value otherwise.

### GetTestResultOk

`func (o *TestResult) GetTestResultOk() (*[]TestResultItem, bool)`

GetTestResultOk returns a tuple with the TestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResult

`func (o *TestResult) SetTestResult(v []TestResultItem)`

SetTestResult sets TestResult field to given value.

### HasTestResult

`func (o *TestResult) HasTestResult() bool`

HasTestResult returns a boolean if a field has been set.

### GetTestResultItems

`func (o *TestResult) GetTestResultItems() []TestResultItem`

GetTestResultItems returns the TestResultItems field if non-nil, zero value otherwise.

### GetTestResultItemsOk

`func (o *TestResult) GetTestResultItemsOk() (*[]TestResultItem, bool)`

GetTestResultItemsOk returns a tuple with the TestResultItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResultItems

`func (o *TestResult) SetTestResultItems(v []TestResultItem)`

SetTestResultItems sets TestResultItems field to given value.

### HasTestResultItems

`func (o *TestResult) HasTestResultItems() bool`

HasTestResultItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


