# SliConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**InitialEvaluationTimestamp** | Pointer to **int64** |  | [optional] 
**MetricConfiguration** | Pointer to [**MetricConfiguration**](MetricConfiguration.md) |  | [optional] 
**SliEntity** | [**SliEntity**](SliEntity.md) |  | 
**SliName** | **string** |  | 

## Methods

### NewSliConfiguration

`func NewSliConfiguration(id string, sliEntity SliEntity, sliName string, ) *SliConfiguration`

NewSliConfiguration instantiates a new SliConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliConfigurationWithDefaults

`func NewSliConfigurationWithDefaults() *SliConfiguration`

NewSliConfigurationWithDefaults instantiates a new SliConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SliConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SliConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SliConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetInitialEvaluationTimestamp

`func (o *SliConfiguration) GetInitialEvaluationTimestamp() int64`

GetInitialEvaluationTimestamp returns the InitialEvaluationTimestamp field if non-nil, zero value otherwise.

### GetInitialEvaluationTimestampOk

`func (o *SliConfiguration) GetInitialEvaluationTimestampOk() (*int64, bool)`

GetInitialEvaluationTimestampOk returns a tuple with the InitialEvaluationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialEvaluationTimestamp

`func (o *SliConfiguration) SetInitialEvaluationTimestamp(v int64)`

SetInitialEvaluationTimestamp sets InitialEvaluationTimestamp field to given value.

### HasInitialEvaluationTimestamp

`func (o *SliConfiguration) HasInitialEvaluationTimestamp() bool`

HasInitialEvaluationTimestamp returns a boolean if a field has been set.

### GetMetricConfiguration

`func (o *SliConfiguration) GetMetricConfiguration() MetricConfiguration`

GetMetricConfiguration returns the MetricConfiguration field if non-nil, zero value otherwise.

### GetMetricConfigurationOk

`func (o *SliConfiguration) GetMetricConfigurationOk() (*MetricConfiguration, bool)`

GetMetricConfigurationOk returns a tuple with the MetricConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricConfiguration

`func (o *SliConfiguration) SetMetricConfiguration(v MetricConfiguration)`

SetMetricConfiguration sets MetricConfiguration field to given value.

### HasMetricConfiguration

`func (o *SliConfiguration) HasMetricConfiguration() bool`

HasMetricConfiguration returns a boolean if a field has been set.

### GetSliEntity

`func (o *SliConfiguration) GetSliEntity() SliEntity`

GetSliEntity returns the SliEntity field if non-nil, zero value otherwise.

### GetSliEntityOk

`func (o *SliConfiguration) GetSliEntityOk() (*SliEntity, bool)`

GetSliEntityOk returns a tuple with the SliEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliEntity

`func (o *SliConfiguration) SetSliEntity(v SliEntity)`

SetSliEntity sets SliEntity field to given value.


### GetSliName

`func (o *SliConfiguration) GetSliName() string`

GetSliName returns the SliName field if non-nil, zero value otherwise.

### GetSliNameOk

`func (o *SliConfiguration) GetSliNameOk() (*string, bool)`

GetSliNameOk returns a tuple with the SliName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliName

`func (o *SliConfiguration) SetSliName(v string)`

SetSliName sets SliName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


