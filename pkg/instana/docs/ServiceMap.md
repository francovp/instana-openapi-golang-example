# ServiceMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | [**[]ServiceMapConnection**](ServiceMapConnection.md) |  | 
**Services** | [**[]Service**](Service.md) |  | 

## Methods

### NewServiceMap

`func NewServiceMap(connections []ServiceMapConnection, services []Service, ) *ServiceMap`

NewServiceMap instantiates a new ServiceMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMapWithDefaults

`func NewServiceMapWithDefaults() *ServiceMap`

NewServiceMapWithDefaults instantiates a new ServiceMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *ServiceMap) GetConnections() []ServiceMapConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ServiceMap) GetConnectionsOk() (*[]ServiceMapConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ServiceMap) SetConnections(v []ServiceMapConnection)`

SetConnections sets Connections field to given value.


### GetServices

`func (o *ServiceMap) GetServices() []Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ServiceMap) GetServicesOk() (*[]Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ServiceMap) SetServices(v []Service)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


