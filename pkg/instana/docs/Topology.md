# Topology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Edges** | Pointer to [**[]GraphEdge**](GraphEdge.md) |  | [optional] 
**Nodes** | Pointer to [**[]GraphNode**](GraphNode.md) |  | [optional] 

## Methods

### NewTopology

`func NewTopology() *Topology`

NewTopology instantiates a new Topology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyWithDefaults

`func NewTopologyWithDefaults() *Topology`

NewTopologyWithDefaults instantiates a new Topology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdges

`func (o *Topology) GetEdges() []GraphEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *Topology) GetEdgesOk() (*[]GraphEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *Topology) SetEdges(v []GraphEdge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *Topology) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetNodes

`func (o *Topology) GetNodes() []GraphNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Topology) GetNodesOk() (*[]GraphNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Topology) SetNodes(v []GraphNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Topology) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


