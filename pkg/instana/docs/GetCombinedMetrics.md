# GetCombinedMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | **[]string** |  | 
**Plugin** | **string** |  | 
**Query** | Pointer to **string** |  | [optional] 
**Rollup** | Pointer to **int32** |  | [optional] 
**SnapshotIds** | Pointer to **[]string** |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 

## Methods

### NewGetCombinedMetrics

`func NewGetCombinedMetrics(metrics []string, plugin string, ) *GetCombinedMetrics`

NewGetCombinedMetrics instantiates a new GetCombinedMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCombinedMetricsWithDefaults

`func NewGetCombinedMetricsWithDefaults() *GetCombinedMetrics`

NewGetCombinedMetricsWithDefaults instantiates a new GetCombinedMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *GetCombinedMetrics) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetCombinedMetrics) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetCombinedMetrics) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.


### GetPlugin

`func (o *GetCombinedMetrics) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *GetCombinedMetrics) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *GetCombinedMetrics) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.


### GetQuery

`func (o *GetCombinedMetrics) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GetCombinedMetrics) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GetCombinedMetrics) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GetCombinedMetrics) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRollup

`func (o *GetCombinedMetrics) GetRollup() int32`

GetRollup returns the Rollup field if non-nil, zero value otherwise.

### GetRollupOk

`func (o *GetCombinedMetrics) GetRollupOk() (*int32, bool)`

GetRollupOk returns a tuple with the Rollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollup

`func (o *GetCombinedMetrics) SetRollup(v int32)`

SetRollup sets Rollup field to given value.

### HasRollup

`func (o *GetCombinedMetrics) HasRollup() bool`

HasRollup returns a boolean if a field has been set.

### GetSnapshotIds

`func (o *GetCombinedMetrics) GetSnapshotIds() []string`

GetSnapshotIds returns the SnapshotIds field if non-nil, zero value otherwise.

### GetSnapshotIdsOk

`func (o *GetCombinedMetrics) GetSnapshotIdsOk() (*[]string, bool)`

GetSnapshotIdsOk returns a tuple with the SnapshotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIds

`func (o *GetCombinedMetrics) SetSnapshotIds(v []string)`

SetSnapshotIds sets SnapshotIds field to given value.

### HasSnapshotIds

`func (o *GetCombinedMetrics) HasSnapshotIds() bool`

HasSnapshotIds returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetCombinedMetrics) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetCombinedMetrics) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetCombinedMetrics) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetCombinedMetrics) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


