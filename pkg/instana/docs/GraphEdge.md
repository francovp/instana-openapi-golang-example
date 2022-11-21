# GraphEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** |  | [optional] 
**Relation** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewGraphEdge

`func NewGraphEdge() *GraphEdge`

NewGraphEdge instantiates a new GraphEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphEdgeWithDefaults

`func NewGraphEdgeWithDefaults() *GraphEdge`

NewGraphEdgeWithDefaults instantiates a new GraphEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *GraphEdge) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GraphEdge) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GraphEdge) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GraphEdge) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetRelation

`func (o *GraphEdge) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphEdge) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphEdge) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *GraphEdge) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetSource

`func (o *GraphEdge) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphEdge) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphEdge) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GraphEdge) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


