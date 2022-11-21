# AvailabilitySliEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**BadEventFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**BadEventFilters** | Pointer to [**[]TagFilter**](TagFilter.md) |  | [optional] 
**BoundaryScope** | **string** |  | 
**EndpointId** | Pointer to **string** |  | [optional] 
**GoodEventFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**GoodEventFilters** | Pointer to [**[]TagFilter**](TagFilter.md) |  | [optional] 
**IncludeInternal** | Pointer to **bool** |  | [optional] 
**IncludeSynthetic** | Pointer to **bool** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 

## Methods

### NewAvailabilitySliEntity

`func NewAvailabilitySliEntity(boundaryScope string, ) *AvailabilitySliEntity`

NewAvailabilitySliEntity instantiates a new AvailabilitySliEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilitySliEntityWithDefaults

`func NewAvailabilitySliEntityWithDefaults() *AvailabilitySliEntity`

NewAvailabilitySliEntityWithDefaults instantiates a new AvailabilitySliEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *AvailabilitySliEntity) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AvailabilitySliEntity) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AvailabilitySliEntity) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AvailabilitySliEntity) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetBadEventFilterExpression

`func (o *AvailabilitySliEntity) GetBadEventFilterExpression() TagFilterExpressionElement`

GetBadEventFilterExpression returns the BadEventFilterExpression field if non-nil, zero value otherwise.

### GetBadEventFilterExpressionOk

`func (o *AvailabilitySliEntity) GetBadEventFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetBadEventFilterExpressionOk returns a tuple with the BadEventFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadEventFilterExpression

`func (o *AvailabilitySliEntity) SetBadEventFilterExpression(v TagFilterExpressionElement)`

SetBadEventFilterExpression sets BadEventFilterExpression field to given value.

### HasBadEventFilterExpression

`func (o *AvailabilitySliEntity) HasBadEventFilterExpression() bool`

HasBadEventFilterExpression returns a boolean if a field has been set.

### GetBadEventFilters

`func (o *AvailabilitySliEntity) GetBadEventFilters() []TagFilter`

GetBadEventFilters returns the BadEventFilters field if non-nil, zero value otherwise.

### GetBadEventFiltersOk

`func (o *AvailabilitySliEntity) GetBadEventFiltersOk() (*[]TagFilter, bool)`

GetBadEventFiltersOk returns a tuple with the BadEventFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadEventFilters

`func (o *AvailabilitySliEntity) SetBadEventFilters(v []TagFilter)`

SetBadEventFilters sets BadEventFilters field to given value.

### HasBadEventFilters

`func (o *AvailabilitySliEntity) HasBadEventFilters() bool`

HasBadEventFilters returns a boolean if a field has been set.

### GetBoundaryScope

`func (o *AvailabilitySliEntity) GetBoundaryScope() string`

GetBoundaryScope returns the BoundaryScope field if non-nil, zero value otherwise.

### GetBoundaryScopeOk

`func (o *AvailabilitySliEntity) GetBoundaryScopeOk() (*string, bool)`

GetBoundaryScopeOk returns a tuple with the BoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaryScope

`func (o *AvailabilitySliEntity) SetBoundaryScope(v string)`

SetBoundaryScope sets BoundaryScope field to given value.


### GetEndpointId

`func (o *AvailabilitySliEntity) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *AvailabilitySliEntity) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *AvailabilitySliEntity) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *AvailabilitySliEntity) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetGoodEventFilterExpression

`func (o *AvailabilitySliEntity) GetGoodEventFilterExpression() TagFilterExpressionElement`

GetGoodEventFilterExpression returns the GoodEventFilterExpression field if non-nil, zero value otherwise.

### GetGoodEventFilterExpressionOk

`func (o *AvailabilitySliEntity) GetGoodEventFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetGoodEventFilterExpressionOk returns a tuple with the GoodEventFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodEventFilterExpression

`func (o *AvailabilitySliEntity) SetGoodEventFilterExpression(v TagFilterExpressionElement)`

SetGoodEventFilterExpression sets GoodEventFilterExpression field to given value.

### HasGoodEventFilterExpression

`func (o *AvailabilitySliEntity) HasGoodEventFilterExpression() bool`

HasGoodEventFilterExpression returns a boolean if a field has been set.

### GetGoodEventFilters

`func (o *AvailabilitySliEntity) GetGoodEventFilters() []TagFilter`

GetGoodEventFilters returns the GoodEventFilters field if non-nil, zero value otherwise.

### GetGoodEventFiltersOk

`func (o *AvailabilitySliEntity) GetGoodEventFiltersOk() (*[]TagFilter, bool)`

GetGoodEventFiltersOk returns a tuple with the GoodEventFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodEventFilters

`func (o *AvailabilitySliEntity) SetGoodEventFilters(v []TagFilter)`

SetGoodEventFilters sets GoodEventFilters field to given value.

### HasGoodEventFilters

`func (o *AvailabilitySliEntity) HasGoodEventFilters() bool`

HasGoodEventFilters returns a boolean if a field has been set.

### GetIncludeInternal

`func (o *AvailabilitySliEntity) GetIncludeInternal() bool`

GetIncludeInternal returns the IncludeInternal field if non-nil, zero value otherwise.

### GetIncludeInternalOk

`func (o *AvailabilitySliEntity) GetIncludeInternalOk() (*bool, bool)`

GetIncludeInternalOk returns a tuple with the IncludeInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInternal

`func (o *AvailabilitySliEntity) SetIncludeInternal(v bool)`

SetIncludeInternal sets IncludeInternal field to given value.

### HasIncludeInternal

`func (o *AvailabilitySliEntity) HasIncludeInternal() bool`

HasIncludeInternal returns a boolean if a field has been set.

### GetIncludeSynthetic

`func (o *AvailabilitySliEntity) GetIncludeSynthetic() bool`

GetIncludeSynthetic returns the IncludeSynthetic field if non-nil, zero value otherwise.

### GetIncludeSyntheticOk

`func (o *AvailabilitySliEntity) GetIncludeSyntheticOk() (*bool, bool)`

GetIncludeSyntheticOk returns a tuple with the IncludeSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSynthetic

`func (o *AvailabilitySliEntity) SetIncludeSynthetic(v bool)`

SetIncludeSynthetic sets IncludeSynthetic field to given value.

### HasIncludeSynthetic

`func (o *AvailabilitySliEntity) HasIncludeSynthetic() bool`

HasIncludeSynthetic returns a boolean if a field has been set.

### GetServiceId

`func (o *AvailabilitySliEntity) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AvailabilitySliEntity) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AvailabilitySliEntity) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *AvailabilitySliEntity) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


