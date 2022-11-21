# ApplicationNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** |  | 
**Inclusive** | Pointer to **bool** |  | [optional] 
**Services** | [**map[string]ServiceNode**](ServiceNode.md) |  | 

## Methods

### NewApplicationNode

`func NewApplicationNode(applicationId string, services map[string]ServiceNode, ) *ApplicationNode`

NewApplicationNode instantiates a new ApplicationNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationNodeWithDefaults

`func NewApplicationNodeWithDefaults() *ApplicationNode`

NewApplicationNodeWithDefaults instantiates a new ApplicationNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApplicationNode) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationNode) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationNode) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetInclusive

`func (o *ApplicationNode) GetInclusive() bool`

GetInclusive returns the Inclusive field if non-nil, zero value otherwise.

### GetInclusiveOk

`func (o *ApplicationNode) GetInclusiveOk() (*bool, bool)`

GetInclusiveOk returns a tuple with the Inclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusive

`func (o *ApplicationNode) SetInclusive(v bool)`

SetInclusive sets Inclusive field to given value.

### HasInclusive

`func (o *ApplicationNode) HasInclusive() bool`

HasInclusive returns a boolean if a field has been set.

### GetServices

`func (o *ApplicationNode) GetServices() map[string]ServiceNode`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ApplicationNode) GetServicesOk() (*map[string]ServiceNode, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ApplicationNode) SetServices(v map[string]ServiceNode)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


