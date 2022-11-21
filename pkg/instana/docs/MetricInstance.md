# MetricInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to **bool** |  | [optional] 
**Description** | **string** |  | 
**Formatter** | **string** |  | 
**Label** | **string** |  | 
**MetricId** | **string** |  | 
**PluginId** | **string** |  | 

## Methods

### NewMetricInstance

`func NewMetricInstance(description string, formatter string, label string, metricId string, pluginId string, ) *MetricInstance`

NewMetricInstance instantiates a new MetricInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricInstanceWithDefaults

`func NewMetricInstanceWithDefaults() *MetricInstance`

NewMetricInstanceWithDefaults instantiates a new MetricInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *MetricInstance) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *MetricInstance) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *MetricInstance) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *MetricInstance) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetDescription

`func (o *MetricInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricInstance) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFormatter

`func (o *MetricInstance) GetFormatter() string`

GetFormatter returns the Formatter field if non-nil, zero value otherwise.

### GetFormatterOk

`func (o *MetricInstance) GetFormatterOk() (*string, bool)`

GetFormatterOk returns a tuple with the Formatter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatter

`func (o *MetricInstance) SetFormatter(v string)`

SetFormatter sets Formatter field to given value.


### GetLabel

`func (o *MetricInstance) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MetricInstance) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MetricInstance) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMetricId

`func (o *MetricInstance) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricInstance) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricInstance) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.


### GetPluginId

`func (o *MetricInstance) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *MetricInstance) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *MetricInstance) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


