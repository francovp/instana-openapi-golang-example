# ServiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**map[string][][]float32**](array.md) |  | 
**Service** | [**Service**](Service.md) |  | 

## Methods

### NewServiceItem

`func NewServiceItem(metrics map[string][][]float32, service Service, ) *ServiceItem`

NewServiceItem instantiates a new ServiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceItemWithDefaults

`func NewServiceItemWithDefaults() *ServiceItem`

NewServiceItemWithDefaults instantiates a new ServiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *ServiceItem) GetMetrics() map[string][][]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ServiceItem) GetMetricsOk() (*map[string][][]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ServiceItem) SetMetrics(v map[string][][]float32)`

SetMetrics sets Metrics field to given value.


### GetService

`func (o *ServiceItem) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceItem) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceItem) SetService(v Service)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


