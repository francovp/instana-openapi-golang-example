# GraphNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**EntityId** | Pointer to [**EntityId**](EntityId.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Plugin** | Pointer to **string** |  | [optional] 

## Methods

### NewGraphNode

`func NewGraphNode() *GraphNode`

NewGraphNode instantiates a new GraphNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphNodeWithDefaults

`func NewGraphNodeWithDefaults() *GraphNode`

NewGraphNodeWithDefaults instantiates a new GraphNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GraphNode) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GraphNode) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GraphNode) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *GraphNode) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEntityId

`func (o *GraphNode) GetEntityId() EntityId`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *GraphNode) GetEntityIdOk() (*EntityId, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *GraphNode) SetEntityId(v EntityId)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *GraphNode) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetId

`func (o *GraphNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphNode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GraphNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GraphNode) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GraphNode) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GraphNode) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GraphNode) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPlugin

`func (o *GraphNode) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *GraphNode) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *GraphNode) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *GraphNode) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


