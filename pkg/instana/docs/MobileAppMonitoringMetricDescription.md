# MobileAppMonitoringMetricDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | **[]string** |  | 
**BeaconTypes** | **[]string** |  | 
**DefaultAggregation** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Formatter** | **string** |  | 
**Label** | **string** |  | 
**MetricId** | **string** |  | 
**PathToValueInBeacon** | Pointer to **[]string** |  | [optional] 
**TagName** | Pointer to **string** |  | [optional] 

## Methods

### NewMobileAppMonitoringMetricDescription

`func NewMobileAppMonitoringMetricDescription(aggregations []string, beaconTypes []string, formatter string, label string, metricId string, ) *MobileAppMonitoringMetricDescription`

NewMobileAppMonitoringMetricDescription instantiates a new MobileAppMonitoringMetricDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileAppMonitoringMetricDescriptionWithDefaults

`func NewMobileAppMonitoringMetricDescriptionWithDefaults() *MobileAppMonitoringMetricDescription`

NewMobileAppMonitoringMetricDescriptionWithDefaults instantiates a new MobileAppMonitoringMetricDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *MobileAppMonitoringMetricDescription) GetAggregations() []string`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *MobileAppMonitoringMetricDescription) GetAggregationsOk() (*[]string, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *MobileAppMonitoringMetricDescription) SetAggregations(v []string)`

SetAggregations sets Aggregations field to given value.


### GetBeaconTypes

`func (o *MobileAppMonitoringMetricDescription) GetBeaconTypes() []string`

GetBeaconTypes returns the BeaconTypes field if non-nil, zero value otherwise.

### GetBeaconTypesOk

`func (o *MobileAppMonitoringMetricDescription) GetBeaconTypesOk() (*[]string, bool)`

GetBeaconTypesOk returns a tuple with the BeaconTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconTypes

`func (o *MobileAppMonitoringMetricDescription) SetBeaconTypes(v []string)`

SetBeaconTypes sets BeaconTypes field to given value.


### GetDefaultAggregation

`func (o *MobileAppMonitoringMetricDescription) GetDefaultAggregation() string`

GetDefaultAggregation returns the DefaultAggregation field if non-nil, zero value otherwise.

### GetDefaultAggregationOk

`func (o *MobileAppMonitoringMetricDescription) GetDefaultAggregationOk() (*string, bool)`

GetDefaultAggregationOk returns a tuple with the DefaultAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAggregation

`func (o *MobileAppMonitoringMetricDescription) SetDefaultAggregation(v string)`

SetDefaultAggregation sets DefaultAggregation field to given value.

### HasDefaultAggregation

`func (o *MobileAppMonitoringMetricDescription) HasDefaultAggregation() bool`

HasDefaultAggregation returns a boolean if a field has been set.

### GetDescription

`func (o *MobileAppMonitoringMetricDescription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MobileAppMonitoringMetricDescription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MobileAppMonitoringMetricDescription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MobileAppMonitoringMetricDescription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormatter

`func (o *MobileAppMonitoringMetricDescription) GetFormatter() string`

GetFormatter returns the Formatter field if non-nil, zero value otherwise.

### GetFormatterOk

`func (o *MobileAppMonitoringMetricDescription) GetFormatterOk() (*string, bool)`

GetFormatterOk returns a tuple with the Formatter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatter

`func (o *MobileAppMonitoringMetricDescription) SetFormatter(v string)`

SetFormatter sets Formatter field to given value.


### GetLabel

`func (o *MobileAppMonitoringMetricDescription) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MobileAppMonitoringMetricDescription) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MobileAppMonitoringMetricDescription) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMetricId

`func (o *MobileAppMonitoringMetricDescription) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MobileAppMonitoringMetricDescription) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MobileAppMonitoringMetricDescription) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.


### GetPathToValueInBeacon

`func (o *MobileAppMonitoringMetricDescription) GetPathToValueInBeacon() []string`

GetPathToValueInBeacon returns the PathToValueInBeacon field if non-nil, zero value otherwise.

### GetPathToValueInBeaconOk

`func (o *MobileAppMonitoringMetricDescription) GetPathToValueInBeaconOk() (*[]string, bool)`

GetPathToValueInBeaconOk returns a tuple with the PathToValueInBeacon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathToValueInBeacon

`func (o *MobileAppMonitoringMetricDescription) SetPathToValueInBeacon(v []string)`

SetPathToValueInBeacon sets PathToValueInBeacon field to given value.

### HasPathToValueInBeacon

`func (o *MobileAppMonitoringMetricDescription) HasPathToValueInBeacon() bool`

HasPathToValueInBeacon returns a boolean if a field has been set.

### GetTagName

`func (o *MobileAppMonitoringMetricDescription) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *MobileAppMonitoringMetricDescription) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *MobileAppMonitoringMetricDescription) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *MobileAppMonitoringMetricDescription) HasTagName() bool`

HasTagName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


