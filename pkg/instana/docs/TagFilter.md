# TagFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BooleanValue** | Pointer to **bool** |  | [optional] 
**Entity** | **string** |  | 
**Key** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**NumberValue** | Pointer to **int64** |  | [optional] 
**Operator** | **string** |  | 
**StringValue** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewTagFilter

`func NewTagFilter(entity string, name string, operator string, ) *TagFilter`

NewTagFilter instantiates a new TagFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagFilterWithDefaults

`func NewTagFilterWithDefaults() *TagFilter`

NewTagFilterWithDefaults instantiates a new TagFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBooleanValue

`func (o *TagFilter) GetBooleanValue() bool`

GetBooleanValue returns the BooleanValue field if non-nil, zero value otherwise.

### GetBooleanValueOk

`func (o *TagFilter) GetBooleanValueOk() (*bool, bool)`

GetBooleanValueOk returns a tuple with the BooleanValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooleanValue

`func (o *TagFilter) SetBooleanValue(v bool)`

SetBooleanValue sets BooleanValue field to given value.

### HasBooleanValue

`func (o *TagFilter) HasBooleanValue() bool`

HasBooleanValue returns a boolean if a field has been set.

### GetEntity

`func (o *TagFilter) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *TagFilter) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *TagFilter) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetKey

`func (o *TagFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TagFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TagFilter) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TagFilter) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *TagFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagFilter) SetName(v string)`

SetName sets Name field to given value.


### GetNumberValue

`func (o *TagFilter) GetNumberValue() int64`

GetNumberValue returns the NumberValue field if non-nil, zero value otherwise.

### GetNumberValueOk

`func (o *TagFilter) GetNumberValueOk() (*int64, bool)`

GetNumberValueOk returns a tuple with the NumberValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberValue

`func (o *TagFilter) SetNumberValue(v int64)`

SetNumberValue sets NumberValue field to given value.

### HasNumberValue

`func (o *TagFilter) HasNumberValue() bool`

HasNumberValue returns a boolean if a field has been set.

### GetOperator

`func (o *TagFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *TagFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *TagFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetStringValue

`func (o *TagFilter) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *TagFilter) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *TagFilter) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *TagFilter) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetValue

`func (o *TagFilter) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TagFilter) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TagFilter) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *TagFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


