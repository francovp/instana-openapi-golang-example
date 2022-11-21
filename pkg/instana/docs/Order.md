# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | **string** |  | 
**Collation** | Pointer to **string** |  | [optional] 
**Direction** | **string** |  | 

## Methods

### NewOrder

`func NewOrder(by string, direction string, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *Order) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *Order) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *Order) SetBy(v string)`

SetBy sets By field to given value.


### GetCollation

`func (o *Order) GetCollation() string`

GetCollation returns the Collation field if non-nil, zero value otherwise.

### GetCollationOk

`func (o *Order) GetCollationOk() (*string, bool)`

GetCollationOk returns a tuple with the Collation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollation

`func (o *Order) SetCollation(v string)`

SetCollation sets Collation field to given value.

### HasCollation

`func (o *Order) HasCollation() bool`

HasCollation returns a boolean if a field has been set.

### GetDirection

`func (o *Order) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Order) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Order) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


