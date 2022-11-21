# WebsiteMonitoringMetricDescription

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

### NewWebsiteMonitoringMetricDescription

`func NewWebsiteMonitoringMetricDescription(aggregations []string, beaconTypes []string, formatter string, label string, metricId string, ) *WebsiteMonitoringMetricDescription`

NewWebsiteMonitoringMetricDescription instantiates a new WebsiteMonitoringMetricDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteMonitoringMetricDescriptionWithDefaults

`func NewWebsiteMonitoringMetricDescriptionWithDefaults() *WebsiteMonitoringMetricDescription`

NewWebsiteMonitoringMetricDescriptionWithDefaults instantiates a new WebsiteMonitoringMetricDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *WebsiteMonitoringMetricDescription) GetAggregations() []string`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *WebsiteMonitoringMetricDescription) GetAggregationsOk() (*[]string, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *WebsiteMonitoringMetricDescription) SetAggregations(v []string)`

SetAggregations sets Aggregations field to given value.


### GetBeaconTypes

`func (o *WebsiteMonitoringMetricDescription) GetBeaconTypes() []string`

GetBeaconTypes returns the BeaconTypes field if non-nil, zero value otherwise.

### GetBeaconTypesOk

`func (o *WebsiteMonitoringMetricDescription) GetBeaconTypesOk() (*[]string, bool)`

GetBeaconTypesOk returns a tuple with the BeaconTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconTypes

`func (o *WebsiteMonitoringMetricDescription) SetBeaconTypes(v []string)`

SetBeaconTypes sets BeaconTypes field to given value.


### GetDefaultAggregation

`func (o *WebsiteMonitoringMetricDescription) GetDefaultAggregation() string`

GetDefaultAggregation returns the DefaultAggregation field if non-nil, zero value otherwise.

### GetDefaultAggregationOk

`func (o *WebsiteMonitoringMetricDescription) GetDefaultAggregationOk() (*string, bool)`

GetDefaultAggregationOk returns a tuple with the DefaultAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAggregation

`func (o *WebsiteMonitoringMetricDescription) SetDefaultAggregation(v string)`

SetDefaultAggregation sets DefaultAggregation field to given value.

### HasDefaultAggregation

`func (o *WebsiteMonitoringMetricDescription) HasDefaultAggregation() bool`

HasDefaultAggregation returns a boolean if a field has been set.

### GetDescription

`func (o *WebsiteMonitoringMetricDescription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebsiteMonitoringMetricDescription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebsiteMonitoringMetricDescription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebsiteMonitoringMetricDescription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormatter

`func (o *WebsiteMonitoringMetricDescription) GetFormatter() string`

GetFormatter returns the Formatter field if non-nil, zero value otherwise.

### GetFormatterOk

`func (o *WebsiteMonitoringMetricDescription) GetFormatterOk() (*string, bool)`

GetFormatterOk returns a tuple with the Formatter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatter

`func (o *WebsiteMonitoringMetricDescription) SetFormatter(v string)`

SetFormatter sets Formatter field to given value.


### GetLabel

`func (o *WebsiteMonitoringMetricDescription) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WebsiteMonitoringMetricDescription) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WebsiteMonitoringMetricDescription) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMetricId

`func (o *WebsiteMonitoringMetricDescription) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *WebsiteMonitoringMetricDescription) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *WebsiteMonitoringMetricDescription) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.


### GetPathToValueInBeacon

`func (o *WebsiteMonitoringMetricDescription) GetPathToValueInBeacon() []string`

GetPathToValueInBeacon returns the PathToValueInBeacon field if non-nil, zero value otherwise.

### GetPathToValueInBeaconOk

`func (o *WebsiteMonitoringMetricDescription) GetPathToValueInBeaconOk() (*[]string, bool)`

GetPathToValueInBeaconOk returns a tuple with the PathToValueInBeacon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathToValueInBeacon

`func (o *WebsiteMonitoringMetricDescription) SetPathToValueInBeacon(v []string)`

SetPathToValueInBeacon sets PathToValueInBeacon field to given value.

### HasPathToValueInBeacon

`func (o *WebsiteMonitoringMetricDescription) HasPathToValueInBeacon() bool`

HasPathToValueInBeacon returns a boolean if a field has been set.

### GetTagName

`func (o *WebsiteMonitoringMetricDescription) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *WebsiteMonitoringMetricDescription) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *WebsiteMonitoringMetricDescription) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *WebsiteMonitoringMetricDescription) HasTagName() bool`

HasTagName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


