# EndpointNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | **string** |  | 
**Inclusive** | Pointer to **bool** |  | [optional] 

## Methods

### NewEndpointNode

`func NewEndpointNode(endpointId string, ) *EndpointNode`

NewEndpointNode instantiates a new EndpointNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointNodeWithDefaults

`func NewEndpointNodeWithDefaults() *EndpointNode`

NewEndpointNodeWithDefaults instantiates a new EndpointNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *EndpointNode) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *EndpointNode) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *EndpointNode) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.


### GetInclusive

`func (o *EndpointNode) GetInclusive() bool`

GetInclusive returns the Inclusive field if non-nil, zero value otherwise.

### GetInclusiveOk

`func (o *EndpointNode) GetInclusiveOk() (*bool, bool)`

GetInclusiveOk returns a tuple with the Inclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusive

`func (o *EndpointNode) SetInclusive(v bool)`

SetInclusive sets Inclusive field to given value.

### HasInclusive

`func (o *EndpointNode) HasInclusive() bool`

HasInclusive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


