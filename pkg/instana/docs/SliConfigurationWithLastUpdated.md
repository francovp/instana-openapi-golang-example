# SliConfigurationWithLastUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**InitialEvaluationTimestamp** | Pointer to **int64** |  | [optional] 
**LastUpdated** | Pointer to **int64** |  | [optional] 
**MetricConfiguration** | Pointer to [**MetricConfiguration**](MetricConfiguration.md) |  | [optional] 
**SliEntity** | [**SliEntity**](SliEntity.md) |  | 
**SliName** | **string** |  | 

## Methods

### NewSliConfigurationWithLastUpdated

`func NewSliConfigurationWithLastUpdated(id string, sliEntity SliEntity, sliName string, ) *SliConfigurationWithLastUpdated`

NewSliConfigurationWithLastUpdated instantiates a new SliConfigurationWithLastUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliConfigurationWithLastUpdatedWithDefaults

`func NewSliConfigurationWithLastUpdatedWithDefaults() *SliConfigurationWithLastUpdated`

NewSliConfigurationWithLastUpdatedWithDefaults instantiates a new SliConfigurationWithLastUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SliConfigurationWithLastUpdated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SliConfigurationWithLastUpdated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SliConfigurationWithLastUpdated) SetId(v string)`

SetId sets Id field to given value.


### GetInitialEvaluationTimestamp

`func (o *SliConfigurationWithLastUpdated) GetInitialEvaluationTimestamp() int64`

GetInitialEvaluationTimestamp returns the InitialEvaluationTimestamp field if non-nil, zero value otherwise.

### GetInitialEvaluationTimestampOk

`func (o *SliConfigurationWithLastUpdated) GetInitialEvaluationTimestampOk() (*int64, bool)`

GetInitialEvaluationTimestampOk returns a tuple with the InitialEvaluationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialEvaluationTimestamp

`func (o *SliConfigurationWithLastUpdated) SetInitialEvaluationTimestamp(v int64)`

SetInitialEvaluationTimestamp sets InitialEvaluationTimestamp field to given value.

### HasInitialEvaluationTimestamp

`func (o *SliConfigurationWithLastUpdated) HasInitialEvaluationTimestamp() bool`

HasInitialEvaluationTimestamp returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SliConfigurationWithLastUpdated) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SliConfigurationWithLastUpdated) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SliConfigurationWithLastUpdated) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SliConfigurationWithLastUpdated) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMetricConfiguration

`func (o *SliConfigurationWithLastUpdated) GetMetricConfiguration() MetricConfiguration`

GetMetricConfiguration returns the MetricConfiguration field if non-nil, zero value otherwise.

### GetMetricConfigurationOk

`func (o *SliConfigurationWithLastUpdated) GetMetricConfigurationOk() (*MetricConfiguration, bool)`

GetMetricConfigurationOk returns a tuple with the MetricConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricConfiguration

`func (o *SliConfigurationWithLastUpdated) SetMetricConfiguration(v MetricConfiguration)`

SetMetricConfiguration sets MetricConfiguration field to given value.

### HasMetricConfiguration

`func (o *SliConfigurationWithLastUpdated) HasMetricConfiguration() bool`

HasMetricConfiguration returns a boolean if a field has been set.

### GetSliEntity

`func (o *SliConfigurationWithLastUpdated) GetSliEntity() SliEntity`

GetSliEntity returns the SliEntity field if non-nil, zero value otherwise.

### GetSliEntityOk

`func (o *SliConfigurationWithLastUpdated) GetSliEntityOk() (*SliEntity, bool)`

GetSliEntityOk returns a tuple with the SliEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliEntity

`func (o *SliConfigurationWithLastUpdated) SetSliEntity(v SliEntity)`

SetSliEntity sets SliEntity field to given value.


### GetSliName

`func (o *SliConfigurationWithLastUpdated) GetSliName() string`

GetSliName returns the SliName field if non-nil, zero value otherwise.

### GetSliNameOk

`func (o *SliConfigurationWithLastUpdated) GetSliNameOk() (*string, bool)`

GetSliNameOk returns a tuple with the SliName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliName

`func (o *SliConfigurationWithLastUpdated) SetSliName(v string)`

SetSliName sets SliName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


