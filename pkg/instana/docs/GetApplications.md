# GetApplications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationBoundaryScope** | Pointer to **string** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**EndpointId** | Pointer to **string** |  | [optional] 
**EndpointTypes** | Pointer to **[]string** |  | [optional] 
**Metrics** | [**[]AppDataMetricConfiguration**](AppDataMetricConfiguration.md) |  | 
**NameFilter** | Pointer to **string** |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**Technologies** | Pointer to **[]string** |  | [optional] 
**TimeFrame** | Pointer to [**TimeFrame**](TimeFrame.md) |  | [optional] 

## Methods

### NewGetApplications

`func NewGetApplications(metrics []AppDataMetricConfiguration, ) *GetApplications`

NewGetApplications instantiates a new GetApplications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApplicationsWithDefaults

`func NewGetApplicationsWithDefaults() *GetApplications`

NewGetApplicationsWithDefaults instantiates a new GetApplications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationBoundaryScope

`func (o *GetApplications) GetApplicationBoundaryScope() string`

GetApplicationBoundaryScope returns the ApplicationBoundaryScope field if non-nil, zero value otherwise.

### GetApplicationBoundaryScopeOk

`func (o *GetApplications) GetApplicationBoundaryScopeOk() (*string, bool)`

GetApplicationBoundaryScopeOk returns a tuple with the ApplicationBoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBoundaryScope

`func (o *GetApplications) SetApplicationBoundaryScope(v string)`

SetApplicationBoundaryScope sets ApplicationBoundaryScope field to given value.

### HasApplicationBoundaryScope

`func (o *GetApplications) HasApplicationBoundaryScope() bool`

HasApplicationBoundaryScope returns a boolean if a field has been set.

### GetApplicationId

`func (o *GetApplications) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GetApplications) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GetApplications) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *GetApplications) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetEndpointId

`func (o *GetApplications) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *GetApplications) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *GetApplications) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *GetApplications) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetEndpointTypes

`func (o *GetApplications) GetEndpointTypes() []string`

GetEndpointTypes returns the EndpointTypes field if non-nil, zero value otherwise.

### GetEndpointTypesOk

`func (o *GetApplications) GetEndpointTypesOk() (*[]string, bool)`

GetEndpointTypesOk returns a tuple with the EndpointTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointTypes

`func (o *GetApplications) SetEndpointTypes(v []string)`

SetEndpointTypes sets EndpointTypes field to given value.

### HasEndpointTypes

`func (o *GetApplications) HasEndpointTypes() bool`

HasEndpointTypes returns a boolean if a field has been set.

### GetMetrics

`func (o *GetApplications) GetMetrics() []AppDataMetricConfiguration`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetApplications) GetMetricsOk() (*[]AppDataMetricConfiguration, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetApplications) SetMetrics(v []AppDataMetricConfiguration)`

SetMetrics sets Metrics field to given value.


### GetNameFilter

`func (o *GetApplications) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *GetApplications) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *GetApplications) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *GetApplications) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetOrder

`func (o *GetApplications) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetApplications) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetApplications) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetApplications) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPagination

`func (o *GetApplications) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetApplications) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetApplications) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetApplications) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetServiceId

`func (o *GetApplications) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *GetApplications) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *GetApplications) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *GetApplications) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetTechnologies

`func (o *GetApplications) GetTechnologies() []string`

GetTechnologies returns the Technologies field if non-nil, zero value otherwise.

### GetTechnologiesOk

`func (o *GetApplications) GetTechnologiesOk() (*[]string, bool)`

GetTechnologiesOk returns a tuple with the Technologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologies

`func (o *GetApplications) SetTechnologies(v []string)`

SetTechnologies sets Technologies field to given value.

### HasTechnologies

`func (o *GetApplications) HasTechnologies() bool`

HasTechnologies returns a boolean if a field has been set.

### GetTimeFrame

`func (o *GetApplications) GetTimeFrame() TimeFrame`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *GetApplications) GetTimeFrameOk() (*TimeFrame, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *GetApplications) SetTimeFrame(v TimeFrame)`

SetTimeFrame sets TimeFrame field to given value.

### HasTimeFrame

`func (o *GetApplications) HasTimeFrame() bool`

HasTimeFrame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


