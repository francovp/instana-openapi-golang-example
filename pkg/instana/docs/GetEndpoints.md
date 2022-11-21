# GetEndpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationBoundaryScope** | Pointer to **string** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**EndpointId** | Pointer to **string** |  | [optional] 
**EndpointTypes** | Pointer to **[]string** |  | [optional] 
**ExcludeSynthetic** | Pointer to **bool** |  | [optional] 
**Metrics** | [**[]AppDataMetricConfiguration**](AppDataMetricConfiguration.md) |  | 
**NameFilter** | Pointer to **string** |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 

## Methods

### NewGetEndpoints

`func NewGetEndpoints(metrics []AppDataMetricConfiguration, ) *GetEndpoints`

NewGetEndpoints instantiates a new GetEndpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEndpointsWithDefaults

`func NewGetEndpointsWithDefaults() *GetEndpoints`

NewGetEndpointsWithDefaults instantiates a new GetEndpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationBoundaryScope

`func (o *GetEndpoints) GetApplicationBoundaryScope() string`

GetApplicationBoundaryScope returns the ApplicationBoundaryScope field if non-nil, zero value otherwise.

### GetApplicationBoundaryScopeOk

`func (o *GetEndpoints) GetApplicationBoundaryScopeOk() (*string, bool)`

GetApplicationBoundaryScopeOk returns a tuple with the ApplicationBoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBoundaryScope

`func (o *GetEndpoints) SetApplicationBoundaryScope(v string)`

SetApplicationBoundaryScope sets ApplicationBoundaryScope field to given value.

### HasApplicationBoundaryScope

`func (o *GetEndpoints) HasApplicationBoundaryScope() bool`

HasApplicationBoundaryScope returns a boolean if a field has been set.

### GetApplicationId

`func (o *GetEndpoints) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GetEndpoints) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GetEndpoints) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *GetEndpoints) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetEndpointId

`func (o *GetEndpoints) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *GetEndpoints) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *GetEndpoints) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *GetEndpoints) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetEndpointTypes

`func (o *GetEndpoints) GetEndpointTypes() []string`

GetEndpointTypes returns the EndpointTypes field if non-nil, zero value otherwise.

### GetEndpointTypesOk

`func (o *GetEndpoints) GetEndpointTypesOk() (*[]string, bool)`

GetEndpointTypesOk returns a tuple with the EndpointTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointTypes

`func (o *GetEndpoints) SetEndpointTypes(v []string)`

SetEndpointTypes sets EndpointTypes field to given value.

### HasEndpointTypes

`func (o *GetEndpoints) HasEndpointTypes() bool`

HasEndpointTypes returns a boolean if a field has been set.

### GetExcludeSynthetic

`func (o *GetEndpoints) GetExcludeSynthetic() bool`

GetExcludeSynthetic returns the ExcludeSynthetic field if non-nil, zero value otherwise.

### GetExcludeSyntheticOk

`func (o *GetEndpoints) GetExcludeSyntheticOk() (*bool, bool)`

GetExcludeSyntheticOk returns a tuple with the ExcludeSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSynthetic

`func (o *GetEndpoints) SetExcludeSynthetic(v bool)`

SetExcludeSynthetic sets ExcludeSynthetic field to given value.

### HasExcludeSynthetic

`func (o *GetEndpoints) HasExcludeSynthetic() bool`

HasExcludeSynthetic returns a boolean if a field has been set.

### GetMetrics

`func (o *GetEndpoints) GetMetrics() []AppDataMetricConfiguration`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetEndpoints) GetMetricsOk() (*[]AppDataMetricConfiguration, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetEndpoints) SetMetrics(v []AppDataMetricConfiguration)`

SetMetrics sets Metrics field to given value.


### GetNameFilter

`func (o *GetEndpoints) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *GetEndpoints) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *GetEndpoints) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *GetEndpoints) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetOrder

`func (o *GetEndpoints) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetEndpoints) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetEndpoints) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetEndpoints) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPagination

`func (o *GetEndpoints) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetEndpoints) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetEndpoints) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetEndpoints) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetServiceId

`func (o *GetEndpoints) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GetEndpoints) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GetEndpoints) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *GetEndpoints) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetEndpoints) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetEndpoints) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetEndpoints) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetEndpoints) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


