# SpanRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | [**[]Application**](Application.md) |  | 
**Endpoint** | Pointer to [**Endpoint**](Endpoint.md) |  | [optional] 
**PhysicalContext** | Pointer to [**PhysicalContext**](PhysicalContext.md) |  | [optional] 
**Service** | Pointer to [**Service**](Service.md) |  | [optional] 

## Methods

### NewSpanRelation

`func NewSpanRelation(applications []Application, ) *SpanRelation`

NewSpanRelation instantiates a new SpanRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanRelationWithDefaults

`func NewSpanRelationWithDefaults() *SpanRelation`

NewSpanRelationWithDefaults instantiates a new SpanRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *SpanRelation) GetApplications() []Application`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *SpanRelation) GetApplicationsOk() (*[]Application, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *SpanRelation) SetApplications(v []Application)`

SetApplications sets Applications field to given value.


### GetEndpoint

`func (o *SpanRelation) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *SpanRelation) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *SpanRelation) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *SpanRelation) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetPhysicalContext

`func (o *SpanRelation) GetPhysicalContext() PhysicalContext`

GetPhysicalContext returns the PhysicalContext field if non-nil, zero value otherwise.

### GetPhysicalContextOk

`func (o *SpanRelation) GetPhysicalContextOk() (*PhysicalContext, bool)`

GetPhysicalContextOk returns a tuple with the PhysicalContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalContext

`func (o *SpanRelation) SetPhysicalContext(v PhysicalContext)`

SetPhysicalContext sets PhysicalContext field to given value.

### HasPhysicalContext

`func (o *SpanRelation) HasPhysicalContext() bool`

HasPhysicalContext returns a boolean if a field has been set.

### GetService

`func (o *SpanRelation) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SpanRelation) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SpanRelation) SetService(v Service)`

SetService sets Service field to given value.

### HasService

`func (o *SpanRelation) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


