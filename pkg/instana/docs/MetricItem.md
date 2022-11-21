# MetricItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **int64** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**map[string][][]float32**](array.md) |  | [optional] 
**Plugin** | Pointer to **string** |  | [optional] 
**SnapshotId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**To** | Pointer to **int64** |  | [optional] 

## Methods

### NewMetricItem

`func NewMetricItem() *MetricItem`

NewMetricItem instantiates a new MetricItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricItemWithDefaults

`func NewMetricItemWithDefaults() *MetricItem`

NewMetricItemWithDefaults instantiates a new MetricItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *MetricItem) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MetricItem) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MetricItem) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MetricItem) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHost

`func (o *MetricItem) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MetricItem) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MetricItem) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MetricItem) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetLabel

`func (o *MetricItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MetricItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MetricItem) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MetricItem) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetrics

`func (o *MetricItem) GetMetrics() map[string][][]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricItem) GetMetricsOk() (*map[string][][]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricItem) SetMetrics(v map[string][][]float32)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MetricItem) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetPlugin

`func (o *MetricItem) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *MetricItem) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *MetricItem) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *MetricItem) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetSnapshotId

`func (o *MetricItem) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *MetricItem) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *MetricItem) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *MetricItem) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetTags

`func (o *MetricItem) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricItem) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricItem) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTo

`func (o *MetricItem) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MetricItem) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MetricItem) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *MetricItem) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


