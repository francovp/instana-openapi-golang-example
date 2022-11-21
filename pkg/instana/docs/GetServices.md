# GetServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationBoundaryScope** | Pointer to **string** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**ContextScope** | Pointer to **string** |  | [optional] 
**Metrics** | [**[]AppDataMetricConfiguration**](AppDataMetricConfiguration.md) |  | 
**NameFilter** | Pointer to **string** |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**Technologies** | Pointer to **[]string** |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 

## Methods

### NewGetServices

`func NewGetServices(metrics []AppDataMetricConfiguration, ) *GetServices`

NewGetServices instantiates a new GetServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServicesWithDefaults

`func NewGetServicesWithDefaults() *GetServices`

NewGetServicesWithDefaults instantiates a new GetServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationBoundaryScope

`func (o *GetServices) GetApplicationBoundaryScope() string`

GetApplicationBoundaryScope returns the ApplicationBoundaryScope field if non-nil, zero value otherwise.

### GetApplicationBoundaryScopeOk

`func (o *GetServices) GetApplicationBoundaryScopeOk() (*string, bool)`

GetApplicationBoundaryScopeOk returns a tuple with the ApplicationBoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBoundaryScope

`func (o *GetServices) SetApplicationBoundaryScope(v string)`

SetApplicationBoundaryScope sets ApplicationBoundaryScope field to given value.

### HasApplicationBoundaryScope

`func (o *GetServices) HasApplicationBoundaryScope() bool`

HasApplicationBoundaryScope returns a boolean if a field has been set.

### GetApplicationId

`func (o *GetServices) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GetServices) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GetServices) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *GetServices) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetContextScope

`func (o *GetServices) GetContextScope() string`

GetContextScope returns the ContextScope field if non-nil, zero value otherwise.

### GetContextScopeOk

`func (o *GetServices) GetContextScopeOk() (*string, bool)`

GetContextScopeOk returns a tuple with the ContextScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextScope

`func (o *GetServices) SetContextScope(v string)`

SetContextScope sets ContextScope field to given value.

### HasContextScope

`func (o *GetServices) HasContextScope() bool`

HasContextScope returns a boolean if a field has been set.

### GetMetrics

`func (o *GetServices) GetMetrics() []AppDataMetricConfiguration`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetServices) GetMetricsOk() (*[]AppDataMetricConfiguration, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetServices) SetMetrics(v []AppDataMetricConfiguration)`

SetMetrics sets Metrics field to given value.


### GetNameFilter

`func (o *GetServices) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *GetServices) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *GetServices) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *GetServices) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetOrder

`func (o *GetServices) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetServices) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetServices) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetServices) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPagination

`func (o *GetServices) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetServices) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetServices) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetServices) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetServiceId

`func (o *GetServices) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GetServices) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GetServices) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *GetServices) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetTechnologies

`func (o *GetServices) GetTechnologies() []string`

GetTechnologies returns the Technologies field if non-nil, zero value otherwise.

### GetTechnologiesOk

`func (o *GetServices) GetTechnologiesOk() (*[]string, bool)`

GetTechnologiesOk returns a tuple with the Technologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologies

`func (o *GetServices) SetTechnologies(v []string)`

SetTechnologies sets Technologies field to given value.

### HasTechnologies

`func (o *GetServices) HasTechnologies() bool`

HasTechnologies returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetServices) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetServices) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetServices) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetServices) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


