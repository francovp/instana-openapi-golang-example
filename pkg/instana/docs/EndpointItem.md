# EndpointItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | [**Endpoint**](Endpoint.md) |  | 
**Metrics** | [**map[string][][]float32**](array.md) |  | 

## Methods

### NewEndpointItem

`func NewEndpointItem(endpoint Endpoint, metrics map[string][][]float32, ) *EndpointItem`

NewEndpointItem instantiates a new EndpointItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointItemWithDefaults

`func NewEndpointItemWithDefaults() *EndpointItem`

NewEndpointItemWithDefaults instantiates a new EndpointItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *EndpointItem) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *EndpointItem) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *EndpointItem) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.


### GetMetrics

`func (o *EndpointItem) GetMetrics() map[string][][]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *EndpointItem) GetMetricsOk() (*map[string][][]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *EndpointItem) SetMetrics(v map[string][][]float32)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


