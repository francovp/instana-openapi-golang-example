# GetMobileAppMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**[]MobileAppMonitoringMetricsConfiguration**](MobileAppMonitoringMetricsConfiguration.md) |  | 
**TagFilters** | Pointer to [**[]DeprecatedTagFilter**](DeprecatedTagFilter.md) |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewGetMobileAppMetrics

`func NewGetMobileAppMetrics(metrics []MobileAppMonitoringMetricsConfiguration, type_ string, ) *GetMobileAppMetrics`

NewGetMobileAppMetrics instantiates a new GetMobileAppMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMobileAppMetricsWithDefaults

`func NewGetMobileAppMetricsWithDefaults() *GetMobileAppMetrics`

NewGetMobileAppMetricsWithDefaults instantiates a new GetMobileAppMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *GetMobileAppMetrics) GetMetrics() []MobileAppMonitoringMetricsConfiguration`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetMobileAppMetrics) GetMetricsOk() (*[]MobileAppMonitoringMetricsConfiguration, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetMobileAppMetrics) SetMetrics(v []MobileAppMonitoringMetricsConfiguration)`

SetMetrics sets Metrics field to given value.


### GetTagFilters

`func (o *GetMobileAppMetrics) GetTagFilters() []DeprecatedTagFilter`

GetTagFilters returns the TagFilters field if non-nil, zero value otherwise.

### GetTagFiltersOk

`func (o *GetMobileAppMetrics) GetTagFiltersOk() (*[]DeprecatedTagFilter, bool)`

GetTagFiltersOk returns a tuple with the TagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilters

`func (o *GetMobileAppMetrics) SetTagFilters(v []DeprecatedTagFilter)`

SetTagFilters sets TagFilters field to given value.

### HasTagFilters

`func (o *GetMobileAppMetrics) HasTagFilters() bool`

HasTagFilters returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetMobileAppMetrics) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetMobileAppMetrics) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetMobileAppMetrics) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetMobileAppMetrics) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.

### GetType

`func (o *GetMobileAppMetrics) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMobileAppMetrics) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMobileAppMetrics) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


