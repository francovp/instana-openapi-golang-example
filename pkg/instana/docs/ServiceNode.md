# ServiceNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | [**map[string]EndpointNode**](EndpointNode.md) |  | 
**Inclusive** | Pointer to **bool** |  | [optional] 
**ServiceId** | **string** |  | 

## Methods

### NewServiceNode

`func NewServiceNode(endpoints map[string]EndpointNode, serviceId string, ) *ServiceNode`

NewServiceNode instantiates a new ServiceNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNodeWithDefaults

`func NewServiceNodeWithDefaults() *ServiceNode`

NewServiceNodeWithDefaults instantiates a new ServiceNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *ServiceNode) GetEndpoints() map[string]EndpointNode`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ServiceNode) GetEndpointsOk() (*map[string]EndpointNode, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ServiceNode) SetEndpoints(v map[string]EndpointNode)`

SetEndpoints sets Endpoints field to given value.


### GetInclusive

`func (o *ServiceNode) GetInclusive() bool`

GetInclusive returns the Inclusive field if non-nil, zero value otherwise.

### GetInclusiveOk

`func (o *ServiceNode) GetInclusiveOk() (*bool, bool)`

GetInclusiveOk returns a tuple with the Inclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusive

`func (o *ServiceNode) SetInclusive(v bool)`

SetInclusive sets Inclusive field to given value.

### HasInclusive

`func (o *ServiceNode) HasInclusive() bool`

HasInclusive returns a boolean if a field has been set.

### GetServiceId

`func (o *ServiceNode) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceNode) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceNode) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


