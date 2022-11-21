# AvailabilitySliEntityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**BadEventFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**BadEventFilters** | Pointer to [**[]TagFilter**](TagFilter.md) |  | [optional] 
**BoundaryScope** | Pointer to **string** |  | [optional] 
**EndpointId** | Pointer to **string** |  | [optional] 
**GoodEventFilterExpression** | Pointer to [**TagFilterExpressionElement**](TagFilterExpressionElement.md) |  | [optional] 
**GoodEventFilters** | Pointer to [**[]TagFilter**](TagFilter.md) |  | [optional] 
**IncludeInternal** | Pointer to **bool** |  | [optional] 
**IncludeSynthetic** | Pointer to **bool** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 

## Methods

### NewAvailabilitySliEntityAllOf

`func NewAvailabilitySliEntityAllOf() *AvailabilitySliEntityAllOf`

NewAvailabilitySliEntityAllOf instantiates a new AvailabilitySliEntityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilitySliEntityAllOfWithDefaults

`func NewAvailabilitySliEntityAllOfWithDefaults() *AvailabilitySliEntityAllOf`

NewAvailabilitySliEntityAllOfWithDefaults instantiates a new AvailabilitySliEntityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *AvailabilitySliEntityAllOf) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AvailabilitySliEntityAllOf) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AvailabilitySliEntityAllOf) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AvailabilitySliEntityAllOf) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetBadEventFilterExpression

`func (o *AvailabilitySliEntityAllOf) GetBadEventFilterExpression() TagFilterExpressionElement`

GetBadEventFilterExpression returns the BadEventFilterExpression field if non-nil, zero value otherwise.

### GetBadEventFilterExpressionOk

`func (o *AvailabilitySliEntityAllOf) GetBadEventFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetBadEventFilterExpressionOk returns a tuple with the BadEventFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadEventFilterExpression

`func (o *AvailabilitySliEntityAllOf) SetBadEventFilterExpression(v TagFilterExpressionElement)`

SetBadEventFilterExpression sets BadEventFilterExpression field to given value.

### HasBadEventFilterExpression

`func (o *AvailabilitySliEntityAllOf) HasBadEventFilterExpression() bool`

HasBadEventFilterExpression returns a boolean if a field has been set.

### GetBadEventFilters

`func (o *AvailabilitySliEntityAllOf) GetBadEventFilters() []TagFilter`

GetBadEventFilters returns the BadEventFilters field if non-nil, zero value otherwise.

### GetBadEventFiltersOk

`func (o *AvailabilitySliEntityAllOf) GetBadEventFiltersOk() (*[]TagFilter, bool)`

GetBadEventFiltersOk returns a tuple with the BadEventFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadEventFilters

`func (o *AvailabilitySliEntityAllOf) SetBadEventFilters(v []TagFilter)`

SetBadEventFilters sets BadEventFilters field to given value.

### HasBadEventFilters

`func (o *AvailabilitySliEntityAllOf) HasBadEventFilters() bool`

HasBadEventFilters returns a boolean if a field has been set.

### GetBoundaryScope

`func (o *AvailabilitySliEntityAllOf) GetBoundaryScope() string`

GetBoundaryScope returns the BoundaryScope field if non-nil, zero value otherwise.

### GetBoundaryScopeOk

`func (o *AvailabilitySliEntityAllOf) GetBoundaryScopeOk() (*string, bool)`

GetBoundaryScopeOk returns a tuple with the BoundaryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaryScope

`func (o *AvailabilitySliEntityAllOf) SetBoundaryScope(v string)`

SetBoundaryScope sets BoundaryScope field to given value.

### HasBoundaryScope

`func (o *AvailabilitySliEntityAllOf) HasBoundaryScope() bool`

HasBoundaryScope returns a boolean if a field has been set.

### GetEndpointId

`func (o *AvailabilitySliEntityAllOf) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *AvailabilitySliEntityAllOf) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *AvailabilitySliEntityAllOf) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *AvailabilitySliEntityAllOf) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetGoodEventFilterExpression

`func (o *AvailabilitySliEntityAllOf) GetGoodEventFilterExpression() TagFilterExpressionElement`

GetGoodEventFilterExpression returns the GoodEventFilterExpression field if non-nil, zero value otherwise.

### GetGoodEventFilterExpressionOk

`func (o *AvailabilitySliEntityAllOf) GetGoodEventFilterExpressionOk() (*TagFilterExpressionElement, bool)`

GetGoodEventFilterExpressionOk returns a tuple with the GoodEventFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodEventFilterExpression

`func (o *AvailabilitySliEntityAllOf) SetGoodEventFilterExpression(v TagFilterExpressionElement)`

SetGoodEventFilterExpression sets GoodEventFilterExpression field to given value.

### HasGoodEventFilterExpression

`func (o *AvailabilitySliEntityAllOf) HasGoodEventFilterExpression() bool`

HasGoodEventFilterExpression returns a boolean if a field has been set.

### GetGoodEventFilters

`func (o *AvailabilitySliEntityAllOf) GetGoodEventFilters() []TagFilter`

GetGoodEventFilters returns the GoodEventFilters field if non-nil, zero value otherwise.

### GetGoodEventFiltersOk

`func (o *AvailabilitySliEntityAllOf) GetGoodEventFiltersOk() (*[]TagFilter, bool)`

GetGoodEventFiltersOk returns a tuple with the GoodEventFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodEventFilters

`func (o *AvailabilitySliEntityAllOf) SetGoodEventFilters(v []TagFilter)`

SetGoodEventFilters sets GoodEventFilters field to given value.

### HasGoodEventFilters

`func (o *AvailabilitySliEntityAllOf) HasGoodEventFilters() bool`

HasGoodEventFilters returns a boolean if a field has been set.

### GetIncludeInternal

`func (o *AvailabilitySliEntityAllOf) GetIncludeInternal() bool`

GetIncludeInternal returns the IncludeInternal field if non-nil, zero value otherwise.

### GetIncludeInternalOk

`func (o *AvailabilitySliEntityAllOf) GetIncludeInternalOk() (*bool, bool)`

GetIncludeInternalOk returns a tuple with the IncludeInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInternal

`func (o *AvailabilitySliEntityAllOf) SetIncludeInternal(v bool)`

SetIncludeInternal sets IncludeInternal field to given value.

### HasIncludeInternal

`func (o *AvailabilitySliEntityAllOf) HasIncludeInternal() bool`

HasIncludeInternal returns a boolean if a field has been set.

### GetIncludeSynthetic

`func (o *AvailabilitySliEntityAllOf) GetIncludeSynthetic() bool`

GetIncludeSynthetic returns the IncludeSynthetic field if non-nil, zero value otherwise.

### GetIncludeSyntheticOk

`func (o *AvailabilitySliEntityAllOf) GetIncludeSyntheticOk() (*bool, bool)`

GetIncludeSyntheticOk returns a tuple with the IncludeSynthetic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSynthetic

`func (o *AvailabilitySliEntityAllOf) SetIncludeSynthetic(v bool)`

SetIncludeSynthetic sets IncludeSynthetic field to given value.

### HasIncludeSynthetic

`func (o *AvailabilitySliEntityAllOf) HasIncludeSynthetic() bool`

HasIncludeSynthetic returns a boolean if a field has been set.

### GetServiceId

`func (o *AvailabilitySliEntityAllOf) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AvailabilitySliEntityAllOf) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AvailabilitySliEntityAllOf) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *AvailabilitySliEntityAllOf) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


