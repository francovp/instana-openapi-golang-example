# ServiceScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ScopedTo** | Pointer to [**ServiceScopedTo**](ServiceScopedTo.md) |  | [optional] 

## Methods

### NewServiceScope

`func NewServiceScope(name string, ) *ServiceScope`

NewServiceScope instantiates a new ServiceScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceScopeWithDefaults

`func NewServiceScopeWithDefaults() *ServiceScope`

NewServiceScopeWithDefaults instantiates a new ServiceScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceScope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceScope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceScope) SetName(v string)`

SetName sets Name field to given value.


### GetScopedTo

`func (o *ServiceScope) GetScopedTo() ServiceScopedTo`

GetScopedTo returns the ScopedTo field if non-nil, zero value otherwise.

### GetScopedToOk

`func (o *ServiceScope) GetScopedToOk() (*ServiceScopedTo, bool)`

GetScopedToOk returns a tuple with the ScopedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedTo

`func (o *ServiceScope) SetScopedTo(v ServiceScopedTo)`

SetScopedTo sets ScopedTo field to given value.

### HasScopedTo

`func (o *ServiceScope) HasScopedTo() bool`

HasScopedTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


