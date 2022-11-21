# SliReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorBudgetRemaining** | Pointer to **int32** |  | [optional] 
**FromTimestamp** | Pointer to **int64** |  | [optional] 
**Sli** | Pointer to **float64** |  | [optional] 
**Slo** | Pointer to **float64** |  | [optional] 
**ToTimestamp** | Pointer to **int64** |  | [optional] 
**TotalErrorBudget** | Pointer to **int32** |  | [optional] 
**ViolationDistribution** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewSliReport

`func NewSliReport() *SliReport`

NewSliReport instantiates a new SliReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliReportWithDefaults

`func NewSliReportWithDefaults() *SliReport`

NewSliReportWithDefaults instantiates a new SliReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorBudgetRemaining

`func (o *SliReport) GetErrorBudgetRemaining() int32`

GetErrorBudgetRemaining returns the ErrorBudgetRemaining field if non-nil, zero value otherwise.

### GetErrorBudgetRemainingOk

`func (o *SliReport) GetErrorBudgetRemainingOk() (*int32, bool)`

GetErrorBudgetRemainingOk returns a tuple with the ErrorBudgetRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorBudgetRemaining

`func (o *SliReport) SetErrorBudgetRemaining(v int32)`

SetErrorBudgetRemaining sets ErrorBudgetRemaining field to given value.

### HasErrorBudgetRemaining

`func (o *SliReport) HasErrorBudgetRemaining() bool`

HasErrorBudgetRemaining returns a boolean if a field has been set.

### GetFromTimestamp

`func (o *SliReport) GetFromTimestamp() int64`

GetFromTimestamp returns the FromTimestamp field if non-nil, zero value otherwise.

### GetFromTimestampOk

`func (o *SliReport) GetFromTimestampOk() (*int64, bool)`

GetFromTimestampOk returns a tuple with the FromTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTimestamp

`func (o *SliReport) SetFromTimestamp(v int64)`

SetFromTimestamp sets FromTimestamp field to given value.

### HasFromTimestamp

`func (o *SliReport) HasFromTimestamp() bool`

HasFromTimestamp returns a boolean if a field has been set.

### GetSli

`func (o *SliReport) GetSli() float64`

GetSli returns the Sli field if non-nil, zero value otherwise.

### GetSliOk

`func (o *SliReport) GetSliOk() (*float64, bool)`

GetSliOk returns a tuple with the Sli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSli

`func (o *SliReport) SetSli(v float64)`

SetSli sets Sli field to given value.

### HasSli

`func (o *SliReport) HasSli() bool`

HasSli returns a boolean if a field has been set.

### GetSlo

`func (o *SliReport) GetSlo() float64`

GetSlo returns the Slo field if non-nil, zero value otherwise.

### GetSloOk

`func (o *SliReport) GetSloOk() (*float64, bool)`

GetSloOk returns a tuple with the Slo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlo

`func (o *SliReport) SetSlo(v float64)`

SetSlo sets Slo field to given value.

### HasSlo

`func (o *SliReport) HasSlo() bool`

HasSlo returns a boolean if a field has been set.

### GetToTimestamp

`func (o *SliReport) GetToTimestamp() int64`

GetToTimestamp returns the ToTimestamp field if non-nil, zero value otherwise.

### GetToTimestampOk

`func (o *SliReport) GetToTimestampOk() (*int64, bool)`

GetToTimestampOk returns a tuple with the ToTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTimestamp

`func (o *SliReport) SetToTimestamp(v int64)`

SetToTimestamp sets ToTimestamp field to given value.

### HasToTimestamp

`func (o *SliReport) HasToTimestamp() bool`

HasToTimestamp returns a boolean if a field has been set.

### GetTotalErrorBudget

`func (o *SliReport) GetTotalErrorBudget() int32`

GetTotalErrorBudget returns the TotalErrorBudget field if non-nil, zero value otherwise.

### GetTotalErrorBudgetOk

`func (o *SliReport) GetTotalErrorBudgetOk() (*int32, bool)`

GetTotalErrorBudgetOk returns a tuple with the TotalErrorBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalErrorBudget

`func (o *SliReport) SetTotalErrorBudget(v int32)`

SetTotalErrorBudget sets TotalErrorBudget field to given value.

### HasTotalErrorBudget

`func (o *SliReport) HasTotalErrorBudget() bool`

HasTotalErrorBudget returns a boolean if a field has been set.

### GetViolationDistribution

`func (o *SliReport) GetViolationDistribution() map[string]int32`

GetViolationDistribution returns the ViolationDistribution field if non-nil, zero value otherwise.

### GetViolationDistributionOk

`func (o *SliReport) GetViolationDistributionOk() (*map[string]int32, bool)`

GetViolationDistributionOk returns a tuple with the ViolationDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationDistribution

`func (o *SliReport) SetViolationDistribution(v map[string]int32)`

SetViolationDistribution sets ViolationDistribution field to given value.

### HasViolationDistribution

`func (o *SliReport) HasViolationDistribution() bool`

HasViolationDistribution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


